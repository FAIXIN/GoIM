package main

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

// 删除字符串切片的自定元素
// func popSlice(users *[]string, name string) {
// 	for i := 0; i < len(*users); i++ {
// 		if (*users)[i] == name {
// 			// 如果元素和name相同，删除元素
// 			*users = append((*users)[:i], (*users)[i+1:]...)
// 			i-- // 减少索引以避免跳过下一个元素
// 		}
// 	}
// }

func removeUsersAvatar(users *[]userListAvatar, name string) {
	// 如果用户名非法直接退出
	if name == "" || len(name) < 1 {
		return
	}

	// 检查是否已经存在相同用户名
	for i := 0; i < len(*users); i++ {
		if (*users)[i].Name == name {
			*users = append((*users)[:i], (*users)[i+1:]...)
			break
		}
	}
}

// 向字符串切片中添加非重复元素
// func appendUsers(users *[]string, name string) {
// 	//如果用户名非法直接退出
// 	if name == "" || len(name) < 1 {
// 		return
// 	}
// 	flag := true
// 	for i := 0; i < len(*users); i++ {
// 		if (*users)[i] == name {
// 			flag = false
// 			break
// 		}
// 	}
// 	if flag {
// 		*users = append(*users, name)
// 	}
// }

// 向字符串切片中添加非重复元素
func appendUsersAvatar(users *[]userListAvatar, user userListAvatar) {
	// 如果用户名非法直接退出
	if user.Name == "" || len(user.Name) < 1 {
		return
	}

	// 检查是否已经存在相同用户名
	exists := false
	for i := 0; i < len(*users); i++ {
		if (*users)[i].Name == user.Name {
			exists = true
			break
		}
	}

	// 如果不存在相同用户名，则添加到切片中
	if !exists {
		*users = append(*users, user)
	}
}

// calc img
func hashImg(str string) string {
	h := sha3.New224()
	h.Write([]byte(str))
	hashImgCode := h.Sum(nil) //[]uint8
	hashImgStr := hex.EncodeToString(hashImgCode)
	return hashImgStr
}
