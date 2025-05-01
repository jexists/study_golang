package main

import (
	"fmt"
	s "strings"
	"unicode"
)

func main() {
	var f = fmt.Printf

	// strings.ToUpper() - 대문자로 변경
	f("ToUpper: %s\n", s.ToUpper("Hello world!")) // ToUpper: HELLO WORLD!

	// strings.ToLower() - 소문자로 변경
	f("ToLower: %s\n", s.ToLower("Hello WORLD!")) // ToLower: hello world!

	// strings.Count() - 같은 문자열 몇개 있는지
	f("Count: %v\n", s.Count("Mihalis", "i"))       // Count: 2
	f("Count: %v\n", s.Count("Mihalis", "I"))       // Count: 0
	f("Count: %v\n", s.Count("Mihalis", "li"))      // Count: 1
	f("Count: %v\n", s.Count("Mihalis", "ls"))      // Count: 0
	f("Count: %v\n", s.Count("Mihalis", "Mihalis")) // Count: 1

	// strings.Repeat() - 문자열 반복 (음수X)
	f("Repeat: %s\n", s.Repeat("ab", 5))   // Repeat: ababababab
	f("Repeat: %s\n", s.Repeat("abc", 0))  // Repeat:
	f("Repeat: %s\n", s.Repeat("abcd", 2)) // Repeat: abcdabcd

	// strings.TrimSpace()
	f("TrimSpace: %s\n", s.TrimSpace(" \tThis is a line1. \t")) // TrimSpace: This is a line1.

	// strings.TrimLeft()
	f("TrimLeft: %s\n", s.TrimLeft(" \tThis is a\t line2. \t", ""))      // TrimLeft:       This is a        line2.
	f("TrimLeft: %s\n", s.TrimLeft(" \tThis is a\t line3. \t", "\n"))    // TrimLeft:       This is a        line3.
	f("TrimLeft: %s\n", s.TrimLeft(" \n This is a\t line4. \t", "\n "))  // TrimLeft: This is a      line4.
	f("TrimLeft: %s\n", s.TrimLeft("\n This is a\t line5. \t", "\n"))    // TrimLeft:  This is a     line5.
	f("TrimLeft: %s\n", s.TrimLeft("\nThis is a\t line6. \t", "\n "))    // TrimLeft: This is a      line6.
	f("TrimLeft: %s\n", s.TrimLeft("\n\tThis is a\t line7. \t", "\n"))   // TrimLeft: This is a      line8.
	f("TrimLeft: %s\n", s.TrimLeft(" \tThis is a\t line8. \t", "\n\t ")) // TrimLeft:       This is a        line7.

	// strings.TrimRight()
	f("TrimRight: %s\n", s.TrimRight(" \tThis is a\t line9. \n", "\n"))    // TrimRight:      This is a        line9.
	f("TrimRight: %s\n", s.TrimRight(" \tThis is a\t line10. \n", "\n\t")) // TrimRight:      This is a        line10.

	// strings.Compare() - 문자열 크기 비교
	f("Compare: %v\n", s.Compare("Mihalis", "MIHALIS")) // Compare: 1
	f("Compare: %v\n", s.Compare("Mihalis", "Mihalis")) // Compare: 0
	f("Compare: %v\n", s.Compare("Mihalis", "MIHalis")) // Compare: 1
	f("Compare: %v\n", s.Compare("Mihalis", "MI"))      // Compare: 1
	f("Compare: %v\n", s.Compare("Mihalis", "IH"))      // Compare: 1
	f("Compare: %v\n", s.Compare("a", "a"))             // Compare: 0
	f("Compare: %v\n", s.Compare("a", "A"))             // Compare: 1
	f("Compare: %v\n", s.Compare("A", "a"))             // Compare: -1
	f("Compare: %v\n", s.Compare("a", "b"))             // Compare: -1
	f("Compare: %v\n", s.Compare("z", "x"))             // Compare: 1
	f("Compare: %v\n", s.Compare("apple", "banana"))    // Compare: -1
	f("Compare: %v\n", s.Compare("zoo", "banana"))      // Compare: 1

	// strings.EqualFold() - 대소문자 구별 없이 비교 (true: 같음 / false: 다름)
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAlis")) // EqualFold: true
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAli"))  // EqualFold: false

	// strings.Index() - 두번째 단어 시작의 인덱스 반환 (-1: 같은 것 없는 경우)
	f("Index: %v\n", s.Index("Mihalis", "ha")) // Index: 2
	f("Index: %v\n", s.Index("Mihalis", "Ha")) // Index: -1

	// strings.HasPrefix() - 시작 단어 체크 (true: 시작 / false: 다른단어)
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "Mi")) // Prefix: true
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "mi")) // Prefix: false
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "Ha")) // Prefix: false

	// strings.HasSuffix() - 마지막 단어 체크 (true: 마지막 / false: 다른단어)
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "is")) // Suffix: true
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "Is")) // Suffix: false

	// strings.Fields() - 공백을 기준으로 분리하고 나눠진 문자열 슬라이스 반환
	t := s.Fields("This is a string!")
	f("Fields: %v\n", t)      // Fields: ["This", "is", "a", "string!"]
	f("Fields: %v\n", len(t)) // Fields: 4
	f("Fields: %v\n", t[0])   // Fields: This
	t = s.Fields("ThisIs a\tstring!")
	f("Fields: %v\n", t)      // Fields: ["ThisIs", "a", "string!"]
	f("Fields: %v\n", len(t)) // Fields: 3
	f("Fields: %v\n", t[0])   // Fields: ThisIs
	f("Fields: %v\n", t[1])   // Fields: a
	f("Fields: %v\n", t[2])   // Fields: string!
	t = s.Fields("")
	f("Fields: %v\n", t)      // Fields: []
	f("Fields: %v\n", len(t)) // Fields: 0

	// strings.Split() - 특정 문자열 기준으로 주어진 문자열 분리
	f("%s\n", s.Split("abcd efg", ""))       // ["a", "b", "c", "d", "", "e", "f", "g"]
	f("%v\n", len(s.Split("abcd efg", "")))  // 8
	f("%s\n", s.Split("abcdefg", ""))        // ["a", "b", "c", "d", "e", "f", "g"]
	f("%v\n", len(s.Split("abcdefg", "")))   // 7
	f("%s\n", s.Split("abcd efg", " "))      // ["abcd", "efg"]
	f("%v\n", len(s.Split("abcd efg", " "))) // 2

	// strings.Replace(원본, 기준, 대체, 최대횟수: 음수(-1)로 지정시 무제한 교체)
	f("%s\n", s.Replace("abcd efg", "", "_", -1))  // _a_b_c_d_ _e_f_g_
	f("%s\n", s.Replace("abcd efg", "", "_", 4))   // _a_b_c_d efg
	f("%s\n", s.Replace("abcd efg", "", "_", 2))   // _a_bcd efg
	f("%s\n", s.Replace("abcd efg", "e", "_", -1)) // abcd _fg

	// strings.SplitAfter() -
	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++"))       // SplitAfter: ["123++", "432++", ""]
	f("SplitAfter: %d\n", len(s.SplitAfter("123++432++", "++")))  // SplitAfter: 3
	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "+"))        // SplitAfter: ["123+", "+", "432+", "+", ""]
	f("SplitAfter: %d\n", len(s.SplitAfter("123++432++", "+")))   // SplitAfter: 5
	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "432"))      // SplitAfter: ["123++432", "++"]
	f("SplitAfter: %d\n", len(s.SplitAfter("123++432++", "432"))) // SplitAfter: 2

	// string.Join()
	lines := []string{"Line 1", "line 2", "line 3"}
	f("Join: %s\n", s.Join(lines, "++")) // Join: Line 1++line 2++line 3
	f("Join: %s\n", s.Join(lines, " "))  // Join: Line 1 line 2 line 3
	f("Join: %s\n", s.Join(lines, ""))   // Join: Line 1line 2line 3

	// 알파벳만
	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction)) // TrimFunc: abc ABC
}
