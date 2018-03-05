package entity

type Relationship struct {
	Id       int `json:"-"`
	Uid      string `json:"-"`
	OtherUid string `json:"user_id"`
	State    string `json:"state"`
	Type     string `json:"type"`
}
