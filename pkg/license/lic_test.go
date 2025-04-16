package license

import (
	"fmt"
	"testing"
)

func TestLic(t *testing.T) {
	lic, err := LicGen("xxx公司", 3600, []string{"13214", "dfadfas"}, []string{"own", "cmdb"})
	fmt.Println("lic:", lic, err)

	//b, err := LicParse("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXN0IjoiZTFjZWRhZGQzOTVmMjdjZmQyZGRkN2Q3MGMyZGQ1ZTE4MzBmYzcyYzdhMmZlMjQ0ODZjMmIwNTgwZWRkZWRhZCIsImV4cCI6MTY3ODM0NjAyMiwiZXhwX2RhdGUiOiIyMDIzLTAzLTA5IDE1OjEzOjQyIiwiaWF0IjoxNjc4MzQ1OTYyLCJpc3MiOiJhaXIgand0LiJ9.qKfWkxoxnkwoeJkg7szCeHztXr1iQXC8O-5oWG6vQ7M")
	//fmt.Println("err", b, err)
}
