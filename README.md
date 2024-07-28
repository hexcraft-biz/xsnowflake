# xsnowflake

## Usage
1. When the system is initializing:
```go
snowflakeNodeId := 0
snowflakeEpoch := time.Parse(time.RFC3339, "2024-07-29T05:00:00Z") // Epoch MUST eariler than current time.
generator, err := xsnowflake.NewGenerator(snowflakeNodeId, snowflakeEpoch)
```

2. Generate a new snowflake ID:
```go
id := generator.Generator()
```
