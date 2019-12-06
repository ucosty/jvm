package jvm

import (
	"fmt"
	"java-hackery/pkg/classloader"
)

type opcode struct {
	Name    string
	Handler func(code []byte, f *Frame) error
}

var opcodes []opcode

func init() {
	opcodes = []opcode{
		{"nop", nop},                 //00
		{"aconst_null", unknown},     //01
		{"iconst_m1", unknown},       //02
		{"iconst_0", unknown},        //03
		{"iconst_1", unknown},        //04
		{"iconst_2", unknown},        //05
		{"iconst_3", unknown},        //06
		{"iconst_4", unknown},        //07
		{"iconst_5", unknown},        //08
		{"lconst_0", unknown},        //09
		{"lconst_1", unknown},        //0a
		{"fconst_0", unknown},        //0b
		{"fconst_1", unknown},        //0c
		{"fconst_2", unknown},        //0d
		{"dconst_0", unknown},        //0e
		{"dconst_1", unknown},        //0f
		{"bipush", unknown},          //10
		{"sipush", unknown},          //11
		{"ldc", ldc},                 //12
		{"ldc_w", unknown},           //13
		{"ldc2_w", unknown},          //14
		{"iload", unknown},           //15
		{"lload", unknown},           //16
		{"fload", unknown},           //17
		{"dload", unknown},           //18
		{"aload", unknown},           //19
		{"iload_0", unknown},         //1a
		{"iload_1", unknown},         //1b
		{"iload_2", unknown},         //1c
		{"iload_3", unknown},         //1d
		{"lload_0", unknown},         //1e
		{"lload_1", unknown},         //1f
		{"lload_2", unknown},         //20
		{"lload_3", unknown},         //21
		{"fload_0", unknown},         //22
		{"fload_1", unknown},         //23
		{"fload_2", unknown},         //24
		{"fload_3", unknown},         //25
		{"dload_0", unknown},         //26
		{"dload_1", unknown},         //27
		{"dload_2", unknown},         //28
		{"dload_3", unknown},         //29
		{"aload_0", unknown},         //2a
		{"aload_1", unknown},         //2b
		{"aload_2", unknown},         //2c
		{"aload_3", unknown},         //2d
		{"iaload", unknown},          //2e
		{"laload", unknown},          //2f
		{"faload", unknown},          //30
		{"daload", unknown},          //31
		{"aaload", unknown},          //32
		{"baload", unknown},          //33
		{"caload", unknown},          //34
		{"saload", unknown},          //35
		{"istore", unknown},          //36
		{"lstore", unknown},          //37
		{"fstore", unknown},          //38
		{"dstore", unknown},          //39
		{"astore", unknown},          //3a
		{"istore_0", unknown},        //3b
		{"istore_1", unknown},        //3c
		{"istore_2", unknown},        //3d
		{"istore_3", unknown},        //3e
		{"lstore_0", unknown},        //3f
		{"lstore_1", unknown},        //40
		{"lstore_2", unknown},        //41
		{"lstore_3", unknown},        //42
		{"fstore_0", unknown},        //43
		{"fstore_1", unknown},        //44
		{"fstore_2", unknown},        //45
		{"fstore_3", unknown},        //46
		{"dstore_0", unknown},        //47
		{"dstore_1", unknown},        //48
		{"dstore_2", unknown},        //49
		{"dstore_3", unknown},        //4a
		{"astore_0", unknown},        //4b
		{"astore_1", unknown},        //4c
		{"astore_2", unknown},        //4d
		{"astore_3", unknown},        //4e
		{"iastore", unknown},         //4f
		{"lastore", unknown},         //50
		{"fastore", unknown},         //51
		{"dastore", unknown},         //52
		{"aastore", unknown},         //53
		{"bastore", unknown},         //54
		{"castore", unknown},         //55
		{"sastore", unknown},         //56
		{"pop", unknown},             //57
		{"pop2", unknown},            //58
		{"dup", unknown},             //59
		{"dup_x1", unknown},          //5a
		{"dup_x2", unknown},          //5b
		{"dup2", unknown},            //5c
		{"dup2_x1", unknown},         //5d
		{"dup2_x2", unknown},         //5e
		{"swap", unknown},            //5f
		{"iadd", unknown},            //60
		{"ladd", unknown},            //61
		{"fadd", unknown},            //62
		{"dadd", unknown},            //63
		{"isub", unknown},            //64
		{"lsub", unknown},            //65
		{"fsub", unknown},            //66
		{"dsub", unknown},            //67
		{"imul", unknown},            //68
		{"lmul", unknown},            //69
		{"fmul", unknown},            //6a
		{"dmul", unknown},            //6b
		{"idiv", unknown},            //6c
		{"ldiv", unknown},            //6d
		{"fdiv", unknown},            //6e
		{"ddiv", unknown},            //6f
		{"irem", unknown},            //70
		{"lrem", unknown},            //71
		{"frem", unknown},            //72
		{"drem", unknown},            //73
		{"ineg", unknown},            //74
		{"lneg", unknown},            //75
		{"fneg", unknown},            //76
		{"dneg", unknown},            //77
		{"ishl", unknown},            //78
		{"lshl", unknown},            //79
		{"ishr", unknown},            //7a
		{"lshr", unknown},            //7b
		{"iushr", unknown},           //7c
		{"lushr", unknown},           //7d
		{"iand", unknown},            //7e
		{"land", unknown},            //7f
		{"ior", unknown},             //80
		{"lor", unknown},             //81
		{"ixor", unknown},            //82
		{"lxor", unknown},            //83
		{"iinc", unknown},            //84
		{"i2l", unknown},             //85
		{"i2f", unknown},             //86
		{"i2d", unknown},             //87
		{"l2i", unknown},             //88
		{"l2f", unknown},             //89
		{"l2d", unknown},             //8a
		{"f2i", unknown},             //8b
		{"f2l", unknown},             //8c
		{"f2d", unknown},             //8d
		{"d2i", unknown},             //8e
		{"d2l", unknown},             //8f
		{"d2f", unknown},             //90
		{"i2b", unknown},             //91
		{"i2c", unknown},             //92
		{"i2s", unknown},             //93
		{"lcmp", unknown},            //94
		{"fcmpl", unknown},           //95
		{"fcmpg", unknown},           //96
		{"dcmpl", unknown},           //97
		{"dcmpg", unknown},           //98
		{"ifeq", unknown},            //99
		{"ifne", unknown},            //9a
		{"iflt", unknown},            //9b
		{"ifge", unknown},            //9c
		{"ifgt", unknown},            //9d
		{"ifle", unknown},            //9e
		{"if_icmpeq", unknown},       //9f
		{"if_icmpne", unknown},       //a0
		{"if_icmplt", unknown},       //a1
		{"if_icmpge", unknown},       //a2
		{"if_icmpgt", unknown},       //a3
		{"if_icmple", unknown},       //a4
		{"if_acmpeq", unknown},       //a5
		{"if_acmpne", unknown},       //a6
		{"goto", unknown},            //a7
		{"jsr", unknown},             //a8
		{"ret", unknown},             //a9
		{"tableswitch", unknown},     //aa
		{"lookupswitch", unknown},    //ab
		{"ireturn", unknown},         //ac
		{"lreturn", unknown},         //ad
		{"freturn", unknown},         //ae
		{"dreturn", unknown},         //af
		{"areturn", unknown},         //b0
		{"return", unknown},          //b1
		{"getstatic", getStatic},     //b2
		{"putstatic", putStatic},     //b3
		{"getfield", unknown},        //b4
		{"putfield", unknown},        //b5
		{"invokevirtual", unknown},   //b6
		{"invokespecial", unknown},   //b7
		{"invokestatic", unknown},    //b8
		{"invokeinterface", unknown}, //b9
		{"invokedynamic", unknown},   //ba
		{"new", unknown},             //bb
		{"newarray", unknown},        //bc
		{"anewarray", unknown},       //bd
		{"arraylength", unknown},     //be
		{"athrow", unknown},          //bf
		{"checkcast", unknown},       //c0
		{"instanceof", unknown},      //c1
		{"monitorenter", unknown},    //c2
		{"monitorexit", unknown},     //c3
		{"wide", unknown},            //c4
		{"multianewarray", unknown},  //c5
		{"ifnull", unknown},          //c6
		{"ifnonnull", unknown},       //c7
		{"goto_w", unknown},          //c8
		{"jsr_w", unknown},           //c9
		{"breakpoint", unknown},      //ca
		{"impdep1", unknown},         //fe
		{"impdep2", unknown},         //ff
	}
}

func ldc(code []byte, f *Frame) error {
	index := code[1]
	entry := f.Class.Header.ConstantPoolTable[index-1]
	return f.PushOperand(entry)
}


func getStatic(code []byte, f *Frame) error {
	staticIndex := code[1]<<8 | code[2]
	_ = f.PushOperand(staticIndex)

	//f.Class.ConstantPool[staticIndex]

	return f.Step(3)
}

func putStatic(code []byte, f *Frame) error {
	staticIndex := code[1]<<8 | code[2]
	value, err := f.PopOperand()
	if err != nil {
		return err
	}

	constantEntry := f.Class.Header.ConstantPoolTable[staticIndex].(classloader.ConstantIntegerEntry)
	constantEntry.Bytes = value.(uint32)
	return f.Step(3)
}

func unknown(code []byte, f *Frame) error {
	instruction := code[0]
	return fmt.Errorf("unhandled opcode %x: %s", instruction, opcodes[instruction].Name)
}

func nop(_ []byte, f *Frame) error {
	return nil
}

// ---- Stack Frame ----

type Frame struct {
	Class        *classloader.JavaClass
	IP           uint32        // Instruction Pointer
	OperandStack []interface{} // Operand Stack
}

func (f *Frame) Step(count uint32) error {
	f.IP += count
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

func (f *Frame) DumpOperandStack() {
	for _, entry := range f.OperandStack {
		fmt.Printf("%T %v\n", entry, entry)
	}
}

func Invoke(c *classloader.JavaClass, method string, args...interface{}) error {
	code := c.Methods[method].AttributeInfo[0].(classloader.CodeAttribute).Code
	return Execute(code, c)
}

func Execute(code []byte, c *classloader.JavaClass) error {
	frame := &Frame{
		Class:        c,
		IP:           0,
		OperandStack: make([]interface{}, 0),
	}

	instruction := code[0]
	if int(instruction) > len(opcodes) {
		return fmt.Errorf("invalid opcode %x", instruction)
	}

	err := opcodes[instruction].Handler(code, frame)
	frame.DumpOperandStack()
	return err
}
