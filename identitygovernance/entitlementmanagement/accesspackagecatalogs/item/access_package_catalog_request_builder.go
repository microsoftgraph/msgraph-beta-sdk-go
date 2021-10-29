package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i022b51f9e3e3184193403e28360545a5de5b30601b8afbe27fa6ded8945ef323 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages"
    i1f9cffb4601faa399fd85c8de6b651cc216602bdc0c0c7fd82237b87ef00c701 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackageresourceroles"
    i667e6adc08d4ab456c2c16fd43f09213855c087d2a0e98a329203bdec41bdb3e "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackageresources"
    if0463ebb43fa9a1191e7a3c314f9b2b32bfd179be9b037338a33a684390e7cfb "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackageresourcescopes"
    i0094b207b40e8cb09f9d32a6382458515f23aab8a8ae50f238c48be456a919cb "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackageresourceroles/item"
    i65f5dc28e0a490505ec4c25929f8f0cd4ed34f866dad25cd8f8d84164466ba9d "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackageresources/item"
    i7684640c78c5bc7c9ae3afd77219ae3a2dfc37dafc063e341fba91ca6bc63d39 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item"
    ia31d35903f75b5412153928c21cef5d403530da3185ff1b360326a3b48175ea2 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackageresourcescopes/item"
)

// Builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageCatalogs\{accessPackageCatalog-id}
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
// Get accessPackageCatalogs from identityGovernance
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
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResourceRoles()(*i1f9cffb4601faa399fd85c8de6b651cc216602bdc0c0c7fd82237b87ef00c701.AccessPackageResourceRolesRequestBuilder) {
    return i1f9cffb4601faa399fd85c8de6b651cc216602bdc0c0c7fd82237b87ef00c701.NewAccessPackageResourceRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.identityGovernance.entitlementManagement.accessPackageCatalogs.item.accessPackageResourceRoles.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResourceRolesById(id string)(*i0094b207b40e8cb09f9d32a6382458515f23aab8a8ae50f238c48be456a919cb.AccessPackageResourceRoleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRole_id"] = id
    }
    return i0094b207b40e8cb09f9d32a6382458515f23aab8a8ae50f238c48be456a919cb.NewAccessPackageResourceRoleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResources()(*i667e6adc08d4ab456c2c16fd43f09213855c087d2a0e98a329203bdec41bdb3e.AccessPackageResourcesRequestBuilder) {
    return i667e6adc08d4ab456c2c16fd43f09213855c087d2a0e98a329203bdec41bdb3e.NewAccessPackageResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.identityGovernance.entitlementManagement.accessPackageCatalogs.item.accessPackageResources.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResourcesById(id string)(*i65f5dc28e0a490505ec4c25929f8f0cd4ed34f866dad25cd8f8d84164466ba9d.AccessPackageResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResource_id"] = id
    }
    return i65f5dc28e0a490505ec4c25929f8f0cd4ed34f866dad25cd8f8d84164466ba9d.NewAccessPackageResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResourceScopes()(*if0463ebb43fa9a1191e7a3c314f9b2b32bfd179be9b037338a33a684390e7cfb.AccessPackageResourceScopesRequestBuilder) {
    return if0463ebb43fa9a1191e7a3c314f9b2b32bfd179be9b037338a33a684390e7cfb.NewAccessPackageResourceScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.identityGovernance.entitlementManagement.accessPackageCatalogs.item.accessPackageResourceScopes.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResourceScopesById(id string)(*ia31d35903f75b5412153928c21cef5d403530da3185ff1b360326a3b48175ea2.AccessPackageResourceScopeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceScope_id"] = id
    }
    return ia31d35903f75b5412153928c21cef5d403530da3185ff1b360326a3b48175ea2.NewAccessPackageResourceScopeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageCatalogRequestBuilder) AccessPackages()(*i022b51f9e3e3184193403e28360545a5de5b30601b8afbe27fa6ded8945ef323.AccessPackagesRequestBuilder) {
    return i022b51f9e3e3184193403e28360545a5de5b30601b8afbe27fa6ded8945ef323.NewAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.identityGovernance.entitlementManagement.accessPackageCatalogs.item.accessPackages.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AccessPackageCatalogRequestBuilder) AccessPackagesById(id string)(*i7684640c78c5bc7c9ae3afd77219ae3a2dfc37dafc063e341fba91ca6bc63d39.AccessPackageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage_id"] = id
    }
    return i7684640c78c5bc7c9ae3afd77219ae3a2dfc37dafc063e341fba91ca6bc63d39.NewAccessPackageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new AccessPackageCatalogRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAccessPackageCatalogRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageCatalogRequestBuilder) {
    m := &AccessPackageCatalogRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/identityGovernance/entitlementManagement/accessPackageCatalogs/{accessPackageCatalog_id}{?select,expand}";
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
// Delete navigation property accessPackageCatalogs for identityGovernance
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
// Get accessPackageCatalogs from identityGovernance
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
// Update the navigation property accessPackageCatalogs in identityGovernance
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
// Delete navigation property accessPackageCatalogs for identityGovernance
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
// Get accessPackageCatalogs from identityGovernance
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
// Update the navigation property accessPackageCatalogs in identityGovernance
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
