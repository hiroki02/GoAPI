package request

type GetTask struct {
	ID string `param:"id" validated:"required"`
}

type ListTask struct {
}
