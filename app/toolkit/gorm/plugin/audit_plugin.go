package plugin

import (
	"gorm.io/gorm"
	"reflect"
)

// AuditPlugin 审计字段插件（用于自动填充 create_by 和 update_by）
type AuditPlugin struct{}

// Name 插件名称
func (ap *AuditPlugin) Name() string {
	return "AuditPlugin"
}

// getUserID 从 context 中获取 user_id（string 类型）
func getUserID(db *gorm.DB) (string, bool) {
	v, ok := db.Statement.Context.Value("user_id").(string)
	return v, ok
}

// Initialize 注册 GORM 插件回调
func (ap *AuditPlugin) Initialize(db *gorm.DB) error {
	// 创建时设置 create_by 和 update_by
	if err := db.Callback().Create().Before("gorm:create").
		Register("audit:create", func(db *gorm.DB) {
			if userID, ok := getUserID(db); ok && db.Statement != nil && db.Statement.Schema != nil {
				if db.Statement.ReflectValue.Kind() == reflect.Struct {
					// 设置 create_by
					if field := db.Statement.Schema.LookUpField("CreateBy"); field != nil {
						if f := db.Statement.ReflectValue.FieldByName("CreateBy"); f.IsValid() && f.CanSet() {
							f.SetString(userID)
						}
					}
					// 设置 update_by
					if field := db.Statement.Schema.LookUpField("UpdateBy"); field != nil {
						if f := db.Statement.ReflectValue.FieldByName("UpdateBy"); f.IsValid() && f.CanSet() {
							f.SetString(userID)
						}
					}
				}
			}
		}); err != nil {
		return err
	}

	// 更新时设置 update_by
	if err := db.Callback().Update().Before("gorm:update").
		Register("audit:update", func(db *gorm.DB) {
			if userID, ok := getUserID(db); ok && db.Statement != nil && db.Statement.Schema != nil {
				// 设置 update_by 字段（兼容结构体更新或 map 更新）
				db.Statement.SetColumn("UpdateBy", userID)
			}
		}); err != nil {
		return err
	}

	return nil
}
