package accesspackageresource

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i060502e07deeba1a883ebbad6018d7b4f76ffcf3a05bf6be8e182cf09b44c838 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackage/accesspackageresourcerolescopes/item/accesspackageresourcescope/accesspackageresource/accesspackageresourceenvironment"
    ib462865f9e5e8f2bc4cb78f97a868149355872311287d1c74d9c36cb07db4919 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackage/accesspackageresourcerolescopes/item/accesspackageresourcescope/accesspackageresource/accesspackageresourceroles"
    id22be877333f9ad23150746b4bd321a5e5d87308aa7101ce19f3cdb7743eb550 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackage/accesspackageresourcerolescopes/item/accesspackageresourcescope/accesspackageresource/accesspackageresourcescopes"
    i681d90d8eec5a17df7dfdaebbede522f6b9488d861187f2468e9bd36c142b412 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackage/accesspackageresourcerolescopes/item/accesspackageresourcescope/accesspackageresource/accesspackageresourcescopes/item"
    i87ebf8d654ce45ee296af3150a344be5fb1f30192b5b535e72acc581d7e44ecf "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentresourceroles/item/accesspackageassignments/item/accesspackage/accesspackageresourcerolescopes/item/accesspackageresourcescope/accesspackageresource/accesspackageresourceroles/item"
)

// Builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignmentResourceRoles\{accessPackageAssignmentResourceRole-id}\accessPackageAssignments\{accessPackageAssignment-id}\accessPackage\accessPackageResourceRoleScopes\{accessPackageResourceRoleScope-id}\accessPackageResourceScope\accessPackageResource
type AccessPackageResourceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type AccessPackageResourceRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type AccessPackageResourceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessPackageResourceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Read-only. Nullable.
type AccessPackageResourceRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type AccessPackageResourceRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageResource;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceEnvironment()(*i060502e07deeba1a883ebbad6018d7b4f76ffcf3a05bf6be8e182cf09b44c838.AccessPackageResourceEnvironmentRequestBuilder) {
    return i060502e07deeba1a883ebbad6018d7b4f76ffcf3a05bf6be8e182cf09b44c838.NewAccessPackageResourceEnvironmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceRoles()(*ib462865f9e5e8f2bc4cb78f97a868149355872311287d1c74d9c36cb07db4919.AccessPackageResourceRolesRequestBuilder) {
    return ib462865f9e5e8f2bc4cb78f97a868149355872311287d1c74d9c36cb07db4919.NewAccessPackageResourceRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.identityGovernance.entitlementManagement.accessPackageAssignmentResourceRoles.item.accessPackageAssignments.item.accessPackage.accessPackageResourceRoleScopes.item.accessPackageResourceScope.accessPackageResource.accessPackageResourceRoles.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceRolesById(id string)(*i87ebf8d654ce45ee296af3150a344be5fb1f30192b5b535e72acc581d7e44ecf.AccessPackageResourceRoleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceRole_id"] = id
    }
    return i87ebf8d654ce45ee296af3150a344be5fb1f30192b5b535e72acc581d7e44ecf.NewAccessPackageResourceRoleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceScopes()(*id22be877333f9ad23150746b4bd321a5e5d87308aa7101ce19f3cdb7743eb550.AccessPackageResourceScopesRequestBuilder) {
    return id22be877333f9ad23150746b4bd321a5e5d87308aa7101ce19f3cdb7743eb550.NewAccessPackageResourceScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.identityGovernance.entitlementManagement.accessPackageAssignmentResourceRoles.item.accessPackageAssignments.item.accessPackage.accessPackageResourceRoleScopes.item.accessPackageResourceScope.accessPackageResource.accessPackageResourceScopes.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AccessPackageResourceRequestBuilder) AccessPackageResourceScopesById(id string)(*i681d90d8eec5a17df7dfdaebbede522f6b9488d861187f2468e9bd36c142b412.AccessPackageResourceScopeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageResourceScope_id"] = id
    }
    return i681d90d8eec5a17df7dfdaebbede522f6b9488d861187f2468e9bd36c142b412.NewAccessPackageResourceScopeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new AccessPackageResourceRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAccessPackageResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageResourceRequestBuilder) {
    m := &AccessPackageResourceRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/identityGovernance/entitlementManagement/accessPackageAssignmentResourceRoles/{accessPackageAssignmentResourceRole_id}/accessPackageAssignments/{accessPackageAssignment_id}/accessPackage/accessPackageResourceRoleScopes/{accessPackageResourceRoleScope_id}/accessPackageResourceScope/accessPackageResource{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new AccessPackageResourceRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAccessPackageResourceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageResourceRequestBuilderInternal(urlParams, requestAdapter)
}
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *AccessPackageResourceRequestBuilder) CreateDeleteRequestInformation(options *AccessPackageResourceRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *AccessPackageResourceRequestBuilder) CreateGetRequestInformation(options *AccessPackageResourceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *AccessPackageResourceRequestBuilder) CreatePatchRequestInformation(options *AccessPackageResourceRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *AccessPackageResourceRequestBuilder) Delete(options *AccessPackageResourceRequestBuilderDeleteOptions)(error) {
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
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *AccessPackageResourceRequestBuilder) Get(options *AccessPackageResourceRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageResource, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAccessPackageResource() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageResource), nil
}
// Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *AccessPackageResourceRequestBuilder) Patch(options *AccessPackageResourceRequestBuilderPatchOptions)(error) {
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
