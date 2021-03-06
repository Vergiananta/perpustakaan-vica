package response

import (
	"github.com/gin-gonic/gin"
)

type jsonResponder struct {
	ctx *gin.Context
}

const (
dataKey      = "data"
pageNoKey    = "pageNo"
totalDataKey = "totalData"
errorKey     = "error"
messageKey   = "message"
)

func (j *jsonResponder) SetContext(ctx interface{}) IResponder {
	j.ctx = ctx.(*gin.Context)
	return j
}
func (j *jsonResponder) PagingResponder(successCode int, data interface{}, pageNo int, totalData int) {
j.ctx.JSON(successCode, gin.H{
dataKey:      data,
pageNoKey:    pageNo,
totalDataKey: totalData,
})
}

func (j *jsonResponder) SingleResponder(successCode int, data interface{}) {
j.ctx.JSON(successCode, gin.H{
dataKey: data,
})
}

func (j *jsonResponder) ErrorResponder(errorCode int, appErrorCode string, errorMessage string) {
	if appErrorCode != "" {
		j.ctx.JSON(errorCode, gin.H{
			errorKey:   appErrorCode,
messageKey: errorMessage,
})
} else {
j.ctx.JSON(errorCode, gin.H{
messageKey: errorMessage,
})
}
}

func NewJsonResponder() IResponder {
return &jsonResponder{}
}