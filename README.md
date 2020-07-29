# Cubeprox

Cubeprox is a simple web service to serve mock json data with inputed routes and json file.

## Installation

1. Install golang 1.14.2 (download and install documentation in golang pages)
2. Install project dependency by running this command
```bash
go mod download
```
3. Start the application by 
```bash
go run app/main.go
```

## Add New Routes

1. Routes configuration located in ```/configs``` folder
2. Create/Update the ```.yml``` file
3. After creating new ```.yml``` file, insert the required key value
```
prefix: <routes-prefix>
routes:
  - route: <route>
    method: <rest-method>
    json: <json-file-name>
    response_code: <response-code>
  .
  .
  .
```
4. Example for the yml file
```
prefix: users
routes:
  - route: /users
    method: POST
    json: users.json
    response_code: 201
  - route: /users/:id
    method: GET
    json: users_id.json
    response_code: 200
```
5. Add the json file in ```/fixtures/<prefix>``` directory
6. Example for the json file
```
{
  "data": {
    "id": 1,
    "invoice_no": "AX23DES",
    "amount": 250,
    "items": [
      {
        "id": 5,
        "product_name": "Shoes"
      },
      {
        "id": 6,
        "product_name": "Dress"
      }
    ]
  }
}
```
7. Run the web service and try it via cli or other apps

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
