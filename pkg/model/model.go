package model



type ClientMsg struct {
	Type string
	SessionId string
	AuthToken string
	Data string
}

type ServerMsg struct {
	Command string
	Data string
	SessionId string
}