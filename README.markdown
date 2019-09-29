# Practical Cryptography for Developers (in Go)

This repository contains examples of how to use various Go libraries to perform cryptographic operations safely and with purpose. 

The primary purposes of cryptography are:

- Confidentiality
- Integrity
- Non-repudiation
- Authentication

## Examples

To build the examples, run `make`

Examples are built into the `./bin` directory

### Confidentiality

Symmetric examples: `cmd/encrypt` & `cmd/decrypt`
Asymmetric examples: `cmd/pkencrypt` & `cmd/pkdecrypt`

### Integrity

Hash example:   `cmd/hash`
Asymmetric example: `cmd/pksign` & `cmd/pkverify`

### Non-repudiation

Asymmetric examples: `pksign` & `pkverify`

### Authentication

TBD
