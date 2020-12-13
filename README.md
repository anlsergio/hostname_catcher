# hostname_catcher

Hostname Catcher is a simple CLI tool for getting IP addresses and Hostnames (server names) from a given public domain.

<br />

## Usage
```bash
➜  ./cli_app --help
NAME:
   CLI Application - Get server IPs and Hostnames on the Internet

USAGE:
   cli_app [global options] command [command options] [arguments...]

COMMANDS:
   ip         Get IPs from addreses from the Internet
   hostnames  Get hostnames from the Internet
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

### Example

Here are some examples on how to use it considering the Amazon domain.

<br />

- Getting server IPs regarding the domain `amazon.com`:

```bash
➜  ./cli_app ip --host amazon.com
176.32.103.205
176.32.98.166
205.251.242.103
```
<br />

- Getting hostnames (server names) regarding the domain `amazon.com`:

```bash
➜  ./cli_app hostnames --host amazon.com
ns4.p31.dynect.net.
pdns1.ultradns.net.
ns1.p31.dynect.net.
pdns6.ultradns.co.uk.
ns2.p31.dynect.net.
ns3.p31.dynect.net.
```
<br />

## Disclaimer

This simple app was made as an exercise from the DevBook course.
