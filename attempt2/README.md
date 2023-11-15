1. Start with hex to bytecode(represents each opcode)

2. Create Stack
- Each item on the stack is 256 bits long. Its called a word. Since numbers 256 bits long are not native within go we will use a special library to deal with these numbers for us.(Show how the holliman setsBytes and use hex.DecodeString() to illustrate when there's more than uint64 bytes in the 256)
- We can push and pop an item on or from the stack