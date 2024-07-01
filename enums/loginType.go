package enums

type LoginType int

const (
	Logout LoginType = iota
	Login
)

func (s LoginType) Text() string {
	switch s {
	case Login:
		return "登入"
	case Logout:
		return "登出"
	default:
		return "未知状态"
	}
}
