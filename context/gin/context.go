package gin

import (
	"net/http"

	ggin "github.com/gin-gonic/gin"
)

// WrapContenxt include gin context and local context
type WrapContenxt struct {
	*ggin.Context
}

// WrapControllerFunction wrap gin.Contenxt and local model.Context
func WrapControllerFunction(ctlFunc func(ctx *WrapContenxt)) ggin.HandlerFunc {
	return func(ctx *ggin.Context) {

		wrapContenxt := &WrapContenxt{
			Context: ctx,
		}

		ctlFunc(wrapContenxt)
	}
}

// Result 返回
type Result map[string]interface{}

// OKJSON 返回状态等于200的JSON
func (w *WrapContenxt) OKJSON(obj interface{}) {
	w.JSON(http.StatusOK, obj)
}

// ErrorJSON 返回状态等于400的Error
func (w *WrapContenxt) ErrorJSON(err string) {
	w.JSON(http.StatusBadRequest, ggin.H{
		"error": err,
	})
}

// BindValidation 绑定验证
func (w *WrapContenxt) BindValidation(v Verifier) bool {
	if err := w.ShouldBind(v); err != nil {
		w.ErrorJSON("参数信息不完整")
		return false
	}
	if err := v.Validation(); err != nil {
		w.ErrorJSON(err.Error())
		return false
	}
	return true
}

// Verifier 验证人
type Verifier interface {
	// Validation 验证
	Validation() error
}
