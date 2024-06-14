package resolvers

import (
	"context"
	"errors"
	"github.com/les-cours/payment-gateway/api/payment"
	"github.com/les-cours/payment-gateway/graph/models"
	"github.com/les-cours/payment-gateway/types"
	"log"
)

func (r *queryResolver) GetAmount(ctx context.Context) (float64, error) {

	user, ok := ctx.Value("user").(*types.UserToken)
	if !ok || *user == (types.UserToken{}) {
		return 0, ErrPermissionDenied
	}

	res, err := r.PaymentClient.GetAmount(ctx, &payment.GetAmountRequest{
		StudentID: user.ID,
	})
	if err != nil {
		return 0, err
	}

	return float64(res.Amount), nil
}

func (r *mutationResolver) ChargeAccount(ctx context.Context, in models.ChargeAccountRequest) (*models.ChargeAccountResponse, error) {

	//user, ok := ctx.Value("user").(*types.UserToken)
	//if !ok || *user == (types.UserToken{}) {
	//	return nil, ErrPermissionDenied
	//}
	//
	//if user.UserType != "student" {
	//	return nil, ErrPermissionDenied
	//}

	_, err := r.PaymentClient.ChargeAccount(ctx, &payment.ChargeAccountRequest{
		StudentID: in.StudentID,
		Code:      in.Code,
	})

	if err != nil {
		log.Println("Err: ", err)
		return nil, ErrApi(err)
	}

	res, err := r.PaymentClient.GetAmount(ctx, &payment.GetAmountRequest{
		StudentID: in.StudentID,
	})
	if err != nil {
		log.Println("err: ", err.Error())
		return nil, errors.New(err.Error())
	}
	res2, err := r.PaymentClient.GetAmountByCode(ctx, &payment.GetAmountRequest{
		StudentID: in.Code,
	})
	if err != nil {
		log.Println("err2: ", err.Error())
		return nil, errors.New(err.Error())
	}

	return &models.ChargeAccountResponse{
		Amount:    float64(res2.Amount),
		NewAmount: float64(res.Amount),
	}, nil
}
