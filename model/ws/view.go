package ws

import (
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

const (
	syserror   = "ERROR"
	connected  = "CONNECTED"
	message    = "MESSAGE"
	discontent = "DISCONTENT"
	broadcast  = "BROADCAST"
	relogin    = "RELOGIN"
)

var (
	requestExecuter = map[string]Executer{
		discontent: discontentFunc{},
		message:    messageFunc{},
	}
)

type Manager interface {
	Connected(*Channel)
	Broadcast(string)
	SendMessage(string, string, string) bool
	Discontent(*Channel)
}

type ChannelManager struct {
	cmap sync.Map
}

func NewChannelManager() *ChannelManager {
	return &ChannelManager{cmap: sync.Map{}}
}

func (cm *ChannelManager) Connected(ch *Channel) {
	if val, ok := cm.cmap.Load(ch.user); ok {
		if old, ok := val.(*Channel); ok {
			_ = old.conn.WriteJSON(ReloginResponse(ch.conn.RemoteAddr().String()))
			_ = old.conn.Close()
			cm.cmap.Delete(old.user)
		}
	}

	// 存入WebSocket Map
	cm.cmap.Store(ch.user, ch)

	// 发送欢迎消息
	_ = ch.conn.WriteJSON(ConnectedResponse())

	log.Printf("%s 连接成功.", ch.user)
}

func (cm *ChannelManager) Broadcast(data string) {

	resp := BroadcastResponse(data)
	cm.cmap.Range(func(key, value interface{}) bool {
		if ch, ok := value.(*Channel); ok {
			_ = ch.conn.WriteJSON(resp)
		}
		return true
	})
}

func (cm *ChannelManager) SendMessage(from, dest, data string) bool {
	if val, ok := cm.cmap.Load(dest); ok {
		if ch, ok := val.(*Channel); ok {
			err := ch.conn.WriteJSON(MessageResponse(from, data))
			return err == nil
		}
	}
	return false
}

func (cm *ChannelManager) Discontent(ch *Channel) {
	if val, ok := cm.cmap.Load(ch.user); ok {
		if old, ok := val.(*Channel); ok {
			if old.connAt.Equal(ch.connAt) {
				_ = old.conn.Close()
				cm.cmap.Delete(old.user)
				log.Printf("%s 断开连接", ch.user)
			}
		}
	}
}

type Executer interface {
	Execute(*Channel, *Request)
}

type Request struct {
	Type string `json:"type"`
	Dest string `json:"dest"`
	Data string `json:"data"`
}

type Response struct {
	Type string `json:"type"`
	From string `json:"from"`
	Time string `json:"time"`
	Data string `json:"data"`
}

type Channel struct {
	manager Manager
	user    string
	conn    *websocket.Conn
	connAt  time.Time
}

func NewChannel(user string, conn *websocket.Conn, manager Manager) *Channel {
	return &Channel{
		manager: manager,
		user:    user,
		conn:    conn,
		connAt:  time.Now(),
	}
}

func (ch *Channel) Read() {
	for {
		var req Request
		err := ch.conn.ReadJSON(&req)
		if err != nil {
			log.Print(err)
			if _, ok := err.(*websocket.CloseError); ok {
				log.Print("WebSocket Closed")
				break
			}
			if _, ok := err.(net.Error); ok {
				log.Print("Network Socket Closed")
				break
			}
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				log.Print("io.EOF Closed")
				break
			}
			_ = ch.conn.WriteJSON(ErrorResponse("消息格式不正确"))
			continue
		}

		executer := requestExecuter[req.Type]
		if executer == nil {
			_ = ch.conn.WriteJSON(ErrorResponse("不支持的消息类型"))
			continue
		}

		log.Print(req)
		executer.Execute(ch, &req)
	}

	ch.manager.Discontent(ch)
}

func timeNow() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func ErrorResponse(data string) *Response {
	return &Response{
		Type: syserror,
		From: "",
		Time: timeNow(),
		Data: data,
	}
}

func DiscontentResponse() *Response {
	return &Response{
		Type: discontent,
		From: "",
		Time: timeNow(),
		Data: "再见",
	}
}

func ConnectedResponse() *Response {
	return &Response{
		Type: connected,
		From: "",
		Time: timeNow(),
		Data: "连接成功",
	}
}

func BroadcastResponse(data string) *Response {
	return &Response{
		Type: broadcast,
		From: "",
		Time: timeNow(),
		Data: data,
	}
}

func MessageResponse(from, data string) *Response {
	return &Response{
		Type: message,
		From: from,
		Time: timeNow(),
		Data: data,
	}
}

func ReloginResponse(data string) *Response {
	return &Response{
		Type: relogin,
		From: "",
		Time: timeNow(),
		Data: data,
	}
}
