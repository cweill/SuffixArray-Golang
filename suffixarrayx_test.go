package suffixarrayx_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"strings"
	. "suffixarrayx"
)

var _ = Describe("suffixarrayx", func() {
	Describe("LongestRepeatingSubstring", func() {
		It("is `` for ``", func() {
			sa := NewSuffixArrayX("")
			Expect(sa.LongestRepeatingSubstring()).To(Equal(""))
		})

		It("is `` for `0123456789`", func() {
			sa := NewSuffixArrayX("0123456789")
			Expect(sa.LongestRepeatingSubstring()).To(Equal(""))
		})

		It("is `a` for `aa`", func() {
			sa := NewSuffixArrayX("aa")
			Expect(sa.LongestRepeatingSubstring()).To(Equal("a"))
		})

		It("is `0101010101` for `0101010101010101010101`", func() {
			sa := NewSuffixArrayX("0101010101010101010101")
			Expect(sa.LongestRepeatingSubstring()).To(Equal("01010101010101010101"))
		})

		It("is `aaaaaaaa` for `aaaaaaaaa`", func() {
			sa := NewSuffixArrayX("aaaaaaaaa")
			Expect(sa.LongestRepeatingSubstring()).To(Equal("aaaaaaaa"))
		})

		It("is `14285714285` for `0.14285714285714285`", func() {
			sa := NewSuffixArrayX("0.14285714285714285")
			Expect(sa.LongestRepeatingSubstring()).To(Equal("14285714285"))
		})

		It("is `st of times it was the ` for Tiny Tale", func() {
			content, err := ioutil.ReadFile("tinyTale.txt")
			if err != nil {
				//Should throw exception
			}
			contentString := strings.Replace(string(content), "\n", " ", -1)
			sa := NewSuffixArrayX(contentString)
			Expect(sa.LongestRepeatingSubstring()).To(Equal("st of times it was the "))
		})

		// This is a really long test. Should take about 30 minutes to finish
		//
		// It("is `,- Such a funny, sporty, gamy, jesty, joky, hoky-poky lad, is the Ocean, oh! Th` for Moby Dick", func() {
		// 	content, err := ioutil.ReadFile("mobydick.txt")
		// 	if err != nil {
		// 		//Should throw exception
		// 	}
		// 	sa := NewSuffixArrayX(strings.Replace(string(content), "\n", " ", -1))
		// 	Expect(sa.LongestRepeatingSubstring()).To(Equal(",- Such a funny, sporty, gamy, jesty, joky, hoky-poky lad, is the Ocean, oh! Th"))
		// })
	})

	Describe("LongestRepeatingNonOverlappingSubstring", func() {
		It("is `a` for `aaaaaaaaa`", func() {
			sa := NewSuffixArrayX("aaaaaaaaa")
			Expect(sa.LongestRepeatingNonOverlappingSubstring()).To(Equal("a"))
		})

		It("is `01` for `0101010101010101010101`", func() {
			sa := NewSuffixArrayX("0101010101010101010101")
			Expect(sa.LongestRepeatingNonOverlappingSubstring()).To(Equal("01"))
		})

		It("is `142857` for `0.142857142857142857142857142857142857142`", func() {
			sa := NewSuffixArrayX("0.142857142857142857142857142857142857142")
			Expect(sa.LongestRepeatingNonOverlappingSubstring()).To(Equal("142857"))
		})

		It("is `` for `0123456789`", func() {
			sa := NewSuffixArrayX("0123456789")
			Expect(sa.LongestRepeatingNonOverlappingSubstring()).To(Equal(""))
		})
	})
})
