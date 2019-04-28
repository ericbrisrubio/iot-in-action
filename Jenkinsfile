node {
    checkout([$class: 'GitSCM', branches: [[name: '*/master']],
     userRemoteConfigs: [[url:'https://github.com/ericbrisrubio/iot-in-action.git'], [credentialsId: 'b57c3c6c-9e1a-4da7-a269-452a81dbf82e']]])
    tools {
            go 'go-1.12'
        }
        environment {
            GO112MODULE = 'on'
        }
    stage("Compile"){
        sh 'go build'
    }
    stage("Test"){
        sh 'go test'
    }
    stage("Deploy"){
            sh "echo Deployment"
        }
}
