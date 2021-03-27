package conf

func Mock() {
	Conf = &Config{
		Addr:                            "10001",
		Port:                            "10001",
		MysqlAmusinguserDatabase:        "amusinguser",
		MysqlAmusinguserUsername:        "amusing",
		MysqlAmusinguserPassword:        "amusing111",
		MysqlAmusinguserHost:            "localhost",
		MysqlAmusinguserPort:            "3306",
		MysqlAmusinguserProtocol:        "tcp",
		MysqlAmusinguserConnMaxLifetime: 2 * 60,
		MysqlAmusinguserMaxIdleConns:    10,
		MysqlAmusinguserMaxOpenConns:    10,
	}
}