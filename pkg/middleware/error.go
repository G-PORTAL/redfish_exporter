package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)

	written, err := r.ResponseWriter.Write(b)
	if err != nil {
		return written, fmt.Errorf("error writing response: %w", err)
	}

	return written, nil
}

func ErrorLogger(c *gin.Context) {
	writer := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
	c.Writer = writer
	c.Next()

	if writer.Status() != http.StatusOK {
		// GIN logger format: 2024/03/27 - 13:54:44
		now := time.Now().Format("2006/01/02 - 15:04:05")
		fmt.Printf("[GIN] %s | %s", now, writer.body.String())
	}
}
