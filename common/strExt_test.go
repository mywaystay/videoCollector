package common

import (
	"fmt"
	"testing"
)

func TestExtractDesc(t *testing.T) {
	desc := "#2019最值得回味的影😃😌🥰😛😝😢🐍视剧# \n《陈情令》大家一定要看这段视频！！！！\n这可以是教科书级别的忘羡剪辑了吧。\n台词衔接、🔥气氛😃😌🥰😛😝😢🐍🐛🌮🌯🤺🎿💰🔌🎚📻📟💾☎️烘托、空境运用全都恰到好处，全程高能！！！\n#肖战"

	s := ExtractDesc(desc, 30)

	fmt.Println(s)
	fmt.Println(ChineseLen(s))
}