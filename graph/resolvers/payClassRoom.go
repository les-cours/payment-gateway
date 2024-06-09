package resolvers

import (
	"context"
	"github.com/les-cours/payment-gateway/api/payment"
	"github.com/les-cours/payment-gateway/graph/models"
	"github.com/les-cours/payment-gateway/types"
	"log"
)

func (r *mutationResolver) PayClassRoom(ctx context.Context, in models.PayClassRoomRequest) (*string, error) {

	var user *types.UserToken
	user, ok := ctx.Value("user").(*types.UserToken)
	if !ok || *user == (types.UserToken{}) {
		return nil, ErrPermissionDenied
	}

	res, err := r.PaymentClient.PayClassRoom(ctx, &payment.PayClassRoomRequest{
		ClassRoomID: in.ClassroomID,
		StudentID:   user.ID,
		Month:       int32(in.Month),
	})

	if err != nil {
		log.Println("Err: ", err)
		return nil, ErrApi(err)
	}

	return &res.Message, nil
}
