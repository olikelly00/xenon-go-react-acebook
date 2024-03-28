package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/makersacademy/go-react-acebook-template/api/src/auth"
	"github.com/makersacademy/go-react-acebook-template/api/src/models"
)

type JSONComment struct {
	ID      uint   `json:"_id"`
	Message string `json:"message"`
	Likes   int    `json:"likes"`
	PostId  int    `json:"post_id"`
	UserId  int    `json:"user_id"`
}

func GetAllCommentsByPostId(ctx *gin.Context) {
	// retrieving a parameter from the request URL here below
	// needing route to be structured like this: /posts/:post_id
	postID := ctx.Param("post_id")
	comments, err := models.FetchAllCommentsByPostId(postID)

	if err != nil {
		SendInternalError(ctx, err)
		return
	}

	val, _ := ctx.Get("userID")
	userID := val.(string)
	token, _ := auth.GenerateToken(userID)

	// Convert comments to JSON Structs
	jsonComments := make([]JSONComment, 0)
	for _, comment := range *comments {
		jsonComments = append(jsonComments, JSONComment{
			Message: comment.Message,
			ID:      comment.ID,
			Likes:   comment.Likes,
			PostId:  comment.PostId,
			UserId:  comment.UserId,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"comments": jsonComments, "token": token})
}

type createCommentRequestBody struct {
	Message string
	Likes   int
	//COMPLETE THIS
}

func CreateComment(ctx *gin.Context) {
	var requestBody createCommentRequestBody
	err := ctx.BindJSON(&requestBody)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if len(requestBody.Message) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Post message empty"})
		return
	}

	postID := ctx.Param("post_id")
	postIdInt, err := strconv.Atoi(postID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	userIDNum := ctx.Param("user_id")
	userIDInt, err := strconv.Atoi(userIDNum)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	newComment := models.Comment{
		Message: requestBody.Message,
		Likes:   requestBody.Likes,
		PostId:  postIdInt,
		UserId:  userIDInt,
	}

	_, err = newComment.Save()
	if err != nil {
		SendInternalError(ctx, err)
		return
	}

	val, _ := ctx.Get("userID")
	userID := val.(string)
	token, _ := auth.GenerateToken(userID)

	ctx.JSON(http.StatusCreated, gin.H{"message": "Comment created", "token": token})
}
