/**
 *  The <tt>suffixarrayx</tt> package represents a suffix array of a string of
 *  length <em>N</em>.
 *  It supports the <em>selecting</em> the <em>i</em>th smallest suffix,
 *  getting the <em>index</em> of the <em>i</em>th smallest suffix,
 *  computing the length of the <em>longest common prefix</em> between the
 *  <em>i</em>th smallest suffix and the <em>i</em>-1st smallest suffix,
 *  and determining the <em>rank</em> of a query string (which is the number
 *  of suffixes strictly less than the query string).
 *  <p>
 *  This implementation uses 3-way radix quicksort to sort the array of suffixes.
 *  For a simpler (but less efficient) implementations of the same API, see
 *  {@link SuffixArray}.
 *  The <em>index</em> and <em>length</em> operations takes constant time
 *  in the worst case. The <em>lcp</em> operation takes time proportional to the
 *  length of the longest common prefix.
 *  The <em>select</em> operation takes time proportional
 *  to the length of the suffix and should be used primarily for debugging.
 *  <p>
 *  For additional documentation, see <a href="http://algs4.cs.princeton.edu/63suffix">Section 6.3</a> of
 *  <i>Algorithms, 4th Edition</i> by Robert Sedgewick and Kevin Wayne.
 *
 *  Source: http://algs4.cs.princeton.edu/63suffix/SuffixArrayX.java.html
 *  Translated to Golang by Charles Weill
 */

package suffixarrayx

import (
	"strings"
)

type suffixarrayx struct {
	CUTOFF int
	text   []rune
	index  []int
	n      int
}

// Constructor
func NewSuffixArrayX(str string) *suffixarrayx {
	str = str + "\n"
	sa := new(suffixarrayx)
	sa.CUTOFF = 5
	sa.text = []rune(str)
	sa.n = len(str)
	sa.index = make([]int, len(str))

	for i := 0; i < sa.n; i++ {
		sa.index[i] = i
	}
	// shuffle
	sa.sort(0, sa.n-1, 0)
	return sa
}

// 3-way string quicksort lo..hi starting at dth character
func (sa *suffixarrayx) sort(lo, hi, d int) {
	// cutoff to insertion sort for small subarrays
	if hi <= lo+sa.CUTOFF {
		sa.insertion(lo, hi, d)
		return
	}
	lt, gt := lo, hi
	v := sa.text[sa.index[lo]+d]
	i := lo + 1
	for i <= gt {
		t := sa.text[sa.index[i]+d]
		if t < v {
			sa.exch(lt, i)
			lt++
			i++
		} else if t > v {
			sa.exch(i, gt)
			gt--
		} else {
			i++
		}
	}

	// a[lo..lt-1] < v = a[lt..gt] < a[gt+1..hi].
	sa.sort(lo, lt-1, d)
	if v > 0 {
		sa.sort(lt, gt, d+1)
	}
	sa.sort(gt+1, hi, d)
}

// sort from a[lo] to a[hi], starting at the dth character
func (sa *suffixarrayx) insertion(lo, hi, d int) {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo && sa.less(sa.index[j], sa.index[j-1], d); j-- {
			sa.exch(j, j-1)
		}
	}
}

// is text[i+d..N) < text[j+d..N) ?
func (sa *suffixarrayx) less(i, j, d int) bool {
	if i == j {
		return false
	}
	i = i + d
	j = j + d
	for i < sa.n && j < sa.n {
		if sa.text[i] < sa.text[j] {
			return true
		}
		if sa.text[i] > sa.text[j] {
			return false
		}
		i++
		j++
	}
	return i > j
}

// exchange index[i] and index[j]
func (sa *suffixarrayx) exch(i, j int) {
	swap := sa.index[i]
	sa.index[i] = sa.index[j]
	sa.index[j] = swap
}

/**
 * Returns the index into the original string of the <em>i</em>th smallest suffix.
 * That is, <tt>text.substring(sa.index(i))</tt> is the <em>i</em> smallest suffix.
 * @param i an integer between 0 and <em>N</em>-1
 * @return the index into the original string of the <em>i</em>th smallest suffix
 */

func (sa *suffixarrayx) Index(i int) int {
	return sa.index[i]
}

/**
 * Returns the length of the longest common prefix of the <em>i</em>th
 * smallest suffix and the <em>i</em>-1st smallest suffix.
 * @param i an integer between 1 and <em>N</em>-1
 * @return the length of the longest common prefix of the <em>i</em>th
 * smallest suffix and the <em>i</em>-1st smallest suffix.
 */
func (sa *suffixarrayx) LongestCommonPrefix(i int) int {
	return sa.lcp(sa.index[i], sa.index[i-1])
}

func (sa *suffixarrayx) lcp(i, j int) int {
	length := 0
	for i < sa.n && j < sa.n {
		if sa.text[i] != sa.text[j] {
			return length
		}
		i++
		j++
		length++
	}
	return length
}

/**
 * Returns the <em>i</em>th smallest suffix as a string.
 * @param i the index
 * @return the <em>i</em> smallest suffix as a string
 */
func (sa *suffixarrayx) Select(i int) string {
	return string(sa.text[sa.index[i] : sa.n-sa.index[i]])
}

/**
 * Returns the number of suffixes strictly less than the <tt>query</tt> string.
 * We note that <tt>rank(select(i))</tt> equals <tt>i</tt> for each <tt>i</tt>
 * between 0 and <em>N</em>-1.
 * @param query the query string
 * @return the number of suffixes strictly less than <tt>query</tt>
 */
func (sa *suffixarrayx) Rank(query string) int {
	lo, hi := 0, sa.n-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		cmp := sa.compare(query, sa.index[mid])
		if cmp < 0 {
			hi = mid - 1
		} else if cmp > 0 {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return lo
}

// is query < text[i..N) ?
func (sa *suffixarrayx) compare(query string, i int) int {
	queryRunes := []rune(query)
	m := len(query)
	j := 0
	for i < sa.n && j < m {
		return int(queryRunes[j]) - int(sa.text[i])
		i++
		j++
	}
	if j < sa.n {
		return -1
	}
	if j < m {
		return 1
	}
	return 0
}

func (sa *suffixarrayx) LongestRepeatingSubstring() string {
	lrs := ""
	for i := 1; i < sa.n; i++ {
		length := sa.LongestCommonPrefix(i)
		if length > len(lrs) {
			lrs = string(sa.text[sa.index[i] : sa.index[i]+length])
		}
	}
	return lrs
}

func (sa *suffixarrayx) LongestRepeatingNonOverlappingSubstring() string {
	text := string(sa.text)
	lrs := sa.LongestRepeatingSubstring()
	lrnos := lrs
	for len(lrs) > 0 {
		sa2 := NewSuffixArrayX(string(lrnos))
		lrs = sa2.LongestRepeatingSubstring()
		newSubstr := lrnos[0 : len(lrnos)-len(lrs)]
		if lrs != lrnos[len(lrnos)-len(lrs):] || !strings.Contains(text, newSubstr+newSubstr) {
			break
		}
		lrnos = newSubstr
	}
	return lrnos
}
