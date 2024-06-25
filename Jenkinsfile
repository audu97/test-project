pipeline {
    agent any

    environment {
        DOCKERHUB_CREDENTIALS = credentials('dockerhub')
        DOCKER_IMAGE = 'ephraimaudu/test-app'
        GITHUB_CREDENTIALS = 'git-secret'
    }

    stages{
        stage('Checkout'){
            steps{
                git url: 'https://github.com/audu97/test-project', branch: 'master',
                credentialsId: "${GITHUB_CREDENTIALS}"
            }
        }
        stage('Build'){
            steps{
                script{
                    docker.build("${DOCKER_IMAGE}:${env.BUILD_ID}")
                }
            }
        }
        stage('push'){
            steps{
                script{
                    docker.withRegistry('https://index.docker.io/v1/', 'dockerhub'){
                        docker.image("${DOCKER_IMAGE}:${env.BUILD_ID}").push()
                    }
                }
            }
        }
    }

    post {
        always{
            cleanWs()
        }
    }
}

