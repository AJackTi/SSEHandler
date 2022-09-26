package v1

import (
	"SSEHandler/ssehandler"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "net/http/pprof"
	"time"
)

type userInfoResponse struct {
	Fullname        string `json:"fullname"`
	CurrentRegister int64  `json:"current_register"`
	Address         string `json:"address"`
	Image           string `json:"image"`
	Age             uint64 `json:"age"`
	Salary          uint64 `json:"salary"`
}

type userInfoResponses struct {
	CurrentUserResponse    userInfoResponse   `json:"current_user_response"`
	LatestUserInfoResponse []userInfoResponse `json:"latest_user_info"`
}

type userRoutes struct {
	sseHandler *ssehandler.SSEHandler
}

func (hand *handler) newUserRoutes(handler *gin.RouterGroup) {
	r := &userRoutes{sseHandler: hand.sseHandler}

	h := handler.Group("/auth")
	{
		h.GET("", r.GetUserInfo)
		h.GET("/push", r.SendInfo)
	}
}

func (u *userRoutes) GetUserInfo(c *gin.Context) {
	u.sseHandler.Subscribe(c)
	c.JSON(http.StatusOK, "response")
}

func (u *userRoutes) SendInfo(c *gin.Context) {
	responseUserData := userInfoResponses{
		CurrentUserResponse: userInfoResponse{
			Fullname:        "Dang Vo Duc Trong",
			CurrentRegister: time.Now().Unix(),
			Address:         "",
			Image:           "",
			Age:             21,
			Salary:          10000000,
		},
		LatestUserInfoResponse: []userInfoResponse{
			{
				Fullname:        "aaaaa",
				CurrentRegister: time.Now().Unix(),
				Address:         "HCMC",
				Image:           "",
				Age:             20,
				Salary:          10000000,
			},
			{
				Fullname:        "bbbbb",
				CurrentRegister: time.Now().Unix(),
				Address:         "HCMC",
				Image:           "",
				Age:             20,
				Salary:          10000000,
			},
			{
				Fullname:        "ccccc",
				CurrentRegister: time.Now().Unix(),
				Address:         "HCMC",
				Image:           "",
				Age:             20,
				Salary:          10000000,
			},
			{
				Fullname:        "ddddd",
				CurrentRegister: time.Now().Unix(),
				Address:         "HCMC",
				Image:           "",
				Age:             20,
				Salary:          10000000,
			},
			{
				Fullname:        "eeeee",
				CurrentRegister: time.Now().Unix(),
				Address:         "HCMC",
				Image:           "",
				Age:             20,
				Salary:          10000000,
			},
		},
	}
	u.sseHandler.SendJSON(responseUserData)
	c.JSON(http.StatusOK, gin.H{"data": "OK"})
}
