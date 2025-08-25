package requests

type IdParamRequest struct {
	Id uint `param:"id" binding:"required"`
}
