# Myers Diff

Demonstrate Myers Diff algorithm in Go. [Git 是怎样生成 diff 的：Myers 算法](https://cjting.me/misc/how-git-generate-diff/)。

![](http://asset.cjting.cn/FvoGcyZRGTsakveEtUr3XxezxmoU.png)

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

