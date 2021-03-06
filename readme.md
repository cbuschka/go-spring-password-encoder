# go-spring-password-encoder

[![Build](https://github.com/cbuschka/go-spring-password-encoder/workflows/build/badge.svg)](https://github.com/cbuschka/go-spring-password-encoder) [![Latest Release](https://img.shields.io/github/release/cbuschka/go-spring-password-encoder.svg)](https://github.com/cbuschka/go-spring-password-encoder/releases) [![Go Report Card](https://goreportcard.com/badge/github.com/cbuschka/go-spring-password-encoder)](https://goreportcard.com/report/github.com/cbuschka/go-spring-password-encoder) [![License](https://img.shields.io/github/license/cbuschka/go-spring-password-encoder.svg)](https://github.com/cbuschka/go-spring-password-encoder/blob/main/license.txt)

### Spring compatible password encoders written in golang

## Usage

Add entry to go.mod file:

```
[...]

go 1.16

require (
    github.com/cbuschka/go-spring-password-encoder v1.0.0
    [...]
)
```

import the package fromn within your code:

```
import "github.com/cbuschka/go-spring-password-encoder"
```

### SHA256PasswordEncoder (called StandardPasswordEncoder by spring)

Hint: This password encoder is considered legacy and too weak.

#### Hash and encode password

```
encoder := spring_password_encoder.NewDefaultSHA256PasswordEncoder()

encodedPasswordHash, err := encoder.Encode("asdfasdf")
```

#### Compare plain text password with hashed password

```
encoder := spring_password_encoder.NewDefaultSHA256PasswordEncoder()

encodedPasswordHash, err := encoder.Encode("asdfasdf")
```

See [sha256 password encoder test](./sha256_password_encoder_test.go)

### BCryptPasswordEncoder

#### Hash and encode password

```
encoder := spring_password_encoder.NewDefaultBCryptPasswordEncoder()

encodedPasswordHash, err := encoder.Encode("asdfasdf")
```

#### Compare plain text password with hashed password

```
encoder := spring_password_encoder.NewDefaultBCryptPasswordEncoder()

encodedPasswordHash, err := encoder.Encode("asdfasdf")
```

See [bcrypt password encoder test](./bcrypt_password_encoder_test.go)

## License

Copyright (c) 2021 by [Cornelius Buschka](https://github.com/cbuschka).

[MIT](./license.txt)


