package resolvers

import (
	"context"
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

func (r *mutationResolver) ChargeAccount(ctx context.Context, in models.ChargeAccountRequest) (*models.OperationStatus, error) {

	res, err := r.PaymentClient.ChargeAccount(ctx, &payment.ChargeAccountRequest{
		StudentID: in.StudentID,
		Code:      in.Code,
	})

	if err != nil {
		log.Println("Err: ", err)
		return nil, ErrApi(err)
	}

	log.Println("res : ", res.GetMessage())
	return &models.OperationStatus{
		Succeeded: true,
	}, nil
}
