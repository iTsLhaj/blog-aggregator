package config

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl    string `json:"db_url"`
	Username string `json:"current_user_name"`
}
