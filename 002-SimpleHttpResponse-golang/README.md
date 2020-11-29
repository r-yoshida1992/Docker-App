## 002-SimpleHttpResponse-golang  

- 単純な Http request を受け付ける(Go)  

- 起動方法  
  - 1. フォルダを移動する  
  `cd {this folder}`  
  - 2. Docker image の生成  
  `docker image build -t go-web-sample:latest .`  
  - 3. 実行  
  `docker container run -d -p 9000:8080 go-web-sample:latest`  
  - 4. 以下のURLにアクセスする  
  `http://localhost:9000/`  

- イメージ  
![image](https://user-images.githubusercontent.com/64537018/100498348-f03f2600-31a4-11eb-93e4-b758cd1e4aab.png)
