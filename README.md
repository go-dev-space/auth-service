# auth-service
![Hero Image](./artifacts/general/img/hero.jpg)

An example of an authorization microservice in Go according to DDD with EDA in a kubernetes cluster.

## Quick Links
| Link | Description | Credential |
|------|-------------|------------|
|http://localhost:8080 | Local development | air-verse |
|http://localhost:50015 | Local development | minikube |
|[Repositry Link](https://github.com/go-chi/chi) | router framework pkg | go-chi/chi |
|[Repositry Link](https://github.com/joho/godotenv) | env files loader pkg | joho/godotenv |
|[Repositry Link](https://github.com/golang-migrate) | postgre migration pkg | golang-migrate |
|[Repositry Link](https://pkg.go.dev/google.golang.org/grpc) | gRPC package | google.golang.org/grpc |
|[Repositry Link](https://github.com/go-playground/validator) | validation pkg | go-playground |
|[Repositry Link](https://github.com/jackc/pgx) | postgresql driver pkg | jackc/pgx |


## Development
### Overview and local setup
This Go project is structured as follows:

```
├── auth-service
│   ├── bin
│   │   └── *                         # Compiled binary
│   ├── cmd
│   │   └── api
│   │        ├── main.go              # Entry point of the app
│   │        ├── router.go            # Router setup
│   │        └── middlewares.go       # Global middlewares
│   ├── internal
│   │   └── auth                      # Bounded context
│   │        ├── application          # Business logic & use cases
│   │        ├── domain               # Entities & aggregates
│   │        ├── infrastructure       # Technology & implementation
│   │        │    ├── crypto          # Crypto services
│   │        │    ├── database        # Database clients
│   │        │    └── grpc            # gRPC services
│   │        │    │   ├── generated   # Generated go files
│   │        │    │   └── proto       # Proto files
│   │        │    ├── mailer          # Mailer services
│   │        │    ├── store           # Store services
│   │        │    └── validator       # Validator services
│   │        └── interfaces           # Interface to the outside world
│   │             ├── dto             # Data transfer objects
│   │             └── middlewares     # Bounded context middlewares
│   ├── pgk
│   │   └── logs                      # Global log service
│   └── .env                          # Environment variables for local dev
```

**Note:** 
This project uses the standard Go `chi router` package for HTTP routing.
All routes and middleware are defined in `/cmd/api`, and actual handler logic is encapsulated in `/internal/[bounded-context]/interfaces`.

### First time initialisation
As soon as the project has been checked out from the Git repository, all required packages must be installed locally. Start in the root directory of your project:

```
# Install the Go packages
$ go mod tidy
```

### Run the local development environment
A Makefile is used to make it easier to set up the local development environment. To do this, please carry out the following steps.

**Step 1:** Start postgres service

```
# Enter in your terminal 
$ make docker_dev_up
```

As soon as the Postgres service is running, the migration of the SQL data can begin. The Docker container name is **`pgauth`** and can be checked with this command: `docker ps -a`.

**Step 2:** Migrate SQL Data

```
# Enter in your terminal 
$ make migrate_up
```
As this is a small sample application, the database has been kept minimalist and the structure is shown below.

![Database](./artifacts/general/img/tables.png)

Before we can start the application, we have to decrypt two files, the **`.env.enc`** and the **`postgres-secret.yml.enc.`** A gpg private key is required for decryption, which must be requested in person and is **`not included`** with the project.

**Step 3:** Decrypt files
```
# Enter in your terminal 
$ make decrypt
$ make decrypt_secret
```

Now the application can be started with **`air-verse`** and the development environment is ready.

**Step 4:** Start the application
```
# Enter in your terminal 
$ air
```

## Call the Healthcheck route
### Local Maschine
You should see the following debug output in your terminal:

```
INFO    2025/04/26 15:59:51 service running:8080
```
Open a http client of your choice and enter the following URL:
```
http://localhost:8080/v1/healthcheck
```
Add the following header for the GET request:
```
key: X-Access-Header
value: 2cf24dba5fb
```
You should then see the following debug output on your terminal:
```
{
    "status_code": 200,
    "message": "pong",
    "data": null,
    "error": false
}
```
## Kubernetes
### Local simulation with minikube

For the Kubernetes simulation in Minikube, the prerequisite is an already installed and running Minikube environment on the local host system. The attached Kubernetes YAML files under **`./kubernetes`** create the Kubernetes cluster in the following figure.

![Kubernetes](./artifacts/general/img/kubernetes.png)

The first thing to do is to start minikube (and the dashboard). I have chosen docker as the driver.
```
# Enter in your terminal:

# start minikube
$ minikube start --driver=docker 

# status check
$ minikube status 

# start dashboard
$ minikube dashboard 
```

The following steps are necessary to create the Kubernetes deployment:

**Step 1:** Create the Kubernetes Secret for ghcr.io
```
# Enter in your terminal:
$ kubectl create secret docker-registry ghcr-secret \
  --docker-server=ghcr.io \
  --docker-username=GITHUB_USERNAME \
  --docker-password=GITHUB_PERSONAL_ACCESS_TOKEN \
  --docker-email=you@example.com
```

**Step 2:** Create persistent volume & persistent volume claim
```
# Enter in your terminal:
$ kubectl apply -f persistent-volume.yml
```

**Step 3:** Create secret
```
# Enter in your terminal:
$ kubectl apply -f postgres-secret.yml
```

**Step 4:** Create postgres deployment & service
```
# Enter in your terminal:
$ kubectl apply -f deployment-postgres.yml service-postgres.yml
```
Wait briefly (approx. 5-10 seconds) until Postgres is running and then check whether the POD has been started
```
# Enter in your terminal:
$ kubectl get pods
```

**Step 5:** Create Go app deployment & service
```
# Enter in your terminal:
$ kubectl apply -f deployment.yml service.yml
```

To summarise, these are all the steps again:
```
# Enter in your terminal:
$ kubectl apply -f persistent-volume.yml
$ kubectl apply -f postgres-secret.yml
$ kubectl apply -f deployment-postgres.yml
$ kubectl apply -f service-postgres.yml

# wait until postgres is running
$ kubectl apply -f deployment.yml
$ kubectl apply -f service.yml
```

To activate the Minikube LoadBalancer, one last command must be entered:

```
# Enter in your terminal:
$ minikube service <METADATA NAME HERE>
```

## Documentation
### For further information 

On my website you will find a complete documentation of the code and all further information under the corresponding headings.

To the official: [Documentation](https://github.com/joho/godotenv)


## Licence
MIT License

Copyright (c) 2025 Gopher

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.