// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
  BindAddr string
  FromAddr string
  ToAddr string
  SMTPServer string
  SMTPPort int
  SMTPUser string
  SMTPPassword string
)

func LoadRuntime() {
  if err := godotenv.Load(); err != nil {
    log.Fatal("Failed to load .env", err)
  }

  BindAddr = os.Getenv("BIND_ADDR") + ":" + os.Getenv("BIND_PORT")
  FromAddr = os.Getenv("FROM_EMAIL")
  ToAddr = os.Getenv("TO_EMAIL")
  SMTPServer = os.Getenv("SMTP_SERVER")
  var err error
  SMTPPort, err = strconv.Atoi(os.Getenv("SMTP_PORT"))
  if err != nil {
    log.Fatal("Failed to convert SMTP_PORT to integer")
  }
  SMTPUser = os.Getenv("SMTP_USER")
  SMTPPassword = os.Getenv("SMTP_PASSWORD")
}
