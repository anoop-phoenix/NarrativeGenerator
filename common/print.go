package common

import (
	"encoding/json"
	"fmt"
)

// PrettyPrint prints the contents of an obj with identation
func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}
