# elimu-api 

but in golang (mostly porting the logic I will do in the java springboot and compare the two, so far i'm enjoying golang more who would've guessed)

## To get started
1. Edit the `.env.example` file to `.env` then fill in your google client id and secrets, you can find out how to get these with a google search "google oauth credentials"
2. To run the server `go run cmd/main.go`
3. To run tests `go test ./...`
4. To generate documentation `swag init -g cmd/api/main.go` make sure you have swaggo installed


## Features
### Modules
- [ ] Identity
- [ ] User management (`student || staff`)
- [ ] Curriculum & Projects
- [ ] Documents
- [ ] Communications
- [ ] Audit

## Test coverage
- Run `go test -cover ./...         `

| Package | Coverage | Status |
|---------|----------|--------|
| `internal/handlers` | 31.3% | ✅ |
| `cmd/api` | 0.0% | ⚠️ No tests |
| `docs` | 0.0% | ✅️ Auto-generated |

# Roadmap

## Core Architecture
- REST API with JSON
- Relational Database
- Caching Layer
- Containerized Code Execution
- WebSocket/Server-Sent Events
- Google OAuth 2.0 Only

---

## 1. Identity & Auth Module

### Authentication
- [x] Google OAuth 2.0 Integration
- [ ] Pre-registration Database Validation (check after Google auth)
- [ ] Session-based Authentication
- [ ] Logout with Session Cleanup

### Authorization
- [ ] Role-based Access Control (different endpoints for students, teachers, admins)
- [ ] Current User Info Endpoint (`GET /api/v1/auth/me`)
- [ ] Admin ability to suspend/reactivate accounts

### User Management
- [ ] User Model with Roles and `is_active` flag
- [ ] Pre-registration Table Sync
- [ ] Login History Tracking
- [ ] Admin User Management Endpoints

---

## 2. Academic Core Module

### Course Management
- [ ] Course CRUD Operations
- [ ] Course Prerequisites System
- [ ] Course Search & Filtering
- [ ] Semester/Year Organization

### Enrollment System
- [ ] Student Enrollment/Unenrollment
- [ ] Course Capacity Validation
- [ ] Enrollment Period Control
- [ ] Course Roster Management

### Academic Records
- [ ] Grade Storage System
- [ ] Transcript Generation
- [ ] GPA Calculation
- [ ] Grade Management API

---

## 3. Exam & Assessment Module

### Exam Management
- [ ] Exam CRUD Operations
- [ ] Question Types: Coding, Multiple Choice, Essay
- [ ] Question Randomization
- [ ] Exam Scheduling

### On-Campus Exam Features
- [ ] Exam Session Management

### Code Execution Sandbox
- [ ] Containerized Code Execution
- [ ] Multi-language Support (Python, JavaScript, Java, etc.)
- [ ] Automated Test Case Execution
- [ ] Resource Limiting (CPU, Memory, Time)

### Basic Integrity Checks
- [ ] Code Similarity Analysis
- [ ] Submission Timing Analysis
- [ ] Manual Review Flagging System

---

## 4. User Profile Module

### Profile Management
- [ ] Profile CRUD Endpoints
- [ ] Avatar Management
- [ ] Academic History
- [ ] Student Information Display

### Preferences
- [ ] Notification Settings
- [ ] Privacy Controls
- [ ] Display Preferences Storage

---

## 5. Admin & System Module

### User Administration
- [ ] Bulk User Management (Import/Export)
- [ ] Role Assignment
- [ ] Account Management (Enable/Disable)
- [ ] Audit Logging

### System Management
- [ ] Health Monitoring Endpoints
- [ ] Usage Analytics
- [ ] Error Tracking & Logging
- [ ] Configuration Management

### Exam Administration
- [ ] Exam Monitoring Dashboard
- [ ] Performance Reports
- [ ] Grade Management Tools
- [ ] Invigilator Management System

---

## 6. Platform Infrastructure

### Configuration
- [ ] Environment-based Configuration
- [ ] Feature Flag Management
- [ ] Rate Limit Configuration

### Deployment
- [ ] Container Support
- [ ] CI/CD Pipeline
- [ ] Database Migration System

### Observability
- [ ] Health Check Endpoints
- [ ] System Metrics Collection
- [ ] Structured Logging

### Security
- [ ] API Security Configuration
- [ ] CORS Policy Management
- [ ] Security Headers Enforcement

---

## Development Phases

### Phase 1: Foundation (
- [ ] Project Setup
- [ ] Database Schema Design
- [ ] Google OAuth with Pre-registration Check
- [ ] Basic User and Course Models
- [ ] Authentication System

### Phase 2: Core Academic 
- [ ] Enrollment System
- [ ] Basic Exam System (Multiple Choice Only)
- [ ] Grade Management
- [ ] Profile System
- [ ] Admin User Management

### Phase 3: Exam System  
- [ ] Code Execution Sandbox
- [ ] Advanced Exam Management
- [ ] Basic Integrity Checks
- [ ] Real-time Monitoring Features

### Phase 4: Polish & Scale  
- [ ] Performance Optimization
- [ ] Comprehensive Testing
- [ ] Security Hardening
- [ ] API Documentation

---

## On-Campus Exam Strategy

### Physical Infrastructure Integration
- [ ] Lab/Classroom Reservation System API
- [ ] Exam Session Logging

### Simplified Integrity Measures
- [ ] Submission Time Analysis
- [ ] Basic Code Similarity Checking
- [ ] Pattern Flagging for Manual Review
- [ ] Session-based Exam Tracking

### Invigilator Features
- [ ] Invigilator Dashboard
- [ ] Student Activity Monitoring
- [ ] Incident Reporting System
- [ ] Manual Override Capabilities

---

## Risk Mitigation

### Exam Integrity (Simplified)
- Physical invigilation as primary control
- Basic automated checks as secondary
- Manual review process for flagged cases
- Session-based tracking for accountability

### Code Execution Security
- Container isolation
- Resource limits
- Network restrictions
- Execution timeouts

### System Performance
- Queue system for code execution
- Caching strategy
- Database optimization
- Load handling for exam periods

---
