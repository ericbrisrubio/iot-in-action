node {
    checkout([$class: 'GitSCM', branches: [[name: '*/master']],
     userRemoteConfigs: [[url:'https://github.com/ericbrisrubio/iot-in-action.git'], [credentialsId: 'b57c3c6c-9e1a-4da7-a269-452a81dbf82e']]])
    tool {
       go 'go-1.12'
    }
    stage("Build"){
        sh "echo testing"
    }
    stage("Test"){
        sh "echo testing 2"
    }
    stage("Deploy"){
            sh "echo testing 3"
        }
}
