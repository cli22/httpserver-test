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

	// todo refactor
	//  func ProcRelationship(*relationship, *relationshipTmp) error {
	//   	...
	// 		ProcState(state, relationship.State, relationshipTmp.State) ([]state, error)
	//
	//		...
	// }

	//  func ProcState(state, state1, state2) ([]states, error) {
	// 		if state == like {
	//			states = ProcLike(state1, state2)
	//		} else if state == unlike {
	//			states = ProcUnlike(state1, state2)
	//		}
	//  }
	switch state {
	case dao.Liked:
		switch relationship.State {
		case dao.Liked:

		case dao.Disliked:
			if relationshipTmp.State == dao.Liked {
				relationship.State = dao.Matched
				relationshipTmp.State = dao.Matched

				_, err := dao.Db.Model(&relationship, &relationshipTmp).Column("state").Update()
				if err != nil {
					log.Warning.Println("UpdateRelationship UPDATE relationship, relationshipTmp state error: ", err)
				}

				log.Info.Println("UpdateRelationship UPDATE relationship,relationshipTmp success, result: ", relationship, relationshipTmp)
			} else {
				relationship.State = dao.Liked

				_, err := dao.Db.Model(&relationship).Column("state").Returning("*").Update()
				if err != nil {
					log.Warning.Println("UpdateRelationship UPDATE relationship state error: ", err)
				}

				log.Info.Println("UpdateRelationship UPDATE relationship success, result: ", relationship)
			}

		case dao.Matched:

		case dao.Default:
			if relationship.Uid == 0 {
				relationship = dao.Relationship{Uid: int64(uid), OtherUid: int64(otherUid), State: state, Type: "relationship"}
				relationshipTmp = dao.Relationship{Uid: int64(otherUid), OtherUid: int64(uid), State: dao.Default, Type: "relationship"}

				err = dao.Db.Insert(&relationship, &relationshipTmp)
				if err != nil {
					log.Warning.Println("UpdateRelationship INSERT relationship, relationshipTmp error: ", err)
				}

				log.Info.Println("UpdateRelationship INSERT relationship, relationshipTmp success, result: ", relationship, relationshipTmp)
			} else {
				if relationshipTmp.State == dao.Liked {
					relationship.State = dao.Matched
					relationshipTmp.State = dao.Matched

					_, err := dao.Db.Model(&relationship, &relationshipTmp).Column("state").Update()
					if err != nil {
						log.Warning.Println("UpdateRelationship UPDATE relationship, relationshipTmp state error: ", err)
					}

					log.Info.Println("UpdateRelationship UPDATE relationship, relationshipTmp state success, result: ", relationship, relationshipTmp)
				} else {
					relationship.State = dao.Liked

					_, err := dao.Db.Model(&relationship).Column("state").Returning("*").Update()
					if err != nil {
						log.Warning.Println("UpdateRelationship UPDATE relationship state error: ", err)
					}

					log.Info.Println("UpdateRelationship UPDATE relationship state success, result: ", relationship)
				}

			}
		}

	case dao.Disliked:
		switch relationship.State {
		case dao.Liked:
			relationship.State = dao.Disliked

			_, err := dao.Db.Model(&relationship).Column("state").Returning("*").Update()
			if err != nil {
				log.Warning.Println("UpdateRelationship UPDATE relationship state error: ", err)
			}

			log.Info.Println("UpdateRelationship UPDATE relationship success, result: ", relationship)

		case dao.Disliked:

		case dao.Matched:
			relationship.State = dao.Disliked
			relationshipTmp.State = dao.Liked

			_, err := dao.Db.Model(&relationship, &relationshipTmp).Column("state").Update()
			if err != nil {
				log.Warning.Println("UpdateRelationship UPDATE relationship,relationshipTmp state error: ", err)
			}

			log.Info.Println("UpdateRelationship UPDATE relationship, relationshipTmp state success, result: ", relationship, relationshipTmp)

		case dao.Default:
			if relationship.Uid == 0 {
				relationship = dao.Relationship{Uid: int64(uid), OtherUid: int64(otherUid), State: state, Type: "relationship"}
				relationshipTmp = dao.Relationship{Uid: int64(otherUid), OtherUid: int64(uid), State: dao.Default, Type: "relationship"}

				err = dao.Db.Insert(&relationship, &relationshipTmp)
				if err != nil {
					log.Warning.Println("UpdateRelationship INSERT relationship, relationshipTmp error: ", err)
				}

				log.Info.Println("UpdateRelationship INSERT relationship, relationshipTmp success, result: ", relationship, relationshipTmp)
			} else {
				relationship.State = dao.Disliked

				_, err := dao.Db.Model(&relationship).Column("state").Returning("*").Update()
				if err != nil {
					log.Warning.Println("UpdateRelationship UPDATE relationship state error: ", err)
				}

				log.Info.Println("UpdateRelationship UPDATE relationship state success, result: ", relationship)
			}
		}
	}

	return
}
