---
subcategory: "SSL"
---

# Resource: sslprofile

The sslprofile resource is used to create SSL profiles.


## Example usage

```hcl
resource "citrixadc_sslprofile" "tf_sslprofile" {
  name = "tf_sslprofile"
}
```

~> **Note:** The following attributes are available for configuration and can be used to control the default bindings of ECC curve and cipher bindings. Detailed descriptions can be found in the `Argument Reference` section below:
`nodefaultecccurvebindings`, `nodefaultcipherbindings`, `nodefaultbindings`

## Argument Reference

* `name` - (Optional) Name for the SSL profile.
* `sslprofiletype` - (Optional) Type of profile. Front end profiles apply to the entity that receives requests from a client. Backend profiles apply to the entity that sends client requests to a server. Possible values: [ BackEnd, FrontEnd ]
* `ssllogprofile` - (Optional) The name of the ssllogprofile.
* `dhcount` - (Optional) Number of interactions, between the client and the Citrix ADC, after which the DH private-public pair is regenerated. A value of zero (0) specifies infinite use (no refresh). This parameter is not applicable when configuring a backend profile. Allowed DH count values are 0 and >= 500.
* `dh` - (Optional) State of Diffie-Hellman (DH) key exchange. This parameter is not applicable when configuring a backend profile. Possible values: [ ENABLED, DISABLED ]
* `dhfile` - (Optional) The file name and path for the DH parameter.
* `ersa` - (Optional) State of Ephemeral RSA (eRSA) key exchange. Ephemeral RSA allows clients that support only export ciphers to communicate with the secure server even if the server certificate does not support export clients. The ephemeral RSA key is automatically generated when you bind an export cipher to an SSL or TCP-based SSL virtual server or service. When you remove the export cipher, the eRSA key is not deleted. It is reused at a later date when another export cipher is bound to an SSL or TCP-based SSL virtual server or service. The eRSA key is deleted when the appliance restarts. This parameter is not applicable when configuring a backend profile. Possible values: [ ENABLED, DISABLED ]
* `ersacount` - (Optional) The  refresh  count  for the re-generation of RSA public-key and private-key pair.
* `sessreuse` - (Optional) State of session reuse. Establishing the initial handshake requires CPU-intensive public key encryption operations. With the ENABLED setting, session key exchange is avoided for session resumption requests received from the client. Possible values: [ ENABLED, DISABLED ]
* `sesstimeout` - (Optional) The Session timeout value in seconds.
* `cipherredirect` - (Optional) State of Cipher Redirect. If this parameter is set to ENABLED, you can configure an SSL virtual server or service to display meaningful error messages if the SSL handshake fails because of a cipher mismatch between the virtual server or service and the client. This parameter is not applicable when configuring a backend profile. Possible values: [ ENABLED, DISABLED ]
* `cipherurl` - (Optional) The redirect URL to be used with the Cipher Redirect feature.
* `clientauth` - (Optional) State of client authentication. In service-based SSL offload, the service terminates the SSL handshake if the SSL client does not provide a valid certificate. This parameter is not applicable when configuring a backend profile. Possible values: [ ENABLED, DISABLED ]
* `clientcert` - (Optional) The rule for client certificate requirement in client authentication. Possible values: [ Mandatory, Optional ]
* `dhkeyexpsizelimit` - (Optional) This option enables the use of NIST recommended (NIST Special Publication 800-56A) bit size for private-key size. For example, for DH params of size 2048bit, the private-key size recommended is 224bits. This is rounded-up to 256bits. Possible values: [ ENABLED, DISABLED ]
* `sslredirect` - (Optional) State of HTTPS redirects for the SSL service. For an SSL session, if the client browser receives a redirect message, the browser tries to connect to the new location. However, the secure SSL session breaks if the object has moved from a secure site (https://) to an unsecure site (http://). Typically, a warning message appears on the screen, prompting the user to continue or disconnect. If SSL Redirect is ENABLED, the redirect message is automatically converted from http:// to https:// and the SSL session does not break. This parameter is not applicable when configuring a backend profile. Possible values: [ ENABLED, DISABLED ]
* `redirectportrewrite` - (Optional) State of the port rewrite while performing HTTPS redirect. If this parameter is set to ENABLED, and the URL from the server does not contain the standard port, the port is rewritten to the standard. Possible values: [ ENABLED, DISABLED ]
* `ssl3` - (Optional) State of SSLv3 protocol support for the SSL profile. Note: On platforms with SSL acceleration chips, if the SSL chip does not support SSLv3, this parameter cannot be set to ENABLED. Possible values: [ ENABLED, DISABLED ]
* `tls1` - (Optional) State of TLSv1.0 protocol support for the SSL profile. Possible values: [ ENABLED, DISABLED ]
* `tls11` - (Optional) State of TLSv1.1 protocol support for the SSL profile. Possible values: [ ENABLED, DISABLED ]
* `tls12` - (Optional) State of TLSv1.2 protocol support for the SSL profile. Possible values: [ ENABLED, DISABLED ]
* `tls13` - (Optional) State of TLSv1.3 protocol support for the SSL profile. Possible values: [ ENABLED, DISABLED ]
* `snienable` - (Optional) State of the Server Name Indication (SNI) feature on the virtual server and service-based offload. SNI helps to enable SSL encryption on multiple domains on a single virtual server or service if the domains are controlled by the same organization and share the same second-level domain name. For example, \*.sports.net can be used to secure domains such as login.sports.net and help.sports.net. Possible values: [ ENABLED, DISABLED ]
* `ocspstapling` - (Optional) State of OCSP stapling support on the SSL virtual server. Supported only if the protocol used is higher than SSLv3. Possible values: ENABLED: The appliance sends a request to the OCSP responder to check the status of the server certificate and caches the response for the specified time. If the response is valid at the time of SSL handshake with the client, the OCSP-based server certificate status is sent to the client during the handshake. DISABLED: The appliance does not check the status of the server certificate. . Possible values: [ ENABLED, DISABLED ]
* `serverauth` - (Optional) State of server authentication support for the SSL Backend profile. Possible values: [ ENABLED, DISABLED ]
* `commonname` - (Optional) Name to be checked against the CommonName (CN) field in the server certificate bound to the SSL server.
* `pushenctrigger` - (Optional) Trigger encryption on the basis of the PUSH flag value. Available settings function as follows: * ALWAYS - Any PUSH packet triggers encryption. * IGNORE - Ignore PUSH packet for triggering encryption. * MERGE - For a consecutive sequence of PUSH packets, the last PUSH packet triggers encryption. * TIMER - PUSH packet triggering encryption is delayed by the time defined in the set ssl parameter command or in the Change Advanced SSL Settings dialog box. Possible values: [ Always, Merge, Ignore, Timer ]
* `sendclosenotify` - (Optional) Enable sending SSL Close-Notify at the end of a transaction. Possible values: [ YES, NO ]
* `cleartextport` - (Optional) Port on which clear-text data is sent by the appliance to the server. Do not specify this parameter for SSL offloading with end-to-end encryption. Range 1 - 65535 * in CLI is represented as 65535 in NITRO API
* `insertionencoding` - (Optional) Encoding method used to insert the subject or issuer's name in HTTP requests to servers. Possible values: [ Unicode, UTF-8 ]
* `denysslreneg` - (Optional) Deny renegotiation in specified circumstances. Available settings function as follows: * NO - Allow SSL renegotiation. * FRONTEND_CLIENT - Deny secure and nonsecure SSL renegotiation initiated by the client. * FRONTEND_CLIENTSERVER - Deny secure and nonsecure SSL renegotiation initiated by the client or the Citrix ADC during policy-based client authentication. * ALL - Deny all secure and nonsecure SSL renegotiation. * NONSECURE - Deny nonsecure SSL renegotiation. Allows only clients that support RFC 5746. Possible values: [ NO, FRONTEND_CLIENT, FRONTEND_CLIENTSERVER, ALL, NONSECURE ]
* `quantumsize` - (Optional) Amount of data to collect before the data is pushed to the crypto hardware for encryption. For large downloads, a larger quantum size better utilizes the crypto resources. Possible values: [ 4096, 8192, 16384 ]
* `strictcachecks` - (Optional) Enable strict CA certificate checks on the appliance. Possible values: [ YES, NO ]
* `encrypttriggerpktcount` - (Optional)
* `pushflag` - (Optional) Insert PUSH flag into decrypted, encrypted, or all records. If the PUSH flag is set to a value other than 0, the buffered records are forwarded on the basis of the value of the PUSH flag. Available settings function as follows: 0 - Auto (PUSH flag is not set.) 1 - Insert PUSH flag into every decrypted record. 2 -Insert PUSH flag into every encrypted record. 3 - Insert PUSH flag into every decrypted and encrypted record.
* `dropreqwithnohostheader` - (Optional) Host header check for SNI enabled sessions. If this check is enabled and the HTTP request does not contain the host header for SNI enabled sessions(i.e vserver or profile bound to vserver has SNI enabled and 'Client Hello' arrived with SNI extension), the request is dropped. Possible values: [ YES, NO ]
* `snihttphostmatch` - (Optional) Controls how the HTTP 'Host' header value is validated. These checks are performed only if the session is SNI enabled (i.e when vserver or profile bound to vserver has SNI enabled and 'Client Hello' arrived with SNI extension) and HTTP request contains 'Host' header. Available settings function as follows: CERT   - Request is forwarded if the 'Host' value is covered by the certificate used to establish this SSL session. Note: 'CERT' matching mode cannot be applied in TLS 1.3 connections established by resuming from a previous TLS 1.3 session. On these connections, 'STRICT' matching mode will be used instead. STRICT - Request is forwarded only if value of 'Host' header in HTTP is identical to the 'Server name' value passed in 'Client Hello' of the SSL connection. NO     - No validation is performed on the HTTP 'Host' header value. Possible values: [ NO, CERT, STRICT ]
* `pushenctriggertimeout` - (Optional) PUSH encryption trigger timeout value. The timeout value is applied only if you set the Push Encryption Trigger parameter to Timer in the SSL virtual server settings.
* `ssltriggertimeout` - (Optional) Time, in milliseconds, after which encryption is triggered for transactions that are not tracked on the Citrix ADC because their length is not known. There can be a delay of up to 10ms from the specified timeout value before the packet is pushed into the queue.
* `clientauthuseboundcachain` - (Optional) Certficates bound on the VIP are used for validating the client cert. Certficates came along with client cert are not used for validating the client cert. Possible values: [ ENABLED, DISABLED ]
* `sslinterception` - (Optional) Enable or disable transparent interception of SSL sessions. Possible values: [ ENABLED, DISABLED ]
* `sslireneg` - (Optional) Enable or disable triggering the client renegotiation when renegotiation request is received from the origin server. Possible values: [ ENABLED, DISABLED ]
* `ssliocspcheck` - (Optional) Enable or disable OCSP check for origin server certificate. Possible values: [ ENABLED, DISABLED ]
* `sslimaxsessperserver` - (Optional)
* `sessionticket` - (Optional) This option enables the use of session tickets, as per the RFC 5077. Possible values: [ ENABLED, DISABLED ]
* `sessionticketlifetime` - (Optional) This option sets the life time of session tickets issued by NS in secs.
* `sessionticketkeyrefresh` - (Optional) This option enables the use of session tickets, as per the RFC 5077. Possible values: [ ENABLED, DISABLED ]
* `sessionticketkeydata` - (Optional) Session ticket enc/dec key , admin can set it.
* `sessionkeylifetime` - (Optional) This option sets the life time of symm key used to generate session tickets issued by NS in secs.
* `prevsessionkeylifetime` - (Optional) This option sets the life time of symm key used to generate session tickets issued by NS in secs.
* `hsts` - (Optional) State of HSTS protocol support for the SSL profile. Using HSTS, a server can enforce the use of an HTTPS connection for all communication with a client. Possible values: [ ENABLED, DISABLED ]
* `maxage` - (Optional) Set the maximum time, in seconds, in the strict transport security (STS) header during which the client must send only HTTPS requests to the server.
* `includesubdomains` - (Optional) Enable HSTS for subdomains. If set to Yes, a client must send only HTTPS requests for subdomains. Possible values: [ YES, NO ]
* `preload` - (Optional) Flag indicates the consent of the site owner to have their domain preloaded. Possible values: [ YES, NO ]
* `skipclientcertpolicycheck` - (Optional) This flag controls the processing of X509 certificate policies. If this option is Enabled, then the policy check in Client authentication will be skipped. This option can be used only when Client Authentication is Enabled and ClientCert is set to Mandatory. Possible values: [ ENABLED, DISABLED ]
* `zerorttearlydata` - (Optional) State of TLS 1.3 0-RTT early data support for the SSL Virtual Server. This setting only has an effect if resumption is enabled, as early data cannot be sent along with an initial handshake. Early application data has significantly different security properties - in particular there is no guarantee that the data cannot be replayed. Possible values: [ ENABLED, DISABLED ]
* `tls13sessionticketsperauthcontext` - (Optional) Number of tickets the SSL Virtual Server will issue anytime TLS 1.3 is negotiated, ticket-based resumption is enabled, and either (1) a handshake completes or (2) post-handhsake client auth completes. This value can be increased to enable clients to open multiple parallel connections using a fresh ticket for each connection. No tickets are sent if resumption is disabled.
* `dhekeyexchangewithpsk` - (Optional) Whether or not the SSL Virtual Server will require a DHE key exchange to occur when a PSK is accepted during a TLS 1.3 resumption handshake. A DHE key exchange ensures forward secrecy even in the event that ticket keys are compromised, at the expense of an additional round trip and resources required to carry out the DHE key exchange. If disabled, a DHE key exchange will be performed when a PSK is accepted but only if requested by the client. If enabled, the server will require a DHE key exchange when a PSK is accepted regardless of whether the client supports combined PSK-DHE key exchange. This setting only has an effect when resumption is enabled. Possible values: [ YES, NO ]
* `alpnprotocol` - (Optional) Protocol to negotiate with client and then send as part of the ALPN extension in the server hello message. Possible values are HTTP1.1, HTTP2 and NONE. Default is none i.e. ALPN extension will not be sent. This parameter is relevant only if ssl connection is handled by the virtual server of type SSL_TCP. This parameter has no effect if TLSv1.3 is negotiated. Possible values: [ NONE, HTTP1.1, HTTP2 ]
* `ciphername` - (Optional) The cipher group/alias/individual cipher configuration.
* `cipherpriority` - (Optional) cipher priority.
* `strictsigdigestcheck` - (Optional) Parameter indicating to check whether peer entity certificate during TLS1.2 handshake is signed with one of signature-hash combination supported by Citrix ADC. Possible values: [ ENABLED, DISABLED ]
* `nodefaultecccurvebindings` - (Optional) (bool) Set to `true` to unbind the default ECC curve, otherwise set to `false`. Possible values: [ true, false ]
* `nodefaultcipherbindings` - (Optional) (bool) Set to `true` to unbind the default ciphers, otherwise set to `false`. Possible values: [ true, false ]
* `nodefaultbindings` - (Optional) Set to "YES" to unbind both the default ECC curve and cipher, or "NO" to keep them bound. Possible values: [ "YES", "NO" ]
* `ecccurvebindings` - (Required) A set of ECC curve names to be bound to this SSL profile. (deprecates soon)

!>
[**DEPRECATED**] Please use `sslprofile_ecccurve_binding` to bind `ecccurve` to `sslprofile` insted of this resource. the support for binding `ecccurve` to `sslprofile` in `sslprofile` resource will get deprecated soon.

* `cipherbindings` - (Optional) A set of ciphersuite bindings to be bound to this SSL profile. Documented below. (deprecates soon)

!>
[**DEPRECATED**] Please use `sslprofile_sslcipher_binding` to bind `sslcipher` to `sslprofile` insted of this resource. the support for binding `sslcipher` to `sslprofile` in `sslprofile` resource will get deprecated soon.



A cipherbindings block supports the following:

* `ciphername` - (Required) Cipher name.
* `cipherpriority` - (Optional) This indicates priority assigned to the particular cipher.



## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslprofile. It has the same value as the `name` attribute.


## Import

A sslprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_sslprofile.tf_sslprofile tf_sslprofile
```
