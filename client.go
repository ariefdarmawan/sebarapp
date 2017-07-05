package sebarapp

import "fmt"

type ClientConfig struct {
	Host string
	Port int
}

type IClient interface {
	Connect() error
	Close()

	SetConfig(*ClientConfig)
	Config() *ClientConfig

	SetCredential(*Credential)
	Send(string, []byte) ([]byte, error)
	Subscribe(string, []byte) (<-chan []byte, error)
}

func NewClient(host string) *ClientBase {
	c := new(ClientBase)
	return c
}

type ClientBase struct {
	*ClientConfig
	*Credential
}

func (c *ClientBase) SetConfig(cfg *ClientConfig) {
	c.ClientConfig = cfg
}

func (c *ClientBase) Config() *ClientConfig {
	return c.ClientConfig
}

func (c *ClientBase) SetCredential(cred *Credential) {
	c.Credential = cred
}

func (c *ClientBase) RegisterMaster(host string, parms []byte) ([]byte, error) {
	return nil, fmt.Errorf("RegisterMaster is not yet implemented")
}

func (c *ClientBase) UnregisterMaster(parms []byte) error {
	return fmt.Errorf("UnregisterMaster is not yet implemented")
}

func (c *ClientBase) Close() {
}

func (c *ClientBase) Connect() error {
	return fmt.Errorf("Connect is not yet implemented")
}

func (c *ClientBase) Send(name string, in []byte) ([]byte, error) {
	return nil, fmt.Errorf("Send is not yet implemented")
}
func (c *ClientBase) Subscribe(name string, in []byte) (<-chan []byte, error) {
	return nil, fmt.Errorf("Subscribe is not yet implemented")
}
