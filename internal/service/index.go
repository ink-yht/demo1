package service

import (
	"context"
	"demo-1/internal/domain"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strings"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (svc *Service) ValidateExcel(ctx context.Context, path string, sheetName string, columnMapping domain.ColumnMapping) ([]domain.ValidationResult, error) {
	// 打开文件
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, fmt.Errorf("无法打开 Excel 文件: %v", err)
	}
	defer f.Close()

	// 获取指定工作表的所有行
	rows, err := f.GetRows(sheetName)
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
		var validationErrors []string

		if err := domain.IsValidIDCard(result.IDCard); err != nil {
			validationErrors = append(validationErrors, "身份证号码无效")
		}

		if err := domain.IsValidDisability(result.DisabilityNo); err != nil {
			validationErrors = append(validationErrors, "残疾证号码无效")
		}

		if err := domain.IsValidPhone(result.Phone); err != nil {
			validationErrors = append(validationErrors, "手机号格式不正确")
		}

		if err := domain.IsValidEmail(result.Email); err != nil {
			validationErrors = append(validationErrors, "邮箱格式不正确")
		}

		if len(validationErrors) == 0 {
			result.ValidationMsg = "验证通过"
		} else {
			// 将所有错误信息拼接成一个字符串
			result.ValidationMsg = strings.Join(validationErrors, "; ")
		}

		// 将结果追加到结果集
		results = append(results, result)
	}
	return results, nil
}
