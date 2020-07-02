# 指针
var a int = 2

var pa *int = &a

*pa = 5

# 参数传递
值传递， 引用传递

{}
# slice (切片)

arr := [...]int{1,2,3,4}

```
func (s []int, a [3]int){
    // s is slice, a is arr
}
```