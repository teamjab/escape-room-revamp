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
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}