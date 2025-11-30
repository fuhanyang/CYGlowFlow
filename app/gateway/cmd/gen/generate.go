package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fuhanyang/CYGlowFlow/app/gateway/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// generate code
func main() {
	// 1. 初始化配置
	// 确保在 app/gateway 目录下运行，或者正确设置工作目录
	// 默认 GO_ENV 为 dev，方便本地开发
	if os.Getenv("GO_ENV") == "" {
		os.Setenv("GO_ENV", "dev")
	}

	// 获取配置
	config := conf.GetConf()
	dsn := config.MySQL.DSN

	fmt.Println("Connecting to database:", dsn)

	// 2. 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %w", err))
	}

	// 3. (可选) 执行 SQL 建表语句
	// 读取 biz/dal/mysql/sql/gateway.sql
	sqlFile := "biz/dal/mysql/sql/gateway.sql"
	content, err := os.ReadFile(sqlFile)
	if err != nil {
		fmt.Printf("Warning: could not read sql file %s: %v. Skipping table creation.\n", sqlFile, err)
	} else {
		fmt.Println("Executing SQL from:", sqlFile)
		sqls := string(content)
		// 简单分割 SQL 语句
		stmts := strings.Split(sqls, ";")
		for _, stmt := range stmts {
			stmt = strings.TrimSpace(stmt)
			if stmt == "" {
				continue
			}
			if err := db.Exec(stmt).Error; err != nil {
				// 忽略一些常见错误，比如表已存在 (如果是 CREATE TABLE IF NOT EXISTS 更好，但这里直接忽略错误继续也行)
				// 为了安全，这里打印警告但不中断，除非是严重错误
				fmt.Printf("Warning executing SQL: %v\nSQL: %s\n", err, stmt)
			}
		}
		fmt.Println("SQL execution completed.")
	}

	// 4. 配置生成器
	g := gen.NewGenerator(gen.Config{
		// 输出路径 (相对于当前执行目录，或者绝对路径)
		// 假设在 app/gateway 下运行 go run cmd/gen/generate.go
		OutPath: "biz/dal/mysql/query",
		// Model 输出路径 (默认是 model 目录，可以自定义)
		// 如果不指定 ModelPkgPath，默认会生成在 OutPath 同级的 model 目录中
		// 这里我们显式指定一下，保持清晰
		ModelPkgPath: "/model",

		// 生成模式
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,

		// 字段类型映射等配置...
		FieldNullable:     true,
		FieldCoverable:    true,
		FieldSignable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	// 使用数据库连接
	g.UseDB(db)

	// 5. 生成模型
	// 生成所有表的模型
	// 也可以指定表名: g.ApplyBasic(g.GenerateModel("sys_users"), ...)

	// 这里我们自动发现所有表
	// 过滤掉不需要的表（如果有）
	// allTables := []string{"sys_users", "sys_roles", "sys_permissions", "sys_user_roles", "sys_role_permissions", "gateway_access_logs", "gateway_rate_limits"}
	// 或者直接 ApplyBasic(g.GenerateAllTable()...)

	g.ApplyBasic(g.GenerateAllTable()...)

	// 6. 执行生成
	g.Execute()

	fmt.Println("Code generation completed successfully!")
}
