package main

import (
	"fmt"

	"github.com/eure/pairs-fs/pairscore/facebook/client"
)

func main() {
	fb := NewFacebook()
	_ = fb.SetAccessToken("ここに有効なアクセストークンをセット")
	me, _, _ := fb.GetMe()
	fmt.Printf("%v", me)
	authData, _ := fb.GetAuthDataWithoutFriendList()
	fmt.Printf("%v", authData)
}
