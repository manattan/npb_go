package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/manattan/npb_go/handler"
	"github.com/manattan/npb_go/usecase"
)

func NewServer(teamUC *usecase.TeamUseCase, playerUC *usecase.PlayerUseCase) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	teamHandler := handler.NewTeamHandler(teamUC)
	playerHandler := handler.NewPlayerHandler(playerUC)

	v0 := e.Group("/api/v0")

	// search team
	v0.GET("/team", teamHandler.GetTeams)
	v0.GET("/team/:id", teamHandler.GetTeam)

	// search player
	v0.GET("/player", playerHandler.GetPlayers)

	return e
}
