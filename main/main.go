package main

import (
	"fmt"

	gorsge "github.com/gogotchuri/go-rs-ge"
)

func main() {
	rs := &gorsge.RsGeClient{
		ServiceUser:     "giorgi kernel:405368212",
		ServicePassword: "GIORGIKERNEL",
		UserID:          1464390,
	}
	resp, err := rs.CheckCredentials()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp)
}
