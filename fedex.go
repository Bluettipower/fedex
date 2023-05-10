package fedex

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"path"
	"strings"

	"golang.org/x/oauth2/clientcredentials"
)

const (
	baseURL    = "https://apis.fedex.com"
	sandboxURL = "https://apis-sandbox.fedex.com"
	// apiVersion = "v1"
	authURL   = "/oauth/token"
	grantType = "client_credentials"
)

type FedExClient struct {
	baseURL    string
	httpClient *http.Client
	conf       clientcredentials.Config

	Shipment *ShipmentServiceOp
}

func NewFedExClient(clientID, clientSecret string, opts ...Option) *FedExClient {

	conf := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	client := &FedExClient{
		conf:       conf,
		httpClient: conf.Client(context.TODO()),
	}

	client.Shipment = &ShipmentServiceOp{client: client}
	for _, opt := range opts {
		opt(client)
	}

	conf.TokenURL = client.baseURL + authURL
	return client
}

func (c *FedExClient) Post(path string, body, v interface{}) error {
	return c.do(http.MethodPost, path, body, v)
}

func (c *FedExClient) Get(path string, body, v interface{}) error {
	return c.do(http.MethodGet, path, body, v)
}

func (c *FedExClient) do(method, relPath string, body, v interface{}) error {
	if strings.HasPrefix(relPath, "/") {
		// make sure it's a relative path
		relPath = strings.TrimLeft(relPath, "/")
	}

	relPath = path.Join("/", relPath)
	req, err := c.newRequest(method, relPath, body)
	if err != nil {
		return err
	}

	return c.doRequest(req, v)
}

func (c *FedExClient) newRequest(method, path string, body interface{}) (*http.Request, error) {
	var js []byte = nil
	err := error(nil)
	if body != nil {
		js, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, c.baseURL+path, bytes.NewBuffer(js))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (c *FedExClient) doRequest(req *http.Request, v interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if v != nil {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&v)
	}
	return err
}

type Option func(c *FedExClient)

func WithSandbox(sandbox bool) Option {
	return func(c *FedExClient) {
		if sandbox {
			c.baseURL = sandboxURL
		} else {
			c.baseURL = baseURL
		}
	}
}
