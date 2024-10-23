pipeline {
    agent any

    stages {

        stage('Build') {
            agent {
                docker {
                    image 'golang:1.23'
                    reuseNode true
                }
            }
            steps {
                sh '''
                    export GOCACHE=/tmp/.cache # Set cache to a writable location
                    mkdir -p /tmp/.cache && chmod -R 777 /tmp/.cache
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
                    echo "Running Tests..."
                    go test ./...
                '''
            }
        }
    }
}
