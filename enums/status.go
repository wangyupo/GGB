package enums

type Status int

const (
	Disable Status = iota
	Enable
)

func (s Status) Text() string {
	switch s {
	case Disable:
		return "禁用"
	case Enable:
		return "启用"
	default:
		return "未知状态"
	}
}
