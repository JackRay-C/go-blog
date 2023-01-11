package initialize

import (
	"blog/pkg/global"
	"blog/pkg/model/po"
	"blog/pkg/utils/encrypt"
	"gorm.io/gorm/clause"
	"time"
)

func InitTable() {
	err := global.DB.AutoMigrate(
		&po.User{},
		&po.Subject{},
		&po.Post{},
		&po.Comment{},
		&po.Tag{},
		&po.Role{},
		&po.File{},
		&po.PostsTags{},
		&po.UsersRoles{},
		&po.Dict{},
		&po.Permissions{},
		&po.RolesPermissions{},
		&po.Draft{},
	)
	if err != nil {
		global.Log.Fatalf("创建数据库表错误：%s", err)
	}
	global.Log.Infof("初始化数据...")
	initData()
}

func initDictData() {
	dicts := []*po.Dict{
		{ID: 1, Name: "active", Code: 1, Value: "active", Description: "激活"},
		{ID: 2, Name: "active", Code: 2, Value: "lock", Description: "锁定"},
		{ID: 3, Name: "publish", Code: 1, Value: "published", Description: "发布"},
		{ID: 4, Name: "publish", Code: 2, Value: "draft", Description: "草稿"},
		{ID: 5, Name: "visibility", Code: 1, Value: "public", Description: "公开"},
		{ID: 6, Name: "visibility", Code: 2, Value: "private", Description: "私有"},
		{ID: 7, Name: "order_by", Code: 1, Value: "created_at desc", Description: "根据创建时间倒序排列"},
		{ID: 8, Name: "order_by", Code: 2, Value: "created_at asc", Description: "根据创建时间正序排列"},
		{ID: 9, Name: "order_by", Code: 1, Value: "updated_at desc", Description: "根据更新时间倒序排列"},
		{ID: 10, Name: "order_by", Code: 2, Value: "updated_at asc", Description: "根据更新时间正序排列"},
		{ID: 11, Name: "order_by", Code: 1, Value: "published_at desc", Description: "根据发布时间倒序排列"},
		{ID: 12, Name: "order_by", Code: 2, Value: "published_at asc", Description: "根据发布时间正序排列"},
	}

	global.Log.Debugf("初始化字典表: %s", dicts)
	if err := global.DB.Model(&po.Dict{}).Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(dicts, 100).Error; err != nil {
		global.Log.Fatalf("初始化字典表失败： %s", err)
	}

	global.Log.Debugf("初始化字典表成功")
}

func initUserData() {
	u := &po.User{
		ID:       1,
		Username: "admin",
		Nickname: "管理员",
		Password: encrypt.Sha256("admin"),
		Active:   1,
		Email:    "18435175817@163.com",
		Avatar:   1,
	}
	global.Log.Debugf("初始化管理员信息: %s", u)

	if err := global.DB.Model(&po.User{}).Clauses(clause.OnConflict{DoNothing: true}).Create(u).Error; err != nil {
		global.Log.Fatalf("初始化管理员失败： %s", err)
	}
	global.Log.Debugf("初始化管理员信息成功. ")
}

func initRoleData() {
	roles := []*po.Role{
		{ID: 1, Name: "Admin", Description: "管理员"},
		{ID: 2, Name: "Editor", Description: "编辑"},
		{ID: 3, Name: "Viewer", Description: "浏览"},
	}
	global.Log.Debugf("初始化角色信息: %s", roles)

	if err := global.DB.Model(&po.Role{}).Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(roles, 1000).Error; err != nil {
		global.Log.Fatalf("初始化角色失败：%s", err)
	}
	global.Log.Debugf("初始化角色成功. ")
}

func initPermissionData() {
	if err := global.DB.Model(&po.Permissions{}).Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches([]*po.Permissions{
		{ID: 1, Name: "Add posts", ObjectType: "posts", ActionType: "add", Description: "add posts", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, Name: "Update posts", ObjectType: "posts", ActionType: "update", Description: "update posts", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 3, Name: "Delete posts", ObjectType: "posts", ActionType: "delete", Description: "delete posts", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 4, Name: "List posts", ObjectType: "posts", ActionType: "list", Description: "list posts", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 5, Name: "Read posts", ObjectType: "posts", ActionType: "read", Description: "read posts", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 6, Name: "Published posts", ObjectType: "posts", ActionType: "published", Description: "published posts", CreatedAt: time.Now(), UpdatedAt: time.Now()},


		{ID: 7, Name: "List settings", ObjectType: "settings", ActionType: "list", Description: "list settings", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 8, Name: "Add settings", ObjectType: "settings", ActionType: "add", Description: "add settings", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 9, Name: "Update settings", ObjectType: "settings", ActionType: "update", Description: "update settings", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 10, Name: "Read settings", ObjectType: "settings", ActionType: "read", Description: "read settings", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		{ID: 11, Name: "Add tags", ObjectType: "tags", ActionType: "add", Description: "add tags", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 12, Name: "List tags", ObjectType: "tags", ActionType: "list", Description: "list tags", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 13, Name: "Update tags", ObjectType: "tags", ActionType: "update", Description: "update tags", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 14, Name: "Delete tags", ObjectType: "tags", ActionType: "delete", Description: "delete tags", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 15, Name: "Read tags", ObjectType: "tags", ActionType: "read", Description: "read tags", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		{ID: 16, Name: "Add subjects", ObjectType: "subjects", ActionType: "add", Description: "add subjects", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 17, Name: "List subjects", ObjectType: "subjects", ActionType: "list", Description: "list subjects", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 18, Name: "Update subjects", ObjectType: "subjects", ActionType: "update", Description: "update subjects", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 19, Name: "Delete subjects", ObjectType: "subjects", ActionType: "delete", Description: "delete subjects", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 20, Name: "Read subjects", ObjectType: "subjects", ActionType: "read", Description: "read subjects", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		{ID: 21, Name: "Add roles", ObjectType: "roles", ActionType: "add", Description: "add roles", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 22, Name: "List roles", ObjectType: "roles", ActionType: "list", Description: "list roles", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 23, Name: "Update roles", ObjectType: "roles", ActionType: "update", Description: "update roles", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 24, Name: "Delete roles", ObjectType: "roles", ActionType: "delete", Description: "delete roles", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 25, Name: "Read roles", ObjectType: "roles", ActionType: "read", Description: "read roles", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 26, Name: "Assign role", ObjectType: "roles", ActionType: "assign", Description: "assign role", CreatedAt: time.Now(), UpdatedAt: time.Now()},


		{ID: 27, Name: "Add permission", ObjectType: "permission", ActionType: "add", Description: "add permission", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 28, Name: "List permission", ObjectType: "permission", ActionType: "list", Description: "list permission", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 29, Name: "Update permission", ObjectType: "permission", ActionType: "update", Description: "update permission", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 30, Name: "Delete permission", ObjectType: "permission", ActionType: "delete", Description: "delete permission", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 31, Name: "Read permission", ObjectType: "permission", ActionType: "read", Description: "read permission", CreatedAt: time.Now(), UpdatedAt: time.Now()},


		{ID: 32, Name: "Add menus", ObjectType: "menus", ActionType: "add", Description: "add menus", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 33, Name: "List menus", ObjectType: "menus", ActionType: "list", Description: "list menus", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 34, Name: "Update menus", ObjectType: "menus", ActionType: "update", Description: "update menus", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 35, Name: "Delete menus", ObjectType: "menus", ActionType: "delete", Description: "delete menus", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 36, Name: "Read menus", ObjectType: "menus", ActionType: "read", Description: "read menus", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		{ID: 37, Name: "Add users", ObjectType: "users", ActionType: "add", Description: "add users", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 38, Name: "List users", ObjectType: "users", ActionType: "list", Description: "list users", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 39, Name: "Update users", ObjectType: "users", ActionType: "update", Description: "update users", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 40, Name: "Delete users", ObjectType: "users", ActionType: "delete", Description: "delete users", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 41, Name: "Read users", ObjectType: "users", ActionType: "read", Description: "read users", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 42, Name: "Authorization users role", ObjectType: "users", ActionType: "authorization", Description: "authorization users role", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		{ID: 43, Name: "Add drafts", ObjectType: "drafts", ActionType: "add", Description: "add drafts", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 44, Name: "List drafts", ObjectType: "drafts", ActionType: "list", Description: "list drafts", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 45, Name: "Update drafts", ObjectType: "drafts", ActionType: "update", Description: "update drafts", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 46, Name: "Delete drafts", ObjectType: "drafts", ActionType: "delete", Description: "delete drafts", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 47, Name: "Read drafts", ObjectType: "drafts", ActionType: "read", Description: "read drafts", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		{ID: 48, Name: "Add files", ObjectType: "files", ActionType: "add", Description: "add files", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 49, Name: "List files", ObjectType: "files", ActionType: "list", Description: "list files", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 50, Name: "Update files", ObjectType: "files", ActionType: "update", Description: "update files", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 51, Name: "Delete files", ObjectType: "files", ActionType: "delete", Description: "delete files", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 52, Name: "Read files", ObjectType: "files", ActionType: "read", Description: "read files", CreatedAt: time.Now(), UpdatedAt: time.Now()},


	}, 1000).Error; err != nil {
		global.Log.Fatalf("初始化权限表失败： %s", err)
	}

	var permissions []*po.RolesPermissions
	for i := 0; i < 52; i++ {
		permissions = append(permissions, &po.RolesPermissions{RoleId: 1, PermissionId: int64(i + 1)})
	}

	permissions = append(permissions, &po.RolesPermissions{RoleId: 2, PermissionId: 1})
	permissions = append(permissions, &po.RolesPermissions{RoleId: 2, PermissionId: 2})
	permissions = append(permissions, &po.RolesPermissions{RoleId: 2, PermissionId: 3})
	permissions = append(permissions, &po.RolesPermissions{RoleId: 2, PermissionId: 4})
	permissions = append(permissions, &po.RolesPermissions{RoleId: 2, PermissionId: 5})
	permissions = append(permissions, &po.RolesPermissions{RoleId: 2, PermissionId: 6})

	permissions = append(permissions, &po.RolesPermissions{RoleId: 2, PermissionId: 11})
	permissions = append(permissions, &po.RolesPermissions{RoleId: 2, PermissionId: 12})
	permissions = append(permissions, &po.RolesPermissions{RoleId: 2, PermissionId: 13})
	permissions = append(permissions, &po.RolesPermissions{RoleId: 2, PermissionId: 14})
	permissions = append(permissions, &po.RolesPermissions{RoleId: 2, PermissionId: 15})
	permissions = append(permissions, &po.RolesPermissions{RoleId: 2, PermissionId: 16})
	permissions = append(permissions, &po.RolesPermissions{RoleId: 2, PermissionId: 17})
	permissions = append(permissions, &po.RolesPermissions{RoleId: 2, PermissionId: 18})
	permissions = append(permissions, &po.RolesPermissions{RoleId: 2, PermissionId: 19})
	permissions = append(permissions, &po.RolesPermissions{RoleId: 2, PermissionId: 20})

	permissions = append(permissions, &po.RolesPermissions{RoleId: 3, PermissionId: 4})
	permissions = append(permissions, &po.RolesPermissions{RoleId: 3, PermissionId: 5})

	global.Log.Debugf("初始化权限表信息: %s", permissions)
	if err := global.DB.Model(&po.RolesPermissions{}).Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(permissions, 1000).Error; err != nil {
		global.Log.Fatalf("初始化权限表失败： %s", err)
	}
	global.Log.Debugf("初始化权限表信息成功. ")

}

func initData() {
	initDictData()
	initRoleData()
	initPermissionData()
	initUserData()
}
