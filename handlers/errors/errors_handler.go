package errors

import (
	"com.dreamfsk/blog/models/common/response"
	"com.dreamfsk/blog/repository/models"
	"com.dreamfsk/blog/services/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateSysError 创建错误日志
// @Tags SysError
// @Summary 创建错误日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.SysError true "创建错误日志"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /sysError/createSysError [post]
func (h *handler) CreateSysError(c *gin.Context) {
	var sysError models.SysError
	err := c.ShouldBindJSON(&sysError)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = h.errService.CreateSysError(c.Request.Context(), &sysError)
	if err != nil {
		h.zl.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSysError 删除错误日志
// @Tags SysError
// @Summary 删除错误日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.SysError true "删除错误日志"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /sysError/deleteSysError [delete]
func (h *handler) DeleteSysError(c *gin.Context) {
	// 创建业务用Context
	ID := c.Query("ID")
	err := h.errService.DeleteSysError(c.Request.Context(), ID)
	if err != nil {
		h.zl.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSysErrorByIds 批量删除错误日志
// @Tags SysError
// @Summary 批量删除错误日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /sysError/deleteSysErrorByIds [delete]
func (h *handler) DeleteSysErrorByIds(c *gin.Context) {
	// 创建业务用Context
	IDs := c.QueryArray("IDs[]")
	err := h.errService.DeleteSysErrorByIds(c.Request.Context(), IDs)
	if err != nil {
		h.zl.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSysError 更新错误日志
// @Tags SysError
// @Summary 更新错误日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body system.SysError true "更新错误日志"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /sysError/updateSysError [put]
func (h *handler) UpdateSysError(c *gin.Context) {

	var sysError models.SysError
	err := c.ShouldBindJSON(&sysError)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = h.errService.UpdateSysError(c.Request.Context(), &sysError)
	if err != nil {
		h.zl.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSysError 用id查询错误日志
// @Tags SysError
// @Summary 用id查询错误日志
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询错误日志"
// @Success 200 {object} response.Response{data=system.SysError,msg=string} "查询成功"
// @Router /sysError/findSysError [get]
func (h *handler) FindSysError(c *gin.Context) {
	ID := c.Query("ID")
	resysError, err := h.errService.GetSysError(c.Request.Context(), ID)
	if err != nil {
		h.zl.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(resysError, c)
}

// GetSysErrorList 分页获取错误日志列表
// @Tags SysError
// @Summary 分页获取错误日志列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query systemReq.SysErrorSearch true "分页获取错误日志列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /sysError/getSysErrorList [get]
func (h *handler) GetSysErrorList(c *gin.Context) {
	// 创建业务用Context
	var pageInfo system.SysErrorSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := h.errService.GetSysErrorInfoList(c.Request.Context(), &pageInfo)
	if err != nil {
		h.zl.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetSysErrorSolution 触发错误日志的异步处理
// @Tags SysError
// @Summary 根据ID触发处理：标记为处理中，1分钟后自动改为处理完成
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query string true "错误日志ID"
// @Success 200 {object} response.Response{msg=string} "处理已提交"
// @Router /sysError/getSysErrorSolution [get]
func (h *handler) GetSysErrorSolution(c *gin.Context) {
	// 兼容 id 与 ID 两种参数
	ID := c.Query("id")
	if ID == "" {
		response.FailWithMessage("缺少参数: id", c)
		return
	}

	err := h.errService.GetSysErrorSolution(c.Request.Context(), ID)
	if err != nil {
		h.zl.Error("处理触发失败!", zap.Error(err))
		response.FailWithMessage("处理触发失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("已提交至AI处理", c)
}
