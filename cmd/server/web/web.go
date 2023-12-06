package web

import (
	"embed"
)

//go:embed dist/*
var SpaFiles embed.FS
