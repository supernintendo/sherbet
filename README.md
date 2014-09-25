# Sherbet #

Sherbet watches files for changes and sends WebSocket messages to a client which acts as a wrapper for your web app. This allows you to do live CSS reloading and trigger other events.

### Ooh, is it good?!

As a delicious frozen treat, yes. As a piece of software, this is probably too early to be of much use.

### Installation

First, build Sherbet with `./build.sh`. Then, add `127.0.0.1   sherbet` to your `etc/hosts` file.

### Running

Run with `./sherbet`. Sherbet will try to load `sherbet.json` from the current directory. You can use `-j` to pass in a JSON file to use instead. See `example/sherbet.json` for an example.

Once Sherbet is up and running, visit `sherbet:7428`. You should see your target site and can begin to send WebSocket messages.
