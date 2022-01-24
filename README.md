## bnfchecker 
### bnfchecker - tool to testing grammar in  [parsegen](https://github.com/kasaderos/parsegen)

### Read guide from grammar.md 

Usage:
```
    ./checker-win.exe -f <filename>
```

Example
```
    ./checker-win.exe -f request.bnf
```

input:

```
GET / 200
```

output:

```
Url :
            /
Method :
            GET
StatusOk :
            200
```