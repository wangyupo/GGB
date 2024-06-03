package response

import "github.com/wangyupo/GGB/model/system"

type LoginResponse struct {
	User      system.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
	Role      Role           `json:"role"`
	Menu      []Menu         `json:"menu"`
}

type Role struct {
	ID          uint   `json:"id"`
	RoleName    string `json:"roleName"`
	RoleCode    string `json:"roleCode"`
	Description string `json:"description"`
	Status      uint   `json:"status"`
}

type Menu struct {
	Label    string `json:"label"`
	Path     string `json:"path"`
	Icon     string `json:"icon"`
	ParentId uint   `json:"parentId"`
	Sort     int    `json:"sort"`
	Type     uint   `json:"type"`
}
