package connect

import (
	"modules/app"
	"net"
	"sync"

	"github.com/herb-go/util"
)

type Manager struct {
	Connects  sync.Map
	listeners sync.Map
	Variables sync.Map
}

func (m *Manager) GetVariable(key string) string {
	v, ok := m.Variables.Load(key)
	if !ok {
		return ""
	}
	return v.(string)

}
func (m *Manager) SetVariable(key, value string) {
	m.Variables.Store(key, value)
}
func (m *Manager) Start(c *app.SystemConfig) {
	server, err := net.Listen("tcp", c.Addr)
	if err != nil {
		panic(err)
	}
	println("Listening " + c.Addr)
	go func() {
		m.NewIncome(server)
	}()
}
func (m *Manager) ConnectClosed(id string) {
	m.Connects.Delete(id)
}
func (m *Manager) OnError(err error) {
	util.LogError(err)
}
func (m *Manager) NewIncome(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		if conn != nil {
			m.OnConnect(conn, app.System.Server)
		}
	}
}
func (m *Manager) Close() error {
	return nil
}
func (m *Manager) OnConnect(rawconn net.Conn, main string) {
	conn := New(rawconn)
	conn.Manager = m
	conn.MainServer = main
	ok, err := conn.Start()
	if err != nil {
		go m.OnError(err)
		return
	}
	if !ok {
		return
	}
	m.Connects.Store(conn.ID, conn)
}

func (m *Manager) Broadcast(msg []byte) {
	m.Connects.Range(func(key, value interface{}) bool {
		c := value.(*Connect)
		if c != nil {
			go c.OnBoradcast(msg)
		}
		return true
	})
}

func NewManager() *Manager {
	return &Manager{}
}

var DefaultManager *Manager
