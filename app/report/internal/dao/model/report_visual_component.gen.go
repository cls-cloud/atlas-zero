// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameReportVisualComponent = "report_visual_component"

// ReportVisualComponent 大屏可视化组件
type ReportVisualComponent struct {
	UUID         string    `gorm:"column:uuid;primaryKey;comment:主键" json:"uuid"`                                 // 主键
	CreatorID    string    `gorm:"column:creator_id;comment:创建人id" json:"creator_id"`                             // 创建人id
	Creator      string    `gorm:"column:creator;comment:创建人" json:"creator"`                                     // 创建人
	CreateTime   time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建人姓名" json:"create_time"` // 创建人姓名
	ModifierID   string    `gorm:"column:modifier_id;comment:修改人id" json:"modifier_id"`                           // 修改人id
	Modifier     string    `gorm:"column:modifier;comment:修改人" json:"modifier"`                                   // 修改人
	ModifyTime   time.Time `gorm:"column:modify_time;comment:修改时间" json:"modify_time"`                            // 修改时间
	CreatorOrgID int32     `gorm:"column:creator_org_id;comment:创建人组织结构id" json:"creator_org_id"`                 // 创建人组织结构id
	TenantUUID   string    `gorm:"column:tenant_uuid;comment:租户UUID" json:"tenant_uuid"`                          // 租户UUID
	Name         string    `gorm:"column:name;comment:组件名称" json:"name"`                                          // 组件名称
	Content      string    `gorm:"column:content;comment:组件内容" json:"content"`                                    // 组件内容
}

// TableName ReportVisualComponent's table name
func (*ReportVisualComponent) TableName() string {
	return TableNameReportVisualComponent
}
