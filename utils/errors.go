package utils
import (
	"log"
	"fmt"
)
func CheckErr(prefix string, err error){
	if err != nil {
		fmt.Println(prefix)
		log.Fatal(err)
	}
}