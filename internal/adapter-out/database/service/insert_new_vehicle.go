package service

func execute() (*string, error) {

	db, err := config.Connect()
	if err != nil {
		return nil, err
	}

}
