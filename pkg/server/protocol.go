package server


// Server commands
const (
	CMD_ReceiveSessionId = "ReceiveSessionId"
)


type ClientMsg struct {
	Command string
	SessionId string
	AuthToken string
	Data string
}

type SyncDataMsg struct {
	CollectionName string
	DataKey string
	DataValue string
	SessionId string
	AuthToken string
}

type ServerMsg struct {
	Command string
	Data string
	SessionId string
}

type TelemetryData struct {
	Type string
	Data string

}