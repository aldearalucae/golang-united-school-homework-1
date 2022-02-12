package solution

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func solution() {
	helloWord := emoji.Sprint("Hello :world_map:")
	fmt.Println(helloWord)
}
