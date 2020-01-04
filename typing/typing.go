package typing

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
)

var (
	no int = 0
	example = []string{"golang", "hands-on", "in", "kagawa"}
	result = []string{"ー", "ー", "ー", "ー"}
	b strings.Builder
	stm time.Time
	ptm time.Duration
	timeStr string = "ー"
	isStart = false
)

func init() {
	rand.Seed(time.Now().UnixNano())
	example = shuffle(example)
}

func shuffle(src []string) []string {
	idxs := rand.Perm(len(src))

	dst := make([]string, 0, cap(src))
	for _, i := range idxs {
		dst = append(dst, src[i])
	}

	return dst
}

func IsContinue() bool {
	return no < len(example)
}

func Input(inputs []rune) {
	if isStart == false {
		stm = time.Now()
		isStart = true
	}
	for _, r := range inputs {
		b.WriteRune(r)
	}
}

func Check() {
	// Enterが離れた
	ptm = time.Since(stm)
	timeStr = fmt.Sprintf("CurrentTime: %.3f(sec)", ptm.Seconds())
	isStart = false

	if example[no] == b.String() {
		result[no] = "○"
	} else {
		result[no] = "×"
	}
	// Enterで次の入力へ
	no++
	if no >= len(example) {
		no = len(example) - 1
	}
	b.Reset()	
}

func CurrentExample() string {
	return example[no]
}

func CurrentTime() string {
	return timeStr
}

func Examples() []string {
	return example
}

func Results() []string {
	return result
}

func UserString() string {
	return b.String()
}