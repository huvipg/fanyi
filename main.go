package main

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"

	. "../../src/dir"
)

var p = fmt.Println

func fy() {
	q := R_file("English.txt")
	qq := url.QueryEscape(q)
	get := Get(`https://translate.google.cn/translate_a/single?client=gtx&sl=en&tl=zh&dt=t&q=` + qq)
	W_file("del.txt", get)
	cn, json, encn := rl_file("del.txt")
	W_file("翻译中文.txt", cn)
	W_file("json.json", json)
	W_file("英文加中文.txt", encn)
	//rl_file("color.txt")
	p(cn)

}
func main() {
	q := R_file("English.txt")
	
}
func fg(str string) {
}
func rl_file(path string) (string, string, string) {
	json := ""
	astr := ""
	encn := ""
	fi, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return "N", "N", "N"
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		stra := string(a)
		if c == io.EOF {
			break
		}
		if Rec(stra) {
			straa := strings.Replace(stra, `[[[`, `,[`, 1)
			saa := strings.Replace(straa, `,null,"en"]`, ``, 1)
			sa := strings.Replace(saa, `\n`, "", -1)

			s := strings.Split(sa, `,null,`)
			p(s[0])
			if len(s[0]) > 0 {
				ress := strings.Replace(s[0], `","`, `"&_&"`, -1)
				chen := strings.Split(ress, `&_&`)
				yen := chen[1]
				qe := strings.Replace(yen, `"`, ``, -1)
				qen := strings.Replace(qe, `\`, ``, -1)
				p(qen)
				//p(fress[0])
				zh := strings.Split(chen[0], `[`)
				yzh := zh[1]
				qzh := strings.Replace(yzh, `"`, ``, -1)

				p(qzh)

				yesstr := `{"zh":` + yzh + `, "en":` + yen + `},`
				json += yesstr
				//		astr += fress[0] + "\n"
				p(yesstr)
				astr += qzh + "\n"
				encn += qzh + "\n" + qen + "\n"
			}

		}

	}
	return astr, json, encn
}

func Restr(s string) bool {
	//	fmt.Println(strings.Contains(s, `//`)) //true
	return strings.Contains(s, `//`)

}

//]如果是1返回false
func Rec(s string) bool {
	//	fmt.Println(strings.Contains(s, `//`)) //true
	nlen := len(s)
	//p(nlen)
	if nlen == 1 || nlen == 0 {
		return false
	} else {

		return true
	}

}
