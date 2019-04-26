node {
    checkout([$class: 'GitSCM', branches: [[name: '*/master']],
     userRemoteConfigs: [[url: 'http://git-server/user/repository.git'], [credentialsId: '9fe8edbd-979d-45b2-812c-8e4e1d8aae17']]])
    stage("go build"){
        sh "echo testing" 
    }
    stage("go test"){
        sh "echo testing 2"
    }
} 
