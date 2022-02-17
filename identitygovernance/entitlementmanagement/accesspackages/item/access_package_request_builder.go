package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2a44e775fb7f737b94d47f06032cc0ce0b4bfd2fbfa2463f91ebde6ab7f3db40 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/incompatibleaccesspackages"
    i7554110a94219c17393947ff9abe170d93bdfb30d153839818ba04ffcc76de4f "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackageresourcerolescopes"
    i76ccc4db9daa784ae822d619ad983c2223a197dcc8d1fde6c1f4e9a639bf2bf2 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackageassignmentpolicies"
    i9f6aa13376612aa214454151aaed75311be92072af5b00d7314fae36adcdfa19 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/incompatiblegroups"
    icf4fae5627b4c00615f25f4291b050076f3e43516464e10c1555e560fa24a900 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackagecatalog"
    idde732b29dba3e8a39bef8255c30333987cd56f4f2de35a68cc15f74a585f58f "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackagesincompatiblewith"
    ifd34e90f17620729556c1246884066ebb12a7b35296a94b49e16baf9acf9aef9 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/getapplicablepolicyrequirements"
    i32cd1387ebc6076b2999c52277e7c8235a84ad82ab68b7e7f5c3787e7e50ad58 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackageassignmentpolicies/item"
    i6f598a4c0565fbb06735b2ed6cb90e45996b75462797a784a018e5156dc02b37 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackages/item/accesspackageresourcerolescopes/item"
)

// AccessPackageRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackages\{accessPackage-id}
type AccessPackageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessPackageRequestBuilderDeleteOptions options for Delete
type AccessPackageRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageRequestBuilderGetOptions options for Get
type AccessPackageRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessPackageRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageRequestBuilderGetQueryParameters represents access package objects.
type AccessPackageRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AccessPackageRequestBuilderPatchOptions options for Patch
type AccessPackageRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackage;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AccessPackageRequestBuilder) AccessPackageAssignmentPolicies()(*i76ccc4db9daa784ae822d619ad983c2223a197dcc8d1fde6c1f4e9a639bf2bf2.AccessPackageAssignmentPoliciesRequestBuilder) {
    return i76ccc4db9daa784ae822d619ad983c2223a197dcc8d1fde6c1f4e9a639bf2bf2.NewAccessPackageAssignmentPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackages.item.accessPackageAssignmentPolicies.item collection
func (m *AccessPackageRequestBuilder) AccessPackageAssignmentPoliciesById(id string)(*i32cd1387ebc6076b2999c52277e7c8235a84ad82ab68b7e7f5c3787e7e50ad58.AccessPackageAssignmentPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentPolicy_id"] = id
    }
    return i32cd1387ebc6076b2999c52277e7c8235a84ad82ab68b7e7f5c3787e7e50ad58.NewAccessPackageAssignmentPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) AccessPackageCatalog()(*icf4fae5627b4c00615f25f4291b050076f3e43516464e10c1555e560fa24a900.AccessPackageCatalogRequestBuilder) {
    return icf4fae5627b4c00615f25f4291b050076f3e43516464e10c1555e560fa24a900.NewAccessPackageCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) AccessPackageResourceRoleScopes()(*i7554110a94219c17393947ff9abe170d93bdfb30d153839818ba04ffcc76de4f.AccessPackageResourceRoleScopesRequestBuilder) {
    return i7554110a94219c17393947ff9abe170d93bdfb30d153839818ba04ffcc76de4f.NewAccessPackageResourceRoleScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRoleScopesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackages.item.accessPackageResourceRoleScopes.item collection
func (m *AccessPackageRequestBuilder) AccessPackageResourceRoleScopesById(id string)(*i6f598a4c0565fbb06735b2ed6cb90e45996b75462797a784a018e5156dc02b37.AccessPackageResourceRoleScopeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRoleScope_id"] = id
    }
    return i6f598a4c0565fbb06735b2ed6cb90e45996b75462797a784a018e5156dc02b37.NewAccessPackageResourceRoleScopeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) AccessPackagesIncompatibleWith()(*idde732b29dba3e8a39bef8255c30333987cd56f4f2de35a68cc15f74a585f58f.AccessPackagesIncompatibleWithRequestBuilder) {
    return idde732b29dba3e8a39bef8255c30333987cd56f4f2de35a68cc15f74a585f58f.NewAccessPackagesIncompatibleWithRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessPackageRequestBuilderInternal instantiates a new AccessPackageRequestBuilder and sets the default values.
func NewAccessPackageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageRequestBuilder) {
    m := &AccessPackageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackageRequestBuilder instantiates a new AccessPackageRequestBuilder and sets the default values.
func NewAccessPackageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation represents access package objects.
func (m *AccessPackageRequestBuilder) CreateDeleteRequestInformation(options *AccessPackageRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation represents access package objects.
func (m *AccessPackageRequestBuilder) CreateGetRequestInformation(options *AccessPackageRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation represents access package objects.
func (m *AccessPackageRequestBuilder) CreatePatchRequestInformation(options *AccessPackageRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete represents access package objects.
func (m *AccessPackageRequestBuilder) Delete(options *AccessPackageRequestBuilderDeleteOptions)(error) {
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
// Get represents access package objects.
func (m *AccessPackageRequestBuilder) Get(options *AccessPackageRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackage, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAccessPackage() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackage), nil
}
func (m *AccessPackageRequestBuilder) GetApplicablePolicyRequirements()(*ifd34e90f17620729556c1246884066ebb12a7b35296a94b49e16baf9acf9aef9.GetApplicablePolicyRequirementsRequestBuilder) {
    return ifd34e90f17620729556c1246884066ebb12a7b35296a94b49e16baf9acf9aef9.NewGetApplicablePolicyRequirementsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) IncompatibleAccessPackages()(*i2a44e775fb7f737b94d47f06032cc0ce0b4bfd2fbfa2463f91ebde6ab7f3db40.IncompatibleAccessPackagesRequestBuilder) {
    return i2a44e775fb7f737b94d47f06032cc0ce0b4bfd2fbfa2463f91ebde6ab7f3db40.NewIncompatibleAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) IncompatibleGroups()(*i9f6aa13376612aa214454151aaed75311be92072af5b00d7314fae36adcdfa19.IncompatibleGroupsRequestBuilder) {
    return i9f6aa13376612aa214454151aaed75311be92072af5b00d7314fae36adcdfa19.NewIncompatibleGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch represents access package objects.
func (m *AccessPackageRequestBuilder) Patch(options *AccessPackageRequestBuilderPatchOptions)(error) {
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
