pipeline {
    agent any
    
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        
        stage('Build') {
            steps {
                sh 'go build -o myapp'
            }
        }
        
        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }
        
        stage('Run') {
            steps {
                sh './myapp'
            }
        }
    }
}