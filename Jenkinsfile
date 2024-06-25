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
                echo "checking out repo"
                git url: 'https://github.com/audu97/test-project', branch: 'master',
                credentialsId: "${GITHUB_CREDENTIALS}"
            }
        }
        stage('Build'){
            steps{
                echo "starting docker build"
                script{
                    def image = docker.build("${DOCKER_IMAGE}:${env.BUILD_ID}", "-f Dockerfile .")
                }
                echo "docker build completed"
            }
        }
        stage('push'){
            steps{
                echo "pushing to docker hub"
                script{
                    docker.withRegistry('https://index.docker.io/v1/', 'dockerhub'){
                        docker.image("${DOCKER_IMAGE}:${env.BUILD_ID}").push()
                    }
                }
                echo "done"
            }
        }
    }

    post {
        always{
            cleanWs()
        }
    }
}

