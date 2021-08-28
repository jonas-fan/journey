# Cryptography

## Classical Cryptography

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
