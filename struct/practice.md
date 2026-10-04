# Struct Practice

| Level | หัวข้อ                    | โจทย์   |
| ----- | ------------------------ | ------- |
| 🟢 1  | ทำความคุ้นเคยกับ Struct     | 1 – 3   |
| 🟡 2  | Struct + Filtering       | 4 – 6   |
| 🟠 3  | Struct + Map             | 7 – 9   |
| 🔴 4  | เริ่มคิดแบบ Backend         | 10      |
| 🔴 5  | Challenge                | 11      |

---

## 🟢 Level 1 — ทำความคุ้นเคยกับ Struct

### โจทย์ 1: หาคนที่อายุมากที่สุด ⭐

กำหนด struct:

```go
type Person struct {
    Name string
    Age  int
}
```

มีข้อมูล:

```go
people := []Person{
    {Name: "John", Age: 20},
    {Name: "Alice", Age: 25},
    {Name: "Bob", Age: 22},
}
```

ให้เขียน Function

```go
func FindOldest(people []Person) Person
```

ให้คืนค่าคนที่อายุมากที่สุด

**Expected:**

```
Alice
25
```

**สิ่งที่ได้ฝึก:** struct · slice · for loop · comparison · function

---

### 🟢 โจทย์ 2: หาคนตาม ID ⭐

```go
type User struct {
    ID   int
    Name string
}
```

ข้อมูล:

```go
users := []User{
    {ID: 101, Name: "John"},
    {ID: 102, Name: "Alice"},
    {ID: 103, Name: "Bob"},
}
```

เขียน:

```go
func FindUser(users []User, id int) User
```

ถ้า `FindUser(users, 102)` ให้ได้:

```
Alice
```

ถ้าไม่เจอ ให้คืน `User{}`

**ฝึกอะไร?** นี่เริ่มเหมือน Backend แล้ว:

```
Request
   ↓
หา User จาก ID
   ↓
Return User
```

---

### 🟢 โจทย์ 3: คำนวณเงินเดือนเฉลี่ย ⭐⭐

```go
type Employee struct {
    Name   string
    Salary float64
}
```

ข้อมูล:

```go
employees := []Employee{
    {Name: "John", Salary: 30000},
    {Name: "Alice", Salary: 40000},
    {Name: "Bob", Salary: 35000},
}
```

เขียน:

```go
func AverageSalary(employees []Employee) float64
```

**Expected:**

```
35000
```

---

## 🟡 Level 2 — Struct + Filtering

### 🟡 โจทย์ 4: หาสินค้าที่มีราคามากกว่า 1,000 ⭐⭐

```go
type Product struct {
    Name  string
    Price float64
}
```

ข้อมูล:

```go
products := []Product{
    {Name: "Keyboard", Price: 1500},
    {Name: "Mouse", Price: 800},
    {Name: "Monitor", Price: 5000},
    {Name: "Cable", Price: 300},
}
```

เขียน:

```go
func FilterExpensiveProducts(products []Product) []Product
```

ให้คืนเฉพาะสินค้าที่ `Price > 1000`

**Expected:**

```
Keyboard
Monitor
```

---

### 🟡 โจทย์ 5: หาสินค้าที่หมด Stock ⭐⭐

```go
type Product struct {
    Name  string
    Price float64
    Stock int
}
```

ให้:

```go
products := []Product{
    {Name: "Keyboard", Price: 1500, Stock: 10},
    {Name: "Mouse", Price: 800, Stock: 0},
    {Name: "Monitor", Price: 5000, Stock: 3},
    {Name: "Cable", Price: 300, Stock: 0},
}
```

สร้าง:

```go
func FindOutOfStock(products []Product) []Product
```

**Expected:**

```
Mouse
Cable
```

---

### 🟡 โจทย์ 6: หาสินค้าราคารวมสูงสุด ⭐⭐

```go
type Product struct {
    Name     string
    Price    float64
    Quantity int
}
```

ตัวอย่าง:

```go
products := []Product{
    {
        Name:     "Keyboard",
        Price:    1500,
        Quantity: 2,
    },
    {
        Name:     "Monitor",
        Price:    5000,
        Quantity: 3,
    },
    {
        Name:     "Mouse",
        Price:    800,
        Quantity: 5,
    },
}
```

ให้คำนวณ `Price × Quantity` แล้วหาสินค้าที่มีมูลค่ารวมสูงสุด

**Expected:**

```
Monitor
15000
```

> อันนี้เริ่มฝึกการคิด Algorithm มากขึ้น

---

## 🟠 Level 3 — Struct + Map

> ตรงนี้ **แนะนำมาก** เพราะจะเริ่มใกล้กับงาน Backend

### 🟠 โจทย์ 7: นับจำนวนสินค้าแต่ละ Category ⭐⭐⭐

```go
type Product struct {
    Name     string
    Category string
}
```

ข้อมูล:

```go
products := []Product{
    {Name: "Keyboard", Category: "Computer"},
    {Name: "Mouse", Category: "Computer"},
    {Name: "iPhone", Category: "Mobile"},
    {Name: "iPad", Category: "Mobile"},
    {Name: "MacBook", Category: "Computer"},
}
```

สร้าง:

```go
func CountCategory(products []Product) map[string]int
```

**Expected:**

```
Computer → 3
Mobile   → 2
```

**สิ่งที่ฝึก:** Struct + Slice + Map + Loop

นี่เป็น combination ที่ควรคล่องมากใน Go

---

### 🟠 โจทย์ 8: Group User ตาม Department ⭐⭐⭐

```go
type Employee struct {
    Name       string
    Department string
}
```

ข้อมูล:

```go
employees := []Employee{
    {Name: "John", Department: "IT"},
    {Name: "Alice", Department: "HR"},
    {Name: "Bob", Department: "IT"},
    {Name: "Jane", Department: "Marketing"},
    {Name: "Tom", Department: "IT"},
}
```

สร้าง:

```go
func GroupByDepartment(
    employees []Employee,
) map[string][]Employee
```

**Expected:**

```
IT:
    John
    Bob
    Tom

HR:
    Alice

Marketing:
    Jane
```

โจทย์นี้ **ดีมากสำหรับคุณ** เพราะเป็น pattern ที่เจอใน Backend บ่อย:

```
Database
   ↓
[]Employee
   ↓
Group
   ↓
map[string][]Employee
   ↓
API Response
```

---

### 🟠 โจทย์ 9: หาคนที่ซื้อของเยอะที่สุด ⭐⭐⭐

```go
type Order struct {
    UserID   int
    Product  string
    Quantity int
}
```

ข้อมูล:

```go
orders := []Order{
    {UserID: 1, Product: "Keyboard", Quantity: 2},
    {UserID: 2, Product: "Mouse", Quantity: 5},
    {UserID: 1, Product: "Monitor", Quantity: 1},
    {UserID: 2, Product: "Keyboard", Quantity: 3},
}
```

รวมจำนวนสินค้าที่แต่ละ user ซื้อ

**Expected:**

```
User 1 → 3
User 2 → 8
```

แล้วหา: `User 2` เพราะซื้อทั้งหมด 8 ชิ้น

**นี่ฝึก:** `map[int]int`

---

## 🔴 Level 4 — เริ่มคิดแบบ Backend

### 🔴 โจทย์ 10: Transaction Summary ⭐⭐⭐⭐

ให้:

```go
type Transaction struct {
    ID     int
    UserID int
    Amount float64
    Type   string
}
```

ตัวอย่าง:

```go
transactions := []Transaction{
    {ID: 1, UserID: 101, Amount: 1000, Type: "deposit"},
    {ID: 2, UserID: 101, Amount: 500, Type: "withdraw"},
    {ID: 3, UserID: 102, Amount: 2000, Type: "deposit"},
    {ID: 4, UserID: 101, Amount: 300, Type: "deposit"},
}
```

สร้าง:

```go
type AccountSummary struct {
    UserID   int
    Deposit  float64
    Withdraw float64
    Balance  float64
}
```

แล้ว:

```go
func Summarize(
    transactions []Transaction,
) []AccountSummary
```

**Expected:**

```
User 101
Deposit  = 1300
Withdraw = 500
Balance  = 800

User 102
Deposit  = 2000
Withdraw = 0
Balance  = 2000
```

> 🔥 **อันนี้ผมอยากให้คุณทำมากที่สุด**
>
> เพราะมันรวม: Struct · Slice · Map · Loop · Condition · Aggregation · Function
>
> และ concept คล้ายกับงานจริงมาก

---

## 🔴 Level 5 — Challenge

### 🔴 โจทย์ 11: Order Summary ⭐⭐⭐⭐⭐

สมมติระบบ E-commerce

```go
type Order struct {
    OrderID  int
    UserID   int
    Product  string
    Price    float64
    Quantity int
}
```

ข้อมูล:

```go
orders := []Order{
    {1, 101, "Keyboard", 1500, 2},
    {2, 102, "Mouse", 800, 1},
    {3, 101, "Monitor", 5000, 1},
    {4, 102, "Keyboard", 1500, 3},
}
```

สร้าง:

```go
type UserOrderSummary struct {
    UserID      int
    TotalOrders int
    TotalItems  int
    TotalAmount float64
}
```

เขียน:

```go
func SummarizeOrders(
    orders []Order,
) []UserOrderSummary
```

**Expected:**

```
User 101
TotalOrders = 2
TotalItems  = 3
TotalAmount = 8000

User 102
TotalOrders = 2
TotalItems  = 4
TotalAmount = 5300
```
