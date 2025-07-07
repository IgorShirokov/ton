package app

func InitApp() error {
	if err := InitConfig(); err != nil {
		return err
	}
	return nil
}
