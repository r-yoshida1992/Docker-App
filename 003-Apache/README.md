## 003-Apache  

- 静的な web page を返す  

- 起動方法  
  - 1. フォルダを移動する  
  `cd {this folder}`  
  - 2. Docker image の生成  
  `docker image build -t apache-sample:latest .`  
  - 3. 実行  
  `docker container run -d -p 8080:80 apache-sample:latest`  
  - 4. 以下のURLにアクセスする  
  `http://localhost:8080/`  