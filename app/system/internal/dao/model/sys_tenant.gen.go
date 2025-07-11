// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysTenant = "sys_tenant"

// SysTenant 租户表
type SysTenant struct {
	ID              string    `gorm:"column:id;primaryKey;comment:id" json:"id"`                                // id
	TenantID        string    `gorm:"column:tenant_id;not null;comment:租户编号" json:"tenant_id"`                  // 租户编号
	ContactUserName string    `gorm:"column:contact_user_name;comment:联系人" json:"contact_user_name"`            // 联系人
	ContactPhone    string    `gorm:"column:contact_phone;comment:联系电话" json:"contact_phone"`                   // 联系电话
	CompanyName     string    `gorm:"column:company_name;comment:企业名称" json:"company_name"`                     // 企业名称
	LicenseNumber   string    `gorm:"column:license_number;comment:统一社会信用代码" json:"license_number"`             // 统一社会信用代码
	Address         string    `gorm:"column:address;comment:地址" json:"address"`                                 // 地址
	Intro           string    `gorm:"column:intro;comment:企业简介" json:"intro"`                                   // 企业简介
	Domain          string    `gorm:"column:domain;comment:域名" json:"domain"`                                   // 域名
	Remark          string    `gorm:"column:remark;comment:备注" json:"remark"`                                   // 备注
	PackageID       string    `gorm:"column:package_id;comment:套餐ID" json:"package_id"`                         // 套餐ID
	ExpireTime      time.Time `gorm:"column:expire_time;comment:过期时间" json:"expire_time"`                       // 过期时间
	AccountCount    int32     `gorm:"column:account_count;default:-1;comment:用户数量（-1不限制）" json:"account_count"` // 用户数量（-1不限制）
	Status          string    `gorm:"column:status;default:0;comment:租户状态（0正常 1停用）" json:"status"`              // 租户状态（0正常 1停用）
	DelFlag         string    `gorm:"column:del_flag;default:0;comment:删除标志（0代表存在 1代表删除）" json:"del_flag"`      // 删除标志（0代表存在 1代表删除）
	CreateDept      int64     `gorm:"column:create_dept;comment:创建部门" json:"create_dept"`                       // 创建部门
	CreateBy        int64     `gorm:"column:create_by;comment:创建者" json:"create_by"`                            // 创建者
	CreateTime      time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`                       // 创建时间
	UpdateBy        int64     `gorm:"column:update_by;comment:更新者" json:"update_by"`                            // 更新者
	UpdateTime      time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                       // 更新时间
}

// TableName SysTenant's table name
func (*SysTenant) TableName() string {
	return TableNameSysTenant
}
