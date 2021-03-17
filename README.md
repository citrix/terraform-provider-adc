[![CircleCI](https://circleci.com/gh/citrix/terraform-provider-citrixadc/tree/master.svg?style=shield)](https://circleci.com/gh/citrix/terraform-provider-citrixadc/tree/master)
# terraform-provider-citrixadc

[Terraform](https://www.terraform.io) Provider for [Citrix
ADC](https://www.citrix.com/products/netscaler-adc/)


## Descritpion
Citrix has developed a custom Terraform provider for automating Citrix ADC deployments and configurations. Using [Terraform](https://www.terraform.io), you can [custom configure your ADCs](https://www.youtube.com/watch?v=IJIIWm5rzpQ&ab_channel=Citrix). Configure your ADCs using Terraform for different use-cases such as Load Balancing, SSL, Content Switching, GSLB, WAF etc. 

For deploying Citrix ADC in Public Cloud - AWS and Azure, check out cloud scripts in github repo [terraform-cloud-scripts](https://github.com/citrix/terraform-cloud-scripts)

All the Citrix ADC modules available for Terraform automation can be found in citrixadc folder. It uses the Nitro API to create/configure LB configurations. To get you started quickly we also have configuration examples in the example folder. You can modify them for your configurations or create your own.

Check out the blog to get an overview on automating Citrix ADC.

**Important note: The provider will not commit the config changes to Citrix ADC's persistent store.**

## Requirement

* [hashicorp/terraform](https://github.com/hashicorp/terraform)

## Contents
	1. General Description - What is Terraform ? What can you do with Citrix ADC ? 
	2. Navigating and Understanding Terraform Repository
	3. Quick Start on using Citrix ADC-Terraform provider:
		a. Installation
			i. Installing Terraform 
			ii. Installing Citrix ADC Provider
		b. Configuring ADC through Terraform
			i. Create terraform resources for basic Load Balancing use-case in Citrix ADC
			ii. Create terraform provider file
			iii. Running terraform commands to configure adc
	4. Use-Case supported through Terraform 
	5. Understanding resources file and repositories
	6. Using remote-exe

## Installation

### **Step 1. Installing Terraform CLI:**
First step is to install Terraform CLI. Refer the https://learn.hashicorp.com/tutorials/terraform/install-cli for installing Terraform CLI. 

### **Step 2. Installing Citrix ADC Provider:**
Terraform provider for Citrix ADC is not available through terrform.registry.io as of now. Hence users have to install the provider manually.

**Follow below steps to install citrix adc provider for Terraform CLI version < 13.0**


**Follow below steps to install citrix adc provider for Terraform CLI version >13.0**
i) Download the citrix adc terraform binary in your local machine where you have terraform installed from the [Releases section of the github repo](https://github.com/citrix/terraform-provider-citrixadc/releases).Untar the files and you can find the binary file terraform-provider-ctxadc.

ii) Create a following directory in your local machine and save the citrix adc terraform binary. e.g. in Ubuntu machine. Note that the directory structure has to be same as below, you can edit the version -0.12.43 to the citrix adc version you downloaded.
```
mkdir -p /home/user/.terraform.d/plugins/registry.terraform.io/citrix/citrixadc/0.12.43/linux_amd64/
```
iii) Copy the terraform-provider-citrixadc to the above created folder as shown below
```
cp terraform-provider-citrixadc /home/user/.terraform.d/plugins/registry.terraform.io/citrix/citrixadc/0.12.43/linux_amd64/
```

## Configuring ADC through Terraform
Here we will configure the basic use-case of setting up server in ADC through Terraform.
Before we configure, clone the github repository in your local machine as follows:
```
git clone https://github.com/citrix/terraform-provider-citrixadc/
```
**Step-1** : Now navigate to examples folder as below. Here you can find many ready to use examples for you to get started:
```
cd terraform-provider-citrixadc/examples/
```
Lets configure a simple server in citrix ADC.
```
cd terraform-provider-citrixadc/examples/simple_server/
```
**Step-2** : Provider.tf contains the details of the target Citrix ADC.Edit the simple_server/provider.tf as follows and add details of your target adc.
```
terraform {
    required_providers {
        citrixadc = {
            source = "citrix/citrixadc"
        }
    }
}
provider "citrixadc" {
  endpoint = "http://10.1.1.3:80"
  username = "UsernameOfYourADC"
  password = "PasswordOfYourADC"
 }
```
**Step-3** : Resources.tf contains the desired state of the resources that you want to manage through terraform.Here we want to create simple server. Edit the simple_server/resources.tf with your configuration values - name,ipaddress as below. 
```
resource "citrixadc_server" "test_server" {
  name      = "test_server"
  ipaddress = "192.168.2.2"
}
```
**Step-4** : Once the provider.tf and resources.tf is edited and saved with the desired values in the simple_server folder, you are good to run terraform and configure ADC.Initialize the terraform by running terraform-init inside the simple_server folder as follow:
```
terraform-provider-citrixadc/examples/simple_server$ terraform init
```
You should see following output if terraform was able to successfully find citrix adc provider and initialize it -
![image](https://user-images.githubusercontent.com/68320753/111422447-ba528d00-8714-11eb-91a6-02a1418b73eb.png)

**Step-5** : Now run the terraform-plan command. This will fetch the true state of your target ADC and will show you the changes/additions it need to make to achieve the desired configuration given in resources.tf. As we see below, terraform plans to add a new resource :
```
terraform-provider-citrixadc/examples/simple_server$ terraform plan
```

![image](https://user-images.githubusercontent.com/68320753/111422516-d5250180-8714-11eb-89e2-bc3d3432c9c7.png)

**Step-6** : If the above plan looks good, then go ahead and run terraform-apply to apply the configurations. Type yes, when prompted.**
```
terraform-provider-citrixadc/examples/simple_server$ terraform apply
```

![image](https://user-images.githubusercontent.com/68320753/111423045-b410e080-8715-11eb-9845-741b6398efbb.png)
![image](https://user-images.githubusercontent.com/68320753/111423077-bf640c00-8715-11eb-835a-fe36b90576db.png)
As you see terraform successfully created server with name test_server and given ipaddress on your target ADC. You can validate this through ADC GUI, checking out Traffic Management -> Load Balancing -> Servers. 

**Repeat Steps 1-6 above for different configurations that you want in Citrix ADC.**

## Usage

### Running
1. Copy the binary (either from the [build](#building) or from the
   [releases](https://github.com/citrix/terraform-provider-citrixadc/releases) page)
   `terraform-provider-citrixadc` to an appropriate location.

   [Configure](https://www.terraform.io/docs/plugins/basics.html) `.terraformrc` to use the
   `citrixadc` provider. An example `.terraformrc`:

```
providers {
    citrixadc = "<path-to-custom-providers>/terraform-provider-citrixadc"
}
```

2. Run `terraform` as usual 

```
terraform plan
terraform apply
```
3. The provider will not commit the config changes to Citrix ADC's persistent store. To do this, run the shell script `ns_commit.sh`:

```
export NS_URL=http://<host>:<port>/
export NS_LOGIN=nsroot
export NS_PASSWORD=nsroot
./ns_commit.sh
```

To ensure that the config is saved on every run, we can use something like `terraform apply && ns_commit.sh`

### Provider Configuration

```
provider "citrixadc" {
    username = "${var.ns_user}"
    password = "${var.ns_password}"
    endpoint = "http://10.71.136.250/"
}
```

We can use a `https` URL and accept the untrusted authority certificate on the Citrix ADC by specifying `insecure_skip_verify = true`

To use `https` without the need to set `insecure_skip_verify = true` follow this [guide](https://support.citrix.com/article/CTX122521) on
how to replace the default TLS certificate with one from a trusted Certifcate Authority.

Use of `https` is preferred. Using `http` will result in all provider configuration variables as well as resource variables
to be transmitted in cleartext. Anyone observing the HTTP data stream will be able to parse sensitive values such as the provider password.

Avoid storing provider credentials in the local state by using a backend that supports encryption.
The hasicorp [vault provider](https://registry.terraform.io/providers/hashicorp/vault/latest/docs) is also recommended for
storing sensitive data.

##### Argument Reference

The following arguments are supported.

* `username` - This is the user name to access to Citrix ADC. Defaults to `nsroot` unless environment variable `NS_LOGIN` has been set
* `password` - This is the password to access to Citrix ADC. Defaults to `nsroot` unless environment variable `NS_PASSWORD` has been set
* `endpoint` - (Required) Nitro API endpoint in the form `http://<NS_IP>/` or `http://<NS_IP>:<PORT>/`. Can be specified in environment variable `NS_URL`
* `insecure_skip_verify` - (Optional, true/false) Whether to accept the untrusted certificate on the Citrix ADC when the Citrix ADC endpoint is `https`
* `proxied_ns` - (Optional, NSIP) The target Citrix ADC NSIP for MAS proxied calls. When this option is defined, `username`, `password` and `endpoint` must refer to the MAS proxy.

The username, password and endpoint can be provided in environment variables `NS_LOGIN`, `NS_PASSWORD` and `NS_URL`. 

### Resource Configuration

#### `citrixadc_lbvserver`

```
resource "citrixadc_lbvserver" "foo" {
  name = "sample_lb"
  ipv46 = "10.71.136.150"
  port = 443
  servicetype = "SSL"
  lbmethod = "ROUNDROBIN"
  persistencetype = "COOKIEINSERT"
  sslcertkey = "${citrixadc_sslcertkey.foo.certkey}"
  sslprofile = "ns_default_ssl_profile_secure_frontend"
}
```

##### Argument Reference
See <https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/12.0/configuration/load-balancing/lbvserver/lbvserver/> for possible values for these arguments and for an exhaustive list of arguments. Additionally, you can specify the SSL `certkey` to be bound to this `lbvserver` using the `sslcertkey` parameter

##### Note
Note that the attribute `state` is not synced with the remote object.
If the state of the lb vserver is out of sync with the terraform configuration you will need to manually taint the resource and apply the configuration again.

#### `citrixadc_service`

```
resource "citrixadc_service" "backend_1" {
  ip = "10.33.44.55"
  port = 80
  servicetype = "HTTP"
  lbvserver = "${citrixadc_lbvserver.foo.name}"
  lbmonitor = "${citrixadc_lbmonitor.foo.name}"
}
```

##### Argument Reference
See <https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/12.0/configuration/basic/service/service/> for possible values for these arguments and for an exhaustive list of arguments. Additionally, you can specify the LB vserver  to be bound to this service  using the `lbvserver` parameter, and the `lbmonitor` parameter specifies the LB monitor to be bound.

##### Note
Note that the attribute `state` is not synced with the remote object.
If the state of the service is out of sync with the terraform configuration you will need to manually taint the resource and apply the configuration again.

#### `citrixadc_servicegroup`

```
resource "citrixadc_servicegroup" "backend_1" {
  servicegroupname = "backend_group_1"
  servicetype = "HTTP"
  lbvservers = ["${citrixadc_lbvserver.foo.name}]"
  lbmonitor = "${citrixadc_lbmonitor.foo.name}"
  servicegroupmembers = ["172.20.0.20:200:50","172.20.0.101:80:10",  "172.20.0.10:80:40"]
  servicegroupmembers_by_servername = ["server_1:200:50","server_2:80:10",  "server_3:80:40"]

}
```

##### Argument Reference
See <https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/12.0/configuration/basic/servicegroup/servicegroup/> for possible values for these arguments and for an exhaustive list of arguments. Additionally, you can specify the LB vservers  to be bound to this service using the `lbvservers` parameter. The `lbmonitor` parameter specifies the LB monitor to be bound.

`servicegroupmembers_by_servername` gives the ability to define servicegroup members by providing the server name. The heuristic rule for assigning members to either `servicegroupmembers_by_servername` or `servicegroupmembers` is whether the `servername` and `ip` property of the binding as read from the Citrix Adc configuration have idetical values. When the values are identical the member is classified as a `servicegroupmembers`. When they differ the member is classified as `servicegroupmembers_by_servername`.

#### `citrixadc_csvserver`

```
resource "citrixadc_csvserver" "foo" {
  name = "sample_cs"
  ipv46 = "10.71.139.151"
  servicetype = "SSL"
  port = 443
  sslprofile = "ns_default_ssl_profile_secure_frontend"
}
```

##### Argument Reference
See <https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/12.0/configuration/content-switching/csvserver/csvserver/> for possible values for these arguments and for an exhaustive list of arguments. Additionally, you can specify the SSL cert to be bound using the `sslcertkey` parameter

##### Note
Note that the attribute `state` is not synced with the remote object.
If the state of the cs vserver is out of sync with the terraform configuration you will need to manually taint the resource and apply the configuration again.

#### `citrixadc_sslcertkey`

```
resource "citrixadc_sslcertkey" "foo" {
  certkey = "sample_ssl_cert"
  cert = "/var/certs/server.crt"
  key = "/var/certs/server.key"
  expirymonitor = "ENABLED"
  notificationperiod = 90
}
```

##### Argument Reference
See <https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/12.0/configuration/ssl/sslcertkey/sslcertkey/> for possible values for these arguments and for an exhaustive list of arguments. 


#### `citrixadc_cspolicy`

```
resource "citrixadc_cspolicy" "foo" {
  policyname = "sample_cspolicy"
  url = "/cart/*"
  csvserver = "${citrixadc_csvserver.foo.name}"
  targetlbvserver = "${citrixadc_lbvserver.foo.name}"
}
```

##### Argument Reference
See <https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/12.0/configuration/content-switching/cspolicy/cspolicy/> for possible values for these arguments and for an exhaustive list of arguments. 


#### `citrixadc_lbmonitor`

```
resource "citrixadc_lbmonitor" "foo" {
  monitorname = "sample_lb_monitor"
  type = "HTTP"
  interval = 350
  resptimeout = 250
}
```

##### Argument Reference
See <https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/12.0/configuration/load-balancing/lbmonitor/lbmonitor/> for possible values for these arguments and for an exhaustive list of arguments. 

#### `citrixadc_gslbvserver`

```
resource "citrixadc_gslbvserver" "foo" {
  
  dnsrecordtype = "A"
  name = "GSLB-East-Coast-Vserver"
  servicetype = "HTTP"
  domain {
	  domainname =  "www.fooco.co"
	  ttl = "60"
  }
  domain {
	  domainname = "www.barco.com"
	  ttl = "55"
  }
  service {
          servicename = "Gslb-EastCoast-Svc"
          weight = "10"
  }
}
```

##### Argument Reference
See <https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/12.0/configuration/global-server-load-balancing/gslbvserver/gslbvserverl> for possible values for these arguments and for an exhaustive list of arguments. Additionally, you can specify the GSLB services  to be bound to this service using the `service` parameter. 

#### `citrixadc_gslbservice`

```
resource "citrixadc_gslbservice" "foo" {
  
  ip = "172.16.1.101"
  port = "80"
  servicename = "gslb1vservice"
  servicetype = "HTTP"
  sitename = "${citrixadc_gslbsite.foo.sitename}"

}
```

##### Argument Reference
See <https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/12.0/configuration/global-server-load-balancing/gslbservice/gslbservice/> for possible values for these arguments and for an exhaustive list of arguments. 


#### `citrixadc_gslbsite`

```
resource "citrixadc_gslbsite" "foo" {
  
  siteipaddress = "172.31.11.20"
  sitename = "Site-GSLB-East-Coast"

}
```

##### Argument Reference
See <https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/12.0/configuration/global-server-load-balancing/gslbsite/gslbsite/> for possible values for these arguments and for an exhaustive list of arguments. 

#### `citrixadc_nsacls`

```
resource "citrixadc_nsacls" "allacls" {
  aclsname = "foo"
  "acl" {
  	aclname = "restrict"
  	protocol = "TCP"
  	aclaction = "DENY"
  	destipval = "192.168.1.20"
  	srcportval = "49-1024"
        priority = 100
	}
  "acl"  {
  	aclname = "restrictvlan"
  	aclaction = "DENY"
  	vlan = "2000"
        priority = 130
  }
}

```

##### Argument Reference
You can have only one element of type `citrixadc_nsacls`. Encapsulating every `nsacl` inside the `citrixadc_nsacls` resource so that Terraform will automatically call `apply` on the `nsacls`.

See <https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/12.0/configuration/ns/nsacl/nsacl/#nsacl> for possible values for these arguments and for an exhaustive list of arguments. 

#### `citrixadc_inat`

```
resource "citrixadc_inat" "foo" {
  
  name = "ip4ip4"
  privateip = "192.168.2.5"
  publicip = "172.17.1.2"
}

```
See <https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/12.0/configuration/network/inat/inat/#inat> for possible values for these arguments and for an exhaustive list of arguments. 

#### `citrixadc_rnat`

```
resource "citrixadc_rnat" "allrnat" {
  depends_on = ["citrixadc_nsacls.allacls"]

  rnatsname = "rnatsall"

  rnat  {
      network = "192.168.88.0"
      netmask = "255.255.255.0"
      natip = "172.17.0.2"
  }

  rnat  {
      aclname = "RNAT_ACL_1"
  }
}

```

##### Argument Reference
You can have only one element of type `citrixadc_rnat`. Encapsulate every `rnat` inside the `citrixadc_rnat` resource.

See <https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/12.0/configuration/network/rnat/rnat/#rnat> for possible values for these arguments and for an exhaustive list of arguments. 

## Using `remote-exec` for one-time tasks
Terraform is useful for maintaining desired state for a set of resources. It is less useful for tasks such as network configuration which don't change. Network configuration is like using a provisioner inside Terraform. The directory `examples/remote-exec` show examples of how Terraform can use ssh to accomplish these one-time tasks.

## Building
### Assumption
* You have (some) experience with Terraform, the different provisioners and providers that come out of the box,
its configuration files, tfstate files, etc.
* You are comfortable with the Go language and its code organization.

1. Install `terraform` from <https://www.terraform.io/downloads.html>
2. Install `dep` (<https://github.com/golang/dep>)
3. Check out this code: `git clone https://<>`
4. Build this code using `make build`



## Samples
See the `examples` directory for various LB topologies that can be driven from this terraform provider.

