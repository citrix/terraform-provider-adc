---
subcategory: "Generic"
---

# Data Source `nitro_info`

The nitro_info data source allows you to retrieve information of various NITRO objects.

## Example Usage

```terraform
# object_by_name example
data "citrixadc_nitro_info" "service_info" {
    workflow = {
        lifecycle = "object_by_name"
        endpoint = "service"
        bound_resource_missing_errorcode = 344
    }
    primary_id = "tf_service"
}

output "nitro_object" {
  value = data.citrixadc_nitro_info.service_info.nitro_object
}

# binding_list example
data "citrixadc_nitro_info" "sample" {
    workflow = {
        lifecycle = "binding_list"
        endpoint = "sslcertkey_sslvserver_binding"
        bound_resource_missing_errorcode: 1540
    }
    primary_id = citrixadc_sslcertkey.tf_sslcertkey.certkey
}

output "list_output" {
    value = data.citrixadc_nitro_info.sample.nitro_list
}

output "object_output" {
    value = [ for item in data.citrixadc_nitro_info.sample.nitro_list: item.object ]
}

# Fetch the content of a file using query_args:
data "citrixadc_nitro_info "my_file" {
    workflow = {
        lifecycle                        = "object_by_name"
        endpoint                         = "systemfile"
        bound_resource_missing_errorcode = "3441"
    }
    query_args = {
        filename    = "my_file_name"
        filecondent = urlencode("/my/file/path")
    }
}

output "my_file" {
    value = data.citrixadc_nitro_info.my_file.nitro_object.filecontent
}
```

## Argument Reference
* `workflow` - (Optional) Dictionary containing the data that will guide the data source execution.

    Currently the following attributes are taken into consideration.

    * `lifecycle` A predetermined list of strings that guide the execution of the data source.
    * `endpoint` The NITRO part of the url to the endpoint.
    * `bound_resource_missing_errorcode` The NITRO error code returned when the resource does not exist.

    The values of these attributes are determined by how the endpoint implements the read functionality.
    A user can provide these values to access data for various endpoints provided the lifecycle is supported.

    A list of such workflows can be found in the github repository example folder for the `nitro_info` data source.

* `query_args` - (Optional) A dictionary of query arguments that will be included when performing the request to the NITRO `endpoint` url.
* `primary_id` - (Optional) Value for the primary id of the nitro endpoint.
* `secondary_id` - (Optional) Value for the secondary id of the nitro endpoint.

## Attributes Reference

The following attributes are exported.

* `nitro_list` -  Contains the result returned by the `binding_list` lifecycle.
* `nitro_object` -  Contains the result returned by the `object_by_name` lifecycle.

