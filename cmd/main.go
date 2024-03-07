package main

import (
	"go.uber.org/fx"
	"kiramishima/credit_assigner/bootstrap"
)

func main() {
	fx.New(bootstrap.Module).Run()
}
