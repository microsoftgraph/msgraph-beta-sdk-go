package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i23eeb299d8c7655a5b8f9d91dcc2657946c4c182d94ebf89a3ccfe073344b511 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/resourcenamespaces/item/resourceactions"
    i57882143c214ec8547e9e29f83c78b6ba8a5a9e82bf6f8cf44bb035bec347ebb "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/resourcenamespaces/item/importresourceactions"
    i0ba042109e088f641da71389cf2ee1d6834428aef79b7f6483d9fafa2f43ab7a "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory/resourcenamespaces/item/resourceactions/item"
)

// Builds and executes requests for operations under \roleManagement\directory\resourceNamespaces\{unifiedRbacResourceNamespace-id}
type UnifiedRbacResourceNamespaceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type UnifiedRbacResourceNamespaceRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type UnifiedRbacResourceNamespaceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UnifiedRbacResourceNamespaceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get resourceNamespaces from roleManagement
type UnifiedRbacResourceNamespaceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type UnifiedRbacResourceNamespaceRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRbacResourceNamespace;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new UnifiedRbacResourceNamespaceRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewUnifiedRbacResourceNamespaceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRbacResourceNamespaceRequestBuilder) {
    m := &UnifiedRbacResourceNamespaceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/directory/resourceNamespaces/{unifiedRbacResourceNamespace_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new UnifiedRbacResourceNamespaceRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewUnifiedRbacResourceNamespaceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRbacResourceNamespaceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRbacResourceNamespaceRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete navigation property resourceNamespaces for roleManagement
// Parameters:
//  - options : Options for the request
func (m *UnifiedRbacResourceNamespaceRequestBuilder) CreateDeleteRequestInformation(options *UnifiedRbacResourceNamespaceRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get resourceNamespaces from roleManagement
// Parameters:
//  - options : Options for the request
func (m *UnifiedRbacResourceNamespaceRequestBuilder) CreateGetRequestInformation(options *UnifiedRbacResourceNamespaceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Update the navigation property resourceNamespaces in roleManagement
// Parameters:
//  - options : Options for the request
func (m *UnifiedRbacResourceNamespaceRequestBuilder) CreatePatchRequestInformation(options *UnifiedRbacResourceNamespaceRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete navigation property resourceNamespaces for roleManagement
// Parameters:
//  - options : Options for the request
func (m *UnifiedRbacResourceNamespaceRequestBuilder) Delete(options *UnifiedRbacResourceNamespaceRequestBuilderDeleteOptions)(error) {
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
// Get resourceNamespaces from roleManagement
// Parameters:
//  - options : Options for the request
func (m *UnifiedRbacResourceNamespaceRequestBuilder) Get(options *UnifiedRbacResourceNamespaceRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRbacResourceNamespace, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUnifiedRbacResourceNamespace() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRbacResourceNamespace), nil
}
func (m *UnifiedRbacResourceNamespaceRequestBuilder) ImportResourceActions()(*i57882143c214ec8547e9e29f83c78b6ba8a5a9e82bf6f8cf44bb035bec347ebb.ImportResourceActionsRequestBuilder) {
    return i57882143c214ec8547e9e29f83c78b6ba8a5a9e82bf6f8cf44bb035bec347ebb.NewImportResourceActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Update the navigation property resourceNamespaces in roleManagement
// Parameters:
//  - options : Options for the request
func (m *UnifiedRbacResourceNamespaceRequestBuilder) Patch(options *UnifiedRbacResourceNamespaceRequestBuilderPatchOptions)(error) {
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
func (m *UnifiedRbacResourceNamespaceRequestBuilder) ResourceActions()(*i23eeb299d8c7655a5b8f9d91dcc2657946c4c182d94ebf89a3ccfe073344b511.ResourceActionsRequestBuilder) {
    return i23eeb299d8c7655a5b8f9d91dcc2657946c4c182d94ebf89a3ccfe073344b511.NewResourceActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.directory.resourceNamespaces.item.resourceActions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *UnifiedRbacResourceNamespaceRequestBuilder) ResourceActionsById(id string)(*i0ba042109e088f641da71389cf2ee1d6834428aef79b7f6483d9fafa2f43ab7a.UnifiedRbacResourceActionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRbacResourceAction_id"] = id
    }
    return i0ba042109e088f641da71389cf2ee1d6834428aef79b7f6483d9fafa2f43ab7a.NewUnifiedRbacResourceActionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
