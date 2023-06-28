package test_test

import (
	"fmt"
	"time"
)

const (
	RateTimeDay   = 1 //天
	RateTimeWeek  = 2 //周
	RateTimeMonth = 3 //月
)

// 获取当天时间戳
func GetDay() (int64, int64) {
	today, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local) //今天
	nextDay := today.AddDate(0, 0, 1)                                                           //明天
	fmt.Println("today=", today, " nextDay=", nextDay)
	return today.Unix(), nextDay.Unix()
}

// 获取本周时间戳
func GetWeekDay() (int64, int64) {
	today, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local) //今天
	offset := int(time.Monday - today.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if offset > 0 {
		offset = -6
	}

	lastoffset := int(time.Saturday - today.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if lastoffset == 6 {
		lastoffset = -1
	}

	firstOfWeek := today.AddDate(0, 0, offset)      //本周一
	lastOfWeeK := today.AddDate(0, 0, lastoffset+2) // 下周一
	fmt.Println("firstOfWeek=", firstOfWeek, " lastOfWeeK=", lastOfWeeK)
	return firstOfWeek.Unix(), lastOfWeeK.Unix()
}

// 获取本月时间戳
func GetMonthDay() (int64, int64) {
	today, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local) //今天
	// now := today

	firstOfMonth := today.AddDate(0, 0, 1-today.Day()) //本月第一天
	lastOfMonth := firstOfMonth.AddDate(0, 1, 0)       //下一个月第一天
	fmt.Println("firstOfMonth=", firstOfMonth, " lastOfMonth=", lastOfMonth)
	return firstOfMonth.Unix(), lastOfMonth.Unix()
}

func GetRateUserValue(STime int64, Type int, Rate int) (int, error) {

	t := time.Now() //获取当前时间，类型是go的时间类型Time

	tY := time.Now().Year()
	tMo := time.Now().Month()
	tD := time.Now().Day()
	tH := time.Now().Hour()
	tMi := time.Now().Minute()
	tS := time.Now().Second()
	//tNaS := time.Now().Nanosecond()

	curTimeDate := time.Date(tY, tMo, tD, 0, 0, 0, 0, time.Local)

	fmt.Println("time.Unix() = ", t.Unix())
	fmt.Println("time.Now() = ", t)
	fmt.Println("tY:tMo:tD:tH:tMi:tS = ", tY, tMo, tD, tH, tMi, tS)
	fmt.Println("curTimeDate = ", curTimeDate)

	GetDay()
	GetWeekDay()
	GetMonthDay()
	switch Type {
	case RateTimeDay: //天

		if STime+86400 < time.Now().Unix() {
			return Rate, nil
		} else {
			return 0, nil
		}
	case RateTimeWeek: //周
		if STime+604800 < time.Now().Unix() {
			return Rate, nil
		} else {
			return 0, nil
		}
	case RateTimeMonth: //月
		if STime+2592000 < time.Now().Unix() {
			return Rate, nil
		} else {
			return 0, nil
		}
	}
	return 0, nil

}
