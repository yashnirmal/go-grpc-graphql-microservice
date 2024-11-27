package resolvers

import "context"

type MutationResolver struct {
	server *Server
}

// func (r *MutationResolver) createAccount(ctx context.Context, input AccountInput) (*Account, error) {
// 	return r.server.AccountClient.CreateAccount(ctx, input)
// }

// func (r *MutationResolver) createProduct(ctx context.Context, input ProductInput) (*Product, error) {
// 	return r.server.AccountClient.CreateAccount(ctx, input)
// }

// func (r *MutationResolver) createOrder(ctx context.Context, input OrderInput) (*Order, error) {
// 	return r.server.AccountClient.CreateAccount(ctx, input)
// }