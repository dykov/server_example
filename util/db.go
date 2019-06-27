package util

import (
	"time"

	"../config"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

func ConnectDatabase(c *config.Config) (*mgo.Database, error) {

	info := &mgo.DialInfo{
		Addrs:    []string{c.Database.Host},
		Timeout:  10 * time.Second,
		Database: c.Database.DatabaseName,
		Username: c.Database.Username,
		Password: c.Database.Password,
	}

	session, err := mgo.DialWithInfo(info)

	if err != nil {
		log.Errorln(DatabaseConnectionFailed)
		log.WithFields(log.Fields{
			"Address":  info.Addrs[0],
			"Timeout":  info.Timeout,
			"Database": info.Database,
			"Username": info.Username,
			"Password": info.Password,
			"Error":    err.Error(),
		}).Debugln(DatabaseConnectionFailed)
		session = nil
	} else {
		log.WithFields(log.Fields{
			"Address":  info.Addrs[0],
			"Timeout":  info.Timeout,
			"Database": info.Database,
			"Username": info.Username,
			"Password": info.Password,
		}).Infoln(SuccessfulDatabaseConnection)

	}

	return session.DB(c.Database.DatabaseName), err

}
