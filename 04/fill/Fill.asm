// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed.
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

// Put your code here.

@LOOP
0;JMP

(LOOP)
  // write initialization
  @KBD
  D=M
  @keyin
  M=D
  @32
  D=A
  @m
  M=D
  @256
  D=A
  @n
  M=D
  @i
  M=0
  @k
  M=0
  @keyin
  D=M
  @SCREEN_BLACK
  D;JGT
  @SCREEN_WHITE
  0;JMP

(SCREEN_BLACK)
  // check stop condition i
  @j
  M=0
  @i
  D=M
  @n
  D=D-M
  @i
  M=M+1
  @LOOP
  D;JGE
  @SCREEN_BLACK_INNER
  0;JMP

(SCREEN_BLACK_INNER)
  // check stop condition j
  @j
  D=M
  @m
  D=D-M
  @SCREEN_BLACK
  D;JGE

  // make SCREEN pointer
  @SCREEN
  D=A
  @addr
  M=D

  // assign new value to pointer value
  @k
  D=M
  @addr
  A=M+D
  M=-1

  // refresh looping variables
  @j
  M=M+1
  @k
  M=M+1

  // goto next loop
  @SCREEN_BLACK_INNER
  0;JMP

(SCREEN_WHITE)
  @j
  M=0
  @i
  D=M
  @n
  D=D-M
  @i
  M=M+1
  @LOOP
  D;JGE
  @SCREEN_WHITE_INNER
  0;JMP

(SCREEN_WHITE_INNER)
  @j
  D=M
  @m
  D=D-M
  @SCREEN_WHITE
  D;JGE

  @SCREEN
  D=A
  @addr
  M=D

  @k
  D=M
  @addr
  A=M+D
  M=0

  @j
  M=M+1
  @k
  M=M+1

  @SCREEN_WHITE_INNER
  0;JMP
