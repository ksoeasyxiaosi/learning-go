package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"webProject/routers"
)

func TestPingRouter(t *testing.T) {

	r := gin.Default()
	routers.Init(r)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/ping", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Println(w.Body.String())
}
