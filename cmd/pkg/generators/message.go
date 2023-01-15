package generators

import (
	"google.golang.org/protobuf/compiler/protogen"
)

type Message struct {
	children []Message
	m        *protogen.Message
	fields   []FieldAttribute
}

func MessageGenerator(m *protogen.Message) Message {
	msg := Message{
		m:        m,
		children: getChildren(m),
	}
	return msg
}

func (m Message) Children() []Message {
	kids := make([]Message, 0, len(m.children))
	kids = append(kids, m.children...)
	return kids
}

func getChildren(protoM *protogen.Message) []Message {
	set := make(messageSet)
	for _, m := range protoM.Messages {
		msg := MessageGenerator(m)
		set.Add(m.GoIdent.GoName, msg)
	}
	return set.Messages()
}

type messageSet map[string]Message

func (m messageSet) Add(k string, v Message) {
	if _, ok := m[k]; !ok {
		m[k] = v
	}
}

func (m messageSet) Keys() []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func (m messageSet) Clone() messageSet {
	clone := make(messageSet, len(m))
	for k, v := range m {
		clone.Add(k, v)
	}
	return clone
}

func (m messageSet) Messages() []Message {
	slice := make([]Message, 0, len(m))
	for _, v := range m {
		slice = append(slice, v)
	}
	return slice
}

func AllMessages(m Message) []Message {
	msgs := make([]Message, 0, len(m.Children()))
	for _, msg := range m.children {
		msgs = append(msgs, AllMessages(msg)...)
	}
	msgs = append(msgs, m)
	return msgs
}
