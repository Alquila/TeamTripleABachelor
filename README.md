# TeamTripleABachelor
This project is an implementation of Elad Barkan, Eli Biham and Nathan Keller's 
attack on the stream cipher A5/2.

This project consists of the following modules.
* cipher.go
* cipher_sym.go
* ciphertext_only_attack.go
* gauss.go
* plaintext_attack.go
* print_module.go
* simple_cipher_stream.go
* simple_sym_cipher_stream.go


The A5/2 algorithm is implemented in ```cipher.go```, while a symbolic version
of the cipher is implemented in ``cipher_sym.go``. 
The symbolic version is needed for the attack.

The attacks Known Plaintext Attack and Ciphertext-only Attack are 
implemented in respectively ``plaintext_attack.go`` and ``ciphertext_only.go``.

To be able to perform the attack we made our own implementation of Gauss elimination.
This is placed in the ``gauss.go`` module. 

The two modules ``simple_cipher_stream.go`` and ``simple_sym_cipher_stream.go`` were
used in the beginning of the implementation process when implementing the attack.
They were used to set up a smaller version of the two attacks using only one register
before implementing the full attacks. 
The module ``print_module.go`` contains multiple printing methods, these are useful 
both for testing and for printing the output and result so it is easier to understand.


To run the Known Plaintext Attack run:
* There's two ways to run the attack. 

To run the Ciphertext-only Attack run:
* ????

## Code Coverage
The testing happened parallel with the implementation of the project.
When running Code Coverage with the current tests (5th of June) the
tests cover 96%. 