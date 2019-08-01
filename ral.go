package main

import (
	"fmt"
	"net/url"
	"flag"
	"os"

	"github.com/wesleycoakley/ral-api"
)

type CommandSet map[string]interface{}

var Commands = map[string]*flag.FlagSet {
	"view": flag.NewFlagSet("view", flag.ExitOnError) }

var CommandArgs = map[string]*CommandSet {
	"view": &ViewFlags }

var ViewFlags = CommandSet {
	"nowrap": Commands["view"].Bool(
		"nowrap",
		false,
		"Do not wrap text (only certain formats)"),
	"wrap": Commands["view"].Int(
		"wrap",
		80,
		"Wrap text to a certain number of characters (only certain formats)"),
	"format": Commands["view"].String(
		"format",
		"simple",
		"Command output format") }

var Flags = map[string]CommandSet {
	"view": ViewFlags }

var Formats = map[string]ral.Format {
	"simple": ral.FormatSimple,
	"csv": ral.FormatCSV,
	"json": ral.FormatJson }

func InitFlags() {
	for name, command := range Commands {
		fset := *CommandArgs[name]
		fset["config"] = command.String(
			"config",
			"",
			"Configuration file") } }

func main() {
	if len(os.Args) < 2 {
		GenericHelp()
		os.Exit(1) }

	InitFlags()
	flagset := Commands[os.Args[1]]
	argset := CommandArgs[os.Args[1]]
	flagset.Parse(os.Args[2:])

	configFile := *(*argset)["config"].(*string)
	if configFile == "" {
		configFile = FindConfig() }

	config, err := ReadConfig(configFile)
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
