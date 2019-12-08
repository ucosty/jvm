package asm

type Opcode struct {
	Mnemonic string
	Length   int
	Operands []Operand
}

type Operand struct {
	Name string
	Type string
}

var Opcodes = []Opcode{
	{"nop", 1, []Operand{}},             //00
	{"aconst_null", 1, []Operand{}},     //01
	{"iconst_m1", 1, []Operand{}},       //02
	{"iconst_0", 1, []Operand{}},        //03
	{"iconst_1", 1, []Operand{}},        //04
	{"iconst_2", 1, []Operand{}},        //05
	{"iconst_3", 1, []Operand{}},        //06
	{"iconst_4", 1, []Operand{}},        //07
	{"iconst_5", 1, []Operand{}},        //08
	{"lconst_0", 1, []Operand{}},        //09
	{"lconst_1", 1, []Operand{}},        //0a
	{"fconst_0", 1, []Operand{}},        //0b
	{"fconst_1", 1, []Operand{}},        //0c
	{"fconst_2", 1, []Operand{}},        //0d
	{"dconst_0", 1, []Operand{}},        //0e
	{"dconst_1", 1, []Operand{}},        //0f
	{"bipush", 1, []Operand{}},          //10
	{"sipush", 1, []Operand{}},          //11
	{"ldc", 2, []Operand{}},             //12
	{"ldc_w", 1, []Operand{}},           //13
	{"ldc2_w", 1, []Operand{}},          //14
	{"iload", 1, []Operand{}},           //15
	{"lload", 1, []Operand{}},           //16
	{"fload", 1, []Operand{}},           //17
	{"dload", 1, []Operand{}},           //18
	{"aload", 1, []Operand{}},           //19
	{"iload_0", 1, []Operand{}},         //1a
	{"iload_1", 1, []Operand{}},         //1b
	{"iload_2", 1, []Operand{}},         //1c
	{"iload_3", 1, []Operand{}},         //1d
	{"lload_0", 1, []Operand{}},         //1e
	{"lload_1", 1, []Operand{}},         //1f
	{"lload_2", 1, []Operand{}},         //20
	{"lload_3", 1, []Operand{}},         //21
	{"fload_0", 1, []Operand{}},         //22
	{"fload_1", 1, []Operand{}},         //23
	{"fload_2", 1, []Operand{}},         //24
	{"fload_3", 1, []Operand{}},         //25
	{"dload_0", 1, []Operand{}},         //26
	{"dload_1", 1, []Operand{}},         //27
	{"dload_2", 1, []Operand{}},         //28
	{"dload_3", 1, []Operand{}},         //29
	{"aload_0", 1, []Operand{}},         //2a
	{"aload_1", 1, []Operand{}},         //2b
	{"aload_2", 1, []Operand{}},         //2c
	{"aload_3", 1, []Operand{}},         //2d
	{"iaload", 1, []Operand{}},          //2e
	{"laload", 1, []Operand{}},          //2f
	{"faload", 1, []Operand{}},          //30
	{"daload", 1, []Operand{}},          //31
	{"aaload", 1, []Operand{}},          //32
	{"baload", 1, []Operand{}},          //33
	{"caload", 1, []Operand{}},          //34
	{"saload", 1, []Operand{}},          //35
	{"istore", 1, []Operand{}},          //36
	{"lstore", 1, []Operand{}},          //37
	{"fstore", 1, []Operand{}},          //38
	{"dstore", 1, []Operand{}},          //39
	{"astore", 1, []Operand{}},          //3a
	{"istore_0", 1, []Operand{}},        //3b
	{"istore_1", 1, []Operand{}},        //3c
	{"istore_2", 1, []Operand{}},        //3d
	{"istore_3", 1, []Operand{}},        //3e
	{"lstore_0", 1, []Operand{}},        //3f
	{"lstore_1", 1, []Operand{}},        //40
	{"lstore_2", 1, []Operand{}},        //41
	{"lstore_3", 1, []Operand{}},        //42
	{"fstore_0", 1, []Operand{}},        //43
	{"fstore_1", 1, []Operand{}},        //44
	{"fstore_2", 1, []Operand{}},        //45
	{"fstore_3", 1, []Operand{}},        //46
	{"dstore_0", 1, []Operand{}},        //47
	{"dstore_1", 1, []Operand{}},        //48
	{"dstore_2", 1, []Operand{}},        //49
	{"dstore_3", 1, []Operand{}},        //4a
	{"astore_0", 1, []Operand{}},        //4b
	{"astore_1", 1, []Operand{}},        //4c
	{"astore_2", 1, []Operand{}},        //4d
	{"astore_3", 1, []Operand{}},        //4e
	{"iastore", 1, []Operand{}},         //4f
	{"lastore", 1, []Operand{}},         //50
	{"fastore", 1, []Operand{}},         //51
	{"dastore", 1, []Operand{}},         //52
	{"aastore", 1, []Operand{}},         //53
	{"bastore", 1, []Operand{}},         //54
	{"castore", 1, []Operand{}},         //55
	{"sastore", 1, []Operand{}},         //56
	{"pop", 1, []Operand{}},             //57
	{"pop2", 1, []Operand{}},            //58
	{"dup", 1, []Operand{}},             //59
	{"dup_x1", 1, []Operand{}},          //5a
	{"dup_x2", 1, []Operand{}},          //5b
	{"dup2", 1, []Operand{}},            //5c
	{"dup2_x1", 1, []Operand{}},         //5d
	{"dup2_x2", 1, []Operand{}},         //5e
	{"swap", 1, []Operand{}},            //5f
	{"iadd", 1, []Operand{}},            //60
	{"ladd", 1, []Operand{}},            //61
	{"fadd", 1, []Operand{}},            //62
	{"dadd", 1, []Operand{}},            //63
	{"isub", 1, []Operand{}},            //64
	{"lsub", 1, []Operand{}},            //65
	{"fsub", 1, []Operand{}},            //66
	{"dsub", 1, []Operand{}},            //67
	{"imul", 1, []Operand{}},            //68
	{"lmul", 1, []Operand{}},            //69
	{"fmul", 1, []Operand{}},            //6a
	{"dmul", 1, []Operand{}},            //6b
	{"idiv", 1, []Operand{}},            //6c
	{"ldiv", 1, []Operand{}},            //6d
	{"fdiv", 1, []Operand{}},            //6e
	{"ddiv", 1, []Operand{}},            //6f
	{"irem", 1, []Operand{}},            //70
	{"lrem", 1, []Operand{}},            //71
	{"frem", 1, []Operand{}},            //72
	{"drem", 1, []Operand{}},            //73
	{"ineg", 1, []Operand{}},            //74
	{"lneg", 1, []Operand{}},            //75
	{"fneg", 1, []Operand{}},            //76
	{"dneg", 1, []Operand{}},            //77
	{"ishl", 1, []Operand{}},            //78
	{"lshl", 1, []Operand{}},            //79
	{"ishr", 1, []Operand{}},            //7a
	{"lshr", 1, []Operand{}},            //7b
	{"iushr", 1, []Operand{}},           //7c
	{"lushr", 1, []Operand{}},           //7d
	{"iand", 1, []Operand{}},            //7e
	{"land", 1, []Operand{}},            //7f
	{"ior", 1, []Operand{}},             //80
	{"lor", 1, []Operand{}},             //81
	{"ixor", 1, []Operand{}},            //82
	{"lxor", 1, []Operand{}},            //83
	{"iinc", 1, []Operand{}},            //84
	{"i2l", 1, []Operand{}},             //85
	{"i2f", 1, []Operand{}},             //86
	{"i2d", 1, []Operand{}},             //87
	{"l2i", 1, []Operand{}},             //88
	{"l2f", 1, []Operand{}},             //89
	{"l2d", 1, []Operand{}},             //8a
	{"f2i", 1, []Operand{}},             //8b
	{"f2l", 1, []Operand{}},             //8c
	{"f2d", 1, []Operand{}},             //8d
	{"d2i", 1, []Operand{}},             //8e
	{"d2l", 1, []Operand{}},             //8f
	{"d2f", 1, []Operand{}},             //90
	{"i2b", 1, []Operand{}},             //91
	{"i2c", 1, []Operand{}},             //92
	{"i2s", 1, []Operand{}},             //93
	{"lcmp", 1, []Operand{}},            //94
	{"fcmpl", 1, []Operand{}},           //95
	{"fcmpg", 1, []Operand{}},           //96
	{"dcmpl", 1, []Operand{}},           //97
	{"dcmpg", 1, []Operand{}},           //98
	{"ifeq", 1, []Operand{}},            //99
	{"ifne", 1, []Operand{}},            //9a
	{"iflt", 1, []Operand{}},            //9b
	{"ifge", 1, []Operand{}},            //9c
	{"ifgt", 1, []Operand{}},            //9d
	{"ifle", 1, []Operand{}},            //9e
	{"if_icmpeq", 1, []Operand{}},       //9f
	{"if_icmpne", 1, []Operand{}},       //a0
	{"if_icmplt", 1, []Operand{}},       //a1
	{"if_icmpge", 1, []Operand{}},       //a2
	{"if_icmpgt", 1, []Operand{}},       //a3
	{"if_icmple", 1, []Operand{}},       //a4
	{"if_acmpeq", 1, []Operand{}},       //a5
	{"if_acmpne", 1, []Operand{}},       //a6
	{"goto", 1, []Operand{}},            //a7
	{"jsr", 1, []Operand{}},             //a8
	{"ret", 1, []Operand{}},             //a9
	{"tableswitch", 1, []Operand{}},     //aa
	{"lookupswitch", 1, []Operand{}},    //ab
	{"ireturn", 1, []Operand{}},         //ac
	{"lreturn", 1, []Operand{}},         //ad
	{"freturn", 1, []Operand{}},         //ae
	{"dreturn", 1, []Operand{}},         //af
	{"areturn", 1, []Operand{}},         //b0
	{"return", 1, []Operand{}},          //b1
	{"getstatic", 1, []Operand{}},       //b2
	{"putstatic", 1, []Operand{}},       //b3
	{"getfield", 1, []Operand{}},        //b4
	{"putfield", 1, []Operand{}},        //b5
	{"invokevirtual", 1, []Operand{}},   //b6
	{"invokespecial", 1, []Operand{}},   //b7
	{"invokestatic", 1, []Operand{}},    //b8
	{"invokeinterface", 1, []Operand{}}, //b9
	{"invokedynamic", 1, []Operand{}},   //ba
	{"new", 1, []Operand{}},             //bb
	{"newarray", 1, []Operand{}},        //bc
	{"anewarray", 1, []Operand{}},       //bd
	{"arraylength", 1, []Operand{}},     //be
	{"athrow", 1, []Operand{}},          //bf
	{"checkcast", 1, []Operand{}},       //c0
	{"instanceof", 1, []Operand{}},      //c1
	{"monitorenter", 1, []Operand{}},    //c2
	{"monitorexit", 1, []Operand{}},     //c3
	{"wide", 1, []Operand{}},            //c4
	{"multianewarray", 1, []Operand{}},  //c5
	{"ifnull", 1, []Operand{}},          //c6
	{"ifnonnull", 1, []Operand{}},       //c7
	{"goto_w", 1, []Operand{}},          //c8
	{"jsr_w", 1, []Operand{}},           //c9
	{"breakpoint", 1, []Operand{}},      //ca
	{"impdep1", 1, []Operand{}},         //fe
	{"impdep2", 1, []Operand{}},         //ff
}
