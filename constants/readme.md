# Constants ใน Go

**Constant** คือค่าคงที่ที่ประกาศด้วย `const` โดยค่าต้องคำนวณได้ตอน compile และไม่สามารถกำหนดค่าใหม่ให้ชื่อเดิมได้

บทนี้ครอบคลุม typed และ untyped constants การนำค่าคงที่ไปใช้กับชนิดข้อมูลต่าง ๆ และการสร้างลำดับค่าด้วย `iota`

> ตัวอย่างแต่ละหัวข้อแยกจากกัน บางตัวอย่างละ `package main`, `import "fmt"` และ `func main()` เพื่อเน้นแนวคิด หากทดลองรัน ให้ใส่คำสั่งใช้งาน เช่น `fmt.Println` ใน `main()` และอย่านำตัวอย่างที่ประกาศชื่อซ้ำมาต่อกันใน scope เดียว

## 1. `const` ต่างจาก `var` อย่างไร?

```go
const maxRetries = 3
var attempts = 0

attempts++ // เปลี่ยนค่าตัวแปรได้
// maxRetries++ // compile error: เปลี่ยนค่าคงที่ไม่ได้
```

| ประเด็น | `const` | `var` |
| --- | --- | --- |
| ความหมาย | ค่าคงที่ | ตัวแปร |
| เปลี่ยนค่าภายหลัง | ไม่ได้ | ได้ |
| ค่าเริ่มต้น | ต้องเป็น constant expression | ใช้ค่าคงที่หรือค่าที่ได้ตอน runtime ก็ได้ |
| ชนิดข้อมูล | เป็น typed หรือ untyped constant ได้ | ตัวแปรมีชนิดแน่นอน แม้ไม่ได้เขียนชนิดไว้เอง |

อย่าจำว่า “`var` รู้ค่าได้เฉพาะตอนโปรแกรมทำงาน” เพราะ `var attempts = 0` ก็เริ่มจากค่าคงที่ได้ จุดต่างคือ `const` ต้องมีค่าที่คำนวณได้ตอน compile และเปลี่ยนค่าภายหลังไม่ได้

## 2. Typed และ untyped constants

### Typed constant: ผูกกับชนิดข้อมูลแล้ว

```go
const age int = 24
```

`age` มีชนิดเป็น `int` จึงต้องปฏิบัติตามกฎการใช้ค่าของชนิดนั้น

### Untyped constant: ยังไม่ผูกกับชนิดข้อมูลเฉพาะ

```go
const age = 24
const pi = 3.141592
const greeting = "hello"
const enabled = true
```

ค่าคงที่เหล่านี้ยังไม่ผูกกับชนิดเฉพาะ เช่น `int`, `int64` หรือ `float64` โดยบริบทที่นำไปใช้จะกำหนดว่าต้องใช้เป็นชนิดใด

คำว่า **kind** ในบทนี้ใช้อธิบายหมวดของค่าคงที่ ไม่ใช่ชื่อ type ที่ต้องเขียนในโค้ด:

| ค่า | หมวดของค่าคงที่ | Default type เมื่อบริบทต้องการชนิด แต่ไม่ได้ระบุไว้ |
| --- | --- | --- |
| `24` | Integer | `int` |
| `3.141592` | Floating-point | `float64` |
| `"hello"` | String | `string` |
| `true` | Boolean | `bool` |

ตัวอย่างการใช้ default type:

```go
const x = 10
value := x

fmt.Printf("%T\n", value) // int
```

`x` ยังคงเป็น untyped constant ส่วน `value` เป็นตัวแปรชนิด `int`

## 3. นำ untyped constant ไปใช้กับหลาย type

```go
const x = 10

var a int = x
var b float64 = x

fmt.Println(a) // 10
fmt.Println(b) // 10
```

ใช้ได้ เพราะ `x` เป็น untyped constant และค่า `10` ใช้ได้กับทั้ง `int` และ `float64` การกำหนดค่าให้ `a` และ `b` ไม่ได้เปลี่ยนให้ `x` มี type ใด type หนึ่งอย่างถาวร

แต่ถ้า `x` มี type เป็น `int` แล้ว:

```go
const x int = 10

var a int = x
var b float64 = x // compile error: ไม่แปลง int เป็น float64 ให้อัตโนมัติ
```

ต้องแปลงชนิดอย่างชัดเจน:

```go
var b float64 = float64(x)
```

### ค่าต้องแทนได้ในชนิดปลายทาง

Untyped ไม่ได้หมายความว่าใช้กับ type ไหนก็ได้:

```go
const small = 127
const large = 128

var a int8 = small // ใช้ได้: int8 มีช่วง -128 ถึง 127
var b int8 = large // compile error: 128 เกินช่วงของ int8
```

เช่นเดียวกับค่าทศนิยมที่มีเศษ ซึ่งนำไปกำหนดให้ตัวแปร `int` โดยตรงไม่ได้:

```go
const pi = 3.141592
var n int = pi // compile error: ค่านี้แทนเป็น int ไม่ได้
```

## 4. Precision และค่าคงที่ขนาดใหญ่

```go
const maxInt = 9223372036854775807 // ค่าสูงสุดของ int64
const bigger = 9223372036854775808543522345
```

แม้ `bigger` จะใหญ่กว่าค่าที่ `int64` เก็บได้ แต่ยังประกาศเป็น untyped constant ได้ เพราะค่าคงที่ตัวเลขไม่ได้ถูกจำกัดด้วยช่วงของ `int64` ตั้งแต่ตอนประกาศ

ข้อจำกัดของชนิดปลายทางเข้ามาเมื่อพยายามนำค่าไปใช้เป็นชนิดนั้น:

```go
var value int64 = bigger // compile error: เกินช่วงของ int64
```

รวมถึงการพิมพ์ค่าคงที่ขนาดใหญ่นี้โดยตรง:

```go
fmt.Println(bigger) // compile error: default type คือ int และค่านี้เกินช่วง
```

`fmt.Println` รับค่าผ่าน interface ทำให้ untyped integer constant ต้องถูกใช้เป็น default type คือ `int` ก่อน จึงเกิด overflow ในตัวอย่างนี้

> ตามความหมายของภาษา ค่าคงที่ตัวเลขมีความแม่นยำแบบ arbitrary precision แต่ compiler สามารถมีขีดจำกัดภายในได้ อย่าสรุปว่าประกาศตัวเลขใหญ่เท่าใดก็ได้โดยไม่มีข้อจำกัด

## 5. ชนิดของค่าคงที่ในนิพจน์

เมื่อใช้ typed constant ร่วมกับ untyped constant ในการคำนวณ:

```go
const one int8 = 1
const two = 2 * one

fmt.Printf("%T\n", two) // int8
fmt.Println(two)        // 2
```

`one` มีชนิด `int8` ส่วน `2` เป็น untyped constant ที่ใช้เป็น `int8` ได้ ผลลัพธ์ `two` จึงมีชนิดเป็น `int8` ด้วย แม้ไม่ได้เขียน type ในบรรทัดประกาศ `two` โดยตรง

ดังนั้น **การไม่เขียน type หลังชื่อ constant ไม่ได้ทำให้เป็น untyped เสมอไป** ต้องดูนิพจน์ที่ใช้กำหนดค่าด้วย

## 6. `iota`: สร้างค่าตามลำดับ

`iota` ใช้ภายในกลุ่มประกาศ `const` เพื่อสร้างค่าจำนวนเต็มตามลำดับ:

```go
const (
    Pending = iota
    Processing
    Success
    Failed
)
```

```go
fmt.Println(Pending)    // 0
fmt.Println(Processing) // 1
fmt.Println(Success)    // 2
fmt.Println(Failed)     // 3
```

ในแต่ละ `const` block ค่า `iota` เริ่มจาก `0` และเพิ่มตามลำดับรายการประกาศ หรือ `ConstSpec` ไม่ใช่ตามจำนวนบรรทัดว่างหรือคอมเมนต์ ส่วนรายการที่ละนิพจน์ไว้จะใช้นิพจน์จากรายการก่อนหน้าซ้ำ โดยประเมิน `iota` เป็นค่าของรายการปัจจุบัน

### เริ่มลำดับจาก 1

```go
const (
    Pending = iota + 1
    Processing
    Success
    Failed
)
```

ได้ค่า `1`, `2`, `3`, `4` ตามลำดับ เพราะแต่ละรายการใช้รูปนิพจน์ `iota + 1` เหมือนกัน

### หลายชื่อในรายการเดียวกัน

```go
const (
    A, B = iota, iota // 0, 0
    C, D             // 1, 1
)
```

`A` และ `B` อยู่ในรายการประกาศเดียวกัน จึงเห็นค่า `iota` เท่ากัน

## 7. `iota` กับ bit shifting

ใช้ `1 << iota` เพื่อสร้างค่าที่มีบิตเปิดอยู่คนละตำแหน่ง:

```go
const (
    Ldate = 1 << iota
    Ltime
    Lmicroseconds
    Llongfile
)
```

| ชื่อ | นิพจน์ | ฐานสอง | ฐานสิบ |
| --- | --- | --- | --- |
| `Ldate` | `1 << 0` | `0001` | `1` |
| `Ltime` | `1 << 1` | `0010` | `2` |
| `Lmicroseconds` | `1 << 2` | `0100` | `4` |
| `Llongfile` | `1 << 3` | `1000` | `8` |

รูปแบบนี้เหมาะกับ flags ที่เปิดหลายตัวพร้อมกันได้ โดยใช้ `|` รวมบิต และใช้ `&` ตรวจบิต:

```go
flags := Ldate | Ltime

fmt.Println(flags)                // 3
fmt.Println(flags&Ldate != 0)     // true
fmt.Println(flags&Llongfile != 0) // false
```

## สรุปสิ่งที่ควรจำ

- `const` ต้องคำนวณค่าได้ตอน compile และกำหนดค่าใหม่ไม่ได้
- Typed constant มีชนิดแน่นอน ส่วน untyped constant ใช้กับชนิดปลายทางได้ตามบริบทและข้อจำกัดของค่านั้น
- เมื่อบริบทไม่ได้ระบุชนิด อาจมีการใช้ default type เช่น `int` หรือ `float64`
- Constant ที่ไม่เขียน type โดยตรงอาจยังเป็น typed constant ได้ ขึ้นอยู่กับนิพจน์
- `iota` เริ่มจาก `0` ใหม่ในแต่ละ `const` block และเพิ่มตามลำดับรายการประกาศ
- `1 << iota` ใช้สร้าง flags ที่มีค่ากำลังของสอง เช่น `1`, `2`, `4`, `8`

## อ่านเพิ่มเติม

- [Go specification: Constants](https://go.dev/ref/spec#Constants)
- [Go specification: Constant expressions](https://go.dev/ref/spec#Constant_expressions)
- [Go specification: Iota](https://go.dev/ref/spec#Iota)
