curl -d '{"brand":"walon", "price":"1000","color":"blue"}' -H "Content-Type: application/json" -X POST http://localhost:8080/shoes | jq