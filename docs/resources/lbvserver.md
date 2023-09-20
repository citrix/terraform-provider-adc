---
subcategory: "Load Balancing"
---

# Resource: lbvserver

The lbvserver resource is used to create load balancing virtual servers.


## Example usage

```hcl
resource "citrixadc_lbvserver" "tf_lbvserver" {

  name = "tf_lbvserver"
  ipv46 = "10.202.11.11"
  lbmethod = "ROUNDROBIN"
  port = 80
  servicetype = "HTTP"
}


resource "citrixadc_lbvserver" "test_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "test_lbvserver"
  port        = 443
  servicetype = "SSL"

  sslpolicybinding {
      policyname = citrixadc_sslpolicy.tf_sslpolicy.name
     priority = 101
  }
  sslpolicybinding {
      policyname = citrixadc_sslpolicy.tf_sslpolicy2.name
     priority = 100
  }

}


resource "citrixadc_lbvserver" "tf_lbvserver_test" {
  name          = "tf_lbvserver"
  ipv46         = "10.202.11.11"
  lbmethod      = "ROUNDROBIN"
  port          = 80
  servicetype   = "HTTP"
  probeport     = 90
  probeprotocol = "HTTP"
}

```


## Argument Reference

* `name` - (Optional) Name for the virtual server.
* `servicetype` - (Optional) Protocol used by the service (also called the service type). Possible values: [ HTTP, FTP, TCP, UDP, SSL, SSL_BRIDGE, SSL_TCP, DTLS, NNTP, DNS, DHCPRA, ANY, SIP_UDP, SIP_TCP, SIP_SSL, DNS_TCP, RTSP, PUSH, SSL_PUSH, RADIUS, RDP, MYSQL, MSSQL, DIAMETER, SSL_DIAMETER, TFTP, ORACLE, SMPP, SYSLOGTCP, SYSLOGUDP, FIX, SSL_FIX, PROXY, USER_TCP, USER_SSL_TCP, QUIC, IPFIX, LOGSTREAM, MONGO, MONGO_TLS ]
* `ipv46` - (Optional) IPv4 or IPv6 address to assign to the virtual server.
* `ippattern` - (Optional) IP address pattern, in dotted decimal notation, for identifying packets to be accepted by the virtual server. The IP Mask parameter specifies which part of the destination IP address is matched against the pattern.  Mutually exclusive with the IP Address parameter. For example, if the IP pattern assigned to the virtual server is 198.51.100.0 and the IP mask is 255.255.240.0 (a forward mask), the first 20 bits in the destination IP addresses are matched with the first 20 bits in the pattern. The virtual server accepts requests with IP addresses that range from 198.51.96.1 to 198.51.111.254.  You can also use a pattern such as 0.0.2.2 and a mask such as 0.0.255.255 (a reverse mask). If a destination IP address matches more than one IP pattern, the pattern with the longest match is selected, and the associated virtual server processes the request. For example, if virtual servers vs1 and vs2 have the same IP pattern, 0.0.100.128, but different IP masks of 0.0.255.255 and 0.0.224.255, a destination IP address of 198.51.100.128 has the longest match with the IP pattern of vs1. If a destination IP address matches two or more virtual servers to the same extent, the request is processed by the virtual server whose port number matches the port number in the request.
* `ipmask` - (Optional) IP mask, in dotted decimal notation, for the IP Pattern parameter. Can have leading or trailing non-zero octets (for example, 255.255.240.0 or 0.0.255.255). Accordingly, the mask specifies whether the first n bits or the last n bits of the destination IP address in a client request are to be matched with the corresponding bits in the IP pattern. The former is called a forward mask. The latter is called a reverse mask.
* `port` - (Optional) Port number for the virtual server. Range 1 - 65535 * in CLI is represented as 65535 in NITRO API
* `ipset` - (Optional) The list of IPv4/IPv6 addresses bound to ipset would form a part of listening service on the current lb vserver.
* `range` - (Optional) Number of IP addresses that the appliance must generate and assign to the virtual server. The virtual server then functions as a network virtual server, accepting traffic on any of the generated IP addresses. The IP addresses are generated automatically, as follows: * For a range of n, the last octet of the address specified by the IP Address parameter increments n-1 times. * If the last octet exceeds 255, it rolls over to 0 and the third octet increments by 1. Note: The Range parameter assigns multiple IP addresses to one virtual server. To generate an array of virtual servers, each of which owns only one IP address, use brackets in the IP Address and Name parameters to specify the range. For example: add lb vserver my_vserver[1-3] HTTP 192.0.2.[1-3] 80.
* `persistencetype` - (Optional) Type of persistence for the virtual server. Available settings function as follows: SOURCEIP - Connections from the same client IP address belong to the same persistence session. COOKIEINSERT - Connections that have the same HTTP Cookie, inserted by a Set-Cookie directive from a server, belong to the same persistence session. SSLSESSION - Connections that have the same SSL Session ID belong to the same persistence session. CUSTOMSERVERID - Connections with the same server ID form part of the same session. For this persistence type, set the Server ID (CustomServerID) parameter for each service and configure the Rule parameter to identify the server ID in a request. RULE - All connections that match a user defined rule belong to the same persistence session. URLPASSIVE - Requests that have the same server ID in the URL query belong to the same persistence session. The server ID is the hexadecimal representation of the IP address and port of the service to which the request must be forwarded. This persistence type requires a rule to identify the server ID in the request. DESTIP - Connections to the same destination IP address belong to the same persistence session. SRCIPDESTIP - Connections that have the same source IP address and destination IP address belong to the same persistence session. CALLID - Connections that have the same CALL-ID SIP header belong to the same persistence session. RTSPSID - Connections that have the same RTSP Session ID belong to the same persistence session. FIXSESSION - Connections that have the same SenderCompID and TargetCompID values belong to the same persistence session. USERSESSION - Persistence session is created based on the persistence parameter value provided from an extension. Possible values: [ SOURCEIP, COOKIEINSERT, SSLSESSION, RULE, URLPASSIVE, CUSTOMSERVERID, DESTIP, SRCIPDESTIP, CALLID, RTSPSID, DIAMETER, FIXSESSION, USERSESSION, NONE ]
* `timeout` - (Optional) Time period for which a persistence session is in effect.
* `persistencebackup` - (Optional) Backup persistence type for the virtual server. Becomes operational if the primary persistence mechanism fails. Possible values: [ SOURCEIP, NONE ]
* `backuppersistencetimeout` - (Optional) Time period for which backup persistence is in effect.
* `lbmethod` - (Optional) Load balancing method.  The available settings function as follows: * ROUNDROBIN - Distribute requests in rotation, regardless of the load. Weights can be assigned to services to enforce weighted round robin distribution. * LEASTCONNECTION (default) - Select the service with the fewest connections. * LEASTRESPONSETIME - Select the service with the lowest average response time. * LEASTBANDWIDTH - Select the service currently handling the least traffic. * LEASTPACKETS - Select the service currently serving the lowest number of packets per second. * CUSTOMLOAD - Base service selection on the SNMP metrics obtained by custom load monitors. * LRTM - Select the service with the lowest response time. Response times are learned through monitoring probes. This method also takes the number of active connections into account. Also available are a number of hashing methods, in which the appliance extracts a predetermined portion of the request, creates a hash of the portion, and then checks whether any previous requests had the same hash value. If it finds a match, it forwards the request to the service that served those previous requests. Following are the hashing methods: * URLHASH - Create a hash of the request URL (or part of the URL). * DOMAINHASH - Create a hash of the domain name in the request (or part of the domain name). The domain name is taken from either the URL or the Host header. If the domain name appears in both locations, the URL is preferred. If the request does not contain a domain name, the load balancing method defaults to LEASTCONNECTION. * DESTINATIONIPHASH - Create a hash of the destination IP address in the IP header. * SOURCEIPHASH - Create a hash of the source IP address in the IP header. * TOKEN - Extract a token from the request, create a hash of the token, and then select the service to which any previous requests with the same token hash value were sent. * SRCIPDESTIPHASH - Create a hash of the string obtained by concatenating the source IP address and destination IP address in the IP header. * SRCIPSRCPORTHASH - Create a hash of the source IP address and source port in the IP header. * CALLIDHASH - Create a hash of the SIP Call-ID header. * USER_TOKEN - Same as TOKEN LB method but token needs to be provided from an extension. Possible values: [ ROUNDROBIN, LEASTCONNECTION, LEASTRESPONSETIME, URLHASH, DOMAINHASH, DESTINATIONIPHASH, SOURCEIPHASH, SRCIPDESTIPHASH, LEASTBANDWIDTH, LEASTPACKETS, TOKEN, SRCIPSRCPORTHASH, LRTM, CALLIDHASH, CUSTOMLOAD, LEASTREQUEST, AUDITLOGHASH, STATICPROXIMITY, USER_TOKEN ]
* `hashlength` - (Optional) Number of bytes to consider for the hash value used in the URLHASH and DOMAINHASH load balancing methods.
* `netmask` - (Optional) IPv4 subnet mask to apply to the destination IP address or source IP address when the load balancing method is DESTINATIONIPHASH or SOURCEIPHASH.
* `v6netmasklen` - (Optional) Number of bits to consider in an IPv6 destination or source IP address, for creating the hash that is required by the DESTINATIONIPHASH and SOURCEIPHASH load balancing methods.
* `backuplbmethod` - (Optional) Backup load balancing method. Becomes operational if the primary load balancing me thod fails or cannot be used. Valid only if the primary method is based on static proximity. Possible values: [ ROUNDROBIN, LEASTCONNECTION, LEASTRESPONSETIME, SOURCEIPHASH, LEASTBANDWIDTH, LEASTPACKETS, CUSTOMLOAD ]
* `cookiename` - (Optional) Use this parameter to specify the cookie name for COOKIE peristence type. It specifies the name of cookie with a maximum of 32 characters. If not specified, cookie name is internally generated.
* `rule` - (Optional) Expression, or name of a named expression, against which traffic is evaluated. The following requirements apply only to the Citrix ADC CLI: * If the expression includes one or more spaces, enclose the entire expression in double quotation marks. * If the expression itself includes double quotation marks, escape the quotations by using the \ character. * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `listenpolicy` - (Optional) Expression identifying traffic accepted by the virtual server. Can be either an expression (for example, CLIENT.IP.DST.IN_SUBNET(192.0.2.0/24) or the name of a named expression. In the above example, the virtual server accepts all requests whose destination IP address is in the 192.0.2.0/24 subnet.
* `listenpriority` - (Optional) Integer specifying the priority of the listen policy. A higher number specifies a lower priority. If a request matches the listen policies of more than one virtual server the virtual server whose listen policy has the highest priority (the lowest priority number) accepts the request.
* `resrule` - (Optional) Expression specifying which part of a server's response to use for creating rule based persistence sessions (persistence type RULE). Can be either an expression or the name of a named expression. Example: HTTP.RES.HEADER("setcookie").VALUE(0).TYPECAST_NVLIST_T('=',';').VALUE("server1").
* `persistmask` - (Optional) Persistence mask for IP based persistence types, for IPv4 virtual servers.
* `v6persistmasklen` - (Optional) Persistence mask for IP based persistence types, for IPv6 virtual servers.
* `pq` - (Optional) Use priority queuing on the virtual server. based persistence types, for IPv6 virtual servers. Possible values: [ on, off ]
* `probeport` - (Optional) Citrix ADC provides support for external health check of the vserver status. Select port for HTTP/TCP monitring. Default value: 0
* `probeprotocol` - (Optional) Citrix ADC provides support for external health check of the vserver status. Select HTTP or TCP probes for healthcheck.. Possible values: [ TCP, HTTP ]
* `sc` - (Optional) Use SureConnect on the virtual server. Possible values: [ on, off ]
* `rtspnat` - (Optional) Use network address translation (NAT) for RTSP data connections. Possible values: [ on, off ]
* `m` - (Optional) Redirection mode for load balancing. Available settings function as follows: * IP - Before forwarding a request to a server, change the destination IP address to the server's IP address. * MAC - Before forwarding a request to a server, change the destination MAC address to the server's MAC address.  The destination IP address is not changed. MAC-based redirection mode is used mostly in firewall load balancing deployments. * IPTUNNEL - Perform IP-in-IP encapsulation for client IP packets. In the outer IP headers, set the destination IP address to the IP address of the server and the source IP address to the subnet IP (SNIP). The client IP packets are not modified. Applicable to both IPv4 and IPv6 packets. * TOS - Encode the virtual server's TOS ID in the TOS field of the IP header. You can use either the IPTUNNEL or the TOS option to implement Direct Server Return (DSR). Possible values: [ IP, MAC, IPTUNNEL, TOS ]
* `tosid` - (Optional) TOS ID of the virtual server. Applicable only when the load balancing redirection mode is set to TOS.
* `datalength` - (Optional) Length of the token to be extracted from the data segment of an incoming packet, for use in the token method of load balancing. The length of the token, specified in bytes, must not be greater than 24 KB. Applicable to virtual servers of type TCP.
* `dataoffset` - (Optional) Offset to be considered when extracting a token from the TCP payload. Applicable to virtual servers, of type TCP, using the token method of load balancing. Must be within the first 24 KB of the TCP payload.
* `sessionless` - (Optional) Perform load balancing on a per-packet basis, without establishing sessions. Recommended for load balancing of intrusion detection system (IDS) servers and scenarios involving direct server return (DSR), where session information is unnecessary. Possible values: [ ENABLED, DISABLED ]
* `trofspersistence` - (Optional) When value is ENABLED, Trofs persistence is honored. When value is DISABLED, Trofs persistence is not honored. Possible values: [ ENABLED, DISABLED ]
* `state` - (Optional) State of the load balancing virtual server. Possible values: [ ENABLED, DISABLED ]
* `connfailover` - (Optional) Mode in which the connection failover feature must operate for the virtual server. After a failover, established TCP connections and UDP packet flows are kept active and resumed on the secondary appliance. Clients remain connected to the same servers. Available settings function as follows: * STATEFUL - The primary appliance shares state information with the secondary appliance, in real time, resulting in some runtime processing overhead. * STATELESS - State information is not shared, and the new primary appliance tries to re-create the packet flow on the basis of the information contained in the packets it receives. * DISABLED - Connection failover does not occur. Possible values: [ DISABLED, STATEFUL, STATELESS ]
* `redirurl` - (Optional) URL to which to redirect traffic if the virtual server becomes unavailable. WARNING! Make sure that the domain in the URL does not match the domain specified for a content switching policy. If it does, requests are continuously redirected to the unavailable virtual server.
* `cacheable` - (Optional) Route cacheable requests to a cache redirection virtual server. The load balancing virtual server can forward requests only to a transparent cache redirection virtual server that has an IP address and port combination of \*:80, so such a cache redirection virtual server must be configured on the appliance. Possible values: [ YES, NO ]
* `clttimeout` - (Optional) Idle time, in seconds, after which a client connection is terminated.
* `somethod` - (Optional) Type of threshold that, when exceeded, triggers spillover. Available settings function as follows: * CONNECTION - Spillover occurs when the number of client connections exceeds the threshold. * DYNAMICCONNECTION - Spillover occurs when the number of client connections at the virtual server exceeds the sum of the maximum client (Max Clients) settings for bound services. Do not specify a spillover threshold for this setting, because the threshold is implied by the Max Clients settings of bound services. * BANDWIDTH - Spillover occurs when the bandwidth consumed by the virtual server's incoming and outgoing traffic exceeds the threshold. * HEALTH - Spillover occurs when the percentage of weights of the services that are UP drops below the threshold. For example, if services svc1, svc2, and svc3 are bound to a virtual server, with weights 1, 2, and 3, and the spillover threshold is 50%, spillover occurs if svc1 and svc3 or svc2 and svc3 transition to DOWN. * NONE - Spillover does not occur. Possible values: [ CONNECTION, DYNAMICCONNECTION, BANDWIDTH, HEALTH, NONE ]
* `sopersistence` - (Optional) If spillover occurs, maintain source IP address based persistence for both primary and backup virtual servers. Possible values: [ ENABLED, DISABLED ]
* `sopersistencetimeout` - (Optional) Timeout for spillover persistence, in minutes.
* `healththreshold` - (Optional) Threshold in percent of active services below which vserver state is made down. If this threshold is 0, vserver state will be up even if one bound service is up.
* `sothreshold` - (Optional) Threshold at which spillover occurs. Specify an integer for the CONNECTION spillover method, a bandwidth value in kilobits per second for the BANDWIDTH method (do not enter the units), or a percentage for the HEALTH method (do not enter the percentage symbol).
* `sobackupaction` - (Optional) Action to be performed if spillover is to take effect, but no backup chain to spillover is usable or exists. Possible values: [ DROP, ACCEPT, REDIRECT ]
* `redirectportrewrite` - (Optional) Rewrite the port and change the protocol to ensure successful HTTP redirects from services. Possible values: [ ENABLED, DISABLED ]
* `downstateflush` - (Optional) Flush all active transactions associated with a virtual server whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions. Possible values: [ ENABLED, DISABLED ]
* `backupvserver` - (Optional) Name of the backup virtual server to which to forward requests if the primary virtual server goes DOWN or reaches its spillover threshold.
* `disableprimaryondown` - (Optional) If the primary virtual server goes down, do not allow it to return to primary status until manually enabled. Possible values: [ ENABLED, DISABLED ]
* `insertvserveripport` - (Optional) Insert an HTTP header, whose value is the IP address and port number of the virtual server, before forwarding a request to the server. The format of the header is <vipHeader>: <virtual server IP address>\_<port number >, where vipHeader is the name that you specify for the header. If the virtual server has an IPv6 address, the address in the header is enclosed in brackets ([ and ]) to separate it from the port number. If you have mapped an IPv4 address to a virtual server's IPv6 address, the value of this parameter determines which IP address is inserted in the header, as follows: * VIPADDR - Insert the IP address of the virtual server in the HTTP header regardless of whether the virtual server has an IPv4 address or an IPv6 address. A mapped IPv4 address, if configured, is ignored. * V6TOV4MAPPING - Insert the IPv4 address that is mapped to the virtual server's IPv6 address. If a mapped IPv4 address is not configured, insert the IPv6 address. * OFF - Disable header insertion. Possible values: [ OFF, VIPADDR, V6TOV4MAPPING ]
* `vipheader` - (Optional) Name for the inserted header. The default name is vip-header.
* `authenticationhost` - (Optional) Fully qualified domain name (FQDN) of the authentication virtual server to which the user must be redirected for authentication. Make sure that the Authentication parameter is set to ENABLED.
* `authentication` - (Optional) Enable or disable user authentication. Possible values: [ on, off ]
* `authn401` - (Optional) Enable or disable user authentication with HTTP 401 responses. Possible values: [ on, off ]
* `authnvsname` - (Optional) Name of an authentication virtual server with which to authenticate users.
* `push` - (Optional) Process traffic with the push virtual server that is bound to this load balancing virtual server. Possible values: [ ENABLED, DISABLED ]
* `pushvserver` - (Optional) Name of the load balancing virtual server, of type PUSH or SSL_PUSH, to which the server pushes updates received on the load balancing virtual server that you are configuring.
* `pushlabel` - (Optional) Expression for extracting a label from the server's response. Can be either an expression or the name of a named expression.
* `pushmulticlients` - (Optional) Allow multiple Web 2.0 connections from the same client to connect to the virtual server and expect updates. Possible values: [ YES, NO ]
* `tcpprofilename` - (Optional) Name of the TCP profile whose settings are to be applied to the virtual server.
* `httpprofilename` - (Optional) Name of the HTTP profile whose settings are to be applied to the virtual server.
* `dbprofilename` - (Optional) Name of the DB profile whose settings are to be applied to the virtual server.
* `comment` - (Optional) Any comments that you might want to associate with the virtual server.
* `l2conn` - (Optional) Use Layer 2 parameters (channel number, MAC address, and VLAN ID) in addition to the 4-tuple (<source IP>:<source port>::<destination IP>:<destination port>) that is used to identify a connection. Allows multiple TCP and non-TCP connections with the same 4-tuple to co-exist on the Citrix ADC. Possible values: [ on, off ]
* `oracleserverversion` - (Optional) Oracle server version. Possible values: [ 10G, 11G ]
* `mssqlserverversion` - (Optional) For a load balancing virtual server of type MSSQL, the Microsoft SQL Server version. Set this parameter if you expect some clients to run a version different from the version of the database. This setting provides compatibility between the client-side and server-side connections by ensuring that all communication conforms to the server's version. Possible values: [ 70, 2000, 2000SP1, 2005, 2008, 2008R2, 2012, 2014 ]
* `mysqlprotocolversion` - (Optional) MySQL protocol version that the virtual server advertises to clients.
* `mysqlserverversion` - (Optional) MySQL server version string that the virtual server advertises to clients.
* `mysqlcharacterset` - (Optional) Character set that the virtual server advertises to clients.
* `mysqlservercapabilities` - (Optional) Server capabilities that the virtual server advertises to clients.
* `appflowlog` - (Optional) Apply AppFlow logging to the virtual server. Possible values: [ ENABLED, DISABLED ]
* `netprofile` - (Optional) Name of the network profile to associate with the virtual server. If you set this parameter, the virtual server uses only the IP addresses in the network profile as source IP addresses when initiating connections with servers.
* `icmpvsrresponse` - (Optional) How the Citrix ADC responds to ping requests received for an IP address that is common to one or more virtual servers. Available settings function as follows: * If set to PASSIVE on all the virtual servers that share the IP address, the appliance always responds to the ping requests. * If set to ACTIVE on all the virtual servers that share the IP address, the appliance responds to the ping requests if at least one of the virtual servers is UP. Otherwise, the appliance does not respond. * If set to ACTIVE on some virtual servers and PASSIVE on the others, the appliance responds if at least one virtual server with the ACTIVE setting is UP. Otherwise, the appliance does not respond. Note: This parameter is available at the virtual server level. A similar parameter, ICMP Response, is available at the IP address level, for IPv4 addresses of type VIP. To set that parameter, use the add ip command in the CLI or the Create IP dialog box in the GUI. Possible values: [ PASSIVE, ACTIVE ]
* `rhistate` - (Optional) Route Health Injection (RHI) functionality of the NetSaler appliance for advertising the route of the VIP address associated with the virtual server. When Vserver RHI Level (RHI) parameter is set to VSVR\_CNTRLD, the following are different RHI behaviors for the VIP address on the basis of RHIstate (RHI STATE) settings on the virtual servers associated with the VIP address: * If you set RHI STATE to PASSIVE on all virtual servers, the Citrix ADC always advertises the route for the VIP address. * If you set RHI STATE to ACTIVE on all virtual servers, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers is in UP state. * If you set RHI STATE to ACTIVE on some and PASSIVE on others, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers, whose RHI STATE set to ACTIVE, is in UP state. Possible values: [ PASSIVE, ACTIVE ]
* `newservicerequest` - (Optional) Number of requests, or percentage of the load on existing services, by which to increase the load on a new service at each interval in slow-start mode. A non-zero value indicates that slow-start is applicable. A zero value indicates that the global RR startup parameter is applied. Changing the value to zero will cause services currently in slow start to take the full traffic as determined by the LB method. Subsequently, any new services added will use the global RR factor.
* `newservicerequestunit` - (Optional) Units in which to increment load at each interval in slow-start mode. Possible values: [ PER_SECOND, PERCENT ]
* `newservicerequestincrementinterval` - (Optional) Interval, in seconds, between successive increments in the load on a new service or a service whose state has just changed from DOWN to UP. A value of 0 (zero) specifies manual slow start.
* `minautoscalemembers` - (Optional)
* `maxautoscalemembers` - (Optional)
* `skippersistency` - (Optional) This argument decides the behavior incase the service which is selected from an existing persistence session has reached threshold. Possible values: [ Bypass, ReLb, None ]
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
* `authnprofile` - (Optional) Name of the authentication profile to be used when authentication is turned on.
* `macmoderetainvlan` - (Optional) This option is used to retain vlan information of incoming packet when macmode is enabled. Possible values: [ ENABLED, DISABLED ]
* `dbslb` - (Optional) Enable database specific load balancing for MySQL and MSSQL service types. Possible values: [ ENABLED, DISABLED ]
* `dns64` - (Optional) This argument is for enabling/disabling the dns64 on lbvserver. Possible values: [ ENABLED, DISABLED ]
* `bypassaaaa` - (Optional) If this option is enabled while resolving DNS64 query AAAA queries are not sent to back end dns server. Possible values: [ YES, NO ]
* `recursionavailable` - (Optional) When set to YES, this option causes the DNS replies from this vserver to have the RA bit turned on. Typically one would set this option to YES, when the vserver is load balancing a set of DNS servers thatsupport recursive queries. Possible values: [ YES, NO ]
* `processlocal` - (Optional) By turning on this option packets destined to a vserver in a cluster will not under go any steering. Turn this option for single packet request response mode or when the upstream device is performing a proper RSS for connection based distribution. Possible values: [ ENABLED, DISABLED ]
* `dnsprofilename` - (Optional) Name of the DNS profile to be associated with the VServer. DNS profile properties will be applied to the transactions processed by a VServer. This parameter is valid only for DNS and DNS-TCP VServers.
* `lbprofilename` - (Optional) Name of the LB profile which is associated to the vserver.
* `redirectfromport` - (Optional) Port number for the virtual server, from which we absorb the traffic for http redirect. Range 1 - 65535 * in CLI is represented as 65535 in NITRO API
* `httpsredirecturl` - (Optional) URL to which to redirect traffic if the traffic is recieved from redirect port.
* `retainconnectionsoncluster` - (Optional) This option enables you to retain existing connections on a node joining a Cluster system or when a node is being configured for passive timeout. By default, this option is disabled. Possible values: [ YES, NO ]
* `weight` - (Optional) Weight to assign to the specified service.
* `servicename` - (Optional) Service to bind to the virtual server.
* `redirurlflags` - (Optional) The redirect URL to be unset.
* `newname` - (Optional) New name for the virtual server.
* `sslcertkey` - (Optoinal) SSL certificate key to be bound to the lb vserver.
* `ciphersuites` - (Optional) A set of ciphersuites to be bound to the lb vserver.
* `ciphers` - (Optional) A set of ciphers to be bound to the lb vserver. This is only relevant for cluster deployments of ADC. In all other cases it is ignored.
* `snisslcertkeys` - (Optional) A set of SNI ssl certificates bound to the lb vserver.
* `sslprofile` - (Optional) A SSL profile to bind to this lb vserver.
* `sslpolicybinding` - (Optional) SSL policy bound to the lb vserver block. Documented below

SSL policy binding supports the following:

* `policyname` - (Optional) The name of the SSL policy binding.
* `priority` - (Optional) The priority of the policies bound to this SSL service.
* `type` - (Optional) Bind point to which to bind the policy. Possible Values: REQUEST, INTERCEPT_REQ and CLIENTHELLO_REQ. These bindpoints mean: 1. REQUEST: Policy evaluation will be done at appplication above SSL. This bindpoint is default and is used for actions based on clientauth and client cert. 2. INTERCEPT_REQ: Policy evaluation will be done during SSL handshake to decide whether to intercept or not. Actions allowed with this type are: INTERCEPT, BYPASS and RESET. 3. CLIENTHELLO_REQ: Policy evaluation will be done during handling of Client Hello Request. Action allowed with this type is: RESET, FORWARD and PICKCACERTGRP. Possible values: [ INTERCEPT_REQ, REQUEST, CLIENTHELLO_REQ ]
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke flag. This attribute is relevant only for ADVANCED policies.
* `labeltype` - (Optional) Type of policy label invocation. Possible values: [ vserver, service, policylabel ]
* `labelname` - (Optional) Name of the label to invoke if the current policy rule evaluates to TRUE.
* `quicbridgeprofilename` - (Optional) Name of the QUIC Bridge profile whose settings are to be applied to the virtual server. Refer [quicbridgeprofile](./quicbridgeprofile.md) for usage.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver. It has the same value as the `name` attribute.


## Import

A lbvserver can be imported using its name, e.g.

```shell
terraform import citrixadc_lbvserver.tf_lbvserver tf_lbvserver
```
