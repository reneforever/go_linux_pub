package conf

type AppConf struct {
	KafkaConf
	TailConf
}

type KafkaConf struct {
	Address string
	Topic   string
}

type TailConf struct {
	FileName string
}
