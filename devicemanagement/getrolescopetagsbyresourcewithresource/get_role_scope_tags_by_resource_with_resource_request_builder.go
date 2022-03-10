package getrolescopetagsbyresourcewithresource

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetRoleScopeTagsByResourceWithResourceRequestBuilder provides operations to call the getRoleScopeTagsByResource method.
type GetRoleScopeTagsByResourceWithResourceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetRoleScopeTagsByResourceWithResourceRequestBuilderGetOptions options for Get
type GetRoleScopeTagsByResourceWithResourceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetRoleScopeTagsByResourceWithResourceRequestBuilderInternal instantiates a new GetRoleScopeTagsByResourceWithResourceRequestBuilder and sets the default values.
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
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetRoleScopeTagsByResourceWithResourceRequestBuilder instantiates a new GetRoleScopeTagsByResourceWithResourceRequestBuilder and sets the default values.
func NewGetRoleScopeTagsByResourceWithResourceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetRoleScopeTagsByResourceWithResourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetRoleScopeTagsByResourceWithResourceRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getRoleScopeTagsByResource
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
// Get invoke function getRoleScopeTagsByResource
func (m *GetRoleScopeTagsByResourceWithResourceRequestBuilder) Get(options *GetRoleScopeTagsByResourceWithResourceRequestBuilderGetOptions)(GetRoleScopeTagsByResourceWithResourceResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetRoleScopeTagsByResourceWithResourceResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetRoleScopeTagsByResourceWithResourceResponseable), nil
}
