## BAD BITCHES ONLY D-:<= 

# GO kommandoer
go test -run TestName (kør specific test)
go test -v (kør alle tests)

## When running the first time run: 
Run this inside the project:
 * go mod init module_name
 * go mod tidy


# Done
Register definition
Make Session Key
Initalise Register
Majority Function
Clock R1-R3

# In Progress

# Notation for extra bit
For a register r.Lenght is the lenght of the actual register. 
In the symbolic version the slices in the register will have lenght r.lenght+1 to account for the extra bit in the end.
This mean that when indexing over a slice we will typically only index over r.Lenght and handle slice[len(r)-1] - the last index - sepperately
When xoring two slices however the last bits are just xored together so there we can loop over len(r)
When clocking we only loop over r.Lenght


# TODO 


# Have been tested
Make Session Key
Majority Function


# Ready To test


# Should be tested
Initialize Registers
Clock R1-R3


# Refactor


# Spørgsmål
- Skal de bits der skal bruges til den sidste XOR, skal de fetches før eller efter der clockes? 
- Framenumber: når framenumber skal repræsenteres som et binært array, skal vi så have least significance bit i index 0? Eller er det bare et valg vi tager, og egentlig ligemeget?
- :'( 



# Noter fra møde med Ivan 

Vi gætter på hvad der står i register4 når det starter