package intrinsics

import (
	"fmt"
	"github.com/ucosty/jvm/pkg/jvm"
)

func PatchSystem(metaspace *jvm.Metaspace) error {
	class, err := metaspace.GetClass("java/lang/System")
	if err != nil {
		return err
	}
	fmt.Printf("PatchSystem:\n%v\n---------\n", class.Fields[0].AccessFlags)


	return nil
}