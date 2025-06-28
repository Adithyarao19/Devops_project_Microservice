pipeline {
    agent any

    environment {
        COMPOSE_FILE = 'docker-compose.yml'
    }

    stages {
        stage('Checkout Code') {
            steps {
                git url: 'https://github.com/Adithyarao19/Adithya_devops_assmt.git', branch: 'main'
            }
        }

        stage('Build and Start Services') {
            steps {
                bat 'docker-compose down || exit 0'
                bat 'docker-compose build'
                bat 'docker-compose up -d'
            }
        }

        stage('Health Check') {
            steps {
                echo "Waiting for services to start..."
                sleep(time: 15, unit: 'SECONDS')

                bat '''
                curl -f http://localhost:8001/ping || exit /b 1
                curl -f http://localhost:8002/ping || exit /b 1
                '''
            }
        }

        stage('Prometheus Check') {
            steps {
                bat '''
                curl -f http://localhost:9090/metrics || exit /b 1
                '''
            }
        }
    }

    post {
        always {
            echo 'Cleaning up containers...'
            bat 'docker-compose down'
        }

        success {
            echo '✅ Pipeline succeeded! All services healthy.'
        }

        failure {
            echo '❌ Pipeline failed. Please check the logs.'
        }
    }
}
