package main

import (
	"fmt"
	"github.com/gogotchuri/go-rs-ge/rs"
)

func main() {
	rs := &rs.Client{
		ServiceUser:     "beka:405217420",
		ServicePassword: "RT1234",
		UserID:          0,
	}
	resp, err := rs.GetUniqueID()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp)
	fmt.Println(rs)
}
