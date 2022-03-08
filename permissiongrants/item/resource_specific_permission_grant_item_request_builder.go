package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i15947f922198c0bab892f4a7c084ebe5f7cbb43f5394d31b937d5157fbe0edba "github.com/microsoftgraph/msgraph-beta-sdk-go/permissiongrants/item/checkmemberobjects"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i751dbbe15e72bbd1122914748b96ee1a20fca8119351121e2c491922693257d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/permissiongrants/item/getmembergroups"
    ibf0a4d81ee4b8a7e4e3ed9c9fb6d0882fdbab481f7b848c7db240c388dfbec2f "github.com/microsoftgraph/msgraph-beta-sdk-go/permissiongrants/item/getmemberobjects"
    ie01778f41a26871b3f51e401f57da7e144998794d08141d92ee03a581180728e "github.com/microsoftgraph/msgraph-beta-sdk-go/permissiongrants/item/checkmembergroups"
    if1160871f4538f1c2881002e619cb11a8c41ca9ef807c8c5f7a0f39c74706b6d "github.com/microsoftgraph/msgraph-beta-sdk-go/permissiongrants/item/restore"
)

// ResourceSpecificPermissionGrantItemRequestBuilder provides operations to manage the collection of resourceSpecificPermissionGrant entities.
type ResourceSpecificPermissionGrantItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ResourceSpecificPermissionGrantItemRequestBuilderDeleteOptions options for Delete
type ResourceSpecificPermissionGrantItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ResourceSpecificPermissionGrantItemRequestBuilderGetOptions options for Get
type ResourceSpecificPermissionGrantItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ResourceSpecificPermissionGrantItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ResourceSpecificPermissionGrantItemRequestBuilderGetQueryParameters get entity from permissionGrants by key
type ResourceSpecificPermissionGrantItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ResourceSpecificPermissionGrantItemRequestBuilderPatchOptions options for Patch
type ResourceSpecificPermissionGrantItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ResourceSpecificPermissionGrantable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) CheckMemberGroups()(*ie01778f41a26871b3f51e401f57da7e144998794d08141d92ee03a581180728e.CheckMemberGroupsRequestBuilder) {
    return ie01778f41a26871b3f51e401f57da7e144998794d08141d92ee03a581180728e.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) CheckMemberObjects()(*i15947f922198c0bab892f4a7c084ebe5f7cbb43f5394d31b937d5157fbe0edba.CheckMemberObjectsRequestBuilder) {
    return i15947f922198c0bab892f4a7c084ebe5f7cbb43f5394d31b937d5157fbe0edba.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewResourceSpecificPermissionGrantItemRequestBuilderInternal instantiates a new ResourceSpecificPermissionGrantItemRequestBuilder and sets the default values.
func NewResourceSpecificPermissionGrantItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResourceSpecificPermissionGrantItemRequestBuilder) {
    m := &ResourceSpecificPermissionGrantItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/permissionGrants/{resourceSpecificPermissionGrant_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewResourceSpecificPermissionGrantItemRequestBuilder instantiates a new ResourceSpecificPermissionGrantItemRequestBuilder and sets the default values.
func NewResourceSpecificPermissionGrantItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResourceSpecificPermissionGrantItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResourceSpecificPermissionGrantItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from permissionGrants
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) CreateDeleteRequestInformation(options *ResourceSpecificPermissionGrantItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from permissionGrants by key
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) CreateGetRequestInformation(options *ResourceSpecificPermissionGrantItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in permissionGrants
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) CreatePatchRequestInformation(options *ResourceSpecificPermissionGrantItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from permissionGrants
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) Delete(options *ResourceSpecificPermissionGrantItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get entity from permissionGrants by key
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) Get(options *ResourceSpecificPermissionGrantItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ResourceSpecificPermissionGrantable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateResourceSpecificPermissionGrantFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ResourceSpecificPermissionGrantable), nil
}
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) GetMemberGroups()(*i751dbbe15e72bbd1122914748b96ee1a20fca8119351121e2c491922693257d9.GetMemberGroupsRequestBuilder) {
    return i751dbbe15e72bbd1122914748b96ee1a20fca8119351121e2c491922693257d9.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) GetMemberObjects()(*ibf0a4d81ee4b8a7e4e3ed9c9fb6d0882fdbab481f7b848c7db240c388dfbec2f.GetMemberObjectsRequestBuilder) {
    return ibf0a4d81ee4b8a7e4e3ed9c9fb6d0882fdbab481f7b848c7db240c388dfbec2f.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in permissionGrants
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) Patch(options *ResourceSpecificPermissionGrantItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) Restore()(*if1160871f4538f1c2881002e619cb11a8c41ca9ef807c8c5f7a0f39c74706b6d.RestoreRequestBuilder) {
    return if1160871f4538f1c2881002e619cb11a8c41ca9ef807c8c5f7a0f39c74706b6d.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
