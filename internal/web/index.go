package web

import (
	"demo-1/internal/domain"
	"demo-1/internal/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

type Handler struct {
	svc *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) RegisterRoutes(server *gin.Engine) {
	server.POST("/upload", h.Upload)
	server.POST("/export", h.ExportCSV)
}

func (h *Handler) Upload(ctx *gin.Context) {
	//获取上传的文件
	file, err := ctx.FormFile("file")
	if err != nil {
		// 如果文件上传失败，返回错误信息
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "文件上传失败，请检查文件"})
		return
	}

	// 保存上传文件到临时路径
	tempDir := os.TempDir()
	tempFilePath := filepath.Join(tempDir, file.Filename)
	if err := ctx.SaveUploadedFile(file, tempFilePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "保存文件失败"})
		return
	}

	// 获取FormData中的JSON数据字符串
	jsonData := ctx.PostForm("columnMapping")
	// 将JSON数据字符串解析为ColumnMapping结构体
	var columnMapping domain.ColumnMapping
	if err := json.Unmarshal([]byte(jsonData), &columnMapping); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的列号映射数据"})
		return
	}
	fmt.Println(columnMapping)

	//调用 service 进行处理
	results, err := h.svc.ValidateExcel(ctx, tempFilePath, columnMapping)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": "文件处理失败：" + err.Error()})
		return
	}
	//返回验证结果（JSON 格式）
	ctx.JSON(http.StatusOK, results)
	fmt.Println(results)
}

// ExportCSV 导出验证结果为 CSV 文件
func (h *Handler) ExportCSV(ctx *gin.Context) {
	// 定义一个 JSON 数组接收验证结果
	var results []domain.ValidationResult

	// 从请求中解析 JSON 数据
	if err := ctx.BindJSON(&results); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的 JSON 数据"})
		return
	}

	// 调用服务层生成 CSV 文件
	filePath, err := h.svc.ExportValidationResults(results)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "CSV 文件导出失败：" + err.Error()})
		return
	}

	// 返回生成的文件
	ctx.File(filePath)
}
