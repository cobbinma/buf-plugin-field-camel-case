# buf-plugin-field-camel-case

🐪

## Usage

A plugin that implements the [Bufplugin API](https://buf.build/bufbuild/bufplugin).

Once you've installed the plugin, using `go install github.com/cobbinma/buf-plugin-field-camel-case` simply add a
reference to it and its rules within your `buf.yaml`.

```yaml
version: v2
lint:
  use:
    - PLUGIN_FIELD_CAMEL_CASE
plugins:
  - plugin: buf-plugin-field-camel-case
```

All [configuration](https://buf.build/docs/configuration/v2/buf-yaml) works as you'd expect: you can
continue to configure `use`, `except`, `ignore`, `ignore_only` and use `// buf:lint:ignore` comment
ignores, just as you would for the builtin rules.

Given the following file:

```protobuf
# foo.proto
syntax = "proto3";

package foo;

import "google/protobuf/timestamp.proto";

message Foo {
  google.protobuf.Timestamp snake_case = 1;
}
```

The following error will be returned from `buf lint`:

```
foo.proto:8:3:Field name snake_case should be camelCase, such as snakeCase. (buf-plugin-field-camel-case)
```
