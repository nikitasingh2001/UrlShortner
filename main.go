package main

import (
	"context"
	"flag"
	"urlshortner/internal/bootstrap"
	"urlshortner/internal/constant"
)

var (
	basePath *string
	env      *string
)

func init() {
	basePath = flag.String(constant.BasePath, constant.DefaultBasePath, "Path to base path")
	env = flag.String(constant.Env, constant.Development, "Application env dev")
}

func main() {

	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bootstrap.Initialize(ctx, *basePath, *env)
}
