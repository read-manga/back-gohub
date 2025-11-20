package controller

import (
	"microgo/core/domain/commit"
	"microgo/core/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CommitController struct {
	usecase *usecase.CommitCase
}

func NewCommitController(c *usecase.CommitCase) *CommitController {
	return &CommitController{usecase: c}
}

func (comm *CommitController) CreateCommit(c *gin.Context) {
	repoId := c.Param("repo_id")

	var body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		UserId      string `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	commitId, err := comm.usecase.SaveCommit(commit.Commit{
		Title:       body.Title,
		Description: body.Description,
		RepoId:      repoId,
		UserId:      body.UserId,
		CreatedAt:   time.Now(),
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"commit_id": commitId})
}
