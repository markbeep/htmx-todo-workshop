# htmx Go Template Workshop

This is a barebones example on how you can create a single-page application todo-list using htmx, Go with the templ library for templates in the backend, and tailwindcss for styling elements.

_Not a single line of javascript will be written._

What you'll be creating:

![Alt text](media/website.png)

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

The website will now be accessible at http://localhost:3000.

## Helpful extensions for your editor

These are some optional and helpful extensions to improve your coding experience with the tech stack:

- TailwindCSS [Intellisense]
  - Additionally add `templ` to the tailwind config so that it runs in templ files.
    Vscode JSON User Settings as an example:
    ```json
    "tailwindCSS.includeLanguages": {
      ...
      "templ": "html"
    }
    ```
- templ[-vscode]
- Go

# Workshop

## What is htmx?

[htmx](https://htmx.org/)' motto is to stay simple. Instead of throwing huge chunks of javascript at a user, we only work with the actual required html. Using htmx we can make buttons responsive and replace elements on the website without reloading the website. The basics are, that when we click a button that should modify the site, we send a request to the backend server. The backend then doesn't return JSON, but the direct HTML which will be used to place wherever defined.

## Base project layout

### Go

This project has a small amount of Go code. `main.go` is the whole webserver. For handling requests we use [chi](https://github.com/go-chi/chi) which is a lightweight http library. It allows for some nice features like easily adding middlewares and getting values from routes.

The other file is `internal/todo_type.go` and this file only defines a single type definition. Go doesn't allow for circular dependencies and we need it in both main.go and the templ templates.

### Templ

We're building a template driven webserver (similar to Django). Go has a template library out of the box (`template/html`), but it is not as clean as it could be. Instead for this project we'll be using [templ](https://templ.guide/). Templ goes in the direction of components, similar to React. Like in the code snippet below, we can easily create a component (like `Hello`) and then easily use that wherever we want (using `@Hello()`). This gives a lot of freedom for how to use our components and it allows for a very clear structuring of the website. The syntax is also basically Go (with the same types), but we sprinkle in some html. Templ will generate Go code which then can be imported in `main.go`. It is important to make sure the generation runs and works.

```go
package main

templ Hello(name string) {
  <div>Hello, { name }</div>
}

templ Greeting(person Person) {
  <div class="greeting">
    @Hello(person.Name)
  </div>
}
```

### Tailwindcss

For styling everything, we use [Tailwindcss](https://tailwindcss.com/). It's a must-have for designing a website. It allows you to more easily write and use css. For this project, most of the tailwindcss should already be added to components though, so you don't really have to touch it.

## Code along

Now it's time to code along. This is basically how the workshop will run:

1.

## Going further

### User specific todo list

The site works as of now, but as soon as you start having more users, you'll quickly notice a problem: The todo list is shared by all users, since it's completely server side in a single slice. You'll then want to start looking into storing data on a per-user level. In the simplest terms this can be done by creating a map with the keys being a user's hostname and the value being their list of todos. [Gorilla sessions](https://github.com/gorilla/sessions) allow this to be implemented in a more secure way.

### Optimized tailwindcss

Currently we load all of tailwindcss which is already 100kB in transfer size alone. This is quite inefficient. Using the tailwind CLI tool we can build a custom css file which only defines the tailwind elements we're actually using. This is described in the [install](https://tailwindcss.com/docs/installation) section of tailwindcss. This reduces the css size to a few kilobytes.
