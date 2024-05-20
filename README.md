# Simple Token Service

A *super* simple token service to go along with this blog post: https://anthony-turcios.dev/posts/a-simple-token-service.

**Warning**: This should not be used in production without more controls in place. It is a good start for anyone looking
to implement API tokens though.

## Usage

You need Docker (and docker-compose) to run this!

```bash
./scripts/init.sh
```

Wait a second for postgres to start then build and run the go binary.

```bash
go build -o sts . && ./sts
```

Navigate to the "/" home page and click the "Get a Token" button. This should redirect you to a page with the token on it.

You can then use that same token to call a protected resource!

```bash
curl -v localhost:8000/api/hello \
 -H 'Authorization: Bearer YOUR_TOKEN_HERE'
```

Here is an example from my machine!

```bash
curl -v http://localhost:8000/api/hello \
 -H 'Authorization: Bearer 9QI1ZQS259OT0Z8fXgSz3aPPHH3SvLxzagLi4IX8'
```
