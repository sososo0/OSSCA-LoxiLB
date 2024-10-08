# Week 1 

## What is Load Balancer? 

Load distribution or load balancing is a type of computer network technology that refers to dividing tasks to computer resources such as two or more central processing units or storage devices. 

In summary, it is a device that eliminates bottlenecks in data traffic, ensuring efficient handling of high user volumes. 

## About LoxiLB 

LoxiLB is a SW Load Balancer with three planes.

- **Management plane** : 
    - **loxicmd**: Command-line tool to manage LoxiLB.
    - **kube-loxilb**: Kubernetes integration tool for managing LoxiLB within Kubernetes environments.

- **Control plane** : 
    - **API server**: Manages interactions and communications between LoxiLB components and external systems. 
    - **LoxiNLP**: A module for network logic processing.
    - **goBGP**: A BGP (Border Gateway Protocol) implementation in Go used for dynamic routing.
    - **netlink**: A communication protocol between the kernel and userspace, used to configure network interfaces, IP addresses, routes, etc. 

- **Data plane** : 
    - **eBPF**: A high-performance in-kernel technology that enables efficient packet processing and forwarding. 

## About REST API 

REST refers to a set of network architecture principles. Here, the term 'network architecture principles' encompasses the overall method of defining a resource and specifying its address.

Simply put, it refers to performing CRUD operations (Create, Read, Update, Delete) using standard methods such as POST, GET, PUT, DELETE, and others.

## About YAML 

To develop APIs using Swagger, you first define the Swagger specificationin YAML format to structure the API. Then, you generate a skeleton server based on that specification and proceed to add your business logic into the generated code. 

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

- JSON is easy for computers to read during programming but can inconvenient for humans to read.

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

- YAML was introduced to address the readability issues of JSON.
- It mainly divided into three parts.

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

5. To add business logic to the Swagger-generated server, navigate to the next files where the API endpoints are defined.

> restapi/configure_company_rest_api.go


1. An example of adding business logic to a GET API in a Swagger-generated Go server. 

```
	if api.GetAccountAllHandler == nil {
		api.GetAccountAllHandler = operations.GetAccountAllHandlerFunc(func(params operations.GetAccountAllParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetAccountAll has not yet been implemented")
		})
	}
```

- Inside the code, you can see the process of registering a handler.
- The handler is designed to accept a single function. 
    - As the logic becomes more complex, it can make the code harder to manage, so it's better to create separate handlers and add them individually.

2. Create a handler folder under the restapi directory, and inside the handler folder, create a account.go file.

```
package handler

import (
	"swaggertest/models"
	"swaggertest/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func ConfigGetAccount(params operations.GetAccountAllParams) middleware.Responder {
	var result []*models.AccountEntry
	result = make([]*models.AccountEntry, 0)
	result = append(result, &models.AccountEntry{
		UserID:   "idid",
		Password: "passwords",
		Email:    "test@test.io",
	})
	return operations.NewGetAccountAllOK().WithPayload(&operations.GetAccountAllOKBody{Attr: result})
}
```

- The handler includes a function that processes the request for GET/account and returns a list of AccountEntry objects. 
- You can add logic to fetch real data.

3. 
```
import (
    ...

	"swaggertest/restapi/handler"
    
    ...
)


func configureAPI(api *operations.CompanyRestAPIAPI) http.Handler {

    ...

	api.GetAccountAllHandler = operations.GetAccountAllHandlerFunc(handler.ConfigGetAccount)

    ...
}

```

- To complete the setup, you'll need to add the handler path in your project and then register the handler function in the ```configure_company_rest_api.go``` file.

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

## About CLI 

The command-line interface (CLI), also known as a command interface, refers to how users interact with a computer by typing text-based commands through a virtual terminal or terminal.

#### Command Shell vs. Prompt Shell

- **Command Shell**: Modern and user-friendly but often associated with lower security. 
- **Prompty Shell**: More traditional and older, but generally offers higher security. 

### CLI Design 

```
# prgramName (command) (condition) (option)
```

- command : (CRUD operations)
- condition : (Name of the REST API)
- option : (parameters, body, etc., in the REST API)

**Example.**

```
loxicmd get lb -o wide 
```

```
kubectl get pod -o wide
```

#### Basic Design Principles 

1. **Consistency**: Maintain a clear and consistent standard for commands (e.g., add, create).
2. **Simplicity**: Ensure the commands are easy for users to understand and use.
3. **Intuitiveness**: Make the commands intuitive so users can understand them easily (e.g., loxicmd get lb).

#### Help Design (-h, --help)

- **Examples**: Providing examples is crucial, especially for actions like adding or deleting resources.
- **Description**: If the design is intuitive, explanations can be brief and still easily understood. 

### About Cobra 

Cobra is a library that provides a simple interface for creating powerful, modern CLI applications, similar to tools like Git and Go.

- A library for building CLI applications in Golang.
- Chosen for its consistency, as it is also used by Kubernetes. 

### Command Code Structure Analysis 

This structure clearly organizes command definites and actions, making it easier to maintain and expand the CLI functionality.  

```
├── cmd
│   ├── root.go        # Defines main actions (get, create, delete)
│   ├── create         # Directory for the create command
│   ├── delete         # Directory for the delete command
│   └── get
│       └── get.go     # Defines the specifics of the get command (e.g., LB, account)
└── main.go            # Main entry point for the program
```

- cmd : This directory contains the code related to the design and structure of commands. 
    - cmd/root.go: This file defines the primary actions (e.g., get, create, delete) for the CLI. 
    - cmd/get/get.go: This file defines the details of the get command (e.g., specifying options like load balancers(LB), accounts, etc.). 

#### CLI Design Practice 

Design a CLI for creating routes using the ```POST /config/route``` API, which accepts a JSON payload like the following: 

```
{
  "destinationIPNet": "DstValue",
  "gateway": "GWValue",
  "protocol": "PROTO"
}
```

**Example**:

To create a new route using this API, the corresponding CLI command would be structured as:

```
loxicmd create route DstValue GWValue --protocol=PROTO
```

- **loxicmd**: The name of the CLI tool.
- **create**: The action to create a new route (similar to a POST request).
- **route**: The resource you are managing (in this case, the route).
- **DstValue**: The destination IP network (destinationIPNet).
- **GWValue**: The gateway (gateway).
- **--protocol=PROTO**: The optional protocol parameter (protocol), passed as a flag.

> This design makes it simple, intuitive, and easy to understand, closely matching the structure of the API request while offering a user-friendly interface for interacting with the system.

<br>

**Example. (examples/cli/cmd/root.go)**

```
var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}
)
```

In a basic Cobra-based CLI application, commands like the root command typically include the following fields. 

- **Use**: This defines how the command is used. It's essentially the command name and any arguments or subcommands that might be used.
    - ```"cobra-cli"``` in the code means the root command will be called **cobra-cli** in the CLI.
- **Short**: A brief description of what the command does. This is typically shown when a user runs ```--help``` or needs a quick overview.
    - "A generator for Cobra based Applications" gives a concise description.
- **Long**: A more detailed explanation of the command and its purpose. It often includes more context about how the command works and what it can be used for.
    - This is the detailed description that explains Cobra's purpose and how the CLI application works. In this case, it describes that Cobra is a CLI library for Go and that this tool helps generate the necessary files to quickly create a Cobra application.

**Example. (examples/cli/cmd/get/get.go)**

The additional part of the command structure is the RUN function, where the actual logic to be executed when the command is run is added. This is where you implement the core functionality of the command, such as processing inputs, performing actions, and interacting with APIs, or other resources. 

```
func GetCmd(restOptions *api.RESTOptions) *cobra.Command {
	//func GetCmd() *cobra.Command {
	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "get a Load balance features in the LoxiLB.",
		Long: `get a Load balance features in the LoxiLB.
Create - Service type external load-balancer, Vlan, Vxlan, Qos Policies, 
	 Endpoint client,FDB, IPaddress, Neighbor, Route,Firewall, Mirror, Session, UlCl
`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("get called!\n", args)
		},
	}

	getCmd.AddCommand(NewGetAccountCmd(restOptions))
	return getCmd
    }
}
```

- This is where the actual logic of the command will be placed. 
    - In the Run section, you can add additional logic like:
    - **Making API calls**: You coud use restOptions to make REST API requests and fetch data.
    - **Error handling**: Catch errors from the API or any other part of your code.
    - **Printing detailed output**: After fetching the data, you can format and print it.
- **AddCommand**: Add subcommands for specific resources (e.g., accounts)

### Command and API Integration Structure 

```
├── cmd
│   ├── root.go                # Defines main command actions (create, delete, get)
│   ├── create                 # Folder for 'create' command logic
│   ├── delete                 # Folder for 'delete' command logic
│   └── get                    # Folder for 'get' command logic
│       └── get_account.go     # Command to handle 'get account' actions
├── api
│   ├── common.go              # Handles common logic to call APIs from commands
│   ├── client.go              # Defines specific API client logic
│   ├── rest.go                # Core logic for making CRUD REST API calls
│   └── account.go             # Handles API calls related to 'account'
└── main.go                    # Entry point of the program
```

The cmd and api directories are structured to organize how commands interact with APIs. 

- **cmd**: Contains command-related logic
- **api**: Handles REST API calls.
    - **api/rest.go**: Manages the general structure and flow of CRUD REST API calls. It defines how the requests are formed, sent, and how responses are handled.
    - **api/common.go**: Contains helper functions that abstract REST API calls so that they can easily be invoked from commands.
    - **api/client.go**: Defines the logic for interacting with specific APIs. It provides functions that allow various API clients to make request and receive responses. 
    - **api/account.go**: The actual code that interacts with the Account-related APIs. It defines the REST calls needed to get, create, update, or delete an account. 
