pipeline {
    agent any

    environment {
        APP_NAME = "escape-room"
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