# wercker api client

[![wercker status](https://app.wercker.com/status/36462973eb9bc9dd7a0f3b9568328af6/s "wercker status")](https://app.wercker.com/project/bykey/36462973eb9bc9dd7a0f3b9568328af6)

werckerclient is the official wercker client for the
[wercker API][api-docs].

> Caution: Both the API and this client are under active development. This
> client may introduce breaking changes, so be sure to vendor in this client.

# Usage

To start using this client, import this package and create a new client:

```golang
import "github.com/wercker/werckerclient"

client := werckerclient.NewClient(nil)
```

This will create a new client with the default config. If you want to override
the default config, then you need to create a `wercker.Config` object and
pass this `wercker.NewClient`:

```golang
import "github.com/wercker/werckerclient"

options := &werckerclient.Config{}
client := werckerclient.NewClient(options)
```

## Authentication

The client uses a `Provider` to get the credentials for the user. By default it
uses the following strategy:

- Retrieve the token from the environment variable`$WERCKER_TOKEN`
- Retrieve the token from file `~/.wercker/credentials`
- Fallback to anonymous user

If you retrieved your wercker token through other means, then you can use the
`credentials.Token` method to create a static `Provider`:

```golang
import "github.com/wercker/werckerclient"
import "github.com/wercker/werckerclient/credentials"

token := "... your token ..."
options := &werckerclient.Config{Credentials: credentials.Token(token)}
client := werckerclient.NewClient(options)
```

[More information][auth-docs]

# FAQ

...

[api-docs]: http://devcenter.wercker.com/api/index.html
[auth-docs]: docs/authentication.md
