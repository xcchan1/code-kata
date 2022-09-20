package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	"code-kata/accounting_providers"
	"code-kata/accounting_providers/model"
	"code-kata/loan_application"
	"code-kata/utils/log"
)

var trans ut.Translator

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		uni := ut.New(en, en)
		trans, _ = uni.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(v, trans)
	}
}

func Index(ctx *gin.Context) {
	providers := accounting_providers.AllProviders()
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"providers": providers,
	})
}

func RetrieveBalanceSheet(ctx *gin.Context) {
	var req RetreiveBalanceSheetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			resp := Response{
				ErrorMessage: handleValidationError(err),
			}
			ctx.PureJSON(http.StatusBadRequest, resp)
		} else {
			log.Error(fmt.Sprintf("Failed to bind request: %s", err.Error()))
			ctx.PureJSON(http.StatusBadRequest, nil)
		}
		return
	}
	params := model.RetrieveBalanceSheetParams{
		BusinessName: req.BusinessName,
	}
	balanceSheet, err := accounting_providers.RetrieveBalanceSheet(ctx, req.AccountingProvider, &params)
	if err != nil {
		resp := Response{
			ErrorMessage: err.Error(),
		}
		if _, ok := err.(*accounting_providers.ProviderNotFoundError); ok {
			log.Error(fmt.Sprintf("Provider %s not found", req.AccountingProvider))
			ctx.PureJSON(http.StatusBadRequest, resp)
			return
		} else {
			ctx.PureJSON(http.StatusInternalServerError, resp)
			return
		}
	}
	resp := toRetrieveBalanceSheetResponse(balanceSheet)
	ctx.PureJSON(http.StatusOK, resp)
}

func SubmitLoanApplication(ctx *gin.Context) {
	var req SubmitLoanApplicationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			resp := Response{
				ErrorMessage: handleValidationError(err),
			}
			ctx.PureJSON(http.StatusBadRequest, resp)
		} else {
			log.Error(fmt.Sprintf("Failed to bind request: %s", err.Error()))
			ctx.PureJSON(http.StatusBadRequest, nil)
		}
		return
	}
	params := loan_application.LoanApplicationParams{
		BusinessName:    req.BusinessName,
		YearEstablished: req.YearEstablished,
		LoanAmount:      req.LoanAmount,
	}
	processor, err := loan_application.NewLoanApplicationProcessor(req.AccountingProvider)
	if err != nil {
		resp := Response{
			ErrorMessage: err.Error(),
		}
		if _, ok := err.(*accounting_providers.ProviderNotFoundError); ok {
			log.Error(fmt.Sprintf("Provider %s not found", req.AccountingProvider))
			ctx.PureJSON(http.StatusBadRequest, resp)
			return
		} else {
			log.Error(fmt.Sprintf("loan_application.NewLoanApplicationProcessor|%s", err.Error()))
			ctx.PureJSON(http.StatusInternalServerError, resp)
			return
		}
	}
	result, err := processor.SubmitLoanApplication(ctx, &params)
	if err != nil {
		log.Error(fmt.Sprintf("processor.SubmitLoanApplication|%s", err.Error()))
		resp := Response{
			ErrorMessage: err.Error(),
		}
		ctx.PureJSON(http.StatusInternalServerError, resp)
		return
	}
	data := LoanApplicationResult{
		Verdict: result.Outcome,
	}
	if result.Outcome {
		data.PreAssessmentValue = &result.PreAssessmentValue
		data.EligibleLoanAmount = &result.EligibleLoanAmount
	}
	resp := Response{
		Data: data,
	}
	ctx.PureJSON(http.StatusOK, resp)
}
