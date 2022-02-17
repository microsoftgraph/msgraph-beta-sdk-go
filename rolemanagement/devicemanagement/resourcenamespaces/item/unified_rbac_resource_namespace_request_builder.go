package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i364de94ee12d8f96dc648e7c3741312fd56499dc3a23d0298bc8da59cf990f04 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/resourcenamespaces/item/importresourceactions"
    i8eb06311e73f2c0da7a5efa55417c3517d5b7b351f6b98ff69e266e1d0c56fa8 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/resourcenamespaces/item/resourceactions"
    i28bbdad243bea8433bbcaf8db4e3551dcedd77a214aab3105a7ef3ad6592115a "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/resourcenamespaces/item/resourceactions/item"
)

// UnifiedRbacResourceNamespaceRequestBuilder builds and executes requests for operations under \roleManagement\deviceManagement\resourceNamespaces\{unifiedRbacResourceNamespace-id}
type UnifiedRbacResourceNamespaceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UnifiedRbacResourceNamespaceRequestBuilderDeleteOptions options for Delete
type UnifiedRbacResourceNamespaceRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRbacResourceNamespaceRequestBuilderGetOptions options for Get
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
// UnifiedRbacResourceNamespaceRequestBuilderGetQueryParameters get resourceNamespaces from roleManagement
type UnifiedRbacResourceNamespaceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UnifiedRbacResourceNamespaceRequestBuilderPatchOptions options for Patch
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
// NewUnifiedRbacResourceNamespaceRequestBuilderInternal instantiates a new UnifiedRbacResourceNamespaceRequestBuilder and sets the default values.
func NewUnifiedRbacResourceNamespaceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRbacResourceNamespaceRequestBuilder) {
    m := &UnifiedRbacResourceNamespaceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/deviceManagement/resourceNamespaces/{unifiedRbacResourceNamespace_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnifiedRbacResourceNamespaceRequestBuilder instantiates a new UnifiedRbacResourceNamespaceRequestBuilder and sets the default values.
func NewUnifiedRbacResourceNamespaceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRbacResourceNamespaceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRbacResourceNamespaceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property resourceNamespaces for roleManagement
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
// CreateGetRequestInformation get resourceNamespaces from roleManagement
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
// CreatePatchRequestInformation update the navigation property resourceNamespaces in roleManagement
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
// Delete delete navigation property resourceNamespaces for roleManagement
func (m *UnifiedRbacResourceNamespaceRequestBuilder) Delete(options *UnifiedRbacResourceNamespaceRequestBuilderDeleteOptions)(error) {
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
// Get get resourceNamespaces from roleManagement
func (m *UnifiedRbacResourceNamespaceRequestBuilder) Get(options *UnifiedRbacResourceNamespaceRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRbacResourceNamespace, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUnifiedRbacResourceNamespace() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRbacResourceNamespace), nil
}
func (m *UnifiedRbacResourceNamespaceRequestBuilder) ImportResourceActions()(*i364de94ee12d8f96dc648e7c3741312fd56499dc3a23d0298bc8da59cf990f04.ImportResourceActionsRequestBuilder) {
    return i364de94ee12d8f96dc648e7c3741312fd56499dc3a23d0298bc8da59cf990f04.NewImportResourceActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property resourceNamespaces in roleManagement
func (m *UnifiedRbacResourceNamespaceRequestBuilder) Patch(options *UnifiedRbacResourceNamespaceRequestBuilderPatchOptions)(error) {
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
func (m *UnifiedRbacResourceNamespaceRequestBuilder) ResourceActions()(*i8eb06311e73f2c0da7a5efa55417c3517d5b7b351f6b98ff69e266e1d0c56fa8.ResourceActionsRequestBuilder) {
    return i8eb06311e73f2c0da7a5efa55417c3517d5b7b351f6b98ff69e266e1d0c56fa8.NewResourceActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourceActionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.deviceManagement.resourceNamespaces.item.resourceActions.item collection
func (m *UnifiedRbacResourceNamespaceRequestBuilder) ResourceActionsById(id string)(*i28bbdad243bea8433bbcaf8db4e3551dcedd77a214aab3105a7ef3ad6592115a.UnifiedRbacResourceActionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRbacResourceAction_id"] = id
    }
    return i28bbdad243bea8433bbcaf8db4e3551dcedd77a214aab3105a7ef3ad6592115a.NewUnifiedRbacResourceActionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
