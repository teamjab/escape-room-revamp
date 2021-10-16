pipeline {
    agent any

    environment {
        APP_NAME = "escape-room"
        API_NAME = "api"
        DOCKER_REGISTRY = "teamjab/escape-room"
        DOCKER_REGISTRY_CREDS = 'dockerhub'
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
        stage('Building image') {
            steps {
                sh '''
                docker build -t escape-room .
                docker tag escape-room $DOCKER_REGISTRY
                '''
            }
        }
    }
}