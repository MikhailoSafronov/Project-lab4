#  Brainfuck Interpreter in Go

Це інтерпретатор мови програмування [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck), написаний на Go з нуля.

Проєкт побудований за принципом: **Lexer → Parser → AST → Interpreter**.

##  Функціональність

-  Лексичний аналізатор (`lexer`)

-  Рекурсивний спусковий парсер (`parser`)

-  Абстрактне синтаксичне дерево (`AST`)

-  Інтерпретатор для виконання Brainfuck-коду

-  Повна підтримка циклів (`[ ]`) і команд (`><+-.,`)

##  Структура проєкту

```

brainfuck-app/

├── cmd/

│   └── bf/

│       └── main.go                

├── internal/

│   ├── ast/                      

│   │   ├── ast.go                         

│   ├── lexer/                     

│   │   ├── lexer.go                     

│   └── parser/                   

│       ├── parser.go             

│       ├── interpreter.go              

├── tests/                        

│   ├── hello.bf                                  

├── go.mod                       

├── .gitignore                    

└── README.md                     

```

## Як запустити

### Передумови

- [Go](https://golang.org/dl/) версії 1.16 або вище

### Встановлення та запуск

1. Клонуйте репозиторій:

```bash

git clone https://github.com/yourusername/brainfuck-app.git

cd brainfuck-app

```

2. Зберіть бінарник:

```bash

go build -o bf ./cmd/bf

```

3. Запустіть Brainfuck-програму:

```bash

./bf tests/hello.bf

```

Вивід:

```

Hello World!

```

### Альтернативний запуск (без збірки)

```bash

go run ./cmd/bf/main.go tests/hello.bf

```

## Приклад програми

Файл `tests/hello.bf`:

```brainfuck

++++++++++[>+++++++>++++++++++>+++>+<<<<-]

>++.>+.+++++++. ..+++.>++.<<+++++++++++++++.>.+++.------.--------.>.

```

##  Компоненти

### Lexer

Перетворює вхідні символи у токени (команди Brainfuck).

### Parser

Будує AST з токенів, враховуючи вкладені цикли.

### Interpreter

Виконує AST, використовуючи байтову стрічку на 30 000 клітинок.

## Автори

- Сафронов Михайло - [посилання](https://github.com/MikhailoSafronov)
- Сюсюков Володимир - [посилання](https://github.com/nevdaha2103)
