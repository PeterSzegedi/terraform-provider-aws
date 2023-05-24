---
subcategory: ""
layout: "aws"
page_title: "Terraform AWS Provider Version 5 Upgrade Guide"
description: |-
  Terraform AWS Provider Version 5 Upgrade Guide
---

# Terraform AWS Provider Version 5 Upgrade Guide

Version 5.0.0 of the AWS provider for Terraform is a major release and includes changes that you need to consider when upgrading. This guide will help with that process and focuses only on changes from version 4.x to version 5.0.0. See the [Version 4 Upgrade Guide](/docs/providers/aws/guides/version-4-upgrade.html) for information on upgrading from 3.x to version 4.0.0.

Upgrade topics:

<!-- TOC depthFrom:2 depthTo:2 -->

- [Provider Version Configuration](#provider-version-configuration)
- [Provider Arguments](#provider-arguments)
- [Default Tags](#default-tags)
- [EC2-Classic Retirement](#ec2-classic-retirement)
- [Macie Classic Retirement](#macie-classic-retirement)
- [resource/aws_acmpca_certificate_authority](#resourceaws_acmpca_certificate_authority)
- [resource/aws_api_gateway_rest_api](#resourceaws_api_gateway_rest_api)
- [resource/aws_autoscaling_attachment](#resourceaws_autoscaling_attachment)
- [resource/aws_autoscaling_group](#resourceaws_autoscaling_group)
- [resource/aws_budgets_budget](#resourceaws_budgets_budget)
- [resource/aws_ce_anomaly_subscription](#resourceaws_ce_anomaly_subscription)
- [resource/aws_cloudwatch_event_target](#resourceaws_cloudwatch_event_target)
- [resource/aws_codebuild_project](#resourceaws_codebuild_project)
- [resource/aws_connect_hours_of_operation](#resourceaws_connect_hours_of_operation)
- [resource/aws_connect_queue](#resourceaws_connect_queue)
- [resource/aws_connect_routing_profile](#resourceaws_connect_routing_profile)
- [resource/aws_db_event_subscription](#resourceaws_db_event_subscription)
- [resource/aws_db_instance_role_association](#resourceaws_db_instance_role_association)
- [resource/aws_db_instance](#resourceaws_db_instance)
- [resource/aws_db_proxy_target](#resourceaws_db_proxy_target)
- [resource/aws_db_security_group](#resourceaws_db_security_group)
- [resource/aws_db_snapshot](#resourceaws_db_snapshot)
- [resource/aws_default_vpc](#resourceaws_default_vpc)
- [resource/aws_dms_endpoint](#resourceaws_dms_endpoint)
- [resource/aws_docdb_cluster](#resourceaws_docdb_cluster)
- [resource/aws_dx_gateway_association](#resourceaws_dx_gateway_association)
- [resource/aws_ec2_client_vpn_endpoint](#resourceaws_ec2_client_vpn_endpoint)
- [resource/aws_ec2_client_vpn_network_association](#resourceaws_ec2_client_vpn_network_association)
- [resource/aws_ecs_cluster](#resourceaws_ecs_cluster)
- [resource/aws_eip](#resourceaws_eip)
- [resource/aws_eip_association](#resourceaws_eip_association)
- [resource/aws_eks_addon](#resourceaws_eks_addon)
- [resource/aws_elasticache_cluster](#resourceaws_elasticache_cluster)
- [resource/aws_elasticache_replication_group](#resourceaws_elasticache_replication_group)
- [resource/aws_elasticache_security_group](#resourceaws_elasticache_security_group)
- [resource/aws_flow_log](#resourceaws_flow_log)
- [resource/aws_guardduty_organization_configuration](#resourceaws_guardduty_organization_configuration)
- [resource/aws_kinesis_firehose_delivery_stream](#resourceaws_kinesis_firehose_delivery_stream)
- [resource/aws_launch_configuration](#resourceaws_launch_configuration)
- [resource/aws_launch_template](#resourceaws_launch_template)
- [resource/aws_lightsail_instance](#resourceaws_lightsail_instance)
- [resource/aws_macie_member_account_association](#resourceaws_macie_member_account_association)
- [resource/aws_macie_s3_bucket_association](#resourceaws_macie_s3_bucket_association)
- [resource/aws_medialive_multiplex_program](#resourceaws_medialive_multiplex_program)
- [resource/aws_msk_cluster](#resourceaws_msk_cluster)
- [resource/aws_neptune_cluster](#resourceaws_neptune_cluster)
- [resource/aws_networkmanager_core_network](#resourceaws_networkmanager_core_network)
- [resource/aws_opensearch_domain](#resourceaws_opensearch_domain)
- [resource/aws_rds_cluster](#resourceaws_rds_cluster)
- [resource/aws_rds_cluster_instance](#resourceaws_rds_cluster_instance)
- [resource/aws_redshift_cluster](#resourceaws_redshift_cluster)
- [resource/aws_redshift_security_group](#resourceaws_redshift_security_group)
- [resource/aws_route](#resourceaws_route)
- [resource/aws_route_table](#resourceaws_route_table)
- [resource/aws_s3_object](#resourceaws_s3_object)
- [resource/aws_s3_object_copy](#resourceaws_s3_object_copy)
- [resource/aws_secretsmanager_secret](#resourceaws_secretsmanager_secret)
- [resource/aws_security_group](#resourceaws_security_group)
- [resource/aws_security_group_rule](#resourceaws_security_group_rule)
- [resource/aws_servicecatalog_product](#resourceaws_servicecatalog_product)
- [resource/aws_ssm_association](#resourceaws_ssm_association)
- [resource/aws_ssm_parameter](#resourceaws_ssm_parameter)
- [resource/aws_vpc](#resourceaws_vpc)
- [resource/aws_vpc_peering_connection](#resourceaws_vpc_peering_connection)
- [resource/aws_vpc_peering_connection_accepter](#resourceaws_vpc_peering_connection_accepter)
- [resource/aws_vpc_peering_connection_options](#resourceaws_vpc_peering_connection_options)
- [resource/aws_wafv2_web_acl](#resourceaws_wafv2_web_acl)
- [resource/aws_wafv2_web_acl_logging_configuration](#resourceaws_wafv2_web_acl_logging_configuration)
- [data-source/aws_api_gateway_rest_api](#data-sourceaws_api_gateway_rest_api)
- [data-source/aws_connect_hours_of_operation](#data-sourceaws_connect_hours_of_operation)
- [data-source/aws_db_instance](#data-sourceaws_db_instance)
- [data-source/aws_elasticache_cluster](#data-sourceaws_elasticache_cluster)
- [data-source/aws_elasticache_replication_group](#data-sourceaws_elasticache_replication_group)
- [data-source/aws_iam_policy_document](#data-sourceaws_iam_policy_document)
- [data-source/aws_identitystore_group](#data-sourceaws_identitystore_group)
- [data-source/aws_identitystore_user](#data-sourceaws_identitystore_user)
- [data-source/aws_launch_configuration](#data-sourceaws_launch_configuration)
- [data-source/aws_opensearch_domain](#data-sourceaws_opensearch_domain)
- [data-source/aws_quicksight_data_set](#data-sourceaws_quicksight_data_set)
- [data-source/aws_redshift_cluster](#data-sourceaws_redshift_cluster)
- [data-source/aws_redshift_service_account](#data-sourceaws_redshift_service_account)
- [data-source/aws_secretsmanager_secret](#data-sourceaws_secretsmanager_secret)
- [data-source/aws_service_discovery_service](#data-sourceaws_service_discovery_service)
- [data-source/aws_subnet_ids](#data-sourceaws_subnet_ids)
- [data-source/aws_vpc_peering_connection](#data-sourceaws_vpc_peering_connection)

<!-- /TOC -->

## Provider Version Configuration

-> Before upgrading to version 5.0.0, upgrade to the most recent 4.X version of the provider and ensure that your environment successfully runs [`terraform plan`](https://www.terraform.io/docs/commands/plan.html). You should not see changes you don't expect or deprecation notices.

Use [version constraints when configuring Terraform providers](https://www.terraform.io/docs/configuration/providers.html#provider-versions). If you are following that recommendation, update the version constraints in your Terraform configuration and run [`terraform init -upgrade`](https://www.terraform.io/docs/commands/init.html) to download the new version.

For example, given this previous configuration:

```terraform
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.65"
    }
  }
}

provider "aws" {
  # Configuration options
}
```

Update to the latest 5.X version:

```terraform
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  # Configuration options
}
```

## Provider Arguments

Version 5.0.0 removes these `provider` arguments:

* `assume_role.duration_seconds` - Use `assume_role.duration` instead
* `assume_role_with_web_identity.duration_seconds` - Use `assume_role_with_web_identity.duration` instead
* `s3_force_path_style` - Use `s3_use_path_style` instead
* `shared_credentials_file` - Use `shared_credentials_files` instead
* `skip_get_ec2_platforms` - Removed following the retirement of EC2-Classic

## Default Tags

The following enhancements are included:

* Duplicate `default_tags` can now be included and will be overwritten by resource `tags`.
* Zero value tags, `""`, can now be included in both `default_tags` and resource `tags`.
* Tags can now be `computed`.

## EC2-Classic Retirement

Following the retirement of EC2-Classic, we removed a number of resources, arguments, and attributes. This list summarizes what we _removed_:

* `aws_db_security_group` resource
* `aws_elasticache_security_group` resource
* `aws_redshift_security_group` resource
* [`aws_db_instance`](/docs/providers/aws/r/db_instance.html) resource's `security_group_names` argument
* [`aws_elasticache_cluster`](/docs/providers/aws/r/elasticache_cluster.html) resource's `security_group_names` argument
* [`aws_redshift_cluster`](/docs/providers/aws/r/redshift_cluster.html) resource's `cluster_security_groups` argument
* [`aws_launch_configuration`](/docs/providers/aws/r/launch_configuration.html) resource's `vpc_classic_link_id` and `vpc_classic_link_security_groups` arguments
* [`aws_vpc`](/docs/providers/aws/r/vpc.html) resource's `enable_classiclink` and `enable_classiclink_dns_support` arguments
* [`aws_default_vpc`](/docs/providers/aws/r/default_vpc.html) resource's `enable_classiclink` and `enable_classiclink_dns_support` arguments
* [`aws_vpc_peering_connection`](/docs/providers/aws/r/vpc_peering_connection.html) resource's `allow_classic_link_to_remote_vpc` and `allow_vpc_to_remote_classic_link` arguments
* [`aws_vpc_peering_connection_accepter`](/docs/providers/aws/r/vpc_peering_connection_accepter.html) resource's `allow_classic_link_to_remote_vpc` and `allow_vpc_to_remote_classic_link` arguments
* [`aws_vpc_peering_connection_options`](/docs/providers/aws/r/vpc_peering_connection_options.html) resource's `allow_classic_link_to_remote_vpc` and `allow_vpc_to_remote_classic_link` arguments
* [`aws_db_instance`](/docs/providers/aws/d/db_instance.html) data source's `db_security_groups` attribute
* [`aws_elasticache_cluster`](/docs/providers/aws/d/elasticache_cluster.html) data source's `security_group_names` attribute
* [`aws_redshift_cluster`](/docs/providers/aws/d/redshift_cluster.html) data source's `cluster_security_groups` attribute
* [`aws_launch_configuration`](/docs/providers/aws/d/launch_configuration.html) data source's `vpc_classic_link_id` and `vpc_classic_link_security_groups` attributes

## Macie Classic Retirement

Following the retirement of Amazon Macie Classic, we removed these resources:

* `aws_macie_member_account_association`
* `aws_macie_s3_bucket_association`

## resource/aws_acmpca_certificate_authority

Remove `status` from configurations as it no longer exists.

## resource/aws_api_gateway_rest_api

The `minimum_compression_size` attribute is now a String type, allowing it to be computed when set via the `body` attribute. Valid values remain the same.

## resource/aws_autoscaling_attachment

Change `alb_target_group_arn`, which no longer exists, to `lb_target_group_arn` in configurations.

## resource/aws_autoscaling_group

Remove `tags` from configurations as it no longer exists. Use the `tag` attribute instead. For use cases requiring dynamic tags, see the [Dynamic Tagging example](../r/autoscaling_group.html.markdown#dynamic-tagging).

## resource/aws_budgets_budget

Remove `cost_filters` from configurations as it no longer exists.

## resource/aws_ce_anomaly_subscription

Remove `threshold` from configurations as it no longer exists.

## resource/aws_cloudwatch_event_target

The `ecs_target.propagate_tags` attribute now has no default value. If no value is specified, the tags are not propagated.

## resource/aws_codebuild_project

Remove `secondary_sources.auth` and `source.auth` from configurations as they no longer exist.

## resource/aws_connect_hours_of_operation

Remove `hours_of_operation_arn` from configurations as it no longer exists.

## resource/aws_connect_queue

Remove `quick_connect_ids_associated` from configurations as it no longer exists.

## resource/aws_connect_routing_profile

Remove `queue_configs_associated` from configurations as it no longer exists.

## resource/aws_db_event_subscription

Configurations that define `source_ids` using the `id` attribute of `aws_db_instance` must be updated to use `identifier` instead. For example, `source_ids = [aws_db_instance.example.id]` must be updated to `source_ids = [aws_db_instance.example.identifier]`.

## resource/aws_db_instance

`aws_db_instance` has had a number of changes:

1. [`id` is no longer the identifier](#aws_db_instanceid-is-no-longer-the-identifier)
2. [Use `db_name` instead of `name`](#use-db_name-instead-of-name)
3. [Remove `db_security_groups`](#remove-db_security_groups)

### aws_db_instance.id is no longer the identifier

**What `id` _is_ has changed and can have far-reaching consequences.** Fortunately, fixing configurations is straightforward.

`id` is _now_ the DBI Resource ID (_i.e._, `dbi-resource-id`), an immutable "identifier" for an instance. `id` is now the same as the `resource_id`. (We recommend using `resource_id` rather than `id` when you need to refer to the DBI Resource ID.) _Previously_, `id` was the DB Identifier. Now when you need to refer to the _DB Identifier_, use `identifier`.

Fixing configurations involves changing any `id` references to `identifier`, where the reference expects the DB Identifier. For example, if you're replicating an `aws_db_instance`, you can no longer use `id` to define the `replicate_source_db`.

This configuration will now result in an error since `replicate_source_db` expects a _DB Identifier_:

```terraform
resource "aws_db_instance" "test" {
  replicate_source_db = aws_db_instance.source.id
  # ...other configuration...
}
```

You can fix the configuration like this:

```terraform
resource "aws_db_instance" "test" {
  replicate_source_db = aws_db_instance.source.identifier
  # ...other configuration...
}
```

### Use `db_name` instead of `name`

Change `name` to `db_name` in configurations as `name` no longer exists.

### Remove `db_security_groups`

Remove `db_security_groups` from configurations as it no longer exists. We removed it as part of the EC2-Classic retirement.

## resource/aws_db_instance_role_association

Configurations that define `db_instance_identifier` using the `id` attribute of `aws_db_instance` must be updated to use `identifier` instead. For example, `db_instance_identifier = aws_db_instance.example.id` must be updated to `db_instance_identifier = aws_db_instance.example.identifier`.

## resource/aws_db_proxy_target

Configurations that define `db_instance_identifier` using the `id` attribute of `aws_db_instance` must be updated to use `identifier` instead. For example, `db_instance_identifier = aws_db_instance.example.id` must be updated to `db_instance_identifier = aws_db_instance.example.identifier`.

## resource/aws_db_security_group

We removed this resource as part of the EC2-Classic retirement.

## resource/aws_db_snapshot

Configurations that define `db_instance_identifier` using the `id` attribute of `aws_db_instance` must be updated to use `identifier` instead. For example, `db_instance_identifier = aws_db_instance.example.id` must be updated to `db_instance_identifier = aws_db_instance.example.identifier`.

## resource/aws_default_vpc

Remove `enable_classiclink` and `enable_classiclink_dns_support` from configurations as they no longer exist. They were part of the EC2-Classic retirement.

## resource/aws_dms_endpoint

Remove `s3_settings.ignore_headers_row` from configurations as it no longer exists. **Be careful to not confuse `ignore_headers_row`, which no longer exists, with `ignore_header_rows`, which still exists.**

## resource/aws_docdb_cluster

Changes to the `snapshot_identifier` attribute will now correctly force re-creation of the resource. Previously, changing this attribute would result in a successful apply, but without the cluster being restored (only the resource state was changed). This change brings behavior of the cluster `snapshot_identifier` attribute into alignment with other RDS resources, such as `aws_db_instance`.

Automated snapshots **should not** be used for this attribute, unless from a different cluster. Automated snapshots are deleted as part of cluster destruction when the resource is replaced.

## resource/aws_dx_gateway_association

The `vpn_gateway_id` attribute has been deprecated. All configurations using `vpn_gateway_id` should be updated to use the `associated_gateway_id` attribute instead.

## resource/aws_ec2_client_vpn_endpoint

Remove `status` from configurations as it no longer exists.

## resource/aws_ec2_client_vpn_network_association

Remove `security_groups` and `status` from configurations as they no longer exist.

## resource/aws_ecs_cluster

Remove `capacity_providers` and `default_capacity_provider_strategy` from configurations as they no longer exist.

## resource/aws_eip

With the retirement of EC2-Classic, the `standard` domain is no longer supported.

## resource/aws_eip_association

With the retirement of EC2-Classic, the `standard` domain is no longer supported.

## resource/aws_eks_addon

The `resolve_conflicts` argument has been deprecated. Use the `resolve_conflicts_on_create` and/or `resolve_conflicts_on_update` arguments instead.

## resource/aws_elasticache_cluster

Remove `security_group_names` from configurations as it no longer exists. We removed it as part of the EC2-Classic retirement.

## resource/aws_elasticache_replication_group

* Remove the `cluster_mode` configuration block. Use top-level `num_node_groups` and `replicas_per_node_group` instead.
* Remove `availability_zones`, `number_cache_clusters`, `replication_group_description` arguments from configurations as they no longer exist.  Use `preferred_cache_cluster_azs`, `num_cache_clusters`, and `description`, respectively, instead.

## resource/aws_elasticache_security_group

We removed this resource as part of the EC2-Classic retirement.

## resource/aws_flow_log

The `log_group_name` attribute has been deprecated. All configurations using `log_group_name` should be updated to use the `log_destination` attribute instead.

## resource/aws_guardduty_organization_configuration

The `auto_enable` argument has been deprecated. Use the `auto_enable_organization_members` argument instead.

## resource/aws_kinesis_firehose_delivery_stream

* Remove the `s3_configuration` attribute from the root of the resource. `s3_configuration` is now a part of the following blocks: `elasticsearch_configuration`, `opensearch_configuration`, `redshift_configuration`, `splunk_configuration`, and `http_endpoint_configuration`.
* Remove `s3` as an option for `destination`. Use `extended_s3` instead
* Rename `extended_s3_configuration.0.s3_backup_configuration.0.buffer_size` and `extended_s3_configuration.0.s3_backup_configuration.0.buffer_interval` to `extended_s3_configuration.0.s3_backup_configuration.0.buffering_size` and `extended_s3_configuration.0.s3_backup_configuration.0.buffering_interval`, respectively.
* Rename `redshift_configuration.0.s3_backup_configuration.0.buffer_size` and `redshift_configuration.0.s3_backup_configuration.0.buffer_interval` to `redshift_configuration.0.s3_backup_configuration.0.buffering_size` and `redshift_configuration.0.s3_backup_configuration.0.buffering_interval`, respectively.
* Rename `s3_configuration.0.buffer_size` and `s3_configuration.0.buffer_interval` to `s3_configuration.0.buffering_size` and `s3_configuration.0.buffering_interval`, respectively.

## resource/aws_launch_configuration

Remove `vpc_classic_link_id` and `vpc_classic_link_security_groups` from configurations as they no longer exist. We removed them as part of the EC2-Classic retirement.

## resource/aws_launch_template

We removed defaults from `metatadata_options`. Launch template metadata options will now default to unset values, which is the AWS default behavior.

## resource/aws_lightsail_instance

Remove `ipv6_address` from configurations as it no longer exists.

## resource/aws_macie_member_account_association

We removed this resource as part of the Macie Classic retirement.

## resource/aws_macie_s3_bucket_association

We removed this resource as part of the Macie Classic retirement.

## resource/aws_medialive_multiplex_program

Change `statemux_settings`, which no longer exists, to `statmux_settings` in configurations.

## resource/aws_msk_cluster

Remove `broker_node_group_info.ebs_volume_size` from configurations as it no longer exists.

## resource/aws_neptune_cluster

Changes to the `snapshot_identifier` attribute will now correctly force re-creation of the resource. Previously, changing this attribute would result in a successful apply, but without the cluster being restored (only the resource state was changed). This change brings behavior of the cluster `snapshot_identifier` attribute into alignment with other RDS resources, such as `aws_db_instance`.

Automated snapshots **should not** be used for this attribute, unless from a different cluster. Automated snapshots are deleted as part of cluster destruction when the resource is replaced.

## resource/aws_networkmanager_core_network

Remove `policy_document` from configurations as it no longer exists. Use the `aws_networkmanager_core_network_policy_attachment` resource instead.

## resource/aws_opensearch_domain

The `kibana_endpoint` attribute has been deprecated. All configurations using `kibana_endpoint` should be updated to use the `dashboard_endpoint` attribute instead.

## resource/aws_rds_cluster

* Update configurations to always include `engine` since it is now required and has no default. Previously, not including `engine` was equivalent to `engine = "aurora"` and created a MySQL-5.6-compatible cluster.
* Changes to the `snapshot_identifier` attribute will now correctly force re-creation of the resource. Previously, changing this attribute would result in a successful apply, but without the cluster being restored (only the resource state was changed). This change brings behavior of the cluster `snapshot_identifier` attribute into alignment with other RDS resources, such as `aws_db_instance`. **NOTE:** Automated snapshots **should not** be used for this attribute, unless from a different cluster. Automated snapshots are deleted as part of cluster destruction when the resource is replaced.

## resource/aws_rds_cluster_instance

Update configurations to always include `engine` since it is now required and has no default. Previously, not including `engine` was equivalent to `engine = "aurora"` and created a MySQL-5.6-compatible cluster.

## resource/aws_redshift_cluster

Remove `cluster_security_groups` from configurations as it no longer exists. We removed it as part of the EC2-Classic retirement.

## resource/aws_redshift_security_group

We removed this resource as part of the EC2-Classic retirement.

## resource/aws_route

Update configurations to use `network_interface_id` rather than `instance_id`, which no longer exists.

For example, this configuration is _no longer valid_:

```terraform
resource "aws_route" "example" {
  instance_id = aws_instance.example.id
  # ...other configuration...
}
```

One possible way to fix this configuration involves referring to the `primary_network_interface_id` of an instance:

```terraform
resource "aws_route" "example" {
  network_interface_id = aws_instance.example.primary_network_interface_id
  # ...other configuration...
}
```

Another fix is to use an ENI:

```terraform
resource "aws_network_interface" "example" {
  # ...other configuration...
}

resource "aws_instance" "example" {
  network_interface {
    network_interface_id = aws_network_interface.example.id
    # ...other configuration...
  }

  # ...other configuration...
}

resource "aws_route" "example" {
  network_interface_id = aws_network_interface.example.id
  # ...other configuration...

  # Wait for the ENI attachment
  depends_on = [aws_instance.example]
}
```

## resource/aws_route_table

Update configurations to use `route.*.network_interface_id` rather than `route.*.instance_id`, which no longer exists.

For example, this configuration is _no longer valid_:

```terraform
resource "aws_route_table" "example" {
  route {
    instance_id = aws_instance.example.id
    # ...other configuration...
  }
  # ...other configuration...
}
```

One possible way to fix this configuration involves referring to the `primary_network_interface_id` of an instance:

```terraform
resource "aws_route_table" "example" {
  route {
    network_interface_id = aws_instance.example.primary_network_interface_id
    # ...other configuration...
  }

  # ...other configuration...
}
```

Another fix is to use an ENI:

```terraform
resource "aws_network_interface" "example" {
  # ...other configuration...
}

resource "aws_instance" "example" {
  network_interface {
    network_interface_id = aws_network_interface.example.id
    # ...other configuration...
  }

  # ...other configuration...
}

resource "aws_route_table" "example" {
  route {
    network_interface_id = aws_network_interface.example.id
    # ...other configuration...
  }

  # ...other configuration...

  # Wait for the ENI attachment
  depends_on = [aws_instance.example]
}
```

## resource/aws_s3_object

The `acl` attribute no longer has a default value. Previously this was set to `private` when omitted. Objects requiring a private ACL should now explicitly set this attribute.

## resource/aws_s3_object_copy

The `acl` attribute no longer has a default value. Previously this was set to `private` when omitted. Object copies requiring a private ACL should now explicitly set this attribute.

## resource/aws_secretsmanager_secret

Remove `rotation_enabled`, `rotation_lambda_arn` and `rotation_rules` from configurations as they no longer exist.

## resource/aws_security_group

With the retirement of EC2-Classic, non-VPC security groups are no longer supported.

## resource/aws_security_group_rule

With the retirement of EC2-Classic, non-VPC security groups are no longer supported.

## resource/aws_servicecatalog_product

Changes to any `provisioning_artifact_parameters` arguments now properly trigger a replacement. This fixes incorrect behavior, but may technically be breaking for configurations expecting non-functional in-place updates.

## resource/aws_ssm_association

The `instance_id` attribute has been deprecated. All configurations using `instance_id` should be updated to use the `targets` attribute instead.

## resource/aws_ssm_parameter

The `overwrite` attribute has been deprecated. Existing parameters should be explicitly imported rather than relying on the "import on create" behavior previously enabled by setting `overwrite = true`. In a future major version the `overwrite` attribute will be removed and attempting to create a parameter that already exists will fail.

## resource/aws_vpc

Remove `enable_classiclink` and `enable_classiclink_dns_support` from configurations as they no longer exist. They were part of the EC2-Classic retirement.

## resource/aws_vpc_peering_connection

Remove `allow_classic_link_to_remote_vpc` and `allow_vpc_to_remote_classic_link` from configurations as they no longer exist. They were part of the EC2-Classic retirement.

## resource/aws_vpc_peering_connection_accepter

Remove `allow_classic_link_to_remote_vpc` and `allow_vpc_to_remote_classic_link` from configurations as they no longer exist. They were part of the EC2-Classic retirement.

## resource/aws_vpc_peering_connection_options

Remove `allow_classic_link_to_remote_vpc` and `allow_vpc_to_remote_classic_link` from configurations as they no longer exist. They were part of the EC2-Classic retirement.

## resource/aws_wafv2_web_acl

* Remove `statement.managed_rule_group_statement.excluded_rule` and `statement.rule_group_reference_statement.excluded_rule` from configurations as they no longer exist.
* The `statement.rule_group_reference_statement.rule_action_override` attribute has been added.

## resource/aws_wafv2_web_acl_logging_configuration

Remove `redacted_fields.all_query_arguments`, `redacted_fields.body` and `redacted_fields.single_query_argument` from configurations as they no longer exist.

## data-source/aws_api_gateway_rest_api

The `minimum_compression_size` attribute is now a String type, allowing it to be computed when set via the `body` attribute.

## data-source/aws_connect_hours_of_operation

Remove `hours_of_operation_arn` from configurations as it no longer exists.

## data-source/aws_db_instance

Remove `db_security_groups` from configurations as it no longer exists. We removed it as part of the EC2-Classic retirement.

## data-source/aws_elasticache_cluster

Remove `security_group_names` from configurations as it no longer exists. We removed it as part of the EC2-Classic retirement.

## data-source/aws_elasticache_replication_group

Rename `number_cache_clusters` and `replication_group_description`, which no longer exist, to `num_cache_clusters`, and `description`, respectively.

## data-source/aws_iam_policy_document

* Remove `source_json` and `override_json` from configurations. Use `source_policy_documents` and `override_policy_documents`, respectively, instead.
* Don't add empty `statement.sid` values to `json` attribute value.

## data-source/aws_identitystore_group

Remove `filter` from configurations as it no longer exists.

## data-source/aws_identitystore_user

Remove `filter` from configurations as it no longer exists.

## data-source/aws_launch_configuration

Remove `vpc_classic_link_id` and `vpc_classic_link_security_groups` from configurations as they no longer exist. They were part of the EC2-Classic retirement.

## data-source/aws_opensearch_domain

The `kibana_endpoint` attribute has been deprecated. All configurations using `kibana_endpoint` should be updated to use the `dashboard_endpoint` attribute instead.

## data-source/aws_quicksight_data_set

The `tags_all` attribute has been deprecated and will be removed in a future version.

## data-source/aws_redshift_cluster

Remove `cluster_security_groups` from configurations as it no longer exists. We removed it as part of the EC2-Classic retirement.

## data-source/aws_redshift_service_account

[AWS document](https://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#db-auditing-bucket-permissions) that [a service principal name](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html#principal-services) be used instead of AWS account ID in any relevant IAM policy.
The [`aws_redshift_service_account`](/docs/providers/aws/d/redshift_service_account.html) data source should now be considered deprecated and will be removed in a future version.

## data-source/aws_service_discovery_service

The `tags_all` attribute has been deprecated and will be removed in a future version.

## data-source/aws_secretsmanager_secret

Remove `rotation_enabled`, `rotation_lambda_arn` and `rotation_rules` from configurations as they no longer exist.

## data-source/aws_subnet_ids

We removed the `aws_subnet_ids` data source. Use the [`aws_subnets`](/docs/providers/aws/d/subnets.html) data source instead.

## data-source/aws_vpc_peering_connection

Remove `allow_classic_link_to_remote_vpc` and `allow_vpc_to_remote_classic_link` from configurations as they no longer exist. They were part of the EC2-Classic retirement.