package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i022b51f9e3e3184193403e28360545a5de5b30601b8afbe27fa6ded8945ef323 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages"
    i1f9cffb4601faa399fd85c8de6b651cc216602bdc0c0c7fd82237b87ef00c701 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackageresourceroles"
    i667e6adc08d4ab456c2c16fd43f09213855c087d2a0e98a329203bdec41bdb3e "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackageresources"
    ib6402235b2e38e1b151b99f91249214d540a38501988e968d4bff7f93a783b52 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/customaccesspackageworkflowextensions"
    if0463ebb43fa9a1191e7a3c314f9b2b32bfd179be9b037338a33a684390e7cfb "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackageresourcescopes"
    i0094b207b40e8cb09f9d32a6382458515f23aab8a8ae50f238c48be456a919cb "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackageresourceroles/item"
    i65f5dc28e0a490505ec4c25929f8f0cd4ed34f866dad25cd8f8d84164466ba9d "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackageresources/item"
    i7684640c78c5bc7c9ae3afd77219ae3a2dfc37dafc063e341fba91ca6bc63d39 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackages/item"
    i806de8bb91fba217829bd350dbe60cbabafef2c2766ce4d625eb368f342897c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/customaccesspackageworkflowextensions/item"
    ia31d35903f75b5412153928c21cef5d403530da3185ff1b360326a3b48175ea2 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackagecatalogs/item/accesspackageresourcescopes/item"
)

// AccessPackageCatalogRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageCatalogs\{accessPackageCatalog-id}
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
// AccessPackageCatalogRequestBuilderGetQueryParameters represents a group of access packages.
type AccessPackageCatalogRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
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
func (m *AccessPackageCatalogRequestBuilder) AccessPackageResourceRoles()(*i1f9cffb4601faa399fd85c8de6b651cc216602bdc0c0c7fd82237b87ef00c701.AccessPackageResourceRolesRequestBuilder) {
    return i1f9cffb4601faa399fd85c8de6b651cc216602bdc0c0c7fd82237b87ef00c701.NewAccessPackageResourceRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageResourceRolesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageCatalogs.item.accessPackageResourceRoles.item collection
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
// AccessPackageResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageCatalogs.item.accessPackageResources.item collection
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
// AccessPackageResourceScopesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageCatalogs.item.accessPackageResourceScopes.item collection
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
// AccessPackagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageCatalogs.item.accessPackages.item collection
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
// NewAccessPackageCatalogRequestBuilderInternal instantiates a new AccessPackageCatalogRequestBuilder and sets the default values.
func NewAccessPackageCatalogRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageCatalogRequestBuilder) {
    m := &AccessPackageCatalogRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageCatalogs/{accessPackageCatalog_id}{?select,expand}";
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
// CreateDeleteRequestInformation represents a group of access packages.
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
// CreateGetRequestInformation represents a group of access packages.
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
// CreatePatchRequestInformation represents a group of access packages.
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
func (m *AccessPackageCatalogRequestBuilder) CustomAccessPackageWorkflowExtensions()(*ib6402235b2e38e1b151b99f91249214d540a38501988e968d4bff7f93a783b52.CustomAccessPackageWorkflowExtensionsRequestBuilder) {
    return ib6402235b2e38e1b151b99f91249214d540a38501988e968d4bff7f93a783b52.NewCustomAccessPackageWorkflowExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomAccessPackageWorkflowExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageCatalogs.item.customAccessPackageWorkflowExtensions.item collection
func (m *AccessPackageCatalogRequestBuilder) CustomAccessPackageWorkflowExtensionsById(id string)(*i806de8bb91fba217829bd350dbe60cbabafef2c2766ce4d625eb368f342897c0.CustomAccessPackageWorkflowExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customAccessPackageWorkflowExtension_id"] = id
    }
    return i806de8bb91fba217829bd350dbe60cbabafef2c2766ce4d625eb368f342897c0.NewCustomAccessPackageWorkflowExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete represents a group of access packages.
func (m *AccessPackageCatalogRequestBuilder) Delete(options *AccessPackageCatalogRequestBuilderDeleteOptions)(error) {
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
// Get represents a group of access packages.
func (m *AccessPackageCatalogRequestBuilder) Get(options *AccessPackageCatalogRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageCatalog, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAccessPackageCatalog() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageCatalog), nil
}
// Patch represents a group of access packages.
func (m *AccessPackageCatalogRequestBuilder) Patch(options *AccessPackageCatalogRequestBuilderPatchOptions)(error) {
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
