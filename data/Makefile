
.PHONY: all
all: classpath/java/lang/System.class classpath/java/lang/String.class

classpath/java/lang/System.class:
	javac --patch-module java.base=. classpath/java/lang/System.java

classpath/java/lang/String.class:
	javac --patch-module java.base=. classpath/java/lang/String.java
