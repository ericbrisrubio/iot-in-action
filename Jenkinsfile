node {
    checkout([$class: 'GitSCM', branches: [[name: '*/master']],
     userRemoteConfigs: [[url:'https://github.com/ericbrisrubio/iot-in-action.git'], [credentialsId: 'b57c3c6c-9e1a-4da7-a269-452a81dbf82e']]])
    def root = tool name: 'go-1.12', type: 'go'
    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
            sh 'go version'
        }
    stage("Test"){
        withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin", "CGO_ENABLED=0", "GO111MODULE=off"]) {
                    sh 'echo RUNNING TESTS'
                    sh 'go test'
                }
    }
    stage("Build"){
        withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
                    sh 'echo BUILDING PROJECT'
                    sh 'go build'
                }
    }
    stage("Deploy"){
         withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
            sshagent(credentials : ['0493bba9-397a-4bfb-a289-d1e10a372476']) {
                sh 'ssh -o StrictHostKeyChecking=no root@ubuntu uptime'
                sh 'ssh -v root@ubuntu'
                //sh 'scp ./source/filename root@ubutnu:/remotehost/target'
            }
         }
     }
}
