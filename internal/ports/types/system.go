package types

// SystemConfig is used for DB migration
type SystemConfig struct {
	Name  string `json:"name" gorm:"primaryKey"` // nolint:tagalign
	Value string `json:"value"`
}

// TableName returns the table name for the SystemConfig model
func (SystemConfig) TableName() string {
	return "system_config"
}
