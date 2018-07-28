package gin

import (
	"net/http"

	ggin "github.com/gin-gonic/gin"
)

// WrapContenxt include gin context and local context
type WrapContenxt struct {
	*ggin.Context
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

// WrapControllerFunction wrap gin.Contenxt and local model.Context
func WrapControllerFunction(ctlFunc func(ctx *WrapContenxt)) ggin.HandlerFunc {
	return func(ctx *ggin.Context) {

		wrapContenxt := &WrapContenxt{
			Context: ctx,
		}

		ctlFunc(wrapContenxt)
	}
}
