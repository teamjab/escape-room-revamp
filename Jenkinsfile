pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh '''
                which npm
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