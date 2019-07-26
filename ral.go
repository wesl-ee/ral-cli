package main

import (
	"fmt"
	"net/url"
	"flag"
	"os"

	"github.com/wesleycoakley/ral-api"
)

type CommandSet map[string]*string

var Commands = map[string]*flag.FlagSet {
	"view": flag.NewFlagSet("view", flag.ExitOnError)}

var ViewFlags = CommandSet {
	"format": Commands["view"].String(
		"format",
		"simple",
		"Command output format") }

var GlobalFlags = CommandSet {
	"config": flag.String(
		"config",
		"config.toml",
		"Configuration file") }

var Flags = map[string]CommandSet {
	"view": ViewFlags }

var Formats = map[string]ral.Format {
	"simple": ral.FormatSimple,
	"json": ral.FormatJson }

func main() {
	if len(os.Args) < 2 {
		GenericHelp()
		os.Exit(1) }

	flag.Parse()
	switch(os.Args[1]) {
		case "view":
			Commands["view"].Parse(os.Args[2:]) }

	config, err := ReadConfig("config.toml")
	if err != nil { panic(err) }

	s := ral.New()
	s.URL = url.URL{
		Scheme: config.Scheme,
		Path: config.Endpoint }

	if Commands["view"].Parsed() {
		View(s, ViewFlags, Commands["view"].Args()) } }

func GenericHelp() {
	fmt.Print(
`ral - Barebones CLI to read / post on an RAL Textboard
usage: ral [command] [arguments]

Commands:
	view - View a Continuity, Year, or Topic
`) }