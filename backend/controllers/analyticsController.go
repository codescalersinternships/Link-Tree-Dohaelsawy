package controllers

import (
	"net/http"
	"strconv"

	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
	"github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/utils"
	"github.com/gin-gonic/gin"
)

func (ds *DBController) GetAnalytics(ctx *gin.Context) {

	var analytics []model.Analytics

	idString := ctx.Params.ByName("user_id")

	user_id, err := strconv.Atoi(idString)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	err = ds.store.GetAllAnalyticsForUser(&analytics, user_id)
	if err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.SuccessRespondJSON(ctx, http.StatusOK, analytics)
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
			utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
			return
		}

	}
	analytics.ClickCount += 1

	if err := ds.store.UpdateAnalytics(&analytics, analytics.ID); err != nil {
		utils.ErrRespondJSON(ctx, http.StatusInternalServerError, err)
		return
	}
}
