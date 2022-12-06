package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i11c002c434a5b4d143e52f3474f0a22379e083e82210f3f169d2093759e97c80 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/approleassignedto"
    i1ba52395740d2fbf9625e50150109aa1e368b2454968c67738f4f2af34aee752 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/delegatedpermissionclassifications"
    i23a06a5bd1dd319cab7683e2769b8ec9b9f5fa2c7ab7c1a6d8384ddaa492014b "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/getmemberobjects"
    i2efb05d83e4aa6ac6147e393e78c353ad0ec9595cbe9160b88d67cdd1d929e87 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/tokenlifetimepolicies"
    i3085d473efc7695c2ec95dbf6722fdb29435eb21e254c6c906e79bee48934977 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/deletepasswordsinglesignoncredentials"
    i33164a67269761d26f8ae32725ec9f06d993d6173a0fd094213e52d827047d04 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization"
    i357c592f2e173cff04efd24060abc2cdb7c037446a1a5fef4a892d849f510985 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/createpasswordsinglesignoncredentials"
    i3ba0e028cf2fd2dd6f6e9af983b18664a1fd2f9f97df4e4f239075c5d82ee6d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/licensedetails"
    i5885ced9dd7e8cbb4309f2263f31f11f6fe6fba0a377430a7a55945a60cfeeba "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/claimsmappingpolicies"
    i69b620b073c121202a57e32daf220f28ce14ca155a3567658e778b4d809c57f8 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/ownedobjects"
    i7c041c069068ca5c40e6ffa0b0c00c8740d0ce4b8cbb17b5f0faa13f0e1afaca "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/createdobjects"
    i871861d7334cfd7aa035f2f94f20cb0ef970009cb980f790f81b88a88b4d4d14 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/federatedidentitycredentials"
    i8c780b5242f53b8ba76b2cb0f6bf0809832c4c71c2d0c657612557cb50d1ccd2 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/approleassignments"
    i940efcd7d4b6ffe938ba761fd33e55282aa3fb1e6a4a8d0e9d8d45244f180cc4 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/tokenissuancepolicies"
    i9419bacd0b3c60b676099cb223cf8bc748a72216ebbe0f24a9bc342a42ac8bf0 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/getmembergroups"
    i954486fee06d06896fcf38cf1c14aca15ee3091055008b27fd0876ead49b1047 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/memberof"
    i9b6204d878ba8d08acc22af10d41ba3d8a46ea87016990ab20b7edd5329f5e9d "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/appmanagementpolicies"
    ia14a94af6e363e7e0fd6365603ac8e9fbb6590d9ae6aca87c947d1f222aa1026 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/endpoints"
    ia7129e26ede98ad7dc73a6205e0b557b4c6b07141b5819da77f2c82e17f98268 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/updatepasswordsinglesignoncredentials"
    ib1f68813a15dd20f8ba94567ec3492119dd7611b59340dd17a087d77a5cd9d6d "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/getpasswordsinglesignoncredentials"
    ib40fceeb10f0f8fde436409edd6fe6812ef0e461aa1d50457964448cc7001831 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/oauth2permissiongrants"
    icd93e098c0af3fad76be5463dae742feef1548ad1abbe576f34223e3d452835f "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/addtokensigningcertificate"
    id30c7593789d23d6eac3f9084377b61498fb31de8edaa02c0bdec09809b55f35 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/homerealmdiscoverypolicies"
    ie33fea0a8ac69b4f1634f4717f1fb2e5201d325d8a31d6543a7fdf01ccb30f7f "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/restore"
    ie7557ac9d3f37fb3cbfd32bd75d54411989729a3308b72cefbc4a42b818f6595 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/checkmemberobjects"
    ie9dc214192a9e2f06b5942f160aa60cd9ca4ef7bf2b5def2df094f16358ee83f "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/checkmembergroups"
    if90a5f8cb729921ddd07df364e6be8db5481a3a4e2f571a7e431c352691d423b "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/owners"
    iff23f4f1ad38ec43511f8fd0c9c34febf0b647f8c4d31b428584d77e243c45c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof"
    i063566d8d9a1b83cc0d5fbfad4c9e22b520c41edfe64764f7ff8d275874b02a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/ownedobjects/item"
    i234e2b441eb6121a8af32a7bac5d32752edca00c3103fbe6cf51997a556c95a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/owners/item"
    i30a88e502a1b36c0ef991fe606f7b7a5334452187e51f492fbbf02a9917ac38c "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/createdobjects/item"
    i5955bca5a205363dbadc089d22ca39765ce0b01bb5f74231714bd9fd47f3f79b "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/tokenlifetimepolicies/item"
    i62a4342854a4da102701b12124cb86be14b4c17b8b8691f57c0ae992cd399674 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/endpoints/item"
    i6a015459333573414a123fba74a4d57f7c21b624c5b2e738ca971c7cebb05410 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/tokenissuancepolicies/item"
    i6b393bae25cbe9b990d8caf535f9f81fc347fdc3062a3518851c4304337a7516 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/memberof/item"
    i73410d76321ebd4b4207e0ffdf54200d41fb342ffc03bc6d921827db529b903c "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/approleassignedto/item"
    i787211dc3896819f655b4fc9e0fa227fb0d5d943d81866ae7ab7ec5c16816da3 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/claimsmappingpolicies/item"
    i8dedea94c6f21c6173a3636499bc6c1b5f870c49de115840573a52af9135a448 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/homerealmdiscoverypolicies/item"
    ib1a73e2744468ed6683d9579b9d268ac4b1655e784fa7abb13d80fc66425cd91 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/appmanagementpolicies/item"
    ib2eed8adad035268bd5db1064969a9c074d122a3ffe5c95dc1b9bee05e0aec50 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/item"
    ieaae3e55e6f40942560f470e7e5fb2380c1742e2c611f1397ebadf664105f6e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/approleassignments/item"
    ied131e9801fe0f7aa043e4bd950bcb9d42a80e13246c9b382d504487c787a1d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/federatedidentitycredentials/item"
    if53b1c603e73bd2e2b1a2f087b7f62d88ef03564750eec23f62906a85212f34d "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/oauth2permissiongrants/item"
    if711a515916176e7c09aa3e5a8f78d5306dd95e446b78ea2cdac1f4050779e27 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/delegatedpermissionclassifications/item"
    if93ec8bb647ba14f89819a5269b0ffbab5940b4c3269233d3ced5ba508bd8f9b "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/licensedetails/item"
)

// ServicePrincipalItemRequestBuilder provides operations to manage the collection of servicePrincipal entities.
type ServicePrincipalItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ServicePrincipalItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServicePrincipalItemRequestBuilderGetQueryParameters retrieve the properties and relationships of a servicePrincipal object.
type ServicePrincipalItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServicePrincipalItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServicePrincipalItemRequestBuilderGetQueryParameters
}
// ServicePrincipalItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddTokenSigningCertificate provides operations to call the addTokenSigningCertificate method.
func (m *ServicePrincipalItemRequestBuilder) AddTokenSigningCertificate()(*icd93e098c0af3fad76be5463dae742feef1548ad1abbe576f34223e3d452835f.AddTokenSigningCertificateRequestBuilder) {
    return icd93e098c0af3fad76be5463dae742feef1548ad1abbe576f34223e3d452835f.NewAddTokenSigningCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppManagementPolicies provides operations to manage the appManagementPolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) AppManagementPolicies()(*i9b6204d878ba8d08acc22af10d41ba3d8a46ea87016990ab20b7edd5329f5e9d.AppManagementPoliciesRequestBuilder) {
    return i9b6204d878ba8d08acc22af10d41ba3d8a46ea87016990ab20b7edd5329f5e9d.NewAppManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppManagementPoliciesById provides operations to manage the appManagementPolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) AppManagementPoliciesById(id string)(*ib1a73e2744468ed6683d9579b9d268ac4b1655e784fa7abb13d80fc66425cd91.AppManagementPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appManagementPolicy%2Did"] = id
    }
    return ib1a73e2744468ed6683d9579b9d268ac4b1655e784fa7abb13d80fc66425cd91.NewAppManagementPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AppRoleAssignedTo provides operations to manage the appRoleAssignedTo property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) AppRoleAssignedTo()(*i11c002c434a5b4d143e52f3474f0a22379e083e82210f3f169d2093759e97c80.AppRoleAssignedToRequestBuilder) {
    return i11c002c434a5b4d143e52f3474f0a22379e083e82210f3f169d2093759e97c80.NewAppRoleAssignedToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignedToById provides operations to manage the appRoleAssignedTo property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) AppRoleAssignedToById(id string)(*i73410d76321ebd4b4207e0ffdf54200d41fb342ffc03bc6d921827db529b903c.AppRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment%2Did"] = id
    }
    return i73410d76321ebd4b4207e0ffdf54200d41fb342ffc03bc6d921827db529b903c.NewAppRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AppRoleAssignments provides operations to manage the appRoleAssignments property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) AppRoleAssignments()(*i8c780b5242f53b8ba76b2cb0f6bf0809832c4c71c2d0c657612557cb50d1ccd2.AppRoleAssignmentsRequestBuilder) {
    return i8c780b5242f53b8ba76b2cb0f6bf0809832c4c71c2d0c657612557cb50d1ccd2.NewAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignmentsById provides operations to manage the appRoleAssignments property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) AppRoleAssignmentsById(id string)(*ieaae3e55e6f40942560f470e7e5fb2380c1742e2c611f1397ebadf664105f6e3.AppRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment%2Did"] = id
    }
    return ieaae3e55e6f40942560f470e7e5fb2380c1742e2c611f1397ebadf664105f6e3.NewAppRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *ServicePrincipalItemRequestBuilder) CheckMemberGroups()(*ie9dc214192a9e2f06b5942f160aa60cd9ca4ef7bf2b5def2df094f16358ee83f.CheckMemberGroupsRequestBuilder) {
    return ie9dc214192a9e2f06b5942f160aa60cd9ca4ef7bf2b5def2df094f16358ee83f.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *ServicePrincipalItemRequestBuilder) CheckMemberObjects()(*ie7557ac9d3f37fb3cbfd32bd75d54411989729a3308b72cefbc4a42b818f6595.CheckMemberObjectsRequestBuilder) {
    return ie7557ac9d3f37fb3cbfd32bd75d54411989729a3308b72cefbc4a42b818f6595.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClaimsMappingPolicies provides operations to manage the claimsMappingPolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) ClaimsMappingPolicies()(*i5885ced9dd7e8cbb4309f2263f31f11f6fe6fba0a377430a7a55945a60cfeeba.ClaimsMappingPoliciesRequestBuilder) {
    return i5885ced9dd7e8cbb4309f2263f31f11f6fe6fba0a377430a7a55945a60cfeeba.NewClaimsMappingPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClaimsMappingPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.servicePrincipals.item.claimsMappingPolicies.item collection
func (m *ServicePrincipalItemRequestBuilder) ClaimsMappingPoliciesById(id string)(*i787211dc3896819f655b4fc9e0fa227fb0d5d943d81866ae7ab7ec5c16816da3.ClaimsMappingPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["claimsMappingPolicy%2Did"] = id
    }
    return i787211dc3896819f655b4fc9e0fa227fb0d5d943d81866ae7ab7ec5c16816da3.NewClaimsMappingPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewServicePrincipalItemRequestBuilderInternal instantiates a new ServicePrincipalItemRequestBuilder and sets the default values.
func NewServicePrincipalItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalItemRequestBuilder) {
    m := &ServicePrincipalItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewServicePrincipalItemRequestBuilder instantiates a new ServicePrincipalItemRequestBuilder and sets the default values.
func NewServicePrincipalItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete a servicePrincipal object.
func (m *ServicePrincipalItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatedObjects provides operations to manage the createdObjects property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) CreatedObjects()(*i7c041c069068ca5c40e6ffa0b0c00c8740d0ce4b8cbb17b5f0faa13f0e1afaca.CreatedObjectsRequestBuilder) {
    return i7c041c069068ca5c40e6ffa0b0c00c8740d0ce4b8cbb17b5f0faa13f0e1afaca.NewCreatedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatedObjectsById provides operations to manage the createdObjects property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) CreatedObjectsById(id string)(*i30a88e502a1b36c0ef991fe606f7b7a5334452187e51f492fbbf02a9917ac38c.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i30a88e502a1b36c0ef991fe606f7b7a5334452187e51f492fbbf02a9917ac38c.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateGetRequestInformation retrieve the properties and relationships of a servicePrincipal object.
func (m *ServicePrincipalItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePasswordSingleSignOnCredentials provides operations to call the createPasswordSingleSignOnCredentials method.
func (m *ServicePrincipalItemRequestBuilder) CreatePasswordSingleSignOnCredentials()(*i357c592f2e173cff04efd24060abc2cdb7c037446a1a5fef4a892d849f510985.CreatePasswordSingleSignOnCredentialsRequestBuilder) {
    return i357c592f2e173cff04efd24060abc2cdb7c037446a1a5fef4a892d849f510985.NewCreatePasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the properties of servicePrincipal object.
func (m *ServicePrincipalItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable, requestConfiguration *ServicePrincipalItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DelegatedPermissionClassifications provides operations to manage the delegatedPermissionClassifications property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) DelegatedPermissionClassifications()(*i1ba52395740d2fbf9625e50150109aa1e368b2454968c67738f4f2af34aee752.DelegatedPermissionClassificationsRequestBuilder) {
    return i1ba52395740d2fbf9625e50150109aa1e368b2454968c67738f4f2af34aee752.NewDelegatedPermissionClassificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DelegatedPermissionClassificationsById provides operations to manage the delegatedPermissionClassifications property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) DelegatedPermissionClassificationsById(id string)(*if711a515916176e7c09aa3e5a8f78d5306dd95e446b78ea2cdac1f4050779e27.DelegatedPermissionClassificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedPermissionClassification%2Did"] = id
    }
    return if711a515916176e7c09aa3e5a8f78d5306dd95e446b78ea2cdac1f4050779e27.NewDelegatedPermissionClassificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete a servicePrincipal object.
func (m *ServicePrincipalItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ServicePrincipalItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeletePasswordSingleSignOnCredentials provides operations to call the deletePasswordSingleSignOnCredentials method.
func (m *ServicePrincipalItemRequestBuilder) DeletePasswordSingleSignOnCredentials()(*i3085d473efc7695c2ec95dbf6722fdb29435eb21e254c6c906e79bee48934977.DeletePasswordSingleSignOnCredentialsRequestBuilder) {
    return i3085d473efc7695c2ec95dbf6722fdb29435eb21e254c6c906e79bee48934977.NewDeletePasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Endpoints provides operations to manage the endpoints property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) Endpoints()(*ia14a94af6e363e7e0fd6365603ac8e9fbb6590d9ae6aca87c947d1f222aa1026.EndpointsRequestBuilder) {
    return ia14a94af6e363e7e0fd6365603ac8e9fbb6590d9ae6aca87c947d1f222aa1026.NewEndpointsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EndpointsById provides operations to manage the endpoints property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) EndpointsById(id string)(*i62a4342854a4da102701b12124cb86be14b4c17b8b8691f57c0ae992cd399674.EndpointItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["endpoint%2Did"] = id
    }
    return i62a4342854a4da102701b12124cb86be14b4c17b8b8691f57c0ae992cd399674.NewEndpointItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FederatedIdentityCredentials provides operations to manage the federatedIdentityCredentials property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) FederatedIdentityCredentials()(*i871861d7334cfd7aa035f2f94f20cb0ef970009cb980f790f81b88a88b4d4d14.FederatedIdentityCredentialsRequestBuilder) {
    return i871861d7334cfd7aa035f2f94f20cb0ef970009cb980f790f81b88a88b4d4d14.NewFederatedIdentityCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FederatedIdentityCredentialsById provides operations to manage the federatedIdentityCredentials property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) FederatedIdentityCredentialsById(id string)(*ied131e9801fe0f7aa043e4bd950bcb9d42a80e13246c9b382d504487c787a1d5.FederatedIdentityCredentialItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["federatedIdentityCredential%2Did"] = id
    }
    return ied131e9801fe0f7aa043e4bd950bcb9d42a80e13246c9b382d504487c787a1d5.NewFederatedIdentityCredentialItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get retrieve the properties and relationships of a servicePrincipal object.
func (m *ServicePrincipalItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ServicePrincipalItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServicePrincipalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable), nil
}
// GetMemberGroups provides operations to call the getMemberGroups method.
func (m *ServicePrincipalItemRequestBuilder) GetMemberGroups()(*i9419bacd0b3c60b676099cb223cf8bc748a72216ebbe0f24a9bc342a42ac8bf0.GetMemberGroupsRequestBuilder) {
    return i9419bacd0b3c60b676099cb223cf8bc748a72216ebbe0f24a9bc342a42ac8bf0.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *ServicePrincipalItemRequestBuilder) GetMemberObjects()(*i23a06a5bd1dd319cab7683e2769b8ec9b9f5fa2c7ab7c1a6d8384ddaa492014b.GetMemberObjectsRequestBuilder) {
    return i23a06a5bd1dd319cab7683e2769b8ec9b9f5fa2c7ab7c1a6d8384ddaa492014b.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetPasswordSingleSignOnCredentials provides operations to call the getPasswordSingleSignOnCredentials method.
func (m *ServicePrincipalItemRequestBuilder) GetPasswordSingleSignOnCredentials()(*ib1f68813a15dd20f8ba94567ec3492119dd7611b59340dd17a087d77a5cd9d6d.GetPasswordSingleSignOnCredentialsRequestBuilder) {
    return ib1f68813a15dd20f8ba94567ec3492119dd7611b59340dd17a087d77a5cd9d6d.NewGetPasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HomeRealmDiscoveryPolicies provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) HomeRealmDiscoveryPolicies()(*id30c7593789d23d6eac3f9084377b61498fb31de8edaa02c0bdec09809b55f35.HomeRealmDiscoveryPoliciesRequestBuilder) {
    return id30c7593789d23d6eac3f9084377b61498fb31de8edaa02c0bdec09809b55f35.NewHomeRealmDiscoveryPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HomeRealmDiscoveryPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.servicePrincipals.item.homeRealmDiscoveryPolicies.item collection
func (m *ServicePrincipalItemRequestBuilder) HomeRealmDiscoveryPoliciesById(id string)(*i8dedea94c6f21c6173a3636499bc6c1b5f870c49de115840573a52af9135a448.HomeRealmDiscoveryPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["homeRealmDiscoveryPolicy%2Did"] = id
    }
    return i8dedea94c6f21c6173a3636499bc6c1b5f870c49de115840573a52af9135a448.NewHomeRealmDiscoveryPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// LicenseDetails provides operations to manage the licenseDetails property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) LicenseDetails()(*i3ba0e028cf2fd2dd6f6e9af983b18664a1fd2f9f97df4e4f239075c5d82ee6d9.LicenseDetailsRequestBuilder) {
    return i3ba0e028cf2fd2dd6f6e9af983b18664a1fd2f9f97df4e4f239075c5d82ee6d9.NewLicenseDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LicenseDetailsById provides operations to manage the licenseDetails property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) LicenseDetailsById(id string)(*if93ec8bb647ba14f89819a5269b0ffbab5940b4c3269233d3ced5ba508bd8f9b.LicenseDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["licenseDetails%2Did"] = id
    }
    return if93ec8bb647ba14f89819a5269b0ffbab5940b4c3269233d3ced5ba508bd8f9b.NewLicenseDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MemberOf provides operations to manage the memberOf property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) MemberOf()(*i954486fee06d06896fcf38cf1c14aca15ee3091055008b27fd0876ead49b1047.MemberOfRequestBuilder) {
    return i954486fee06d06896fcf38cf1c14aca15ee3091055008b27fd0876ead49b1047.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById provides operations to manage the memberOf property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) MemberOfById(id string)(*i6b393bae25cbe9b990d8caf535f9f81fc347fdc3062a3518851c4304337a7516.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i6b393bae25cbe9b990d8caf535f9f81fc347fdc3062a3518851c4304337a7516.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Oauth2PermissionGrants provides operations to manage the oauth2PermissionGrants property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) Oauth2PermissionGrants()(*ib40fceeb10f0f8fde436409edd6fe6812ef0e461aa1d50457964448cc7001831.Oauth2PermissionGrantsRequestBuilder) {
    return ib40fceeb10f0f8fde436409edd6fe6812ef0e461aa1d50457964448cc7001831.NewOauth2PermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Oauth2PermissionGrantsById provides operations to manage the oauth2PermissionGrants property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) Oauth2PermissionGrantsById(id string)(*if53b1c603e73bd2e2b1a2f087b7f62d88ef03564750eec23f62906a85212f34d.OAuth2PermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["oAuth2PermissionGrant%2Did"] = id
    }
    return if53b1c603e73bd2e2b1a2f087b7f62d88ef03564750eec23f62906a85212f34d.NewOAuth2PermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OwnedObjects provides operations to manage the ownedObjects property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) OwnedObjects()(*i69b620b073c121202a57e32daf220f28ce14ca155a3567658e778b4d809c57f8.OwnedObjectsRequestBuilder) {
    return i69b620b073c121202a57e32daf220f28ce14ca155a3567658e778b4d809c57f8.NewOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnedObjectsById provides operations to manage the ownedObjects property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) OwnedObjectsById(id string)(*i063566d8d9a1b83cc0d5fbfad4c9e22b520c41edfe64764f7ff8d275874b02a2.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i063566d8d9a1b83cc0d5fbfad4c9e22b520c41edfe64764f7ff8d275874b02a2.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Owners provides operations to manage the owners property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) Owners()(*if90a5f8cb729921ddd07df364e6be8db5481a3a4e2f571a7e431c352691d423b.OwnersRequestBuilder) {
    return if90a5f8cb729921ddd07df364e6be8db5481a3a4e2f571a7e431c352691d423b.NewOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.servicePrincipals.item.owners.item collection
func (m *ServicePrincipalItemRequestBuilder) OwnersById(id string)(*i234e2b441eb6121a8af32a7bac5d32752edca00c3103fbe6cf51997a556c95a7.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i234e2b441eb6121a8af32a7bac5d32752edca00c3103fbe6cf51997a556c95a7.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the properties of servicePrincipal object.
func (m *ServicePrincipalItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable, requestConfiguration *ServicePrincipalItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServicePrincipalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable), nil
}
// Restore provides operations to call the restore method.
func (m *ServicePrincipalItemRequestBuilder) Restore()(*ie33fea0a8ac69b4f1634f4717f1fb2e5201d325d8a31d6543a7fdf01ccb30f7f.RestoreRequestBuilder) {
    return ie33fea0a8ac69b4f1634f4717f1fb2e5201d325d8a31d6543a7fdf01ccb30f7f.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Synchronization provides operations to manage the synchronization property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) Synchronization()(*i33164a67269761d26f8ae32725ec9f06d993d6173a0fd094213e52d827047d04.SynchronizationRequestBuilder) {
    return i33164a67269761d26f8ae32725ec9f06d993d6173a0fd094213e52d827047d04.NewSynchronizationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenIssuancePolicies provides operations to manage the tokenIssuancePolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) TokenIssuancePolicies()(*i940efcd7d4b6ffe938ba761fd33e55282aa3fb1e6a4a8d0e9d8d45244f180cc4.TokenIssuancePoliciesRequestBuilder) {
    return i940efcd7d4b6ffe938ba761fd33e55282aa3fb1e6a4a8d0e9d8d45244f180cc4.NewTokenIssuancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenIssuancePoliciesById provides operations to manage the tokenIssuancePolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) TokenIssuancePoliciesById(id string)(*i6a015459333573414a123fba74a4d57f7c21b624c5b2e738ca971c7cebb05410.TokenIssuancePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tokenIssuancePolicy%2Did"] = id
    }
    return i6a015459333573414a123fba74a4d57f7c21b624c5b2e738ca971c7cebb05410.NewTokenIssuancePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TokenLifetimePolicies provides operations to manage the tokenLifetimePolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) TokenLifetimePolicies()(*i2efb05d83e4aa6ac6147e393e78c353ad0ec9595cbe9160b88d67cdd1d929e87.TokenLifetimePoliciesRequestBuilder) {
    return i2efb05d83e4aa6ac6147e393e78c353ad0ec9595cbe9160b88d67cdd1d929e87.NewTokenLifetimePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenLifetimePoliciesById provides operations to manage the tokenLifetimePolicies property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) TokenLifetimePoliciesById(id string)(*i5955bca5a205363dbadc089d22ca39765ce0b01bb5f74231714bd9fd47f3f79b.TokenLifetimePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tokenLifetimePolicy%2Did"] = id
    }
    return i5955bca5a205363dbadc089d22ca39765ce0b01bb5f74231714bd9fd47f3f79b.NewTokenLifetimePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TransitiveMemberOf provides operations to manage the transitiveMemberOf property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) TransitiveMemberOf()(*iff23f4f1ad38ec43511f8fd0c9c34febf0b647f8c4d31b428584d77e243c45c4.TransitiveMemberOfRequestBuilder) {
    return iff23f4f1ad38ec43511f8fd0c9c34febf0b647f8c4d31b428584d77e243c45c4.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById provides operations to manage the transitiveMemberOf property of the microsoft.graph.servicePrincipal entity.
func (m *ServicePrincipalItemRequestBuilder) TransitiveMemberOfById(id string)(*ib2eed8adad035268bd5db1064969a9c074d122a3ffe5c95dc1b9bee05e0aec50.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return ib2eed8adad035268bd5db1064969a9c074d122a3ffe5c95dc1b9bee05e0aec50.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UpdatePasswordSingleSignOnCredentials provides operations to call the updatePasswordSingleSignOnCredentials method.
func (m *ServicePrincipalItemRequestBuilder) UpdatePasswordSingleSignOnCredentials()(*ia7129e26ede98ad7dc73a6205e0b557b4c6b07141b5819da77f2c82e17f98268.UpdatePasswordSingleSignOnCredentialsRequestBuilder) {
    return ia7129e26ede98ad7dc73a6205e0b557b4c6b07141b5819da77f2c82e17f98268.NewUpdatePasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
