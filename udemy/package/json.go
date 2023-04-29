package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// json
// 構造体からJSONテキストへ変換
type A struct{}

type User struct {
	ID      int       `json:"id"` // `json:"id,string"`JSONに変換した際に文字列型になる。
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       *A        `json:"A,omitempty"` // {}の場合は、ポインタ型にすると消える。
}

func (u *User) UnmarshalJSON(b []byte) error {
	type User2 struct {
		Name string
	}
	var u2 User2
	err := json.Unmarshal(b, &u2)
	if err != nil {
		fmt.Println(err)
	}
	u.Name = u2.Name + "!"
	return err
}

func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr " + u.Name,
	})
	return v, err
}

func main() {
	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "example@example.com"
	u.Created = time.Now()

	// Marshal JSONに変換
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bs)
	fmt.Println(string(bs))

	fmt.Printf("%T\n", bs)

	// JSONから構造体に
	u2 := new(User)

	// UnMarshal JSONをデータに変換
	if err := json.Unmarshal(bs, u2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(u2)

}
