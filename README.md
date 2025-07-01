# Course Management System

A full-stack web application that allows users to explore courses, add them to favourites, search for courses, and generate learning roadmaps using AI. Includes an admin interface to manage courses. The backend is built fully in Go with Gin.

---

## Features

### User Functionality
- **Register and Login** with JWT-based authentication.
- **Browse Courses** on the home page.
- **Search Courses** by keywords.
- **Add to Favourites** and manage your saved courses.
- **Generate Learning Roadmaps** powered by **Gemini 1.5 Flash AI model.**

### Admin Functionality
- Login as an **Admin**.
- **Add new courses** with title, image, and link.
- **Delete existing courses.**

### Authentication & Authorization
- **JWT tokens** for secure authentication.
- Authorization middleware to protect admin routes.

---

## Tech Stack

### Frontend
- React.js
- Chakra UI
- Redux Toolkit
- Axios

### Backend
- Go (Golang)
- Gin Web Framework
- MongoDB with the official Go driver
- JWT for authentication & authorization
- CORS middleware

### AI Integration
- Roadmap generation using **Gemini 1.5 Flash (Google AI model).**

---

## 🗂️ Project Structure

### Backend
```
go_backend/
├── controllers/
│   ├── course_controller.go
│   ├── favourite_controller.go
│   └── user_controller.go
├── db/
│   └── mongo.go
├── middleware/
│   └── jwt.go
├── models/
│   ├── course.go
│   ├── fav.go
│   └── user.go
├── routes/
│   ├── course_routes.go
│   ├── fav_routes.go
│   └── user_routes.go
├── go.mod
├── go.sum
└── main.go

```

### Frontend
```
frontend/
├── public/
├── src/
│   ├── assets/
│   ├── pages/
│   │   ├── favourites.jsx
│   │   ├── home.jsx
│   │   ├── login.jsx
│   │   ├── register.jsx
│   │   ├── roadmap.jsx
│   │   └── navbar.jsx
│   ├── slice/
│   ├── App.jsx
│   ├── main.jsx
│   └── store.js
└── package.json
```

---

## Getting Started

### Backend
1. **Install Go modules:**
   ```bash
   go mod tidy
   ```
2. **Configure `.env`:**
   ```
   MONGO_URI=your_mongodb_connection_string
   JWT_SECRET=your_jwt_secret
   ```
3. **Run the server:**
   ```bash
   go run main.go
   ```

### Frontend
1. **Install dependencies:**
   ```bash
   npm install
   ```
2. **Run the app:**
   ```bash
   npm start
   ```

---

## 📌 Notes
- Admin credentials can be configured in your MongoDB user collection or via a seeding script.
- JWT tokens are required for protected routes (like adding/deleting courses).

---

## 🌟 License
This project is open source and available under the [Fake License](LICENSE).

---
