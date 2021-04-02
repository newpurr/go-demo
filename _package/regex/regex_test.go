package regex

import (
	"fmt"
	"regexp"
	"testing"
)

// Go正则表达式示例: https://colobu.com/2020/11/11/golang-regex-replace-example/

// https://gowalker.org/regexp#MatchString
func TestRegexMatchString(t *testing.T) {
	str := "Golang regular expressions example"
	match, err := regexp.MatchString(`^Golang`, str)
	fmt.Println("Match: ", match, " Error: ", err)
}

func TestCompile(t *testing.T) {
	str := "Golang expressions example"

	// Compile() 或者 MustCompile()创建一个编译好的正则表达式对象。假如正则表达式非法，那么Compile()方法回返回error,而MustCompile()编译非法正则表达式时不会返回error，而是回panic。
	// 如果你想要很好的性能，不要在使用的时候才调用Compile()临时进行编译，而是预先调用Compile()编译好正则表达式对象
	compile, err := regexp.Compile("Gola([a-z]+)g")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(compile.FindString(str))
}

// FindAllString
// FindString方法的All版本，它返回所有匹配的字符串的slice。如果返回nil值代表没有匹配的字符串
func TestFindAllString(t *testing.T) {
	str := "Golang regular expressions example"

	compile, err := regexp.Compile(`p([a-z]+)e`)
	if err != nil {
		fmt.Println(err)
		return
	}
	match := compile.FindAllString(str, 3)

	fmt.Println("Match: ", match, " Error: ", err)
}
