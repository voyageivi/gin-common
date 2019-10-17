package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/voyageivi/gin-common/e"
	"github.com/voyageivi/gin-common/logging"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}
	config := &validator.Config{TagName: "validate"}
	validate := validator.New(config)
	errs := validate.Struct(form)
	if errs != nil {
		return http.StatusInternalServerError, e.ERROR
	}
	return http.StatusOK, e.SUCCESS
}

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}

	return
}
