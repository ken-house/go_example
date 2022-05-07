package negotiate

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func JSON(code int, data interface{}) (int, gin.Negotiate) {
	return code, gin.Negotiate{
		Offered: []string{binding.MIMEJSON},
		Data:    data,
	}
}
