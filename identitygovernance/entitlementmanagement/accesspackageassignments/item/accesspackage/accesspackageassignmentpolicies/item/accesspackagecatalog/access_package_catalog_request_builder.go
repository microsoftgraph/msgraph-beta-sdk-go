package accesspackagecatalog

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i07a9ef5f9b0f40f8ebf893d8d44a46489643ca2292b15118310ce6f606686d13 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackageresources"
    i11205d712848fdbfa02ee78c0745d4d18b3b1500538a8551884f641f44badfa7 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackageresourcescopes"
    i768014ea8ea492d76aa1798ea1dbd5b7639ebff10140a5332b3f71fa08127d23 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackages"
    if786c788dd63819264409c3170147b1844bfa63d54547441b23fb95b7165d9dd "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackageresourceroles"
    i5288af25440b607fff6171b3b19b2d15ee0ed2206dea4eaccf26f44da1ebe69b "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackageresourceroles/item"
    i58f1a967e08d5762218747e4f8d392526cc0d30b5257629ddc8c86c10a823152 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackages/item"
    i5c5e19df77ffb5cf91bf3eae5b42b4a923ad5a683a41186f3f3deb6d1e09eaed "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackageresources/item"
    i71d45d936521056143746a13be3de8026a030d47aa7da36494920bccf0e32ec0 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackage/accesspackageassignmentpolicies/item/accesspackagecatalog/accesspackageresourcescopes/item"
)

// AccessPackageCatalogRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignments\{accessPackageAssignment-id}\accessPackage\accessPackageAssignmentPolicies\{accessPackageAssignmentPolicy-id}\accessPackageCatalog
type AccessPackageCatalogRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessPackageCatalogRequestBuilderDeleteOptions options for Delete
type AccessPackageCatalogRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageCatalogRequestBuilderGetOptions options for Get
type AccessPackageCatalogRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessPackageCatalogRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageCatalogRequestBuilderGetQueryParameters get accessPackageCatalog from identityGovernance
type AccessPackageCatalogRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// AccessPackageCatalogRequestBuilderPatchOptions options for Patch
type AccessPackageCatalogRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageCatalog;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResourceRoles()(*if786c788dd63819264409c3170147b1844bfa63d54547441b23fb95b7165d9dd.AccessPackageResourceRolesRequestBuilder) {
    return if786c788dd63819264409c3170147b1844bfa63d54547441b23fb95b7165d9dd.NewAccessPackageResourceRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRolesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignments.item.accessPackage.accessPackageAssignmentPolicies.item.accessPackageCatalog.accessPackageResourceRoles.item collection
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResourceRolesById(id string)(*i5288af25440b607fff6171b3b19b2d15ee0ed2206dea4eaccf26f44da1ebe69b.AccessPackageResourceRoleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRole_id"] = id
    }
    return i5288af25440b607fff6171b3b19b2d15ee0ed2206dea4eaccf26f44da1ebe69b.NewAccessPackageResourceRoleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResources()(*i07a9ef5f9b0f40f8ebf893d8d44a46489643ca2292b15118310ce6f606686d13.AccessPackageResourcesRequestBuilder) {
    return i07a9ef5f9b0f40f8ebf893d8d44a46489643ca2292b15118310ce6f606686d13.NewAccessPackageResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignments.item.accessPackage.accessPackageAssignmentPolicies.item.accessPackageCatalog.accessPackageResources.item collection
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResourcesById(id string)(*i5c5e19df77ffb5cf91bf3eae5b42b4a923ad5a683a41186f3f3deb6d1e09eaed.AccessPackageResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResource_id"] = id
    }
    return i5c5e19df77ffb5cf91bf3eae5b42b4a923ad5a683a41186f3f3deb6d1e09eaed.NewAccessPackageResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResourceScopes()(*i11205d712848fdbfa02ee78c0745d4d18b3b1500538a8551884f641f44badfa7.AccessPackageResourceScopesRequestBuilder) {
    return i11205d712848fdbfa02ee78c0745d4d18b3b1500538a8551884f641f44badfa7.NewAccessPackageResourceScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceScopesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignments.item.accessPackage.accessPackageAssignmentPolicies.item.accessPackageCatalog.accessPackageResourceScopes.item collection
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResourceScopesById(id string)(*i71d45d936521056143746a13be3de8026a030d47aa7da36494920bccf0e32ec0.AccessPackageResourceScopeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceScope_id"] = id
    }
    return i71d45d936521056143746a13be3de8026a030d47aa7da36494920bccf0e32ec0.NewAccessPackageResourceScopeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageCatalogRequestBuilder) AccessPackages()(*i768014ea8ea492d76aa1798ea1dbd5b7639ebff10140a5332b3f71fa08127d23.AccessPackagesRequestBuilder) {
    return i768014ea8ea492d76aa1798ea1dbd5b7639ebff10140a5332b3f71fa08127d23.NewAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignments.item.accessPackage.accessPackageAssignmentPolicies.item.accessPackageCatalog.accessPackages.item collection
func (m *AccessPackageCatalogRequestBuilder) AccessPackagesById(id string)(*i58f1a967e08d5762218747e4f8d392526cc0d30b5257629ddc8c86c10a823152.AccessPackageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage_id"] = id
    }
    return i58f1a967e08d5762218747e4f8d392526cc0d30b5257629ddc8c86c10a823152.NewAccessPackageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAccessPackageCatalogRequestBuilderInternal instantiates a new AccessPackageCatalogRequestBuilder and sets the default values.
func NewAccessPackageCatalogRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageCatalogRequestBuilder) {
    m := &AccessPackageCatalogRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment_id}/accessPackage/accessPackageAssignmentPolicies/{accessPackageAssignmentPolicy_id}/accessPackageCatalog{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackageCatalogRequestBuilder instantiates a new AccessPackageCatalogRequestBuilder and sets the default values.
func NewAccessPackageCatalogRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageCatalogRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageCatalogRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessPackageCatalog for identityGovernance
func (m *AccessPackageCatalogRequestBuilder) CreateDeleteRequestInformation(options *AccessPackageCatalogRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get accessPackageCatalog from identityGovernance
func (m *AccessPackageCatalogRequestBuilder) CreateGetRequestInformation(options *AccessPackageCatalogRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property accessPackageCatalog in identityGovernance
func (m *AccessPackageCatalogRequestBuilder) CreatePatchRequestInformation(options *AccessPackageCatalogRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property accessPackageCatalog for identityGovernance
func (m *AccessPackageCatalogRequestBuilder) Delete(options *AccessPackageCatalogRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get get accessPackageCatalog from identityGovernance
func (m *AccessPackageCatalogRequestBuilder) Get(options *AccessPackageCatalogRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageCatalog, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAccessPackageCatalog() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageCatalog), nil
}
// Patch update the navigation property accessPackageCatalog in identityGovernance
func (m *AccessPackageCatalogRequestBuilder) Patch(options *AccessPackageCatalogRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
