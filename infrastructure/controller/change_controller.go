package controller

import (
	"io"
	"microgo/core/domain/change"
	"microgo/core/usecase"
	"microgo/infrastructure/adapters"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChangeController struct {
	usecase *usecase.ChangeCase
}

func NewChangeController(c *usecase.ChangeCase) *ChangeController {
	return &ChangeController{usecase: c}
}

func (cha *ChangeController) CreateChange(c *gin.Context) {
	commitId := c.PostForm("commit_id")

	form, _ := c.MultipartForm()
	files := form.File["files"]

	var current []change.NewFilesBody

	for _, file := range files {
		f, _ := file.Open()
		data, _ := io.ReadAll(f)
		f.Close()

		s3 := adapters.NewS3Adapter()
		hash := s3.CreateFileHash(data)

		current = append(current, change.NewFilesBody{
			Path:    file.Filename,
			Hash:    hash,
			Content: data,
		})
	}

	previuos := usecase.CommitCase{}
	commitIdPrevious, err := previuos.GetCommitByDate()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Não foi possível subir arquivos."})
		return
	}

	changesResult, err := cha.usecase.GetChanges(commitIdPrevious.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Não foi possível subir arquivos."})
	}

	// next

}
