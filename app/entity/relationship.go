package entity

type Relationship struct {
	Uid      string `json:"-"`
	OtherUid string `json:"user_id"`
	State    string `json:"state"`
	Type     string `json:"type"`
}
