package domain

type A struct {
	ID        int64   `json:"id" gorm:"primaryKey"`
	ShortData *string `json:"short_data"`
	LongData  *string `json:"long_data"`
}

// TODO: input validation pipeline
type ACreateBody struct {
	ShortData *string `json:"short_data"`
	LongData  *string `json:"long_data"`
}

func CreateAFromBody(body ACreateBody) A {
	return A{
		ID:        GenID(),
		ShortData: body.ShortData,
		LongData:  body.LongData,
	}
}
