package resolvers

type QueryResolver struct {
	server *Server
}

func (r *QueryResolver) Accounts(ctx context.Context, pagination *PaginationInput, id *string) ([]*Account, error) {
	return r.server.AccountClient.GetAccounts(ctx)
}

func (r *QueryResolver) Products(ctx context.Context, pagination *PaginationInput,query *string, id *string) ([]*Product, error) {
	return r.server.AccountClient.GetAccounts(ctx)
}