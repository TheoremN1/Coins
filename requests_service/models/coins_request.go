package models

type CoinsRequest struct {
	Id            int
	UserId        int
	UserMessage   int
	HrId          int
	HrMessage     string
	AchievementId int
	StatusKey     string
}
