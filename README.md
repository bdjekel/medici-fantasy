TODO: Refine AI-Generated README below.

# medici-fantasy

> Below are the proposed features for this package. The package is in the early stages of construction and has no functioning parts yet. This readme will be updated as the package progresses.

# Fantasy Football League Analytics

A custom fantasy football league management application that provides historical data and analytics not available on ESPN's platform. Built specifically for league commissioners and managers who want deeper insights into their league's history and performance.

## 🏆 Project Overview

This application focuses on tracking and displaying historical fantasy football data, with emphasis on:

- **Manager Career Statistics**: Championships won, win/loss records, career performance
- **Historical Records**: Single matchup high scores, season point totals, all-time achievements
- **Advanced Analytics**: Head-to-head records, consistency metrics, playoff success rates
- **League Management**: Commissioner tools for data management and validation

## 🎯 Key Features (Planned)

### Core Features

- Manager profiles with career statistics and achievements
- Historical records tracking (single matchup highs, season totals)
- Championship history and playoff performance analysis
- Head-to-head manager comparisons
- Season-by-season performance trends

### Commissioner Tools

- Admin CLI for data management and imports
- Historical data import from ESPN Fantasy API
- Manual data entry and validation tools
- Automated daily updates from ESPN

### Advanced Analytics

- Manager efficiency ratings and consistency scores
- Multi-season trend analysis
- Draft and waiver wire success tracking
- Playoff performance metrics

## 🛠️ Technology Stack

### Backend

- **Language**: Go 1.21+
- **Database**: PostgreSQL with pgx driver
- **API**: REST API with JSON responses
- **Code Generation**: sqlc for type-safe database queries
- **Scheduling**: robfig/cron for automated data fetching
- **Logging**: zap for structured logging

### Frontend

- **Framework**: React with Vite
- **Styling**: Tailwind CSS
- **Deployment**: Vercel

### Infrastructure

- **Backend Deployment**: Fly.io (planned)
- **Database**: Managed PostgreSQL
- **Frontend**: Vercel CDN
- **Monitoring**: Basic health checks and logging

## 📁 Project Structure

```
/
├── cmd/
│   ├── api/          # HTTP server entry point
│   └── admin/        # CLI for manual data management
├── internal/
│   ├── service/      # Business logic layer
│   ├── handler/      # HTTP handlers
│   ├── repository/   # Database access layer
│   └── config/       # Configuration management
├── fetcher/          # ESPN Fantasy API integration
├── pkg/
│   └── models/       # Shared data models
├── web/              # React frontend application
├── migrations/       # Database migration files
├── scripts/          # Utility scripts
└── docs/             # Documentation
```

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 14+
- Node.js 18+ (for frontend)
- Docker (for local development)

### Local Development Setup

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd fantasy-football-analytics
   ```

2. **Set up environment variables**

   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Start local database**

   ```bash
   docker-compose up -d postgres
   ```

4. **Run database migrations**

   ```bash
   go run cmd/admin/main.go migrate up
   ```

5. **Start the backend server**

   ```bash
   go run cmd/api/main.go
   ```

6. **Start the frontend development server**
   ```bash
   cd web
   npm install
   npm run dev
   ```

## 📊 Data Sources

- **ESPN Fantasy API**: Primary source for league data, matchups, and scores
- **Manual Entry**: Commissioner tools for historical data import
- **CSV Import**: Bulk import capabilities for legacy data

## 🔧 Development Status

This project is currently in early development. See the [development roadmap](docs/ROADMAP.md) for detailed progress tracking.

### Current Phase: Foundation Setup

- [x] Project structure and repository setup
- [x] Go module initialization
- [x] Basic README documentation
- [ ] Database schema design
- [ ] ESPN Fantasy API integration research
- [ ] Core data models definition

## 🤝 Contributing

This is a private league management tool. If you're interested in similar functionality for your league, feel free to fork and adapt for your needs.

## 📝 License

This project is for personal use within our fantasy football league.

## 🔮 Future Enhancements

- Mobile-responsive design optimization
- Advanced data visualization with charts
- Player-level historical analysis
- Draft performance tracking
- Trade analysis and success metrics
- League comparison tools

---

**Note**: This application is designed specifically for fantasy football league management and historical data analysis. It complements but does not replace ESPN's fantasy platform.
