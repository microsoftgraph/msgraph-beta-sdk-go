package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i099ff8300577d054f828e3660ff7c64555088ad12039a74d607fb15d91b11b22 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item/accesspackagecatalog"
    i68d71cd4e8be98e26bfe2cf5e9d6195d7b890492d0376612b2bb2386875d6980 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item/accesspackagesincompatiblewith"
    i7d7f233f0ad689059ae551b42cbdcf651191bbaffe672acf45069955346d609d "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item/accesspackageresourcerolescopes"
    iba95de12c1af7af27f26180e0467cc5ec03c9bc57d1bffc7b60b209a3879a30e "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item/accesspackageassignmentpolicies"
    ic758e20d3da2bb69484030a405ef6779480d7d1f8198a6e5781c2b82f1270aeb "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item/getapplicablepolicyrequirements"
    idf8be4049a06eba3c02cea82c1e467235689bbb72e7f5a9316debb4a2f08bdbe "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item/incompatibleaccesspackages"
    ie1a8c7d364379750d0ba3c645628705d279bdaaa8552b9a79d5ebdb5897d9af1 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item/incompatiblegroups"
    i39c5e89e204906c062687a386a0d848e6611ef1843b6369276fa3ba4665c98c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item/accesspackageassignmentpolicies/item"
    i679b728c2ea18c020e946e877c8f683d2811733265eeb5287486c65ca1c521bd "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item/accesspackageresourcerolescopes/item"
)

// AccessPackageItemRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageCatalogs\{accessPackageCatalog-id}\accessPackages\{accessPackage-id}
type AccessPackageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessPackageItemRequestBuilderDeleteOptions options for Delete
type AccessPackageItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageItemRequestBuilderGetOptions options for Get
type AccessPackageItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessPackageItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageItemRequestBuilderGetQueryParameters the access packages in this catalog. Read-only. Nullable.
type AccessPackageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AccessPackageItemRequestBuilderPatchOptions options for Patch
type AccessPackageItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackage;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AccessPackageItemRequestBuilder) AccessPackageAssignmentPolicies()(*iba95de12c1af7af27f26180e0467cc5ec03c9bc57d1bffc7b60b209a3879a30e.AccessPackageAssignmentPoliciesRequestBuilder) {
    return iba95de12c1af7af27f26180e0467cc5ec03c9bc57d1bffc7b60b209a3879a30e.NewAccessPackageAssignmentPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageCatalogs.item.accessPackages.item.accessPackageAssignmentPolicies.item collection
func (m *AccessPackageItemRequestBuilder) AccessPackageAssignmentPoliciesById(id string)(*i39c5e89e204906c062687a386a0d848e6611ef1843b6369276fa3ba4665c98c9.AccessPackageAssignmentPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentPolicy_id"] = id
    }
    return i39c5e89e204906c062687a386a0d848e6611ef1843b6369276fa3ba4665c98c9.NewAccessPackageAssignmentPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageItemRequestBuilder) AccessPackageCatalog()(*i099ff8300577d054f828e3660ff7c64555088ad12039a74d607fb15d91b11b22.AccessPackageCatalogRequestBuilder) {
    return i099ff8300577d054f828e3660ff7c64555088ad12039a74d607fb15d91b11b22.NewAccessPackageCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageItemRequestBuilder) AccessPackageResourceRoleScopes()(*i7d7f233f0ad689059ae551b42cbdcf651191bbaffe672acf45069955346d609d.AccessPackageResourceRoleScopesRequestBuilder) {
    return i7d7f233f0ad689059ae551b42cbdcf651191bbaffe672acf45069955346d609d.NewAccessPackageResourceRoleScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRoleScopesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageCatalogs.item.accessPackages.item.accessPackageResourceRoleScopes.item collection
func (m *AccessPackageItemRequestBuilder) AccessPackageResourceRoleScopesById(id string)(*i679b728c2ea18c020e946e877c8f683d2811733265eeb5287486c65ca1c521bd.AccessPackageResourceRoleScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRoleScope_id"] = id
    }
    return i679b728c2ea18c020e946e877c8f683d2811733265eeb5287486c65ca1c521bd.NewAccessPackageResourceRoleScopeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageItemRequestBuilder) AccessPackagesIncompatibleWith()(*i68d71cd4e8be98e26bfe2cf5e9d6195d7b890492d0376612b2bb2386875d6980.AccessPackagesIncompatibleWithRequestBuilder) {
    return i68d71cd4e8be98e26bfe2cf5e9d6195d7b890492d0376612b2bb2386875d6980.NewAccessPackagesIncompatibleWithRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessPackageItemRequestBuilderInternal instantiates a new AccessPackageItemRequestBuilder and sets the default values.
func NewAccessPackageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageItemRequestBuilder) {
    m := &AccessPackageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageCatalogs/{accessPackageCatalog_id}/accessPackages/{accessPackage_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackageItemRequestBuilder instantiates a new AccessPackageItemRequestBuilder and sets the default values.
func NewAccessPackageItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the access packages in this catalog. Read-only. Nullable.
func (m *AccessPackageItemRequestBuilder) CreateDeleteRequestInformation(options *AccessPackageItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the access packages in this catalog. Read-only. Nullable.
func (m *AccessPackageItemRequestBuilder) CreateGetRequestInformation(options *AccessPackageItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the access packages in this catalog. Read-only. Nullable.
func (m *AccessPackageItemRequestBuilder) CreatePatchRequestInformation(options *AccessPackageItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the access packages in this catalog. Read-only. Nullable.
func (m *AccessPackageItemRequestBuilder) Delete(options *AccessPackageItemRequestBuilderDeleteOptions)(error) {
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
// Get the access packages in this catalog. Read-only. Nullable.
func (m *AccessPackageItemRequestBuilder) Get(options *AccessPackageItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackage, error) {
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
func (m *AccessPackageItemRequestBuilder) GetApplicablePolicyRequirements()(*ic758e20d3da2bb69484030a405ef6779480d7d1f8198a6e5781c2b82f1270aeb.GetApplicablePolicyRequirementsRequestBuilder) {
    return ic758e20d3da2bb69484030a405ef6779480d7d1f8198a6e5781c2b82f1270aeb.NewGetApplicablePolicyRequirementsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageItemRequestBuilder) IncompatibleAccessPackages()(*idf8be4049a06eba3c02cea82c1e467235689bbb72e7f5a9316debb4a2f08bdbe.IncompatibleAccessPackagesRequestBuilder) {
    return idf8be4049a06eba3c02cea82c1e467235689bbb72e7f5a9316debb4a2f08bdbe.NewIncompatibleAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageItemRequestBuilder) IncompatibleGroups()(*ie1a8c7d364379750d0ba3c645628705d279bdaaa8552b9a79d5ebdb5897d9af1.IncompatibleGroupsRequestBuilder) {
    return ie1a8c7d364379750d0ba3c645628705d279bdaaa8552b9a79d5ebdb5897d9af1.NewIncompatibleGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the access packages in this catalog. Read-only. Nullable.
func (m *AccessPackageItemRequestBuilder) Patch(options *AccessPackageItemRequestBuilderPatchOptions)(error) {
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
