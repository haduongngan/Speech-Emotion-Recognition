package model

type Segment struct {
	Id      int    `json:"id" gorm:"primaryKey"`
	Emotion string `json:"emotion" gorm:"emotion"`
	SegNum  int    `json:"segNum" gorm:"segNum"`
	CallId  int    `json:"callId" gorm:"callId"`
}

type SegmentRepository interface {
	GetAll() ([]Segment, error)
	GetById(id int) (*Segment, error)
	GetByCallId(callId int) ([]Segment, error)
	CreateSegment(new *Segment) error
	DeleteSegment(id int) error
	GetEmotion(id int) (string, error)
}
