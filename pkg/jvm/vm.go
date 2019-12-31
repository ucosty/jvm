package jvm

import (
	"fmt"
	strings "strings"
)

type opcode struct {
	Name    string
	Size    int
	Returns bool
	Handler func(code []byte, f *StackFrame) error
}

var opcodes []opcode

func init() {
	opcodes = []opcode{
		{"nop", 1, false, nop},                         //00
		{"aconst_null", 1, false, unknown},             //01
		{"iconst_m1", 1, false, unknown},               //02
		{"iconst_0", 1, false, unknown},                //03
		{"iconst_1", 1, false, unknown},                //04
		{"iconst_2", 1, false, unknown},                //05
		{"iconst_3", 1, false, unknown},                //06
		{"iconst_4", 1, false, unknown},                //07
		{"iconst_5", 1, false, unknown},                //08
		{"lconst_0", 1, false, unknown},                //09
		{"lconst_1", 1, false, unknown},                //0a
		{"fconst_0", 1, false, unknown},                //0b
		{"fconst_1", 1, false, unknown},                //0c
		{"fconst_2", 1, false, unknown},                //0d
		{"dconst_0", 1, false, unknown},                //0e
		{"dconst_1", 1, false, unknown},                //0f
		{"bipush", 1, false, bipush},                   //10
		{"sipush", 1, false, sipush},                   //11
		{"ldc", 1, false, ldc},                         //12
		{"ldc_w", 1, false, unknown},                   //13
		{"ldc2_w", 1, false, unknown},                  //14
		{"iload", 1, false, unknown},                   //15
		{"lload", 1, false, unknown},                   //16
		{"fload", 1, false, unknown},                   //17
		{"dload", 1, false, unknown},                   //18
		{"aload", 1, false, unknown},                   //19
		{"iload_0", 1, false, unknown},                 //1a
		{"iload_1", 1, false, unknown},                 //1b
		{"iload_2", 1, false, unknown},                 //1c
		{"iload_3", 1, false, unknown},                 //1d
		{"lload_0", 1, false, unknown},                 //1e
		{"lload_1", 1, false, unknown},                 //1f
		{"lload_2", 1, false, unknown},                 //20
		{"lload_3", 1, false, unknown},                 //21
		{"fload_0", 1, false, unknown},                 //22
		{"fload_1", 1, false, unknown},                 //23
		{"fload_2", 1, false, unknown},                 //24
		{"fload_3", 1, false, unknown},                 //25
		{"dload_0", 1, false, unknown},                 //26
		{"dload_1", 1, false, unknown},                 //27
		{"dload_2", 1, false, unknown},                 //28
		{"dload_3", 1, false, unknown},                 //29
		{"aload_0", 1, false, unknown},                 //2a
		{"aload_1", 1, false, unknown},                 //2b
		{"aload_2", 1, false, unknown},                 //2c
		{"aload_3", 1, false, unknown},                 //2d
		{"iaload", 1, false, unknown},                  //2e
		{"laload", 1, false, unknown},                  //2f
		{"faload", 1, false, unknown},                  //30
		{"daload", 1, false, unknown},                  //31
		{"aaload", 1, false, unknown},                  //32
		{"baload", 1, false, unknown},                  //33
		{"caload", 1, false, unknown},                  //34
		{"saload", 1, false, unknown},                  //35
		{"istore", 1, false, unknown},                  //36
		{"lstore", 1, false, unknown},                  //37
		{"fstore", 1, false, unknown},                  //38
		{"dstore", 1, false, unknown},                  //39
		{"astore", 1, false, unknown},                  //3a
		{"istore_0", 1, false, unknown},                //3b
		{"istore_1", 1, false, unknown},                //3c
		{"istore_2", 1, false, unknown},                //3d
		{"istore_3", 1, false, unknown},                //3e
		{"lstore_0", 1, false, unknown},                //3f
		{"lstore_1", 1, false, unknown},                //40
		{"lstore_2", 1, false, unknown},                //41
		{"lstore_3", 1, false, unknown},                //42
		{"fstore_0", 1, false, unknown},                //43
		{"fstore_1", 1, false, unknown},                //44
		{"fstore_2", 1, false, unknown},                //45
		{"fstore_3", 1, false, unknown},                //46
		{"dstore_0", 1, false, unknown},                //47
		{"dstore_1", 1, false, unknown},                //48
		{"dstore_2", 1, false, unknown},                //49
		{"dstore_3", 1, false, unknown},                //4a
		{"astore_0", 1, false, unknown},                //4b
		{"astore_1", 1, false, unknown},                //4c
		{"astore_2", 1, false, unknown},                //4d
		{"astore_3", 1, false, unknown},                //4e
		{"iastore", 1, false, unknown},                 //4f
		{"lastore", 1, false, unknown},                 //50
		{"fastore", 1, false, unknown},                 //51
		{"dastore", 1, false, unknown},                 //52
		{"aastore", 1, false, unknown},                 //53
		{"bastore", 1, false, unknown},                 //54
		{"castore", 1, false, unknown},                 //55
		{"sastore", 1, false, unknown},                 //56
		{"pop", 1, false, unknown},                     //57
		{"pop2", 1, false, unknown},                    //58
		{"dup", 1, false, unknown},                     //59
		{"dup_x1", 1, false, unknown},                  //5a
		{"dup_x2", 1, false, unknown},                  //5b
		{"dup2", 1, false, unknown},                    //5c
		{"dup2_x1", 1, false, unknown},                 //5d
		{"dup2_x2", 1, false, unknown},                 //5e
		{"swap", 1, false, unknown},                    //5f
		{"iadd", 1, false, unknown},                    //60
		{"ladd", 1, false, unknown},                    //61
		{"fadd", 1, false, unknown},                    //62
		{"dadd", 1, false, unknown},                    //63
		{"isub", 1, false, unknown},                    //64
		{"lsub", 1, false, unknown},                    //65
		{"fsub", 1, false, unknown},                    //66
		{"dsub", 1, false, unknown},                    //67
		{"imul", 1, false, unknown},                    //68
		{"lmul", 1, false, unknown},                    //69
		{"fmul", 1, false, unknown},                    //6a
		{"dmul", 1, false, unknown},                    //6b
		{"idiv", 1, false, unknown},                    //6c
		{"ldiv", 1, false, unknown},                    //6d
		{"fdiv", 1, false, unknown},                    //6e
		{"ddiv", 1, false, unknown},                    //6f
		{"irem", 1, false, unknown},                    //70
		{"lrem", 1, false, unknown},                    //71
		{"frem", 1, false, unknown},                    //72
		{"drem", 1, false, unknown},                    //73
		{"ineg", 1, false, unknown},                    //74
		{"lneg", 1, false, unknown},                    //75
		{"fneg", 1, false, unknown},                    //76
		{"dneg", 1, false, unknown},                    //77
		{"ishl", 1, false, unknown},                    //78
		{"lshl", 1, false, unknown},                    //79
		{"ishr", 1, false, unknown},                    //7a
		{"lshr", 1, false, unknown},                    //7b
		{"iushr", 1, false, unknown},                   //7c
		{"lushr", 1, false, unknown},                   //7d
		{"iand", 1, false, unknown},                    //7e
		{"land", 1, false, unknown},                    //7f
		{"ior", 1, false, unknown},                     //80
		{"lor", 1, false, unknown},                     //81
		{"ixor", 1, false, unknown},                    //82
		{"lxor", 1, false, unknown},                    //83
		{"iinc", 1, false, unknown},                    //84
		{"i2l", 1, false, unknown},                     //85
		{"i2f", 1, false, unknown},                     //86
		{"i2d", 1, false, unknown},                     //87
		{"l2i", 1, false, unknown},                     //88
		{"l2f", 1, false, unknown},                     //89
		{"l2d", 1, false, unknown},                     //8a
		{"f2i", 1, false, unknown},                     //8b
		{"f2l", 1, false, unknown},                     //8c
		{"f2d", 1, false, unknown},                     //8d
		{"d2i", 1, false, unknown},                     //8e
		{"d2l", 1, false, unknown},                     //8f
		{"d2f", 1, false, unknown},                     //90
		{"i2b", 1, false, unknown},                     //91
		{"i2c", 1, false, unknown},                     //92
		{"i2s", 1, false, unknown},                     //93
		{"lcmp", 1, false, unknown},                    //94
		{"fcmpl", 1, false, unknown},                   //95
		{"fcmpg", 1, false, unknown},                   //96
		{"dcmpl", 1, false, unknown},                   //97
		{"dcmpg", 1, false, unknown},                   //98
		{"ifeq", 1, false, unknown},                    //99
		{"ifne", 1, false, unknown},                    //9a
		{"iflt", 1, false, unknown},                    //9b
		{"ifge", 1, false, unknown},                    //9c
		{"ifgt", 1, false, unknown},                    //9d
		{"ifle", 1, false, unknown},                    //9e
		{"if_icmpeq", 1, false, unknown},               //9f
		{"if_icmpne", 1, false, unknown},               //a0
		{"if_icmplt", 1, false, unknown},               //a1
		{"if_icmpge", 1, false, unknown},               //a2
		{"if_icmpgt", 1, false, unknown},               //a3
		{"if_icmple", 1, false, unknown},               //a4
		{"if_acmpeq", 1, false, unknown},               //a5
		{"if_acmpne", 1, false, unknown},               //a6
		{"goto", 1, false, unknown},                    //a7
		{"jsr", 1, false, unknown},                     //a8
		{"ret", 1, false, unknown},                     //a9
		{"tableswitch", 1, false, unknown},             //aa
		{"lookupswitch", 1, false, unknown},            //ab
		{"ireturn", 1, true, ireturn},                 //ac
		{"lreturn", 1, true, lreturn},                 //ad
		{"freturn", 1, true, freturn},                 //ae
		{"dreturn", 1, true, dreturn},                 //af
		{"areturn", 1, true, areturn},                 //b0
		{"return", 1, true, vreturn},                  //b1
		{"getstatic", 1, false, getStatic},             //b2
		{"putstatic", 1, false, putStatic},             //b3
		{"getfield", 1, false, unknown},                //b4
		{"putfield", 1, false, unknown},                //b5
		{"invokevirtual", 1, false, invokevirtual},     //b6
		{"invokespecial", 1, false, invokespecial},     //b7
		{"invokestatic", 1, false, invokestatic},       //b8
		{"invokeinterface", 1, false, invokeinterface}, //b9
		{"invokedynamic", 1, false, invokedynamic},     //ba
		{"new", 1, false, unknown},                     //bb
		{"newarray", 1, false, unknown},                //bc
		{"anewarray", 1, false, unknown},               //bd
		{"arraylength", 1, false, unknown},             //be
		{"athrow", 1, false, unknown},                  //bf
		{"checkcast", 1, false, unknown},               //c0
		{"instanceof", 1, false, unknown},              //c1
		{"monitorenter", 1, false, unknown},            //c2
		{"monitorexit", 1, false, unknown},             //c3
		{"wide", 1, false, unknown},                    //c4
		{"multianewarray", 1, false, unknown},          //c5
		{"ifnull", 1, false, unknown},                  //c6
		{"ifnonnull", 1, false, unknown},               //c7
		{"goto_w", 1, false, unknown},                  //c8
		{"jsr_w", 1, false, unknown},                   //c9
		{"breakpoint", 1, false, unknown},              //ca
		{"impdep1", 1, false, unknown},                 //fe
		{"impdep2", 1, false, unknown},                 //ff
	}
}

/*
ldc: Push item from run-time constant pool
*/
func ldc(code []byte, f *StackFrame) error {
	index := code[f.IP+1]
	entry := f.Class.Constants[index]
	if entry.Instance == nil {
		return fmt.Errorf("constant %d instance is null", index)
	}
	if err := f.PushOperand(entry.Instance); err != nil {
		return err
	}
	return f.Step(2)
}

/*
Get static field from class
*/
func getStatic(code []byte, f *StackFrame) error {
	index := uint16(code[f.IP+1])<<8 | uint16(code[f.IP+2])
	if err := f.PushOperand(index); err != nil {
		return err
	}
	return f.Step(3)
}

// bipush: Push byte to operand stack
func bipush(code []byte, f *StackFrame) error {
	value := int(code[f.IP+1])
	if err := f.PushOperand(value); err != nil {
		return err
	}
	return f.Step(2)
}

// sipush: Push short to operand stack
func sipush(code []byte, f *StackFrame) error {
	value := int(code[f.IP+1]<<8 | code[f.IP+2])
	if err := f.PushOperand(value); err != nil {
		return err
	}
	return f.Step(3)
}

// invokedynamic
func invokedynamic(code []byte, f *StackFrame) error {
	fmt.Printf("Exec: invokedynamic\n")
	targetIndex := code[f.IP+1]<<8 | code[f.IP+2]
	targetConstant := f.Class.Constants[targetIndex]
	if targetConstant.Type != "CONSTANT_InvokeDynamic" {
		return fmt.Errorf("invokedynamic target unexpcted type, found %s", targetConstant.Type)
	}
	if code[f.IP+3] != 0 || code[f.IP+4] != 0 {
		return fmt.Errorf("non-zero byte found in invokedynamic instruction")
	}

	invokeDynamic := f.Class.Constants[targetIndex].Value.(ConstantInvokeDynamic)

	name, descriptor, err := f.Class.GetNameAndType(invokeDynamic.NameAndTypeIndex)
	if err != nil {
		return err
	}

	//invokeDynamic.BootstrapMethodAttrIndex

	fmt.Printf("%s %s\n", name, descriptor)

	//fmt.Printf()

	return f.Step(5)
}

func getInputDescriptor(descriptor string) string {
	inputStart := strings.Index(descriptor, "(")
	inputEnd := strings.Index(descriptor, ")")
	return descriptor[inputStart+1:inputEnd-1]
}

func getInvokeArgs(descriptor string, f *StackFrame) (args []interface{}, err error) {
	inputDescriptor := getInputDescriptor(descriptor)
	argumentTypes := strings.Split(inputDescriptor, ";")
	args = make([]interface{}, len(argumentTypes))

	for i, argType := range argumentTypes {
		fmt.Printf("Getting argument for type %s\n", argType)
		if args[i], err = f.PopOperand(); err != nil {
			return nil, err
		}
	}
	return args, nil
}

func getFieldref(index uint16, c *Class) error {
	fieldRef := c.Constants[index].Value.(ConstantFieldref)

	class := c.Constants[fieldRef.ClassIndex].Value.(ConstantClass)
	className, err := c.GetUTF8Constant(class.NameIndex)
	if err != nil {
		return err
	}

	nameAndType := c.Constants[fieldRef.NameAndTypeIndex].Value.(ConstantNameAndType)
	fieldName, err := c.GetUTF8Constant(nameAndType.NameIndex)
	if err != nil {
		return err
	}
	descriptor, err := c.GetUTF8Constant(nameAndType.DescriptorIndex)
	if err != nil {
		return err
	}

	fmt.Printf("Fieldref: className = %s, fieldName = %s, descriptor = %s\n", className, fieldName, descriptor)

	return nil
}

// invokevirtual
func invokevirtual(code []byte, f *StackFrame) error {
	fmt.Printf("Exec: invokevirtual\n")
	methodRefIndex := code[f.IP+1]<<8 | code[f.IP+2]
	methodRef := f.Class.Constants[methodRefIndex].Value.(ConstantMethodref)

	class := f.Class.Constants[methodRef.ClassIndex].Value.(ConstantClass)
	className, err := f.Class.GetUTF8Constant(class.NameIndex)
	if err != nil {
		return err
	}

	nameAndType := f.Class.Constants[methodRef.NameAndTypeIndex].Value.(ConstantNameAndType)
	name, err := f.Class.GetUTF8Constant(nameAndType.NameIndex)
	if err != nil {
		return err
	}
	descriptor, err := f.Class.GetUTF8Constant(nameAndType.DescriptorIndex)
	if err != nil {
		return err
	}

	// Get the method args
	args, err := getInvokeArgs(descriptor, f)
	if err != nil {
		return err
	}

	fmt.Printf("Got %d arguments\n", len(args))

	// Get the callee object reference
	objectRefIndex, err := f.PopOperand()
	if err != nil {
		return err
	}
	fmt.Printf("Got object ref: %d\n", objectRefIndex)

	_ = getFieldref(objectRefIndex.(uint16), f.Class)

	fmt.Printf("%s.%s%s\n", className, name, descriptor)

	return f.Step(3)
}

// invokespecial
func invokespecial(void []byte, f *StackFrame) error {
	fmt.Printf("Exec: invokespecial\n")
	return f.Step(1)
}

// invokestatic
func invokestatic(void []byte, f *StackFrame) error {
	fmt.Printf("Exec: invokestatic\n")
	return f.Step(1)
}

// invokeinterface
func invokeinterface(void []byte, f *StackFrame) error {
	fmt.Printf("Exec: invokeinterface\n")
	return f.Step(1)
}

// return: Return void from method
func vreturn(code []byte, f *StackFrame) error {
	f.ClearOperandStack()
	return f.Step(1)
}

// ireturn: Return int from method
func ireturn(void []byte, f *StackFrame) error {
	value, err := f.PopOperand()
	if err != nil {
		return err
	}
	// Ensure value is an integer, or otherwise try to convert value to integer if possible


	if err := f.Parent.PushOperand(value); err != nil {
		return err
	}
	return f.Step(1)
}

// lreturn: Return long from method
func lreturn(void []byte, f *StackFrame) error {
	value, err := f.PopOperand()
	if err != nil {
		return err
	}
	// Ensure value is a long, or otherwise try to convert value to a long if possible


	if err := f.Parent.PushOperand(value); err != nil {
		return err
	}
	return f.Step(1)
}

// freturn: Return float from method
func freturn(void []byte, f *StackFrame) error {
	value, err := f.PopOperand()
	if err != nil {
		return err
	}
	// Ensure value is a float, or otherwise try to convert value to a float if possible


	if err := f.Parent.PushOperand(value); err != nil {
		return err
	}
	return f.Step(1)
}

// dreturn: Return double from method
func dreturn(void []byte, f *StackFrame) error {
	value, err := f.PopOperand()
	if err != nil {
		return err
	}
	// Ensure value is a double, or otherwise try to convert value to a double if possible

	if err := f.Parent.PushOperand(value); err != nil {
		return err
	}
	return f.Step(1)
}

// areturn: Return reference from method
func areturn(void []byte, f *StackFrame) error {
	value, err := f.PopOperand()
	if err != nil {
		return err
	}
	// Ensure value is a reference

	if err := f.Parent.PushOperand(value); err != nil {
		return err
	}
	return f.Step(1)
}

/*
Set static field in class
*/
func putStatic(code []byte, f *StackFrame) error {
	targetIndex := code[f.IP+1]<<8 | code[f.IP+2]

	source, err := f.PopOperand()
	if err != nil {
		return err
	}
	target := f.Class.Constants[targetIndex]

	switch target.Type {
	case "CONSTANT_Fieldref":
		fieldref := target.Value.(ConstantFieldref)
		nameAndType := f.Class.Constants[fieldref.NameAndTypeIndex].Value.(ConstantNameAndType)
		fieldName, err := f.Class.GetUTF8Constant(nameAndType.NameIndex)
		if err != nil {
			return err
		}
		classType, err := f.Class.GetUTF8Constant(nameAndType.DescriptorIndex)
		if err != nil {
			return err
		}
		sourceInstance := source.(*ClassInstance)
		if classType != fmt.Sprintf("L%s;", sourceInstance.Class.Name) {
			return fmt.Errorf("could not assign %s.%s: expected type %s, found %s", f.Class.Name, fieldName, sourceInstance.Class.Name, classType)
		}
		target.Instance = sourceInstance
	}

	return f.Step(3)
}

func unknown(code []byte, f *StackFrame) error {
	instruction := code[f.IP]
	return fmt.Errorf("unhandled opcode %x: %s", instruction, opcodes[instruction].Name)
}

func nop(_ []byte, f *StackFrame) error {
	return nil
}

// ---- Stack StackFrame Operations ----

func (f *StackFrame) Step(count int) error {
	f.IP = f.IP + count
	return nil
}

func (f *StackFrame) PushOperand(value interface{}) error {
	f.OperandStack = append(f.OperandStack, value)
	return nil
}

func (f *StackFrame) PushOperands(values... interface{}) error {
	f.OperandStack = append(f.OperandStack, values)
	return nil
}

func (f *StackFrame) PopOperand() (value interface{}, err error) {
	value = f.OperandStack[len(f.OperandStack)-1]
	f.OperandStack = f.OperandStack[:len(f.OperandStack)-1]
	return value, nil
}

func (f *StackFrame) DumpFrame() {
	fmt.Printf("ip = %d\n", f.IP)
	f.DumpOperandStack()
}

func (f *StackFrame) ClearOperandStack() {
	f.OperandStack = nil
}

func (f *StackFrame) DumpOperandStack() {
	for idx, entry := range f.OperandStack {
		fmt.Printf("%d: %T %v\n", idx, entry, entry)
	}
}

func Execute(parent *StackFrame, code []byte, c *Class) error {
	frame := &StackFrame{
		Parent:       parent,
		Class:        c,
		IP:           0,
		OperandStack: make([]interface{}, 0),
	}

	_ = frame.PushOperand(nil)

	for frame.IP < len(code) {
		instruction := code[frame.IP]
		if int(instruction) > len(opcodes) {
			return fmt.Errorf("invalid opcode %x", instruction)
		}
		if err := opcodes[instruction].Handler(code, frame); err != nil {
			return err
		}
		if opcodes[instruction].Returns {
			fmt.Printf("return from function with %s\n", opcodes[instruction].Name)
			return nil
		}
	}
	return fmt.Errorf("ran out of code")
}
