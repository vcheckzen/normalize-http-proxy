package server

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/vcheckzen/normalize-http-proxy/log"
)

type KingCardServerManager struct {
	TokenApi       string
	KingCardServer *Server
	sync.Mutex
}

type ManagerOption func(*KingCardServerManager)

func WithTokenApi(api string) ManagerOption {
	return func(m *KingCardServerManager) {
		m.TokenApi = api
	}
}

func WithKingCardServer(s *Server) ManagerOption {
	return func(m *KingCardServerManager) {
		m.KingCardServer = s
	}
}

func NewKingCardServerManager(opts ...ManagerOption) *KingCardServerManager {
	m := &KingCardServerManager{TokenApi: "http://kc.iikira.com/kingcard"}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func (m *KingCardServerManager) NewServer(address, listen string, verbose bool) *Server {
	m.KingCardServer = &Server{
		Listen:        listen,
		HttpUpstream:  NewUpstream(WithAddress(fmt.Sprintf("%s:8090", address))),
		HttpsUpstream: NewUpstream(WithAddress(fmt.Sprintf("%s:8091", address))),
		Verbose:       verbose,
	}

	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/30 * * * * ?", func() { m.updateToken() })
	c.Start()

	return m.KingCardServer
}

func (m *KingCardServerManager) checkToken() (bool, error) {
	if m.KingCardServer.HttpUpstream.Headers["Q-GUID"] == "" {
		return false, errors.New("kingcard token not initialized")
	}

	req, err := http.NewRequest("GET", "http://qq.com", nil)
	if err != nil {
		return false, err
	}

	tr := &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(fmt.Sprintf("http://%s", m.KingCardServer.Listen))
	}}
	client := &http.Client{
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 5 * time.Second,
	}
	rsp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer rsp.Body.Close()

	return strings.Contains(rsp.Header.Get("Location"), "www.qq.com"), err
}

func (m *KingCardServerManager) updateToken() error {
	m.Lock()
	defer m.Unlock()

	valid, err := m.checkToken()
	if err == nil && valid {
		log.Info("Kingcard token is still valid")
		return nil
	}

	log.Warn("Kingcard token is expired, updating")
	rsp, err := http.Get(m.TokenApi)
	if err != nil {
		log.Error("Kingcard token api failing request")
		return err
	}
	defer rsp.Body.Close()

	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Error("Kingcard token api failing data reading")
		return err
	}

	auth := string(data)
	if len(auth) < 50 || !strings.Contains(auth, ",") {
		log.Error("Kingcard token data returned from api is invalid")
		return errors.New("kingcard token data returned from api is invalid")
	}

	token := strings.Split(auth, ",")
	headers := Headers{
		"Q-GUID":  token[0],
		"Q-Token": token[1],
	}
	m.KingCardServer.HttpUpstream.Headers = headers
	m.KingCardServer.HttpsUpstream.Headers = headers
	log.Warn("Kingcard token updated: ", headers)
	return nil
}
