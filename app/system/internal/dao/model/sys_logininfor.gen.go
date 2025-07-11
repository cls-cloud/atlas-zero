// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysLogininfor = "sys_logininfor"

// SysLogininfor 系统访问记录
type SysLogininfor struct {
	InfoID        string    `gorm:"column:info_id;primaryKey;comment:访问ID" json:"info_id"`         // 访问ID
	TenantID      string    `gorm:"column:tenant_id;default:000000;comment:租户编号" json:"tenant_id"` // 租户编号
	UserName      string    `gorm:"column:user_name;comment:用户账号" json:"user_name"`                // 用户账号
	ClientKey     string    `gorm:"column:client_key;comment:客户端" json:"client_key"`               // 客户端
	DeviceType    string    `gorm:"column:device_type;comment:设备类型" json:"device_type"`            // 设备类型
	Ipaddr        string    `gorm:"column:ipaddr;comment:登录IP地址" json:"ipaddr"`                    // 登录IP地址
	LoginLocation string    `gorm:"column:login_location;comment:登录地点" json:"login_location"`      // 登录地点
	Browser       string    `gorm:"column:browser;comment:浏览器类型" json:"browser"`                   // 浏览器类型
	Os            string    `gorm:"column:os;comment:操作系统" json:"os"`                              // 操作系统
	Status        string    `gorm:"column:status;default:0;comment:登录状态（0成功 1失败）" json:"status"`   // 登录状态（0成功 1失败）
	Msg           string    `gorm:"column:msg;comment:提示消息" json:"msg"`                            // 提示消息
	LoginTime     time.Time `gorm:"column:login_time;comment:访问时间" json:"login_time"`              // 访问时间
}

// TableName SysLogininfor's table name
func (*SysLogininfor) TableName() string {
	return TableNameSysLogininfor
}
