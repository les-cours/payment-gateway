package resolvers

import (
	"github.com/les-cours/payment-gateway/api/payment"
	"github.com/les-cours/payment-gateway/graph"
)

type Resolver struct {
	PaymentClient payment.PaymentServiceClient
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
