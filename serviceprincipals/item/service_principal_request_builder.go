package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i11c002c434a5b4d143e52f3474f0a22379e083e82210f3f169d2093759e97c80 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/approleassignedto"
    i1ba52395740d2fbf9625e50150109aa1e368b2454968c67738f4f2af34aee752 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/delegatedpermissionclassifications"
    i23a06a5bd1dd319cab7683e2769b8ec9b9f5fa2c7ab7c1a6d8384ddaa492014b "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/getmemberobjects"
    i2efb05d83e4aa6ac6147e393e78c353ad0ec9595cbe9160b88d67cdd1d929e87 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/tokenlifetimepolicies"
    i3085d473efc7695c2ec95dbf6722fdb29435eb21e254c6c906e79bee48934977 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/deletepasswordsinglesignoncredentials"
    i33164a67269761d26f8ae32725ec9f06d993d6173a0fd094213e52d827047d04 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization"
    i357c592f2e173cff04efd24060abc2cdb7c037446a1a5fef4a892d849f510985 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/createpasswordsinglesignoncredentials"
    i3ba0e028cf2fd2dd6f6e9af983b18664a1fd2f9f97df4e4f239075c5d82ee6d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/licensedetails"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
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
    i62a4342854a4da102701b12124cb86be14b4c17b8b8691f57c0ae992cd399674 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/endpoints/item"
    i73410d76321ebd4b4207e0ffdf54200d41fb342ffc03bc6d921827db529b903c "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/approleassignedto/item"
    ieaae3e55e6f40942560f470e7e5fb2380c1742e2c611f1397ebadf664105f6e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/approleassignments/item"
    ied131e9801fe0f7aa043e4bd950bcb9d42a80e13246c9b382d504487c787a1d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/federatedidentitycredentials/item"
    if711a515916176e7c09aa3e5a8f78d5306dd95e446b78ea2cdac1f4050779e27 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/delegatedpermissionclassifications/item"
)

// ServicePrincipalRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}
type ServicePrincipalRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ServicePrincipalRequestBuilderDeleteOptions options for Delete
type ServicePrincipalRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ServicePrincipalRequestBuilderGetOptions options for Get
type ServicePrincipalRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ServicePrincipalRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ServicePrincipalRequestBuilderGetQueryParameters get entity from servicePrincipals by key
type ServicePrincipalRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ServicePrincipalRequestBuilderPatchOptions options for Patch
type ServicePrincipalRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServicePrincipal;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ServicePrincipalRequestBuilder) AddTokenSigningCertificate()(*icd93e098c0af3fad76be5463dae742feef1548ad1abbe576f34223e3d452835f.AddTokenSigningCertificateRequestBuilder) {
    return icd93e098c0af3fad76be5463dae742feef1548ad1abbe576f34223e3d452835f.NewAddTokenSigningCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) AppManagementPolicies()(*i9b6204d878ba8d08acc22af10d41ba3d8a46ea87016990ab20b7edd5329f5e9d.AppManagementPoliciesRequestBuilder) {
    return i9b6204d878ba8d08acc22af10d41ba3d8a46ea87016990ab20b7edd5329f5e9d.NewAppManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) AppRoleAssignedTo()(*i11c002c434a5b4d143e52f3474f0a22379e083e82210f3f169d2093759e97c80.AppRoleAssignedToRequestBuilder) {
    return i11c002c434a5b4d143e52f3474f0a22379e083e82210f3f169d2093759e97c80.NewAppRoleAssignedToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignedToById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.servicePrincipals.item.appRoleAssignedTo.item collection
func (m *ServicePrincipalRequestBuilder) AppRoleAssignedToById(id string)(*i73410d76321ebd4b4207e0ffdf54200d41fb342ffc03bc6d921827db529b903c.AppRoleAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment_id"] = id
    }
    return i73410d76321ebd4b4207e0ffdf54200d41fb342ffc03bc6d921827db529b903c.NewAppRoleAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) AppRoleAssignments()(*i8c780b5242f53b8ba76b2cb0f6bf0809832c4c71c2d0c657612557cb50d1ccd2.AppRoleAssignmentsRequestBuilder) {
    return i8c780b5242f53b8ba76b2cb0f6bf0809832c4c71c2d0c657612557cb50d1ccd2.NewAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.servicePrincipals.item.appRoleAssignments.item collection
func (m *ServicePrincipalRequestBuilder) AppRoleAssignmentsById(id string)(*ieaae3e55e6f40942560f470e7e5fb2380c1742e2c611f1397ebadf664105f6e3.AppRoleAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment_id"] = id
    }
    return ieaae3e55e6f40942560f470e7e5fb2380c1742e2c611f1397ebadf664105f6e3.NewAppRoleAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) CheckMemberGroups()(*ie9dc214192a9e2f06b5942f160aa60cd9ca4ef7bf2b5def2df094f16358ee83f.CheckMemberGroupsRequestBuilder) {
    return ie9dc214192a9e2f06b5942f160aa60cd9ca4ef7bf2b5def2df094f16358ee83f.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) CheckMemberObjects()(*ie7557ac9d3f37fb3cbfd32bd75d54411989729a3308b72cefbc4a42b818f6595.CheckMemberObjectsRequestBuilder) {
    return ie7557ac9d3f37fb3cbfd32bd75d54411989729a3308b72cefbc4a42b818f6595.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) ClaimsMappingPolicies()(*i5885ced9dd7e8cbb4309f2263f31f11f6fe6fba0a377430a7a55945a60cfeeba.ClaimsMappingPoliciesRequestBuilder) {
    return i5885ced9dd7e8cbb4309f2263f31f11f6fe6fba0a377430a7a55945a60cfeeba.NewClaimsMappingPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewServicePrincipalRequestBuilderInternal instantiates a new ServicePrincipalRequestBuilder and sets the default values.
func NewServicePrincipalRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServicePrincipalRequestBuilder) {
    m := &ServicePrincipalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewServicePrincipalRequestBuilder instantiates a new ServicePrincipalRequestBuilder and sets the default values.
func NewServicePrincipalRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServicePrincipalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from servicePrincipals
func (m *ServicePrincipalRequestBuilder) CreateDeleteRequestInformation(options *ServicePrincipalRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ServicePrincipalRequestBuilder) CreatedObjects()(*i7c041c069068ca5c40e6ffa0b0c00c8740d0ce4b8cbb17b5f0faa13f0e1afaca.CreatedObjectsRequestBuilder) {
    return i7c041c069068ca5c40e6ffa0b0c00c8740d0ce4b8cbb17b5f0faa13f0e1afaca.NewCreatedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entity from servicePrincipals by key
func (m *ServicePrincipalRequestBuilder) CreateGetRequestInformation(options *ServicePrincipalRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ServicePrincipalRequestBuilder) CreatePasswordSingleSignOnCredentials()(*i357c592f2e173cff04efd24060abc2cdb7c037446a1a5fef4a892d849f510985.CreatePasswordSingleSignOnCredentialsRequestBuilder) {
    return i357c592f2e173cff04efd24060abc2cdb7c037446a1a5fef4a892d849f510985.NewCreatePasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update entity in servicePrincipals
func (m *ServicePrincipalRequestBuilder) CreatePatchRequestInformation(options *ServicePrincipalRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ServicePrincipalRequestBuilder) DelegatedPermissionClassifications()(*i1ba52395740d2fbf9625e50150109aa1e368b2454968c67738f4f2af34aee752.DelegatedPermissionClassificationsRequestBuilder) {
    return i1ba52395740d2fbf9625e50150109aa1e368b2454968c67738f4f2af34aee752.NewDelegatedPermissionClassificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DelegatedPermissionClassificationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.servicePrincipals.item.delegatedPermissionClassifications.item collection
func (m *ServicePrincipalRequestBuilder) DelegatedPermissionClassificationsById(id string)(*if711a515916176e7c09aa3e5a8f78d5306dd95e446b78ea2cdac1f4050779e27.DelegatedPermissionClassificationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedPermissionClassification_id"] = id
    }
    return if711a515916176e7c09aa3e5a8f78d5306dd95e446b78ea2cdac1f4050779e27.NewDelegatedPermissionClassificationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete entity from servicePrincipals
func (m *ServicePrincipalRequestBuilder) Delete(options *ServicePrincipalRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ServicePrincipalRequestBuilder) DeletePasswordSingleSignOnCredentials()(*i3085d473efc7695c2ec95dbf6722fdb29435eb21e254c6c906e79bee48934977.DeletePasswordSingleSignOnCredentialsRequestBuilder) {
    return i3085d473efc7695c2ec95dbf6722fdb29435eb21e254c6c906e79bee48934977.NewDeletePasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) Endpoints()(*ia14a94af6e363e7e0fd6365603ac8e9fbb6590d9ae6aca87c947d1f222aa1026.EndpointsRequestBuilder) {
    return ia14a94af6e363e7e0fd6365603ac8e9fbb6590d9ae6aca87c947d1f222aa1026.NewEndpointsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EndpointsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.servicePrincipals.item.endpoints.item collection
func (m *ServicePrincipalRequestBuilder) EndpointsById(id string)(*i62a4342854a4da102701b12124cb86be14b4c17b8b8691f57c0ae992cd399674.EndpointRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["endpoint_id"] = id
    }
    return i62a4342854a4da102701b12124cb86be14b4c17b8b8691f57c0ae992cd399674.NewEndpointRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) FederatedIdentityCredentials()(*i871861d7334cfd7aa035f2f94f20cb0ef970009cb980f790f81b88a88b4d4d14.FederatedIdentityCredentialsRequestBuilder) {
    return i871861d7334cfd7aa035f2f94f20cb0ef970009cb980f790f81b88a88b4d4d14.NewFederatedIdentityCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FederatedIdentityCredentialsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.servicePrincipals.item.federatedIdentityCredentials.item collection
func (m *ServicePrincipalRequestBuilder) FederatedIdentityCredentialsById(id string)(*ied131e9801fe0f7aa043e4bd950bcb9d42a80e13246c9b382d504487c787a1d5.FederatedIdentityCredentialRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["federatedIdentityCredential_id"] = id
    }
    return ied131e9801fe0f7aa043e4bd950bcb9d42a80e13246c9b382d504487c787a1d5.NewFederatedIdentityCredentialRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get entity from servicePrincipals by key
func (m *ServicePrincipalRequestBuilder) Get(options *ServicePrincipalRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServicePrincipal, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewServicePrincipal() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServicePrincipal), nil
}
func (m *ServicePrincipalRequestBuilder) GetMemberGroups()(*i9419bacd0b3c60b676099cb223cf8bc748a72216ebbe0f24a9bc342a42ac8bf0.GetMemberGroupsRequestBuilder) {
    return i9419bacd0b3c60b676099cb223cf8bc748a72216ebbe0f24a9bc342a42ac8bf0.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) GetMemberObjects()(*i23a06a5bd1dd319cab7683e2769b8ec9b9f5fa2c7ab7c1a6d8384ddaa492014b.GetMemberObjectsRequestBuilder) {
    return i23a06a5bd1dd319cab7683e2769b8ec9b9f5fa2c7ab7c1a6d8384ddaa492014b.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) GetPasswordSingleSignOnCredentials()(*ib1f68813a15dd20f8ba94567ec3492119dd7611b59340dd17a087d77a5cd9d6d.GetPasswordSingleSignOnCredentialsRequestBuilder) {
    return ib1f68813a15dd20f8ba94567ec3492119dd7611b59340dd17a087d77a5cd9d6d.NewGetPasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) HomeRealmDiscoveryPolicies()(*id30c7593789d23d6eac3f9084377b61498fb31de8edaa02c0bdec09809b55f35.HomeRealmDiscoveryPoliciesRequestBuilder) {
    return id30c7593789d23d6eac3f9084377b61498fb31de8edaa02c0bdec09809b55f35.NewHomeRealmDiscoveryPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) LicenseDetails()(*i3ba0e028cf2fd2dd6f6e9af983b18664a1fd2f9f97df4e4f239075c5d82ee6d9.LicenseDetailsRequestBuilder) {
    return i3ba0e028cf2fd2dd6f6e9af983b18664a1fd2f9f97df4e4f239075c5d82ee6d9.NewLicenseDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LicenseDetailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.servicePrincipals.item.licenseDetails.item collection
func (m *ServicePrincipalRequestBuilder) LicenseDetailsById(id string)(*i3ba0e028cf2fd2dd6f6e9af983b18664a1fd2f9f97df4e4f239075c5d82ee6d9.LicenseDetailsRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["licenseDetails_id"] = id
    }
    return i3ba0e028cf2fd2dd6f6e9af983b18664a1fd2f9f97df4e4f239075c5d82ee6d9.NewLicenseDetailsRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) MemberOf()(*i954486fee06d06896fcf38cf1c14aca15ee3091055008b27fd0876ead49b1047.MemberOfRequestBuilder) {
    return i954486fee06d06896fcf38cf1c14aca15ee3091055008b27fd0876ead49b1047.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) Oauth2PermissionGrants()(*ib40fceeb10f0f8fde436409edd6fe6812ef0e461aa1d50457964448cc7001831.Oauth2PermissionGrantsRequestBuilder) {
    return ib40fceeb10f0f8fde436409edd6fe6812ef0e461aa1d50457964448cc7001831.NewOauth2PermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) OwnedObjects()(*i69b620b073c121202a57e32daf220f28ce14ca155a3567658e778b4d809c57f8.OwnedObjectsRequestBuilder) {
    return i69b620b073c121202a57e32daf220f28ce14ca155a3567658e778b4d809c57f8.NewOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) Owners()(*if90a5f8cb729921ddd07df364e6be8db5481a3a4e2f571a7e431c352691d423b.OwnersRequestBuilder) {
    return if90a5f8cb729921ddd07df364e6be8db5481a3a4e2f571a7e431c352691d423b.NewOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in servicePrincipals
func (m *ServicePrincipalRequestBuilder) Patch(options *ServicePrincipalRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ServicePrincipalRequestBuilder) Restore()(*ie33fea0a8ac69b4f1634f4717f1fb2e5201d325d8a31d6543a7fdf01ccb30f7f.RestoreRequestBuilder) {
    return ie33fea0a8ac69b4f1634f4717f1fb2e5201d325d8a31d6543a7fdf01ccb30f7f.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) Synchronization()(*i33164a67269761d26f8ae32725ec9f06d993d6173a0fd094213e52d827047d04.SynchronizationRequestBuilder) {
    return i33164a67269761d26f8ae32725ec9f06d993d6173a0fd094213e52d827047d04.NewSynchronizationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) TokenIssuancePolicies()(*i940efcd7d4b6ffe938ba761fd33e55282aa3fb1e6a4a8d0e9d8d45244f180cc4.TokenIssuancePoliciesRequestBuilder) {
    return i940efcd7d4b6ffe938ba761fd33e55282aa3fb1e6a4a8d0e9d8d45244f180cc4.NewTokenIssuancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) TokenLifetimePolicies()(*i2efb05d83e4aa6ac6147e393e78c353ad0ec9595cbe9160b88d67cdd1d929e87.TokenLifetimePoliciesRequestBuilder) {
    return i2efb05d83e4aa6ac6147e393e78c353ad0ec9595cbe9160b88d67cdd1d929e87.NewTokenLifetimePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) TransitiveMemberOf()(*iff23f4f1ad38ec43511f8fd0c9c34febf0b647f8c4d31b428584d77e243c45c4.TransitiveMemberOfRequestBuilder) {
    return iff23f4f1ad38ec43511f8fd0c9c34febf0b647f8c4d31b428584d77e243c45c4.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalRequestBuilder) UpdatePasswordSingleSignOnCredentials()(*ia7129e26ede98ad7dc73a6205e0b557b4c6b07141b5819da77f2c82e17f98268.UpdatePasswordSingleSignOnCredentialsRequestBuilder) {
    return ia7129e26ede98ad7dc73a6205e0b557b4c6b07141b5819da77f2c82e17f98268.NewUpdatePasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
