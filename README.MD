### DESCRIPTION

##### Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of ASCII-art project.

#### Authors: Nursultan aka @lunarnuts

```
 _    _   ______   _        _         ____         __          __                 _       _   _
| |  | | |  ____| | |      | |       / __ \        \ \        / /                | |     | | | |
| |__| | | |__    | |      | |      | |  | |        \ \  /\  / /    ___    _ __  | |   __| | | |
|  __  | |  __|   | |      | |      | |  | |         \ \/  \/ /    / _ \  | '__| | |  / _` | | |
| |  | | | |____  | |____  | |____  | |__| |          \  /\  /    | (_) | | |    | | | (_| | |_|
|_|  |_| |______| |______| |______|  \____/            \/  \/      \___/  |_|    |_|  \__,_| (_)


```

#### Usage

##### copy git to your pc, in your command line, change into ascii-art-web directory and run `go build` command. Now you should have an executable named `ascii-art-web`. run it by typing `\.\/ascii-art-web`. Web-server should be hosted at `localhost:8080`, type it to your browser's adress bar. in the web-page you can chose 3 types of font for ASCII. Type your input into the text area and click on `Enter` button. ASCII text is displayed below.

###### only ascii characters from 32 to 126 are supported, else server responds with ERROR 422 Unprocessable entity

#### Dockerize

##### to test Docker, run docker.sh bash-script, to stop docker run docker-stop.sh.

#### Implementation

##### Code is simple, there is a simple handler func that handles main page, if different page is requested, server responds with `Error 404 Not Found`. Function handles only two methods `GET` and `POST`. Get methods fetches pages from server and `POST` payload data as json to server. If other methods are requsted, server responds with `Error 400 Bad Request`. If data is invalid server responds with `Error 422 Unprocessable Entity`. For processing text code from previous project was reused, with slight modifications. Namely, main() was turned into AsciiArt function which received an `array` of `strings` and returns an ASCII text as `string` and error code as `int`. any errors during processing is sent as error codes 404,422 or 200 for `OK`.
