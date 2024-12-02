---
subcategory: "SSL"
---

# Resource: sslrsakey

The sslrsakey resource is used to create ssl rsakey file.


## Example usage

```hcl
resource "citrixadc_sslrsakey" "tf_sslrsakey" {
  keyfile          = "/nsconfig/ssl/key1.pem"
  bits             = 2048
  aes256           = true
  password         = "MySuperSecretPassword"
}
```


## Argument Reference

* `keyfile` - (Required) Name for and, optionally, path to the RSA key. /nsconfig/ssl/ is the default path. Maximum length =  63
* `bits` - (Required) Size, in bits, of the RSA key. Minimum value: 512 Maximum value: 4096
* `exponent` - (Optional) Public exponent for the RSA key. The exponent is part of the cipher algorithm and is required for creating the RSA key.  Possible values: 3, F4 Default value: F4
* `keyform` - (Optional) Format in which the key is stored on the appliance. Possible values: [ DER, PEM ] Default value: PEM
* `aes256` - (Optional) Encrypt the generated RSA key by using the AES algorithm.
* `des` - (Optional) Encrypt the generated RSA key by using the DES algorithm.
* `des3` - (Optional) Encrypt the generated RSA key by using the Triple-DES algorithm.
* `password` - (Optional) Pass phrase to use for encryption if AES256, DES or DES3 option is selected. Maximum value: 31
* `pkcs8` - (Optional) Create the private key in PKCS#8 format.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslrsakey. It is a unique string prefixed with "tf-sslrsakey-"

