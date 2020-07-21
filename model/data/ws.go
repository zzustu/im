package data

import (
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type Manager interface {
	// 新连接
	Join(*Channel)

	// 发送消息
	Send(string, string, string) bool

	// 发送广播消息
	Broadcast(string)

	// 删除一个通道
	Remove(string)
}

type ChannelManager struct {
	Channels map[string]*Channel
	mutex    sync.RWMutex
}

func (m *ChannelManager) Join(ch *Channel) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	old := m.Channels[ch.Username]
	_ = old.Close()
	m.Channels[ch.Username] = ch
}

func (m *ChannelManager) Send(from, dest, message string) bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	ch := m.Channels[dest]
	if ch == nil {
		return false
	}

	return ch.Conn.WriteMessage(websocket.TextMessage, []byte(message)) == nil
}

func (m *ChannelManager) Broadcast(message string) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for _, ch := range m.Channels {
		_ = ch.Conn.WriteMessage(websocket.TextMessage, []byte(message))
	}
}

func (m *ChannelManager) Remove(username string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	ch := m.Channels[username]
	if ch != nil {
		_ = ch.Close()
		delete(m.Channels, username)
	}
}

type Channel struct {
	Username string
	ConnAt   time.Time
	Conn     *websocket.Conn
}

func NewChannel(username string, conn *websocket.Conn) *Channel {
	return &Channel{
		Username: username,
		ConnAt:   time.Now(),
		Conn:     conn,
	}
}

// 关闭WebSocket通道
func (c *Channel) Close() error {
	return c.Conn.Close()
}
