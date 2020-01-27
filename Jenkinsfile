pipeline {
  agent {
    docker {
      image 'golang'
    }
  }

  stages {
    stage('Build') {
      steps {
          sh 'go version'
      }
    }
    stage('Test') {
      steps {
        sh 'go test'
      }
      }
  }
}