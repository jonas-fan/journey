# Cryptography

- [Classic cryptography](#classic-cryptography)
- [Morden cryptography](#morden-cryptography)
    - [Symmetric encryption](#symmetric-encryption)
    - [Asymmetric encryption](#asymmetric-encryption)
- [Integrity, Authentication, and Non-repudiation](#integrity-authentication-and-non-repudiation)
    - [Hash](#hash)
    - [Message Authentication Code (MAC)](#message-authentication-code-mac)
    - [Digital Signature](#digital-signature)

## Classic Cryptography

Entire system MUST be a secret.

### Caesar chiper

Rotating each letter in a plaintext with a fixed number.

```
encode(x, n)  = (x + n) mod 26
deconde(x, n) = (x - n) mod 26
```

An outstanding example is `rot13(x)` (aka `encode(x, 13)`)

```
s = rot13(x)
x = rot13(s)
```

## Morden Cryptography

System can be public, but keys MUST be a secret.

### Symmetric Encryption

```
  key
  +-----------+                 +------------+
A | plaintext |  -- encrypt ->  | ciphertext |
  +-----------+     w/ key      +------------+
                                     ////
                                     \\\\
                                    channel
                                     ////
                                     \\\\
  +-----------+                 +------------+
B | plaintext |  <- decrypt --  | ciphertext |
  +-----------+     w/ key      +------------+
  key
```

### Asymmetric Encryption

```
  A's public and private key
  B's public key
  +-----------+                 +------------+
A | plaintext |  -- encrypt ->  | ciphertext |
  +-----------+  w/ B's pub key +------------+
                                     ////
                                     \\\\
                                    channel
                                     ////
                                     \\\\
  +-----------+                 +------------+
B | plaintext |  <- decrypt --  | ciphertext |
  +-----------+  w/ B's pri key +------------+
  B's public and private key
  A's public key
```

## Integrity, Authentication, and Non-repudiation

| Approach                          | Integrity | Authentication | Non-repudiation |
|-----------------------------------|-----------|----------------|-----------------|
| Hash                              | O         | X              | X               |
| Message Authentication Code (MAC) | O         | O              | X               |
| Digital Signature                 | O         | O              | O               |

### Hash

`MD5`, `SHA-1`, `SHA-256`, etc.

### Message Authentication Code (MAC)

```
  key
  +---------+                       +-----+
A | message |  -- mac algorithm ->  | mac |
  +---------+        w/ key         +-----+
       |                               |
       |       +---------------+       |
       +---->  | message | mac |  <----+
               +---------------+
                     ////
                     \\\\
                    channel
                     ////
                     \\\\
               +---------------+
       +-----  | message | mac |  -----+
       |       +---------------+       |
  key  |                               |
  +---------+                       +-----+
B | message |                       | mac |
  +---------+                       +-----+
       |                               |
       |             +-----+           v
  mac algorithm ---> | mac | -------> same?
     w/ key          +-----+
```

### Digital Signature

```
  A's public and private key
  B's public key
  +---------+              +----------+                 +-----------+
A | message |  -- hash ->  | hash sum |  -- encrypt ->  | signature |
  +---------+              +----------+  w/ A's pri key +-----------+
       |                                                      |
       |              +---------------------+                 |
       +----------->  | message | signature |  <--------------+
                      +---------------------+
                               ////
                               \\\\
                              channel
                               ////
                               \\\\
                      +---------------------+
       +------------  | message | signature |  ---------------+
       |              +---------------------+                 |
       |                                                      |
       |                   +----------+                 +-----------+
       |                   | hash sum |  <- decrypt --  | signature |
       |                   +----------+  w/ A's pub key +-----------+
       |                        |
       |                        v
       |                       same?
       |                        ^
       |                        |
  +---------+              +----------+
B | message |  -- hash ->  | hash sum |
  +---------+              +----------+
  B's public and private key
  A's public key
```
