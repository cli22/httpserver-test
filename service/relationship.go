package service

import (
	"fmt"
	"httpserver-test/dao"
	"sync"
)

func ListUserRelationship(uid int) (relationships []dao.Relationship, err error) {
	err = dao.Db.Model(&relationships).Where("relationship.uid=?", uid).Select()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(relationships)
	return
}

func UpdateRelationship(uid, anotherUid int, state string) (relationshipUid dao.Relationship, err error) {
	var (
		mu                     sync.Mutex
		relationshipanotherUid dao.Relationship
	)

	mu.Lock()
	defer mu.Unlock()
	err = dao.Db.Model(&relationshipUid).Where("relationship.uid=?", uid).Select()
	if err != nil {
		fmt.Println(err)
	}
	err = dao.Db.Model(&relationshipanotherUid).Where("relationship.uid=?", anotherUid).Select()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(relationshipUid)
	fmt.Println(relationshipanotherUid)
	switch state {
	case string(dao.Liked):
		switch relationshipUid.State {
		case string(dao.Liked):
		case string(dao.Disliked):
			if relationshipanotherUid.State == string(dao.Liked) {
				relationshipUid.State = string(dao.Matched)
				relationshipanotherUid.State = string(dao.Matched)
				res, err := dao.Db.Model(&relationshipUid, &relationshipanotherUid).Column("state").Update()
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(res)
			} else {
				relationshipUid.State = string(dao.Liked)
				fmt.Println(relationshipUid)
				_, err := dao.Db.Model(&relationshipUid).Column("state").Returning("*").Update()
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(relationshipUid)
			}
		case string(dao.Matched):
		case string(dao.Default):
			if relationshipUid.Uid == 0 {
				relationshipUid = dao.Relationship{Uid: int64(uid), AnotherUid: int64(anotherUid), State: state, Type: "relationship"}
				relationshipanotherUid = dao.Relationship{Uid: int64(anotherUid), AnotherUid: int64(uid), State: string(dao.Default), Type: "relationship"}
				err = dao.Db.Insert(&relationshipUid, &relationshipanotherUid)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				relationshipUid.State = string(dao.Liked)
			}
		}

	case string(dao.Disliked):
		switch relationshipUid.State {
		case string(dao.Liked):
			relationshipUid.State = string(dao.Disliked)
			fmt.Println(relationshipUid)
			_, err := dao.Db.Model(&relationshipUid).Column("state").Returning("*").Update()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(relationshipUid)
		case string(dao.Disliked):
		case string(dao.Matched):
			relationshipUid.State = string(dao.Disliked)
			relationshipanotherUid.State = string(dao.Liked)
			res, err := dao.Db.Model(&relationshipUid, &relationshipanotherUid).Column("state").Update()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(res)
		case string(dao.Default):
			if relationshipUid.Uid == 0 {
				relationshipUid = dao.Relationship{Uid: int64(uid), AnotherUid: int64(anotherUid), State: state, Type: "relationship"}
				relationshipanotherUid = dao.Relationship{Uid: int64(anotherUid), AnotherUid: int64(uid), State: string(dao.Default), Type: "relationship"}
				err = dao.Db.Insert(&relationshipUid, &relationshipanotherUid)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				relationshipUid.State = string(dao.Disliked)
			}
		}
	}

	return
}
