package resolvers

import (
	"context"
	"github.com/les-cours/payment-gateway/api/payment"
	"log"
)

func (r *mutationResolver) GeneratePaymentCode(ctx context.Context, amount float64) (*string, error) {

	res, err := r.PaymentClient.GeneratePaymentCode(ctx, &payment.GeneratePaymentCodeRequest{
		Amount: float32(amount),
	})

	if err != nil {
		log.Println("Err: ", err)
		return nil, ErrApi(err)
	}

	return &res.Message, nil
}
