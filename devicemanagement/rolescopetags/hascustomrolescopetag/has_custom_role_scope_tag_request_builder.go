package hascustomrolescopetag

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// HasCustomRoleScopeTagRequestBuilder provides operations to call the hasCustomRoleScopeTag method.
type HasCustomRoleScopeTagRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// HasCustomRoleScopeTagRequestBuilderGetOptions options for Get
type HasCustomRoleScopeTagRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewHasCustomRoleScopeTagRequestBuilderInternal instantiates a new HasCustomRoleScopeTagRequestBuilder and sets the default values.
func NewHasCustomRoleScopeTagRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*HasCustomRoleScopeTagRequestBuilder) {
    m := &HasCustomRoleScopeTagRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/roleScopeTags/microsoft.graph.hasCustomRoleScopeTag()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewHasCustomRoleScopeTagRequestBuilder instantiates a new HasCustomRoleScopeTagRequestBuilder and sets the default values.
func NewHasCustomRoleScopeTagRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*HasCustomRoleScopeTagRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHasCustomRoleScopeTagRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function hasCustomRoleScopeTag
func (m *HasCustomRoleScopeTagRequestBuilder) CreateGetRequestInformation(options *HasCustomRoleScopeTagRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function hasCustomRoleScopeTag
func (m *HasCustomRoleScopeTagRequestBuilder) Get(options *HasCustomRoleScopeTagRequestBuilderGetOptions)(HasCustomRoleScopeTagResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateHasCustomRoleScopeTagResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(HasCustomRoleScopeTagResponseable), nil
}
