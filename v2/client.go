package v2

type BybitV2ClientInterface interface {
	// AppDefinitions() v1.AppDefinitionsInterface
}

type BybitV2Client struct {
	// appDefinition v1.AppDefinitionsInterface
}

//func (c *BybitV2Client) AppDefinitions() v1.AppDefinitionsInterface {
//	return c.appDefinition
//}

func New(url, apiKey, apiSecret string) BybitV2ClientInterface {
	return &BybitV2Client{
		// appDefinition: v1.New(url, token),
	}
}
