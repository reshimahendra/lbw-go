package helper

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	e "github.com/reshimahendra/lbw-go/internal/pkg/errors"
)

// TestAPIResponse will test ApiResponse wrapper function
func TestAPIResponse(t *testing.T) {
    type data struct{
        Code int
        Data string
    }
    cases := []struct{
        name string
        code int
        message string
        data data 
    }{
        {"SuccessA", 200, "seccessfull", data{Code: 200, Data: "test success"}},
        {"SuccessB", 201, "seccessfull", data{Code: 201, Data: "test success"}},
    }

    gin.SetMode(gin.TestMode)
    r := gin.Default()

    for _, tt := range cases {
        t.Run(tt.name, func(t *testing.T){
            url := fmt.Sprintf("/%s", strings.ToLower(tt.name))
            r.GET(url, func(c *gin.Context){
                APIResponse(c, tt.code, tt.message, tt.data)
            })

            req, err := http.NewRequest(http.MethodGet, url, nil)
            if err != nil {
                t.Fatalf("error creating test request: %v\n", err)
            }

            w := httptest.NewRecorder()

            r.ServeHTTP(w, req)

            if w.Code != tt.code {
                t.Fatalf("expecting got status'%d' but got '%d'", tt.code, w.Code)
            }
        })
    }
} 

func TestApiErrorResponse(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := gin.Default()
    
    r.GET("/test", func(c *gin.Context){
        err := e.New(e.ErrDataNotFound)
        APIErrorResponse(c, http.StatusBadRequest, err)
    })

    req, err := http.NewRequest(http.MethodGet, "/test", nil)
    if err != nil {
        t.Fatalf("error creating test request: %v\n", err)
    }

    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)

    if w.Code != http.StatusBadRequest {
        t.Fatalf("expecting status '%d' but got '%d'", http.StatusBadRequest, w.Code)
    }
}
