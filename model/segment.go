package model

type Segment struct {
	Id      int    `json:"id" gorm:"primaryKey"`
	Emotion string `json:"emotion" gorm:"emotion"`
	SegNum  int    `json:"segNum" gorm:"segNum"`
	CallId  int    `json:"callId" gorm:"callId"`
}
