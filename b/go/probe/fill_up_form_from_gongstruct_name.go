// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/thomaspeugeot/testprobe/b/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "B":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "B Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BFormCallback(
			nil,
			probe,
			formGroup,
		)
		b := new(models.B)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(b, formGroup, probe)
	}
	formStage.Commit()
}
