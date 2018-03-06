package dao

import (
	"strconv"

	"github.com/httpserver-test/app/entity"
	"github.com/httpserver-test/log"
)

type Relationship struct {
	Id       int    `json:"id"`
	Uid      int    `json:"uid"`
	OtherUid int    `json:"user_id"`
	State    string `json:"state"`
	Type     string `json:"type"`
}

type RelationshipState = string //typedef

const (
	Liked    RelationshipState = "liked"
	Disliked RelationshipState = "disliked"
	Matched  RelationshipState = "matched"
	Default  RelationshipState = "default"
)

type RelationshipTypeGroup = string

const (
	RelationshipType RelationshipTypeGroup = "relationship"
)

type MyRelationship struct{}

func NewMyRelationship() *MyRelationship {
	return &MyRelationship{}
}

func (this *MyRelationship) resToRelationship(res []*Relationship) (relationships []*entity.Relationship, err error) {
	if len(res) == 0 {
		return nil, nil
	}

	for _, r := range res {
		relationship := new(entity.Relationship)
		relationship.Uid = strconv.Itoa(r.Uid)
		relationship.OtherUid = strconv.Itoa(r.OtherUid)
		relationship.State = r.State
		relationship.Type = r.Type

		relationships = append(relationships, relationship)
	}

	return relationships, nil
}

func (this *MyRelationship) GetByUid(data *entity.Relationship) (relationships []*entity.Relationship, err error) {
	res := make([]*Relationship, 0)

	err = Db.Model(&res).Where("relationship.uid=?", data.Uid).Order("relationship.other_uid").Select()
	if err != nil {
		log.Warning.Println("SELECT error: ", err)
	}

	relationships, err = this.resToRelationship(res)

	log.Info.Println("SELECT result: ", relationships)

	return
}

func (this *MyRelationship) GetByUidOtherUid(data *entity.Relationship) (relationship *entity.Relationship, err error) {
	daoRelationship := &Relationship{}

	err = Db.Model(daoRelationship).Where("relationship.uid=?", data.Uid).Where("relationship.other_uid=?", data.OtherUid).Select()

	if err != nil {
		log.Warning.Println("SELECT error: ", err)
	}

	res := make([]*Relationship, 0)
	res = append(res, daoRelationship)
	relationships, err := this.resToRelationship(res)

	if len(relationships) == 1 {
		relationship = relationships[0]
	}

	log.Info.Println("SELECT result: ", relationship)

	return
}

func (this *MyRelationship) UpdateRelationshipByState(data *entity.Relationship) (relationship *entity.Relationship, err error) {
	daoRelationship := &Relationship{}
	daoRelationship.State = data.State

	_, err = Db.Model(daoRelationship).Where("relationship.uid=?", data.Uid).Where("relationship.other_uid=?", data.OtherUid).Column("state").Update()
	if err != nil {
		log.Warning.Println("Update error: ", err)
	}

	res := make([]*Relationship, 0)
	res = append(res, daoRelationship)
	relationships, err := this.resToRelationship(res)

	if len(relationships) == 1 {
		relationship = relationships[0]
	}

	log.Info.Println("Update result: ")

	return
}

func (this *MyRelationship) Add(data *entity.Relationship) (relationship *entity.Relationship, err error) {
	daoRelationship := &Relationship{}
	daoRelationship.Uid, _ = strconv.Atoi(data.Uid)
	daoRelationship.OtherUid, _ = strconv.Atoi(data.OtherUid)
	daoRelationship.State = data.State
	daoRelationship.Type = RelationshipType

	err = Db.Insert(daoRelationship)
	if err != nil {
		log.Warning.Println("INSERT db error: ", err)
	}

	res := make([]*Relationship, 0)
	res = append(res, daoRelationship)

	relationships, err := this.resToRelationship(res)
	if len(relationships) == 1 {
		relationship = relationships[0]
	}

	log.Info.Println("Update result: ", relationship)

	return relationship, err
}
