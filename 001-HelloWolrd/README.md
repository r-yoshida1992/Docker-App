## 001-HelloWolrd

- sh で hello world する  

- 起動方法
  - 1. フォルダを移動する  
  `cd {this folder}`  
  - 2. Docker image の生成  
  `docker image build -t helloworld:latest .`  
  - 3. 実行  
  `docker container run helloworld:latest`  
