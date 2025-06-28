pipeline {
    agent any

    environment {
        COMPOSE_FILE = 'docker-compose.yml'
    }

    stages {
        stage('Clone Repo') {
    steps {
        git url: 'https://github.com/Adithyarao19/Adithya_devops_assmt.git',
            branch: 'main' 
    }
}


        stage('Build and Start Services') {
            steps {
                echo "Starting Docker Compose..."
                sh 'docker-compose down || true'
                sh 'docker-compose build'
                sh 'docker-compose up -d'
            }
        }

        stage('Health Check') {
            steps {
                echo "Waiting for containers to be healthy..."
                sh 'sleep 20'
                sh 'docker ps --filter "health=healthy"'
            }
        }

        stage('Prometheus Check') {
            steps {
                echo "Scraping Prometheus target..."
                sh 'curl -s http://localhost:9090/metrics | head -n 10'
            }
        }
    }

    post {
        always {
            echo "Cleaning up..."
            sh 'docker-compose down'
        }
    }
}
