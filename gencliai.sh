A=$(cat .github/copilot-instructions.md)
B=$(cat prd.md)
cliai sc -p "Generate the application based on the following intructions and product specifications: $A $B"
