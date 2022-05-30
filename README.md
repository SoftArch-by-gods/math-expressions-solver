# Math expressions solver
Program works with the prefix form of mathematical expressions  

To run program use:
```
go run ./cmd/example
```
with some flags:
```
 -e "- 4 2"         - read expression from cli
 -f inputFile.txt   - read expression from file
 -o outputFile.txt  - output file for expression result
```

```
Input:                            Output:
/ * + 0.5 1.5 7 8                 1.75
+ 2 / 31.35 * 5 + - 5 3 ^ 5 4     2.01
+ 10 9 6                          expression syntax error
h 9 6                             undefined symbol
```

Tests work out:  
https://github.com/SoftArch-by-gods/math-expressions-solver/actions/runs/2386188659  
Tests dropped:  
https://github.com/SoftArch-by-gods/math-expressions-solver/actions/runs/2386195017  
Testing pull request:  
https://github.com/SoftArch-by-gods/math-expressions-solver/actions/runs/2386215712
