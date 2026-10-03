# Struct ใน Go

> **Struct** = กล่องที่เอาข้อมูลหลายชนิดมารวมกันเป็นสิ่งเดียว

## 1. Struct คืออะไร

สมมติเรามี User ที่ประกอบไปด้วย `name`, `age`, `email`

แต่ละตัวจะมี type เป็นของตัวเอง

```go
name  string
age   int
email string
```

เราสามารถรวมมันเป็น User ได้ ที่เก็บไว้ใน struct

```go
type User struct {
    name  string
    age   int
    email string
}
```

ตอนนี้ `User` จะนับว่าเป็น type ใหม่ของเราเอง หรือเรียกว่า **type struct**

| ส่วน      | ความหมาย                     |
| -------- | --------------------------- |
| `type`   | บอกว่าเรากำลังสร้าง type ใหม่   |
| `User`   | ชื่อ type                     |
| `struct` | บอกว่า type นี้เป็น struct      |

ส่วนที่อยู่ใน `{ }` เรียกว่า **field**

```go
{
    name  string
    age   int
    email string
}
```

**ดังนั้น** `User` จะเป็น data type ใหม่ ที่เหมือน `string`, `int`, `bool` แต่ type `User` คือ type ที่เราออกแบบเองสร้างเอง

## 2. สร้างตัวแปรจาก struct

การสร้างตัวแปรจาก struct ก็เหมือนสร้างตัวแปรเหมือน type อื่น

**การสร้างตัวแปรจาก type built-in**

```go
var a int
var b string
```

**การสร้างตัวแปรจาก struct** — เมื่อเราสร้าง struct `User` แล้ว เราประกาศ

```go
var u User
```

ถ้าเรา `fmt.Printf("%+v\n", u)` ออกมาก็จะได้ **zero value** ของแต่ละ field เพราะเรายังไม่ได้กำหนดค่าให้มัน

```
name  → ""
age   → 0
email → ""
```

## 3. การเข้าถึง field

การเข้าถึง field เราจะใช้ `.` เช่น `u.age` หมายถึง เอา field `age` ของ `u`

```go
var u User

u.name = "John"
u.age = 25
u.email = "john@gmail.com"

fmt.Println(u.name)
fmt.Println(u.age)
fmt.Println(u.email)
```

Output

```
John
25
john@gmail.com
```

## 4. Anonymous Struct

คือ สร้าง struct โดยไม่ตั้งชื่อ type

```go
user := struct {
    name string
    age  int
}{
    name: "John",
    age:  25,
}
```

### Anonymous Struct ใช้ทำอะไร

มันเหมาะกับข้อมูลที่ใช้ **เฉพาะจุดนั้น**

```go
config := struct {
    Host string
    Port int
}{
    Host: "localhost",
    Port: 8080,
}
```

ถ้าไม่ได้จำเป็นต้องเอาไปใช้ที่อื่นก็สามารถใช้ Anonymous Struct ได้เลย

## 5. Named vs Unnamed Types

**Named**

```go
type User struct {
    Name string
    Age  int
}
```

**Unnamed**

```go
struct {
    Name string
    Age  int
}
```

## 6. Struct กิน Memory เท่าไหร่

```go
type Example struct {
    flag    bool
    counter int16
    pi      float32
}
```

เรารู้ว่า

| Type      | ขนาด    |
| --------- | ------- |
| `bool`    | 1 byte  |
| `int16`   | 2 bytes |
| `float32` | 4 bytes |

ถ้าคิดแบบง่าย ๆ `1 + 2 + 4 = 7 bytes`

แต่ Go บอกว่า

```go
fmt.Println(unsafe.Sizeof(Example{})) // 8
```

Output เป็น **8** เพราะมีสิ่งที่เรียกว่า **Padding**

### Padding คืออะไร

ลองจินตนาการว่า Memory เป็นบ้านที่มีเลขที่ `100` – `107`

1. `bool` ใช้แค่ 1 byte — สมมติเริ่มที่ `100` → `bool` ใช้บ้านที่ `100`
2. ต่อไปคือ `int16` — ต้องใช้ 2 bytes และต้องเริ่มที่ address ที่เหมาะสมกับ **2-byte alignment**
3. ดังนั้นมันไม่สามารถเริ่มที่ `101` แต่ต้องขยับไปที่ `102` จึงเกิดช่องว่าง:

```
100 → bool
101 → padding
102 → int16
103 → int16
104 → float32
105 → float32
106 → float32
107 → float32
```

```
1 byte bool
1 byte padding
2 byte int16
4 byte float32
──────────────
= 8 bytes
```

![Memory layout ของ Example: flag, padding, counter, pi รวม 8 bytes](https://tour.ardanlabs.com/tour/eng/static/img/f1.png)

นี่คือเหตุผลที่ `1 + 2 + 4` ไม่ได้เท่ากับขนาด struct จริงเสมอไป

### Alignment คืออะไร

ข้อมูลบางชนิดอยากเริ่มต้นที่ address ที่**หารด้วยขนาดของมันได้ลงตัว**

| Type      | ขนาด    | ชอบ address แบบ               |
| --------- | ------- | ---------------------------- |
| `int16`   | 2 bytes | `100`, `102`, `104`, `106`, `108` |
| `float32` | 4 bytes | `100`, `104`, `108`, `112`   |

ดังนั้นถ้าเราอยู่ตรง `101` แล้วต้องวาง `float32` ต้องขยับไป `104` ทำให้ `101`, `102`, `103` กลายเป็น padding

### ลำดับ field มีผลกับ memory

นี่คือเหตุผลว่าทำไมลำดับการเรียง field ใน struct ถึงมีผลกับการใช้ memory

```go
type Example2 struct {
    flag    bool
    counter int16
    flag2   bool
    pi      float32
}
```

เรามาลองวาง memory กัน สมมติเริ่มที่ `0x20`

1. **flag** → `0x20` ใช้ 1 byte
2. **counter** (`int16`) → ต่อไปคือ `0x21` แต่ `int16` ต้องการ 2-byte alignment, `0x21` ไม่เหมาะ จึงต้องขยับไป `0x22`
3. **flag2** → `0x24` ใช้ 1 byte
4. **pi** (`float32`) → ต่อไปคือ `0x25` แต่ `float32` ต้องการ 4-byte alignment เราต้องไป address ถัดไปที่เหมาะสมคือ `0x28` ทำให้ `0x25`–`0x27` เป็น padding

สุดท้าย:

```
0x20 ─ flag
0x21 ─ padding
0x22 ─ counter
0x23 ─ counter
0x24 ─ flag2
0x25 ─ padding
0x26 ─ padding
0x27 ─ padding
0x28 ─ pi
0x29 ─ pi
0x2A ─ pi
0x2B ─ pi
```

รวม: **12 bytes**

![Memory layout ของ Example2: flag, padding, counter, flag2, padding 3 bytes, pi รวม 12 bytes](https://tour.ardanlabs.com/tour/eng/static/img/f2.png)

### Rule ที่ควรจำ

สำหรับ struct ที่ต้องการประหยัด memory: โดยทั่วไปให้วาง field ที่มี alignment ใหญ่ก่อน แล้วเล็กลงมา

```go
type Data struct {
    A int64 // 8
    B int32 // 4
    C int16 // 2
    D bool  // 1
}
```

ตัวอย่าง: ถ้าเอา `Example2` มาเรียง field ใหม่จากใหญ่ไปเล็ก จะไม่มี padding เลย และเหลือแค่ **8 bytes** (จากเดิม 12 bytes)

```go
type Example2 struct {
    pi      float32 // 4
    counter int16   // 2
    flag    bool    // 1
    flag2   bool    // 1
}
```

![Memory layout ของ Example2 หลังเรียง field ใหม่: pi, counter, flag, flag2 รวม 8 bytes](https://tour.ardanlabs.com/tour/eng/static/img/f3.png)

> **แต่ขอเน้น:** อย่าเอา rule นี้ไปจัดทุก struct แบบไม่คิด
> ใน application ปกติ readability ก็สำคัญกว่า memory optimization

เรื่องนี้จะมีประโยชน์มากขึ้นเมื่อคุณทำ:

- high-throughput service
- large arrays/slices ของ struct
- memory-sensitive application
- performance optimization

## 7. Named Type กับการ Assign

```go
type Example1 struct {
    flag    bool
    counter int16
    pi      float32
}

type Example2 struct {
    flag    bool
    counter int16
    pi      float32
}
```

หน้าตาเหมือนกัน 100% แต่ **named type ไม่เหมือนกัน**

```go
ex1 = ex2 // Compile error
```

แต่เราสามารถ **Convert** ได้ ถ้า struct มี structure ที่ compatible กัน:

```go
ex1 = Example1(ex2)
```

ตรงนี้เราแปลง type ไม่ใช่การ Assignment value แต่เป็นการแปลง `ex2` ให้เป็น `Example1` ก่อน แล้วค่อยเอาไปใส่ `ex1`

### แล้ว Anonymous Struct ล่ะ?

นี่เป็นจุดที่ Ardan Labs ต้องการให้เห็นความแตกต่าง

สมมติ:

```go
type Example1 struct {
    flag    bool
    counter int16
    pi      float32
}

var ex1 Example1
```

แต่ตัวที่สอง:

```go
var ex2 struct {
    flag    bool
    counter int16
    pi      float32
}
```

`ex2` เป็น **unnamed struct type** และ structure เหมือนกับ `Example1` ดังนั้น:

```go
ex1 = ex2 // ✅ สามารถทำได้
```

เพราะ Go สามารถ assign ระหว่าง unnamed struct กับ named struct ที่มี underlying structure หรือก็คือ field เหมือนกันนั่นแหละ ที่เข้ากันได้ในกรณีนี้

## สรุป

| Code                     | ความหมาย                       |
| ------------------------ | ------------------------------ |
| `type User struct {...}` | สร้าง named type               |
| `var u User`             | สร้าง User zero value          |
| `User{...}`              | struct literal                 |
| `u.Name`                 | เข้าถึง field                  |
| `struct {...}`           | anonymous/unnamed struct       |
| `Example1` vs `Example2` | คนละ named type                |
| `Example1(ex2)`          | type conversion                |
| padding                  | memory ที่เพิ่มเพื่อ alignment |

## แบบฝึกหัดรอบนี้

ลองทำโดย **ยังไม่เปิดเฉลย**

### ข้อ 1 — อันไหนกิน memory มากกว่า?

```go
type A struct {
    A bool
    B int64
    C bool
}
```

หรือ

```go
type B struct {
    A int64
    B bool
    C bool
}
```

ก่อนรัน ให้เดาขนาดก่อน

### ข้อ 2 — อันนี้ compile ได้ไหม?

```go
type UserA struct {
    Name string
}

type UserB struct {
    Name string
}

func main() {
    var a UserA
    var b UserB

    a = b
}
```

ถ้าไม่ได้ ทำไม?

### ข้อ 3 — อันนี้ล่ะ?

```go
type User struct {
    Name string
}

func main() {
    var a User

    var b struct {
        Name string
    }

    a = b

    fmt.Println(a)
}
```
