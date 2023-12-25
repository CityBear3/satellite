package shcema

type ClientSchema struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	Secrets   string `db:"secrets"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

type ClientWithDevicesSchema struct {
	ClientSchema
	Device DeviceSchema `db:"device"`
}
