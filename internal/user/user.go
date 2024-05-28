package user

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	crypto "github.com/Altermilk/cryptoMath"
)

type User struct {
	Name string
	p, q, N int
	c, D, f int
	msg int
}


func NewUser(name string, rnd *rand.Rand) *User{
	user := User{Name : name}
	user.SetKeys(rnd)
	return &user
}

func (u*User) SetKeys(rnd *rand.Rand){
	u.p = crypto.GetRandomSimpleNum(*rnd)
	u.q = crypto.GetRandomSimpleNum(*rnd)
	u.N = u.p*u.q
	u.f = (u.p - 1)*(u.q - 1)
	for {
		d := crypto.GetRandomSimpleNum(*rnd)
		if d < u.f{
			u.D = d
			break
		}
	}
	u.c = crypto.ModInv(int(u.D), int(u.f))

}

func (u*User) ShareD() (d int){
	return u.D
}

func (u*User) ShareN() (N int){
	return u.N
}

func (u*User) CheckN(Nb int) bool{
	return u.msg < Nb
}

func (u *User) SendMsg(msg, Nb, db int) (e []int, ok bool){
	fmt.Println(u.Name, " Sent msg: \"", msg, "\"")
	fmt.Println(" key: \"", Nb, "\"")
	if msg < Nb{
		return []int{encrypt(msg, db, Nb)}, true
	}
	fmt.Println("Invalid key")
	partsInt := cutMsg(msg, Nb, encrypt, [2]int{db, Nb})
	
	return partsInt, false
}

func (u*User) GetMsg(e []int){
	fmt.Println(u.Name, " Recieved encrypted msg: \"", e, "\"")

	u.msg = decrypt(e[0], u.c, u.N)
	
	fmt.Println(u.Name, " Decrypted msg: \"", u.msg, "\"")
}

func cutMsg(msg, Nb int, f func(int, int, int) int, args [2]int) []int {
	m := strconv.FormatInt(int64(msg), 2)
	n := strconv.FormatInt(int64(Nb), 2)

	size := len(m) / len(n)
	remainder := len(m) % len(n)

	parts := make([]string, size)
	for i := 0; i < size; i++ {
		start := i * len(n)
		end := (i + 1) * len(n)
		parts[i] = m[start:end]
	}

	// Обработка остатка
	// Обработка остатка
// Обработка остатка
if remainder > 0 {
	sb := strings.Builder{}
	for i := 0; i < len(n); i++ {
		sb.WriteString("0")
	}
	str := strings.ReplaceAll(sb.String(), sb.String()[:remainder], m[size*len(n):])
	parts = append(parts, str)
}

partsInt := make([]int, len(parts))
for i := len(parts) - 1; i >= 0; i-- {
	M, _ := strconv.ParseInt(parts[i], 2, 64)
	partsInt[i] = f(int(M), args[0], args[1])
}



	return partsInt
}



func buildInt(e []int) int {
	sb := strings.Builder{}
	for i := range e {
		sb.WriteString(strconv.FormatInt(int64(e[i]), 2))
	}
	E, _ := strconv.ParseInt(sb.String(), 10, 64)
	return int(E)
}

func encrypt(M, db, Nb int) int{
	return crypto.Modularizate(M, db, Nb)
}

func decrypt(e, cb, nb int) int{
	return crypto.Modularizate(e, cb, nb)
}

// func buildInt(e []int) int{
// 	sb := strings.Builder{}
// 		for i := range e{
// 			sb.WriteString(strconv.Itoa(e[i]))
// 		}
// 		E, _ := strconv.Atoi(sb.String())
// 	return E
// }