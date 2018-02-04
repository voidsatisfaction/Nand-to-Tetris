// R[0]에 있는 숫자가 소수인지 아닌지 판별
// 소수이면 R[1]에 1을
// 소수가 아니면 R[1]에 0을 저장
// check whether R[0] is prime or not
// the result is saved to R[1]
@R0
D=M
@n
M=D

@n
D=M
@STOP_FALSE
D-1;JEQ

@2
D=A
@i
M=D

@LOOP
0;JMP

(LOOP)
  @i
  D=M
  @n
  // i*i <= n
  D=D-M
  @STOP_TRUE
  D+1;JEQ

  @temp // temp for calculating D%M
  M=0
  @LOOP_JUDGE
  0;JMP

(LOOP_JUDGE)
  @i
  D=M
  @temp
  M=M+D
  @temp
  D=M
  @n
  D=D-M
  @STOP_FALSE
  D;JEQ

  @LOOP_JUDGE
  D;JLT

  @i
  M=M+1
  @LOOP
  0;JMP

(STOP_TRUE)
  @R1
  M=1
  @END
  0;JMP

(STOP_FALSE)
  @R1
  M=0
  @END
  0;JMP

(END)
  @END
  0;JMP

```
