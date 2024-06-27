package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	A := "hello20"
	B := "hello20"
	fmt.Println(strings.Compare(A, B))
	fmt.Println(strings.Contains(A, "2"))
	fmt.Println(strings.ContainsAny(A, "1200"))
	fmt.Println(strings.Count(A, "l"))
	fmt.Println(strings.Cut(A, "lo"))

	bf, af, fo := strings.Cut(A, "e")
	fmt.Println(bf)
	fmt.Println(af)
	fmt.Println(fo)

	// Compare only for strings not numbers
	fmt.Println(strings.EqualFold(A, "HELLO20"))

	fmt.Printf("Fields are: %q\n", strings.Fields("he ll o20"))
	C := strings.Fields("he ll o")
	fmt.Println(C[0])
	fmt.Println(C[1])
	fmt.Println(C[2])

	//to remove any special chars between words
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q\n", strings.FieldsFunc("  foo1;bar2,baz3...", f))

	//only for first letters
	fmt.Println(strings.HasPrefix("Bola", "B"))
	fmt.Println(strings.HasPrefix("Gopher", "C"))
	fmt.Println(strings.HasPrefix("Gopher", ""))
	fmt.Println("---------")

	//only for last letters
	fmt.Println(strings.HasSuffix("Amigo", "go"))
	fmt.Println(strings.HasSuffix("Amigo", "O"))
	fmt.Println(strings.HasSuffix("Amigo", "Ami"))
	fmt.Println(strings.HasSuffix("Amigo", ""))
	fmt.Println("---------")

	//waiking on every element of the string
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))

	s := []string{"foo", "bar", "baz"} //slice are string into some strings "ah" "me" "d"
	fmt.Println(strings.Join(s, " - "))

	//repeat string
	fmt.Println("ba" + strings.Repeat("na", 2))

	//index of the word or letters
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
	fmt.Println("---------")

	fmt.Println(strings.IndexAny("chicken", "aeiouy"))
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))
	fmt.Println("---------")

	fmt.Println(strings.LastIndex("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "rodent"))
	fmt.Println("---------")

	fmt.Println(strings.LastIndexAny("go gopher", "go"))
	fmt.Println(strings.LastIndexAny("go gopher", "rodent"))
	fmt.Println(strings.LastIndexAny("go gopher", "fail"))
	fmt.Println("---------")

	fmt.Println(strings.IndexByte("golang", 'g'))
	fmt.Println(strings.IndexByte("golang", 'x'))
	fmt.Println("---------")

	d := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", d))
	fmt.Println(strings.IndexFunc("Hello, world", d))
	fmt.Println("---------")

	fmt.Println(strings.LastIndexFunc("go 123", unicode.IsNumber))
	fmt.Println(strings.LastIndexFunc("123 go", unicode.IsNumber))
	fmt.Println(strings.LastIndexFunc("go", unicode.IsNumber))
	fmt.Println("---------")

	fmt.Println(strings.IndexRune("chicken", 'k'))
	fmt.Println(strings.IndexRune("chicken", 'd'))
	fmt.Println("---------")

	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	fmt.Println("---------")

	fmt.Println(strings.Split("a,b,c", ","))
	fmt.Println(strings.Split("a man a plan a canal panama", "a "))
	fmt.Println(strings.Split(" xyz ", ""))
	fmt.Println(strings.Split("", "Bernardo"))
	fmt.Println("---------")

	fmt.Println(strings.SplitAfter("a,b,c", ","))
	fmt.Println("---------")

	// Capitalize first letter of every word
	fmt.Println(strings.Title("her royal highness"))
	fmt.Println("---------")

	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın Aizonai"))
	fmt.Println("---------")

	// Returns all letters to lower case.
	fmt.Println(strings.ToLower("GoPher"))
	fmt.Println("---------")

	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "ÖnNek İş"))
	fmt.Println("---------")

	// Returns all letters to upper case.
	fmt.Println(strings.ToUpper("Gopher"))
	fmt.Println("---------")

	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))
	fmt.Println("---------")

	//Remove any unwanted chars
	fmt.Println(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))
	fmt.Println("---------")

	fmt.Println(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
	fmt.Println("---------")

	fmt.Println(strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡"))
	fmt.Println("---------")

	fmt.Println(strings.TrimLeftFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
	fmt.Println("---------")

	fmt.Println(strings.TrimRight("¡¡¡Hello, Gophers!!!", "!¡"))
	fmt.Println("---------")

	fmt.Println(strings.TrimRightFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
	fmt.Println("---------")

	var v = "¡¡¡Hello, Gophers!!!"
	v = strings.TrimSuffix(v, ", Gophers!!!")
	v = strings.TrimSuffix(v, ", Marmots!!!")
	fmt.Println(v)
	fmt.Println("---------")

	//Remove any white spaces
	fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))
	fmt.Println("---------")

}
