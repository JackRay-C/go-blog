package setting

type Snowflake struct {
	WorkID       int64 `mapstructure:"work-id"`
	DataCenterID int64 `mapstructure:"data-center-id"`
}
