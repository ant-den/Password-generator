# Password Generator

A command-line Go application that generates **unique passwords** based on user-defined parameters.

---

##  Task Description

Test task for **SIA "Scandinavian World"**.

Create an application that generates passwords based on user-specified parameters:

###  Password Parameters:

- **Password length**: A number between 1 and the total number of selected characters.
- **Character sets** (user can choose one or more):
  - Digits (`0-9`)
  - Lowercase Latin letters (`a-z`)
  - Uppercase Latin letters (`A-Z`)

###  Password Requirements:

- Characters in the password **must not repeat**.
- Each generated password must be **unique** (not repeated across generations).
- If multiple character sets are selected, the password must contain **at least one character** from each selected set.

---

##  Interface

The application uses a simple **CLI (Command-Line Interface)** for interaction.

Other interface options that would be accepted:
- Web interface
- Native OS windows

>  Graphical design is not important.

---

## How to Run 

Make sure you have Go installed. Then run:
 go run main.go
