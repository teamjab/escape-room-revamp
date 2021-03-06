pipeline {
    agent any

    environment {
        APP_NAME = "escape-room"
        API_NAME = "api"
        DOCKER_REGISTRY = "teamjab/escape-room"
        DOCKER_REGISTRY_CREDS = 'dockerhub'
        HEROKU_APP_NAME="escape-room-revamp"
        HEROKU_API_NAME="escape-room-revamp-api"
    }

    stages {
        // stage('NPM package install') {
        //     steps {
        //         sh '''
        //         echo "Installing Packages..."
        //         cd "${APP_NAME}"
        //         npm install -y
        //         '''
        //     }
        // }
        // stage('JS Unit Test') {
        //     steps {
        //         sh '''
        //         echo "JS unit test..."
        //         cd "${APP_NAME}"
        //         npm run test
        //         '''
        //     }
        // }
        stage('Go Unit Test') {
            steps {
                sh '''
                echo "Testing backend service..."
                export PATH=$PATH:/usr/local/go/bin
                cd "${API_NAME}"
                go test ./...
                '''
            }
        }
        stage('Build & Deploy FE application') {
            steps {
                withCredentials([string(credentialsId: 'heroku', variable: 'HEROKU_PASS')]) {
                sh '''
                echo "Deploying to https://escape-room-revamp.herokuapp.com"
                docker login --username=_ --password=$HEROKU_PASS registry.heroku.com
                heroku container:push web -a $HEROKU_APP_NAME
                heroku container:release web -a $HEROKU_APP_NAME
                '''
            }
          }
        }

        stage('Build & Deploy BE application') {
            steps {
                withCredentials([string(credentialsId: 'heroku', variable: 'HEROKU_PASS')]) {
                sh '''
                echo "Deploying to https://escape-room-revamp-api.herokuapp.com"
                cd api
                docker login --username=_ --password=$HEROKU_PASS registry.heroku.com
                heroku container:push web -a $HEROKU_API_NAME
                heroku container:release web -a $HEROKU_API_NAME
                '''
            }
          }
        }
    }
}