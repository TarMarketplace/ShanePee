package domain

type A struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	LongData string `json:"long_data"`
}

// TODO: input validation pipeline
type ACreateBody struct {
	LongData string `json:"long_data"`
}

func CreateAFromBody(body ACreateBody) A {
	return A{
		ID:       GenID(),
		LongData: body.LongData,
	}
}
