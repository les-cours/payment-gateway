package resolvers

import (
	"context"
	"github.com/les-cours/payment-gateway/api/payment"
	"github.com/les-cours/payment-gateway/graph/models"
	"log"
)

func (r *mutationResolver) PayClassRoom(ctx context.Context, in models.PayClassRoomRequest) (*string, error) {

	//var user *types.UserToken
	//if user, _ = ctx.Value("user").(*types.UserToken); user == nil {
	//	return nil, ErrPermissionDenied
	//}
	//
	//if !user.Read.PAYMENT {
	//	return nil, ErrPermissionDenied
	//}

	res, err := r.PaymentClient.PayClassRoom(ctx, &payment.PayClassRoomRequest{
		ClassRoomID: in.ClassroomID,
		StudentID:   "5a7e2874-0b2c-4acd-ba08-155a03fe6a0c", //user.ID
		Month:       int32(in.Month),
	})

	if err != nil {
		log.Println("Err: ", err)
		return nil, ErrApi(err)
	}

	return &res.Message, nil
}
