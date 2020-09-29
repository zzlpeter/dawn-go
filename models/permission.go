package models

import "time"

type DTStruct struct {
	InsertTime 		time.Time 		`json:"insert_time" comment:"插入时间" gorm:"column:insert_time;default:null"`
	UpdateTime 		time.Time		`json:"update_time" comment:"更新时间" gorm:"column:update_time;default:null"`
}

type Role struct {
	DTStruct
	Id 			uint 		`json:"id" comment:"主键自增" gorm:"primary_key"`
	RoleId 		string		`json:"role_id" comment:"roleID"`
	RoleDesc 	string		`json:"role_desc" comment:"role描述"`
	SiteUrl 	string		`json:"site_url"`
}

func (r *Role) TableName() string {
	return "t_role"
}

type Permission struct {
	DTStruct
	Id 				uint 		`json:"id" comment:"主键自增" gorm:"primary_key"`
	RoleId 			string		`json:"role_id" comment:"roleID"`
	PermissionId 	string		`json:"permission_id"`
	PermissionDesc 	string		`json:"permission_desc"`
}

func (p *Permission) TableName() string {
	return "t_permission"
}