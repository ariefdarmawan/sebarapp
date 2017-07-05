package sebarapp

import (
	"fmt"
	"sync"
)

type Cluster struct {
	Name string

	servers []IServer
	masters []IServer

	sclients map[string]IClient
}

func NewCluster() *Cluster {
	c := new(Cluster)
	c.servers = []IServer{}
	c.masters = []IServer{}
	return c
}

func (c *Cluster) initProp() {
	if c.servers == nil {
		c.servers = []IServer{}
	}

	if c.masters == nil {
		c.masters = []IServer{}
	}
}

func (c *Cluster) Server(id string) IServer {
	c.initProp()
	for _, v := range c.servers {
		if v.Id() == id {
			return v
		}
	}
	return nil
}

func (c *Cluster) AddServer(s IServer) error {
	id := s.Id()
	c.initProp()
	if s := c.Server(id); s != nil {
		return fmt.Errorf("Server %s already exist", id)
	}
	c.servers = append(c.servers, s)
	return nil
}

func (c *Cluster) RemoveServer(serverid string) error {
	servers := []IServer{}
	for _, v := range c.servers {
		if v.Id() != serverid {
			servers = append(servers, v)
		}
	}
	c.servers = servers
	return nil
}

type BroadcastResult struct {
	Name  string
	Error string
	Data  []byte
}

func (c *Cluster) Broadcast(target, name string, in []byte) (<-chan *BroadcastResult, *sync.WaitGroup) {
	result := make(chan *BroadcastResult)
	wg := new(sync.WaitGroup)

	go func() {
		servers := c.findServers(target)
		for _, v := range servers {
			if client, cok := c.sclients[v.Id()]; cok {
				wg.Add(1)
				go func() {
					sendout, err := client.Send(name, in)
					bres := new(BroadcastResult)
					bres.Data = sendout
					if err != nil {
						bres.Error = err.Error()
					}
					wg.Done()
					result <- bres
				}()
			}
		}
		close(result)
	}()

	return result, wg
}

func (c *Cluster) findServers(target string) []IServer {
	rets := []IServer{}
	for _, v := range c.servers {
		if _, b := v.Tag(target); b {
			rets = append(rets, v)
		}
	}
	return rets
}
