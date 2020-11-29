"# Docker-App"  

- Sample-HelloWorld  
  - 1. フォルダを移動する  
  `cd {クローン先フォルダ配下のSample-HelloWorld}`  
  - 2. Docker image の生成  
  `docker image build -t helloworld:latest .`  
  - 3. 実行  
  `docker container run helloworld:latest`  
