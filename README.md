### CredParser:

This is a simple script that parses a space delimited file into json.

### Build:
`CGO_ENABLED=0 go build -a -installsuffix . -o credparser`

### Usage:
Inside `bin` there's a binary, just run :)<br>
(It's not a good pratice, I know)<br>
Example:
Contents of `passwords.txt`:
```
user_a pass_a domaina.com
user_b pass_b domainb.com
user_c pass_c domainc.com
user_d pass_d domaind.com
user_e pass_e domaine.com
user_f pass_f domainf.com
user_g pass_g domaing.com

```

`./credparser -f passwords.txt -o data.json`

The contents of `data.json` will be:
```
[
    {
        "username": "user_a@domaina.com",
        "passwords": {
            "password": "f5c27ee0023fc2d3d081c1bb33e9808f",
            "password_format": "md5"
        }
    },
    {
        "username": "user_b@domainb.com",
        "passwords": {
            "password": "c3fe6e989d0cbc857704a5e651a48e48",
            "password_format": "md5"
        }
    },
    {
        "username": "user_c@domainc.com",
        "passwords": {
            "password": "dd2304eedb82a8115840498d2847082c",
            "password_format": "md5"
        }
    },
    {
        "username": "user_d@domaind.com",
        "passwords": {
            "password": "c3d90f329e02cc97350fe8637bd1ba83",
            "password_format": "md5"
        }
    },
    {
        "username": "user_e@domaine.com",
        "passwords": {
            "password": "b4464d4bd4852fde42c7d7e610558805",
            "password_format": "md5"
        }
    },
    {
        "username": "user_f@domainf.com",
        "passwords": {
            "password": "4a57d8467cdb55700aa45217266e6bd4",
            "password_format": "md5"
        }
    },
    {
        "username": "user_g@domaing.com",
        "passwords": {
            "password": "d02189d197f7e08aa8c4c1d38e6382dd",
            "password_format": "md5"
        }
    }
]```
