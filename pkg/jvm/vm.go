package jvm

import (
	"fmt"
)

type opcode struct {
	Name    string
	Size    int
	Handler func(code []byte, f *Frame) error
}

var opcodes []opcode

func init() {
	opcodes = []opcode{
		{"nop", 1, nop},                         //00
		{"aconst_null", 1, unknown},             //01
		{"iconst_m1", 1, unknown},               //02
		{"iconst_0", 1, unknown},                //03
		{"iconst_1", 1, unknown},                //04
		{"iconst_2", 1, unknown},                //05
		{"iconst_3", 1, unknown},                //06
		{"iconst_4", 1, unknown},                //07
		{"iconst_5", 1, unknown},                //08
		{"lconst_0", 1, unknown},                //09
		{"lconst_1", 1, unknown},                //0a
		{"fconst_0", 1, unknown},                //0b
		{"fconst_1", 1, unknown},                //0c
		{"fconst_2", 1, unknown},                //0d
		{"dconst_0", 1, unknown},                //0e
		{"dconst_1", 1, unknown},                //0f
		{"bipush", 1, unknown},                  //10
		{"sipush", 1, unknown},                  //11
		{"ldc", 1, ldc},                         //12
		{"ldc_w", 1, unknown},                   //13
		{"ldc2_w", 1, unknown},                  //14
		{"iload", 1, unknown},                   //15
		{"lload", 1, unknown},                   //16
		{"fload", 1, unknown},                   //17
		{"dload", 1, unknown},                   //18
		{"aload", 1, unknown},                   //19
		{"iload_0", 1, unknown},                 //1a
		{"iload_1", 1, unknown},                 //1b
		{"iload_2", 1, unknown},                 //1c
		{"iload_3", 1, unknown},                 //1d
		{"lload_0", 1, unknown},                 //1e
		{"lload_1", 1, unknown},                 //1f
		{"lload_2", 1, unknown},                 //20
		{"lload_3", 1, unknown},                 //21
		{"fload_0", 1, unknown},                 //22
		{"fload_1", 1, unknown},                 //23
		{"fload_2", 1, unknown},                 //24
		{"fload_3", 1, unknown},                 //25
		{"dload_0", 1, unknown},                 //26
		{"dload_1", 1, unknown},                 //27
		{"dload_2", 1, unknown},                 //28
		{"dload_3", 1, unknown},                 //29
		{"aload_0", 1, unknown},                 //2a
		{"aload_1", 1, unknown},                 //2b
		{"aload_2", 1, unknown},                 //2c
		{"aload_3", 1, unknown},                 //2d
		{"iaload", 1, unknown},                  //2e
		{"laload", 1, unknown},                  //2f
		{"faload", 1, unknown},                  //30
		{"daload", 1, unknown},                  //31
		{"aaload", 1, unknown},                  //32
		{"baload", 1, unknown},                  //33
		{"caload", 1, unknown},                  //34
		{"saload", 1, unknown},                  //35
		{"istore", 1, unknown},                  //36
		{"lstore", 1, unknown},                  //37
		{"fstore", 1, unknown},                  //38
		{"dstore", 1, unknown},                  //39
		{"astore", 1, unknown},                  //3a
		{"istore_0", 1, unknown},                //3b
		{"istore_1", 1, unknown},                //3c
		{"istore_2", 1, unknown},                //3d
		{"istore_3", 1, unknown},                //3e
		{"lstore_0", 1, unknown},                //3f
		{"lstore_1", 1, unknown},                //40
		{"lstore_2", 1, unknown},                //41
		{"lstore_3", 1, unknown},                //42
		{"fstore_0", 1, unknown},                //43
		{"fstore_1", 1, unknown},                //44
		{"fstore_2", 1, unknown},                //45
		{"fstore_3", 1, unknown},                //46
		{"dstore_0", 1, unknown},                //47
		{"dstore_1", 1, unknown},                //48
		{"dstore_2", 1, unknown},                //49
		{"dstore_3", 1, unknown},                //4a
		{"astore_0", 1, unknown},                //4b
		{"astore_1", 1, unknown},                //4c
		{"astore_2", 1, unknown},                //4d
		{"astore_3", 1, unknown},                //4e
		{"iastore", 1, unknown},                 //4f
		{"lastore", 1, unknown},                 //50
		{"fastore", 1, unknown},                 //51
		{"dastore", 1, unknown},                 //52
		{"aastore", 1, unknown},                 //53
		{"bastore", 1, unknown},                 //54
		{"castore", 1, unknown},                 //55
		{"sastore", 1, unknown},                 //56
		{"pop", 1, unknown},                     //57
		{"pop2", 1, unknown},                    //58
		{"dup", 1, unknown},                     //59
		{"dup_x1", 1, unknown},                  //5a
		{"dup_x2", 1, unknown},                  //5b
		{"dup2", 1, unknown},                    //5c
		{"dup2_x1", 1, unknown},                 //5d
		{"dup2_x2", 1, unknown},                 //5e
		{"swap", 1, unknown},                    //5f
		{"iadd", 1, unknown},                    //60
		{"ladd", 1, unknown},                    //61
		{"fadd", 1, unknown},                    //62
		{"dadd", 1, unknown},                    //63
		{"isub", 1, unknown},                    //64
		{"lsub", 1, unknown},                    //65
		{"fsub", 1, unknown},                    //66
		{"dsub", 1, unknown},                    //67
		{"imul", 1, unknown},                    //68
		{"lmul", 1, unknown},                    //69
		{"fmul", 1, unknown},                    //6a
		{"dmul", 1, unknown},                    //6b
		{"idiv", 1, unknown},                    //6c
		{"ldiv", 1, unknown},                    //6d
		{"fdiv", 1, unknown},                    //6e
		{"ddiv", 1, unknown},                    //6f
		{"irem", 1, unknown},                    //70
		{"lrem", 1, unknown},                    //71
		{"frem", 1, unknown},                    //72
		{"drem", 1, unknown},                    //73
		{"ineg", 1, unknown},                    //74
		{"lneg", 1, unknown},                    //75
		{"fneg", 1, unknown},                    //76
		{"dneg", 1, unknown},                    //77
		{"ishl", 1, unknown},                    //78
		{"lshl", 1, unknown},                    //79
		{"ishr", 1, unknown},                    //7a
		{"lshr", 1, unknown},                    //7b
		{"iushr", 1, unknown},                   //7c
		{"lushr", 1, unknown},                   //7d
		{"iand", 1, unknown},                    //7e
		{"land", 1, unknown},                    //7f
		{"ior", 1, unknown},                     //80
		{"lor", 1, unknown},                     //81
		{"ixor", 1, unknown},                    //82
		{"lxor", 1, unknown},                    //83
		{"iinc", 1, unknown},                    //84
		{"i2l", 1, unknown},                     //85
		{"i2f", 1, unknown},                     //86
		{"i2d", 1, unknown},                     //87
		{"l2i", 1, unknown},                     //88
		{"l2f", 1, unknown},                     //89
		{"l2d", 1, unknown},                     //8a
		{"f2i", 1, unknown},                     //8b
		{"f2l", 1, unknown},                     //8c
		{"f2d", 1, unknown},                     //8d
		{"d2i", 1, unknown},                     //8e
		{"d2l", 1, unknown},                     //8f
		{"d2f", 1, unknown},                     //90
		{"i2b", 1, unknown},                     //91
		{"i2c", 1, unknown},                     //92
		{"i2s", 1, unknown},                     //93
		{"lcmp", 1, unknown},                    //94
		{"fcmpl", 1, unknown},                   //95
		{"fcmpg", 1, unknown},                   //96
		{"dcmpl", 1, unknown},                   //97
		{"dcmpg", 1, unknown},                   //98
		{"ifeq", 1, unknown},                    //99
		{"ifne", 1, unknown},                    //9a
		{"iflt", 1, unknown},                    //9b
		{"ifge", 1, unknown},                    //9c
		{"ifgt", 1, unknown},                    //9d
		{"ifle", 1, unknown},                    //9e
		{"if_icmpeq", 1, unknown},               //9f
		{"if_icmpne", 1, unknown},               //a0
		{"if_icmplt", 1, unknown},               //a1
		{"if_icmpge", 1, unknown},               //a2
		{"if_icmpgt", 1, unknown},               //a3
		{"if_icmple", 1, unknown},               //a4
		{"if_acmpeq", 1, unknown},               //a5
		{"if_acmpne", 1, unknown},               //a6
		{"goto", 1, unknown},                    //a7
		{"jsr", 1, unknown},                     //a8
		{"ret", 1, unknown},                     //a9
		{"tableswitch", 1, unknown},             //aa
		{"lookupswitch", 1, unknown},            //ab
		{"ireturn", 1, unknown},                 //ac
		{"lreturn", 1, unknown},                 //ad
		{"freturn", 1, unknown},                 //ae
		{"dreturn", 1, unknown},                 //af
		{"areturn", 1, unknown},                 //b0
		{"return", 1, returnVoid},               //b1
		{"getstatic", 1, getStatic},             //b2
		{"putstatic", 1, putStatic},             //b3
		{"getfield", 1, unknown},                //b4
		{"putfield", 1, unknown},                //b5
		{"invokevirtual", 1, invokevirtual},     //b6
		{"invokespecial", 1, invokespecial},     //b7
		{"invokestatic", 1, invokestatic},       //b8
		{"invokeinterface", 1, invokeinterface}, //b9
		{"invokedynamic", 1, invokedynamic},     //ba
		{"new", 1, unknown},                     //bb
		{"newarray", 1, unknown},                //bc
		{"anewarray", 1, unknown},               //bd
		{"arraylength", 1, unknown},             //be
		{"athrow", 1, unknown},                  //bf
		{"checkcast", 1, unknown},               //c0
		{"instanceof", 1, unknown},              //c1
		{"monitorenter", 1, unknown},            //c2
		{"monitorexit", 1, unknown},             //c3
		{"wide", 1, unknown},                    //c4
		{"multianewarray", 1, unknown},          //c5
		{"ifnull", 1, unknown},                  //c6
		{"ifnonnull", 1, unknown},               //c7
		{"goto_w", 1, unknown},                  //c8
		{"jsr_w", 1, unknown},                   //c9
		{"breakpoint", 1, unknown},              //ca
		{"impdep1", 1, unknown},                 //fe
		{"impdep2", 1, unknown},                 //ff
	}
}

/*
ldc: Push item from run-time constant pool
*/
func ldc(code []byte, f *Frame) error {
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
func getStatic(code []byte, f *Frame) error {
	index := code[f.IP+1]<<8 | code[f.IP+2]
	_ = f.PushOperand(index)
	return f.Step(3)
}

// Return void from method
func returnVoid(code []byte, f *Frame) error {
	f.ClearOperandStack()
	return f.Step(1)
}

// invokedynamic
func invokedynamic(code []byte, f *Frame) error {
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

// invokevirtual
func invokevirtual(code []byte, f *Frame) error {
	fmt.Printf("Exec: invokevirtual\n")
	methodRefIndex := code[f.IP+1]<<8 | code[f.IP+2]
	methodref := f.Class.Constants[methodRefIndex].Value.(ConstantMethodref)

	nameAndType := f.Class.Constants[methodref.NameAndTypeIndex].Value.(ConstantNameAndType)
	name, err := f.Class.GetUTF8Constant(nameAndType.NameIndex)
	if err != nil {
		return err
	}
	descriptor, err := f.Class.GetUTF8Constant(nameAndType.DescriptorIndex)
	if err != nil {
		return err
	}

	objectRefIndex, err := f.PopOperand()
	if err != nil {
		return err
	}
	fmt.Printf("Got object ref: %v\n", objectRefIndex)

	fmt.Printf("%s %s\n", name, descriptor)


	return f.Step(3)
}

// invokespecial
func invokespecial(void []byte, f *Frame) error {
	fmt.Printf("Exec: invokespecial\n")
	return f.Step(1)
}

// invokestatic
func invokestatic(void []byte, f *Frame) error {
	fmt.Printf("Exec: invokestatic\n")
	return f.Step(1)
}

// invokeinterface
func invokeinterface(void []byte, f *Frame) error {
	fmt.Printf("Exec: invokeinterface\n")
	return f.Step(1)
}

/*
Set static field in class
*/
func putStatic(code []byte, f *Frame) error {
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

func unknown(code []byte, f *Frame) error {
	instruction := code[f.IP]
	return fmt.Errorf("unhandled opcode %x: %s", instruction, opcodes[instruction].Name)
}

func nop(_ []byte, f *Frame) error {
	return nil
}

// ---- Stack Frame ----

type Frame struct {
	Class        *Class
	IP           int           // Instruction Pointer
	OperandStack []interface{} // Operand Stack
}

func (f *Frame) Step(count int) error {
	f.IP = f.IP + count
	return nil
}

func (f *Frame) PushOperand(value interface{}) error {
	f.OperandStack = append(f.OperandStack, value)
	return nil
}

func (f *Frame) PopOperand() (value interface{}, err error) {
	value = f.OperandStack[len(f.OperandStack)-1]
	f.OperandStack = f.OperandStack[:len(f.OperandStack)-1]
	return value, nil
}

func (f *Frame) DumpFrame() {
	fmt.Printf("ip = %d\n", f.IP)
	f.DumpOperandStack()
}

func (f *Frame) ClearOperandStack() {
	f.OperandStack = nil
}

func (f *Frame) DumpOperandStack() {
	for idx, entry := range f.OperandStack {
		fmt.Printf("%d: %T %v\n", idx, entry, entry)
	}
}

func Execute(code []byte, c *Class) error {
	frame := &Frame{
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
	}
	return nil
}
