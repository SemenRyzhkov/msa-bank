// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.2 DO NOT EDIT.
package models

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// AllPayment defines model for AllPayment.
type AllPayment = []PaymentPlan

// Credit defines model for Credit.
type Credit struct {
	Amount   float32            `json:"amount"`
	ClientId openapi_types.UUID `json:"clientId"`
	Id       openapi_types.UUID `json:"id"`
	Months   int                `json:"months"`
	Rate     float32            `json:"rate"`
}

// EarlyRepayment defines model for EarlyRepayment.
type EarlyRepayment struct {
	Amount   float32            `json:"amount"`
	ClientId openapi_types.UUID `json:"clientId"`
}

// FullRepayment defines model for FullRepayment.
type FullRepayment struct {
	ClientId openapi_types.UUID `json:"clientId"`
}

// PaymentPlan defines model for PaymentPlan.
type PaymentPlan struct {
	Amount       float32 `json:"amount"`
	Month        int     `json:"month"`
	MonthPayment float32 `json:"monthPayment"`
}

// Rate defines model for Rate.
type Rate struct {
	Amount *float32 `json:"amount,omitempty"`
	Months *int     `json:"months,omitempty"`
	Rate   *float32 `json:"rate,omitempty"`
}

// PostCreditJSONRequestBody defines body for PostCredit for application/json ContentType.
type PostCreditJSONRequestBody = Credit

// PostCreditEarlyRepaymentJSONRequestBody defines body for PostCreditEarlyRepayment for application/json ContentType.
type PostCreditEarlyRepaymentJSONRequestBody = EarlyRepayment

// PostCreditFullRepaymentJSONRequestBody defines body for PostCreditFullRepayment for application/json ContentType.
type PostCreditFullRepaymentJSONRequestBody = FullRepayment

// PostCreditPaymentPlanJSONRequestBody defines body for PostCreditPaymentPlan for application/json ContentType.
type PostCreditPaymentPlanJSONRequestBody = Rate