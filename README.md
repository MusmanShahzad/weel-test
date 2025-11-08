# Weel Test - Full Stack Application

A modern full-stack application built with Go (Gin) backend and Next.js frontend, featuring authentication, order management, and AI-powered product suggestions.

## 🏗️ Architecture

- **Backend**: Go 1.23+ with Gin framework, GORM, PostgreSQL
- **Frontend**: Next.js 16, React 19, Redux Toolkit, Redux Saga, React Query
- **Database**: PostgreSQL 15
- **Authentication**: JWT-based authentication
- **AI Integration**: OpenAI API for product suggestions

## 📁 Project Structure

```
Weel-test/
├── backend/              # Go backend application
│   ├── cmd/             # Application entry points
│   ├── config/          # Configuration management
│   ├── internal/        # Internal packages
│   │   ├── app/        # Application initialization
│   │   ├── domain/     # Domain models
│   │   ├── handler/    # HTTP handlers
│   │   ├── middleware/ # HTTP middleware
│   │   ├── module/     # Feature modules
│   │   ├── repository/ # Data access layer
│   │   ├── router/     # Routing
│   │   ├── seed/       # Database seeders
│   │   └── service/    # Business logic
│   ├── migrations/     # Database migrations
│   ├── Dockerfile      # Backend Docker image
│   └── docker-compose.yml  # Backend-only Docker setup
├── frontend/           # Next.js frontend application
│   └── weel-test/
│       ├── app/        # Next.js app directory
│       ├── components/ # React components (Atomic Design)
│       ├── lib/         # Utilities and configurations
│       └── Dockerfile    # Frontend Docker image
├── docker-compose.yml   # Root Docker Compose (all services)
└── README.md           # This file
```

## 🚀 Quick Start with Docker

### Prerequisites

- Docker and Docker Compose installed
- Git

### Steps

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd Weel-test
   ```

2. **Create environment file**
   ```bash
   cp .env.example .env
   ```

3. **Edit `.env` file** (optional, defaults work for development)
   ```env
   # Database
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=weel_db
   DB_PORT=5432

   # Backend
   BACKEND_PORT=8080
   JWT_SECRET=your-super-secret-jwt-key-change-in-production

   # Frontend
   FRONTEND_PORT=3000
   NEXT_PUBLIC_API_URL=http://localhost:8080

   # OpenAI (Optional)
   OPEN_AI_SECRET=your-openai-api-key-here
   ```

4. **Start all services**
   ```bash
   docker-compose up -d
   ```

   This will:
   - Start PostgreSQL database
   - Build and start the backend (runs migrations and seeds automatically)
   - Build and start the frontend

5. **Access the application**
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080
   - Health Check: http://localhost:8080/health

6. **View logs**
   ```bash
   docker-compose logs -f
   ```

7. **Stop all services**
   ```bash
   docker-compose down
   ```

8. **Stop and remove volumes** (clean slate)
   ```bash
   docker-compose down -v
   ```

## 🛠️ Local Development Setup

### Backend Setup

1. **Prerequisites**
   - Go 1.23 or higher
   - PostgreSQL 15 (or use Docker for database only)

2. **Install dependencies**
   ```bash
   cd backend
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cd backend
   cp .env.example .env
   ```
   Edit `.env` with your database credentials.

4. **Start PostgreSQL** (if using Docker)
   ```bash
   cd backend
   docker-compose up -d postgres
   ```

5. **Run migrations**
   ```bash
   cd backend
   make migrate
   # or
   go run ./cmd/migrate
   ```

6. **Seed database**
   ```bash
   cd backend
   make seed
   # or
   go run ./cmd/seed
   ```

7. **Run the server**
   ```bash
   cd backend
   make run
   # or
   go run ./cmd/server
   ```

   Server will start at `http://localhost:8080`

### Frontend Setup

1. **Prerequisites**
   - Node.js 20 or higher
   - npm or yarn

2. **Install dependencies**
   ```bash
   cd frontend/weel-test
   npm install
   ```

3. **Set up environment variables**
   ```bash
   cd frontend/weel-test
   cp .env.example .env.local
   ```
   Edit `.env.local`:
   ```env
   NEXT_PUBLIC_API_URL=http://localhost:8080
   ```

4. **Run development server**
   ```bash
   cd frontend/weel-test
   npm run dev
   ```

   Frontend will start at `http://localhost:3000`

## 📝 Default Credentials

After seeding the database, you can use these credentials:

- **Admin User**
  - Email: `admin@example.com`
  - Password: `password123`

- **Test User**
  - Email: `user@example.com`
  - Password: `password123`

## 🔌 API Endpoints

### Authentication
- `POST /api/v1/auth/login` - Login
- `GET /api/v1/me` - Get current user (protected)

### Orders
- `GET /api/v1/orders` - List orders (protected, with filters)
- `POST /api/v1/orders` - Create order (protected)
- `GET /api/v1/orders/:id` - Get order by ID (protected)
- `PUT /api/v1/orders/:id` - Update order (protected)
- `POST /api/v1/orders/suggestions` - Get AI product suggestions (public)

### Feature Flags
- `GET /api/v1/feature-flags` - Get all feature flags (protected)

### Health Check
- `GET /health` - Health check endpoint

## 🧪 Testing

### Backend Tests
```bash
cd backend
make test
# or
go test ./...
```

### Test Coverage
```bash
cd backend
make test-coverage
```

## 🐳 Docker Commands

### Start services
```bash
docker-compose up -d
```

### Stop services
```bash
docker-compose down
```

### View logs
```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f backend
docker-compose logs -f frontend
docker-compose logs -f postgres
```

### Rebuild services
```bash
docker-compose up -d --build
```

### Execute commands in containers
```bash
# Backend shell
docker-compose exec backend sh

# Database shell
docker-compose exec postgres psql -U postgres -d weel_db

# Run migrations manually
docker-compose exec backend ./migrate

# Run seed manually
docker-compose exec backend ./seed
```

## 🔧 Environment Variables

### Root `.env` (for Docker Compose)
```env
# Database
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=weel_db
DB_SSLMODE=disable

# Backend
BACKEND_PORT=8080
SERVER_PORT=8080
SERVER_HOST=0.0.0.0
JWT_SECRET=your-super-secret-jwt-key-change-in-production
OPEN_AI_SECRET=your-openai-api-key-here

# Frontend
FRONTEND_PORT=3000
NEXT_PUBLIC_API_URL=http://localhost:8080
```

### Backend `.env` (for local development)
```env
SERVER_PORT=8080
SERVER_HOST=localhost
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=weel_db
DB_SSLMODE=disable
JWT_SECRET=your-super-secret-jwt-key-change-in-production
OPEN_AI_SECRET=your-openai-api-key-here
```

### Frontend `.env.local` (for local development)
```env
NEXT_PUBLIC_API_URL=http://localhost:8080
```

## 📦 Database Migrations

The backend uses GORM's AutoMigrate for code-first migrations. Models are defined in `backend/internal/domain/` and migrations run automatically on server start.

### Manual Migration
```bash
# Local
cd backend
make migrate

# Docker
docker-compose exec backend ./migrate
```

## 🌱 Database Seeding

Seed data includes:
- Admin and test users
- Sample orders
- Feature flags

### Run Seed
```bash
# Local
cd backend
make seed

# Docker (runs automatically on startup)
docker-compose exec backend ./seed

# Reset and reseed
cd backend
make seed-reset
```

## 🏃 Make Commands (Backend)

```bash
cd backend
make help          # Show all available commands
make build         # Build binaries
make run           # Run server
make test          # Run tests
make migrate       # Run migrations
make seed          # Seed database
make seed-reset    # Reset and reseed
make docker-up     # Start Docker (backend only)
make docker-down   # Stop Docker
make docker-logs   # View Docker logs
```

## 🎨 Frontend Features

- **Authentication**: Login/Signup with JWT
- **Dashboard**: Order management with statistics
- **Order Creation**: AI-powered product suggestions
- **Order Management**: View, edit, filter, and sort orders
- **Feature Flags**: Dynamic feature toggling
- **Responsive Design**: Mobile-friendly UI

## 🏛️ Architecture Patterns

### Backend
- **Clean Architecture**: Domain, Repository, Service, Handler layers
- **Module System**: Feature-based modular architecture
- **Dependency Injection**: Container-based DI
- **Code-First Migrations**: GORM AutoMigrate

### Frontend
- **Atomic Design**: Component structure (atoms, molecules, organisms)
- **State Management**: Redux Toolkit + Redux Saga
- **API Management**: React Query + Axios
- **Form Handling**: Formik + Yup validation
- **UI Components**: Shadcn UI

## 🐛 Troubleshooting

### Backend won't start
- Check database connection in `.env`
- Ensure PostgreSQL is running
- Check port 8080 is not in use

### Frontend can't connect to backend
- Verify `NEXT_PUBLIC_API_URL` is set correctly
- Check backend is running on port 8080
- Check CORS settings in backend

### Database connection errors
- Verify PostgreSQL is running
- Check database credentials in `.env`
- Ensure database exists

### Docker issues
- Check Docker is running
- Try rebuilding: `docker-compose up -d --build`
- Check logs: `docker-compose logs -f`

## 👥 Authors

[Usman shahzad]

