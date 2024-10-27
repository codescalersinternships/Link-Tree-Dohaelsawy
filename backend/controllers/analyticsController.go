package controllers

import (
	"net/http"
	"strconv"

	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	"github.com/gin-gonic/gin"
)


//	@Summary		Get Analytics
//	@Description	Get Analytics properties of how many users view the tree links and their users name 
//	@Tags			analytics
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path	int	true	"user ID"
//	@Security		basic
//	@Success		200	{object}	SuccessResponse
//	@Failure		400	{object}	ErrResponse
//	@Failure		401	{object}	ErrResponse
//	@Failure		404	{object}	ErrResponse
//	@Failure		500	{object}	ErrResponse
//	@Router			/analytics/get_analytics/{user_id} [get]
func (ds *DBController) GetAnalytics(ctx *gin.Context) {

	var analytics []model.Analytics

	idString := ctx.Params.ByName("user_id")

	user_id, err := strconv.Atoi(idString)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	err = ds.store.GetAllAnalyticsForUser(&analytics, user_id)
	if err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	SuccessRespondJSON(ctx, http.StatusOK, gin.H{"analytics": analytics} )
}

func (ds *DBController) CalculateAnalytics(ctx *gin.Context, guestUsername string, user_id int) {

	var analytics model.Analytics

	if err := ds.store.GetAnalyticsForGuestUsername(&analytics, guestUsername, user_id); err != nil {

		analytics = model.Analytics{
			ClickCount:    1,
			UserID:        user_id,
			GuestUsername: guestUsername,
		}

		if err = ds.store.AddNewVisitor(&analytics); err != nil {
			ErrRespondJSON(ctx, http.StatusInternalServerError, err)
			return
		}

	}
	analytics.ClickCount += 1

	if err := ds.store.UpdateAnalytics(&analytics, analytics.ID); err != nil {
		ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}
}
