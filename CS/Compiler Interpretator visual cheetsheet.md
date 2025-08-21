***Компилятор***
```
source.c ──> [Frontend (Clang) + Optimizer + Backend (LLVM)] ──> binary.exe ──> run
```

***Интерпретатор***
```
source.py ──> [Frontend] ──> AST/байткод ──> VM исполняет на лету
```

***JIT***
```
source.js ──> [Frontend] ──> байткод ──> VM исполняет + компилит "горячие" куски в машинный код
```
