pipeline {
    agent any

    tools {
       go 'golang-1.22.4'
    }
    environment {
        DOCKERHUB_CREDENTIALS = credentials('dockerhub')
//         DOCKER_IMAGE = 'ephraimaudu/test-app'
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
                    sh 'docker build -t ephraimaudu/test-app .'
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

