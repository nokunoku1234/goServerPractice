package validator

import (
	"goServerPractice/internal/transport"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{
		validator: validator.New(),
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return cv.translateError(err)
	}
	return nil
}

func (cv *CustomValidator) translateError(err error) error {
	var details []transport.FieldIssue

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			details = append(details, transport.FieldIssue{
				Field:  e.Field(),
				Reason: getErrorMessage(e),
			})
		}
	}

	return &ValidationError{
		Details: details,
	}
}

func getErrorMessage(e validator.FieldError) string {
	field := e.Field()
	tag := e.Tag()

	switch field {
	case "Email":
		switch tag {
		case "required":
			return "メールアドレスは必須です"
		case "email":
			return "有効なメールアドレスを入力してください"
		}
	case "Password":
		switch tag {
		case "required":
			return "パスワードは必須です"
		case "min":
			return "パスワードは8文字以上で入力してください"
		case "max":
			return "パスワードは72文字以内で入力してください"
		}
	}
	
	return "入力値が正しくありません"
}
