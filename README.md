# Spinner

> Spinner starts and then stops when you tell it to.

## Getting started

The following procedure illustrates how to build and run the _spinner_ Go app. App _spinner_ exits with a 0 exit code because it handles the SIGTERM Unix signal.

1. Get the code.

   ```
   git clone git@github.com:mbigras/spinner.git
   cd spinner
   ```

1. Build the app.

   ```
   go build
   ```

1. Run spinner.

   ```
   ./spinner && echo $?
   ```

1. In a new terminal, send the _`SIGTERM`_ Unix signal to spinner.

   ```
   pkill spinner
   ```

1. Notice that spinner cleanly exits.

   Your output should look like the following:

   ```
   $ ./spinner && echo $?
   Running as pid: 14897
   Got args: []
   Got signal: terminated
   0
   ```

## Docker example

The following example illustrates how to build and run a _spinner_ Docker container. You can stop the container because _spinner_ handles the SIGTERM Unix signal that Docker sends to your container when you run the `docker stop` command.

1. Get the code.

   ```
   git clone git@github.com:mbigras/spinner.git
   cd spinner
   ```

1. Build the Docker image.

   ```
   docker build -t spinner .
   ```

1. Run the Docker container.

   ```
   docker run -it --rm --name spinner spinner foo bar baz && echo $?
   ```

   You pass the foo, bar, and baz command-line arguments. Docker passes those arguments to spinner because the [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint) Docker directive tells Docker to treat spinner as the main exectuable to run.

   Also, the exec ENTRYPOINT form—`ENTRYPOINT ["/app/spinner"]`—ensures that `/bin/sh` shell running inside your Docker container won't block Unix signals.

1. In a new terminal, stop the container.

   ```
   docker stop spinner
   ```

1. Notice that spinner cleanly exits.

   Your output should look like the following:

   ```
   $ docker run -it --rm --name spinner spinner foo bar baz && echo $?
   Running as pid: 1
   Got args: [foo bar baz]
   Got signal: terminated
   0
   ```
