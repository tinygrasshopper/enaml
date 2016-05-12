package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/xchapter7x/enaml/generators"
	"github.com/xchapter7x/enaml/pull"
	"github.com/xchapter7x/enaml/run"
)

func init() {
	if c := os.Getenv("ENAML_CACHE_DIR"); c != "" {
		cacheDir = c
	}

	if o := os.Getenv("ENAML_OUTPUT_DIR"); o != "" {
		outputDir = o
	}

	os.MkdirAll(cacheDir, 0755)
}

const (
	// CacheDir is the location downloaded releases are stored
	CacheDir = ".cache"
	// OutputDir is where the generate command outputs generated structs
	OutputDir = "./enaml-gen"
)

var (
	// Version is the enaml version
	Version     string
	cacheDir    = CacheDir
	outputDir   = OutputDir
	releaseRepo = pull.Release{CacheDir: cacheDir}
)

func main() {
	var err error
	app := cli.NewApp()
	app.Name = "enaml"
	app.Usage = "Because (EN)ough with the y(AML) already"
	app.Authors = []cli.Author{
		cli.Author{
			Name: "John Calabrese",
		},
		cli.Author{
			Name: "Caleb Washburn",
		},
		cli.Author{
			Name:  "Shawn Neal",
			Email: "sneal@sneal.net",
		},
	}
	app.Version = Version
	app.Commands = []cli.Command{
		{
			Name:      "generate",
			Usage:     "Generate golang structs for a given release",
			ArgsUsage: "<releaseURL>",

			Action: func(c *cli.Context) {
				generators.GenerateReleaseJobsPackage(c.Args().First(), cacheDir, OutputDir)
				println("completed generating release job structs for ", c.Args().First())
			},
		},
		{
			Name:      "diff",
			Usage:     "Show changes between two releases",
			ArgsUsage: "<release1URL> <release2URL>",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "job, j",
					Usage: "Focus the command to a specific named `JOB`",
				},
			},
			Action: func(c *cli.Context) {
				d := run.NewDiffCmd(releaseRepo, c.Args()[0], c.Args()[1])
				if len(c.String("job")) > 0 {
					err = d.Job(c.String("job"), os.Stdout)
				} else {
					err = d.All(os.Stdout)
				}
				ifErrorDisplayAndExit(err)
			},
		},
		{
			Name:      "show",
			Usage:     "Show details about the release",
			ArgsUsage: "<releaseURL>",
			Action: func(c *cli.Context) {
				s := run.NewShowCmd(releaseRepo, c.Args()[0])
				err := s.All(os.Stdout)
				ifErrorDisplayAndExit(err)
			},
		},
	}
	app.Run(os.Args)
}

func ifErrorDisplayAndExit(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
