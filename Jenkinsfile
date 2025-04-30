@Library('Shared') _

pipeline {
    stages {
        stage('Code') {
            steps {
                script {
                        git()
                }
            }
        }
        stage('Test') {
            steps {
                echo 'Testing....'
                echo 'Done'
            }
        }
        stage('Deploy') {
            steps {
                script {
                    docker()
                }
            }
        }
    }
}
