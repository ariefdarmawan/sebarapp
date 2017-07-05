package test

import (
	"eaciit/sebarapp"
	"testing"

	"github.com/eaciit/toolkit"
)

var (
	s1, s2 sebarapp.IServer
	c1     sebarapp.IClient

	serverCred = &sebarapp.Credential{UserId: "user01", Password: "abcdefghjiklmn"}
	clientCred = &sebarapp.Credential{UserId: "user01", Password: "something"}
)

func TestPrepareMasterServer(t *testing.T) {
	s1 = sebarapp.NewServer()
	s1.SetConfig("port", 9001)
	if e := s1.Start(); e != nil {
		t.Errorf("Start-1 fail: %s", e.Error())
	}
}

func TestSvr2(t *testing.T) {
	s2 = sebarapp.NewServer()
	s2.SetConfig("port", 9002)
	s2.RegisterMaster(":9001", serverCred, nil)
	if e := s2.Start(); e != nil {
		t.Errorf("Start-2 fail: %s", e.Error())
	}
}

func TestClient(t *testing.T) {
	c1 = sebarapp.NewClient(":9001")
	c1.SetCredential(clientCred)
	bs, e := c1.Send("serverlists", nil)
	if e != nil {
		t.Fatalf("Unable to invoke command serverlists: %s", e.Error())
	}

	ms := []toolkit.M{}
	e = toolkit.DecodeByte(bs, &ms)
	if e != nil {
		t.Fatalf("Unable to decode: %s", e.Error())
	}
	if len(ms) != 2 {
		t.Fatalf("Wrong output")
	}
}

func TestClose(t *testing.T) {
	if s1 != nil {
		s1.Stop()
	} else {
		t.Error("Server is not ready")
	}

	if c1 != nil {
		c1.Close()
	}
}
