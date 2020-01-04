package typing

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
)

// var (
// 	no int = 0
// 	example = []string{"golang", "hands-on", "in", "kagawa"}
// 	result = []string{"ー", "ー", "ー", "ー"}
// 	b strings.Builder
// 	stm time.Time
// 	ptm time.Duration
// 	timeStr string = "ー"
// 	isStart = false
// )

type Typing struct {
	Example []string
	Result []string
	TimeStr string
	b strings.Builder
	no int
	stm time.Time
	ptm time.Duration
	isStart bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func shuffle(src []string) []string {
	idxs := rand.Perm(len(src))

	dst := make([]string, 0, cap(src))
	for _, i := range idxs {
		dst = append(dst, src[i])
	}

	return dst
}

func NewTyping() *Typing {
	t := &Typing {
		Example: []string{"golang", "hands-on", "in", "kagawa"},
		Result: []string{"ー", "ー", "ー", "ー"},
		TimeStr: "ー",
		no: 0,
		isStart: false,
	}
	t.Example = shuffle(t.Example)

	return t
}

func (t *Typing) IsContinue() bool {
	return t.no < len(t.Example)
}

func (t *Typing) Input(inputs []rune) {
	if t.isStart == false {
		t.stm = time.Now()
		t.isStart = true
	}
	for _, r := range inputs {
		t.b.WriteRune(r)
	}
}

func (t *Typing) Check() {
	// Enterが離れた
	t.ptm = time.Since(t.stm)
	t.TimeStr = fmt.Sprintf("CurrentTime: %.3f(sec)", t.ptm.Seconds())
	t.isStart = false

	if t.Example[t.no] == t.b.String() {
		t.Result[t.no] = "○"
	} else {
		t.Result[t.no] = "×"
	}
	// Enterで次の入力へ
	t.no++
	if t.no >= len(t.Example) {
		t.no = len(t.Example) - 1
	}
	t.b.Reset()	
}

func (t *Typing) CurrentExample() string {
	return t.Example[t.no]
}

func (t *Typing)  CurrentTime() string {
	return t.TimeStr
}

func (t *Typing)  Examples() []string {
	return t.Example
}

func (t *Typing)  Results() []string {
	return t.Result
}

func (t *Typing)  UserString() string {
	return t.b.String()
}