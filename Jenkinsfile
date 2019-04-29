node {
    checkout([$class: 'GitSCM', branches: [[name: '*/master']],
     userRemoteConfigs: [[url:'https://github.com/ericbrisrubio/iot-in-action.git'], [credentialsId: 'b57c3c6c-9e1a-4da7-a269-452a81dbf82e']]])
    def root = tool name: 'go-1.12', type: 'go'
    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
            sh 'go version'
        }
    stage("Test"){
        withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin", "CGO_ENABLED=0", "GO111MODULE=off"]) {
                    sh 'echo $PATH'
                    sh 'echo $GOROOT'
                    sh 'go test'
                }
    }
    stage("Build"){
        withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
                    sh 'go build'
                }
    }
    stage("Deploy"){
            sh "echo testing 3"
        }
}
