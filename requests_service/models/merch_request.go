package models

type MerchRequest struct {
	Id          int
	UserId      int
	UserMessage int
	HrId        int
	HrMessage   string
	MerchId     int
	StatusKey   string
	ImageUrl    string
}
