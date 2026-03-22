package models

import (
	"com.dreamfsk/blog/commons"
	"database/sql/driver"
	"fmt"
	"time"
)

// CustomTime 自定义时间类型
type CustomTime struct {
	time.Time
}

// MarshalJSON 实现 JSON 序列化
func (t CustomTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	formatted := fmt.Sprintf("\"%s\"", t.Format(commons.TimeLayout))
	return []byte(formatted), nil
}

// UnmarshalJSON 实现 JSON 反序列化（可选）
func (t *CustomTime) UnmarshalJSON(data []byte) error {
	// 去掉引号
	str := string(data)
	if str == "null" || str == `""` {
		t.Time = time.Time{}
		return nil
	}

	// 去掉两边的引号
	str = str[1 : len(str)-1]

	// 解析时间
	parsed, err := time.Parse(commons.TimeLayout, str)
	if err != nil {
		return err
	}
	t.Time = parsed
	return nil
}

// Value 实现数据库驱动接口（写入数据库时）
func (t CustomTime) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan 实现数据库驱动接口（从数据库读取时）
func (t *CustomTime) Scan(value interface{}) error {
	if value == nil {
		t.Time = time.Time{}
		return nil
	}

	// 根据实际数据库返回的类型进行处理
	switch v := value.(type) {
	case time.Time:
		t.Time = v
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
}
