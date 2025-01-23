package domain

type A struct {
	ID       string `json:"id" gorm:"primaryKey"`
	LongData string `json:"long_data"`
}

// TODO: input validation pipeline
type ACreateBody struct {
	LongData string `json:"long_data"`
}

func CreateAFromBody(body ACreateBody) A {
	return A{
		// TODO: ID generation
		ID:       "1",
		LongData: body.LongData,
	}
}
