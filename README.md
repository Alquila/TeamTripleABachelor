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


## Run the two attacks
There are two ways to run the two attacks. As both attacks take up to several hours, it is possible
to run the attack within the interval of the correct solution. This will decrease the time it takes
to run the attack to a few minutes. 

The Known Plaintext Attack run:
* To run the short version of the Known Plaintext Attack you need to do the following:
  * First go to ``plaintext_attack.go`` and check that line 282 is not commented out, but line 283 is. 
  * Then run in the terminal: go test -timeout 0 -run TestKnownPlaintextAttack 
  * The code will now run within a set interval of the correct solution. 
* To run the normal version with 2^16 guesses do the following:
  * go to ``plaintext_attack.go`` and check that line 283 is not commented out, but line 282 is.
  * The run in the terminal: go test -timeout 0 -run TestKnownPlaintextAttack
  * Expect that the code might run up to 3 and a half hours. 

To run the Ciphertext-only Attack run:
* To run the short version of the Ciphertext-only Attack you need to do the following:
    * First go to ``ciphertext_only_attack_test.go`` and check that line 241 is not commented out, but line 242 is.
    * Then run in the terminal: go test -timeout 0 -run TestCiphertextOnlyAttack
    * The code will now run within a set interval of the correct solution.
* To run the normal version with 2^16 guesses do the following:
    * go to ``ciphertext_only_attack_test.go`` and check that line 242 is not commented out, but line 241 is.
    * The run in the terminal: go test -timeout 0 -run TestCiphertextOnlyAttack
    * Expect that the code might run up to 7 and a half hours.

NOTE: All tests were run on a MacBook Pro with a 2,3 GHz 8-Core Intel Core i9 CPU, and 16 GB 2667 MHz DDR4 RAM, 
and lastly, an internal Intel UHD Graphics 630 1536 MB GPU. 
So the time it takes to run the experiment may vary dependent on specs.


## Code Coverage
The testing happened parallel with the implementation of the project.
When running Code Coverage with the current tests (5th of June) the
tests cover 96%. 