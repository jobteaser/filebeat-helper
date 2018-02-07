filebeat-helper is a small utility to help you create pattern to handle multiline

More info on elastic.co documentation https://www.elastic.co/guide/en/beats/filebeat/current/multiline-examples.html

# Build

```
make build.linux
```

# Run the example

examples/elastic-1.log provides some stack trace to validate against your pattern

```
dist/filebeat-helper-linux -logfile examples/elastic-1.log -pattern '^\[\d{4}-\d{2}-\d{2}[^]]*\]' -negate
```
