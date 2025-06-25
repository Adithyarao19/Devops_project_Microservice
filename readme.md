

# 🚀 DevOps Observability Assignment — by Adithya A.

Welcome to my mini DevOps project that ties together reverse proxy, container orchestration, and observability in a clean, modular setup. This project runs two microservices behind an NGINX reverse proxy and uses Prometheus + Grafana for monitoring.
![Screenshot 2025-06-25 161634](https://github.com/user-attachments/assets/75998b3b-04cd-4823-a97b-8705c396d88c)
this is the custom dashboard i have created 
which shows service1 total requests,service 2 latancy,service 2 total requests.

---

## 📁 Project Overview

This project is built with:

- 🐳 Docker Compose for orchestrating all services  
- 🔀 NGINX as a reverse proxy  
- ⚙️ Two microservices:
  - **Service 1:** Written in Go
  - **Service 2:** Built with Python (Flask)
- 📊 Prometheus for scraping metrics
- 📈 Grafana to visualize them

---

## 🔧 Folder Structure

```

Adithya\_devops\_assmt/
├── docker-compose.yml
├── nginx/
│   ├── nginx.conf
│   └── Dockerfile
├── service\_1/
│   ├── main.go
│   └── Dockerfile
├── service\_2/
│   ├── app.py
│   ├── requirements.txt
│   └── Dockerfile
├── prometheus/
│   └── prometheus.yml
└── README.md

````

---

## 🛠️ How to Run

1. Clone the repository:
  

2. Launch everything:

   ```bash
   docker-compose up --build
   ```

3. Access the services:

| Service    | URL                                                                        |
| ---------- | -------------------------------------------------------------------------- |
| Service 1  | [http://localhost:8080/service1/ping](http://localhost:8080/service1/ping) |
| Service 2  | [http://localhost:8080/service2/ping](http://localhost:8080/service2/ping) |
| Prometheus | [http://localhost:9090](http://localhost:9090)                             |
| Grafana    | [http://localhost:4000](http://localhost:4000)                             |

🧠 **Default Grafana login:**
`Username: admin`
`Password: admin`

---

## 🔍 Observability Details

### ✔️ Prometheus

* Automatically scrapes metrics from both services at `/metrics`.
* Defined scrape configs in `prometheus/prometheus.yml`.

### 📊 Grafana Dashboards

* Prometheus is already set up as a data source.
* You can create panels or import pre-built dashboards.
* Example Dashboard ID for Prometheus stats: **1860**

### Example Metrics Exposed:

* `service1_http_requests_total`
* `flask_http_request_duration_seconds`
* System metrics like CPU, memory, GC are scraped but here it is not necessary 





## ✅ Health Checks

Each service includes Docker health checks that verify `/ping` endpoints regularly to ensure they're alive before proceeding with reverse proxy setup.

---

## 🔐 Security Highlights

* Basic NGINX headers to prevent clickjacking and XSS
* Services are isolated via a custom Docker network
* Lightweight base images (e.g., `alpine`, `python:slim`)

---

## 💡 Why this project?

I built this as part of a DevOps learning path to demonstrate my understanding of containerized environments, reverse proxying, service monitoring, and real-time visualization.




🧑‍💻 Made with care — Adithya A.

