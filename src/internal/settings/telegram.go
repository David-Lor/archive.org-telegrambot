package settings

type TelegramSettings struct {
	Bot struct {
		Token string `yaml:"token"`
	} `yaml:"bot"`
}
