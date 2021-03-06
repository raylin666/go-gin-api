package validator

import (
	"gin-api/internal/constant"
	"gin-api/internal/http"
	"gin-api/internal/utils"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

func Validate(ctx *utils.Context, validate interface{}) (ok bool) {
	// 注册翻译器
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")

	// 注册验证器
	valid := validator.New()

	//注册一个函数，获取struct tag里自定义的label作为字段名
	valid.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	// 验证器注册翻译器
	err := zh_translations.RegisterDefaultTranslations(valid, trans)
	if err != nil {
		http.ErrorResponse(ctx, constant.StatusBusinessValidationHandleError)
		return false
	}

	err = valid.Struct(validate)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			http.Response(ctx, &http.Output{
				Builder: http.Builder{
					Code: constant.StatusUnprocessableEntity,
					Message: err.Translate(trans),
				},
			})

			return false
		}
	}

	return true
}