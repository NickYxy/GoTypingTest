package service

import (
	"Tiktok/dao"
	"fmt"
	"testing"
)

func TestCountComment(t *testing.T) {
	dao.Init()
	impl := CommentServiceImpl{}
	count, err := impl.CountFromVideoId(1)
	if err != nil {
		fmt.Println("count err:", err)
	}
	fmt.Println(count)
}
