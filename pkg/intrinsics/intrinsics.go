package intrinsics

import (
	"github.com/ucosty/jvm/pkg/jvm"
)

func Classes() []jvm.Class {
	return []jvm.Class{
		{Name: "java/lang/Object", Superclass: ""},
		{Name: "java/lang/String", Superclass: "java/lang/Object"},
		{Name: "java/lang/invoke/StringConcatFactory", Superclass: "java/lang/Object"},
		{Name: "java/lang/invoke/MethodType", Superclass: "java/lang/Object"},
		{Name: "java/lang/invoke/MethodHandles", Superclass: "java/lang/Object"},
		{Name: "java/lang/invoke/CallSite", Superclass: "java/lang/Object"},
	}
}
