package cache

type Cache struct {
	LastTimeSync LastTimeSync `yaml:"last_time_sync"`
}

type LastTimeSync struct {
	C2COrders         uint64 `yaml:"c2c_orders"`
	SpotOrders        uint64 `yaml:"spot_orders"`
	FiatDepositOrders uint64 `yaml:"fiat_deposit_orders"`
}
