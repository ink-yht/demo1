package service

import (
	"context"
	"demo-1/internal/domain"
	"encoding/csv"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"path/filepath"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (svc *Service) ValidateExcel(ctx context.Context, path string, columnMapping domain.ColumnMapping) ([]domain.ValidationResult, error) {
	// 打开文件
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, fmt.Errorf("无法打开 Excel 文件: %v", err)
	}
	defer f.Close()

	// 获取 Sheet1 的所有行
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return nil, fmt.Errorf("无法读取表格数据: %v", err)
	}

	var results []domain.ValidationResult
	// 读取每一行数据
	for i, row := range rows {
		// 跳过表头（第一行）
		if i == 0 {
			continue
		}
		// 使用列号读取对应列的数据
		result := domain.ValidationResult{
			Row:   i + 1,
			Index: i,
		}
		// 假设columnMapping中的列号是从0开始计数的，与Go语言数组下标一致
		if columnMapping.IDCardCol >= 0 && columnMapping.IDCardCol < len(row) {
			result.IDCard = row[columnMapping.IDCardCol]
		}

		if columnMapping.DisabilityNoCol >= 0 && columnMapping.DisabilityNoCol < len(row) {
			result.DisabilityNo = row[columnMapping.DisabilityNoCol]
		}

		if columnMapping.PhoneCol >= 0 && columnMapping.PhoneCol < len(row) {
			result.Phone = row[columnMapping.PhoneCol]
		}

		if columnMapping.EmailCol >= 0 && columnMapping.EmailCol < len(row) {
			result.Email = row[columnMapping.EmailCol]
		}
		// 验证每一列数据
		if !domain.IsValidIDCard(result.IDCard) {
			result.ValidationMsg = "身份证号码无效"
		} else if !domain.IsValidPhone(result.Phone) {
			result.ValidationMsg = "手机号格式不正确"
		} else if !domain.IsValidEmail(result.Email) {
			result.ValidationMsg = "邮箱格式不正确"
		} else {
			result.ValidationMsg = "验证通过"
		}

		// 将结果追加到结果集
		results = append(results, result)
	}
	return results, nil
}

// ExportValidationResults 将验证结果导出为 CSV 文件
func (svc *Service) ExportValidationResults(results []domain.ValidationResult) (string, error) {
	// 确保根目录下的 tmp 文件夹存在
	tmpDir := "./tmp"
	if err := os.MkdirAll(tmpDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("无法创建 tmp 目录: %v", err)
	}

	// 文件路径
	filePath := filepath.Join(tmpDir, "validation_results.csv")

	// 创建文件
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("无法创建 CSV 文件: %v", err)
	}
	defer file.Close()

	// 创建 CSV 写入器
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 写入表头
	writer.Write([]string{"行号", "身份证号", "残疾证号码", "手机号", "邮箱", "验证信息"})

	// 写入数据
	for _, result := range results {
		writer.Write([]string{
			fmt.Sprintf("%d", result.Row),
			result.IDCard,
			result.DisabilityNo,
			result.Phone,
			result.Email,
			result.ValidationMsg,
		})
	}

	return filePath, nil
}
