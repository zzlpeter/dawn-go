package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zzlpeter/dawn-go/libs/mysql"
	"github.com/zzlpeter/dawn-go/libs/utils"
	"github.com/zzlpeter/dawn-go/middlewares"
	"github.com/zzlpeter/dawn-go/models"
)

// @Tags Tasks
// @Summary Get tasks list
// @Accept x-www-form-urlencoded
// @Produce json
// @Param page query string true "page"
// @Param size query string true "size"
// @Success 200 {object} SwaggerResponse
// @Failure 500 {object} SwaggerResponse
// @Router /assassin/task/tasks [get]
func GetTasks(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("size")
	pageInt, _ := utils.String2Int(page)
	sizeInt, _ := utils.String2Int(size)

	db, _ := mysql.GetDbConn("default")
	tasks := []models.Task{}
	// 查询总数
	var cnt int
	db.Model(&models.Task{}).Count(&cnt)
	db.Offset((pageInt - 1) * sizeInt).Limit(sizeInt).Find(&tasks)
	data := map[string]interface{}{
		"total": cnt,
		"data": tasks,
		"page": pageInt,
		"size": sizeInt,
	}
	middlewares.Response200("success", data, c)
}

// @Tags Tasks
// @Summary Get special task
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} SwaggerResponse
// @Failure 500 {object} SwaggerResponse
// @Router /assassin/task/task [get]
func GetTask(c *gin.Context) {
	id := c.Query("id")

	db, _ := mysql.GetDbConn("default")
	task := models.Task{}
	db.Where("id = ?", id).First(&task)
	middlewares.Response200("success", task, c)
}

// @Tags Tasks
// @Summary Add special task
// @Accept x-www-form-urlencoded
// @Produce json
// @Param task_key formData string true "任务key"
// @Param execute_func formData string true "执行方法"
// @Param desc formData string true "任务描述"
// @Param trigger formData string true "调度类型"
// @Param spec formData string true "调度频次"
// @Param args formData string false "执行参数"
// @Param is_valid formData string true "是否有效"
// @Param status formData string true "执行状态"
// @Param extra formData string false "额外信息"
// @Success 200 {object} SwaggerResponse
// @Failure 500 {object} SwaggerResponse
// @Router /assassin/task/task [post]
func PostTask(c *gin.Context) {
	taskKey := c.PostForm("task_key")
	executeFunc := c.PostForm("execute_func")
	desc := c.PostForm("desc")
	trigger := c.PostForm("trigger")
	spec := c.PostForm("spec")
	args := c.PostForm("args")
	isValid := c.PostForm("is_valid")
	var valid bool
	if isValid == "1" {
		valid = true
	}
	extra := c.PostForm("extra")
	task := models.Task{
		TaskKey: taskKey,
		Desc: desc,
		ExecuteFunc: executeFunc,
		Trigger: trigger,
		Spec: spec,
		Args: args,
		IsValid: valid,
		Status: "ready",
		Extra: extra,
	}
	db, _ := mysql.GetDbConn("default")
	if err := db.Create(&task).Error; err != nil {
		middlewares.ResponseError(400, err.Error(), c)
		return
	}
	middlewares.Response200("success", task, c)
}

// @Tags Tasks
// @Summary Update special task
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id formData int true "任务ID"
// @Param task_key formData string true "任务key"
// @Param execute_func formData string true "执行方法"
// @Param desc formData string true "任务描述"
// @Param trigger formData string true "调度类型"
// @Param spec formData string true "调度频次"
// @Param args formData string false "执行参数"
// @Param is_valid formData string true "是否有效"
// @Param extra formData string false "额外信息(json格式)"
// @Param status formData string true "执行状态"
// @Success 200 {object} SwaggerResponse
// @Failure 500 {object} SwaggerResponse
// @Router /assassin/task/task [put]
func PutTask(c *gin.Context) {
	id := c.PostForm("id")
	taskKey := c.PostForm("task_key")
	executeFunc := c.PostForm("execute_func")
	desc := c.PostForm("desc")
	trigger := c.PostForm("trigger")
	spec := c.PostForm("spec")
	args := c.PostForm("args")
	isValid := c.PostForm("is_valid")
	status := c.PostForm("status")
	extra := c.PostForm("extra")

	db, _ := mysql.GetDbConn("default")
	mapper := map[string]interface{}{
		"task_key": taskKey,
		"desc": desc,
		"execute_func": executeFunc,
		"trigger": trigger,
		"spec": spec,
		"args": args,
		"is_valid": isValid,
		"status": status,
		"extra": extra,
	}

	if err := db.Model(&models.Task{}).Where("id = ?", id).Updates(mapper).Error; err != nil {
		middlewares.ResponseError(400, err.Error(), c)
		return
	}
	middlewares.Response200("success", mapper, c)
}

type SwaggerResponse struct {
	Code 	int
	Data 	map[string]interface{}
	Message string
}
