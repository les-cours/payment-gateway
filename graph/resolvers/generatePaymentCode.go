package resolvers

import (
	"context"
	"github.com/les-cours/payment-gateway/api/payment"
	"github.com/les-cours/payment-gateway/types"
	"log"
)

func (r *mutationResolver) GeneratePaymentCode(ctx context.Context, amount float64) (*string, error) {

	var user *types.UserToken
	if user, _ = ctx.Value("user").(*types.UserToken); user == nil {
		return nil, ErrPermissionDenied
	}

	if !user.Read.PAYMENT {
		return nil, ErrPermissionDenied
	}

	res, err := r.PaymentClient.GeneratePaymentCode(ctx, &payment.GeneratePaymentCodeRequest{
		Amount: float32(amount),
	})

	if err != nil {
		log.Println("Err: ", err)
		return nil, ErrApi(err)
	}

	return &res.Message, nil
}
