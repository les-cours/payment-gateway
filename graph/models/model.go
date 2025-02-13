// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type ChargeAccountRequest struct {
	StudentID string `json:"studentID"`
	Code      string `json:"code"`
}

type ChargeAccountResponse struct {
	Amount    float64 `json:"amount"`
	NewAmount float64 `json:"newAmount"`
}

type Mutation struct {
}

type OperationStatus struct {
	Succeeded bool `json:"succeeded"`
}

type PayClassRoomRequest struct {
	ClassroomID string `json:"classroomID"`
	Month       int    `json:"month"`
}

type Query struct {
}
