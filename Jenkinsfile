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
                    export GOCACHE=/tmp/.cache
                    mkdir -p $GOCACHE
                    chmod -R 777 $GOCACHE
                    echo "Running Tests..."
                    go test ./...
                '''
            }
        }
    }
}
