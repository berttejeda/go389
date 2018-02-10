## go389
A Simple LDAP Server

### Features
- Pluggable backend
  - LDAP
  - YAML
- Pluggable Auth
  - Simple (hash based - currently sha256)
  - PAM

### Build
```
go get
go build -i -tags "yaml pam"
```

##### Supported Build Tags
- Backend
  - yaml
- Auth
  - pam
- App Config
  - ini

### Sample App Config
```yaml
server: ldap
backend: yaml

servers:
  ldap:
    Bind: "localhost:8389"

backends:
  yaml:
    Path: db.yml
    BaseDN: "dc=example,dc=com"

auths:
  pam:
    Service: go389
```

### Sample Backend YAML DB
```yaml
groups:
  - name: sadmin
    uid:  9001
  - name: admin
    uid:  9002

users:
  - name: suser1
    uid: 9001
    gid: 9001
    auths:
      - "sha256:06004fd4f328fab028833d5156da66649be95afd61d41b690de33c1e3e3941a6" # suser1
      - pam
  - name: user1
    uid: 9002
    gid: 9002
    auths:
      - "sha256:0a041b9462caa4a31bac3567e0b6e6fd9100787db2ab433d96f6d178cabfce90" # user1
```

### Usage


```

go389 - A simple LDAP server

USAGE:
   go389 [global options] command [command options] [arguments...]

VERSION:
   0.1.0

AUTHOR:
   kernel164

COMMANDS:
     server   Server
     hash     hash
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

```

#### Examples

* You want to start serving ldap requests

`go389 server -c {{ APPCONFIG }}.yml` e.g.

`go389 server -c app.yml`

The above command will read in settings as defined in `app.yml`

The provided configuration references a yaml backend, a flat file named `db.yaml`

You can configure new users fairly easily.

To generate new sha256 passwords, you can utilize a free web service or do so via your choice of programmatic approach, e.g. python, go, bash, etc

Here's the website I used:

[http://passwordsgenerator.net/sha256-hash-generator/](http://passwordsgenerator.net/sha256-hash-generator/)

### References
- https://github.com/nmcclain/ldap
- https://github.com/nmcclain/glauth
