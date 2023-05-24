This repository is a template with integration of hex architecture, gqlgen and Docker... neo4j as database example.

To run this project you must run the command:

```bash
go mod tidy
```

This downloads the dependencies, after that run:

```bash
go build . && go run .
```

don't forget to have the environment variables set, which you find in the `.sample_env` file.

Finally, to run the project on docker-compose, execute the command:

```bash
docker compose up --build -d
```

## Languages and Tools

<p align="left"> <a href="https://www.docker.com/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/docker/docker-original-wordmark.svg" alt="docker" width="40" height="60"/> </a> <a href="https://graphql.org" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/graphql/graphql-icon.svg" alt="graphql" width="40" height="60"/> </a> <a href="https://www.linux.org/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/linux/linux-original.svg" alt="linux" width="40" height="60"/> </a> <a href="https://go.dev/doc/" target="_blank" rel="noreferrer"><img src="https://www.vectorlogo.zone/logos/golang/golang-icon.svg" width="40" height="60" alt="Go"></a> </p>