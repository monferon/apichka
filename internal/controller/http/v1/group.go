package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/monferon/apichka/pkg/logger"
	"net/http"
)

type testRoutes struct {
	l logger.Interface
}

type testResponse struct {
	str string
}

// @Summary     Show hello world
// @Description Show hello world
// @ID          testID
// @Tags  	    test
// @Accept      json
// @Produce     json
// @Success     200 {object} testResponse
// @Failure     500 {object} response
// @Router      /test/hello [get]
func (t *testRoutes) helloWorld(c *gin.Context) {

	c.JSON(http.StatusOK, "Hello World")
}

func newTestRouter(handler *gin.RouterGroup, l logger.Interface) {

	t := testRoutes{l}

	h := handler.Group("/test")
	{
		h.GET("/hello", t.helloWorld)
		h.GET("/world", t.helloWorld)
	}
}
