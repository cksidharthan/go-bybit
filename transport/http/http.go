package http

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/cksidharthan/go-bybit/transport"
)

type HTTP struct {
	BaseURL   string
	APIKey    string
	APISecret string
}

func New(url, apiKey, apiSecret string) *HTTP {
	return &HTTP{
		BaseURL:   url,
		APIKey:    apiKey,
		APISecret: apiSecret,
	}
}

func (h *HTTP) SignedPostForm(path string, params url.Values, response interface{}) (err error) {
	u, err := url.Parse(h.BaseURL)
	if err != nil {
		return err
	}
	u.Path = path

	params = h.populateSignature(params)

	resp, err := http.Post(u.String(), "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}
	return
}

func (h *HTTP) SignedRequest(ctx context.Context, method string, path string, params interface{}, response interface{}) (err error) {
	apiPath, err := url.Parse(path)
	if err != nil {
		return
	}

	queryString, err := query.Values(params)
	if err != nil {
		return
	}
	apiPath.RawQuery = queryString.Encode()

	payload, err := h.signedRequestCall(ctx, apiPath, method, nil, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(payload, &response)
	if err != nil {
		return
	}
	return
}

func (h *HTTP) UnsignedRequest(ctx context.Context, method string, path string, params interface{}, response interface{}) (err error) {
	apiPath, err := url.Parse(path)
	if err != nil {
		return
	}

	queryString, err := query.Values(params)
	if err != nil {
		return
	}
	apiPath.RawQuery = queryString.Encode()

	payload, err := h.unSignedRequestCall(ctx, apiPath, method, nil, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(payload, &response)
	if err != nil {
		return
	}
	return
}

func (h *HTTP) unSignedRequestCall(ctx context.Context, apiPath *url.URL, method string, payload []byte, headers map[string]string) ([]byte, error) {
	c := http.Client{
		Timeout: 5 * time.Second,
	}

	base, err := url.Parse(h.BaseURL)
	if err != nil {
		return nil, &transport.Error{
			ErrorMsg: err.Error(),
			HTTPCode: -1,
		}
	}
	apiURL := base.ResolveReference(apiPath)

	apiURL.RawQuery = apiURL.Query().Encode()

	// replace encodings to make it work with http.NewRequest
	urlString := strings.ReplaceAll(apiURL.String(), "%2C", ",")

	req, err := http.NewRequestWithContext(ctx, method, urlString, bytes.NewReader(payload))
	if err != nil {
		return nil, &transport.Error{
			ErrorMsg: err.Error(),
			HTTPCode: -1,
		}
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, &transport.Error{
			ErrorMsg: err.Error(),
			HTTPCode: -1,
		}
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, &transport.Error{
			ErrorMsg: err.Error(),
			HTTPCode: -1,
		}
	}

	// TODO: Handle error codes
	if resp.StatusCode > 399 {
		var bybitError transport.BybitError
		err = json.Unmarshal(respData, &bybitError)
		if err != nil {
			return nil, &transport.Error{
				ErrorMsg: err.Error(),
				HTTPCode: -1,
			}
		}
		return nil, &transport.Error{
			ErrorMsg: bybitError.RetMsg,
			HTTPCode: resp.StatusCode,
		}
	}

	return respData, nil
}

func (h *HTTP) signedRequestCall(ctx context.Context, apiPath *url.URL, method string, payload []byte, headers map[string]string) ([]byte, error) {
	// Check if the auth values are present
	if h.APIKey == "" || h.APISecret == "" {
		return nil, &transport.Error{
			ErrorMsg: "APIKey and APISecret are required",
			HTTPCode: -1,
		}
	}

	c := http.Client{
		Timeout: 5 * time.Second,
	}

	base, err := url.Parse(h.BaseURL)
	if err != nil {
		return nil, &transport.Error{
			ErrorMsg: err.Error(),
			HTTPCode: -1,
		}
	}
	apiURL := base.ResolveReference(apiPath)

	// populate with signature
	query := h.populateSignature(apiURL.Query())
	apiURL.RawQuery = query.Encode()

	// replace encodings to make it work with http.NewRequest
	urlString := strings.ReplaceAll(apiURL.String(), "%2C", ",")

	req, err := http.NewRequestWithContext(ctx, method, urlString, bytes.NewBuffer(payload))
	if err != nil {
		return nil, &transport.Error{
			ErrorMsg: err.Error(),
			HTTPCode: -1,
		}
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, &transport.Error{
			ErrorMsg: err.Error(),
			HTTPCode: -1,
		}
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, &transport.Error{
			ErrorMsg: err.Error(),
			HTTPCode: -1,
		}
	}

	// TODO: Handle error codes
	if resp.StatusCode > 399 {
		var bybitError transport.BybitError
		err = json.Unmarshal(respData, &bybitError)
		if err != nil {
			return nil, &transport.Error{
				ErrorMsg: err.Error(),
				HTTPCode: -1,
			}
		}
		return nil, &transport.Error{
			ErrorMsg: bybitError.RetMsg,
			HTTPCode: resp.StatusCode,
		}
	}

	return respData, nil
}

func (h *HTTP) populateSignature(query url.Values) url.Values {
	intNow := int(time.Now().UTC().UnixNano() / int64(time.Millisecond))
	now := strconv.Itoa(intNow)

	query.Add("api_key", h.APIKey)
	query.Add("timestamp", now)
	query.Add("sign", h.sign(query, h.APISecret))
	return query
}

func (h *HTTP) sign(query url.Values, secret string) string {
	keys := make([]string, len(query))
	i := 0
	_val := ""
	for k := range query {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		_val += k + "=" + query.Get(k) + "&"
	}
	_val = _val[0 : len(_val)-1]
	hash := hmac.New(sha256.New, []byte(secret))
	_, err := io.WriteString(hash, _val)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", hash.Sum(nil))
}
