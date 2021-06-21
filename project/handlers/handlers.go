package handlers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go-reverse-proxy/project/utils"
	"io"
	"net/http"
	"os"
)

func GetProxy(c *gin.Context) {
	c.GetHeader("Host")
	backend := os.Getenv("backend")
	res := utils.CheckCache(backend)
	buf := bytes.NewBuffer(res.Body())
	b, _ := io.ReadAll(buf)
	c.Data(http.StatusOK, "text/html; charset=utf-8", b)
}
