# AFEScraper

A Go application for scraping product availability and prices from various online stores.

## Features

- Scrapes product information from Fozzy, Auchan, and Epicentrk.
- Outputs availability and price to the console.

## Project Structure

- `main.go` — Entry point, loads environment variables and runs scrapers.
- `sites/` — Contains site-specific scraping logic.
- `internal/types.go` — Shared types for parsing JSON responses.

## Setup

1. **Clone the repository**

2. **Install dependencies**

   ```sh
   go mod tidy
   ```
