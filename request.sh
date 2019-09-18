# jajajajja walon jajjjajajajaja chavex xdxd jajajajaj
curl -d '{"brand":"walon", "price":1000,"color":"blue"}' -H "Content-Type: application/json" -X POST http://localhost:8080/shoes -v | jq

#Login
curl -d '{"email":"thomtwd@gmail.com", "password":"thom"}' -H "Content-Type: application/json" -X POST http://localhost:8080/users -v | jq