package resolvers

import (
	"context"
	"github.com/les-cours/payment-gateway/api/payment"
	"github.com/les-cours/payment-gateway/graph/models"
	"log"
)

func (r *queryResolver) GetAmount(ctx context.Context, studentID string) (int, error) {

	var x = 1500
	return x, nil
	//
	////var user *types.UserToken
	////if user, _ = ctx.Value("user").(*types.UserToken); user == nil {
	////	return nil, ErrPermissionDenied
	////}
	////
	////if !user.Read.USER {
	////	return nil, ErrPermissionDenied
	////}
	//
	//var students []*models.Student
	//res, err := r.UserClient.GetStudents(ctx, &users.GetStudentsRequest{
	//	FilterType:  in.FilterType,
	//	FilterValue: in.FilterValue,
	//})
	//if err != nil {
	//	return nil, ErrApi(err)
	//}
	//for _, student := range res.Students {
	//	students = append(students, gprcToGraph.Student(student))
	//}
	//
	//return students, nil
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
