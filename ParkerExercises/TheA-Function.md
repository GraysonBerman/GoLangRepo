## Rules

```
a(0,n) = n+1
a(m+1,0) = a(m,1)
a(m+1,n+1) = a(m, a(m+1,n))
```

#### Examples
* a(1,2)

`a(1,2) -> a(0,a(1,1)) -> a(1,1)+1 -> a(0,a(1,0))+1 -> a(0,1)+1+1 -> 2+2 -> 4`

* a(3,4)
`125`

* a(2,2)
`7`

* a(3,12)
`32765` This was the biggest number I was able to produce using a computer


