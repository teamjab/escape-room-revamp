pipeline {
    agent any
    stages {
        stage('Jest Unit Test') {
            agent { docker { image 'node:14-alpine' } }
            steps {
                sh 'cd "escape-room"'
            }
        }
        stage('Build') {
            steps {
                echo 'Building..'
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