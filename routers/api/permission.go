package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zzlpeter/dawn-go/libs/mysql"
	"github.com/zzlpeter/dawn-go/libs/utils"
	"github.com/zzlpeter/dawn-go/middlewares"
	"github.com/zzlpeter/dawn-go/models"
)

// @Tags Account
// @Summary Get roles list
// @Accept x-www-form-urlencoded
// @Produce json
// @Param page query string true "page"
// @Param size query string true "size"
// @Param role_id query string false "role_id"
// @Param role_desc query string false "role_desc"
// @Success 200 {object} SwaggerResponse
// @Failure 500 {object} SwaggerResponse
// @Router /assassin/account/roles [get]
func GetRoles(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("size")
	pageInt, _ := utils.String2Int(page)
	sizeInt, _ := utils.String2Int(size)
	roleId := c.Query("role_id")
	roleDesc := c.Query("role_desc")

	db, _ := mysql.GetDbConn("permission")
	roles := []models.Role{}
	// 拼接过滤条件
	var cnt int
	query := db.Model(&models.Role{})
	if roleId != "" {
		query = query.Where("role_id LIKE ?", "%"+roleId+"%")
	}
	if roleDesc != "" {
		query = query.Where("role_desc LIKE ?", "%"+roleDesc+"%")
	}
	// 查询总数
	query.Count(&cnt)
	query.Offset((pageInt - 1) * sizeInt).Limit(sizeInt).Find(&roles)
	data := map[string]interface{}{
		"total": cnt,
		"data": roles,
		"page": pageInt,
		"size": sizeInt,
	}
	middlewares.Response200("success", data, c)
}

// @Tags Account
// @Summary Update special role
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id formData string true "id"
// @Param role_id formData string true "role_id"
// @Param role_desc formData string true "role_desc"
// @Success 200 {object} SwaggerResponse
// @Failure 500 {object} SwaggerResponse
// @Router /assassin/account/role [put]
func PutRole(c *gin.Context) {
	id := c.PostForm("id")
	roleId := c.PostForm("role_id")
	roleDesc := c.PostForm("role_desc")

	db, _ := mysql.GetDbConn("permission")
	mapper := map[string]interface{}{
		"role_id": roleId,
		"role_desc": roleDesc,
	}
	if err := db.Model(&models.Role{}).Where("id = ?", id).Updates(mapper).Error; err != nil {
		middlewares.ResponseError(400, err.Error(), c)
		return
	}
	middlewares.Response200("success", mapper, c)
}

// @Tags Account
// @Summary Add special role
// @Accept x-www-form-urlencoded
// @Produce json
// @Param role_id formData string true "role_id"
// @Param role_desc formData string true "role_desc"
// @Success 200 {object} SwaggerResponse
// @Failure 500 {object} SwaggerResponse
// @Router /assassin/account/role [post]
func PostRole(c *gin.Context) {
	roleId := c.PostForm("role_id")
	roleDesc := c.PostForm("role_desc")

	role := models.Role{
		RoleId: roleId,
		RoleDesc: roleDesc,
	}
	db, _ := mysql.GetDbConn("permission")
	if err := db.Create(&role).Error; err != nil {
		middlewares.ResponseError(400, err.Error(), c)
		return
	}
	middlewares.Response200("success", role, c)
}

// @Tags Account
// @Summary Delete special role
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} SwaggerResponse
// @Failure 500 {object} SwaggerResponse
// @Router /assassin/account/role [delete]
func DeleteRole(c *gin.Context) {
	id := c.Query("id")

	db, _ := mysql.GetDbConn("permission")
	if err := db.Where("id = ?", id).Delete(&models.Role{}).Error; err != nil {
		middlewares.ResponseError(400, err.Error(), c)
		return
	}
	middlewares.Response200("success", id, c)
}

// @Tags Account
// @Summary Get permissions list
// @Accept x-www-form-urlencoded
// @Produce json
// @Param page query string true "page"
// @Param size query string true "size"
// @Param role_id query string true "role_id"
// @Param permission_id query string false "permission_id"
// @Param permission_desc query string false "permission_desc"
// @Success 200 {object} SwaggerResponse
// @Failure 500 {object} SwaggerResponse
// @Router /assassin/account/permissions [get]
func GetPermissions(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("size")
	pageInt, _ := utils.String2Int(page)
	sizeInt, _ := utils.String2Int(size)
	roleId := c.Query("role_id")
	permissionId := c.Query("permission_id")
	permissionDesc := c.Query("permission_desc")

	db, _ := mysql.GetDbConn("permission")
	perms := []models.Permission{}
	// 拼接过滤条件
	var cnt int
	query := db.Model(&models.Permission{}).Where("role_id = ?", roleId)
	if permissionId != "" {
		query = query.Where("permission_id LIKE ?", "%"+permissionId+"%")
	}
	if permissionDesc != "" {
		query = query.Where("permission_desc LIKE ?", "%"+permissionDesc+"%")
	}
	// 查询总数
	query.Count(&cnt)
	query.Offset((pageInt - 1) * sizeInt).Limit(sizeInt).Find(&perms)
	data := map[string]interface{}{
		"total": cnt,
		"data": perms,
		"page": pageInt,
		"size": sizeInt,
	}
	middlewares.Response200("success", data, c)
}

// @Tags Account
// @Summary Update special permission
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id formData string true "id"
// @Param role_id formData string true "role_id"
// @Param permission_id formData string true "permission_id"
// @Param permission_desc formData string true "permission_desc"
// @Success 200 {object} SwaggerResponse
// @Failure 500 {object} SwaggerResponse
// @Router /assassin/account/permission [put]
func PutPermission(c *gin.Context) {
	id := c.PostForm("id")
	roleId := c.PostForm("role_id")
	permissionId := c.PostForm("permission_id")
	permissionDesc := c.PostForm("permission_desc")

	db, _ := mysql.GetDbConn("permission")
	mapper := map[string]interface{}{
		"role_id": roleId,
		"permission_desc": permissionDesc,
		"permission_id": permissionId,
	}
	if err := db.Model(&models.Permission{}).Where("id = ?", id).Updates(mapper).Error; err != nil {
		middlewares.ResponseError(400, err.Error(), c)
		return
	}
	middlewares.Response200("success", mapper, c)
}

// @Tags Account
// @Summary Add special permission
// @Accept x-www-form-urlencoded
// @Produce json
// @Param role_id formData string true "role_id"
// @Param permission_id formData string true "permission_id"
// @Param permission_desc formData string true "permission_desc"
// @Success 200 {object} SwaggerResponse
// @Failure 500 {object} SwaggerResponse
// @Router /assassin/account/permission [post]
func PostPermission(c *gin.Context) {
	roleId := c.PostForm("role_id")
	permissionId := c.PostForm("permission_id")
	permissionDesc := c.PostForm("permission_desc")

	perm := models.Permission{
		RoleId: roleId,
		PermissionDesc: permissionDesc,
		PermissionId: permissionId,
	}
	db, _ := mysql.GetDbConn("permission")
	if err := db.Create(&perm).Error; err != nil {
		middlewares.ResponseError(400, err.Error(), c)
		return
	}
	middlewares.Response200("success", perm, c)
}

// @Tags Account
// @Summary Delete special permission
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} SwaggerResponse
// @Failure 500 {object} SwaggerResponse
// @Router /assassin/account/permission [delete]
func DeletePermission(c *gin.Context) {
	id := c.Query("id")

	db, _ := mysql.GetDbConn("permission")
	if err := db.Where("id = ?", id).Delete(&models.Permission{}).Error; err != nil {
		middlewares.ResponseError(400, err.Error(), c)
		return
	}
	middlewares.Response200("success", id, c)
}