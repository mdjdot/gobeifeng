package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	{
		s := "aba张天"
		h := md5.New()
		fmt.Printf("%+v\n", md5.Sum([]byte(s)))
		fmt.Printf("%x\n", md5.Sum([]byte(s)))
		fmt.Printf("%s\n", md5.Sum([]byte(s)))
		fmt.Printf("%+v\n", h.Sum([]byte(s)))
		fmt.Printf("%x\n", h.Sum([]byte(s)))
		fmt.Printf("%s\n", h.Sum([]byte(s)))
		/*
			[129 55 58 125 106 139 24 97 131 154 82 4 225 165 5 225]
			81373a7d6a8b1861839a5204e1a505e1
			�7:}j�a��R��
			[97 98 97 229 188 160 229 164 169 212 29 140 217 143 0 178 4 233 128 9 152 236 248 66 126]
			616261e5bca0e5a4a9d41d8cd98f00b204e9800998ecf8427e
			aba张天��ُ��	���B~
		*/
	}

	{
		s1 := "aba张天"
		s2 := "salt"
		h1 := md5.New()
		h1.Write([]byte(s1))
		fmt.Printf("%+v\n", md5.Sum([]byte(s1)))
		fmt.Printf("%x\n", md5.Sum([]byte(s1)))
		fmt.Printf("%s\n", md5.Sum([]byte(s1)))
		fmt.Printf("%+v\n", h1.Sum([]byte(s2)))
		fmt.Printf("%x\n", h1.Sum([]byte(s2)))
		fmt.Printf("%s\n", h1.Sum([]byte(s2)))
		/*
			[129 55 58 125 106 139 24 97 131 154 82 4 225 165 5 225]
			81373a7d6a8b1861839a5204e1a505e1
			�7:}j�a��R��
			[115 97 108 116 129 55 58 125 106 139 24 97 131 154 82 4 225 165 5 225]
			73616c7481373a7d6a8b1861839a5204e1a505e1
			salt�7:}j�a��R��
		*/
	}
}
