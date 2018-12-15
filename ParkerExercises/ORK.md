### Rewrite Rules
```
* xR -> xRK
* Ox -> Oxx
* RRR -> K
* KK -> 
```

In the above rules, `x` can be any string.


## Some Rewrites
```
OR
* OR -> ORR
* ORR -> ORRRR
* ORRRR -> ORRRRRRRR
* ORRRRRRRR -> OKKRR
* OKKRR -> ORR
* ORR -> ORKR
* ORKR -> ORKKR
* ORKKR -> ...
```

From the above rewriting, it looks like OR can not be reduced to OK, it gets stuck in an infinite loop.
