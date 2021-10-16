pipeline {
    agent { docker { image 'node:14-alpine' } }
    stages {
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
                echo 'Deploying.... '
            }
        }
    }
}