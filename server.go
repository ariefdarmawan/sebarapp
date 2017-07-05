package sebarapp

import "strings"
import "github.com/eaciit/toolkit"

type ServerHandler func([]byte) ([]byte, error)

type IServer interface {
	Id() string
	SetId(id string)
	RegisterMaster(string, *Credential, []byte) ([]byte, error)
	UnregisterMaster() error
	SetCredential(*Credential)

	Config(string) (interface{}, bool)
	SetConfig(string, interface{})

	Start() error
	Stop()

	AddHandler(string, ServerHandler)
	Tag(string) (interface{}, bool)
	SetTag(string, interface{})
}

type ServerBase struct {
	Tags map[string]interface{}

	uId        string
	cluster    *Cluster
	handlers   map[string]ServerHandler
	log        *toolkit.LogEngine
	config     map[string]interface{}
	credential *Credential
}

func NewServer() IServer {
	s := new(ServerBase)
	s.SetId(toolkit.RandomString(32))
	return s
}

func (s *ServerBase) Id() string {
	if s.uId == "" {
		s.uId = toolkit.RandomString(32)
	}
	return s.uId
}

func (s *ServerBase) SetId(id string) {
	s.uId = id
}

func (s *ServerBase) SetCredential(cred *Credential) {
	s.credential = cred
}

func (s *ServerBase) Config(name string) (interface{}, bool) {
	s.validateInit()
	v, b := s.config[name]
	return v, b
}

func (s *ServerBase) SetConfig(name string, value interface{}) {
	s.validateInit()
	s.config[name] = value
}

func (s *ServerBase) Tag(tagid string) (interface{}, bool) {
	s.validateInit()
	tagid = strings.ToLower(tagid)
	v, b := s.Tags[tagid]
	return v, b
}

func (s *ServerBase) SetTag(tagid string, value interface{}) {
	s.validateInit()
	tagid = strings.ToLower(tagid)
	s.Tags[tagid] = value
}

func (s *ServerBase) RegisterMaster(host string, cred *Credential, in []byte) ([]byte, error) {
	return nil, nil
}
func (s *ServerBase) UnregisterMaster() error { return nil }

func (s *ServerBase) AddHandler(name string, handler ServerHandler) {
	s.validateInit()
	s.handlers[strings.ToLower(name)] = handler
}

func (s *ServerBase) validateInit() {
	if s.handlers == nil {
		s.handlers = make(map[string]ServerHandler)
	}

	if s.log == nil {
		s.log, _ = toolkit.NewLog(true, false, "", "", "")
	}

	if s.config == nil {
		s.config = map[string]interface{}{}
	}
}

func (s *ServerBase) Start() error {
	return nil
}

func (s *ServerBase) Stop() {
	if ismaster, has := s.Tag("role:master"); has && ismaster.(bool) {
		cs, wg := s.cluster.Broadcast("", "stop", nil)
		go func() {
			for c := range cs {
				s.log.Infof("%s has been stopped", c.Name)
			}
		}()

		wg.Wait()
	} else {
		s.doStop(nil)
	}
}

func (s *ServerBase) doStop(in []byte) ([]byte, error) {
	return nil, nil
}
