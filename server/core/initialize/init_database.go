package initialize

import (
	"blog/app/domain"
	"blog/app/encrypt"
	"blog/core/database/mysql"
	"blog/core/global"
	"fmt"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

/**
初始化数据库连接信息
*/
func SetupDatabase() {
	var err error
	log.Println("初始化数据库连接...")
	switch global.Setting.App.DBType {
	case "mysql":
		global.DB, err = mysql.NewEngine(global.Setting.Mysql)
	default:
		global.DB, err = mysql.NewEngine(global.Setting.Mysql)
	}
	if err != nil {
		panic(fmt.Sprintf("初始化数据库连接错误: %s\n", err))
	}
	initTable()
	initData()

	log.Println("初始化数据库连接完成.")
}

func initTable() {
	err := global.DB.AutoMigrate(
		&domain.User{},
		&domain.Subject{},
		&domain.Post{},
		&domain.Comment{},
		&domain.Tag{},
		&domain.Role{},
		&domain.File{},
		&domain.PostsTags{},
		&domain.UsersRoles{},
		&domain.Dict{},
		&domain.Permissions{},
		&domain.RolesPermissions{},
		&domain.Repository{},
		&domain.History{},
		&domain.Head{},
	)
	if err != nil {
		panic(fmt.Sprintf("初始化数据库表错误： %s\n", err))
	}

}

func initDictData() {
	d := &domain.Dict{}
	if err := d.InsertAll([]domain.Dict{
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
	}); err != nil {
		panic(err)
	}
}

func initUserData() {
	u := &domain.User{
		ID:       1,
		Username: "admin",
		Nickname: "管理员",
		Password: encrypt.Sha256("admin"),
		Active:   1,
		Email:    "18435175817@163.com",
		Avatar:   1,
	}
	if err := global.DB.Model(&domain.User{}).Clauses(clause.OnConflict{DoNothing: true}).Create(u).Error; err != nil {
		panic(err)
	}
}

func initRoleData() {
	if err := global.DB.Model(&domain.Role{}).Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches([]*domain.Role{
		{ID: 1, Name: "Admin", Description: "管理员"},
		{ID: 2, Name: "Editor", Description: "编辑"},
		{ID: 3, Name: "Viewer", Description: "浏览"},
	}, 1000).Error; err != nil {
		panic(err)
	}
}

func initPermissionData() {
	if err := global.DB.Model(&domain.Permissions{}).Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches([]*domain.Permissions{
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
	}, 1000).Error; err != nil {
		panic(err)
	}

	var permissions []*domain.RolesPermissions
	for i := 0; i < 42; i++ {
		permissions = append(permissions, &domain.RolesPermissions{RoleId: 1, PermissionId: i + 1})
	}
	permissions = append(permissions, &domain.RolesPermissions{RoleId: 2, PermissionId: 1})
	permissions = append(permissions, &domain.RolesPermissions{RoleId: 2, PermissionId: 2})
	permissions = append(permissions, &domain.RolesPermissions{RoleId: 2, PermissionId: 3})
	permissions = append(permissions, &domain.RolesPermissions{RoleId: 2, PermissionId: 4})
	permissions = append(permissions, &domain.RolesPermissions{RoleId: 2, PermissionId: 5})
	permissions = append(permissions, &domain.RolesPermissions{RoleId: 2, PermissionId: 6})

	permissions = append(permissions, &domain.RolesPermissions{RoleId: 2, PermissionId: 11})
	permissions = append(permissions, &domain.RolesPermissions{RoleId: 2, PermissionId: 12})
	permissions = append(permissions, &domain.RolesPermissions{RoleId: 2, PermissionId: 13})
	permissions = append(permissions, &domain.RolesPermissions{RoleId: 2, PermissionId: 14})
	permissions = append(permissions, &domain.RolesPermissions{RoleId: 2, PermissionId: 15})
	permissions = append(permissions, &domain.RolesPermissions{RoleId: 2, PermissionId: 16})
	permissions = append(permissions, &domain.RolesPermissions{RoleId: 2, PermissionId: 17})
	permissions = append(permissions, &domain.RolesPermissions{RoleId: 2, PermissionId: 18})
	permissions = append(permissions, &domain.RolesPermissions{RoleId: 2, PermissionId: 19})
	permissions = append(permissions, &domain.RolesPermissions{RoleId: 2, PermissionId: 20})

	permissions = append(permissions, &domain.RolesPermissions{RoleId: 3, PermissionId: 4})
	permissions = append(permissions, &domain.RolesPermissions{RoleId: 3, PermissionId: 5})

	if err := global.DB.Model(&domain.RolesPermissions{}).Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(permissions, 1000).Error; err != nil {
		panic(err)
	}

}

func initData() {
	initDictData()
	initRoleData()
	initPermissionData()
	initUserData()
}
