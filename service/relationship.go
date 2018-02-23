package service

import (
	"sync"

	"httpserver-test/dao"
	"httpserver-test/log"
)

func ListUserRelationship(uid int) (relationships []dao.Relationship, err error) {
	err = dao.Db.Model(&relationships).Where("relationship.uid=?", uid).OrderExpr("relationship.another_uid ASC").Select()
	if err != nil {
		log.Warning.Println("ListUserRelationship SELECT error: ", err)
	}
	log.Info.Println("ListUserRelationship SELECT success, result relationships: ", relationships)
	return
}

func UpdateRelationship(uid, anotherUid int, state string) (relationshipUid dao.Relationship, err error) {
	var (
		mu                     sync.Mutex
		relationshipAnotherUid dao.Relationship
	)
	// Todo change to transaction, not lock
	mu.Lock()
	defer mu.Unlock()
	err = dao.Db.Model(&relationshipUid).Where("relationship.uid=?", uid).Where("relationship.another_uid=?", anotherUid).Select()
	if err != nil {
		log.Warning.Println("UpdateRelationship SELECT uid error: ", err)
	}
	err = dao.Db.Model(&relationshipAnotherUid).Where("relationship.uid=?", anotherUid).Where("relationship.another_uid=?", uid).Select()
	if err != nil {
		log.Warning.Println("UpdateRelationship SELECT anotherUid error: ", err)
	}
	log.Info.Println("UpdateRelationship SELECT uid success, result relationshipUid: ", relationshipUid)
	log.Info.Println("UpdateRelationship SELECT anotherUid success, result relationshipAnotherUid: ", relationshipAnotherUid)
	switch state {
	case string(dao.Liked):
		switch relationshipUid.State {
		case string(dao.Liked):
		case string(dao.Disliked):
			if relationshipAnotherUid.State == string(dao.Liked) {
				relationshipUid.State = string(dao.Matched)
				relationshipAnotherUid.State = string(dao.Matched)
				_, err := dao.Db.Model(&relationshipUid, &relationshipAnotherUid).Column("state").Update()
				if err != nil {
					log.Warning.Println("UpdateRelationship UPDATE relationshipUid, relationshipAnotherUid state error: ", err)
				}
				log.Info.Println("UpdateRelationship UPDATE relationshipUid,relationshipAnotherUid success, result: ", relationshipUid, relationshipAnotherUid)
			} else {
				relationshipUid.State = string(dao.Liked)
				_, err := dao.Db.Model(&relationshipUid).Column("state").Returning("*").Update()
				if err != nil {
					log.Warning.Println("UpdateRelationship UPDATE relationshipUid state error: ", err)
				}
				log.Info.Println("UpdateRelationship UPDATE relationshipUid success, result: ", relationshipUid)
			}
		case string(dao.Matched):
		case string(dao.Default):
			if relationshipUid.Uid == 0 {
				relationshipUid = dao.Relationship{Uid: int64(uid), AnotherUid: int64(anotherUid), State: state, Type: "relationship"}
				relationshipAnotherUid = dao.Relationship{Uid: int64(anotherUid), AnotherUid: int64(uid), State: string(dao.Default), Type: "relationship"}
				err = dao.Db.Insert(&relationshipUid, &relationshipAnotherUid)
				if err != nil {
					log.Warning.Println("UpdateRelationship INSERT relationshipUid, relationshipAnotherUid error: ", err)
				}
				log.Info.Println("UpdateRelationship INSERT relationshipUid, relationshipAnotherUid success, result: ", relationshipUid, relationshipAnotherUid)
			} else {
				if relationshipAnotherUid.State == string(dao.Liked) {
					relationshipUid.State = string(dao.Matched)
					relationshipAnotherUid.State = string(dao.Matched)
					_, err := dao.Db.Model(&relationshipUid, &relationshipAnotherUid).Column("state").Update()
					if err != nil {
						log.Warning.Println("UpdateRelationship UPDATE relationshipUid, relationshipAnotherUid state error: ", err)
					}
					log.Info.Println("UpdateRelationship UPDATE relationshipUid, relationshipAnotherUid state success, result: ", relationshipUid, relationshipAnotherUid)
				} else {
					relationshipUid.State = string(dao.Liked)
					_, err := dao.Db.Model(&relationshipUid).Column("state").Returning("*").Update()
					if err != nil {
						log.Warning.Println("UpdateRelationship UPDATE relationshipUid state error: ", err)
					}
					log.Info.Println("UpdateRelationship UPDATE relationshipUid state success, result: ", relationshipUid)
				}

			}
		}

	case string(dao.Disliked):
		switch relationshipUid.State {
		case string(dao.Liked):
			relationshipUid.State = string(dao.Disliked)
			_, err := dao.Db.Model(&relationshipUid).Column("state").Returning("*").Update()
			if err != nil {
				log.Warning.Println("UpdateRelationship UPDATE relationshipUid state error: ", err)
			}
			log.Info.Println("UpdateRelationship UPDATE relationshipUid success, result: ", relationshipUid)
		case string(dao.Disliked):
		case string(dao.Matched):
			relationshipUid.State = string(dao.Disliked)
			relationshipAnotherUid.State = string(dao.Liked)
			_, err := dao.Db.Model(&relationshipUid, &relationshipAnotherUid).Column("state").Update()
			if err != nil {
				log.Warning.Println("UpdateRelationship UPDATE relationshipUid,relationshipAnotherUid state error: ", err)
			}
			log.Info.Println("UpdateRelationship UPDATE relationshipUid, relationshipAnotherUid state success, result: ", relationshipUid, relationshipAnotherUid)
		case string(dao.Default):
			if relationshipUid.Uid == 0 {
				relationshipUid = dao.Relationship{Uid: int64(uid), AnotherUid: int64(anotherUid), State: state, Type: "relationship"}
				relationshipAnotherUid = dao.Relationship{Uid: int64(anotherUid), AnotherUid: int64(uid), State: string(dao.Default), Type: "relationship"}
				err = dao.Db.Insert(&relationshipUid, &relationshipAnotherUid)
				if err != nil {
					log.Warning.Println("UpdateRelationship INSERT relationshipUid, relationshipAnotherUid error: ", err)
				}
				log.Info.Println("UpdateRelationship INSERT relationshipUid, relationshipAnotherUid success, result: ", relationshipUid, relationshipAnotherUid)
			} else {
				relationshipUid.State = string(dao.Disliked)
				_, err := dao.Db.Model(&relationshipUid).Column("state").Returning("*").Update()
				if err != nil {
					log.Warning.Println("UpdateRelationship UPDATE relationshipUid state error: ", err)
				}
				log.Info.Println("UpdateRelationship UPDATE relationshipUid state success, result: ", relationshipUid)
			}
		}
	}

	return
}
