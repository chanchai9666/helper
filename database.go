package helper

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	Host         string //Host Database
	Port         int    //Port ที่ใช้
	UserName     string //ชื่อผู้ใช้
	Password     string //รหัสผ่าน
	DatabaseName string //ชื่อฐานข้อมูล
	DriverName   string //ชื่อของ DRIVER ของฐานข้อมูล (postgres,mysql,sqlserver)
	ConnectName  string //ชื่อการเชื่อมต่อ (กำหนดเองได้)
}

// สร้าง connect string การเชื่อมต่อของฐานข้อมูล
func DBConnectionString(dbConfig DBConfig) (dsn string, driver string) {
	// ตรวจสอบประเภทของฐานข้อมูลและสร้าง DSN ที่เหมาะสม
	switch dbConfig.DriverName {
	case "postgres":
		dsn = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.UserName,
			dbConfig.Password,
			dbConfig.DatabaseName,
		)
		driver = dbConfig.DriverName
	case "mysql":
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			dbConfig.UserName,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.DatabaseName,
		)
		driver = dbConfig.DriverName
	case "sqlserver":
		dsn = fmt.Sprintf(
			"sqlserver://%s:%s@%s:%d?database=%s",
			dbConfig.UserName,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.DatabaseName,
		)
		driver = dbConfig.DriverName
	default:
		dsn = "Invalid DriverName"
		driver = "NoConnect"
	}

	return dsn, driver
}

// เชื่อมต่อฐานข้อมูล
func DBConnects(dsn string, dbType string) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)
	fmt.Println(dsn)
	fmt.Println(dbType)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, //Slow SQL threshold
			LogLevel:      logger.Info, //Log level
			Colorful:      true,
		},
	)

	gormConfig := &gorm.Config{
		Logger:          newLogger,
		DryRun:          false,
		CreateBatchSize: 10000,
		QueryFields:     false,
	}

	switch dbType {
	case "postgres":
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN: dsn + " TimeZone=Asia/Bangkok",
			// PreferSimpleProtocol: true,
		}),
			gormConfig,
		)

	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), gormConfig)

	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dsn), gormConfig)

	case "sqlserver":
		db, err = gorm.Open(sqlserver.Open(dsn), gormConfig)

	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection: %w", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed Ping database : %w", err)
	}

	// ตั้งค่า connection pool
	sqlDB.SetMaxIdleConns(50)  //กำหนดจำนวนการเชื่อมต่อฐานข้อมูลที่ไม่ได้ใช้งานสูงสุดที่จะถูกเก็บไว้ในการเชื่อมต่อฐานข้อมูล (หมายความว่า หากกระบวนการของคุณสร้างการเชื่อมต่อฐานข้อมูลมากกว่า MaxIdleConns การเชื่อมต่อฐานข้อมูลเหล่านั้นจะถูกปิดจนกว่ากระบวนการของคุณจะต้องการใช้งานอีกครั้ง)
	sqlDB.SetMaxOpenConns(100) //กำหนดจำนวนการเชื่อมต่อฐานข้อมูลสูงสุดที่กระบวนการของคุณจะเปิดไว้พร้อมกัน
	sqlDB.SetConnMaxIdleTime(time.Minute * 10)
	sqlDB.SetConnMaxLifetime(time.Minute * 10) //กำหนดระยะเวลาสูงสุดที่การเชื่อมต่อฐานข้อมูลจะยังคงเปิดอยู่
	return db, nil
}
