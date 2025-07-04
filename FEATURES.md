# Tính năng của dự án Project2 Microservice Go

## Tổng quan kiến trúc
Dự án này sử dụng kiến trúc microservice với 5 service chính:
- **Auth Service** (Port 3001) - Xác thực và ủy quyền
- **User Service** (Port 3000) - Quản lý người dùng
- **Task Service** (Port 3002) - Quản lý công việc
- **Notification Service** (Port 3003) - Hệ thống thông báo
- **Dashboard Service** (Port 3004) - Giám sát và điều khiển service

## 1. 🔐 Hệ thống Authentication & Authorization

### Tính năng chính:
- **JWT Token Authentication** với Access Token và Refresh Token
- **Middleware bảo mật** cho tất cả API endpoints cần xác thực
- **Password hashing** với bcrypt
- **Session management** với thời gian hết hạn có thể cấu hình

### API Endpoints:
```
POST /api/auth/login           - Đăng nhập
POST /api/auth/register        - Đăng ký tài khoản
POST /api/auth/refresh-token   - Làm mới token
POST /api/auth/change-password - Đổi mật khẩu (yêu cầu auth)
```

### Cấu trúc JWT Claims:
- UserID, Username, Email
- Thời gian hết hạn có thể cấu hình
- Refresh token kéo dài 7 ngày

## 2. 👥 Hệ thống User Management

### Tính năng chính:
- **CRUD operations** cho quản lý người dùng
- **Profile management** với thông tin cá nhân
- **Password management** với validation
- **User listing** và tìm kiếm

### API Endpoints:
```
GET    /api/user          - Lấy danh sách tất cả users
POST   /api/user          - Tạo user mới
GET    /api/user/me       - Lấy thông tin user hiện tại (yêu cầu auth)
GET    /api/user/:id      - Lấy thông tin user theo ID
PATCH  /api/user/:id      - Cập nhật thông tin user
DELETE /api/user/:id      - Xóa user
POST   /api/user/change-password/:id - Đổi mật khẩu
```

### Data Model:
- ID, Username (unique), Email (unique)
- Password (được hash, không trả về trong response)
- Timestamps (CreatedAt, UpdatedAt)

## 3. 📋 Hệ thống Task Management

### Tính năng chính:
- **CRUD operations** cho tasks
- **Status tracking** (To Do, In Progress, Done)
- **User assignment** và ownership
- **Date management** (start date, due date)
- **Filtering** theo status và user
- **Real-time notifications** khi tạo/cập nhật task

### API Endpoints:
```
GET    /api/task                - Lấy tất cả tasks (có thể filter theo status)
POST   /api/task                - Tạo task mới (yêu cầu auth)
GET    /api/task/:id            - Lấy task theo ID (yêu cầu auth)
PUT    /api/task/:id            - Cập nhật task (yêu cầu auth)
DELETE /api/task/:id            - Xóa task (yêu cầu auth)
GET    /api/task/user/:id       - Lấy tasks theo User ID (yêu cầu auth)
```

### Data Model:
- ID, UserID (foreign key), Title, Content
- StartDate, DueDate, Status
- Timestamps (CreatedAt, UpdatedAt)

### Tích hợp với Notification:
- Tự động gửi notification email khi tạo task mới
- Integration với RabbitMQ để xử lý async

## 4. 📧 Hệ thống Notification

### Tính năng chính:
- **Email notification system** với SMTP
- **RabbitMQ integration** cho message queuing
- **Notification history** và tracking
- **Multiple notification types** (task, auth, user events)
- **Status management** (pending, sent, failed)

### API Endpoints:
```
POST   /api/notifications           - Tạo notification mới (yêu cầu auth)
GET    /api/notifications           - Lấy tất cả notifications (yêu cầu auth)
GET    /api/notifications/:id       - Lấy notification theo ID (yêu cầu auth)
GET    /api/notifications/user/:user_id - Lấy notifications theo User (yêu cầu auth)
PUT    /api/notifications/:id/status - Cập nhật status notification (yêu cầu auth)
DELETE /api/notifications/:id       - Xóa notification (yêu cầu auth)
POST   /api/sendemail               - Gửi email trực tiếp (yêu cầu auth)
```

### Message Queue System:
- **RabbitMQ Exchange**: notification_exchange
- **Routing Keys**: {type}.{action} (vd: task.create, user.update)
- **Consumer** tự động xử lý và gửi email
- **Publisher** từ các service khác

### Email Integration:
- SMTP configuration với Ethereal Email (test)
- Template-based email content
- Error handling và retry logic

## 5. 🖥️ Dashboard Service - Service Monitoring

### Tính năng chính:
- **Docker Container Management** - Start/Stop/Restart services
- **Health Check** cho tất cả microservices
- **Service Status Monitoring** real-time
- **Container Orchestration** thông qua Docker API

### API Endpoints:
```
GET    /api/services/status        - Lấy status tất cả services
POST   /api/services/:name/restart - Restart service
POST   /api/services/:name/stop    - Stop service  
POST   /api/services/:name/start   - Start service
GET    /api/health                 - Health check dashboard
```

### Monitoring Features:
- Docker container status (running/stopped)
- Service health endpoints check
- Port connectivity verification
- Last checked timestamps
- Error logging và reporting

## 6. 🗄️ Database Management

### Tính năng chính:
- **PostgreSQL database** với GORM ORM
- **Auto-migration system** cho database schema
- **Connection pooling** và optimization
- **Transaction support**

### Migration System:
```bash
go run cmd/migrate/main.go --migrate
```

### Database Models:
- **Users**: Quản lý tài khoản người dùng
- **Tasks**: Quản lý công việc và assignment
- **Notifications**: Log và tracking thông báo
- **RefreshTokens**: Quản lý session và security

## 7. 🐳 Infrastructure & DevOps

### Docker Containerization:
- **Multi-container setup** với Docker Compose
- **Service isolation** và scalability
- **Health checks** cho tất cả services
- **Volume persistence** cho database
- **Network isolation** với custom bridge

### Load Balancing:
- **Nginx reverse proxy** (Port 3005)
- **Service discovery** configuration
- **CORS handling** cho frontend integration

### Message Queue:
- **RabbitMQ** với management UI (Port 15672)
- **Exchange/Queue setup** tự động
- **Connection retry logic**

### Monitoring Tools:
- **Service health endpoints** (/api/health)
- **Docker stats integration**
- **Real-time status dashboard**

## 8. 🛠️ Development Tools

### Build và Development:
- **Makefile** với common commands
- **Air live reloading** cho development
- **Environment configuration** với .env
- **Go modules** dependency management

### API Documentation:
- RESTful API design
- Consistent error handling
- JSON request/response format
- Authentication middleware integration

## 9. 🔒 Security Features

### Bảo mật:
- **JWT token-based authentication**
- **Password hashing** với bcrypt
- **CORS configuration** cho frontend
- **Environment variables** cho sensitive data
- **Request validation** và sanitization

### Middleware:
- **Authentication middleware** cho protected routes
- **Error handling middleware**
- **CORS middleware** configuration

## 10. 📊 Logging và Error Handling

### Logging:
- **Structured logging** throughout services
- **Error tracking** và reporting
- **Request/Response logging**
- **Service interaction logs**

### Error Handling:
- **Custom error types** (NotFoundError, ValidationError)
- **Consistent error responses**
- **Graceful degradation**
- **Rollback mechanisms**

## Kết luận

Dự án này cung cấp một hệ thống microservice hoàn chỉnh với:
- ✅ Authentication & Authorization
- ✅ User Management  
- ✅ Task Management với notifications
- ✅ Email notification system
- ✅ Service monitoring dashboard
- ✅ Database migration
- ✅ Docker containerization
- ✅ Message queue integration
- ✅ Load balancing
- ✅ Security best practices

Đây là một foundation mạnh mẽ có thể mở rộng để xây dựng các ứng dụng enterprise-level với Go microservices.