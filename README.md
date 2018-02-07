# FileBeat Helper

This library is a small utility to help you create pattern to handle multiline.

> **INFO:** documentation https://www.elastic.co/guide/en/beats/filebeat/current/multiline-examples.html

# Build

```sh
# For Linux
make build.linux

# For MacOS
make build.darwin
```

# Run the example

examples/elastic-1.log provides some stack trace to validate against your pattern

```sh
dist/filebeat-helper-linux -logfile examples/elastic-1.log -pattern '^\[\d{4}-\d{2}-\d{2}[^]]*\]' -negate
```
