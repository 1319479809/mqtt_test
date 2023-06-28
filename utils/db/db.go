package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/1319479809/mqtt_test/utils"
	"github.com/1319479809/mqtt_test/utils/slog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	sqlTag = "gorm"
)

const (
	IsEnable  = 1
	IsDisable = 2
)

const (
	IsNotDelete = 0
	IsDelete    = 1
)

var (
	DBERROR      = errors.New("数据库错误")
	DTNOFNDERROR = errors.New("数据不存在")
)

type DB struct {
	*gorm.DB
}

const (
// RoleTypeSystem   = 99
// RoleTypeWeb      = 1
// RoleTypeMember   = 2
// RoleTypePersonal = 3
// RoleTypeManger   = 10

)

type ReqPage struct {
	Offset int `json:"offset" form:"offset" binding:"gte=0"`
	Number int `json:"number" form:"number" binding:"required,gte=1,lte=100"`
}

type ReqPageUnRequired struct {
	Offset int `json:"offset" form:"offset" binding:"gte=0"`
	Number int `json:"number" form:"number" binding:"gte=0,lte=100"`
}

type Page struct {
	Total   int  `json:"total"`
	Offset  int  `json:"offset"`
	Number  int  `json:"number"`
	HasMore bool `json:"has_more"`
}

var (
	Gormdb *gorm.DB
)

func init() {
	dbconf, err := utils.Cfg.GetSection("db")
	if err != nil {
		slog.Error("GetSection", err)
	}
	sqlstr := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		dbconf.Key("user").String(),
		dbconf.Key("password").String(),
		dbconf.Key("host").String(),
		dbconf.Key("datebase").String(),
		dbconf.Key("sqlparam").String(),
	)

	Gormdb, err = NewDB(sqlstr)
	if err != nil {
		slog.Error("NewDB", err)
	}

}

func GetSystemOrmDb() (db *DB) {
	db = &DB{

		DB: Gormdb.WithContext(context.Background()),
	}
	return
}

func GetSystemBeginOrm() (db *DB) {
	db = &DB{}
	tx := Gormdb.WithContext(context.Background())
	tx = tx.Begin()
	db.DB = tx
	return
}

func GetBeginOrmDB(ctx context.Context) (db *DB, err error) {

	tx := Gormdb.WithContext(ctx)
	tx = tx.Begin()
	db = &DB{
		DB: tx,
	}
	return
}

func MapHas(d map[string]string, key string) bool {
	_, ok := d[key]
	return ok
}

func NewDB(sqlstr string) (db *gorm.DB, err error) {
	lConf := logger.Config{
		SlowThreshold:             100 * time.Millisecond, // 慢 SQL 阈值
		LogLevel:                  logger.Info,            // 日志级别
		IgnoreRecordNotFoundError: true,
		Colorful:                  false,
	}
	switch os.Getenv("GORM_LOG_LEVEL") {
	case "info":
		lConf.LogLevel = logger.Info
		lConf.Colorful = true
	case "warn":
		lConf.LogLevel = logger.Warn
	case "error":
		lConf.LogLevel = logger.Error
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		lConf,
	)

	db, err = gorm.Open(mysql.Open(sqlstr), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return nil, err
	}
	return db, nil
}

// GetSelectSQL 获取修改sql
func GetSelectSQL(tableName string, data interface{}) (sql string) {
	typ := reflect.TypeOf(data)
	if typ.Kind() != reflect.Ptr {
		return ""
	}
	sql = "SELECT "
	var names []string
	for i := 0; i < typ.Elem().NumField(); i++ {
		name := GetFieldName(typ.Elem().Field(i), sqlTag)
		if name["COLUMN"] == "-" {
			continue
		}
		names = append(names, name["COLUMN"])
	}
	sql += strings.Join(names, ",")
	if len(names) > 0 {
		sql += " FROM " + tableName + " "
	}
	return
}

// GetSelectsSQL 获取修改sql
func GetSelectsSQL(tableName string, data interface{}) (sql string) {
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Ptr && val.Kind() != reflect.Slice {
		return ""
	}
	sInd := reflect.Indirect(val)
	typ := sInd.Type().Elem()

	sql = "SELECT "
	var names []string
	for i := 0; i < typ.NumField(); i++ {
		name := GetFieldName(typ.Field(i), sqlTag)
		if name["COLUMN"] == "-" {
			continue
		}
		names = append(names, name["COLUMN"])
	}
	sql += strings.Join(names, ",")
	if len(names) > 0 {
		sql += " FROM " + tableName + " "
	}
	return
}

// SELECTNUM 查询表NUM
func SelectCountSQL(tablename, key string) (sql string) {
	sql = "SELECT COUNT(" + key + ") FROM " + tablename
	return sql
}

// GetSelectsLIMITSQL 获取修改sql
func GetSelectsLIMITSQL(tableName string, data interface{}) (sql string) {
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Ptr && val.Kind() != reflect.Slice {
		return ""
	}
	sInd := reflect.Indirect(val)
	typ := sInd.Type().Elem()

	sql = ""
	var names []string
	for i := 0; i < typ.NumField(); i++ {
		name := GetFieldName(typ.Field(i), sqlTag)
		if name["COLUMN"] == "-" {
			continue
		}
		names = append(names, name["COLUMN"])
	}
	sql += strings.Join(names, ",")
	if len(names) > 0 {
		sql += " FROM " + tableName + " "
	}
	return
}

func ParseTagSetting(str string, sep string) map[string]string {
	settings := map[string]string{}
	names := strings.Split(str, sep)

	for i := 0; i < len(names); i++ {
		j := i
		if len(names[j]) > 0 {
			for {
				if names[j][len(names[j])-1] == '\\' {
					i++
					names[j] = names[j][0:len(names[j])-1] + sep + names[i]
					names[i] = ""
				} else {
					break
				}
			}
		}

		values := strings.Split(names[j], ":")
		k := strings.TrimSpace(strings.ToUpper(values[0]))

		if len(values) >= 2 {
			settings[k] = strings.Join(values[1:], ":")
		} else if k != "" {
			settings[k] = k
		}
	}

	return settings
}

// GetFieldName GetFieldName
func GetFieldName(filed reflect.StructField, tagName string) map[string]string {
	name := filed.Tag.Get(tagName)
	setting := ParseTagSetting(name, ";")
	if _, ok := setting["COLUMN"]; !ok {
		setting["COLUMN"] = strings.ToLower(filed.Name)
	}
	if setting["TABLE"] != "" {
		setting["COLUMN"] = setting["TABLE"] + "." + setting["COLUMN"]
	}

	return setting
}

// SpliceCondition 拼接Condition
func SpliceCondition(condition, str string) string {
	if strings.Contains(condition, "WHERE") {
		return condition + " AND " + str
	}
	return condition + " WHERE " + str
}
