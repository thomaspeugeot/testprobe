// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/thomaspeugeot/testprobe/b/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__BFormCallback(
	b *models.B,
	probe *Probe,
	formGroup *table.FormGroup,
) (bFormCallback *BFormCallback) {
	bFormCallback = new(BFormCallback)
	bFormCallback.probe = probe
	bFormCallback.b = b
	bFormCallback.formGroup = formGroup

	bFormCallback.CreationMode = (b == nil)

	return
}

type BFormCallback struct {
	b *models.B

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (bFormCallback *BFormCallback) OnSave() {

	// log.Println("BFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bFormCallback.probe.formStage.Checkout()

	if bFormCallback.b == nil {
		bFormCallback.b = new(models.B).Stage(bFormCallback.probe.stageOfInterest)
	}
	b_ := bFormCallback.b
	_ = b_

	for _, formDiv := range bFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(b_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if bFormCallback.formGroup.HasSuppressButtonBeenPressed {
		b_.Unstage(bFormCallback.probe.stageOfInterest)
	}

	bFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.B](
		bFormCallback.probe,
	)
	bFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if bFormCallback.CreationMode || bFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(bFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BFormCallback(
			nil,
			bFormCallback.probe,
			newFormGroup,
		)
		b := new(models.B)
		FillUpForm(b, newFormGroup, bFormCallback.probe)
		bFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(bFormCallback.probe)
}
