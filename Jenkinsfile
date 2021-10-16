pipeline {
    agent any
    stages {
        stage('setup phase') {
            steps {
                sh '''
                export PATH=/usr/bin
                '''
            }
        }
        stage('Build') {
            steps {
                sh '''
                cd "escape-room"
                /usr/bin/npm install -y
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