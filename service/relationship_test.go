package service

import (
	"testing"
	"httpserver-test/dao"
)

// when relationship changed, this testcase cannot work
//type relationListCase struct {
//	uid          int
//	relationship []dao.Relationship
//}
//
//var relations = []relationListCase{
//	{1, []dao.Relationship{{Id: 1, Uid: 1, OtherUid: 2, State: dao.Matched, Type: dao.RelationshipType}, {Id: 3, Uid: 1, OtherUid: 3, State: dao.Liked, Type: dao.RelationshipType}}},
//}
//
//func TestGetUserRelationship(t *testing.T) {
//	for _, relation := range relations {
//		v, _ := GetUserRelationship(relation.uid)
//		for i, vi := range v {
//			if vi.Uid != relation.relationship[i].Uid || vi.OtherUid != relation.relationship[i].OtherUid || vi.State != relation.relationship[i].State {
//				t.Error(
//					"For", relation.uid,
//					"expected", relation.relationship,
//					"got", v,
//				)
//			}
//		}
//	}
//}

type relationUpdateCase struct {
	uid          int
	otherUid     int
	state        string
	relationship dao.Relationship
}

var relationUpdates = []relationUpdateCase{
	{1, 2, dao.Disliked, dao.Relationship{Id: 1, Uid: 1, OtherUid: 2, State: dao.Disliked, Type: dao.RelationshipType}},
	{2, 1, dao.Liked, dao.Relationship{Id: 2, Uid: 2, OtherUid: 1, State: dao.Liked, Type: dao.RelationshipType}},
	{1, 2, dao.Liked, dao.Relationship{Id: 1, Uid: 1, OtherUid: 2, State: dao.Matched, Type: dao.RelationshipType}},
	{2, 1, dao.Liked, dao.Relationship{Id: 2, Uid: 2, OtherUid: 1, State: dao.Matched, Type: dao.RelationshipType}},
	{1, 3, dao.Liked, dao.Relationship{Id: 3, Uid: 1, OtherUid: 3, State: dao.Liked, Type: dao.RelationshipType}},
}

func TestUpdateRelationship(t *testing.T) {
	for _, relation := range relationUpdates {
		v, _ := UpdateRelationship(relation.uid, relation.otherUid, relation.state)
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
