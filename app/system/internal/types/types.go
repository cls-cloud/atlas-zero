// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.3

package types

type Resp struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type LoginReq struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	ClientId  string `json:"clientId"`
	TenantId  string `json:"tenantId"`
	GrantType string `json:"grantType"`
	Code      string `json:"code,optional"`
	Uuid      string `json:"uuid,optional"`
}

type LoginResp struct {
	AccessToken     string `json:"access_token"`
	RefreshToken    string `json:"refresh_token"`
	ClientId        string `json:"client_id"`
	ExpireIn        int64  `json:"expire_in"`
	RefreshExpireIn int64  `json:"refresh_expire_in"`
}

type CaptchaResp struct {
	CaptchaEnabled bool   `json:"captchaEnabled"`
	Img            string `json:"img"`
	Uuid           string `json:"uuid"`
}

type TenantResp struct {
	TenantEnabled bool       `json:"tenantEnabled"`
	VoList        []TenantVo `json:"voList"`
}

type TenantVo struct {
	TenantId    string `json:"tenantId"`
	CompanyName string `json:"companyName"`
	Domain      string `json:"domain"`
}

type AddOrUpdateUserReq struct {
	RoleIds []string `json:"roleIds"`
	PostIds []string `json:"postIds"`
	UserBase
}

type UpdateUserStatusReq struct {
	UserId string `json:"userId"`
	Status string `json:"status"`
}

type UserDetailResp struct {
	PostIds []string    `json:"postIds"`
	Posts   []*PostBase `json:"posts"`
	RoleIds []string    `json:"roleIds"`
	Roles   []*RoleBase `json:"roles"`
	User    *UserRoles  `json:"user"`
}

type UserQuery struct {
	UserId      string `form:"userId,optional"`
	TenantId    string `form:"tenantId,optional"`
	DeptId      string `form:"deptId,optional"`
	UserName    string `form:"userName,optional"`
	NickName    string `form:"nickName,optional"`
	UserType    string `form:"userType,optional"`
	Email       string `form:"email,optional"`
	PhoneNumber string `form:"phonenumber,optional"`
	Sex         string `form:"sex,optional"`
	Avatar      int64  `form:"avatar,optional"`
	Status      string `form:"status,optional"`
	DelFlag     string `form:"delFlag,optional"`
	LoginIp     string `form:"loginIp,optional"`
}

type QueryPageUserListReq struct {
	PageReq
	UserQuery
}

type QueryPageUserListResp struct {
	Total int64             `json:"total"`
	Rows  []UserAndDeptName `json:"rows"`
}

type UserAndDeptName struct {
	UserRoles
	DeptName   string `json:"deptName"`
	CreateTime string `json:"createTime"`
}

type QueryUserListReq struct {
	UserQuery
}

type UserInfoResp struct {
	User        UserRoles `json:"user"`
	Roles       []string  `json:"roles"`
	Permissions []string  `json:"permissions"`
}

type UserRoles struct {
	UserBase
	Roles []RoleBase `json:"roles,omitempty"`
}

type ResetPwdReq struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}

type UserProfileResp struct {
	User      UserAndDeptName `json:"user"`
	PostGroup string          `json:"postGroup"`
	RoleGroup string          `json:"roleGroup"`
}

type PageReq struct {
	PageNum  int64 `form:"pageNum,default=1"`
	PageSize int64 `form:"pageSize,default=20"`
	TimeReq
}

type IdReq struct {
	Id string `path:"id"`
}

type CodeReq struct {
	Code string `path:"code"`
}

type IdsReq struct {
	Id string `json:"ids"`
}

type TimeReq struct {
	BeginTime string `form:"params[beginTime],optional"`
	EndTime   string `form:"params[endTime],optional"`
}

type UserBase struct {
	UserID      string `json:"userId,optional"`
	TenantID    string `json:"tenantId,optional"`
	DeptID      string `json:"deptId,optional"`
	UserName    string `json:"userName"`
	NickName    string `json:"nickName"`
	UserType    string `json:"userType,optional"`
	Email       string `json:"email,optional"`
	PhoneNumber string `json:"phonenumber,optional"`
	Sex         string `json:"sex,optional"`
	Avatar      string `json:"avatar,optional"`
	Password    string `json:"password,optional"`
	Status      string `json:"status,optional"`
	LoginIp     string `json:"loginIp,optional"`
	LoginDate   string `json:"loginDate,optional"`
	Remark      string `json:"remark,optional"`
}

type UserRoleBase struct {
	UserID string `json:"userId"`
	RoleID string `json:"roleId"`
}

type RoleBase struct {
	RoleID            string `json:"roleId,optional"`
	TenantID          string `json:"tenantId,optional"`
	RoleName          string `json:"roleName"`
	RoleKey           string `json:"roleKey"`
	RoleSort          int32  `json:"roleSort"`
	DataScope         string `json:"dataScope,optional"`
	MenuCheckStrictly bool   `json:"menuCheckStrictly,optional"`
	DeptCheckStrictly bool   `json:"deptCheckStrictly,optional"`
	Status            string `json:"status,default=0,optional"`
	Remark            string `json:"remark,optional"`
	CreateTime        string `json:"createTime,optional"`
	SuperAdmin        bool   `json:"superAdmin,omitempty,optional"`
}

type UserPostBase struct {
	UserID string `json:"userId"`
	PostID string `json:"postId"`
}

type PostBase struct {
	PostID       string `json:"postId,optional"`
	TenantID     string `json:"tenantId,optional"`
	DeptID       string `json:"deptId"`
	PostCode     string `json:"postCode"`
	PostCategory string `json:"postCategory,optional"`
	PostName     string `json:"postName"`
	PostSort     int32  `json:"postSort"`
	Status       string `json:"status,default=0"`
	Remark       string `json:"remark,optional"`
	CreateTime   string `json:"createTime,optional"`
}

type RoleMenuBase struct {
	RoleID string `json:"roleId"`
	MenuID string `json:"menuId"`
}

type MenuBase struct {
	MenuID     string `json:"menuId,optional"`
	MenuName   string `json:"menuName"`
	ParentID   string `json:"parentId,optional"`
	OrderNum   int32  `json:"orderNum,optional"`
	Path       string `json:"path,optional"`
	Component  string `json:"component,optional"`
	QueryParam string `json:"queryParam,optional"`
	IsFrame    string `json:"isFrame,optional"`
	IsCache    string `json:"isCache,optional"`
	MenuType   string `json:"menuType,optional"`
	Visible    string `json:"visible,optional"`
	Status     string `json:"status,optional"`
	Perms      string `json:"perms,optional"`
	Icon       string `json:"icon,optional"`
	Remark     string `json:"remark,optional"`
	CreateTime string `json:"createTime,optional"`
}

type RoleDeptBase struct {
	RoleID string `json:"roleId"`
	DeptID string `json:"deptId"`
}

type DeptBase struct {
	DeptID       string      `json:"deptId,optional"`
	TenantID     string      `json:"tenantId,optional"`
	ParentID     string      `json:"parentId,optional"`
	Ancestors    string      `json:"ancestors,optional"`
	DeptName     string      `json:"deptName,optional"`
	DeptCategory string      `json:"deptCategory,optional"`
	OrderNum     int32       `json:"orderNum,optional"`
	Leader       *string     `json:"leader,optional"`
	Phone        string      `json:"phone,optional"`
	Email        string      `json:"email,optional"`
	Status       string      `json:"status,optional"`
	CreateTime   string      `json:"createTime,optional"`
	Children     []*DeptBase `json:"children,optional,omitempty"`
}

type DictDataBase struct {
	DictCode   string `json:"dictCode,optional"`
	TenantId   string `json:"tenantId,optional"`
	DictSort   int32  `json:"dictSort,optional"`
	DictLabel  string `json:"dictLabel,optional"`
	DictValue  string `json:"dictValue,optional"`
	DictType   string `json:"dictType,optional"`
	CssClass   string `json:"cssClass,optional"`
	ListClass  string `json:"listClass,optional"`
	IsDefault  string `json:"isDefault,optional"`
	CreateTime string `json:"createTime,optional"`
	Remark     string `json:"remark,optional"`
}

type DictTypeBase struct {
	DictID     string `json:"dictId,optional"`
	TenantID   string `json:"tenantId,optional"`
	DictName   string `json:"dictName,optional"`
	DictType   string `json:"dictType,optional"`
	CreateTime string `json:"createTime,optional"`
	Remark     string `json:"remark,optional"`
}

type ConfigBase struct {
	ConfigID    string `json:"configId,optional"`
	TenantID    string `json:"tenantId,optional"`
	ConfigName  string `json:"configName,optional"`
	ConfigKey   string `json:"configKey,optional"`
	ConfigValue string `json:"configValue,optional"`
	ConfigType  string `json:"configType,optional"`
	Remark      string `json:"remark,optional"`
	CreateTime  string `json:"createTime,optional"`
}

type LogininforBase struct {
	InfoID        string `json:"infoId,optional"`
	TenantID      string `json:"tenantId,optional"`
	UserName      string `json:"userName,optional"`
	ClientKey     string `json:"clientKey,optional"`
	DeviceType    string `json:"deviceType,optional"`
	Ipaddr        string `json:"ipaddr,optional"`
	LoginLocation string `json:"loginLocation,optional"`
	Browser       string `json:"browser,optional"`
	Os            string `json:"os,optional"`
	Status        string `json:"status,optional"`
	Msg           string `json:"msg,optional"`
	LoginTime     string `json:"loginTime,optional"`
}

type NoticeBase struct {
	NoticeID      string `json:"noticeId,optional"`
	TenantID      string `json:"tenantId,optional"`
	NoticeTitle   string `json:"noticeTitle"`
	NoticeType    string `json:"noticeType"`
	NoticeContent string `json:"noticeContent,optional"`
	Status        string `json:"status,optional"`
	Remark        string `json:"remark,optional"`
	CreateTime    string `json:"createTime,optional,omitempty"`
	CreateByName  string `json:"createByName,optional,omitempty"`
}

type OperLogBase struct {
	OperID        string `json:"operId,optional"`
	TenantID      string `json:"tenantId,optional"`
	Title         string `json:"title,optional"`
	BusinessType  int32  `json:"businessType,optional"`
	Method        string `json:"method,optional"`
	RequestMethod string `json:"requestMethod,optional"`
	OperatorType  int32  `json:"operatorType,optional"`
	OperName      string `json:"operName,optional"`
	DeptName      string `json:"deptName,optional"`
	OperUrl       string `json:"operUrl,optional"`
	OperIp        string `json:"operIp,optional"`
	OperLocation  string `json:"operLocation,optional"`
	OperParam     string `json:"operParam,optional"`
	JsonResult    string `json:"jsonResult,optional"`
	Status        int32  `json:"status,optional"`
	ErrorMsg      string `json:"errorMsg,optional"`
	OperTime      string `json:"operTime,optional"`
	CostTime      int64  `json:"costTime,optional"`
}

type OssBase struct {
	OssID        string `json:"ossId,optional"`
	TenantID     string `json:"tenantId,optional"`
	FileName     string `json:"fileName"`
	OriginalName string `json:"originalName"`
	FileSuffix   string `json:"fileSuffix"`
	Url          string `json:"url"`
	Ext1         string `json:"ext1,optional"`
	Service      string `json:"service"`
}

type OssConfigBase struct {
	OssConfigID  string `json:"ossConfigId,optional"`
	TenantID     string `json:"tenantId,optional"`
	ConfigKey    string `json:"configKey"`
	AccessKey    string `json:"accessKey,optional"`
	SecretKey    string `json:"secretKey,optional"`
	BucketName   string `json:"bucketName,optional"`
	Prefix       string `json:"prefix,optional"`
	Endpoint     string `json:"endpoint,optional"`
	Domain       string `json:"domain,optional"`
	IsHttps      string `json:"isHttps,optional"`
	Region       string `json:"region,optional"`
	AccessPolicy string `json:"accessPolicy"`
	Status       string `json:"status,optional"`
	Ext1         string `json:"ext1,optional"`
	Remark       string `json:"remark,optional"`
}

type TenantBase struct {
	ID              string `json:"id,optional"`
	TenantID        string `json:"tenantId,optional"`
	ContactUserName string `json:"contactUserName,optional"`
	ContactPhone    string `json:"contactPhone,optional"`
	CompanyName     string `json:"companyName,optional"`
	LicenseNumber   string `json:"licenseNumber,optional"`
	Address         string `json:"address,optional"`
	Intro           string `json:"intro,optional"`
	Domain          string `json:"domain,optional"`
	Remark          string `json:"remark,optional"`
	PackageId       string `json:"packageId,optional"`
	ExpireTime      string `json:"expireTime,optional"`
	AccountCount    int32  `json:"accountCount,optional"`
	Status          string `json:"status,optional"`
}

type TenantPackageBase struct {
	PackageID         string `json:"packageId,optional"`
	PackageName       string `json:"packageName,optional"`
	MenuIds           string `json:"menuIds,optional"`
	Remark            string `json:"remark,optional"`
	MenuCheckStrictly bool   `json:"menuCheckStrictly,optional"`
	Status            string `json:"status,optional"`
}

type ClientBase struct {
	ID            string `json:"id,optional"`
	ClientID      string `json:"clientId,optional"`
	ClientKey     string `json:"clientKey,optional"`
	ClientSecret  string `json:"clientSecret,optional"`
	GrantType     string `json:"grantType,optional"`
	DeviceType    string `json:"deviceType,optional"`
	ActiveTimeout int32  `json:"activeTimeout,optional"`
	Timeout       int32  `json:"timeout,optional"`
	Status        string `json:"status,optional"`
}

type SocialBase struct {
	ID               string `json:"id,optional"`
	UserID           string `json:"userId"`
	TenantID         string `json:"tenantId,optional"`
	AuthID           string `json:"authId"`
	Source           string `json:"source"`
	OpenID           string `json:"openId,optional"`
	UserName         string `json:"userName"`
	NickName         string `json:"nickName,optional"`
	Email            string `json:"email,optional"`
	Avatar           string `json:"avatar,optional"`
	AccessToken      string `json:"accessToken"`
	ExpireIn         int32  `json:"expireIn,optional"`
	RefreshToken     string `json:"refreshToken,optional"`
	AccessCode       string `json:"accessCode,optional"`
	UnionId          string `json:"unionId,optional"`
	Scope            string `json:"scope,optional"`
	TokenType        string `json:"tokenType,optional"`
	IdToken          string `json:"idToken,optional"`
	MacAlgorithm     string `json:"macAlgorithm,optional"`
	MacKey           string `json:"macKey,optional"`
	Code             string `json:"code,optional"`
	OauthToken       string `json:"oauthToken,optional"`
	OauthTokenSecret string `json:"oauthTokenSecret,optional"`
}

type DeptTree struct {
	Disabled bool       `json:"disabled"`
	Id       string     `json:"id"`
	Label    string     `json:"label"`
	ParentId string     `json:"parentId"`
	Weight   int64      `json:"weight"`
	Children []DeptTree `json:"children"`
}

type AddOrUpdateRoleReq struct {
	RoleBase
	MenuIds []string `json:"menuIds"`
}

type UpdateRoleStatusReq struct {
	RoleID string `json:"roleId"`
	Status string `json:"status,default=0"` //角色状态（0正常 1停用）
}

type RoleDetailResp struct {
	RoleBase
	CreateDept int64  `json:"createDept"` //创建部门
	CreateBy   int64  `json:"createBy"`   //创建者
	CreateTime string `json:"createTime"` //创建时间
	UpdateBy   int64  `json:"updateBy"`   //更新者
	UpdateTime string `json:"updateTime"` //更新时间
	Remark     string `json:"remark"`     //备注
}

type RoleQuery struct {
	RoleId            string `form:"roleId,optional"`            //角色ID
	TenantId          string `form:"tenantId,optional"`          //租户编号
	RoleName          string `form:"roleName,optional"`          //角色名称
	RoleKey           string `form:"roleKey,optional"`           //角色权限字符串
	RoleSort          int32  `form:"roleSort,optional"`          //显示顺序
	DataScope         string `form:"dataScope,optional"`         //数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限 5：仅本人数据权限 6：部门及以下或本人数据权限）
	MenuCheckStrictly int32  `form:"menuCheckStrictly,optional"` //菜单树选择项是否关联显示
	DeptCheckStrictly int32  `form:"deptCheckStrictly,optional"` //部门树选择项是否关联显示
	Status            string `form:"status,optional"`            //角色状态（0正常 1停用）
}

type RolePageSetReq struct {
	PageReq
	RoleQuery
}

type RolePageSetResp struct {
	Total int64      `json:"total"`
	Rows  []RoleBase `json:"rows"`
}

type DeptTreeResp struct {
	CheckedKeys []string   `json:"checkedKeys"`
	Depts       []DeptTree `json:"depts"`
}

type AllocatedReq struct {
	PageReq
	RoleId      string `form:"roleId"`
	UserName    string `form:"userName,optional"`
	PhoneNumber string `form:"phonenumber,optional"`
}

type AllocatedResp struct {
	Total int64      `json:"total"`
	Rows  []UserBase `json:"rows"`
}

type SelectAllReq struct {
	RoleId  string `form:"roleId"`
	UserIds string `form:"userIds"`
}

type CancelReq struct {
	RoleId string `json:"roleId"`
	UserId string `json:"userId"`
}

type ModifyDeptReq struct {
	DeptBase
}

type DeptQuery struct {
	DeptId       string `form:"deptId,optional"`       // 部门ID
	TenantId     string `form:"tenantId,optional"`     // 租户编号
	ParentId     string `form:"parentId,optional"`     // 父部门ID
	Ancestors    string `form:"ancestors,optional"`    // 祖级列表
	DeptName     string `form:"deptName,optional"`     // 部门名称
	DeptCategory string `form:"deptCategory,optional"` // 部门类别编码
	OrderNum     int32  `form:"orderNum,optional"`     // 显示顺序
	Leader       int64  `form:"leader,optional"`       // 负责人
	Phone        string `form:"phone,optional"`        // 联系电话
	Email        string `form:"email,optional"`        // 邮箱
	Status       string `form:"status,optional"`       // 部门状态（0正常 1停用）
	DelFlag      string `form:"delFlag,optional"`      // 删除标志（0代表存在 1代表删除）
	CreateDept   int64  `form:"createDept,optional"`   // 创建部门
	CreateBy     int64  `form:"createBy,optional"`     // 创建者
	CreateTime   string `form:"createTime,optional"`   // 创建时间
	UpdateBy     int64  `form:"updateBy,optional"`     // 更新者
	UpdateTime   string `form:"updateTime,optional"`   // 更新时间
}

type PageSetDeptResp struct {
	Total int64      `json:"total"`
	Rows  []DeptBase `json:"rows"`
}

type ModifyMenuReq struct {
	MenuBase
}

type MenuQuery struct {
	MenuId     string `form:"menuId,optional"`     // 菜单ID
	MenuName   string `form:"menuName,optional"`   // 菜单名称
	ParentId   string `form:"parentId,optional"`   // 父菜单ID
	OrderNum   int32  `form:"orderNum,optional"`   // 显示顺序
	Path       string `form:"path,optional"`       // 路由地址
	Component  string `form:"component,optional"`  // 组件路径
	QueryParam string `form:"queryParam,optional"` // 路由参数
	IsFrame    int32  `form:"isFrame,optional"`    // 是否为外链（0是 1否）
	IsCache    int32  `form:"isCache,optional"`    // 是否缓存（0缓存 1不缓存）
	MenuType   string `form:"menuType,optional"`   // 菜单类型（M目录 C菜单 F按钮）
	Visible    string `form:"visible,optional"`    // 显示状态（0显示 1隐藏）
	Status     string `form:"status,optional"`     // 菜单状态（0正常 1停用）
	Perms      string `form:"perms,optional"`      // 权限标识
}

type QueryMenuListReq struct {
	MenuQuery
}

type MenuDetailResp struct {
	MenuBase
	CreateDept int64  `json:"createDept"`
	CreateBy   int64  `json:"createBy"`
	CreateTime string `json:"createTime"`
	UpdateBy   int64  `json:"updateBy"`
	UpdateTime string `json:"updateTime"`
}

type RouterMenuResp struct {
	MenuId     string            `json:"menuId"`
	Name       string            `json:"name"`
	Path       string            `json:"path"`
	Hidden     bool              `json:"hidden"`
	Redirect   string            `json:"redirect,optional"` // 可选
	Component  string            `json:"component"`
	AlwaysShow bool              `json:"alwaysShow,optional"` // 可选
	Meta       *RouterMenuMeta   `json:"meta"`
	ParentId   string            `json:"parentId"`
	Children   []*RouterMenuResp `json:"children,omitempty"`
}

type RouterMenuMeta struct {
	Title   string `json:"title"`
	Icon    string `json:"icon"`
	NoCache bool   `json:"noCache"`
	Link    string `json:"link,optional"`
}

type SelectMenuTreeResp struct {
	CheckedKeys []string          `json:"checkedKeys"`
	Menus       []*SelectMenuTree `json:"menus"`
}

type SelectMenuTree struct {
	Id       string            `json:"id"`
	ParentId string            `json:"parentId"`
	MenuType string            `json:"menuType"`
	Icon     string            `json:"icon"`
	Weight   int32             `json:"weight"`
	Label    string            `json:"label"`
	Children []*SelectMenuTree `json:"children,omitempty"`
}

type ModifyPostReq struct {
	PostBase
}

type PostQuery struct {
	PostId       string `form:"postId,optional"`       // 岗位ID
	TenantId     string `form:"tenantId,optional"`     // 租户编号
	DeptId       string `form:"deptId,optional"`       // 部门ID
	PostCode     string `form:"postCode,optional"`     // 岗位编码
	PostCategory string `form:"postCategory,optional"` // 岗位类别编码
	PostName     string `form:"postName,optional"`     // 岗位名称
	PostSort     int32  `form:"postSort,optional"`     // 显示顺序
	Status       string `form:"status,optional"`       // 状态（0正常 1停用）
	CreateDept   int64  `form:"createDept,optional"`   // 创建部门
	CreateBy     int64  `form:"createBy,optional"`     // 创建者
	CreateTime   string `form:"createTime,optional"`   // 创建时间
	UpdateBy     int64  `form:"updateBy,optional"`     // 更新者
	UpdateTime   string `form:"updateTime,optional"`   // 更新时间
	Remark       string `form:"remark,optional"`       // 备注
}

type PageSetRoleReq struct {
	PageReq
	PostQuery
	BelongDeptId string `form:"belongDeptId, optional"`
}

type PageSetRoleResp struct {
	Total int64       `json:"total"`
	Rows  []*PostBase `json:"rows"`
}

type ModifyNoticeReq struct {
	NoticeBase
}

type NoticeQuery struct {
	NoticeId      string `form:"noticeId,optional"`      // 公告ID
	TenantId      string `form:"tenantId,optional"`      // 租户编号
	NoticeTitle   string `form:"noticeTitle,optional"`   // 公告标题
	NoticeType    string `form:"noticeType,optional"`    // 公告类型（1通知 2公告）
	NoticeContent string `form:"noticeContent,optional"` // 公告内容
	Status        string `form:"status,optional"`        // 公告状态（0正常 1关闭）
	CreateDept    int64  `form:"createDept,optional"`    // 创建部门
	CreateBy      string `form:"createBy,optional"`      // 创建者
	CreateTime    string `form:"createTime,optional"`    // 创建时间
	UpdateBy      int64  `form:"updateBy,optional"`      // 更新者
	UpdateTime    string `form:"updateTime,optional"`    // 更新时间
	Remark        string `form:"remark,optional"`        // 备注
}

type PageSetNoticeReq struct {
	PageReq
	NoticeQuery
}

type PageSetNoticeResp struct {
	Rows  []*NoticeBase `json:"rows"`
	Total int64         `json:"total"`
}

type ModifyClientReq struct {
	ClientBase
}

type ClientQuery struct {
	Id            string `form:"id,optional"`            // id
	ClientId      string `form:"clientId,optional"`      // 客户端id
	ClientKey     string `form:"clientKey,optional"`     // 客户端key
	ClientSecret  string `form:"clientSecret,optional"`  // 客户端秘钥
	GrantType     string `form:"grantType,optional"`     // 授权类型
	DeviceType    string `form:"deviceType,optional"`    // 设备类型
	ActiveTimeout int32  `form:"activeTimeout,optional"` // token活跃超时时间
	Timeout       int32  `form:"timeout,optional"`       // token固定超时
	Status        string `form:"status,optional"`        // 状态（0正常 1停用）
	DelFlag       string `form:"delFlag,optional"`       // 删除标志（0代表存在 1代表删除）
	CreateDept    int64  `form:"createDept,optional"`    // 创建部门
	CreateBy      int64  `form:"createBy,optional"`      // 创建者
	CreateTime    string `form:"createTime,optional"`    // 创建时间
	UpdateBy      int64  `form:"updateBy,optional"`      // 更新者
	UpdateTime    string `form:"updateTime,optional"`    // 更新时间
}

type PageSetClientReq struct {
	PageReq
	ClientQuery
}

type PageSetClientResp struct {
	Rows  []*ClientBase `json:"rows"`
	Total int64         `json:"total"`
}

type ChangeStatusClientReq struct {
	ClientID string `json:"clientId"`
	Status   string `json:"status"`
}

type ModifyTenantReq struct {
	TenantBase
	UserName string `json:"username"`
	Password string `json:"password"`
}

type TenantQuery struct {
	Id              string `form:"id,optional"`              // id
	TenantId        string `form:"tenantId,optional"`        // 租户编号
	ContactUserName string `form:"contactUserName,optional"` // 联系人
	ContactPhone    string `form:"contactPhone,optional"`    // 联系电话
	CompanyName     string `form:"companyName,optional"`     // 企业名称
	LicenseNumber   string `form:"licenseNumber,optional"`   // 统一社会信用代码
	Address         string `form:"address,optional"`         // 地址
	Intro           string `form:"intro,optional"`           // 企业简介
	Domain          string `form:"domain,optional"`          // 域名
	Remark          string `form:"remark,optional"`          // 备注
	PackageId       int64  `form:"packageId,optional"`       // 租户套餐编号
	ExpireTime      string `form:"expireTime,optional"`      // 过期时间
	AccountCount    int32  `form:"accountCount,optional"`    // 用户数量（-1不限制）
	Status          string `form:"status,optional"`          // 租户状态（0正常 1停用）
	DelFlag         string `form:"delFlag,optional"`         // 删除标志（0代表存在 1代表删除）
	CreateDept      int64  `form:"createDept,optional"`      // 创建部门
	CreateBy        int64  `form:"createBy,optional"`        // 创建者
	CreateTime      string `form:"createTime,optional"`      // 创建时间
	UpdateBy        int64  `form:"updateBy,optional"`        // 更新者
	UpdateTime      string `form:"updateTime,optional"`      // 更新时间
}

type PageSetTenantReq struct {
	PageReq
	TenantQuery
}

type PageSetTenantResp struct {
	Rows  []*TenantBase `json:"rows"`
	Total int64         `json:"total"`
}

type ModifyTenantPackageReq struct {
	PackageID         string   `json:"packageId,optional"`
	PackageName       string   `json:"packageName,optional"`
	MenuIds           []string `json:"menuIds,optional"`
	Remark            string   `json:"remark,optional"`
	MenuCheckStrictly bool     `json:"menuCheckStrictly,optional"`
}

type TenantPackageQuery struct {
	PackageId         string `form:"packageId,optional"`         // 租户套餐id
	PackageName       string `form:"packageName,optional"`       // 套餐名称
	MenuIds           string `form:"menuIds,optional"`           // 关联菜单id
	Remark            string `form:"remark,optional"`            // 备注
	MenuCheckStrictly bool   `form:"menuCheckStrictly,optional"` // 菜单树选择项是否关联显示
	Status            string `form:"status,optional"`            // 状态（0正常 1停用）
	DelFlag           string `form:"delFlag,optional"`           // 删除标志（0代表存在 1代表删除）
	CreateDept        int64  `form:"createDept,optional"`        // 创建部门
	CreateBy          int64  `form:"createBy,optional"`          // 创建者
	CreateTime        string `form:"createTime,optional"`        // 创建时间
	UpdateBy          int64  `form:"updateBy,optional"`          // 更新者
	UpdateTime        string `form:"updateTime,optional"`        // 更新时间
}

type PageSetTenantPackageReq struct {
	PageReq
	TenantPackageQuery
}

type PageSetTenantPackageResp struct {
	Rows  []*TenantPackageBase `json:"rows"`
	Total int64                `json:"total"`
}

type ModifyConfigReq struct {
	ConfigBase
}

type ConfigQuery struct {
	ConfigId    string `form:"configId,optional"`    // 参数主键
	TenantId    string `form:"tenantId,optional"`    // 租户编号
	ConfigName  string `form:"configName,optional"`  // 参数名称
	ConfigKey   string `form:"configKey,optional"`   // 参数键名
	ConfigValue string `form:"configValue,optional"` // 参数键值
	ConfigType  string `form:"configType,optional"`  // 系统内置（Y是 N否）
	CreateDept  int64  `form:"createDept,optional"`  // 创建部门
	CreateBy    int64  `form:"createBy,optional"`    // 创建者
	CreateTime  string `form:"createTime,optional"`  // 创建时间
	UpdateBy    int64  `form:"updateBy,optional"`    // 更新者
	UpdateTime  string `form:"updateTime,optional"`  // 更新时间
	Remark      string `form:"remark,optional"`      // 备注
}

type PageSetConfigReq struct {
	PageReq
	ConfigQuery
}

type PageSetConfigResp struct {
	Total int64         `json:"total"`
	Rows  []*ConfigBase `json:"rows"`
}

type ConfigKeyResp struct {
	Data string `json:"data"`
}

type ModifyDictDataReq struct {
	DictDataBase
}

type DictDataQuery struct {
	DictCode   int64  `form:"dictCode,optional"`   // 字典编码
	TenantId   string `form:"tenantId,optional"`   // 租户编号
	DictSort   int32  `form:"dictSort,optional"`   // 字典排序
	DictLabel  string `form:"dictLabel,optional"`  // 字典标签
	DictValue  string `form:"dictValue,optional"`  // 字典键值
	DictType   string `form:"dictType,optional"`   // 字典类型
	CssClass   string `form:"cssClass,optional"`   // 样式属性（其他样式扩展）
	ListClass  string `form:"listClass,optional"`  // 表格回显样式
	IsDefault  string `form:"isDefault,optional"`  // 是否默认（Y是 N否）
	CreateDept int64  `form:"createDept,optional"` // 创建部门
	CreateBy   int64  `form:"createBy,optional"`   // 创建者
	CreateTime string `form:"createTime,optional"` // 创建时间
	UpdateBy   int64  `form:"updateBy,optional"`   // 更新者
	UpdateTime string `form:"updateTime,optional"` // 更新时间
	Remark     string `form:"remark,optional"`     // 备注
}

type PageSetDictDataReq struct {
	PageReq
	DictDataQuery
}

type PageSetDictDataResp struct {
	Total int64           `json:"total"`
	Rows  []*DictDataBase `json:"rows"`
}

type ModifyDictTypeReq struct {
	DictTypeBase
}

type DictTypeQuery struct {
	DictId     string `form:"dictId,optional"`     // 字典主键
	TenantId   string `form:"tenantId,optional"`   // 租户编号
	DictName   string `form:"dictName,optional"`   // 字典名称
	DictType   string `form:"dictType,optional"`   // 字典类型
	CreateDept int64  `form:"createDept,optional"` // 创建部门
	CreateBy   int64  `form:"createBy,optional"`   // 创建者
	CreateTime string `form:"createTime,optional"` // 创建时间
	UpdateBy   int64  `form:"updateBy,optional"`   // 更新者
	UpdateTime string `form:"updateTime,optional"` // 更新时间
	Remark     string `form:"remark,optional"`     // 备注
}

type PageSetDictTypeReq struct {
	PageReq
	DictTypeQuery
}

type PageSetDictTypeResp struct {
	Total int64           `json:"total"`
	Rows  []*DictTypeBase `json:"rows"`
}
