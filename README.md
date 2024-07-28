# xsnowflake

## Original godoc
https://pkg.go.dev/github.com/bwmarrin/snowflake

## Usage
1. When the system is initializing:
```go
nodeId := 0 // nodeId MUST be 0~1023
epoch := time.Parse(time.RFC3339, "2024-07-29T05:00:00Z") // Epoch MUST eariler than current time.
generator, err := xsnowflake.NewGenerator(nodeId, epoch)
```

2. Generate a new snowflake ID:
```go
id := generator.Generator()
```
