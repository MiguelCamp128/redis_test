pipeline {
     agent { label 'agente_monitor' }
     stages {
         stage("init") {
            steps {
                sh 'cp -r ./* /home/jenkins'
            }
        }
         stage("get libraries") {
             steps {
                 sh 'cd /home/jenkins && go get'
             }
         }
         stage("SonarQube Analysis") {
             environment {
                 scannerHome = tool 'sonar-scanner'
             }
        //      steps {
        //          withSonarQubeEnv('sonarQube') {
        //              sh "${scannerHome}/bin/sonar-scanner \
        //                      -D sonar.projectKey=Go_ClientRedis:test \
        //                      -D sonar.projectName=Go_ClientRedis \
        //                      -D sonar.host.url=http://10.0.103.75:9000 \
        //                      -D sonar.sources=."
        //          }
        //          timeout(time: 3, unit: 'MINUTES') {
        //              waitForQualityGate abortPipeline: true
        //          }
        //      }
        //  }
         stage("start monitor") {
             steps {
                 sh 'cd /home/jenkins && go run .'
             }
         }
     }
}