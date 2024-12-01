# Go-Wallet API

โปรเจค Go-Wallet เป็น REST API สำหรับจัดการรายรับ-รายจ่าย โดยใช้ Go และ MongoDB พัฒนาด้วยแนวคิด Clean Architecture

## เทคโนโลยีที่ใช้
- Go (Golang)
- MongoDB
- JWT สำหรับการยืนยันตัวตน
- Gin Framework
- godotenv สำหรับจัดการ environment variables

## โครงสร้างโปรเจค
```
go-wallet/
├── cmd/
│   └── api/
│       └── main.go            # จุดเริ่มต้นของแอปพลิเคชัน
├── internal/
│   ├── config/               # การตั้งค่าต่างๆ
│   ├── domain/              # โมเดลและ interface
│   ├── usecase/             # business logic
│   ├── repository/          # การจัดการข้อมูลกับ database
│   └── delivery/            # HTTP handlers และ middleware
├── pkg/
│   ├── auth/                # เครื่องมือสำหรับ JWT และ password
│   ├── database/           # การเชื่อมต่อ database
│   └── response/           # รูปแบบ response มาตรฐาน
├── .env                    # ไฟล์ตั้งค่าสภาพแวดล้อม
├── .env.example           # ตัวอย่างไฟล์ตั้งค่า
└── README.md
```

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