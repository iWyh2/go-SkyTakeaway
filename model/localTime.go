package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

// LocalTime 自定义时间类型
type LocalTime time.Time

// MarshalJSON 序列化时间数据
func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

// UnmarshalJSON 反序列化时间数据
func (t *LocalTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = LocalTime(t1)
	return err
}

// Value 实现数据存储前对数据进⾏相关操作
//
//	（为解决sql:
//		converting argument $6 type:
//			unsupported type model.LocalTime, a struct错误，需要是值类型方法）
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

// Scan 实现在数据查询出来之前对数据进⾏相关操作
func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
