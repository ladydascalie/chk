# chk
chk your file hashes

## Install

`go get github.com/ladydascalie/chk`

Or just grab one of the [releases](https://github.com/ladydascalie/chk/releases)

## Examples

```
Usage of chk:
  -c string
    	chk -c {hash} {file or files}
```

### Display a file's hashes

```
chk file.txt

# output:
file:    file.txt
md5:     01234516b98bac202ecaee7ac4001234
sha1:    012343ee5e6b4b0d3255bfef95601890afd01234
sha256:  0123444298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b78501234
```

### Compare to a known hash

```
chk -c 01234516b98bac202ecaee7ac4001234 file.txt

# output will be the same, but any matching hash will be printed in green
```


### Checking multiple files

Of course this works with the `-c` flag as well
```
chk file.txt file2.txt

# output:
file:    file.txt
md5:     01234516b98bac202ecaee7ac4001234
sha1:    012343ee5e6b4b0d3255bfef95601890afd01234
sha256:  0123444298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b78501234
---
file:    file2.txt
md5:     01234516b98bac202ecaee7ac4001234
sha1:    012343ee5e6b4b0d3255bfef95601890afd01234
sha256:  0123444298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b78501234

```
