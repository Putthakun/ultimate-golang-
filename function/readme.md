# Functions ใน Go

ฟังก์ชันคือชุดคำสั่งที่เรียกใช้ซ้ำได้ โดยอาจรับข้อมูลผ่าน **parameters** และส่งผลลัพธ์กลับผ่าน **return values** บทนี้เริ่มจากการประกาศฟังก์ชัน แล้วต่อด้วยการคืนค่าหลายตัว, closures และการจัดการลำดับการทำงาน

> ตัวอย่างแต่ละหัวข้อแยกจากกัน บางตัวอย่างละ `package main`, `import` และ `func main()` เพื่อเน้นแนวคิด หากทดลองรัน ให้ใส่คำสั่งเรียกใช้ใน `main()` และ import แพ็กเกจที่ใช้ เช่น `fmt` หรือ `errors` อย่านำตัวอย่างที่ประกาศชื่อซ้ำมาต่อกันใน scope เดียว

## 1. การประกาศและเรียกใช้ฟังก์ชัน

```go
func add(x int, y int) int {
    return x + y
}
```

| ส่วนของโค้ด | ความหมาย |
| --- | --- |
| `func` | คำสั่งประกาศฟังก์ชัน |
| `add` | ชื่อฟังก์ชัน |
| `x int, y int` | Parameters หรือตัวแปรที่รับค่าจากผู้เรียก |
| `int` หลังวงเล็บ | ชนิดของค่าที่ส่งกลับ |
| `return x + y` | ส่งผลบวกกลับให้ผู้เรียก |

เรียกใช้และเก็บผลลัพธ์:

```go
result := add(10, 20)
fmt.Println(result) // 30
```

`x` และ `y` คือ **parameters** ส่วน `10` และ `20` คือ **arguments** หรือค่าที่ส่งเข้าไปตอนเรียก

### Parameters ที่มีชนิดเดียวกัน

ถ้า parameters ที่อยู่ติดกันมีชนิดเดียวกัน สามารถเขียนย่อได้:

```go
func add(x, y int) int {
    return x + y
}
```

### ฟังก์ชันที่ไม่คืนค่า

ฟังก์ชันไม่จำเป็นต้องคืนค่าเสมอไป หากไม่มีผลลัพธ์ ให้ละชนิดผลลัพธ์หลังวงเล็บ:

```go
func greet(name string) {
    fmt.Println("Hello", name)
}
```

เรียกใช้โดยไม่ต้องมีตัวแปรรับผลลัพธ์:

```go
greet("Mod") // Hello Mod
```

## 2. การคืนค่าหลายตัวและจัดการ error

Go สามารถคืนค่าหลายตัวจากฟังก์ชันเดียวได้ รูปแบบที่พบบ่อยคือ **ผลลัพธ์ + error**

```go
func divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, errors.New("division by zero")
    }

    return x / y, nil
}
```

ฟังก์ชันนี้คืนค่าสองตัวตามลำดับ:

1. `float64`: ผลหาร หรือ `0` ในกรณีที่ตัวอย่างนี้ทำงานไม่สำเร็จ
2. `error`: ข้อผิดพลาด หรือ `nil` เมื่อสำเร็จ

ผู้เรียกควรตรวจ error ก่อนใช้ผลลัพธ์:

```go
result, err := divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
    return
}

fmt.Println(result) // 5
```

ถ้าเปลี่ยนเป็น `divide(10, 0)` จะเข้าเงื่อนไข `err != nil` และพิมพ์ `Error: division by zero`

### `error` เป็น type ไม่ใช่ค่าข้อผิดพลาด

`error` เป็น interface ที่ Go มีให้ในภาษา มีนิยามดังนี้:

```go
type error interface {
    Error() string
}
```

นิยามด้านบนมีไว้ให้อ่านทำความเข้าใจ ไม่ต้องประกาศซ้ำในโปรแกรม ชนิดข้อมูลที่มี method `Error() string` สามารถใช้เป็นค่า `error` ได้

ต้องแยก **ชื่อชนิดข้อมูล** ออกจาก **ค่า** ที่เป็นชนิดนั้น:

```go
var message string = "hello"
var err error = errors.New("ไม่พบผู้ใช้")
```

ตัวอย่างคำสั่ง `return` สำหรับฟังก์ชันที่ประกาศผลลัพธ์เป็น `(string, error)`:

```go
// ผิด: 0 ใช้เป็น string ไม่ได้ และ error เป็นชื่อ type
return 0, error

// ถูก: คืน string และค่า error ที่สร้างขึ้น
return "", errors.New("ไม่พบผู้ใช้")

// เมื่อสำเร็จ: คืนผลลัพธ์และ nil
return "hello", nil
```

แต่ละ `return` ด้านบนเป็นตัวอย่างแยกกัน ใช้ `errors.New` เพื่อสร้าง error จากข้อความ หรือใช้ `fmt.Errorf` เมื่อต้องการจัดรูปแบบ เช่น `fmt.Errorf("invalid user ID: %d", userID)`

## 3. Blank identifier และการใช้ `:=` ซ้ำ

### ใช้ `_` ทิ้งค่าที่ไม่ต้องการ

เมื่อนำผลลัพธ์หลายค่ามากำหนดให้ตัวแปร ต้องมีตำแหน่งรับค่าครบ ถ้าไม่ใช้บางค่า ให้ใช้ **blank identifier** หรือ `_` ได้:

```go
func getUser() (string, int) {
    return "Mod", 24
}
```

ถ้าต้องการเฉพาะชื่อ:

```go
name, _ := getUser()
fmt.Println(name) // Mod
```

> ไม่ควรใช้ `_` ทิ้ง error โดยไม่มีเหตุผลที่เหมาะสม เพราะอาจทำให้ไม่รู้ว่าการทำงานล้มเหลว

### ใช้ `:=` กับตัวแปรเดิมและตัวแปรใหม่

```go
name := "Mod"
name, age := getUser()

fmt.Println(name, age) // Mod 24
```

ใช้ได้ เพราะ `age` เป็นตัวแปรใหม่ ส่วน `name` เป็นตัวแปรเดิมใน scope เดียวกันและยังมีชนิดเดิม

แต่ตัวอย่างนี้ใช้ไม่ได้:

```go
name := "Mod"
name := "John" // compile error: ไม่มีตัวแปรใหม่ใน scope เดิม
```

ถ้าต้องการเปลี่ยนค่าตัวแปรเดิม ให้ใช้ `=`:

```go
name := "Mod"
name = "John"
fmt.Println(name) // John
```

จำว่า `:=` ต้องมีตัวแปรใหม่อย่างน้อยหนึ่งตัวที่ไม่ใช่ `_` กฎการใช้ตัวแปรเดิมนี้อ้างถึง **scope เดียวกัน** หากประกาศชื่อเดิมใน block ชั้นใน อาจเป็นตัวแปรคนละตัวที่บังชื่อเดิม

## 4. Variadic functions: รับค่าได้หลายตัวด้วย `...`

ฟังก์ชันปกติกำหนดจำนวน arguments ไว้ เช่น `add(a, b int)` รับสองตัว แต่ variadic function รับจำนวน arguments ที่เปลี่ยนไปได้:

```go
func sum(nums ...int) int {
    total := 0

    for _, num := range nums {
        total += num
    }

    return total
}
```

เรียกใช้ได้หลายรูปแบบ รวมถึงไม่ส่งค่าเลย:

```go
fmt.Println(sum())              // 0
fmt.Println(sum(10, 20))        // 30
fmt.Println(sum(10, 20, 30))    // 60
fmt.Println(sum(1, 2, 3, 4, 5)) // 15
```

ภายในฟังก์ชัน `nums` มีชนิดเป็น `[]int` จึงใช้ `range` วนอ่านสมาชิกได้

ถ้ามี slice อยู่แล้ว ให้เติม `...` ตอนส่งเข้าไป:

```go
numbers := []int{10, 20, 30}
result := sum(numbers...)

fmt.Println(result) // 60
```

`numbers...` ส่งสมาชิกของ slice เป็น variadic arguments และ variadic parameter ต้องอยู่ท้ายสุดของรายการ parameters

## 5. Anonymous functions และ closures

### ฟังก์ชันที่ไม่มีชื่อ

นอกจากการประกาศ `func add(...)` แล้ว ยังสร้างฟังก์ชันที่ไม่มีชื่อและเก็บไว้ในตัวแปรได้:

```go
add := func(a, b int) int {
    return a + b
}

fmt.Println(add(10, 20)) // 30
```

ในที่นี้ `add` คือชื่อตัวแปรที่เก็บค่าฟังก์ชัน และเรียกใช้ผ่านตัวแปรนั้นได้

### Closure: เข้าถึงตัวแปรจาก scope ภายนอก

```go
func counter() func() int {
    count := 0

    return func() int {
        count++
        return count
    }
}
```

`counter` คืนฟังก์ชันชนิด `func() int` ซึ่งไม่รับ arguments และคืน `int` หนึ่งตัว:

```go
next := counter()

fmt.Println(next()) // 1
fmt.Println(next()) // 2
fmt.Println(next()) // 3
```

`count` ไม่กลับเป็น `0` ทุกครั้ง เพราะ `next()` เรียกฟังก์ชันที่คืนออกมา ไม่ได้เรียก `counter()` ใหม่ ฟังก์ชันที่คืนออกมายังเข้าถึงตัวแปร `count` ตัวเดิมได้ แม้ `counter()` จะทำงานจบแล้ว

หากเรียก `counter()` อีกครั้ง จะได้ตัวนับที่มี `count` แยกจากชุดแรก:

```go
another := counter()
fmt.Println(another()) // 1
fmt.Println(next())    // 4
```

Closures ใช้ได้กับ callbacks, middleware และฟังก์ชันที่ต้องเก็บสถานะไว้ระหว่างการเรียก

## 6. Function types: ส่งฟังก์ชันเป็นค่า

ฟังก์ชันเป็น value จึงเก็บในตัวแปรหรือส่งเป็น argument ได้ และสามารถตั้งชื่อชนิดฟังก์ชันไว้ใช้ซ้ำ:

```go
type MathFunc func(int, int) int
```

`MathFunc` คือชนิดของฟังก์ชันที่รับ `int` สองตัวและคืน `int` หนึ่งตัว

```go
func add(a, b int) int {
    return a + b
}

func multiply(a, b int) int {
    return a * b
}

func calculate(a, b int, operation MathFunc) int {
    return operation(a, b)
}
```

เรียกใช้โดยส่งตัวฟังก์ชันเข้าไป:

```go
fmt.Println(calculate(10, 5, add))      // 15
fmt.Println(calculate(10, 5, multiply)) // 50
```

`add` หมายถึงตัวฟังก์ชัน ส่วน `add(10, 5)` หมายถึงการเรียกแล้วนำผลลัพธ์มาใช้ ในตัวอย่างนี้ `calculate` ต้องการรับฟังก์ชัน จึงส่ง `add` โดยไม่เติมวงเล็บเรียก

ฟังก์ชันที่รับหรือคืนฟังก์ชันเรียกว่า **higher-order function** ตัวอย่างในบทนี้คือ `calculate` และ `counter` แนวคิดนี้เป็นพื้นฐานหนึ่งของการออกแบบ middleware

## 7. Defer, panic และ recover

### `defer`: เลื่อนการเรียกไว้ก่อนฟังก์ชันปัจจุบันจบ

```go
func main() {
    defer fmt.Println("Third")

    fmt.Println("First")
    fmt.Println("Second")
}
```

ผลลัพธ์:

```text
First
Second
Third
```

ถ้ามีหลาย `defer` จะทำงานย้อนลำดับแบบ **LIFO — Last In, First Out**:

```go
func main() {
    defer fmt.Println("A")
    defer fmt.Println("B")
    defer fmt.Println("C")
}
```

ผลลัพธ์:

```text
C
B
A
```

### `panic`: หยุดลำดับการทำงานตามปกติ

`panic` หยุดการทำงานตามปกติ แล้วคลาย call stack พร้อมเรียก deferred functions ที่เกี่ยวข้อง หากไม่มีการ recover โปรแกรมจะจบด้วยข้อผิดพลาด

```go
func main() {
    fmt.Println("Start")
    panic("something went wrong")

    // คำสั่งที่วางต่อจาก panic ตรงนี้จะไม่ทำงาน
}
```

โดยทั่วไปควรคืน `error` สำหรับปัญหาที่ผู้เรียกคาดการณ์และจัดการได้ เช่น input ไม่ถูกต้อง ไม่ใช้ `panic` แทน error ทุกกรณี

### `recover`: รับมือ panic ภายใน goroutine เดียวกัน

`recover()` ต้องถูกเรียกโดยตรงภายใน deferred function ของ goroutine ที่กำลัง panic จึงจะหยุด panic นั้นได้:

```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()

    panic("unexpected error")
}
```

ผลลัพธ์:

```text
Recovered: unexpected error
```

หลัง recover สำเร็จ จะไม่กลับไปทำคำสั่งถัดจากจุดที่ panic ในฟังก์ชันนั้น ในตัวอย่างนี้ deferred function ทำงานเสร็จแล้ว `main` ก็จบ

ใน web server มักใช้ recovery middleware จัดการ panic จาก request handler โดยต้องอยู่ใน goroutine เดียวกัน ไม่สามารถใช้ `recover` จากอีก goroutine มาจับแทนได้

## สรุปสิ่งที่ควรจำ

| เรื่อง | หลักสำคัญ |
| --- | --- |
| Parameters / arguments | ตัวแปรรับค่า / ค่าที่ส่งตอนเรียก |
| Return values | คืนค่าให้ตรงจำนวน ชนิด และลำดับที่ประกาศ |
| `error` | เป็น type; คืนค่า error หรือ `nil` ตามผลการทำงาน |
| `_` | ทิ้งค่าที่ไม่ใช้ โดยระวังการทิ้ง error |
| `:=` | ต้องมีตัวแปรใหม่อย่างน้อยหนึ่งตัวที่ไม่ใช่ `_` |
| `...` | Variadic parameter ต้องอยู่ท้ายสุด |
| Closure | ฟังก์ชันยังเข้าถึงตัวแปรจาก scope ภายนอกได้ |
| Function type | กำหนดรูปแบบ parameters และ return values |
| `defer` | ทำงานก่อนฟังก์ชันจบ เรียงย้อนลำดับที่ลงทะเบียนไว้ |
| `panic` / `recover` | หยุดลำดับปกติ / จัดการ panic ผ่าน deferred function |

## อ่านเพิ่มเติม

- [Go specification: Function types](https://go.dev/ref/spec#Function_types)
- [Go specification: Short variable declarations](https://go.dev/ref/spec#Short_variable_declarations)
- [Go specification: Handling panics](https://go.dev/ref/spec#Handling_panics)
- [Go built-in type: error](https://pkg.go.dev/builtin#error)
