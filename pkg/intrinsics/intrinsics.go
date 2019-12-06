package intrinsics

import "java-hackery/pkg/classloader"

func Classes() []classloader.JavaClass {
	return []classloader.JavaClass{
		{Name: "java/lang/Object", Superclass: ""},
		{Name: "java/lang/String", Superclass: "java/lang/Object"},
	}
}
