package getrolescopetagsbyresourcewithresource

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Builds and executes requests for operations under \deviceManagement\microsoft.graph.getRoleScopeTagsByResource(resource='{resource}')
type GetRoleScopeTagsByResourceWithResourceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type GetRoleScopeTagsByResourceWithResourceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new GetRoleScopeTagsByResourceWithResourceRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
//  - resource : Usage: resource={resource}
func NewGetRoleScopeTagsByResourceWithResourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, resource *string)(*GetRoleScopeTagsByResourceWithResourceRequestBuilder) {
    m := &GetRoleScopeTagsByResourceWithResourceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoft.graph.getRoleScopeTagsByResource(resource='{resource}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if resource != nil {
        urlTplParams["resource"] = *resource
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new GetRoleScopeTagsByResourceWithResourceRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewGetRoleScopeTagsByResourceWithResourceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetRoleScopeTagsByResourceWithResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetRoleScopeTagsByResourceWithResourceRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Invoke function getRoleScopeTagsByResource
// Parameters:
//  - options : Options for the request
func (m *GetRoleScopeTagsByResourceWithResourceRequestBuilder) CreateGetRequestInformation(options *GetRoleScopeTagsByResourceWithResourceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
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
// Invoke function getRoleScopeTagsByResource
// Parameters:
//  - options : Options for the request
func (m *GetRoleScopeTagsByResourceWithResourceRequestBuilder) Get(options *GetRoleScopeTagsByResourceWithResourceRequestBuilderGetOptions)([]GetRoleScopeTagsByResourceWithResource, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendCollectionAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGetRoleScopeTagsByResourceWithResource() }, nil)
    if err != nil {
        return nil, err
    }
    val := make([]GetRoleScopeTagsByResourceWithResource, len(res))
    for i, v := range res {
        val[i] = *(v.(*GetRoleScopeTagsByResourceWithResource))
    }
    return val, nil
}
