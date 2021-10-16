pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh '''
                cd "escape-room"
                sudo apt install npm
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