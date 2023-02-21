# Arrayfuck

### Brainfuck, with a twist

## What is this?

Arrayfuck is a superset of brainfuck. That means it extends brainfuck, every valid brainfuck program is a valid
arrayfuck program too.

## Extension

Arrayfuck features arrays, which can be used in the stack of the program.
Programms have two pointers now. One For choosing the position in the stack(the first) and one for choosing the position
in the array(the second).
Programms have a (virtually) infinit sized stack, considering out of arrays.
When moving the first pointer, the second pointer gets reset to 0.

## Extended Commands

| Command   | Description                                                 |
|-----------|-------------------------------------------------------------|
| ^         | Moves the pointer in the array up                           |
| v         | Moves the pointer in the array down                         |
| *         | Outputs the entire array                                    |
| u         | Copy the content of the current index to the higher index   |
| d         | Copy the content of the current index to the lower index    |

## Why would I need Arrayfuck

Well, you can create mor understandable code, without the need to use the extra features.
You also have more control about, where you store your data, for example you can choose the 5th array of your programm
to only contain user inputs, the 4th to contain Outputs, ect.

