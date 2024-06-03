package request

type MoveMenu struct {
	OriginID uint   `json:"originId"`
	TargetID uint   `json:"targetId"`
	DropType string `json:"dropType"`
}
