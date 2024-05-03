package gohttp

type httpClient struct {

func New() HttpClient {
									
		client := &httpClient{}
		return client
}

	type HttpClient interface {
		Get()
		Post()
		Put()
		Patch()
		Delete()
	
}

func (c *HttpClient) Get() {}

func (c *HttpClient) Post() {}

func (c *HttpClient) Putt() {}

func (c *HttpClient) Patch() {}

func (c *HttpClient) Delete() {}

