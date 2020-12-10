grpc文件生成：protoc --go_out=plugins=grpc:../services grpctest.proto 
gateway文件生成：protoc --grpc-gateway_out=logtostderr=true:../services grpctest.proto 


自签证书  
服务端
1.生成RSA私钥： openssl genrsa -out ca.key 2048  
2.生成pem: openssl req -new -x509 -days 3650 -key ca.key -out ca.pem  
3.生成服务器证书: openssl genrsa -out server.key 2048  
4.生成CSR证书签名请求文件：openssl req -new -key server.key -out server.csr   
5.生成服务器证书：  
openssl x509 -req \
  -CA ca.pem \
  -CAkey ca.key \
  -CAcreateserial \
  -days 365 \
  -in server.csr \
  -out server.pem


客户端
1.openssl genrsa -out client.key 2048  
2.openssl req -new -key client.key -out client.csr  
3.openssl x509 -req \
    -CA ca.pem \
    -CAkey ca.key \
    -CAcreateserial \
    -days 365 \
    -in client.csr\
    -out client.pem

