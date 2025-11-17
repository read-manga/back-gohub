package controller

import (
	"microgo/core/domain/repo"
	"microgo/core/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RepoController struct {
	usecase *usecase.RepoUseCase
}

func NewRepoController(r *usecase.RepoUseCase) *RepoController {
	return &RepoController{r}
}

func (r *RepoController) CreateRepository(c *gin.Context) {
	var repository repo.Repo

	if err := c.ShouldBindJSON(&repository); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := r.usecase.SaveRepo(repository)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Repositório foi criado com sucesso."})
}
