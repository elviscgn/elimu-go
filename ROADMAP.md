
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
- [x] Pre-registration Database Validation (check after Google auth)
- [x] Session-based Authentication
- [x] Logout with Session Cleanup

### Authorization
- [x] Role-based Access Control (different endpoints for students, teachers, admins)
- [x] Current User Info Endpoint (`GET /api/me`)


### User Management
- [x] User Model with Roles
- [x] Pre-registration Table Sync

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
- [ ] Audit Logging (logins, everything)


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
