/* groovylint-disable DuplicateStringLiteral, LineLength, NestedBlockDepth, ParameterCount */
@Library('d11-jenkins-lib@master') _
pipeline {
    agent {
        label "devx-auto"
    }

    options {
        ansiColor('xterm')
    }

    stages {
        stage('CheckoutCode') {
            steps {
                script {
                    cleanWs()
                    checkout scm
                }
            }
        }
    }

    stages {
        stage('Installation') {
            steps {
                script {
                  sh """
                    make install
                    go build .
                    sudo mv ./odin /usr/local/bin
                    odin --version
                  """
                }
            }
        }
    }
}

