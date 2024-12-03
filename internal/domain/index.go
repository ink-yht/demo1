package domain

import "regexp"

// ValidationResult 验证结果结构体
type ValidationResult struct {
	Row           int    `json:"row"`            // 行号
	Index         int    `json:"index"`          // 序号
	IDCard        string `json:"id_card"`        // 身份证号
	DisabilityNo  string `json:"disability_no"`  // 残疾证号码
	Phone         string `json:"phone"`          // 手机号码
	Email         string `json:"email"`          // 邮箱
	ValidationMsg string `json:"validation_msg"` // 验证结果信息
}

type ColumnMapping struct {
	IDCardCol       int `json:"id_card_col"`
	DisabilityNoCol int `json:"disability_no_col"`
	PhoneCol        int `json:"phone_col"`
	EmailCol        int `json:"email_col"`
}

// IsValidIDCard 验证身份证号码是否合法
func IsValidIDCard(idCard string) bool {
	// 简单判断身份证号码长度为 18 位
	return len(idCard) == 18
}

// IsValidPhone 验证手机号格式
func IsValidPhone(phone string) bool {
	// 使用正则表达式验证手机号
	phoneRegex := regexp.MustCompile(`^\d{11}$`)
	return phoneRegex.MatchString(phone)
}

// IsValidEmail 验证邮箱格式
func IsValidEmail(email string) bool {
	// 使用正则表达式验证邮箱
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
