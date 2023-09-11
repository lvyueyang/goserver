package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenCaptcha 生成6位随机验证码
func GenCaptcha() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}
