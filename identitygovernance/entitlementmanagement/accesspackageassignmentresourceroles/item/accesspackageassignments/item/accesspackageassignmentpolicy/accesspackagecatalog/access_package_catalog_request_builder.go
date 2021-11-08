package accesspackagecatalog

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i37eb72050fb28d46d11ae54b42db111aab72ecf5aa9fd4c470866ba104fce7b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackagecatalog/accesspackageresourceroles"
    i7e2cbe90c8879e155340a710c82fff884137ac87bba0e04dcdc5a342de234272 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages"
    i803a7b36907f71072fade484e0cf793a7743c53abed3fdcfd96cca7266386172 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackagecatalog/accesspackageresourcescopes"
    i947925599ae0a7818a8ac8f43e5af58b852d18a55724243459e0582b3125a721 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackagecatalog/accesspackageresources"
    i741704e31029f8c4b1ade1c3c152b3734d08209443fa9c0525e33dfc87e5ae8c "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackagecatalog/accesspackageresourceroles/item"
    i9869f160c0a64630a23067ffac5b738a3353dde18defbb7b776ead7c63435a70 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackagecatalog/accesspackageresources/item"
    ib45646ddc06c9c5d86c40a823edd375a56e7ac16f318db8a9032a6004cd309ba "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackagecatalog/accesspackageresourcescopes/item"
    icc55437f342dcb997b3b1b90d83560d39d1edaac916c081ebe4cbd360fc47aeb "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackageassignmentpolicy/accesspackagecatalog/accesspackages/item"
)

// Builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignmentResourceRoles\{accessPackageAssignmentResourceRole-id}\accessPackageAssignments\{accessPackageAssignment-id}\accessPackageAssignmentPolicy\accessPackageCatalog
type AccessPackageCatalogRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type AccessPackageCatalogRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
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
// Get accessPackageCatalog from identityGovernance
type AccessPackageCatalogRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
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
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResourceRoles()(*i37eb72050fb28d46d11ae54b42db111aab72ecf5aa9fd4c470866ba104fce7b1.AccessPackageResourceRolesRequestBuilder) {
    return i37eb72050fb28d46d11ae54b42db111aab72ecf5aa9fd4c470866ba104fce7b1.NewAccessPackageResourceRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.identityGovernance.entitlementManagement.accessPackageAssignmentResourceRoles.item.accessPackageAssignments.item.accessPackageAssignmentPolicy.accessPackageCatalog.accessPackageResourceRoles.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResourceRolesById(id string)(*i741704e31029f8c4b1ade1c3c152b3734d08209443fa9c0525e33dfc87e5ae8c.AccessPackageResourceRoleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRole_id"] = id
    }
    return i741704e31029f8c4b1ade1c3c152b3734d08209443fa9c0525e33dfc87e5ae8c.NewAccessPackageResourceRoleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResources()(*i947925599ae0a7818a8ac8f43e5af58b852d18a55724243459e0582b3125a721.AccessPackageResourcesRequestBuilder) {
    return i947925599ae0a7818a8ac8f43e5af58b852d18a55724243459e0582b3125a721.NewAccessPackageResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.identityGovernance.entitlementManagement.accessPackageAssignmentResourceRoles.item.accessPackageAssignments.item.accessPackageAssignmentPolicy.accessPackageCatalog.accessPackageResources.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResourcesById(id string)(*i9869f160c0a64630a23067ffac5b738a3353dde18defbb7b776ead7c63435a70.AccessPackageResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResource_id"] = id
    }
    return i9869f160c0a64630a23067ffac5b738a3353dde18defbb7b776ead7c63435a70.NewAccessPackageResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResourceScopes()(*i803a7b36907f71072fade484e0cf793a7743c53abed3fdcfd96cca7266386172.AccessPackageResourceScopesRequestBuilder) {
    return i803a7b36907f71072fade484e0cf793a7743c53abed3fdcfd96cca7266386172.NewAccessPackageResourceScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.identityGovernance.entitlementManagement.accessPackageAssignmentResourceRoles.item.accessPackageAssignments.item.accessPackageAssignmentPolicy.accessPackageCatalog.accessPackageResourceScopes.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResourceScopesById(id string)(*ib45646ddc06c9c5d86c40a823edd375a56e7ac16f318db8a9032a6004cd309ba.AccessPackageResourceScopeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceScope_id"] = id
    }
    return ib45646ddc06c9c5d86c40a823edd375a56e7ac16f318db8a9032a6004cd309ba.NewAccessPackageResourceScopeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageCatalogRequestBuilder) AccessPackages()(*i7e2cbe90c8879e155340a710c82fff884137ac87bba0e04dcdc5a342de234272.AccessPackagesRequestBuilder) {
    return i7e2cbe90c8879e155340a710c82fff884137ac87bba0e04dcdc5a342de234272.NewAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.identityGovernance.entitlementManagement.accessPackageAssignmentResourceRoles.item.accessPackageAssignments.item.accessPackageAssignmentPolicy.accessPackageCatalog.accessPackages.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AccessPackageCatalogRequestBuilder) AccessPackagesById(id string)(*icc55437f342dcb997b3b1b90d83560d39d1edaac916c081ebe4cbd360fc47aeb.AccessPackageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage_id"] = id
    }
    return icc55437f342dcb997b3b1b90d83560d39d1edaac916c081ebe4cbd360fc47aeb.NewAccessPackageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new AccessPackageCatalogRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAccessPackageCatalogRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageCatalogRequestBuilder) {
    m := &AccessPackageCatalogRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/{accessPackageAssignmentResourceRole_id}/accessPackageAssignments/{accessPackageAssignment_id}/accessPackageAssignmentPolicy/accessPackageCatalog{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new AccessPackageCatalogRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAccessPackageCatalogRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageCatalogRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageCatalogRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete navigation property accessPackageCatalog for identityGovernance
// Parameters:
//  - options : Options for the request
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
// Get accessPackageCatalog from identityGovernance
// Parameters:
//  - options : Options for the request
func (m *AccessPackageCatalogRequestBuilder) CreateGetRequestInformation(options *AccessPackageCatalogRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// Update the navigation property accessPackageCatalog in identityGovernance
// Parameters:
//  - options : Options for the request
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
// Delete navigation property accessPackageCatalog for identityGovernance
// Parameters:
//  - options : Options for the request
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
// Get accessPackageCatalog from identityGovernance
// Parameters:
//  - options : Options for the request
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
// Update the navigation property accessPackageCatalog in identityGovernance
// Parameters:
//  - options : Options for the request
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
