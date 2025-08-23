package main

import (
	"flag"
	"log"
	"strconv"

	// insertion point for models import
	b_models "github.com/thomaspeugeot/testprobe/b/go/models"
	b_stack "github.com/thomaspeugeot/testprobe/b/go/stack"

	// static
	split_static "github.com/fullstack-lang/gong/lib/split/go/static"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("b: ")
	log.SetFlags(log.Lmicroseconds)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := split_static.ServeStaticFiles(*logGINFlag)

	// setup model stack with its probe
	stack := b_stack.NewStack(r, "b", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	splitStage := split_stack.NewStack(r, "", "", "", "", false, false).Stage

	stager := b_models.NewStager(r, stack.Stage, splitStage)

	// one for the probe of the
	split.StageBranch(splitStage, &split.View{
		Name: stack.Stage.GetName() + "with Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Size: 50,
				AsSplit: (&split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						stager.GetAsSplitArea(),
					},
				}),
			}),
			(&split.AsSplitArea{
				Size: 50,
				Split: (&split.Split{
					StackName: stack.Stage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	// commit the split stage (this will initiate the front components)
	splitStage.Commit()

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
