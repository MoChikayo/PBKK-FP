package utils

import (
    "github.com/gin-gonic/gin"
)

type Response struct {
    Success bool        `json:"success"`
    Message string      `json:"message,omitempty"`
    Data    interface{} `json:"data,omitempty"`
}

func RespondSuccess(c *gin.Context, status int, data interface{}) {
    c.JSON(status, Response{
        Success: true,
        Data:    data,
    })
}

func RespondError(c *gin.Context, status int, msg string) {
    c.JSON(status, Response{
        Success: false,
        Message: msg,
    })
}
