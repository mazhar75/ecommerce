 
POST  /create-product api end point:

 C:\Users\USER>curl -X POST http://localhost:9090/create-product -H "Content-Type: application/json" -d "{\"name\":\"Laptop\",\"type\":\"Electronics\",\"price\":1200,\"img_url\":\"https://example.com/laptop.png\"}"
Resource created successfully

GET /products api end point:

C:\Users\USER>curl -X GET "http://localhost:9090/products"
[{"id":1,"name":"Mango","type":"Fruits","price":70,"img_url":""},{"id":2,"name":"Mouse","type":"Electronics","price":100,"img_url":""},{"id":3,"name":"Sharee","type":"Shopping","price":999,"img_url":""},{"id":0,"name":"Laptop","type":"Electronics","price":1200,"img_url":"https://example.com/laptop.png"}]
