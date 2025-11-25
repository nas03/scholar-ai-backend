# Scholar AI Backend - Development Roadmap

## 📋 Overview
Golang backend using Gin + GORM for the Scholar AI application. This roadmap organizes work into clear milestones with priorities and acceptance criteria.

## 🎯 Legend
- **Priority**: 🔴 P0 (urgent) | 🟡 P1 (important) | 🟢 P2 (nice-to-have)
- **Status**: ⬜ Todo | 🟡 In Progress | ✅ Done

## 🏗️ Key Entities
- **Core**: User, Course, Semester, Schedule/Timetable
- **Productivity**: Reminder, Lecture Notes, Materials, Quiz

---

## 🚀 Milestone M0: Foundation & Conventions

### 🔴 P0 - Critical Infrastructure
- [ ] **Standardized Project Configuration**
  - ✅ Load from env and .env files
  - ✅ Required variables validated with clear errors
  - ✅ Sample `.env.example` exists
  - 📁 `internal/config/` (types.go, database.go, server.go, mail.go)

- [ ] **Structured Logging**
  - [x] Central logger (zap/logrus) in `global/global.go`
  - [x] Request logging middleware
  - [x] Log levels via environment
  - [ ] JSON output in production

- [ ] **Error & Response Contract**
  - [x] Use `pkg/response` consistently
  - [ ] Add error codes and request-id correlation
  - [ ] Uniform error payload structure
  - [ ] Update controllers to use helpers

- [ ] **CORS & Security Headers**
  - [x] `internal/middleware/cors.go` with configurable origins
  - [ ] Secure headers (no-sniff, frameguard)
  - [ ] Rate limiting middleware

### 🟡 P1 - Enhanced Security
- [ ] **Rate Limiting**
  - [ ] Global/token-bucket per IP
  - [ ] Per-user limits on auth endpoints

---

## 🔐 Milestone M1: Authentication & User Management

### 🔴 P0 - Core Auth
- [ ] **User Registration & Login**
  - [ ] POST `/api/v1/auth/register`
  - [ ] POST `/api/v1/auth/login`
  - [ ] POST `/api/v1/auth/logout`
  - [ ] Password hashing (bcrypt/argon2id)
  - [ ] JWT access + refresh tokens
  - [ ] Refresh token rotation

- [ ] **Email Verification**
  - [x] Token generation + confirm endpoint
  - [ ] Resend with cooldown
  - [x] Mock provider for development
  - [x] Interface for real provider

- [ ] **Login Security**
  - [ ] Failed attempt tracking
  - [ ] Temporary lockout with exponential backoff
  - [ ] Audit logs

### 🟡 P1 - Advanced Auth
- [ ] **Password Reset Flow**
  - [ ] Request reset endpoint
  - [ ] Token generation and validation
  - [ ] Reset endpoint with token invalidation
  - [ ] Minimum password policy

- [ ] **Phone Verification** (Optional)
  - [ ] Store E.164 format numbers
  - [ ] OTP verification
  - [ ] Pluggable SMS provider
  - [ ] Rate limiting

- [ ] **2FA (TOTP)**
  - [ ] Enable/disable TOTP
  - [ ] QR code provisioning
  - [ ] Recovery codes
  - [ ] Step-up authentication

### 🟢 P2 - SSO Integration
- [ ] **OAuth2 SSO**
  - [ ] Google/Microsoft OIDC
  - [ ] Account linking
  - [ ] New-user onboarding with provider claims

---

## 📚 Milestone M2: Core Domain CRUD

### 🔴 P0 - Essential Features
- [ ] **Courses Management**
  - [ ] CRUD operations (name, description, credits)
  - [ ] Routes in `internal/router/user.route.go`
  - [ ] Service + repository methods
  - [ ] Unit tests

### 🟡 P1 - Academic Structure
- [ ] **Semesters Management**
  - [ ] CRUD operations (name, start/end dates)
  - [ ] Course-semester mapping validation

- [ ] **Schedule/Timetable**
  - [ ] CRUD for time blocks
  - [ ] Day-of-week, start/end times, location
  - [ ] Conflict detection on create/update

---

## 🎯 Milestone M3: Productivity Features

### 🟡 P1 - Core Productivity
- [ ] **Reminders System**
  - [ ] CRUD operations
  - [ ] Schedule engine (cron/worker)
  - [ ] Pluggable channels (email, push)

- [ ] **Lecture Notes**
  - [ ] CRUD operations
  - [ ] Rich text support (markdown/JSON)
  - [ ] Versioning metadata
  - [ ] Basic search by title/tags

- [ ] **Materials Management**
  - [ ] Upload & list functionality
  - [ ] Storage provider interface (local/S3)
  - [ ] Signed URLs for download
  - [ ] Size/type limits

### 🟢 P2 - Advanced Features
- [ ] **Quiz System**
  - [ ] CRUD for quizzes and questions
  - [ ] Assignment to courses/notes
  - [ ] Simple scoring endpoint

---

## 🗄️ Milestone M4: Data & Persistence

### 🔴 P0 - Database Foundation
- [x] **Database Migrations** ✅
  - [x] Atlas migration tool setup
  - [x] GORM integration
  - [x] Makefile targets for up/down
  - 📁 `migrations/` directory

- [ ] **GORM Models Finalized**
  - [x] Models in `internal/models/*.go`
  - [x] Constraints and indexes
  - [ ] AutoMigrate only in development

### 🟡 P1 - Caching & Performance
- [ ] **Redis Integration**
  - [ ] Connection pool from config
  - [ ] Health check endpoint
  - [ ] Rate limiting and token blacklisting
  - [ ] Hot read caching

---

## 📊 Milestone M5: Observability & Reliability

### 🔴 P0 - Health Monitoring
- [ ] **Health Endpoints**
  - [ ] `/healthz` (process health)
  - [ ] `/readyz` (DB + Redis readiness)
  - [ ] `/livez` (liveness check)
  - [ ] Container probe integration

### 🟡 P1 - Metrics & Monitoring
- [ ] **Prometheus Metrics**
  - [ ] Process metrics
  - [ ] HTTP latency metrics
  - [ ] Database latency metrics
  - [ ] Cache hit rate metrics
  - [ ] `/metrics` endpoint

### 🟢 P2 - Advanced Observability
- [ ] **OpenTelemetry Tracing**
  - [ ] HTTP handler traces
  - [ ] Database call traces
  - [ ] Configurable exporter

---

## 🧪 Milestone M6: Quality Assurance

### 🔴 P0 - Testing Foundation
- [ ] **Unit Tests**
  - [ ] >60% coverage for core packages
  - [ ] Table-driven tests
  - [ ] SQLite-in-memory for GORM repos

- [ ] **Integration Tests**
  - [ ] Auth flow: register → verify → login → refresh → logout
  - [ ] CI pipeline integration

### 🟡 P1 - Code Quality
- [ ] **Static Analysis**
  - [ ] `golangci-lint` configuration
  - [ ] CI job fails on lint errors
  - [ ] `go fmt` enforcement

---

## 📖 Milestone M7: Documentation & Developer Experience

### 🟡 P1 - API Documentation
- [ ] **OpenAPI/Swagger**
  - [ ] `swaggo/swag` integration
  - [ ] `/swagger/index.html` in development
  - [ ] CI validation for spec builds

- [ ] **Developer Onboarding**
  - [ ] Updated README with setup instructions
  - [ ] Makefile targets for common actions
  - [ ] Environment setup guide

---

## 🐳 Milestone M8: Packaging & Deployment

### 🟡 P1 - Containerization
- [ ] **Docker Setup**
  - [ ] Multi-stage Dockerfile
  - [ ] Minimal image size
  - [ ] Non-root user
  - [ ] Health check integration

- [ ] **Development Environment**
  - [ ] Docker Compose (app + DB + Redis)
  - [ ] Hot-reload in development

### 🟢 P2 - CI/CD Pipeline
- [ ] **GitHub Actions**
  - [ ] Build, test, lint jobs
  - [ ] Docker build and publish
  - [ ] Artifact management

---

## 💡 Backlog & Future Ideas

### 🟢 P2 - Advanced Features
- [ ] Full-text search for notes/materials
- [ ] Webhook callbacks for reminders
- [ ] Background worker separation
- [ ] Real-time notifications
- [ ] Advanced analytics dashboard

---

## 📁 Current Code Structure
	
```
backend/
├── internal/
│   ├── config/          # Configuration management
│   ├── models/          # GORM models
│   ├── repositories/    # Data access layer
│   ├── services/        # Business logic
│   ├── controllers/     # HTTP handlers
│   ├── router/          # Route definitions
│   └── initialize/      # GORM/router initialization
├── pkg/response/        # Response utilities
├── migrations/          # Database migrations (Atlas)
└── sql/                # Legacy SQL files
```

---

## 📝 Development Notes

- **Framework**: Gin for routing and middleware
- **Database**: GORM with Atlas migrations
- **Security**: JWT tokens, never commit secrets
- **Testing**: Small, composable PRs with tests per milestone
- **Environment**: Use `.env` files and CI secrets

---

## 🎯 Current Status
- ✅ **Database Migration Setup**: Atlas with GORM integration
- ✅ **Project Structure**: Clean architecture with proper separation
- 🟡 **Authentication**: Basic user model ready
- ⬜ **API Endpoints**: Ready for implementation

**Last Updated**: $(date)