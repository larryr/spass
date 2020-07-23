# spass - secret password manager

Manage a single password by using Shamir's algorithm to break the secret 
into N parts and then to combine the parts and recover the secret using K 
parts where K <= N.

- [Shamir's Secret Sharing](https://en.wikipedia.org/wiki/Shamir%27s_Secret_Sharing)
- [Hashicorp Vault Imlpementation](https://github.com/hashicorp/vault)

# Example

## Split the secret into parts
This example submits the `secret` to be split into `5` parts with a recovery threshold of `2`:
```
$ echo "secret(c15845ad00ae)" | spass -p 5 -t 2 split 
c87492403e02bec7c066e069867e8ab850dbd2b334f5
4e157ac431192cd1284a80b423eef29ddf7aeb9e0f81
1a6ac595a5af5af6970470f6ad9d3e4810fe085fe953
ef71f01d7e4b90b4a280d8de072abc90b9ae00689cc4
98267292dc838d2620bff353edcfdb1e7faf292b4d85
```

## Combine the parts to recover the secret
Any `2` of the `5` generated parts are used to recover the secret:
```
$ cat <<EOF | spass combine
> ef71f01d7e4b90b4a280d8de072abc90b9ae00689cc4
> c87492403e02bec7c066e069867e8ab850dbd2b334f5
> EOF
secret(c15845ad00ae)
```

# Building
Requires Go installed ()> 1.12).  The `go build` command is used and will install the target binary `spass` in the default `bin` directory (e.g. `~/go/bin`)
```
make
```
to test 
```
make test
```

# History
Code copied from:
- https://github.com/kinvolk/go-shamir  (Apache License)
- [Hashicorp Imlpementation](https://github.com/hashicorp/vault/shamir) - (Mozilla)
 