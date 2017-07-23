# chk
chk your file hashes

## Install

`go get github.com/LadyDascalie/chk`

Or just grab one of the [releases](https://github.com/LadyDascalie/chk/releases)

## Examples

```
chk file.txt

# output:
file:    file.txt
md5:     01234516b98bac202ecaee7ac4001234
sha1:    012343ee5e6b4b0d3255bfef95601890afd01234
sha256:  0123444298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b78501234
```

Checking multiple files works too:

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
