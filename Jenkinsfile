pipeline {
    agent any

    environment {
        APP_IMAGE = 'catfact-app'
        MONGO_IMAGE = 'fur'
        NETWORK_NAME = 'catfact-network'
    }

    stages {
        

        // stage('Run Tests') {
        //     steps {
        //         // Run tests in the Go application
        //         sh 'go test ./... -v'
        //     }
        // }

        stage('Build Docker Image') {
            steps {
                // Build the Docker image for the Go application
                script {
                    dockerImage = docker.build(APP_IMAGE)
                }
            }
        }

        stage('Setup MongoDB Container') {
            steps {
                // Create a Docker network if it doesn't exist
                sh '''
                docker network create ${NETWORK_NAME} || true
                docker run -d --name ${MONGO_IMAGE} --network ${NETWORK_NAME} mongo:latest
                '''
            }
        }

        stage('Run Go App Container') {
            steps {
                // Run the Go application container
                script {
                    dockerImage.run("--network ${NETWORK_NAME} -p 3000:3000")
                }
            }
        }
    }

    post {
        always {
            // Cleanup: Stop and remove the MongoDB container after the build
            sh '''
            docker stop ${MONGO_IMAGE} || true
            docker rm ${MONGO_IMAGE} || true
            '''
        }
    }
}
