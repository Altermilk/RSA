package main

import (
	"fmt"
	"math/rand"
	user "rsa/internal/user"
	"time"
)

func main(){
	var rnd = rand.New(rand.NewSource(time.Now().Unix()))
	u1 := user.NewUser("Alice", rnd)
	u2 := user.NewUser("Bob", rnd)
	e, ok := u1.SendMsg(1234567, u2.ShareN(), u2.ShareD())
	if ok{
		u2.GetMsg(e)
	}else{
		for i:=0;i<10;i++{
			u2.SetKeys(rnd)
			if u1.CheckN(u2.N){
				e, ok := u1.SendMsg(1234567, u2.ShareN(), u2.ShareD())
				if ok{
					u2.GetMsg(e)
					break
				}else{
					fmt.Println("attempt ", i + 1, "/10 : Invalid key (", u2.N, ")")
				}
				
			
		}
	}

}
}