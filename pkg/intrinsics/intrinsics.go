package intrinsics

import (
	"java-hackery/pkg/jvm"
)

func Classes() []jvm.Class {
	return []jvm.Class{
		{Name: "java/lang/Object", Superclass: ""},
		{Name: "java/lang/String", Superclass: "java/lang/Object"},
	}
}
