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

// AccessPackageRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageCatalogs\{accessPackageCatalog-id}\accessPackages\{accessPackage-id}
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
// AccessPackageRequestBuilderGetQueryParameters the access packages in this catalog. Read-only. Nullable.
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
func (m *AccessPackageRequestBuilder) AccessPackageAssignmentPolicies()(*iba95de12c1af7af27f26180e0467cc5ec03c9bc57d1bffc7b60b209a3879a30e.AccessPackageAssignmentPoliciesRequestBuilder) {
    return iba95de12c1af7af27f26180e0467cc5ec03c9bc57d1bffc7b60b209a3879a30e.NewAccessPackageAssignmentPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageCatalogs.item.accessPackages.item.accessPackageAssignmentPolicies.item collection
func (m *AccessPackageRequestBuilder) AccessPackageAssignmentPoliciesById(id string)(*i39c5e89e204906c062687a386a0d848e6611ef1843b6369276fa3ba4665c98c9.AccessPackageAssignmentPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentPolicy_id"] = id
    }
    return i39c5e89e204906c062687a386a0d848e6611ef1843b6369276fa3ba4665c98c9.NewAccessPackageAssignmentPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) AccessPackageCatalog()(*i099ff8300577d054f828e3660ff7c64555088ad12039a74d607fb15d91b11b22.AccessPackageCatalogRequestBuilder) {
    return i099ff8300577d054f828e3660ff7c64555088ad12039a74d607fb15d91b11b22.NewAccessPackageCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) AccessPackageResourceRoleScopes()(*i7d7f233f0ad689059ae551b42cbdcf651191bbaffe672acf45069955346d609d.AccessPackageResourceRoleScopesRequestBuilder) {
    return i7d7f233f0ad689059ae551b42cbdcf651191bbaffe672acf45069955346d609d.NewAccessPackageResourceRoleScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRoleScopesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageCatalogs.item.accessPackages.item.accessPackageResourceRoleScopes.item collection
func (m *AccessPackageRequestBuilder) AccessPackageResourceRoleScopesById(id string)(*i679b728c2ea18c020e946e877c8f683d2811733265eeb5287486c65ca1c521bd.AccessPackageResourceRoleScopeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRoleScope_id"] = id
    }
    return i679b728c2ea18c020e946e877c8f683d2811733265eeb5287486c65ca1c521bd.NewAccessPackageResourceRoleScopeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) AccessPackagesIncompatibleWith()(*i68d71cd4e8be98e26bfe2cf5e9d6195d7b890492d0376612b2bb2386875d6980.AccessPackagesIncompatibleWithRequestBuilder) {
    return i68d71cd4e8be98e26bfe2cf5e9d6195d7b890492d0376612b2bb2386875d6980.NewAccessPackagesIncompatibleWithRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessPackageRequestBuilderInternal instantiates a new AccessPackageRequestBuilder and sets the default values.
func NewAccessPackageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageRequestBuilder) {
    m := &AccessPackageRequestBuilder{
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
// NewAccessPackageRequestBuilder instantiates a new AccessPackageRequestBuilder and sets the default values.
func NewAccessPackageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the access packages in this catalog. Read-only. Nullable.
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
// CreateGetRequestInformation the access packages in this catalog. Read-only. Nullable.
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
// CreatePatchRequestInformation the access packages in this catalog. Read-only. Nullable.
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
// Delete the access packages in this catalog. Read-only. Nullable.
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
// Get the access packages in this catalog. Read-only. Nullable.
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
func (m *AccessPackageRequestBuilder) GetApplicablePolicyRequirements()(*ic758e20d3da2bb69484030a405ef6779480d7d1f8198a6e5781c2b82f1270aeb.GetApplicablePolicyRequirementsRequestBuilder) {
    return ic758e20d3da2bb69484030a405ef6779480d7d1f8198a6e5781c2b82f1270aeb.NewGetApplicablePolicyRequirementsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) IncompatibleAccessPackages()(*idf8be4049a06eba3c02cea82c1e467235689bbb72e7f5a9316debb4a2f08bdbe.IncompatibleAccessPackagesRequestBuilder) {
    return idf8be4049a06eba3c02cea82c1e467235689bbb72e7f5a9316debb4a2f08bdbe.NewIncompatibleAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) IncompatibleGroups()(*ie1a8c7d364379750d0ba3c645628705d279bdaaa8552b9a79d5ebdb5897d9af1.IncompatibleGroupsRequestBuilder) {
    return ie1a8c7d364379750d0ba3c645628705d279bdaaa8552b9a79d5ebdb5897d9af1.NewIncompatibleGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the access packages in this catalog. Read-only. Nullable.
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
