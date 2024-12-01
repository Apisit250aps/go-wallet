# Go-Wallet API

โปรเจค Go-Wallet เป็น REST API สำหรับจัดการรายรับ-รายจ่าย โดยใช้ Go และ MongoDB พัฒนาด้วยแนวคิด Clean Architecture

## เทคโนโลยีที่ใช้
- Go (Golang)
- MongoDB
- JWT สำหรับการยืนยันตัวตน
- Gin Framework
- godotenv สำหรับจัดการ environment variables

# โครงสร้างโปรเจคและความหมาย

```
go-wallet/
├── cmd/                    # โฟลเดอร์สำหรับไฟล์ที่ใช้รันแอปพลิเคชัน
│   └── api/               # โฟลเดอร์สำหรับ API service
│       └── main.go        # จุดเริ่มต้นของแอปพลิเคชัน เป็นที่รวม dependencies ทั้งหมด
├── internal/              # โค้ดที่ใช้เฉพาะภายในโปรเจคนี้เท่านั้น ไม่สามารถ import ไปใช้ที่อื่นได้
│   ├── config/           # การตั้งค่าต่างๆ ของแอปพลิเคชัน เช่น database, server, JWT
│   ├── domain/          # เก็บ entities และ interfaces หลักของระบบ (business rules)
│   ├── usecase/         # เก็บ business logic หรือกฎการทำงานต่างๆ ของระบบ
│   ├── repository/      # ส่วนติดต่อกับฐานข้อมูล จัดการการอ่าน/เขียนข้อมูล
│   └── delivery/        # ส่วนรับส่งข้อมูลกับภายนอก (HTTP, gRPC) รวมถึง middleware
├── pkg/                 # โค้ดที่สามารถนำไปใช้โปรเจคอื่นได้ (reusable packages)
│   ├── auth/            # เครื่องมือสำหรับจัดการ authentication (JWT, password hashing)
│   ├── database/       # utility functions สำหรับเชื่อมต่อฐานข้อมูล
│   └── response/       # รูปแบบ response มาตรฐานที่ส่งกลับไปยัง client
├── .env                # ไฟล์เก็บค่า environment variables สำหรับ development
├── .env.example       # ตัวอย่างไฟล์ .env เพื่อให้ developer อื่นๆ รู้ว่าต้องตั้งค่าอะไรบ้าง
└── README.md          # เอกสารอธิบายโปรเจค วิธีติดตั้ง และการใช้งาน
```

### หลักการแบ่งโฟลเดอร์:
- `cmd/` - เก็บ entry points ของแอปพลิเคชัน แยกตามประเภทของ service
- `internal/` - โค้ดที่ใช้เฉพาะภายในโปรเจคนี้ Go จะป้องกันไม่ให้โปรเจคอื่นนำไป import
- `pkg/` - แพ็คเกจที่สามารถแชร์ไปใช้ในโปรเจคอื่นได้ มักเป็น utilities ต่างๆ
- `.env` files - ไฟล์สำหรับตั้งค่าที่แตกต่างกันในแต่ละสภาพแวดล้อม (development, production)

### Clean Architecture:
โครงสร้างนี้ออกแบบตามหลัก Clean Architecture โดย:
1. `domain` - เป็นชั้นในสุด ไม่ขึ้นกับชั้นอื่นๆ
2. `usecase` - ใช้ domain และมี business logic
3. `repository` - จัดการข้อมูลตาม interface ที่กำหนดใน domain
4. `delivery` - ชั้นนอกสุด จัดการการรับส่งข้อมูลกับภายนอก
## การติดตั้ง

1. Clone โปรเจค:
```bash
git clone https://github.com/yourusername/go-wallet.git
cd go-wallet
```

2. ติดตั้ง dependencies:
```bash
go mod download
```

3. สร้างไฟล์ .env:
```bash
cp .env.example .env
```

4. แก้ไขไฟล์ .env ตามความเหมาะสม:
```env
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=go_wallet
JWT_SECRET=your-secret-key
JWT_EXPIRY_HOURS=24
SERVER_PORT=:8080
GIN_MODE=release
```

5. รันแอปพลิเคชัน:
```bash
go run cmd/api/main.go
```

## API Endpoints

### การจัดการผู้ใช้
- `POST /api/register` - ลงทะเบียนผู้ใช้ใหม่
- `POST /api/login` - เข้าสู่ระบบ

### การจัดการธุรกรรม (ต้องยืนยันตัวตน)
- `POST /api/transactions` - สร้างธุรกรรมใหม่
- `GET /api/transactions` - ดูรายการธุรกรรมทั้งหมด
- `GET /api/transactions/:id` - ดูรายละเอียดธุรกรรม
- `PUT /api/transactions/:id` - แก้ไขธุรกรรม
- `DELETE /api/transactions/:id` - ลบธุรกรรม

## รูปแบบข้อมูล

### การลงทะเบียน
```json
{
  "username": "testuser",
  "password": "password123",
  "email": "test@example.com"
}
```

### การเข้าสู่ระบบ
```json
{
  "username": "testuser",
  "password": "password123"
}
```

### การสร้าง/แก้ไขธุรกรรม
```json
{
  "amount": 100.50,
  "type": "income",     // "income" หรือ "expense"
  "category": "salary", // หมวดหมู่ธุรกรรม
  "description": "เงินเดือน"
}
```

## การรักษาความปลอดภัย
- รหัสผ่านถูกเข้ารหัสด้วย bcrypt
- ใช้ JWT สำหรับการยืนยันตัวตน
- ต้องส่ง token ในรูปแบบ Bearer token ในส่วน Authorization header

## การพัฒนาเพิ่มเติม
1. เพิ่มระบบรายงานสรุป
2. เพิ่มการจัดการหมวดหมู่
3. เพิ่มระบบแจ้งเตือน
4. เพิ่ม Unit Tests

## หมายเหตุ
- ควรเปลี่ยน JWT_SECRET ในไฟล์ .env เมื่อใช้งานจริง
- ใน production ควรตั้งค่า GIN_MODE=release
- ควรกำหนดค่า trusted proxies ที่เหมาะสมในสภาพแวดล้อมจริง

## ผู้พัฒนา
[Apisit Saithong]()

[aps.apisit250@gmail.com]()