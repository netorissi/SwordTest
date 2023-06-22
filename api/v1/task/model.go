package task

type createRequest struct {
	Summary string `json:"summary" validate:"required,max=2500"`
}

type updateRequest struct {
	Summary string `json:"summary" validate:"required,max=2500"`
}
