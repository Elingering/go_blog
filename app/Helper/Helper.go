package Helper

import (
	"fmt"
	"math/rand"
	"time"
)

//生成6位随机数
func Random6() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	str := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return str
}
