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
	{1, []dao.Relationship{{Id: 1, Uid: 1, OtherUid: 2, State: dao.Matched, Type: "relationship"}, {Id: 3, Uid: 1, OtherUid: 3, State: dao.Liked, Type: "relationship"}}},
}

func TestListUserRelationship(t *testing.T) {
	for _, relation := range relations {
		v, _ := ListUserRelationship(relation.uid)
		for i, vi := range v {
			if vi != relation.relationship[i] {
				t.Error(
					"For", relation.uid,
					"expected", relation.relationship,
					"got", v,
				)
			}
		}
	}
}

type relationUpdateCase struct {
	uid          int
	otherUid     int
	state        string
	relationship dao.Relationship
}

var relationUpdates = []relationUpdateCase{
	{1, 2, dao.Disliked, dao.Relationship{Id: 1, Uid: 1, OtherUid: 2, State: dao.Disliked, Type: "relationship"}},
	{2, 1, dao.Liked, dao.Relationship{Id: 2, Uid: 2, OtherUid: 1, State: dao.Liked, Type: "relationship"}},
	{1, 2, dao.Liked, dao.Relationship{Id: 1, Uid: 1, OtherUid: 2, State: dao.Matched, Type: "relationship"}},
	{2, 1, dao.Liked, dao.Relationship{Id: 2, Uid: 2, OtherUid: 1, State: dao.Matched, Type: "relationship"}},
	{1, 3, dao.Liked, dao.Relationship{Id: 3, Uid: 1, OtherUid: 3, State: dao.Liked, Type: "relationship"}},
}

func TestUpdateRelationship(t *testing.T) {
	for _, relation := range relationUpdates {
		v, _ := UpdateRelationship(relation.uid, relation.otherUid, relation.state)
		if v != relation.relationship {
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
