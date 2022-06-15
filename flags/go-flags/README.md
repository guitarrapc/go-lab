## Windows

will not use `-` or `--` but `/` will be used.


## output

```
Usage:
  C:\Users\IKIRU~1.YOS\AppData\Local\Temp\go-build550775398\b001\exe\main.exe [OPTIONS]

Create Options:
  /i, /title:<title>            The title of an issue
  /m, /message:<message>        The message of an issue

List Options:
  /n, /num:<num>                Limit the number of issue to output. (default:
                                20)
      /state:<state>            Print only issue of the state just those that
                                are "opened", "closed" or "all" (default: all)
      /scope:<scope>            Print only given scope. "created-by-me",
                                "assigned-to-me" or "all". (default: all)
      /orderby:<orderby>        Print issue ordered by "created_at" or
                                "updated_at" fields. (default: updated_at)
      /sort:<sort>              Print issue ordered in "asc" or "desc" order.
                                (default: desc)
  /s, /search:<search word>     Search issues against their title and
                                description.
```