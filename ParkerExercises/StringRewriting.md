# String Rewriting Exercise

#### Rewrite Rules
```go
ab -> ba
ba -> ab
aa -> 
b -> 
```

Strings to Rewrite :


* abba:
```
-> baba
-> baab
-> bb
-> b
-> b
```



* bababa:
```
-> abbaba
-> abbbaa
-> abbb
-> abb
-> ab
-> ba
...
```

In order for an ARS to be considered terminating, there needs to be a defined measure function. The ARS above does not contain
a measure function and therefore does not terminate. Because `ab -> ba` and `ba -> ab`, the string potentially remians the same length and stuck in a cycle of rewriting.
