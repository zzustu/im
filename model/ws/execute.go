package ws

import "log"

type discontentFunc struct {
}

func (discontentFunc) Execute(ch *Channel, _ *Request) {
	_ = ch.conn.WriteJSON(DiscontentResponse())
	ch.manager.Discontent(ch)
}

type messageFunc struct {
}

func (messageFunc) Execute(ch *Channel, req *Request) {
	log.Printf("%s >>>> %s: %s", ch.user, req.Dest, req.Data)
	if ok := ch.manager.SendMessage(ch.user, req.Dest, req.Data); !ok {
		_ = ch.conn.WriteJSON(ErrorResponse("发送失败"))
	}
}
