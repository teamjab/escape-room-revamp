pipeline {
    agent any
    stages {
        stage('setup phase') {
            steps {
                sh '''
                export PATH=/usr/local/bin
                '''
            }
        }
        stage('Build') {
            steps {
                sh '''
                cd "escape-room"
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