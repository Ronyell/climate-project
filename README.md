Rules/Commands Go

### Commands:
 - `go run <file>` -> Run file
 - `go build` -> Compile project
 - `go install` -> Compile project (save it in root)
 - `go get <package>` -> install external package

### Rules
- Point of enter is in package main
- First letter Cap in func is public func
- First letter is not Cap in func is not public func
- Variable implicit is used with `:=` and attribute a value
- Variable explicit is used with a type after variable name and word var before variable name
- byte = uint32 and rune == int32
- To declare pointer use `var variable *type`, to reference another use `variable = *variable0`, to dereferencing use `&variable`
- Array use a fix length. Ex `var array1[5] int`
- Slice is similar to Array but dont have a prefixed length Ex `var slice1[] int`
