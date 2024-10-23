pipeline {
    agent any

    environment {
        APP_IMAGE = 'catfact-app'
        MONGO_IMAGE = 'fur'
        NETWORK_NAME = 'catfact-network'
        GOCACHE = '/var/jenkins_home/.cache/go-build'
    }

    stages {
        stage('Setup Cache Directory') {
            steps {
                sh '''
                    mkdir -p /var/jenkins_home/.cache
                    chmod -R 777 /var/jenkins_home/.cache
                '''
            }
        }

        stage('Build') {
            agent {
                docker {
                    image 'golang:1.23'
                    reuseNode true
                }
            }
            steps {
                sh '''
                    echo "Using GOCACHE at: $GOCACHE"
                    go version
                    go mod tidy
                    go build -v
                '''
            }
        }

        stage('Test') {
            agent {
                docker {
                    image 'golang:1.23'
                    reuseNode true
                }
            }
            steps {
                sh '''
                    echo "Running tests..."
                    go test ./... -v
                '''
            }
        }
    }
}

