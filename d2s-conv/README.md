# Diablo 2 Save editor CLI

CLI for convert .d2s to JSON and vice-versa

## Installation

To install command line program, use the following:

```bash
go install github.com/vitalick/d2s/d2s-conv@latest
```

## Usage
### CLI

For convert JSON to .d2s, use the following:
```bash
d2s-conv -fromjson <input files>
```

For convert .d2s to JSON, use the following:
```bash
d2s-conv -tojson <input files>
```