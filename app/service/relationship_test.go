package service

import (
	"testing"

	"httpserver-test/app/dao"
	"httpserver-test/app/entity"
)

type relationUpdateCase struct {
	uid          string
	otherUid     string
	state        string
	relationship *entity.Relationship
}

var relationUpdates = []relationUpdateCase{
	{"1", "2", dao.Disliked, &entity.Relationship{Uid: "1", OtherUid: "2", State: dao.Disliked, Type: dao.RelationshipType}},
	{"2", "1", dao.Liked, &entity.Relationship{Uid: "2", OtherUid: "1", State: dao.Liked, Type: dao.RelationshipType}},
	{"1", "2", dao.Liked, &entity.Relationship{Uid: "1", OtherUid: "2", State: dao.Matched, Type: dao.RelationshipType}},
	{"2", "1", dao.Liked, &entity.Relationship{Uid: "2", OtherUid: "1", State: dao.Matched, Type: dao.RelationshipType}},
	{"1", "3", dao.Liked, &entity.Relationship{Uid: "1", OtherUid: "3", State: dao.Liked, Type: dao.RelationshipType}},
}

func TestRelationship_UpdateRelationship(t *testing.T) {
	relationship := NewRelationship()
	for _, relation := range relationUpdates {
		inputRelation := &entity.Relationship{Uid: relation.uid, OtherUid: relation.otherUid, State: relation.state}
		v, _ := relationship.UpdateRelationship(inputRelation)
		if v.Uid != relation.relationship.Uid || v.OtherUid != relation.relationship.OtherUid || v.State != relation.relationship.State {
			t.Error(
				"For uid", relation.uid,
				"For otherUid", relation.otherUid,
				"For state", relation.state,
				"expected", relation.relationship,
				"got", v,
			)
		}
	}
}
