package service

import (
	"sync"

	"github.com/httpserver-test/app/dao"
	"github.com/httpserver-test/app/entity"
	"github.com/httpserver-test/log"
)

var RelationshipSvc *Relationship

type Relationship struct {
	myDaoRelationship *dao.MyRelationship
}

func NewRelationship() *Relationship {
	relationship := new(Relationship)
	relationship.myDaoRelationship = dao.NewMyRelationship()
	return relationship
}

func (r *Relationship) GetUserRelationship(data *entity.Relationship) (relationships []*entity.Relationship, err error) {
	relationships, err = r.myDaoRelationship.GetByUid(data)
	if err != nil {
		log.Warning.Println("GetUserRelationship error: ", err)
	}

	return
}

func (r *Relationship) UpdateRelationship(data *entity.Relationship) (relationship *entity.Relationship, err error) {
	var mu sync.Mutex

	// Todo change to transaction, not lock
	mu.Lock()
	defer mu.Unlock()

	relationship, err = r.myDaoRelationship.GetByUidOtherUid(data)

	if err != nil {
		log.Warning.Println("SELECT uid error: ", err)
	}

	relationshipTmp := new(entity.Relationship)
	relationshipTmp.Uid = data.OtherUid
	relationshipTmp.OtherUid = data.Uid

	relationshipTmp, err = r.myDaoRelationship.GetByUidOtherUid(relationshipTmp)

	if err != nil {
		log.Warning.Println("SELECT other_uid error: ", err)
	}

	log.Info.Println("SELECT result: ", relationship, relationshipTmp)

	err = r.ProcRelationship(data, relationship, relationshipTmp)
	if err != nil {
		log.Warning.Println("ProcRelationship error: ", err)
	}

	log.Info.Println("UpdateRelationship result: ", relationship)

	return
}

func (r *Relationship) ProcRelationship(data, relationship, relationshipTmp *entity.Relationship) (err error) {
	// not exists, it's a new relationship
	if relationship.State == "" {

		// insert relationship
		relationship.State = data.State
		relationship.Uid = data.Uid
		relationship.OtherUid = data.OtherUid

		relationship, err = r.myDaoRelationship.Add(relationship)
		if err != nil {
			log.Warning.Println("INSERT relationship error: ", err)
		}

		relationshipTmp.State = dao.Default
		relationshipTmp.Uid = data.OtherUid
		relationshipTmp.OtherUid = data.Uid
		relationship, err = r.myDaoRelationship.Add(relationshipTmp)
		if err != nil {
			log.Warning.Println("INSERT relationshipTmp error: ", err)
		}

		log.Info.Println("ProcRelationship INSERT result: ", relationship, relationshipTmp)

		// update relationship
	} else {
		// 获取改变状态后的结果
		states := r.ProcState(data.State, relationship.State, relationshipTmp.State)

		relationship.State = states[0]
		relationshipTmp.State = states[1]

		relationship, err = r.myDaoRelationship.UpdateRelationshipByState(relationship)
		if err != nil {
			log.Warning.Println("UPDATE relationship error: ", err)
		}

		relationshipTmp, err = r.myDaoRelationship.UpdateRelationshipByState(relationshipTmp)
		if err != nil {
			log.Warning.Println("UPDATE relationshipTmp error: ", err)
		}

		log.Info.Println("ProcRelationship UPDATE result: ", relationship, relationshipTmp)

	}

	return
}

// 此函数负责获得state1和state2的最终状态，结果states中：
// state1 = states[0],  state2 = states[1],
func (r *Relationship) ProcState(state, state1, state2 string) (states []string) {
	states = []string{state1, state2}

	// 此处处理state状态为liked时逻辑。
	// 只有当state1是disliked或default时，输入liked才会改变状态；
	// 其余状态不变
	if state == dao.Liked {
		if state1 == dao.Disliked || state1 == dao.Default {
			if state2 == dao.Liked {
				states[0] = dao.Matched
				states[1] = dao.Matched
			} else {
				states[0] = dao.Liked
			}
		}

		// 此处处理state状态为disliked时逻辑。
		// 只有当state1是liked或default时，state1变成disliked；
	} else {
		if state1 == dao.Liked || state1 == dao.Default {
			states[0] = dao.Disliked
		} else if state1 == dao.Matched {
			states[0] = dao.Disliked
			states[1] = dao.Liked
		}
	}

	log.Info.Println("ProcState success, result: ", states)

	return
}
