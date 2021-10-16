pipeline {
    agent any

    environment {
        APP_NAME = "escape-room"
        API_NAME = "api"
    }

    stages {
        stage('npm package install') {
            steps {
                sh '''
                echo "Installing Packages..."
                cd "${APP_NAME}"
                npm install -y
                '''
            }
        }
        stage('JS Unit Test') {
            steps {
                sh '''
                echo "JS unit test..."
                cd "${APP_NAME}"
                npm run test
                '''
            }
        }
        stage('Go Unit Test') {
            steps {
                sh '''
                export PATH=$PATH:/usr/local/go/bin
                cd "${API_NAME}"
                go test ./...
                '''
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying...'
            }
        }
    }
}