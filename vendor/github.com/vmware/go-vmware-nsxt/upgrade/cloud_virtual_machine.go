/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package upgrade

import (
	"github.com/vmware/go-vmware-nsxt/common"
)

type CloudVirtualMachine struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected.
	Revision int64 `json:"_revision"`

	// Timestamp of resource creation
	CreateTime int64 `json:"_create_time,omitempty"`

	// ID of the user who created this resource
	CreateUser string `json:"_create_user,omitempty"`

	// Timestamp of last modification
	LastModifiedTime int64 `json:"_last_modified_time,omitempty"`

	// ID of the user who last modified this resource
	LastModifiedUser string `json:"_last_modified_user,omitempty"`

	// Indicates system owned resource
	SystemOwned bool `json:"_system_owned,omitempty"`

	// Description of this resource
	Description string `json:"description,omitempty"`

	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`

	// Unique identifier of this resource
	Id string `json:"id,omitempty"`

	// The type of this resource
	ResourceType string `json:"resource_type"`

	// Opaque identifiers meaningful to the API user
	Tags []common.Tag `json:"tags,omitempty"`

	// Agent Status
	AgentStatus string `json:"agent_status,omitempty"`

	// Agent version details
	AgentVersion string `json:"agent_version,omitempty"`

	// Cloud tags for the virtual machine
	CloudTags []CloudTag `json:"cloud_tags,omitempty"`

	// List of Error Messages obtained from PCM related to VM
	ErrorMessages []ComputeInstanceErrorMessage `json:"error_messages,omitempty"`

	// Index of HA that indicates whether gateway is primary or secondary. If index is 0, then it is primary gateway. Else secondary gateway.
	GatewayHaIndex int64 `json:"gateway_ha_index,omitempty"`

	// Gateway Status
	GatewayStatus string `json:"gateway_status,omitempty"`

	// Flag to identify if this VM is a gateway node
	IsGateway bool `json:"is_gateway,omitempty"`

	// Flag to identify if this VM is an active gateway node
	IsGatewayActive bool `json:"is_gateway_active,omitempty"`

	// Logical Switch display name
	LogicalSwitchDisplayName string `json:"logical_switch_display_name,omitempty"`

	// Logical Switch ID
	LogicalSwitchId string `json:"logical_switch_id,omitempty"`

	// Indicate if vm is managed by NSX or not
	ManagedByNsx bool `json:"managed_by_nsx,omitempty"`

	// IP address provided by Nsx
	NsxIp string `json:"nsx_ip,omitempty"`

	// Operating system details
	OsDetails string `json:"os_details,omitempty"`

	// Operating system of the virtual machine
	OsType string `json:"os_type,omitempty"`

	// Private IP address of the virtual machine
	PrivateIp string `json:"private_ip,omitempty"`

	// Public IP address of the virtual machine
	PublicIp string `json:"public_ip,omitempty"`

	// Quarantine State of VM
	QuarantineState string `json:"quarantine_state,omitempty"`
}