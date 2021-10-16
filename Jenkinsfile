pipeline {
    agent any

    environment {
        APP_NAME: "escape-room"
    }

    stages {
        stage('Build') {
            steps {
                sh '''
                cd "${APP_NAME}"
                npm install -y
                echo "Done Installing packages"
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