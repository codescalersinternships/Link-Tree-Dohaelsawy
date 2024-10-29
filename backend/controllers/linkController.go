package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/utils"
	"github.com/gin-gonic/gin"
)

var (
	ErrNotFound = errors.New("not found")
)

type LinkReq struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// @Summary		Create Link
// @Description	Create Link
// @Tags			link
// @Accept			json
// @Produce		json
// @Param			LinkReq	body	LinkReq	true	"link attribute"
// @Security		basic
// @Success		200	{object}	SuccessResponse
// @Failure		400	{object}	ErrResponse
// @Failure		401	{object}	ErrResponse
// @Failure		404	{object}	ErrResponse
// @Failure		500	{object}	ErrResponse
// @Router			/link/create_link [post]
func (ds *DBController) CreateLink(ctx *gin.Context) {

	var reqBody LinkReq

	if err := ctx.BindJSON(&reqBody); err != nil {
		ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	if err := ds.Validate.Struct(reqBody); err != nil {
		ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	user_id, err := utils.ExtractTokenID(ctx, *ds.Config)

	if err != nil {
		errorMessage := fmt.Errorf("can't find your token %s", err)
		ErrRespondJSON(ctx, http.StatusInternalServerError, errorMessage)
		return
	}

	link := model.Link{
		Name:   reqBody.Name,
		Url:    reqBody.Url,
		UserID: user_id,
	}

	err = ds.store.AddNewLink(&link)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	SuccessRespondJSON(ctx, http.StatusOK, gin.H{"link": link})
}

// @Summary		Delete Link
// @Description	Delete Link
// @Tags			link
// @Accept			json
// @Produce		json
// @Param			link_id	path	int	true	"link id"
// @Security		basic
// @Success		200	{object}	SuccessResponse
// @Failure		400	{object}	ErrResponse
// @Failure		401	{object}	ErrResponse
// @Failure		404	{object}	ErrResponse
// @Failure		500	{object}	ErrResponse
// @Router			/link/delete_link/{link_id} [delete]
func (ds *DBController) DeleteLink(ctx *gin.Context) {

	var link model.Link

	idString := ctx.Params.ByName("link_id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}
	if err := ds.store.GetOneLink(&link, id); err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	err = ds.store.DeleteLink(&link, id)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	SuccessRespondJSON(ctx, http.StatusOK, "deleted")
}

// @Summary		Update Link
// @Description	Update Link
// @Tags			link
// @Accept			json
// @Produce		json
// @Param			link_id	path	int	true	"link_id"
// @Param			LinkReq	body	LinkReq	true	"Link Req"
// @Security		basic
// @Success		200	{object}	SuccessResponse
// @Failure		400	{object}	ErrResponse
// @Failure		401	{object}	ErrResponse
// @Failure		404	{object}	ErrResponse
// @Failure		500	{object}	ErrResponse
// @Router			/link/update_link/{link_id} [put]
func (ds *DBController) UpdateLink(ctx *gin.Context) {

	var reqBody LinkReq

	if err := ctx.BindJSON(&reqBody); err != nil {
		ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	if err := ds.Validate.Struct(reqBody); err != nil {
		ErrRespondJSON(ctx, http.StatusBadRequest, err)
		return
	}

	idString := ctx.Params.ByName("link_id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	var link model.Link

	err = ds.store.GetOneLink(&link, id)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	if !checkEmpty(reqBody.Name) {
		link.Name = reqBody.Name
	}

	if !checkEmpty(reqBody.Url) {
		link.Url = reqBody.Url
	}

	err = ds.store.PutOneLink(&link, link.ID)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	SuccessRespondJSON(ctx, http.StatusOK, gin.H{"link": link})
}

// @Summary		Get Links
// @Description	Get Links
// @Tags			link
// @Accept			json
// @Produce		json
// @Param			username	path	string	true	"username"
// @Security		basic
// @Success		200	{object}	SuccessResponse
// @Failure		400	{object}	ErrResponse
// @Failure		401	{object}	ErrResponse
// @Failure		404	{object}	ErrResponse
// @Failure		500	{object}	ErrResponse
// @Router			/link_tree/{username} [get]
func (ds *DBController) GetLinks(ctx *gin.Context) {

	var user model.User
	username := ctx.Params.ByName("username")

	if err := ds.store.GetUserUsername(&user, username); err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	var links []model.Link
	if err := ds.store.GetAllLinksForUser(&links, user.ID); err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	getGuestUsername(ctx, ds, user.ID)

	SuccessRespondJSON(ctx, http.StatusOK, gin.H{"links": links})
}

func getGuestUsername(ctx *gin.Context, ds *DBController, user_id int) {

	guest_id, err := utils.ExtractTokenID(ctx, *ds.Config)
	if err != nil {
		ds.CalculateAnalytics(ctx, "unknown", user_id)
		return
	}

	if guest_id == user_id {
		return
	}

	var guest model.User
	if err := ds.store.GetUserID(&guest, guest_id); err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	ds.CalculateAnalytics(ctx, guest.Username, user_id)
}

func checkEmpty(item string) bool {
	return item == ""
}
