package handler

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

// Upload upload files.
func get_test(c *gin.Context) {
    form, _ := c.MultipartForm()
    files := form.File["file"]

    for _, file := range files {
        err := c.SaveUploadedFile(file, "images/"+file.Filename)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        }
    }

    c.JSON(http.StatusOK, gin.H{"message": "success!!"})
}