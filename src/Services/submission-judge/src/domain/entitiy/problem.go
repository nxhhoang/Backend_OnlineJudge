package domain

type Problem struct {
	ProblemId string   `bson:"_id" json:"problem_id,omitempty"`
	Index     string   `bson:"index" json:"index,omitempty"`
	Name      string   `bson:"name" json:"name,omitempty"`
	Points    int      `bson:"points" json:"points,omitempty"`
	Rating    int      `bson:"rating" json:"rating,omitempty"`
	Tags      []string `bson:"tags" json:"tags,omitempty"`
}
