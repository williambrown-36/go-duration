# Duration parse/format

```
duration.go
```
Check out the Go Duration test suite right next to the core logic to see how it handles edge cases in practice.

We often need to parse standard '1h30m' string formats into raw seconds and format them back. Doing this without pulling in heavy external packages keeps our eval harness fast and cheap. As Python developers building agent pipelines, we really appreciate tools that keep the dependency tree shallow.

This Go Duration utility relies strictly on the standard library. You do not have to install any extra services or third-party modules to get it running.

## Common questions

**Why skip the extra client library?**  
We simply do not need one. The network request is just a basic HTTPS call handled inside `duration.go`, with `go run .` acting as the sole tooling component. When you are just parsing time strings, keeping dependencies this minimal saves you a lot of headache during deployment.