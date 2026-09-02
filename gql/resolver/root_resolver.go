package resolver

type RootResolver struct{}

func (*RootResolver) Mutation() *mutationResolver {
	return &mutationResolver{}
}

func (*RootResolver) Query() *queryResolver {
	return &queryResolver{}
}
