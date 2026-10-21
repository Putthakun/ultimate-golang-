# Pointer ใน Go

## 1. เริ่มจาก "Value" ก่อน

ก่อนเรียน Pointer ต้องเข้าใจคำว่า **"Value"** ก่อน

```go
func main() {
    count := 10

    increment(count)

    fmt.Println(count)
}

func increment(n int) {
    n++
}
```

หลายคนคิดว่า: `increment(count)` ส่ง `count` เข้าไป แล้ว `n++` ก็น่าจะทำให้ `count` เป็น `11`

แต่ผลคือ:

```
10
```

เพราะ **Go pass by value เสมอ**

## 2. Pass by Value คืออะไร?

```go
x := 10
```

- `x` คือชื่อกล่อง
- `10` คือ value ที่อยู่ในกล่อง `x`

เวลาส่งค่า `x` ไปที่อื่น มันทำการ **copy** ค่า `10` ที่อยู่ในกล่อง `x` ส่งไป แสดงว่าจะมีค่าใหม่คือ `10` อีก value นึง ดังนั้นเมื่อ `++` กับค่าใหม่ที่ copy ไป value ในกล่องจะไม่มีการเปลี่ยนแปลง นั่นก็คือ `10` เหมือนเดิม แต่ตัวที่ copy ใหม่จะเป็น `11` แทน

## 3. แล้ว Pointer เข้ามาช่วยยังไง

เราไม่ส่งตัว value ไป แต่จะส่ง **address ของกล่อง** `x` ไป

```go
increment(&count)
```

เครื่องหมาย `&` คือ **address of**

### ลองดู `&`

```go
count := 10

fmt.Println(count)
fmt.Println(&count)
```

อาจได้:

```
10
0xc0000120a0
```

### แล้ว `*` คืออะไร?

```go
func increment(n *int)
```

`*int` → `n` เป็น pointer ที่ชี้ไปยัง `int`

`*` มี 2 หน้าที่หลัก:

```go
var p *int
      ↑
   บอก TYPE
```

แต่

```go
fmt.Println(*p)
            ↑
     เอาค่าที่ pointer ชี้
```

> **จำง่าย ๆ:**
>
> | เขียน   | อ่านว่า            | ความหมาย                   |
> | ------ | ----------------- | -------------------------- |
> | `*int` | "pointer to int"  | บอกชนิด                     |
> | `*p`   | "value at p"      | เข้าไปเอาค่าที่ pointer ชี้     |

## 4. Struct + Pointer

```go
type User struct {
    Name string
    Age  int
}

user := User{
    Name: "John",
    Age:  20,
}

func Birthday(u User) {
    u.Age++
}

Birthday(user)
fmt.Println(user.Age) // 20
```

เพราะ struct ถูก copy:

```
user                 u
┌─────────────┐      ┌─────────────┐
│ John        │      │ John        │
│ 20          │ ──→  │ 20          │
└─────────────┘      └─────────────┘
```

หลัง `u.Age++` กลายเป็น:

```
user                 u
┌─────────────┐      ┌─────────────┐
│ John        │      │ John        │
│ 20          │      │ 21          │
└─────────────┘      └─────────────┘
```

## 5. Pointer ไม่ได้มีไว้แค่แก้ค่า

อีกเหตุผลหนึ่งคือ **ไม่ต้อง copy value ขนาดใหญ่โดยไม่จำเป็น** มันแค่ส่ง address ขนาดเล็กไปได้เลย

> ⚠️ อย่างไรก็ตาม อย่าเพิ่งสรุปว่า pointer เร็วกว่าเสมอ เพราะ compiler มี optimization และ pointer ก็มี trade-off เรื่อง allocation/GC

## 6. Escape Analysis คืออะไร

Go ต้องตัดสินใจว่า value ควรอยู่ที่ **Stack** หรือ **Heap**

Compiler ใช้สิ่งที่เรียกว่า **Escape Analysis**
