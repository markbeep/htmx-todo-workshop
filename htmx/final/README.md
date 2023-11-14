# HTMX Go Template Workshop

## Requirements

There are two ways you can code along:

- Fully use Docker.
- Download all the dependencies manually:

  - Go | [Install](https://go.dev/doc/install): The newer the better. The `go install` command requires Go >=1.18.
  - Templ | [Install](https://templ.guide/quick-start/installation): Used for generating Go code from templates.
    Can be installed as follows:

    ```sh
    go install github.com/a-h/templ/cmd/templ@latest # install via go
    curl -SLO https://github.com/a-h/templ/releases/download/v0.2.432/templ_Linux_x86_64.tar.gz # install binary
    nix develop . # installs all required dependencies
    ```

  - Air | [Install](https://github.com/cosmtrek/air#installation): Allows for hot-reloading the webserver on save.

    ```sh
    go install github.com/cosmtrek/air@latest
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh # then execute the sh script
    nix develop .
    ```

  - Go dependencies: Execute `go get .` while in this directory.

## Starting the webserver

How the server is to be started depends on if you went the Docker or manual route:

- Docker: `docker compose up --build`
- Manual: `air`

The website will not be accessible at http://localhost:3000.

## Helpful extensions for your editor

These are some optional and helpful extensions to improve your coding experience with the tech stack:

- TailwindCSS [Intellisense]
- templ[-vscode]
- Go
