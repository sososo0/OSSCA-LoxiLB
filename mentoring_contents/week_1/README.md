# Week 1 

## What is Load Balancer? 

Load distribution or load balancing is a type of computer network technology that refers to dividing tasks to computer resources such as two or more central processing units or storage devices. 

In summary, it is a device that eliminates bottlenecks in data traffic, ensuring efficient handling of high user volumes. 

## About LoxiLB 

LoxiLB is a SW Load Balancer. 

3가지 plane이 있다. 

- Management plane : command를 날릴 수 있는 loxicmd, 쿠버네티스 환경에서 manage할 수 있는 곳 
- Control plane : 
- Data plane : 

## About REST API 

REST refers to a set of network architecture principles. Here, the term 'network architecture principles' encompasses the overall method of defining a resource and specifying its address.

Simply put, it refers to performing CRUD operations (Create, Read, Update, Delete) using standard methods such as POST, GET, PUT, DELETE, and others.

## About YAML 

API를 개발하기 위해서 Swagger를 사용하는데 Swagger를 사용하기 위해서 YAML로 Swagger에 대한 포멧을 맞춰놓고 껍데기 서버를 만들고 로직을 넣는다. 

YAML is a human-readable data serialization format, inspired by concepts from email formats defined in XML, C, Python, Perl, and RFC2822. 

- A type of data structure. 
    - It's a data type structured around key-value pairs. 
- Similar to Python's Dict or JSON, but created as an alternative due to JSON's limitations. 

### JSON vs. YAML 

#### JSON 

```
{
  "apiVersion": "v1",
  "kind": "Pod",
  "metadata": {
    "name": "nginx"
  },
  "spec": {
    "containers": [
      {
        "name": "nginx",
        "image": "nginx:1.14.2",
        "ports": [
          {
            "containerPort": 80
          }
        ]
      }
    ]
  }
}
```

- JSON은 프로그래밍 할 때 컴퓨터가 읽기는 쉽지만, 사람이 읽기에는 불편하다.


<br>

#### YAML 

```
apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  containers:
  - name: nginx
    image: nginx:1.14.2
    ports:
    - containerPort: 80
```

- 위의 JSON의 단점을 보완하기 위해 나온 것이 YAML이다. 
- 크게 3가지로 나뉜다. 

##### Type 1. Map 

```
apiVersion: v1
```

- Rules : Combination of Key, Value, and spaces. (Here, **spaces are crucial.**)
- Syntax : 
    - A key-value pair forms a map, with a space required after the colon in ```key: value.```
    - Spaces can be included in both the key and value, though it's uncommon, especially for keys. Example: ```key key: value value``` ( ```{"key key" : "value value"}``` )
    - When a value is another map, it's represented on the next line with indentation. Example: 
        ```
        key:
        key2: value
        ```
        (Equivalent to : ```{"key": {"key2" : "value"}}``` )
    - If ```key2: value``` is written without spaces, it is treated as a single string. Example: ```{"key": "key2:value"}```

##### Type 2. Arrays 

- Rules : Combination of key, value, spaces, and hyphens(-). Used when a key has multiple values. 
- Syntax : 
    - Represent arrays similar to Python's array (using ```[ ]``` ) Example.
        ```
        key: [ value1, value2 ]
        ```
        (Equivalent to:  ```{"key" : ["value1", "value2"]}``` )
    - Use a hyphen (-) to list values: Example.
        ```
        key: 
          - value1
          - value2 
        ```
    
**Case 1.**

When an array contains maps as values: Example.
```
key: 
  - value1
  - key2: value2 
```

**Case 2.**

If the second value is not prefixed with a hyphen(-), it's treated as part of the previous value: Example.
```
key: 
  - value1
    value2
```

- Equivalent to : ```{"key": ["value1 value2"]}```

##### Type 3. ref 

- Rules: ```$ref``` is used to reference a defined component. 

Example. 
```
components:
  schemas:
    User:
      properties:
        id:
          type: integer
        name:
          type: string
```

- ```$ref``` refers to the predefined User schema under components/schemas. 

```
responses:
  '200':
    description: The response
    schema:
      $ref: '#/components/schemas/User'
```

## About Swagger 

Swagger is an open-source software framework supported by a large ecosystem of tools that helps developers design, build, document, and consume RESTful web services. 

LoxiLB uses Swagger for creating and managing API documentation. Swagger is used to create documentation with YAML, which also facilitates server generation.

Swagger can be broadly divided into **three main categories**: 

#### Basic 

Basic Information in Swagger:
- **Version**: Specifies the API version.
- **General Information**: Provides details about the API, such as its name, description, and purpose.
- **Documentation Version**: Indicates the version of the documentation itself.
- **Protocols**: Defines the supported communication protocols (e.g., HTTP, HTTPS).
- **Host**: Specifies the server hosting the API (e.g., api.example.com).

```
swagger: '2.0'
info:
  title: Company REST API
  description: Company REST API for Baremetal Scenarios
  version: 0.0.1
schemes:
  - http
  - https
host: "0.0.0.0:11111"
basePath: /company_name
produces:
  - application/json
consumes:
  - application/json
```

#### Path 

In Swagger, a typical URI path includes the following components:
- **Actual Path**: Represents the URI endpoint.
- **Description**:: Provides a brief and detailed explanation of what the endpoint does. 
- **Parameters**: 
    - **Query Parameters**: For query strings (e.g., ```?paramName=value``` ). 
    - **Body Parameters**: Data sent in the body of the request (e.g., POST/PUT requests).
    - **Path Parameters**: Used for dynamic URI parts (e.g., ```/user/{id}``` ).
- **Response** 

```
paths:
  /account:
    post:
      summary: Create a new account
      description: Create a new account with the specified attributes.
      parameters:
        - name: attr
          in: body
          required: true
          description: Attributes for load balance service
          schema:
            $ref: '#/definitions/AccountEntry'
      responses:
        '200':
          description: OK
          schema:
            $ref: '#/definitions/PostSuccess'
```

#### Definition 

In Swagger, Definitions are used to describe the structure of objects, especially those that will be sent in the body of a request or received in a response.  

```
definitions:
  Error:
    type: object
    properties:
      code:
        type: integer
        format: int32
        description: Main error code
      sub-code:
        type: integer
        format: int32
        description: Additional sub-code for more detailed error categorization
      message:
        type: string
        description: Description of the error
      fields:
        type: array
        items:
          type: string
        description: List of fields involved in the error (if applicable)
      details:
        type: string
        description: Additional information about the error
```

### Practice to create a server using Swagger

1. Navigate to the directory where **swagger.yaml** is located and Run the command to generate the Server.

```
# go mod init swaggertest
```

2. To build and run Swagger-based Go server in Docker, Run the command after generating the go.mod.

```
# sudo docker run --rm -it  --user $(id -u):$(id -g) -e GOPATH=$(go env GOPATH):/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger:0.30.3 generate server
```

3. After generating the Go server code with Swagger, Run the command to ensure all necessary dependencies are downloaded and included in go.mod file.

```
# go mode tidy
```

4. Run the command to generate a server using Swagger, it essentially create the "skeleton" of the API server.

```
# go build cmd/company-rest-api-server/main.go
```

<br>

> While the server structure is generated, it is up to developer to fill in the logic behind the API to make it fully functional. \
In short, the **Swagger-generated server is a starting point a functional skeleton** that needs further development to meet your application's needs.

### Reference - RESTful API Design 

- **Basic Design Principles**: 
    - **Convention**: Use lowercase, hyphens(-), avoid trailing slaches, and use plural nouns.
    - **Consistency**: Establish and follow your own consistent standards.
    - **Simplicity**: Ensure it is easy for users to interact with the API. 
    - **Intuitiveness**: Design the API so that users can understand it at a glance.

- **Practical Options**:
    - Use a single entity from the database schema as the foundation (e.g., User, Product, Config).
    - Display the ID first, followed by details (e.g., ```user/detail/{ID}``` becomes ```/user/{id}/detail``` ).
    - When details don't change, arrange endpoints from general to specific (e.g., ```/v1/config/...``` ).





