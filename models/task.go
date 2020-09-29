package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/zzlpeter/dawn-go/libs/utils"
	"time"
)

type DateTimeStruct struct {
	CreateAt 	time.Time 	`json:"create_at" comment:"创建时间" gorm:"column:create_at;default:null"`
	UpdateAt 	time.Time	`json:"update_at" comment:"更新时间" gorm:"column:update_at;default:null"`
}

type Task struct {
	DateTimeStruct
	Id 			uint 		`json:"id" comment:"主键自增" gorm:"primary_key"`
	TaskKey 	string 		`json:"task_key" comment:"任务唯一key" gorm:"index"`
	Desc 		string 		`json:"desc" comment:"任务描述"`
	ExecuteFunc string		`json:"execute_func" comment:"执行方法"`
	Trigger 	string 		`json:"trigger" comment:"触发类型cron/interval/date"`
	Spec 		string		`json:"spec" comment:"执行时间"`
	Args 		string		`json:"args" comment:"执行参数"`
	IsValid 	bool		`json:"is_valid" comment:"是否有效"`
	Status 		string		`json:"status" comment:"任务执行状态ready/doing"`
	Extra 		string		`json:"extra" comment:"额外信息(json格式)"`
}

func (t *Task) TableName() string {
	return "task"
}

func (t *Task) BeforeCreate(tx *gorm.DB) error {
	// 校验extra是否有值
	if t.Extra == "" {
		t.Extra = "{}"
	}
	// 校验trigger是否为cron/date/interval
	switch t.Trigger {
	case "cron":
		return t.isCronValid()
	case "interval":
		return t.isIntervalValid()
	case "date":
		return t.isDateValid()
	default:
		err := fmt.Sprintf("Trigger: %v is invalid", t.Trigger)
		return errors.New(err)
	}
}

func (t *Task) BeforeUpdate(tx *gorm.DB) error {
	return t.BeforeCreate(tx)
}

func (t Task) isCronValid() error {
 	if t.Spec != "1" {
		err := errors.New("cron format is invalid")
		return err
	}
	return nil
}

func (t Task) isIntervalValid() error {
	_, err := utils.String2Int(t.Spec)
	return err
}

func (t Task) isDateValid() error {
	_, err := time.Parse("2006-01-02 15:04:05", t.Spec)
	return err
}

type TaskExecute struct {
	DateTimeStruct
	Id 			uint 		`json:"id" comment:"主键自增" gorm:"primary_key"`
	TaskId 		uint		`json:"task_id" comment:"任务ID" gorm:"index"`
	Extra 		string 		`json:"extra" comment:"额外信息"`
	TraceId 	string 		`json:"trace_id" comment:"traceID"`
}

func (t *TaskExecute) TableName() string {
	return "task_execute"
}