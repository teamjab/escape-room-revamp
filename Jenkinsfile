pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh '''
                cd "escape-room"
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