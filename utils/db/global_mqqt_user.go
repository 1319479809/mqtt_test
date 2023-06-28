package db

import (
	"fmt"
	"time"

	"github.com/1319479809/mqtt_test/utils"
	"github.com/1319479809/mqtt_test/utils/slog"
	"gorm.io/gorm"
)

type GlobalMqttUser struct {
	BusinessId int64  `json:"business_id" gorm:"column:business_id;primaryKey"` //商户ID
	Username   string `json:"username" gorm:"column:username"`                  //使用ddy做前缀+13为随机小写字母和数字
	Password   string `json:"password" gorm:"column:password"`                  //随机生成32位大小写数字的字符串
	Remark     string `json:"remark" gorm:"column:remark"`                      //备注信息 默认**创建
	Status     int    `json:"status" gorm:"column:status"`                      //状态1:可用 2停用
	CTime      int64  `json:"c_time" gorm:"column:c_time"`                      //创建时间
}

// 通过用户查找数据
func (db *DB) GetUsernameMqttUserData(username string) (data *GlobalMqttUser, err error) {
	data = &GlobalMqttUser{}
	err = db.Where("username = ?", username).First(data).Error
	if err != nil {
		return data, err
	}
	return data, err
}

// 查找创建 BusinessId
func (db *DB) FindCreatMqttUserData(BusinessId int64) (data *GlobalMqttUser, err error) {
	data = &GlobalMqttUser{}
	err = db.Where("business_id = ?", BusinessId).First(data).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err { //没有找到数据 创建数据
			data, err = db.CreatMqttUserData(BusinessId)
			return data, err
		} else {
			return &GlobalMqttUser{}, err
		}
	}
	return data, err
}

func (db *DB) CreatMqttUserData(BusinessId int64) (data *GlobalMqttUser, err error) {
	data = &GlobalMqttUser{}
	data.BusinessId = BusinessId
	data.CTime = time.Now().Unix()
	data.Status = 1
	data.Remark = slog.Info(BusinessId, "sys", "创建", "mqtt用户")
	data.Username = fmt.Sprintf("ddy%s", utils.RandomStr2(13, 3, 2)) //生成13位随机数 (数字 小写字母)
	data.Password = utils.RandomStr2(32, 3, 3)                       //生成32位随机数 (数字 大小写字母)
	err = db.Create(&data).Error
	if err != nil { //创建失败
		return &GlobalMqttUser{}, err
	}
	return data, nil
}
