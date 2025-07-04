# Nuntius

Nuntius (Latin for "messenger") is a modern, full-stack blogging platform featuring a Go backend and a Svelte + TypeScript frontend. It provides secure user authentication, markdown-based blog management, and a responsive admin dashboard.

---

## Features

- **User Authentication**: Secure signup, login, JWT-based sessions.
- **Blog Management**: Create, edit, delete, and view blogs with markdown support.
- **RESTful API**: Go backend with robust routing and middleware.
- **Admin Dashboard**: Svelte-based UI for managing blogs and users.
- **Markdown Storage**: Blogs are stored as markdown files for easy editing and portability.
- **Modern UI**: Built with Svelte, Tailwind CSS, and Vite for fast, beautiful interfaces.

---

## Monorepo Structure

```
backend/         # Go backend server (API, DB, static serving)
frontend_admin/  # Svelte + TS admin dashboard
blogs_md_folder/ # Markdown files for blogs
```

---

## Technologies Used

- **Backend**: Go, go-chi, GORM, SQLite, JWT
- **Frontend**: Svelte, TypeScript, Vite, Tailwind CSS
- **Package Management**: Go modules, pnpm

---

## Getting Started

### Prerequisites
- Go 1.24+
- Node.js 18+
- pnpm (or npm/yarn)

### Backend Setup
```sh
cd backend
# Run the server
# (Ensure Go is installed)
go run ./cmd/main.go
```

### Frontend Setup
```sh
cd frontend_admin
pnpm install
pnpm dev
```

- The Application runs on [http://localhost:3000](http://localhost:3000)

---

## API Overview
- `POST /api/v1/auth/signup` — Register
- `POST /api/v1/auth/login` — Login
- `GET /api/v1/blogs` — List blogs (auth required)
- `POST /api/v1/blogs` — Create blog (auth required)
- ...and more (see source for full API)

---

## Makefile Usage

The project includes a `MakeFile` to automate common tasks for both Windows and Linux environments.

### Common Targets

- **Frontend**
  - `make frontend-windows` — Build frontend and copy static files (Windows)
  - `make frontend-linux` — Build frontend and copy static files (Linux)
- **Backend**
  - `make backend-windows` — Build Go backend binary (Windows)
  - `make backend-linux` — Build Go backend binary (Linux)
- **Build All**
  - `make build-windows` — Build both frontend and backend (Windows)
  - `make build-linux` — Build both frontend and backend (Linux)
- **Clean**
  - `make clean-windows` — Remove built files and static assets (Windows)
  - `make clean-linux` — Remove built files and static assets (Linux)

> **Note:** On Windows, run these commands in a terminal that supports `make` (e.g., Git Bash, or use `mingw32-make`).

---

## License
MIT © 2025 Nwankwo Onyedika Benson
