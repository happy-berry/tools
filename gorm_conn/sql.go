package gorm_conn

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

var gdb *gorm.DB

func Db() *gorm.DB {
	return gdb
}

type ConnParam struct {
	User        string // mysql用户名
	Pwd         string // mysql 密码
	Host        string // mysql服务器地址
	Db          string // 数据库
	IdleConn    int    // 空闲连接数
	MaxConn     int    // 最大连接数
	MaxLiftTime int64  // 最大链接时间
	Retry       bool   // mysql 重试机制
	RetryFun    func() // 方法
}

// MySqlConn mysql链接
func MySqlConn(p ConnParam) {
	fmt.Println("[mysql] 初始化 Database ")
	var err error
	// 单例模式创建数据库链接
	if gdb == nil {
		c, err := conn(p)
		if err != nil {
			fmt.Printf("[mysql] db connection error : %v \n", err)
			os.Exit(0)
		}
		gdb = c
	}
	sqlDB, err := gdb.DB()
	if err != nil {
		fmt.Printf("[mysql] ErrInvalidDB invalid db : %v \n", err)
		os.Exit(0)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(p.IdleConn)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(p.MaxConn)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Duration(p.MaxLiftTime))
	fmt.Println("[mysql] Database 初始化结束 end")
	if p.Retry {
		go retry(sqlDB, p)
	}
}

func conn(p ConnParam) (*gorm.DB, error) {
	fmt.Println("[mysql] db conn info : ", p.Host, p.User, p.Db)
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8mb4&loc=Local&timeout=10s&parseTime=true`, p.User, p.Pwd, p.Host, p.Db)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

// 每分钟去尝试Ping mysql 服务器
func retry(sqlDB *sql.DB, p ConnParam) {
	ticker := time.NewTicker(1 * time.Minute) // 每隔1秒执行一次任务
	defer ticker.Stop()                       // 程序结束前记得关闭定时器
	for range ticker.C {
		if err := sqlDB.Ping(); err != nil {
			fmt.Println("[mysql] server conn fail : ", err)
			c, err := conn(p)
			if err != nil {
				fmt.Printf("[mysql] db connection error : %v \n", err)
				p.RetryFun()
			} else {
				gdb = c
			}
		}
	}
}
