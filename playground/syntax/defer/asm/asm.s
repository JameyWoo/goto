"".returnButDefer STEXT size=146 args=0x8 locals=0x60
	0x0000 00000 (main.go:5)	TEXT	"".returnButDefer(SB), ABIInternal, $96-8
	0x0000 00000 (main.go:5)	MOVQ	TLS, CX
	0x0009 00009 (main.go:5)	PCDATA	$0, $-2
	0x0009 00009 (main.go:5)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (main.go:5)	PCDATA	$0, $-1
	0x0010 00016 (main.go:5)	CMPQ	SP, 16(CX)
	0x0014 00020 (main.go:5)	PCDATA	$0, $-2
	0x0014 00020 (main.go:5)	JLS	136
	0x0016 00022 (main.go:5)	PCDATA	$0, $-1
	0x0016 00022 (main.go:5)	SUBQ	$96, SP
	0x001a 00026 (main.go:5)	MOVQ	BP, 88(SP)
	0x001f 00031 (main.go:5)	LEAQ	88(SP), BP
	0x0024 00036 (main.go:5)	PCDATA	$0, $-2
	0x0024 00036 (main.go:5)	PCDATA	$1, $-2
	0x0024 00036 (main.go:5)	FUNCDATA	$0, gclocals路33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0024 00036 (main.go:5)	FUNCDATA	$1, gclocals路33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0024 00036 (main.go:5)	FUNCDATA	$2, gclocals路9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x0024 00036 (main.go:5)	PCDATA	$0, $0
	0x0024 00036 (main.go:5)	PCDATA	$1, $0
	0x0024 00036 (main.go:5)	MOVQ	$0, "".it+104(SP)
	0x002d 00045 (main.go:7)	MOVL	$8, ""..autotmp_2+8(SP)
	0x0035 00053 (main.go:7)	PCDATA	$0, $1
	0x0035 00053 (main.go:7)	LEAQ	"".returnButDefer.func1路f(SB), AX
	0x003c 00060 (main.go:7)	PCDATA	$0, $0
	0x003c 00060 (main.go:7)	MOVQ	AX, ""..autotmp_2+32(SP)
	0x0041 00065 (main.go:7)	PCDATA	$0, $1
	0x0041 00065 (main.go:7)	LEAQ	"".it+104(SP), AX
	0x0046 00070 (main.go:7)	PCDATA	$0, $0
	0x0046 00070 (main.go:7)	MOVQ	AX, ""..autotmp_2+80(SP)
	0x004b 00075 (main.go:7)	PCDATA	$0, $1
	0x004b 00075 (main.go:7)	LEAQ	""..autotmp_2+8(SP), AX
	0x0050 00080 (main.go:7)	PCDATA	$0, $0
	0x0050 00080 (main.go:7)	MOVQ	AX, (SP)
	0x0054 00084 (main.go:7)	CALL	runtime.deferprocStack(SB)
	0x0059 00089 (main.go:7)	TESTL	AX, AX
	0x005b 00091 (main.go:7)	JNE	120
	0x005d 00093 (main.go:7)	JMP	95
	0x005f 00095 (main.go:11)	MOVQ	$891, "".it+104(SP)
	0x0068 00104 (main.go:11)	XCHGL	AX, AX
	0x0069 00105 (main.go:11)	CALL	runtime.deferreturn(SB)
	0x006e 00110 (main.go:11)	MOVQ	88(SP), BP
	0x0073 00115 (main.go:11)	ADDQ	$96, SP
	0x0077 00119 (main.go:11)	RET
	0x0078 00120 (main.go:7)	XCHGL	AX, AX
	0x0079 00121 (main.go:7)	CALL	runtime.deferreturn(SB)
	0x007e 00126 (main.go:7)	MOVQ	88(SP), BP
	0x0083 00131 (main.go:7)	ADDQ	$96, SP
	0x0087 00135 (main.go:7)	RET
	0x0088 00136 (main.go:7)	NOP
	0x0088 00136 (main.go:5)	PCDATA	$1, $-1
	0x0088 00136 (main.go:5)	PCDATA	$0, $-2
	0x0088 00136 (main.go:5)	CALL	runtime.morestack_noctxt(SB)
	0x008d 00141 (main.go:5)	PCDATA	$0, $-1
	0x008d 00141 (main.go:5)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 3b 61 10 76 72 48 83 ec 60 48 89 6c 24 58 48  H;a.vrH..`H.l$XH
	0x0020 8d 6c 24 58 48 c7 44 24 68 00 00 00 00 c7 44 24  .l$XH.D$h.....D$
	0x0030 08 08 00 00 00 48 8d 05 00 00 00 00 48 89 44 24  .....H......H.D$
	0x0040 20 48 8d 44 24 68 48 89 44 24 50 48 8d 44 24 08   H.D$hH.D$PH.D$.
	0x0050 48 89 04 24 e8 00 00 00 00 85 c0 75 1b eb 00 48  H..$.......u...H
	0x0060 c7 44 24 68 7b 03 00 00 90 e8 00 00 00 00 48 8b  .D$h{.........H.
	0x0070 6c 24 58 48 83 c4 60 c3 90 e8 00 00 00 00 48 8b  l$XH..`.......H.
	0x0080 6c 24 58 48 83 c4 60 c3 e8 00 00 00 00 e9 6e ff  l$XH..`.......n.
	0x0090 ff ff                                            ..
	rel 12+4 t=17 TLS+0
	rel 56+4 t=16 "".returnButDefer.func1路f+0
	rel 85+4 t=8 runtime.deferprocStack+0
	rel 106+4 t=8 runtime.deferreturn+0
	rel 122+4 t=8 runtime.deferreturn+0
	rel 137+4 t=8 runtime.morestack_noctxt+0
"".main STEXT size=198 args=0x0 locals=0x78
	0x0000 00000 (main.go:14)	TEXT	"".main(SB), ABIInternal, $120-0
	0x0000 00000 (main.go:14)	MOVQ	TLS, CX
	0x0009 00009 (main.go:14)	PCDATA	$0, $-2
	0x0009 00009 (main.go:14)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (main.go:14)	PCDATA	$0, $-1
	0x0010 00016 (main.go:14)	CMPQ	SP, 16(CX)
	0x0014 00020 (main.go:14)	PCDATA	$0, $-2
	0x0014 00020 (main.go:14)	JLS	188
	0x001a 00026 (main.go:14)	PCDATA	$0, $-1
	0x001a 00026 (main.go:14)	SUBQ	$120, SP
	0x001e 00030 (main.go:14)	MOVQ	BP, 112(SP)
	0x0023 00035 (main.go:14)	LEAQ	112(SP), BP
	0x0028 00040 (main.go:14)	PCDATA	$0, $-2
	0x0028 00040 (main.go:14)	PCDATA	$1, $-2
	0x0028 00040 (main.go:14)	FUNCDATA	$0, gclocals路7d2d5fca80364273fb07d5820a76fef4(SB)
	0x0028 00040 (main.go:14)	FUNCDATA	$1, gclocals路a8c779e02c34a144ad683cd0297269a5(SB)
	0x0028 00040 (main.go:14)	FUNCDATA	$2, gclocals路f6aec3988379d2bd21c69c093370a150(SB)
	0x0028 00040 (main.go:14)	FUNCDATA	$3, "".main.stkobj(SB)
	0x0028 00040 (main.go:15)	PCDATA	$0, $0
	0x0028 00040 (main.go:15)	PCDATA	$1, $0
	0x0028 00040 (main.go:15)	MOVQ	$321, "".a+48(SP)
	0x0031 00049 (main.go:16)	CALL	"".returnButDefer(SB)
	0x0036 00054 (main.go:16)	MOVQ	(SP), AX
	0x003a 00058 (main.go:16)	MOVQ	AX, "".a+48(SP)
	0x003f 00063 (main.go:17)	MOVQ	AX, (SP)
	0x0043 00067 (main.go:17)	CALL	runtime.convT64(SB)
	0x0048 00072 (main.go:17)	PCDATA	$0, $1
	0x0048 00072 (main.go:17)	MOVQ	8(SP), AX
	0x004d 00077 (main.go:17)	PCDATA	$0, $0
	0x004d 00077 (main.go:17)	PCDATA	$1, $1
	0x004d 00077 (main.go:17)	MOVQ	AX, ""..autotmp_2+64(SP)
	0x0052 00082 (main.go:17)	PCDATA	$1, $2
	0x0052 00082 (main.go:17)	XORPS	X0, X0
	0x0055 00085 (main.go:17)	MOVUPS	X0, ""..autotmp_1+72(SP)
	0x005a 00090 (main.go:17)	PCDATA	$0, $1
	0x005a 00090 (main.go:17)	PCDATA	$1, $1
	0x005a 00090 (main.go:17)	LEAQ	""..autotmp_1+72(SP), AX
	0x005f 00095 (main.go:17)	MOVQ	AX, ""..autotmp_4+56(SP)
	0x0064 00100 (main.go:17)	TESTB	AL, (AX)
	0x0066 00102 (main.go:17)	PCDATA	$0, $2
	0x0066 00102 (main.go:17)	PCDATA	$1, $0
	0x0066 00102 (main.go:17)	MOVQ	""..autotmp_2+64(SP), CX
	0x006b 00107 (main.go:17)	PCDATA	$0, $3
	0x006b 00107 (main.go:17)	LEAQ	type.int(SB), DX
	0x0072 00114 (main.go:17)	PCDATA	$0, $2
	0x0072 00114 (main.go:17)	MOVQ	DX, ""..autotmp_1+72(SP)
	0x0077 00119 (main.go:17)	PCDATA	$0, $1
	0x0077 00119 (main.go:17)	MOVQ	CX, ""..autotmp_1+80(SP)
	0x007c 00124 (main.go:17)	TESTB	AL, (AX)
	0x007e 00126 (main.go:17)	JMP	128
	0x0080 00128 (main.go:17)	MOVQ	AX, ""..autotmp_3+88(SP)
	0x0085 00133 (main.go:17)	MOVQ	$1, ""..autotmp_3+96(SP)
	0x008e 00142 (main.go:17)	MOVQ	$1, ""..autotmp_3+104(SP)
	0x0097 00151 (main.go:17)	PCDATA	$0, $0
	0x0097 00151 (main.go:17)	MOVQ	AX, (SP)
	0x009b 00155 (main.go:17)	MOVQ	$1, 8(SP)
	0x00a4 00164 (main.go:17)	MOVQ	$1, 16(SP)
	0x00ad 00173 (main.go:17)	CALL	fmt.Println(SB)
	0x00b2 00178 (main.go:18)	MOVQ	112(SP), BP
	0x00b7 00183 (main.go:18)	ADDQ	$120, SP
	0x00bb 00187 (main.go:18)	RET
	0x00bc 00188 (main.go:18)	NOP
	0x00bc 00188 (main.go:14)	PCDATA	$1, $-1
	0x00bc 00188 (main.go:14)	PCDATA	$0, $-2
	0x00bc 00188 (main.go:14)	CALL	runtime.morestack_noctxt(SB)
	0x00c1 00193 (main.go:14)	PCDATA	$0, $-1
	0x00c1 00193 (main.go:14)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 3b 61 10 0f 86 a2 00 00 00 48 83 ec 78 48 89  H;a.......H..xH.
	0x0020 6c 24 70 48 8d 6c 24 70 48 c7 44 24 30 41 01 00  l$pH.l$pH.D$0A..
	0x0030 00 e8 00 00 00 00 48 8b 04 24 48 89 44 24 30 48  ......H..$H.D$0H
	0x0040 89 04 24 e8 00 00 00 00 48 8b 44 24 08 48 89 44  ..$.....H.D$.H.D
	0x0050 24 40 0f 57 c0 0f 11 44 24 48 48 8d 44 24 48 48  $@.W...D$HH.D$HH
	0x0060 89 44 24 38 84 00 48 8b 4c 24 40 48 8d 15 00 00  .D$8..H.L$@H....
	0x0070 00 00 48 89 54 24 48 48 89 4c 24 50 84 00 eb 00  ..H.T$HH.L$P....
	0x0080 48 89 44 24 58 48 c7 44 24 60 01 00 00 00 48 c7  H.D$XH.D$`....H.
	0x0090 44 24 68 01 00 00 00 48 89 04 24 48 c7 44 24 08  D$h....H..$H.D$.
	0x00a0 01 00 00 00 48 c7 44 24 10 01 00 00 00 e8 00 00  ....H.D$........
	0x00b0 00 00 48 8b 6c 24 70 48 83 c4 78 c3 e8 00 00 00  ..H.l$pH..x.....
	0x00c0 00 e9 3a ff ff ff                                ..:...
	rel 12+4 t=17 TLS+0
	rel 50+4 t=8 "".returnButDefer+0
	rel 68+4 t=8 runtime.convT64+0
	rel 110+4 t=16 type.int+0
	rel 174+4 t=8 fmt.Println+0
	rel 189+4 t=8 runtime.morestack_noctxt+0
"".returnButDefer.func1 STEXT nosplit size=21 args=0x8 locals=0x0
	0x0000 00000 (main.go:7)	TEXT	"".returnButDefer.func1(SB), NOSPLIT|ABIInternal, $0-8
	0x0000 00000 (main.go:7)	PCDATA	$0, $-2
	0x0000 00000 (main.go:7)	PCDATA	$1, $-2
	0x0000 00000 (main.go:7)	FUNCDATA	$0, gclocals路1a65e721a2ccc325b382662e7ffee780(SB)
	0x0000 00000 (main.go:7)	FUNCDATA	$1, gclocals路69c1753bd5f81501d95132d08af04464(SB)
	0x0000 00000 (main.go:7)	FUNCDATA	$2, gclocals路bfec7e55b3f043d1941c093912808913(SB)
	0x0000 00000 (main.go:8)	PCDATA	$0, $1
	0x0000 00000 (main.go:8)	PCDATA	$1, $0
	0x0000 00000 (main.go:8)	MOVQ	"".&it+8(SP), AX
	0x0005 00005 (main.go:8)	PCDATA	$0, $2
	0x0005 00005 (main.go:8)	PCDATA	$1, $1
	0x0005 00005 (main.go:8)	MOVQ	"".&it+8(SP), CX
	0x000a 00010 (main.go:8)	PCDATA	$0, $1
	0x000a 00010 (main.go:8)	MOVQ	(CX), CX
	0x000d 00013 (main.go:8)	ADDQ	$71, CX
	0x0011 00017 (main.go:8)	PCDATA	$0, $0
	0x0011 00017 (main.go:8)	MOVQ	CX, (AX)
	0x0014 00020 (main.go:9)	RET
	0x0000 48 8b 44 24 08 48 8b 4c 24 08 48 8b 09 48 83 c1  H.D$.H.L$.H..H..
	0x0010 47 48 89 08 c3                                   GH...
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.loc."".returnButDefer SDWARFLOC size=0
go.info."".returnButDefer SDWARFINFO size=55
	0x0000 03 22 22 2e 72 65 74 75 72 6e 42 75 74 44 65 66  ."".returnButDef
	0x0010 65 72 00 00 00 00 00 00 00 00 00 00 00 00 00 00  er..............
	0x0020 00 00 00 01 9c 00 00 00 00 01 0f 69 74 00 01 05  ...........it...
	0x0030 00 00 00 00 01 9c 00                             .......
	rel 0+0 t=24 type.int+0
	rel 0+0 t=24 type.noalg.struct { "".siz uint32; "".started bool; "".heap bool; "".openDefer bool; "".sp uintptr; "".pc uintptr; "".fn uintptr; ""._panic uintptr; "".link uintptr; "".framepc uintptr; "".varp uintptr; "".fd uintptr; "".args [8]uint8 }+0
	rel 19+8 t=1 "".returnButDefer+0
	rel 27+8 t=1 "".returnButDefer+146
	rel 37+4 t=30 gofile..E:\virtualShare\gopath3\src\github.com\Jameywoo\goto\playground\syntax\defer\asm\main.go+0
	rel 48+4 t=29 go.info.int+0
go.range."".returnButDefer SDWARFRANGE size=0
go.debuglines."".returnButDefer SDWARFMISC size=26
	0x0000 04 02 13 0a eb f7 06 5f 06 08 87 06 23 06 3b 06  ......._....#.;.
	0x0010 69 ab 06 4b 71 04 01 03 7c 01                    i..Kq...|.
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=45
	0x0000 03 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 0a 61 00 0f 00 00 00 00 03 91 b0 7f 00           .a...........
	rel 0+0 t=24 type.*[1]interface {}+0
	rel 0+0 t=24 type.[1]interface {}+0
	rel 0+0 t=24 type.[]interface {}+0
	rel 0+0 t=24 type.int+0
	rel 0+0 t=24 type.unsafe.Pointer+0
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+198
	rel 27+4 t=30 gofile..E:\virtualShare\gopath3\src\github.com\Jameywoo\goto\playground\syntax\defer\asm\main.go+0
	rel 36+4 t=29 go.info.int+0
go.range."".main SDWARFRANGE size=0
go.debuglines."".main SDWARFMISC size=26
	0x0000 04 02 03 08 14 0a 08 23 9c 6a 06 41 06 6a 06 37  .......#.j.A.j.7
	0x0010 06 02 58 f6 6f 04 01 03 73 01                    ..X.o...s.
runtime.nilinterequal路f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.nilinterequal+0
runtime.memequal64路f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=16
	0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64路f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.nilinterequal路f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=6 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=18
	0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface 
	0x0010 7b 7d                                            {}
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64路f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=6 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=19
	0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
	0x0010 20 7b 7d                                          {}
type.*[1]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 bf 03 a8 35 08 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64路f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} SRODATA dupok size=72
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.nilinterequal路f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 44+4 t=6 type.*[1]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
go.loc."".returnButDefer.func1 SDWARFLOC size=0
go.info."".returnButDefer.func1 SDWARFINFO size=62
	0x0000 03 22 22 2e 72 65 74 75 72 6e 42 75 74 44 65 66  ."".returnButDef
	0x0010 65 72 2e 66 75 6e 63 31 00 00 00 00 00 00 00 00  er.func1........
	0x0020 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0030 0f 26 69 74 00 00 07 00 00 00 00 01 9c 00        .&it..........
	rel 0+0 t=24 type.*int+0
	rel 25+8 t=1 "".returnButDefer.func1+0
	rel 33+8 t=1 "".returnButDefer.func1+21
	rel 43+4 t=30 gofile..E:\virtualShare\gopath3\src\github.com\Jameywoo\goto\playground\syntax\defer\asm\main.go+0
	rel 55+4 t=29 go.info.*int+0
go.range."".returnButDefer.func1 SDWARFRANGE size=0
go.debuglines."".returnButDefer.func1 SDWARFMISC size=14
	0x0000 04 02 03 02 14 06 41 06 a6 04 01 03 78 01        ......A.....x.
""..inittask SNOPTRDATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 fmt..inittask+0
"".returnButDefer.func1路f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".returnButDefer.func1+0
type..namedata.*[]uint8- SRODATA dupok size=11
	0x0000 00 00 08 2a 5b 5d 75 69 6e 74 38                 ...*[]uint8
type.*[]uint8 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 a5 8e d0 69 08 08 08 36 00 00 00 00 00 00 00 00  ...i...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64路f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]uint8-+0
	rel 48+8 t=1 type.[]uint8+0
type.[]uint8 SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 df 7e 2e 38 02 08 08 17 00 00 00 00 00 00 00 00  .~.8............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]uint8-+0
	rel 44+4 t=6 type.*[]uint8+0
	rel 48+8 t=1 type.uint8+0
type..namedata.*[8]uint8- SRODATA dupok size=12
	0x0000 00 00 09 2a 5b 38 5d 75 69 6e 74 38              ...*[8]uint8
type.*[8]uint8 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 a9 89 a5 7a 08 08 08 36 00 00 00 00 00 00 00 00  ...z...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64路f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[8]uint8-+0
	rel 48+8 t=1 type.[8]uint8+0
runtime.gcbits. SRODATA dupok size=0
type.[8]uint8 SRODATA dupok size=72
	0x0000 08 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 3e f9 30 b4 0a 01 01 11 00 00 00 00 00 00 00 00  >.0.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 08 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64路f+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*[8]uint8-+0
	rel 44+4 t=6 type.*[8]uint8+0
	rel 48+8 t=1 type.uint8+0
	rel 56+8 t=1 type.[]uint8+0
type..namedata.*struct { siz uint32; started bool; heap bool; openDefer bool; sp uintptr; pc uintptr; fn uintptr; _panic uintptr; link uintptr; framepc uintptr; varp uintptr; fd uintptr; args [8]uint8 }- SRODATA dupok size=190
	0x0000 00 00 bb 2a 73 74 72 75 63 74 20 7b 20 73 69 7a  ...*struct { siz
	0x0010 20 75 69 6e 74 33 32 3b 20 73 74 61 72 74 65 64   uint32; started
	0x0020 20 62 6f 6f 6c 3b 20 68 65 61 70 20 62 6f 6f 6c   bool; heap bool
	0x0030 3b 20 6f 70 65 6e 44 65 66 65 72 20 62 6f 6f 6c  ; openDefer bool
	0x0040 3b 20 73 70 20 75 69 6e 74 70 74 72 3b 20 70 63  ; sp uintptr; pc
	0x0050 20 75 69 6e 74 70 74 72 3b 20 66 6e 20 75 69 6e   uintptr; fn uin
	0x0060 74 70 74 72 3b 20 5f 70 61 6e 69 63 20 75 69 6e  tptr; _panic uin
	0x0070 74 70 74 72 3b 20 6c 69 6e 6b 20 75 69 6e 74 70  tptr; link uintp
	0x0080 74 72 3b 20 66 72 61 6d 65 70 63 20 75 69 6e 74  tr; framepc uint
	0x0090 70 74 72 3b 20 76 61 72 70 20 75 69 6e 74 70 74  ptr; varp uintpt
	0x00a0 72 3b 20 66 64 20 75 69 6e 74 70 74 72 3b 20 61  r; fd uintptr; a
	0x00b0 72 67 73 20 5b 38 5d 75 69 6e 74 38 20 7d        rgs [8]uint8 }
type.*struct { "".siz uint32; "".started bool; "".heap bool; "".openDefer bool; "".sp uintptr; "".pc uintptr; "".fn uintptr; ""._panic uintptr; "".link uintptr; "".framepc uintptr; "".varp uintptr; "".fd uintptr; "".args [8]uint8 } SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 05 65 a5 79 08 08 08 36 00 00 00 00 00 00 00 00  .e.y...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64路f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*struct { siz uint32; started bool; heap bool; openDefer bool; sp uintptr; pc uintptr; fn uintptr; _panic uintptr; link uintptr; framepc uintptr; varp uintptr; fd uintptr; args [8]uint8 }-+0
	rel 48+8 t=1 type.noalg.struct { "".siz uint32; "".started bool; "".heap bool; "".openDefer bool; "".sp uintptr; "".pc uintptr; "".fn uintptr; ""._panic uintptr; "".link uintptr; "".framepc uintptr; "".varp uintptr; "".fd uintptr; "".args [8]uint8 }+0
type..namedata.siz- SRODATA dupok size=6
	0x0000 00 00 03 73 69 7a                                ...siz
type..namedata.started- SRODATA dupok size=10
	0x0000 00 00 07 73 74 61 72 74 65 64                    ...started
type..namedata.heap- SRODATA dupok size=7
	0x0000 00 00 04 68 65 61 70                             ...heap
type..namedata.openDefer- SRODATA dupok size=12
	0x0000 00 00 09 6f 70 65 6e 44 65 66 65 72              ...openDefer
type..namedata.sp- SRODATA dupok size=5
	0x0000 00 00 02 73 70                                   ...sp
type..namedata.pc- SRODATA dupok size=5
	0x0000 00 00 02 70 63                                   ...pc
type..namedata.fn- SRODATA dupok size=5
	0x0000 00 00 02 66 6e                                   ...fn
type..namedata._panic- SRODATA dupok size=9
	0x0000 00 00 06 5f 70 61 6e 69 63                       ..._panic
type..namedata.link- SRODATA dupok size=7
	0x0000 00 00 04 6c 69 6e 6b                             ...link
type..namedata.framepc- SRODATA dupok size=10
	0x0000 00 00 07 66 72 61 6d 65 70 63                    ...framepc
type..namedata.varp- SRODATA dupok size=7
	0x0000 00 00 04 76 61 72 70                             ...varp
type..namedata.fd- SRODATA dupok size=5
	0x0000 00 00 02 66 64                                   ...fd
type..namedata.args- SRODATA dupok size=7
	0x0000 00 00 04 61 72 67 73                             ...args
type.noalg.struct { "".siz uint32; "".started bool; "".heap bool; "".openDefer bool; "".sp uintptr; "".pc uintptr; "".fn uintptr; ""._panic uintptr; "".link uintptr; "".framepc uintptr; "".varp uintptr; "".fd uintptr; "".args [8]uint8 } SRODATA dupok size=392
	0x0000 50 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  P...............
	0x0010 e4 2b b1 79 02 08 08 19 00 00 00 00 00 00 00 00  .+.y............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 0d 00 00 00 00 00 00 00 0d 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0090 0a 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00a0 00 00 00 00 00 00 00 00 0c 00 00 00 00 00 00 00  ................
	0x00b0 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00c0 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00d0 00 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00  ........ .......
	0x00e0 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00f0 30 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  0...............
	0x0100 00 00 00 00 00 00 00 00 40 00 00 00 00 00 00 00  ........@.......
	0x0110 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0120 50 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  P...............
	0x0130 00 00 00 00 00 00 00 00 60 00 00 00 00 00 00 00  ........`.......
	0x0140 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0150 70 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  p...............
	0x0160 00 00 00 00 00 00 00 00 80 00 00 00 00 00 00 00  ................
	0x0170 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0180 90 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*struct { siz uint32; started bool; heap bool; openDefer bool; sp uintptr; pc uintptr; fn uintptr; _panic uintptr; link uintptr; framepc uintptr; varp uintptr; fd uintptr; args [8]uint8 }-+0
	rel 44+4 t=6 type.*struct { "".siz uint32; "".started bool; "".heap bool; "".openDefer bool; "".sp uintptr; "".pc uintptr; "".fn uintptr; ""._panic uintptr; "".link uintptr; "".framepc uintptr; "".varp uintptr; "".fd uintptr; "".args [8]uint8 }+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type.noalg.struct { "".siz uint32; "".started bool; "".heap bool; "".openDefer bool; "".sp uintptr; "".pc uintptr; "".fn uintptr; ""._panic uintptr; "".link uintptr; "".framepc uintptr; "".varp uintptr; "".fd uintptr; "".args [8]uint8 }+80
	rel 80+8 t=1 type..namedata.siz-+0
	rel 88+8 t=1 type.uint32+0
	rel 104+8 t=1 type..namedata.started-+0
	rel 112+8 t=1 type.bool+0
	rel 128+8 t=1 type..namedata.heap-+0
	rel 136+8 t=1 type.bool+0
	rel 152+8 t=1 type..namedata.openDefer-+0
	rel 160+8 t=1 type.bool+0
	rel 176+8 t=1 type..namedata.sp-+0
	rel 184+8 t=1 type.uintptr+0
	rel 200+8 t=1 type..namedata.pc-+0
	rel 208+8 t=1 type.uintptr+0
	rel 224+8 t=1 type..namedata.fn-+0
	rel 232+8 t=1 type.uintptr+0
	rel 248+8 t=1 type..namedata._panic-+0
	rel 256+8 t=1 type.uintptr+0
	rel 272+8 t=1 type..namedata.link-+0
	rel 280+8 t=1 type.uintptr+0
	rel 296+8 t=1 type..namedata.framepc-+0
	rel 304+8 t=1 type.uintptr+0
	rel 320+8 t=1 type..namedata.varp-+0
	rel 328+8 t=1 type.uintptr+0
	rel 344+8 t=1 type..namedata.fd-+0
	rel 352+8 t=1 type.uintptr+0
	rel 368+8 t=1 type..namedata.args-+0
	rel 376+8 t=1 type.[8]uint8+0
type..importpath.fmt. SRODATA dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
gclocals路33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals路9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
gclocals路7d2d5fca80364273fb07d5820a76fef4 SRODATA dupok size=8
	0x0000 03 00 00 00 00 00 00 00                          ........
gclocals路a8c779e02c34a144ad683cd0297269a5 SRODATA dupok size=11
	0x0000 03 00 00 00 07 00 00 00 00 02 0a                 ...........
gclocals路f6aec3988379d2bd21c69c093370a150 SRODATA dupok size=12
	0x0000 04 00 00 00 03 00 00 00 00 01 03 07              ............
"".main.stkobj SRODATA size=24
	0x0000 01 00 00 00 00 00 00 00 d8 ff ff ff ff ff ff ff  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
	rel 16+8 t=1 type.[1]interface {}+0
gclocals路1a65e721a2ccc325b382662e7ffee780 SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 01 00                    ..........
gclocals路69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals路bfec7e55b3f043d1941c093912808913 SRODATA dupok size=11
	0x0000 03 00 00 00 02 00 00 00 00 01 03                 ...........
