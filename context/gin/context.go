package gin

import (
	"errors"
	"net/http"

	ggin "github.com/gin-gonic/gin"
)

// WrapContext include gin context and local context
type WrapContext struct {
	*ggin.Context
}

// WrapControllerFunction wrap gin.Contenxt and local model.Context
func WrapControllerFunction(ctlFunc func(ctx *WrapContext)) ggin.HandlerFunc {
	return func(ctx *ggin.Context) {

		wrapContenxt := &WrapContext{
			Context: ctx,
		}

		ctlFunc(wrapContenxt)
	}
}

// Result 返回
type Result map[string]interface{}

// OKJSON 返回状态等于200的JSON
func (w *WrapContext) OKJSON(obj interface{}) {
	w.JSON(http.StatusOK, obj)
}

var (
	errIncomplete = errors.New("参数信息不完整")
)

// ErrorJSON 返回状态等于400的Error
func (w *WrapContext) ErrorJSON(err error) {
	w.JSON(http.StatusBadRequest, ggin.H{
		"error": err.Error(),
	})
}

// BindValidation 绑定验证
func (w *WrapContext) BindValidation(v Verifier, obj interface{}) bool {
	if err := w.ShouldBind(v); err != nil {
		w.ErrorJSON(errIncomplete)
		return false
	}
	if err := v.Validation(obj); err != nil {
		w.ErrorJSON(err)
		return false
	}
	return true
}

// Verifier 验证人
type Verifier interface {
	// Validation 验证
	Validation(obj interface{}) error
}
