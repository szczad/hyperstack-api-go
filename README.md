# Go API client for hyperstack

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Generator version: 7.7.0-SNAPSHOT
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import hyperstack "github.com/szczad/hyperstack-api-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `hyperstack.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), hyperstack.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `hyperstack.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), hyperstack.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `hyperstack.ContextOperationServerIndices` and `hyperstack.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), hyperstack.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), hyperstack.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AliveAPI* | [**GetAlive**](docs/AliveAPI.md#getalive) | **Get** /billing/alive | GET: Alive
*ApiKeyAPI* | [**GenerateAPIKey**](docs/ApiKeyAPI.md#generateapikey) | **Post** /api-key/generate | Generate API Key
*ApiKeyAPI* | [**GetAPIKey**](docs/ApiKeyAPI.md#getapikey) | **Get** /api-key | Get API Key
*AssigningMemberRoleAPI* | [**AssignRBACRoles**](docs/AssigningMemberRoleAPI.md#assignrbacroles) | **Put** /auth/users/{user_id}/assign-roles | Assign RBAC Roles
*AssigningMemberRoleAPI* | [**RemoveRoleFromAUser**](docs/AssigningMemberRoleAPI.md#removerolefromauser) | **Delete** /auth/users/{user_id}/roles | Remove role from a user
*AuthAPI* | [**AuthUserInformation**](docs/AuthAPI.md#authuserinformation) | **Get** /auth/me | Get me information
*BillingAPI* | [**GetLastDayCost**](docs/BillingAPI.md#getlastdaycost) | **Get** /billing/billing/last-day-cost | GET: Last Day Cost
*BillingAPI* | [**GetUsage**](docs/BillingAPI.md#getusage) | **Get** /billing/billing/usage | GET: Billing usage
*CallbacksAPI* | [**AttachCallbackToVirtualMachine**](docs/CallbacksAPI.md#attachcallbacktovirtualmachine) | **Post** /core/virtual-machines/{id}/attach-callback | Attach callback to virtual machine
*CallbacksAPI* | [**AttachCallbackToVolume**](docs/CallbacksAPI.md#attachcallbacktovolume) | **Post** /core/volumes/{id}/attach-callback | Attach callback to volume
*CallbacksAPI* | [**DeleteVirtualMachineCallback**](docs/CallbacksAPI.md#deletevirtualmachinecallback) | **Delete** /core/virtual-machines/{id}/delete-callback | Delete virtual machine callback
*CallbacksAPI* | [**DeleteVolumeCallback**](docs/CallbacksAPI.md#deletevolumecallback) | **Delete** /core/volumes/{id}/delete-callback | Delete volume callback
*CallbacksAPI* | [**UpdateVirtualMachineCallback**](docs/CallbacksAPI.md#updatevirtualmachinecallback) | **Put** /core/virtual-machines/{id}/update-callback | Update virtual machine callback
*CallbacksAPI* | [**UpdateVolumeCallback**](docs/CallbacksAPI.md#updatevolumecallback) | **Put** /core/volumes/{id}/update-callback | Update volume callback
*ComplianceAPI* | [**CreateCompliance**](docs/ComplianceAPI.md#createcompliance) | **Post** /core/compliance | Create compliance
*ComplianceAPI* | [**DeleteACompliance**](docs/ComplianceAPI.md#deleteacompliance) | **Delete** /core/compliance/{gpu_model} | Delete a compliance
*ComplianceAPI* | [**RetrieveCompliance**](docs/ComplianceAPI.md#retrievecompliance) | **Get** /core/compliance | Retrieve GPU compliance
*ComplianceAPI* | [**UpdateACompliance**](docs/ComplianceAPI.md#updateacompliance) | **Put** /core/compliance | Update a compliance
*CreditAPI* | [**CheckBalanceAsAnOrganization**](docs/CreditAPI.md#checkbalanceasanorganization) | **Get** /billing/user-credit/credit | GET: View credit and threshold
*CustomerContractAPI* | [**DetailsOfContractByIDForCustomer**](docs/CustomerContractAPI.md#detailsofcontractbyidforcustomer) | **Get** /pricebook/contracts/{contract_id} | Details of Contract by ID for Customer
*CustomerContractAPI* | [**ListContractsForCustomer**](docs/CustomerContractAPI.md#listcontractsforcustomer) | **Get** /pricebook/contracts | List Contracts for Customer
*DashboardAPI* | [**RetrieveDashboard**](docs/DashboardAPI.md#retrievedashboard) | **Get** /core/dashboard | Retrieve Dashboard
*DeploymentAPI* | [**DeleteDeployment**](docs/DeploymentAPI.md#deletedeployment) | **Delete** /core/marketplace/deployments/{id} | Delete Deployment
*DeploymentAPI* | [**DetailsOfDeploymentByID**](docs/DeploymentAPI.md#detailsofdeploymentbyid) | **Get** /core/marketplace/deployments/{id} | Details of Deployment by ID
*DeploymentAPI* | [**ListDeployments**](docs/DeploymentAPI.md#listdeployments) | **Get** /core/marketplace/deployments | List Deployments
*DeploymentAPI* | [**StartDeployment**](docs/DeploymentAPI.md#startdeployment) | **Post** /core/marketplace/deployments | Start Deployment
*EnvironmentAPI* | [**CreateEnvironment**](docs/EnvironmentAPI.md#createenvironment) | **Post** /core/environments | Create environment
*EnvironmentAPI* | [**DeleteEnvironment**](docs/EnvironmentAPI.md#deleteenvironment) | **Delete** /core/environments/{id} | Delete environment
*EnvironmentAPI* | [**ListEnvironments**](docs/EnvironmentAPI.md#listenvironments) | **Get** /core/environments | List environments
*EnvironmentAPI* | [**RetrieveEnvironment**](docs/EnvironmentAPI.md#retrieveenvironment) | **Get** /core/environments/{id} | Retrieve environment
*EnvironmentAPI* | [**UpdateEnvironment**](docs/EnvironmentAPI.md#updateenvironment) | **Put** /core/environments/{id} | Update environment
*FirewallAttachmentAPI* | [**AttachFirewallToVMs**](docs/FirewallAttachmentAPI.md#attachfirewalltovms) | **Post** /core/firewalls/{firewall_id}/update-attachments | Attach Firewalls to VMs
*FirewallsAPI* | [**AddRulesToFirewall**](docs/FirewallsAPI.md#addrulestofirewall) | **Post** /core/firewalls/{firewall_id}/firewall-rules | Add Rules to Firewall
*FirewallsAPI* | [**CreateFirewall**](docs/FirewallsAPI.md#createfirewall) | **Post** /core/firewalls | Create Firewall
*FirewallsAPI* | [**DeleteFirewall**](docs/FirewallsAPI.md#deletefirewall) | **Delete** /core/firewalls/{id} | Delete Firewall
*FirewallsAPI* | [**DeleteFirewallRulesFromFirewall**](docs/FirewallsAPI.md#deletefirewallrulesfromfirewall) | **Delete** /core/firewalls/{firewall_id}/firewall-rules/{firewall_rule_id} | Delete Firewall Rules from Firewall
*FirewallsAPI* | [**DetailsOfFirewallByID**](docs/FirewallsAPI.md#detailsoffirewallbyid) | **Get** /core/firewalls/{id} | Details of Firewall by ID
*FirewallsAPI* | [**RetrieveFirewalls**](docs/FirewallsAPI.md#retrievefirewalls) | **Get** /core/firewalls | Retrieve Firewalls
*FlavorAPI* | [**ListFlavors**](docs/FlavorAPI.md#listflavors) | **Get** /core/flavors | List flavors
*FloatingIpAPI* | [**AttachPublicIPToVirtualMachine**](docs/FloatingIpAPI.md#attachpubliciptovirtualmachine) | **Post** /core/virtual-machines/{id}/attach-floatingip | Attach public IP to virtual machine
*FloatingIpAPI* | [**DetachPublicIPFromVirtualMachine**](docs/FloatingIpAPI.md#detachpublicipfromvirtualmachine) | **Post** /core/virtual-machines/{id}/detach-floatingip | Detach public IP from virtual machine
*GpuAPI* | [**ListGPUs**](docs/GpuAPI.md#listgpus) | **Get** /core/gpus | List GPUs
*ImageAPI* | [**ListImages**](docs/ImageAPI.md#listimages) | **Get** /core/images | List images
*InviteAPI* | [**DeleteInvite**](docs/InviteAPI.md#deleteinvite) | **Delete** /auth/invites/{id} | Delete Invite
*InviteAPI* | [**InviteAnUserToOrganization**](docs/InviteAPI.md#inviteanusertoorganization) | **Post** /auth/invites | Invite an user to organization
*InviteAPI* | [**ListInvites**](docs/InviteAPI.md#listinvites) | **Get** /auth/invites | List Invites
*KeypairAPI* | [**DeleteKeyPair**](docs/KeypairAPI.md#deletekeypair) | **Delete** /core/keypair/{id} | Delete key pair
*KeypairAPI* | [**ImportKeyPair**](docs/KeypairAPI.md#importkeypair) | **Post** /core/keypairs | Import key pair
*KeypairAPI* | [**ListKeyPairs**](docs/KeypairAPI.md#listkeypairs) | **Get** /core/keypairs | List key pairs
*KeypairAPI* | [**UpdateKeyPairName**](docs/KeypairAPI.md#updatekeypairname) | **Put** /core/keypair/{id} | Update key pair name
*OrganizationAPI* | [**GetOrganizationInfo**](docs/OrganizationAPI.md#getorganizationinfo) | **Get** /auth/organizations | Organization Info
*OrganizationAPI* | [**RemoveAMemberFromOrganization**](docs/OrganizationAPI.md#removeamemberfromorganization) | **Post** /auth/organizations/remove-member | Remove a member from organization
*OrganizationAPI* | [**UpdateOrganizationInfo**](docs/OrganizationAPI.md#updateorganizationinfo) | **Put** /auth/organizations/update | Update organization info
*PaymentAPI* | [**GetDetails**](docs/PaymentAPI.md#getdetails) | **Get** /billing/payment/payment-details | GET: View payment details
*PaymentAPI* | [**PostPayment**](docs/PaymentAPI.md#postpayment) | **Post** /billing/payment/payment-initiate | POST: Initiate payment
*PermissionAPI* | [**ListPermissions**](docs/PermissionAPI.md#listpermissions) | **Get** /auth/permissions | List Permissions
*PolicyAPI* | [**ListPolicies**](docs/PolicyAPI.md#listpolicies) | **Get** /auth/policies | List Policies
*PricebookAPI* | [**RetrivePricebook**](docs/PricebookAPI.md#retrivepricebook) | **Get** /pricebook | 
*ProfileAPI* | [**CreateProfile**](docs/ProfileAPI.md#createprofile) | **Post** /core/profiles | Create profile
*ProfileAPI* | [**DeleteProfile**](docs/ProfileAPI.md#deleteprofile) | **Delete** /core/profiles/{id} | Delete profile
*ProfileAPI* | [**ListProfiles**](docs/ProfileAPI.md#listprofiles) | **Get** /core/profiles | List profiles
*ProfileAPI* | [**RetrieveProfileDetails**](docs/ProfileAPI.md#retrieveprofiledetails) | **Get** /core/profiles/{id} | Retrieve profile details
*RbacRoleAPI* | [**CreateRBACRole**](docs/RbacRoleAPI.md#createrbacrole) | **Post** /auth/roles | Create RBAC Role
*RbacRoleAPI* | [**DeleteARBACRole**](docs/RbacRoleAPI.md#deletearbacrole) | **Delete** /auth/roles/{id} | Delete a RBAC Role
*RbacRoleAPI* | [**GetARBACRoleDetail**](docs/RbacRoleAPI.md#getarbacroledetail) | **Get** /auth/roles/{id} | Get a RBAC Role Detail
*RbacRoleAPI* | [**ListRBACRoles**](docs/RbacRoleAPI.md#listrbacroles) | **Get** /auth/roles | List RBAC Roles
*RbacRoleAPI* | [**UpdateARBACRole**](docs/RbacRoleAPI.md#updatearbacrole) | **Put** /auth/roles/{id} | Update a RBAC Role
*RegionAPI* | [**ListRegions**](docs/RegionAPI.md#listregions) | **Get** /core/regions | List regions
*SecurityRulesAPI* | [**ListFirewallRuleProtocols**](docs/SecurityRulesAPI.md#listfirewallruleprotocols) | **Get** /core/sg-rules-protocols | List firewall rule protocols
*StockAPI* | [**RetrieveGPUStocks**](docs/StockAPI.md#retrievegpustocks) | **Get** /core/stocks | 
*TemplateAPI* | [**CreateTemplate**](docs/TemplateAPI.md#createtemplate) | **Post** /core/marketplace/templates | Create template
*TemplateAPI* | [**DeleteTemplate**](docs/TemplateAPI.md#deletetemplate) | **Delete** /core/marketplace/templates/{id} | Delete template
*TemplateAPI* | [**ListTemplates**](docs/TemplateAPI.md#listtemplates) | **Get** /core/marketplace/templates | List templates
*TemplateAPI* | [**RetrieveTemplateDetails**](docs/TemplateAPI.md#retrievetemplatedetails) | **Get** /core/marketplace/templates/{id} | Retrieve template details
*TemplateAPI* | [**UpdateTemplate**](docs/TemplateAPI.md#updatetemplate) | **Put** /core/marketplace/templates/{id} | Update template
*UserAPI* | [**GetUser**](docs/UserAPI.md#getuser) | **Get** /billing/user/info | GET: Fetch User Info
*UserAPI* | [**PostUser**](docs/UserAPI.md#postuser) | **Post** /billing/user/info | POST: Insert user info
*UserAPI* | [**UpdateUserInfo**](docs/UserAPI.md#updateuserinfo) | **Put** /billing/user/info | PUT: Update user info
*UserDetailChoiceAPI* | [**RetrieveDefaultFlavorsAndImagesForUser**](docs/UserDetailChoiceAPI.md#retrievedefaultflavorsandimagesforuser) | **Get** /core/user/resources/defaults | Retrieve default flavors and images for user
*UserPermissionAPI* | [**ListCurrentUserPermissions**](docs/UserPermissionAPI.md#listcurrentuserpermissions) | **Get** /auth/users/me/permissions | List Current User Permissions
*UserPermissionAPI* | [**ListUserPermissions**](docs/UserPermissionAPI.md#listuserpermissions) | **Get** /auth/users/{id}/permissions | List User Permissions
*VirtualMachineAPI* | [**AddFirewallRuleToVirtualMachine**](docs/VirtualMachineAPI.md#addfirewallruletovirtualmachine) | **Post** /core/virtual-machines/{id}/sg-rules | Add firewall rule to virtual machine
*VirtualMachineAPI* | [**AttachFirewallsToAVM**](docs/VirtualMachineAPI.md#attachfirewallstoavm) | **Post** /core/virtual-machines/{vm_id}/attach-firewalls | Attach Firewalls to a VM
*VirtualMachineAPI* | [**CreateVirtualMachine**](docs/VirtualMachineAPI.md#createvirtualmachine) | **Post** /core/virtual-machines | Create virtual machine
*VirtualMachineAPI* | [**DeleteFirewallRuleFromVirtualMachine**](docs/VirtualMachineAPI.md#deletefirewallrulefromvirtualmachine) | **Delete** /core/virtual-machines/{virtual_machine_id}/sg-rules/{sg_rule_id} | Delete firewall rule from virtual machine
*VirtualMachineAPI* | [**DeleteVirtualMachine**](docs/VirtualMachineAPI.md#deletevirtualmachine) | **Delete** /core/virtual-machines/{id} | Delete virtual machine
*VirtualMachineAPI* | [**EditLabelsOfAnExistingVM**](docs/VirtualMachineAPI.md#editlabelsofanexistingvm) | **Put** /core/virtual-machines/{virtual_machine_id}/label | Edit labels of an existing VM
*VirtualMachineAPI* | [**HardRebootVirtualMachine**](docs/VirtualMachineAPI.md#hardrebootvirtualmachine) | **Get** /core/virtual-machines/{id}/hard-reboot | Hard reboot virtual machine
*VirtualMachineAPI* | [**HibernateVirtualMachine**](docs/VirtualMachineAPI.md#hibernatevirtualmachine) | **Get** /core/virtual-machines/{virtual_machine_id}/hibernate | Hibernate virtual machine
*VirtualMachineAPI* | [**ListVirtualMachines**](docs/VirtualMachineAPI.md#listvirtualmachines) | **Get** /core/virtual-machines | List all virtual machines
*VirtualMachineAPI* | [**ResizeVirtualMachine**](docs/VirtualMachineAPI.md#resizevirtualmachine) | **Post** /core/virtual-machines/{virtual_machine_id}/resize | Resize virtual machine
*VirtualMachineAPI* | [**RestoreVirtualMachineFromHibernation**](docs/VirtualMachineAPI.md#restorevirtualmachinefromhibernation) | **Get** /core/virtual-machines/{virtual_machine_id}/hibernate-restore | Restore virtual machine from hibernation
*VirtualMachineAPI* | [**RetrieveVirtualMachineDetails**](docs/VirtualMachineAPI.md#retrievevirtualmachinedetails) | **Get** /core/virtual-machines/{id} | Retrieve virtual machine details
*VirtualMachineAPI* | [**RetrieveVirtualMachinePerformanceMetrics**](docs/VirtualMachineAPI.md#retrievevirtualmachineperformancemetrics) | **Get** /core/virtual-machines/{virtual_machine_id}/metrics | Retrieve virtual machine performance metrics
*VirtualMachineAPI* | [**RetrieveVirtualMachinesAssociatedWithAContract**](docs/VirtualMachineAPI.md#retrievevirtualmachinesassociatedwithacontract) | **Get** /core/virtual-machines/contract/{contract_id}/virtual-machines | Retrieve virtual machines associated with a contract
*VirtualMachineAPI* | [**StartVirtualMachine**](docs/VirtualMachineAPI.md#startvirtualmachine) | **Get** /core/virtual-machines/{id}/start | Start virtual machine
*VirtualMachineAPI* | [**StopVirtualMachine**](docs/VirtualMachineAPI.md#stopvirtualmachine) | **Get** /core/virtual-machines/{id}/stop | Stop virtual machine
*VirtualMachineEventsAPI* | [**ListVirtualMachineEvents**](docs/VirtualMachineEventsAPI.md#listvirtualmachineevents) | **Get** /core/virtual-machines/{virtual_machine_id}/events | List virtual machine events
*VncUrlAPI* | [**GetVNCConsoleLink**](docs/VncUrlAPI.md#getvncconsolelink) | **Get** /core/virtual-machines/{virtual_machine_id}/console/{job_id} | Get VNC Console Link
*VncUrlAPI* | [**RequestConsole**](docs/VncUrlAPI.md#requestconsole) | **Get** /core/virtual-machines/{id}/request-console | Request Instance Console
*VolumeAPI* | [**CreateVolume**](docs/VolumeAPI.md#createvolume) | **Post** /core/volumes | Create volume
*VolumeAPI* | [**DeleteVolume**](docs/VolumeAPI.md#deletevolume) | **Delete** /core/volumes/{id} | Delete volume
*VolumeAPI* | [**ListVolumeTypes**](docs/VolumeAPI.md#listvolumetypes) | **Get** /core/volume-types | List volume types
*VolumeAPI* | [**ListVolumes**](docs/VolumeAPI.md#listvolumes) | **Get** /core/volumes | List volumes
*VolumeAttachmentAPI* | [**AttachVolumesToVirtualMachine**](docs/VolumeAttachmentAPI.md#attachvolumestovirtualmachine) | **Post** /core/virtual-machines/{virtual_machine_id}/attach-volumes | Attach volumes to virtual machine
*VolumeAttachmentAPI* | [**DetachVolumesFromVirtualMachine**](docs/VolumeAttachmentAPI.md#detachvolumesfromvirtualmachine) | **Post** /core/virtual-machines/{virtual_machine_id}/detach-volumes | Detach volumes from virtual machine


## Documentation For Models

 - [AddUpdateFlavorOrganizationPayload](docs/AddUpdateFlavorOrganizationPayload.md)
 - [AddUserInfoSuccessResponseModel](docs/AddUserInfoSuccessResponseModel.md)
 - [AdminAddUpdateImageOrganizationPayload](docs/AdminAddUpdateImageOrganizationPayload.md)
 - [AdminBootstrapEnvironmentPayload](docs/AdminBootstrapEnvironmentPayload.md)
 - [AdminClusterResource](docs/AdminClusterResource.md)
 - [AdminContainerResource](docs/AdminContainerResource.md)
 - [AdminContractEventFields](docs/AdminContractEventFields.md)
 - [AdminContractFields](docs/AdminContractFields.md)
 - [AdminCountResourcesOrganization](docs/AdminCountResourcesOrganization.md)
 - [AdminCountResourcesOrganizations](docs/AdminCountResourcesOrganizations.md)
 - [AdminCreateContractResponseModel](docs/AdminCreateContractResponseModel.md)
 - [AdminEnvrionmentResources](docs/AdminEnvrionmentResources.md)
 - [AdminFlavorDetailFields](docs/AdminFlavorDetailFields.md)
 - [AdminFlavorDetailNodeFields](docs/AdminFlavorDetailNodeFields.md)
 - [AdminFlavorResource](docs/AdminFlavorResource.md)
 - [AdminFlavorResourcesList](docs/AdminFlavorResourcesList.md)
 - [AdminGetContractDetailFields](docs/AdminGetContractDetailFields.md)
 - [AdminGetVersionResponse](docs/AdminGetVersionResponse.md)
 - [AdminHibernationRestorationPayloadModel](docs/AdminHibernationRestorationPayloadModel.md)
 - [AdminImageAdminFields](docs/AdminImageAdminFields.md)
 - [AdminImageAdminResponse](docs/AdminImageAdminResponse.md)
 - [AdminImageAdminResponseImage](docs/AdminImageAdminResponseImage.md)
 - [AdminImageListAdminResponse](docs/AdminImageListAdminResponse.md)
 - [AdminInstanceResources](docs/AdminInstanceResources.md)
 - [AdminNodeResource](docs/AdminNodeResource.md)
 - [AdminOrganizationResourceList](docs/AdminOrganizationResourceList.md)
 - [AdminOrganizationResourceResponse](docs/AdminOrganizationResourceResponse.md)
 - [AdminOrganizationResources](docs/AdminOrganizationResources.md)
 - [AdminOrganizationResponseModel](docs/AdminOrganizationResponseModel.md)
 - [AdminOrganizationSummaryFields](docs/AdminOrganizationSummaryFields.md)
 - [AdminOrganizationsResponseModel](docs/AdminOrganizationsResponseModel.md)
 - [AdminOrganizationsSummaryResponseModel](docs/AdminOrganizationsSummaryResponseModel.md)
 - [AdminUserFields](docs/AdminUserFields.md)
 - [AdminUserResponseModel](docs/AdminUserResponseModel.md)
 - [AdminUsersResponseModel](docs/AdminUsersResponseModel.md)
 - [AdminVersionResponseModel](docs/AdminVersionResponseModel.md)
 - [AdminVolumeResource](docs/AdminVolumeResource.md)
 - [Adminaddorganizationpayload](docs/Adminaddorganizationpayload.md)
 - [Admincreditpostpayload](docs/Admincreditpostpayload.md)
 - [Adminpaymenthistoryfields](docs/Adminpaymenthistoryfields.md)
 - [Adminpaymenthistoryresponse](docs/Adminpaymenthistoryresponse.md)
 - [Adminrechargehistoryfields](docs/Adminrechargehistoryfields.md)
 - [Adminrechargehistoryresponse](docs/Adminrechargehistoryresponse.md)
 - [Adminthresholdpostpayload](docs/Adminthresholdpostpayload.md)
 - [ApiKeyFields](docs/ApiKeyFields.md)
 - [ApiKeyVerifyFields](docs/ApiKeyVerifyFields.md)
 - [AssignRbacRolePayload](docs/AssignRbacRolePayload.md)
 - [AttachCallbackPayload](docs/AttachCallbackPayload.md)
 - [AttachCallbackResponse](docs/AttachCallbackResponse.md)
 - [AttachFirewallWithVM](docs/AttachFirewallWithVM.md)
 - [AttachFirewallsToVMPayload](docs/AttachFirewallsToVMPayload.md)
 - [AttachVolumeFields](docs/AttachVolumeFields.md)
 - [AttachVolumes](docs/AttachVolumes.md)
 - [AttachVolumesPayload](docs/AttachVolumesPayload.md)
 - [AuthGetTokenResponseModel](docs/AuthGetTokenResponseModel.md)
 - [AuthRequestLoginFields](docs/AuthRequestLoginFields.md)
 - [AuthRequestLoginResponseModel](docs/AuthRequestLoginResponseModel.md)
 - [AuthUserFields](docs/AuthUserFields.md)
 - [AuthUserInfoResponseModel](docs/AuthUserInfoResponseModel.md)
 - [BillingResponse](docs/BillingResponse.md)
 - [Billingimmuneresources](docs/Billingimmuneresources.md)
 - [Billingimmuneresourcesresponse](docs/Billingimmuneresourcesresponse.md)
 - [Billingmetricesfields](docs/Billingmetricesfields.md)
 - [Billingmetricesresponse](docs/Billingmetricesresponse.md)
 - [CommonResponseModel](docs/CommonResponseModel.md)
 - [ComplianceFields](docs/ComplianceFields.md)
 - [ComplianceModelFields](docs/ComplianceModelFields.md)
 - [CompliancePayload](docs/CompliancePayload.md)
 - [ComplianceResponse](docs/ComplianceResponse.md)
 - [ContainerOverviewFields](docs/ContainerOverviewFields.md)
 - [ContractInstanceFields](docs/ContractInstanceFields.md)
 - [ContractInstancesResponse](docs/ContractInstancesResponse.md)
 - [CreateContarctFields](docs/CreateContarctFields.md)
 - [CreateContractPayload](docs/CreateContractPayload.md)
 - [CreateDiscountResponse](docs/CreateDiscountResponse.md)
 - [CreateDiscountsPayload](docs/CreateDiscountsPayload.md)
 - [CreateEnvironment](docs/CreateEnvironment.md)
 - [CreateFirewallPayload](docs/CreateFirewallPayload.md)
 - [CreateFirewallRulePayload](docs/CreateFirewallRulePayload.md)
 - [CreateGPU](docs/CreateGPU.md)
 - [CreateInstancesPayload](docs/CreateInstancesPayload.md)
 - [CreateProfilePayload](docs/CreateProfilePayload.md)
 - [CreateProfileResponse](docs/CreateProfileResponse.md)
 - [CreateSecurityRulePayload](docs/CreateSecurityRulePayload.md)
 - [CreateUpdateComplianceResponse](docs/CreateUpdateComplianceResponse.md)
 - [CreateUpdatePermissionPayload](docs/CreateUpdatePermissionPayload.md)
 - [CreateUpdatePermissionResponseModel](docs/CreateUpdatePermissionResponseModel.md)
 - [CreateUpdatePolicyPayload](docs/CreateUpdatePolicyPayload.md)
 - [CreateUpdatePolicyResponseModel](docs/CreateUpdatePolicyResponseModel.md)
 - [CreateUpdateRbacRolePayload](docs/CreateUpdateRbacRolePayload.md)
 - [CreateVolumePayload](docs/CreateVolumePayload.md)
 - [CustomerContractDetailResponseModel](docs/CustomerContractDetailResponseModel.md)
 - [CustomerContractFields](docs/CustomerContractFields.md)
 - [CustomerFields](docs/CustomerFields.md)
 - [CustomerPayload](docs/CustomerPayload.md)
 - [DashboardInfoResponse](docs/DashboardInfoResponse.md)
 - [DeploymentFields](docs/DeploymentFields.md)
 - [DeploymentFieldsforstartdeployments](docs/DeploymentFieldsforstartdeployments.md)
 - [Deployments](docs/Deployments.md)
 - [DetachVolumes](docs/DetachVolumes.md)
 - [DetachVolumesPayload](docs/DetachVolumesPayload.md)
 - [DiscountEntityModel](docs/DiscountEntityModel.md)
 - [DiscountFields](docs/DiscountFields.md)
 - [DiscountPlanFields](docs/DiscountPlanFields.md)
 - [DiscountResourceFields](docs/DiscountResourceFields.md)
 - [EditlabelofanexistingVMPayload](docs/EditlabelofanexistingVMPayload.md)
 - [Environment](docs/Environment.md)
 - [EnvironmentFields](docs/EnvironmentFields.md)
 - [EnvironmentFieldsforVolume](docs/EnvironmentFieldsforVolume.md)
 - [Environments](docs/Environments.md)
 - [ErrorResponseModel](docs/ErrorResponseModel.md)
 - [Excludebillingpostpayload](docs/Excludebillingpostpayload.md)
 - [Excludebillingpostresponse](docs/Excludebillingpostresponse.md)
 - [FirewallAttachmentModel](docs/FirewallAttachmentModel.md)
 - [FirewallAttachmentVMModel](docs/FirewallAttachmentVMModel.md)
 - [FirewallDetailFields](docs/FirewallDetailFields.md)
 - [FirewallDetailResponse](docs/FirewallDetailResponse.md)
 - [FirewallEnvironmentFields](docs/FirewallEnvironmentFields.md)
 - [FirewallFields](docs/FirewallFields.md)
 - [FirewallResponse](docs/FirewallResponse.md)
 - [FirewallRule](docs/FirewallRule.md)
 - [FirewallsListResponse](docs/FirewallsListResponse.md)
 - [FlavorAdminResponse](docs/FlavorAdminResponse.md)
 - [FlavorAdminResponseFlavors](docs/FlavorAdminResponseFlavors.md)
 - [FlavorDetailResponse](docs/FlavorDetailResponse.md)
 - [FlavorFields](docs/FlavorFields.md)
 - [FlavorItemGetResponse](docs/FlavorItemGetResponse.md)
 - [FlavorListResponse](docs/FlavorListResponse.md)
 - [FlavorObjectFields](docs/FlavorObjectFields.md)
 - [FlavorPayload](docs/FlavorPayload.md)
 - [FlavorResponse](docs/FlavorResponse.md)
 - [FlavorVMFields](docs/FlavorVMFields.md)
 - [FlavorVMsResponse](docs/FlavorVMsResponse.md)
 - [FutureNodeModel](docs/FutureNodeModel.md)
 - [FutureNodeResponseModel](docs/FutureNodeResponseModel.md)
 - [FutureNodeStockModel](docs/FutureNodeStockModel.md)
 - [FutureNodeUpdateModel](docs/FutureNodeUpdateModel.md)
 - [FutureNodesStockModel](docs/FutureNodesStockModel.md)
 - [GPU](docs/GPU.md)
 - [GPUFields](docs/GPUFields.md)
 - [GPUList](docs/GPUList.md)
 - [GPURegionFields](docs/GPURegionFields.md)
 - [GenerateApiKeyResponseModel](docs/GenerateApiKeyResponseModel.md)
 - [GetAllDiscountForAllOrganizationResponse](docs/GetAllDiscountForAllOrganizationResponse.md)
 - [GetAllDiscountsFields](docs/GetAllDiscountsFields.md)
 - [GetApiKeyResponseModel](docs/GetApiKeyResponseModel.md)
 - [GetContractDetailResponseModel](docs/GetContractDetailResponseModel.md)
 - [GetContractEventsResponseModel](docs/GetContractEventsResponseModel.md)
 - [GetContractsListResponseModel](docs/GetContractsListResponseModel.md)
 - [GetCustomerContractsListResponseModel](docs/GetCustomerContractsListResponseModel.md)
 - [GetDiscountDetailResponse](docs/GetDiscountDetailResponse.md)
 - [GetDiscountResponse](docs/GetDiscountResponse.md)
 - [GetEntityDiscountDetailResponse](docs/GetEntityDiscountDetailResponse.md)
 - [GetInvitesResponseModel](docs/GetInvitesResponseModel.md)
 - [GetOrganizationResponseModel](docs/GetOrganizationResponseModel.md)
 - [GetPermissionsResponseModel](docs/GetPermissionsResponseModel.md)
 - [GetPoliciesResponseModel](docs/GetPoliciesResponseModel.md)
 - [GetRbacRolesResponseModel](docs/GetRbacRolesResponseModel.md)
 - [GetTokenPayload](docs/GetTokenPayload.md)
 - [GetUserPermissionsResponseModel](docs/GetUserPermissionsResponseModel.md)
 - [GetVersionResponse](docs/GetVersionResponse.md)
 - [Getcreditandthresholdinfo](docs/Getcreditandthresholdinfo.md)
 - [Getcreditandthresholdinfoinresponse](docs/Getcreditandthresholdinfoinresponse.md)
 - [ImageFields](docs/ImageFields.md)
 - [ImageGetResponse](docs/ImageGetResponse.md)
 - [ImageLogos](docs/ImageLogos.md)
 - [Images](docs/Images.md)
 - [ImportKeypairPayload](docs/ImportKeypairPayload.md)
 - [ImportKeypairResponse](docs/ImportKeypairResponse.md)
 - [InfrahubResourceObjectResponse](docs/InfrahubResourceObjectResponse.md)
 - [InsertDiscountPlanFields](docs/InsertDiscountPlanFields.md)
 - [Instance](docs/Instance.md)
 - [InstanceAdmin](docs/InstanceAdmin.md)
 - [InstanceAdminFields](docs/InstanceAdminFields.md)
 - [InstanceEnvironmentFields](docs/InstanceEnvironmentFields.md)
 - [InstanceEvents](docs/InstanceEvents.md)
 - [InstanceEventsFields](docs/InstanceEventsFields.md)
 - [InstanceFields](docs/InstanceFields.md)
 - [InstanceFlavorFields](docs/InstanceFlavorFields.md)
 - [InstanceImageFields](docs/InstanceImageFields.md)
 - [InstanceKeypairFields](docs/InstanceKeypairFields.md)
 - [InstanceOverviewFields](docs/InstanceOverviewFields.md)
 - [InstanceResizePayload](docs/InstanceResizePayload.md)
 - [Instances](docs/Instances.md)
 - [InstancesAdmin](docs/InstancesAdmin.md)
 - [InstancesSummaryAdmin](docs/InstancesSummaryAdmin.md)
 - [InstancesSummaryFields](docs/InstancesSummaryFields.md)
 - [InternalEnvironmentFields](docs/InternalEnvironmentFields.md)
 - [InternalInstanceFields](docs/InternalInstanceFields.md)
 - [InternalInstanceFlavorFields](docs/InternalInstanceFlavorFields.md)
 - [InternalInstanceImageFields](docs/InternalInstanceImageFields.md)
 - [InternalInstanceKeypairFields](docs/InternalInstanceKeypairFields.md)
 - [InternalInstancesResponse](docs/InternalInstancesResponse.md)
 - [InternalSecurityRulesFieldsForInstance](docs/InternalSecurityRulesFieldsForInstance.md)
 - [InternalVolumeAttachmentFields](docs/InternalVolumeAttachmentFields.md)
 - [InternalVolumeFields](docs/InternalVolumeFields.md)
 - [InternalVolumesResponse](docs/InternalVolumesResponse.md)
 - [InviteFields](docs/InviteFields.md)
 - [InviteUserPayload](docs/InviteUserPayload.md)
 - [InviteUserResponseModel](docs/InviteUserResponseModel.md)
 - [KeypairFields](docs/KeypairFields.md)
 - [Keypairs](docs/Keypairs.md)
 - [LableResonse](docs/LableResonse.md)
 - [Lastdaycostfields](docs/Lastdaycostfields.md)
 - [Lastdaycostresponse](docs/Lastdaycostresponse.md)
 - [LogoGetResponse](docs/LogoGetResponse.md)
 - [LogoutPayload](docs/LogoutPayload.md)
 - [MetricItemFields](docs/MetricItemFields.md)
 - [MetricsFields](docs/MetricsFields.md)
 - [NewConfigurationsResponse](docs/NewConfigurationsResponse.md)
 - [NewModelResponse](docs/NewModelResponse.md)
 - [NewStockResponse](docs/NewStockResponse.md)
 - [NewStockRetriveResponse](docs/NewStockRetriveResponse.md)
 - [NewStockUpdateResponseModel](docs/NewStockUpdateResponseModel.md)
 - [NodeModel](docs/NodeModel.md)
 - [NodePayloadModel](docs/NodePayloadModel.md)
 - [NodePowerUsageModel](docs/NodePowerUsageModel.md)
 - [NodeResponseModel](docs/NodeResponseModel.md)
 - [NodeStockPayloadModel](docs/NodeStockPayloadModel.md)
 - [NodeStockResponseModel](docs/NodeStockResponseModel.md)
 - [NodeStocksPayload](docs/NodeStocksPayload.md)
 - [OrganizationFields](docs/OrganizationFields.md)
 - [OrganizationObjectResponse](docs/OrganizationObjectResponse.md)
 - [OrganizationUserResponseModel](docs/OrganizationUserResponseModel.md)
 - [OverviewInfo](docs/OverviewInfo.md)
 - [PaymentDetailsFields](docs/PaymentDetailsFields.md)
 - [PaymentDetailsResponse](docs/PaymentDetailsResponse.md)
 - [PaymentInitiateFields](docs/PaymentInitiateFields.md)
 - [PaymentInitiatePayload](docs/PaymentInitiatePayload.md)
 - [PaymentInitiateResponse](docs/PaymentInitiateResponse.md)
 - [PermissionFields](docs/PermissionFields.md)
 - [PolicyFields](docs/PolicyFields.md)
 - [PolicyPermissionFields](docs/PolicyPermissionFields.md)
 - [PowerUsageModel](docs/PowerUsageModel.md)
 - [PricebookModel](docs/PricebookModel.md)
 - [PricebookResourceObjectResponse](docs/PricebookResourceObjectResponse.md)
 - [ProfileFields](docs/ProfileFields.md)
 - [ProfileListResponse](docs/ProfileListResponse.md)
 - [ProfileObjectFields](docs/ProfileObjectFields.md)
 - [RbacRoleDetailResponseModel](docs/RbacRoleDetailResponseModel.md)
 - [RbacRoleField](docs/RbacRoleField.md)
 - [RbacRoleFields](docs/RbacRoleFields.md)
 - [RefreshTokenPayload](docs/RefreshTokenPayload.md)
 - [RegionFields](docs/RegionFields.md)
 - [RegionPayload](docs/RegionPayload.md)
 - [RegionResponse](docs/RegionResponse.md)
 - [Regions](docs/Regions.md)
 - [RemoveMemberFromOrganizationResponseModel](docs/RemoveMemberFromOrganizationResponseModel.md)
 - [RemoveMemberPayload](docs/RemoveMemberPayload.md)
 - [RequestConsole](docs/RequestConsole.md)
 - [ResourcePayload](docs/ResourcePayload.md)
 - [ResponseModel](docs/ResponseModel.md)
 - [RolePermissionFields](docs/RolePermissionFields.md)
 - [RolePolicyFields](docs/RolePolicyFields.md)
 - [SecurityGroupRule](docs/SecurityGroupRule.md)
 - [SecurityGroupRuleFields](docs/SecurityGroupRuleFields.md)
 - [SecurityRulesFieldsforInstance](docs/SecurityRulesFieldsforInstance.md)
 - [SecurityRulesProtocolFields](docs/SecurityRulesProtocolFields.md)
 - [SetDefaultsPayload](docs/SetDefaultsPayload.md)
 - [SingleVisibilityUserResponse](docs/SingleVisibilityUserResponse.md)
 - [StartDeployment](docs/StartDeployment.md)
 - [StartDeploymentPayload](docs/StartDeploymentPayload.md)
 - [StockVisibilityUserListResponse](docs/StockVisibilityUserListResponse.md)
 - [StockVisibilityUserPayload](docs/StockVisibilityUserPayload.md)
 - [SuccessResponseModel](docs/SuccessResponseModel.md)
 - [Template](docs/Template.md)
 - [TemplateFields](docs/TemplateFields.md)
 - [Templates](docs/Templates.md)
 - [TokenFields](docs/TokenFields.md)
 - [UpdateContractPayload](docs/UpdateContractPayload.md)
 - [UpdateDiscountsPayload](docs/UpdateDiscountsPayload.md)
 - [UpdateDiscountsStatusPayload](docs/UpdateDiscountsStatusPayload.md)
 - [UpdateEnvironment](docs/UpdateEnvironment.md)
 - [UpdateGPU](docs/UpdateGPU.md)
 - [UpdateKeypairName](docs/UpdateKeypairName.md)
 - [UpdateKeypairnameresponse](docs/UpdateKeypairnameresponse.md)
 - [UpdateOrganizationPayload](docs/UpdateOrganizationPayload.md)
 - [UpdateOrganizationResponseModel](docs/UpdateOrganizationResponseModel.md)
 - [UpdateTemplate](docs/UpdateTemplate.md)
 - [UserDefaultChoiceForAdminFields](docs/UserDefaultChoiceForAdminFields.md)
 - [UserDefaultChoiceForUserFields](docs/UserDefaultChoiceForUserFields.md)
 - [UserDefaultChoicesForAdminResponse](docs/UserDefaultChoicesForAdminResponse.md)
 - [UserDefaultChoicesForUserResponse](docs/UserDefaultChoicesForUserResponse.md)
 - [UserPermissionFields](docs/UserPermissionFields.md)
 - [UserTransferPayload](docs/UserTransferPayload.md)
 - [Userinfopostpayload](docs/Userinfopostpayload.md)
 - [UsersInfoFields](docs/UsersInfoFields.md)
 - [UsersInfoListResponse](docs/UsersInfoListResponse.md)
 - [VNCURL](docs/VNCURL.md)
 - [VNCURLFields](docs/VNCURLFields.md)
 - [VerifyApiKeyPayload](docs/VerifyApiKeyPayload.md)
 - [VerifyApiKeyResponseModel](docs/VerifyApiKeyResponseModel.md)
 - [Volume](docs/Volume.md)
 - [VolumeAttachmentFields](docs/VolumeAttachmentFields.md)
 - [VolumeFields](docs/VolumeFields.md)
 - [VolumeFieldsforInstance](docs/VolumeFieldsforInstance.md)
 - [VolumeOverviewFields](docs/VolumeOverviewFields.md)
 - [VolumeTypes](docs/VolumeTypes.md)
 - [Volumes](docs/Volumes.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### accessToken

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		hyperstack.ContextAPIKeys,
		map[string]hyperstack.APIKey{
			"Authorization": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### apiKey

- **Type**: API key
- **API key parameter name**: api_key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: api_key and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		hyperstack.ContextAPIKeys,
		map[string]hyperstack.APIKey{
			"api_key": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



