// 08_random_avatar_generator.go

package main

import (
	"github.com/landaiqing/go-pixelnebula"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	savePath = "random_generated_avatars.svg"
)

// 随机生成头像
// Note: 随机传入id即可随机生成头像，固定的id会生成相同的头像
func main() {

	// 创建随机数
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(100)

	pixelNebula := pixelnebula.NewPixelNebula().WithDefaultCache()
	svg, err := pixelNebula.Generate(strconv.Itoa(randomInt), false).ToSVG()
	if err != nil {
		panic(err)
	}
	// 保存图片
	os.WriteFile(savePath, []byte(svg), 0644)
	defer os.Remove(savePath)
}
