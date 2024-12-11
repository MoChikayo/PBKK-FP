package utils

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ParseBody(r *http.Request, x interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, x)
}

type respFormat int

const (
	FormatHTML respFormat = iota
	FormatJSON
)

var formatMap = map[string]respFormat{
	"html": FormatHTML,
	"json": FormatJSON,
}

func GetFormat(c *gin.Context) respFormat {
	format := c.DefaultQuery("format", "json")
	if f, exists := formatMap[format]; exists {
		return f
	}
	return FormatJSON // Default to JSON if no valid format is specified
}
