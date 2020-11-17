package main

import (
	"fmt"
	"strings"

	serialize "github.com/smeyerhot/serialize/treeserial"
)

func main() {
	obj := serialize.Constructor()
	input := []string{"1", "2", "3", "4", "5", "6"}
	data := strings.Join(input, ",")
	ans := obj.Deserialize(data)
	output := obj.Serialize(ans)
	fmt.Print(output)

}
