SuffixTree-Golang-LCP-LRS
=========================

A replacement for the regular suffix tree packaged with Go. It computes the longest repeated substring (non-overlapping too).

This is a translation from the java suffixtreex.java from Robert Sedgewick and Kevin Wayne at Princeton. Source http://algs4.cs.princeton.edu/63suffix/SuffixArrayX.java.html

The test suite is written in ginkgo. To run the tests you'll first need to run:
```
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega
```
No licence until I get the original author's approval. Until then, this is code only available for academic purposes.
