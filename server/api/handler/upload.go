package handler

import (
	"github.com/BioChemML/SIC50/server/config"
	"github.com/BioChemML/SIC50/server/utils/image"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
	"net/http"
	"path"
)

func Upload(c *gin.Context) {
	mf, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var mfh *multipart.FileHeader
	for _, m := range mf.File {
		for _, i := range m {
			mfh = i
		}
	}
	fPath := path.Join(config.Cfg.UploadPath, uuid.New().String())
	err = image.Save(mfh, fPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, fPath)
}
