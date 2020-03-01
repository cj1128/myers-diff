# Myers Diff

Demonstrate Myers Diff algorithm in Go. [Git 是怎样生成 diff 的：Myers 算法](http://cjting.me/misc/how-git-generate-diff/?hmsr=toutiao.io&utm_medium=toutiao.io&utm_source=toutiao.io)。

![](http://ww1.sinaimg.cn/large/9b85365dgy1fg4d65ntaij21cu0r00tv)

## Usage

```bash
# char mode: myers-diff -char src_text dst_text
myers-diff -char ABCABBA CBABAC
# file mode: myers-idff src_file dst_file
myers-diff file1.txt file2.txt
```

## Install

```bash
go get -v github.com/cj1128/myers-diff
```

