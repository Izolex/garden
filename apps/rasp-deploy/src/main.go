package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
	"gorm.io/gorm"
	"log"
	"os"
	"shared/app"
	"shared/model/database"
	"shared/model/entity/raspberry"
)

type Config struct {
	MySQLDsn   string
	SentryDsn  string
	ApiUrl     string
	ApiSignKey string
}

var ansiblePlaybookFile = flag.String("ansible", "/app/src/ansible.yaml", "ansible playbook file")
var raspBinFile = flag.String("bin", "/app/main", "compiled binary")
var dbFile = flag.String("db", "/app/db.sql", "DB file")

func main() {
	fmt.Println("Hi!")
	flag.Parse()

	checkFile(*ansiblePlaybookFile, "ansible playbook")
	checkFile(*raspBinFile, "raspberry bin")
	checkFile(*dbFile, "DB file")

	config := &Config{
		app.MustEnv("MYSQL_DSN"),
		app.MustEnv("SENTRY_DSN"),
		app.MustEnv("API_URL"),
		app.MustEnv("API_SIGN_KEY"),
	}

	db, sqlDb := database.NewMySQL(config.MySQLDsn)
	defer sqlDb.Close()

	raspberries := selectRaspberries(db)

	fmt.Printf("Going to deploy on %d raspberries\n", len(raspberries))

	inventory := newInventory(raspberries)
	err := newPlaybook(inventory, config).Run(context.TODO())

	if err == nil {
		fmt.Println("Success! Bye honey:*")
	} else {
		fmt.Println(err)
		fmt.Println("Fail! Repair it bro :/")
		os.Exit(1)
	}
}

func checkFile(path, name string) string {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		log.Fatalf("%s file \"%s\" does not exists", name, path)
	}
	return path
}

func selectRaspberries(db *gorm.DB) []*raspberry.Entity {
	var raspberries []*raspberry.Entity
	db.Where("is_active = 1").Find(&raspberries)
	return raspberries
}

func newInventory(raspberries []*raspberry.Entity) string {
	var inventory string
	for _, rasp := range raspberries {
		inventory += rasp.Address + ","
	}
	return inventory
}

func newPlaybook(inventory string, config *Config) *playbook.AnsiblePlaybookCmd {
	return &playbook.AnsiblePlaybookCmd{
		Playbooks: []string{*ansiblePlaybookFile},
		PrivilegeEscalationOptions: &options.AnsiblePrivilegeEscalationOptions{
			Become: true,
		},
		Options: &playbook.AnsiblePlaybookOptions{
			Inventory: inventory,
			ExtraVars: map[string]interface{}{
				"api_sign_key":   config.ApiSignKey,
				"api_url":        config.ApiUrl,
				"sentry_dsn":     config.SentryDsn,
				"service_bin":    *raspBinFile,
				"source_db_path": *dbFile,
			},
		},
	}
}
