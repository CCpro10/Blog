package util
import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"Blog/pkg/setting"
)

//获取query的page
func GetPage(c *gin.Context) int {
	result := 0
	//引入外部包,string转换成int的方法
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}