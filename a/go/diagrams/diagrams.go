package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/thomaspeugeot/testprobe/a/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration
var _ ref_models.Stage

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	const __write__local_time = "2025-08-23 17:08:07.095935 CEST"
	const __write__utc_time__ = "2025-08-23 15:08:07.095935 UTC"

	const __commitId__ = "0000000001"

	// Declaration of instances to stage

	__DiagramPackage__000000_Diagram_Package_created_the_2025_08_23T15_08_07Z := (&models.DiagramPackage{}).Stage(stage)

	// Setup of values

	__DiagramPackage__000000_Diagram_Package_created_the_2025_08_23T15_08_07Z.Name = `Diagram Package created the 2025-08-23T15:08:07Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_08_23T15_08_07Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_08_23T15_08_07Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_08_23T15_08_07Z.AbsolutePathToDiagramPackage = ``

	// Setup of pointers
	// setup of DiagramPackage instances pointers
}

