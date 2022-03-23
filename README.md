# Spinner

> Spinner starts and then stops when you tell it to.

## Getting started

The following example illustrates building and running the spinner app. Spinner exists with a 0 exit code because it handles the SIGTERM Unix signal.

1. Get the code.

   ```
   git clone ...
   cd spinner
   ```

1. Build the app.

   ```
   go build -o spinner spinner.go
   ```

1. Run spinner.

   ```
   ./spinner && echo $?
   ```

1. Open a new terminal.

1. Send the SIGTERM Unix signal to spinner.

   ```
   pkill spinner
   ```

1. Notice that spinner cleanly exits.

   The output should look like the following:

   ```
   $ ./spinner && echo $?
   Got args: []
   Got signal: terminated
   0
   ```

## Docker example

The following example illustrates building a spinner Docker image and passing command-line arguments. You can stop the Docker image because spinner handles the SIGTERM Unix signal that Docker sends to your container when you run the `docker stop` command.

1. Get the code.

   ```
   git clone ...
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

1. Open a new terminal.

1. Stop the container.

   ```
   docker stop spinner
   ```

1. Notice that spinner cleanly exits.

   The output should look like the following:

   ```
   $ docker run -it --rm --name spinner spinner foo bar baz && echo $?
   Got args: [foo bar baz]
   Got signal: terminated
   0
   ```
