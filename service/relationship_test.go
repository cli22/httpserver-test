package service

import (
	"testing"
	"httpserver-test/dao"
)

type relationListCase struct {
	uid          int
	relationship []dao.Relationship
}

var relations = []relationListCase{
	{1, []dao.Relationship{{Id: 1, Uid: 1, AnotherUid: 2, State: string(dao.Matched), Type: "relationship"}, {Id: 3, Uid: 1, AnotherUid: 3, State: string(dao.Liked), Type: "relationship"}}},
}

func TestListUserRelationship(t *testing.T) {
	for _, relation := range relations {
		v, _ := ListUserRelationship(relation.uid)
		if v[0] != relation.relationship[0] {
			t.Error(
				"For", relation.uid,
				"expected", relation.relationship,
				"got", v,
			)
		}
	}
}

type relationUpdateCase struct {
	uid          int
	anotherUid   int
	state        string
	relationship dao.Relationship
}

var relationUpdates = []relationUpdateCase{
	{1, 2, string(dao.Disliked), dao.Relationship{Id: 1, Uid: 1, AnotherUid: 2, State: string(dao.Disliked), Type: "relationship"}},
	{2, 1, string(dao.Liked), dao.Relationship{Id: 2, Uid: 2, AnotherUid: 1, State: string(dao.Liked), Type: "relationship"}},
	{1, 2, string(dao.Liked), dao.Relationship{Id: 1, Uid: 1, AnotherUid: 2, State: string(dao.Matched), Type: "relationship"}},
	{2, 1, string(dao.Liked), dao.Relationship{Id: 2, Uid: 2, AnotherUid: 1, State: string(dao.Matched), Type: "relationship"}},
	{1, 3, string(dao.Liked), dao.Relationship{Id: 3, Uid: 1, AnotherUid: 3, State: string(dao.Liked), Type: "relationship"}},
}

func TestUpdateRelationship(t *testing.T) {
	for _, relation := range relationUpdates {
		v, _ := UpdateRelationship(relation.uid, relation.anotherUid, relation.state)
		if v != relation.relationship {
			t.Error(
				"For uid", relation.uid,
				"For anotherUid", relation.anotherUid,
				"For state", relation.state,
				"expected", relation.relationship,
				"got", v,
			)
		}
	}
}
