package service

import (
	"fmt"
	"httpserver-test/dao"
)

func ListUserRelationship(uid int) (relationships []dao.Relationship, err error) {
	err = dao.Db.Model(&relationships).Column("relationship.uid", "relationship.state", "relationship.type").Where("uid=%d", uid).Select()
	if err != nil {
		fmt.Print(err)
	}
	return
}

func UpdateRelationship(uid, anotherUid int, state dao.RelationshipState) (relationship dao.Relationship, err error) {
	var (
		relationshipUid        dao.Relationship
		relationshipanotherUid dao.Relationship
	)
	// 查询uid，anthorUid 状态r1
	// 查询anthorUid，uid 状态r2
	// 如果存在，
	// r1.state = state, 不更新，并返回
	// r1.state != state, if state=liked and r1.state=matched, 不更新
	// r1.state != state, if state=disliked and r1.state=matched, 更新为disliked， 并且判断r2.state==matched, 更新为liked
	// 如果不存在，插入2条新数据
	err = dao.Db.Model(&relationshipUid).Column("relationship.state").Where("uid=%d", uid).Select()
	if err != nil {
		fmt.Print(err)
	}
	err = dao.Db.Model(&relationshipanotherUid).Column("relationship.state").Where("uid=%d", anotherUid).Select()
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(relationshipUid)
	fmt.Print(relationshipanotherUid)

	return
}
