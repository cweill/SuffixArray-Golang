SuffixArray-Golang
=========================

A replacement for the regular suffix array packaged with Go. It computes the longest repeated substring (non-overlapping too).

This is a translation of the java implementation by Robert Sedgewick and Kevin Wayne at Princeton. Source http://algs4.cs.princeton.edu/63suffix/SuffixArrayX.java.html

The test suite is written in ginkgo. To run the tests you'll first need to run:
```
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega
```
## License

SuffixArray-Golang is released under the [MIT License](http://www.opensource.org/licenses/MIT).
