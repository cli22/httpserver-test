package service

import (
	"sync"

	"httpserver-test/dao"
	"httpserver-test/log"
)

func ListUserRelationship(uid int) (relationships []dao.Relationship, err error) {
	err = dao.Db.Model(&relationships).Where("relationship.uid=?", uid).OrderExpr("relationship.other_uid ASC").Select()
	if err != nil {
		log.Warning.Println("ListUserRelationship SELECT error: ", err)
	}

	log.Info.Println("ListUserRelationship SELECT success, result relationships: ", relationships)

	return
}

func UpdateRelationship(uid, otherUid int, state string) (relationship dao.Relationship, err error) {
	var (
		mu              sync.Mutex
		relationshipTmp dao.Relationship
	)
	// Todo change to transaction, not lock
	mu.Lock()
	defer mu.Unlock()

	err = dao.Db.Model(&relationship).Where("relationship.uid=?", uid).Where("relationship.other_uid=?", otherUid).Select()
	if err != nil {
		log.Warning.Println("UpdateRelationship SELECT uid error: ", err)
	}

	err = dao.Db.Model(&relationshipTmp).Where("relationship.uid=?", otherUid).Where("relationship.other_uid=?", uid).Select()
	if err != nil {
		log.Warning.Println("UpdateRelationship SELECT otherUid error: ", err)
	}

	log.Info.Println("UpdateRelationship SELECT uid success, result relationship: ", relationship)
	log.Info.Println("UpdateRelationship SELECT otherUid success, result relationshipTmp: ", relationshipTmp)

	err = ProcRelationship(state, &relationship, &relationshipTmp, uid, otherUid)
	if err != nil {
		log.Warning.Println("UpdateRelationship ProcRelationship error: ", err)
	}

	log.Info.Println("UpdateRelationship success, result: ", relationship)

	return
}

func ProcRelationship(state string, relationship, relationshipTmp *dao.Relationship, uid, otherUid int) (err error) {
	// not exists, it's a new relationship
	if relationship.Uid == 0 {

		states := ProcState(state, relationship.State, relationshipTmp.State, false)

		relationship = &dao.Relationship{Uid: int64(uid), OtherUid: int64(otherUid), State: states[0], Type: "relationship"}
		relationshipTmp = &dao.Relationship{Uid: int64(otherUid), OtherUid: int64(uid), State: states[1], Type: "relationship"}

		err = dao.Db.Insert(relationship, relationshipTmp)
		if err != nil {
			log.Warning.Println("UpdateRelationship INSERT relationship, relationshipTmp error: ", err)
		}

		log.Info.Println("UpdateRelationship INSERT relationship, relationshipTmp success, result: ", relationship, relationshipTmp)

	} else {
		states := ProcState(state, relationship.State, relationshipTmp.State, true)

		relationship.State = states[0]
		relationshipTmp.State = states[1]

		_, err := dao.Db.Model(relationship, relationshipTmp).Column("state").Update()
		if err != nil {
			log.Warning.Println("UpdateRelationship UPDATE relationship,relationshipTmp state error: ", err)
		}

		log.Info.Println("UpdateRelationship UPDATE relationship, relationshipTmp state success, result: ", relationship, relationshipTmp)

	}

	return
}

func ProcState(state, state1, state2 string, exists bool) (states [2]string) {
	if state == dao.Liked {
		states = ProcLiked(state1, state2, exists)
	} else {
		states = ProcDisliked(state1, state2, exists)
	}

	log.Info.Println("ProcState success, result: ", states)

	return
}

func ProcLiked(state1, state2 string, exists bool) (states [2]string) {
	switch state1 {
	case dao.Liked:
		states[0] = state1
		states[1] = state2

	case dao.Disliked:
		if state2 == dao.Liked {
			states[0] = dao.Matched
			states[1] = dao.Matched
		} else {
			states[0] = dao.Liked
			states[1] = state2
		}

	case dao.Matched:
		states[0] = state1
		states[1] = state2

	case dao.Default:
		if exists {
			if state2 == dao.Liked {
				states[0] = dao.Matched
				states[1] = dao.Matched
			} else {
				states[0] = dao.Liked
				states[1] = state2
			}
		} else {
			states[0] = dao.Liked
			states[1] = dao.Default
		}

	}

	return
}

func ProcDisliked(state1, state2 string, exists bool) (states [2]string) {
	switch state1 {
	case dao.Liked:
		states[0] = dao.Disliked
		states[1] = state2

	case dao.Disliked:
		states[0] = state1
		states[1] = state2

	case dao.Matched:
		states[0] = dao.Disliked
		states[1] = dao.Liked

	case dao.Default:
		if exists {
			states[0] = dao.Disliked
			states[1] = state2
		} else {
			states[0] = dao.Disliked
			states[1] = dao.Default
		}

	}

	return
}
