package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"code-kata/api"
	"code-kata/utils/logger"
)

func main() {
	engine := gin.Default()

	engine.Static("/codekata/js", "./static/js")
	engine.LoadHTMLGlob("static/*.html")

	uiGroup := engine.Group("/codekata")
	{
		uiGroup.GET("/", api.Index)
	}

	apiGroup := engine.Group("/codekata/api")
	{
		apiGroup.POST("/retrieve_balance_sheet", api.RetrieveBalanceSheet)
		apiGroup.POST("/submit_loan_application", api.SubmitLoanApplication)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	instance := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: engine,
	}

	logger.Info(fmt.Sprintf("Starting server on port %s", port))
	if err := instance.ListenAndServe(); err != nil {
		logger.Error(fmt.Sprintf("Failed to start server: %s", err.Error()))
	}
}
