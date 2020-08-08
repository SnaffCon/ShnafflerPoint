package output

import "time"

type ErrClass int

const (
	FATAL ErrClass = iota
	WARN
	SUPPRESS
)

type MsgClass int

const (
	TRACE MsgClass = iota
	DEGUB
	INFO
	IMPORTANT
)

type TriageClass int

const (
	GREEN TriageClass = iota
	YELLOW
	RED
	BLACK
)

type Message struct {
	Triaged TriageClass
	Time    time.Time
	Err     error
	ErrType ErrClass
	Msg     string
	MsgType MsgClass
	Action  func()
	Payload interface{}
}

// TODO: MOAR decorators for building a Message inline????
func New(msg string) *Message {
	return &Message{
		Triaged: 0,
		Time:    time.Now(),
		Err:     nil,
		ErrType: 0,
		Msg:     "",
		MsgType: 0,
		Action:  nil,
		Payload: nil,
	}
}

func NewErr(msg string, err error) *Message {
	return &Message{
		Triaged: 0,
		Time:    time.Now(),
		Err:     err,
		ErrType: 0,
		Msg:     msg,
		MsgType: 0,
		Action:  nil,
		Payload: nil,
	}
}

func (m *Message) ResetMsg(msg string) *Message {
	m.Msg = msg
	return m
}

func (m *Message) Do(f func()) *Message {
	m.Action = f
	return m
}

func (m *Message) Data(p interface{}) *Message {
	m.Payload = p
	return m
}

///////////////////////////
// Error types
func (m *Message) Fatal() *Message {
	if m.Err != nil {
		m.ErrType = FATAL
	}
	return m
}

func (m *Message) Warning() *Message {
	if m.Err != nil {
		m.ErrType = WARN
	}
	return m
}

func (m *Message) Suppress() *Message {
	if m.Err != nil {
		m.ErrType = SUPPRESS
	}
	return m
}

// END Error Types
///////////////////////////

///////////////////////////
// Triage Types
func (m *Message) Triage(t TriageClass) *Message {
	m.Triaged = t
	return m
}

func (m *Message) Green() *Message {
	m.Triaged = GREEN
	return m
}

func (m *Message) Yellow() *Message {
	m.Triaged = YELLOW
	return m
}

func (m *Message) Red() *Message {
	m.Triaged = RED
	return m
}

func (m *Message) Black() *Message {
	m.Triaged = BLACK
	return m
}

// END Triage Types
///////////////////////////

///////////////////////////
// Message Types
func (m *Message) Typing(t MsgClass) *Message {
	m.MsgType = t
	return m
}
func (m *Message) Trace() *Message {
	m.MsgType = TRACE
	return m
}
func (m *Message) Degub() *Message {
	m.MsgType = DEGUB
	return m
}
func (m *Message) Info() *Message {
	m.MsgType = INFO
	return m
}
func (m *Message) Force() *Message {
	m.MsgType = IMPORTANT
	return m
}

// END Message Types
///////////////////////////
