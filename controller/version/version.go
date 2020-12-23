package version

import (
	"GOGOGO/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Version struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Version     uint   `json:"version"`
	Description string `json:"description"`
	URL         string `json:"url"`
	MD5         string `json:"md_5"`
	SHA1        string `json:"sha1"`
}

func GetVersion(c *gin.Context) {
	var version Version
	model.DB.Order("id desc").First(&version)

	c.JSON(http.StatusOK, version)
}
