package domain

type SourceCode struct {
	SourceCodeId string `json:"id,omitempty" bson:"_id"`
	Language     string `json:"language,omitempty" bson:"language"`
	CreatedAt    string `json:"create_at,omitempty" bson:"create_at"`
}
