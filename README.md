# Go LLM Concurrency & Semantic Caching System

A high-performance LLM backend architecture using Go, Redis, Python FastAPI, embeddings, and concurrency deduplication.

## Overview

This project implements a scalable LLM backend that prevents duplicate LLM calls, reduces latency, and improves throughput by combining:

Go concurrency (singleflight)

Redis exact + semantic caching

MiniLM embedding similarity matching

Redis-backed job queue

Python FastAPI LLM microservice

React Router frontend

When multiple users send identical or similar prompts, the system ensures only one LLM request runs — all others reuse cached or shared results.

## Architecture (High-Level)
Client → Go API → Exact Cache → Semantic Cache → singleflight
       → Job Queue → Go Worker → Python LLM Service → Redis Cache


Exact cache: constant-time LLM response lookup

Semantic cache: embeddings + cosine similarity

singleflight: concurrency deduplication

Worker queue: async LLM processing

Python FastAPI: embeddings + model inference

## Tech Stack
Backend

Go (Gin)

Redis

Python FastAPI

Ollama (local LLM runtime)

PostgreSQL 

Frontend

React Router + TypeScript

## Project Structure
```
backend/
  go/
    cmd/
      api/
        main.go
      worker/
        main.go
    internal/
      auth/
      common/
      llmclient/
      users/
      worker/
      llm/
      queue/
  python/
    models/
    chat.py
    embedding.py
    main.py
frontend/
  app/
```

## Installation & Setup
### Install Dependencies
**Go backend**
```
cd backend/go
go mod download
```
**Python backend**
```
cd backend/python
pip install -r requirements.txt
```


(Make sure FastAPI, uvicorn, requests, and sentence-transformers are installed.)

**Frontend**
```
cd frontend
npm install
```

**Start Redis**
```
redis-server
```

**Ensure Ollama model is installed**
```
ollama pull <model-name>
```

##Copy the example env files:
```
cd backend/go
cp .env.sample .env
```
```
cd backend/python
cp .env.sample .env
```
```
cd frontend
cp .env.sample .env
```

## Running the Project

### Run each service in a separate terminal.

**1. Start the Go API Server**
```
cd backend/go
go run cmd/api/main.go
```


Runs on port 8000

**2. Start the Go Worker Service**
```
cd backend/go
go run cmd/worker/main.go
```


This worker pulls LLM tasks from Redis and stores results in cache.

**3. Start the Python FastAPI LLM Backend**
```
cd backend/python
uvicorn main:app --host 0.0.0.0 --port [port]
```


Provides:

/embed → embeddings

/chat → LLM response

**4. Start the React Router Frontend**
```
cd frontend
npm run dev
```

## Key Features
**1. Exact Cache**

Redis resp:<prompt> lookup prevents recomputation.

**2. Semantic Cache**

Embedding-based similarity search using MiniLM.

**3. singleflight Concurrency Deduplication**

Ensures only one goroutine performs LLM work for matching prompts.

**4. Redis Job Queue**

Asynchronous processing pipeline for LLM workloads.

**5. Go Worker**

Consumes tasks → calls Python → updates cache.

**6. Python LLM Runtime**

Handles embeddings + LLM responses through FastAPI.

**7. Performance Benefits**

Avoids redundant LLM calls

Reduces average latency

Saves compute cost

Improves throughput under load

Handles high concurrency reliably

**8. Architecture Diagram**
<img width="889" height="487" alt="Screenshot 2025-11-19 at 10 44 17" src="https://github.com/user-attachments/assets/83982391-ab3a-4c6e-979e-bf83f91e467b" />
