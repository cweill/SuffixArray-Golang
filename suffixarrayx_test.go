package suffixarrayx_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"regexp"
	. "suffixarrayx"
)

var _ = Describe("suffixarrayx", func() {
	/*  ABRACADABRA!
	 *    i ind lcp rnk  select
	 *  ---------------------------
	 *    0  11   -   0  !
	 *    1  10   0   1  A!
	 *    2   7   1   2  ABRA!
	 *    3   0   4   3  ABRACADABRA!
	 *    4   3   1   4  ACADABRA!
	 *    5   5   1   5  ADABRA!
	 *    6   8   0   6  BRA!
	 *    7   1   3   7  BRACADABRA!
	 *    8   4   0   8  CADABRA!
	 *    9   6   0   9  DABRA!
	 *   10   9   0  10  RA!
	 *   11   2   2  11  RACADABRA!
	 */

	Describe("ABRACADABRA!", func() {
		var (
			ind []int
			lcp []int
			rnk []int
			sel []string
		)
		sa := NewSuffixArrayX("ABRACADABRA!")

		BeforeEach(func() {
			ind = []int{11, 10, 7, 0, 3, 5, 8, 1, 4, 6, 9, 2}
			lcp = []int{0, 0, 1, 4, 1, 1, 0, 3, 0, 0, 0, 2}
			rnk = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
			sel = []string{"!", "A!", "ABRA!", "ABRACADABRA!", "ACADABRA!", "ADABRA!",
				"BRA!", "BRACADABRA!", "CADABRA!", "DABRA!", "RA!", "RACADABRA!"}
		})

		Describe("Index", func() {
			It("is correct", func() {
				for i := 0; i < 12; i++ {
					Expect(sa.Index(i)).To(Equal(ind[i]))
				}
			})
		})

		Describe("LongestCommonPrefix", func() {
			It("is correct", func() {
				for i := 1; i < 12; i++ {
					Expect(sa.LongestCommonPrefix(i)).To(Equal(lcp[i]))
				}
			})
		})

		Describe("Rank", func() {
			It("is correct", func() {
				for i := 0; i < 12; i++ {
					Expect(sa.Rank(sel[i])).To(Equal(rnk[i]))
				}
			})
		})

		Describe("Select", func() {
			It("is correct", func() {
				for i := 0; i < 12; i++ {
					Expect(sa.Select(i)).To(Equal(sel[i]))
				}
			})
		})
	})

	Describe("LongestRepeatingSubstring", func() {
		It("is `` for ``", func() {
			sa := NewSuffixArrayX("")
			Expect(sa.LongestRepeatingSubstring()).To(Equal(""))
		})

		It("is `` for `0123456789`", func() {
			sa := NewSuffixArrayX("0123456789")
			Expect(sa.LongestRepeatingSubstring()).To(Equal(""))
		})

		It("is `ana` for `banana`", func() {
			sa := NewSuffixArrayX("banana")
			Expect(sa.LongestRepeatingSubstring()).To(Equal("ana"))
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
			re := regexp.MustCompile("\\s+")
			contentString := re.ReplaceAllString(string(content), " ")
			sa := NewSuffixArrayX(contentString)
			Expect(sa.LongestRepeatingSubstring()).To(Equal("st of times it was the "))
		})

		It("is `,- Such a funny, sporty, gamy, jesty, joky, hoky-poky lad, is the Ocean, oh! Th` for Moby Dick", func() {
			content, err := ioutil.ReadFile("mobydick.txt")
			if err != nil {
				//Should throw exception
			}
			re := regexp.MustCompile("\\s+")
			contentString := re.ReplaceAllString(string(content), " ")
			sa := NewSuffixArrayX(contentString)
			Expect(sa.LongestRepeatingSubstring()).To(Equal(",- Such a funny, sporty, gamy, jesty, joky, hoky-poky lad, is the Ocean, oh! Th"))
		})
	})

	Describe("LongestRepeatingNonOverlappingSubstring", func() {
		It("is `a` for `aaaaaaaaa`", func() {
			sa := NewSuffixArrayX("aaaaaaaaa")
			Expect(sa.LongestRepeatingNonOverlappingSubstring()).To(Equal("a"))
		})

		It("is `ana` for `banana`", func() {
			sa := NewSuffixArrayX("banana")
			Expect(sa.LongestRepeatingNonOverlappingSubstring()).To(Equal("an"))
		})

		It("is `01` for `0101010101010101010101`", func() {
			sa := NewSuffixArrayX("0101010101010101010101")
			Expect(sa.LongestRepeatingNonOverlappingSubstring()).To(Equal("01"))
		})

		It("is `142857` for `0.142857142857142857142857142857142857142`", func() {
			sa := NewSuffixArrayX("0.142857142857142857142857142857142857142")
			Expect(sa.LongestRepeatingNonOverlappingSubstring()).To(Equal("142857"))
		})

		It("is `010989` for `0.0109890109890109890`", func() {
			sa := NewSuffixArrayX("0.0109890109890109890")
			Expect(sa.LongestRepeatingNonOverlappingSubstring()).To(Equal("010989"))
		})

		It("is `` for `0123456789`", func() {
			sa := NewSuffixArrayX("0123456789")
			Expect(sa.LongestRepeatingNonOverlappingSubstring()).To(Equal(""))
		})

		It("is `ATGT` for `ATGTATGT`", func() {
			sa := NewSuffixArrayX("ATGTATGT")
			Expect(sa.LongestRepeatingNonOverlappingSubstring()).To(Equal("ATGT"))
		})

		It("is `ATTCCCGTT` for `ATTGTTCCCATTGTT`", func() {
			sa := NewSuffixArrayX("ATTGTTCCCATTGTT")
			Expect(sa.LongestRepeatingNonOverlappingSubstring()).To(Equal("ATTGTT"))
		})

		It("is `,- Such a funny, sporty, gamy, jesty, joky, hoky-poky lad, is the Ocean, oh! Th` for Moby Dick", func() {
			content, err := ioutil.ReadFile("mobydick.txt")
			if err != nil {
				//Should throw exception
			}
			re := regexp.MustCompile("\\s+")
			contentString := re.ReplaceAllString(string(content), " ")
			sa := NewSuffixArrayX(contentString)
			Expect(sa.LongestRepeatingNonOverlappingSubstring()).To(Equal(",- Such a funny, sporty, gamy, jesty, joky, hoky-poky lad, is the Ocean, oh! Th"))
		})
	})
})
