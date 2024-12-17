# Golang rest
## 
```
GET /users 
--list of users 
200, 404, 500
```
```
GET /users/:id 
--user by id 
200, 404, 500
```
```
POST /users/:id 
--create a new user 
204, 4xx
Header location: url
```
```
PUT /users/:id 
--fully update a user 
204/200, 404, 400, 500
```
```
PATCH /users/:id 
--partially update a user 
204/200, 404, 400, 500
```
```
DELETE /users/:id
--delete a user 
204, 404, 400
```