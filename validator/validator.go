package validator

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/yyliziqiu/waf/logs"
)

const HIdLen = 24

var hid validator.Func = func(fl validator.FieldLevel) bool {
	if hid, ok := fl.Field().Interface().(string); ok {
		return len(hid) == HIdLen
	}
	return false
}

var mhid validator.Func = func(fl validator.FieldLevel) bool {
	if hidStr, ok := fl.Field().Interface().(string); ok {
		hids := strings.Split(hidStr, ",")
		for _, v := range hids {
			if len(v) != HIdLen {
				return false
			}
		}
		max, err := strconv.Atoi(fl.Param())
		if err != nil {
			return false
		}
		return len(hids) <= max
	}
	return false
}

func Initialize() {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		logs.Fatal("初始化验证器失败")
	}
	if err := v.RegisterValidation("hid", hid); err != nil {
		logs.Fatal(err)
	}
	if err := v.RegisterValidation("mhid", mhid); err != nil {
		logs.Fatal(err)
	}
}
