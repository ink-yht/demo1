package domain

import (
	"errors"
	"github.com/dlclark/regexp2"
)

var (
	emailRegex      = regexp2.MustCompile(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`, regexp2.None)
	ErrEmail        = errors.New("邮箱格式无效")
	phoneRegex      = regexp2.MustCompile(`^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0-8]|18[0-9]|19[89])\d{8}$`, regexp2.None)
	ErrPhone        = errors.New("邮箱格式无效")
	idCordRegex     = regexp2.MustCompile(`^[1-9]\d{5}(18|19|20)\d{2}(0[1-9]|1[0-2])(0[1-9]|[12][0-9]|3[01])\d{3}(\d|X|x)$`, regexp2.None)
	ErrIdCord       = errors.New("身份证格式无效")
	disabilityRegex = regexp2.MustCompile(`^[1-9]\d{5}(18|19|20)\d{2}(0[1-9]|1[0-2])(0[1-9]|[12]\d|3[01])\d{3}[0-9Xx][0-9a-zA-Z]{2}$`, regexp2.None)
	ErrDisability   = errors.New("残疾证号码无效")
)

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
	IDCardCol       int    `json:"id_card_col"`
	DisabilityNoCol int    `json:"disability_no_col"`
	PhoneCol        int    `json:"phone_col"`
	EmailCol        int    `json:"email_col"`
	SheetName       string `json:"sheet_name"`
}

// IsValidIDCard 验证身份证号码是否合法
func IsValidIDCard(idCard string) error {
	// 校验邮箱格式
	if match, _ := idCordRegex.MatchString(idCard); !match {
		return ErrIdCord
	}
	return nil
}

// IsValidDisability 验证身份证号码是否合法
func IsValidDisability(disability string) error {
	// 校验邮箱格式
	if match, _ := disabilityRegex.MatchString(disability); !match {
		return ErrDisability
	}
	return nil
}

// IsValidPhone 验证手机号格式
func IsValidPhone(phone string) error {
	// 使用正则表达式验证手机号
	if match, _ := phoneRegex.MatchString(phone); !match {
		return ErrPhone
	}
	return nil
}

// IsValidEmail 验证邮箱格式
func IsValidEmail(email string) error {
	// 校验邮箱格式
	if match, _ := emailRegex.MatchString(email); !match {
		return ErrEmail
	}
	return nil
}
