package postgres

import (
	"strconv"

	entities "authentication/internal/entities"
	"github.com/D3vR4pt0rs/logger"

	"github.com/jackc/pgx"
)

type Config struct {
	Username string
	Password string
	Ip       string
	Port     string
	Database string
}

type dbClient struct {
	client *pgx.Conn
}

func New(cnfg Config) *dbClient {
	port, _ := strconv.Atoi(cnfg.Port)
	postgressConfig := pgx.ConnConfig{Host: cnfg.Ip, Port: uint16(port), User: cnfg.Username, Password: cnfg.Password, Database: cnfg.Database}
	conn, err := pgx.Connect(postgressConfig)
	if err != nil {
		logger.Error.Println(err.Error())
	}
	return &dbClient{
		client: conn,
	}
}

func (postgres *dbClient) GetAllProfile() ([]entities.Profile, error) {
	var profiles []entities.Profile
	rows, err := postgres.client.Query("SELECT * FROM profiles")
	if err != nil {
		logger.Error.Println("Error while executing query")
		return []entities.Profile{}, err
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			logger.Error.Println(err.Error())
			return []entities.Profile{}, err
		}
		profile := entities.Profile{
			ID: int(values[0].(int32)),
			Credentials: entities.Credentials{
				Email:    values[1].(string),
				Password: values[2].(string),
			},
			Balance: float64(values[3].(float32)),
		}
		profiles = append(profiles, profile)
	}
	return profiles, nil
}

func (postgres *dbClient) GetProfileByEmail(value string) (entities.Profile, error) {
	var profile Profile
	err := postgres.client.QueryRow("select * from profiles where email=$1", value).Scan(&profile.ID, &profile.Email, &profile.Password, &profile.Balance)
	if err != nil {
		logger.Error.Println(err.Error())
		return entities.Profile{}, err
	}

	return entities.Profile{
		Credentials: entities.Credentials{
			Email:    profile.Email,
			Password: profile.Password,
		},
		ID:      int(profile.ID),
		Balance: float64(profile.Balance),
	}, nil
}

func (postgres *dbClient) AddProfile(credentials entities.Credentials) error {
	_, err := postgres.client.Exec("insert into profiles (email,password) values ($1,$2)", credentials.Email, credentials.Password)
	if err != nil {
		logger.Error.Println(err.Error())
		return err
	}
	return nil
}
