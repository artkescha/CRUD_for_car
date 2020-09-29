# CRUD_for_car
implement jsonAPI CRUD for car

---
**Create a new car**:
	curl -X POST localhost:81/v1/cars -d '{"data" : {"type" : "cars" , "attributes": {"vendor":"1", "model" : "1", "status": "in transit"}}}'
---	
**List users**:
	curl -X GET localhost:81/v1/cars
---	
**List paginated users**:
	curl -X GET 'http://localhost:81/v1/cars?sort=id'	
OR
	curl -X GET 'http://localhost:81/v1/cars?sort=id&page[offset]=0&page[limit]=5'
---
**Update**:
	curl -vX PATCH 'http://localhost:81/v1/cars/9bfe969c-30e4-428a-8a2a-12714e6df5f8 -d '{"data" : {"type" : "cars", "id": "9bfe969c-30e4-428a-8a2a-12714e6df5f8", "attributes": {"model" : "12123"}}}'
---
**Delete**:
	curl -vX DELETE http://localhost:81/v1/cars/af0646d1-b7e9-44c2-8c78-0e7816276fef
