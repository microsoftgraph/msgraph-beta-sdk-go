package getglobalscripthighestavailableversion

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetGlobalScriptHighestAvailableVersionRequestBuilder provides operations to call the getGlobalScriptHighestAvailableVersion method.
type GetGlobalScriptHighestAvailableVersionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetGlobalScriptHighestAvailableVersionRequestBuilderPostOptions options for Post
type GetGlobalScriptHighestAvailableVersionRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetGlobalScriptHighestAvailableVersionRequestBuilderInternal instantiates a new GetGlobalScriptHighestAvailableVersionRequestBuilder and sets the default values.
func NewGetGlobalScriptHighestAvailableVersionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetGlobalScriptHighestAvailableVersionRequestBuilder) {
    m := &GetGlobalScriptHighestAvailableVersionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceHealthScripts/{deviceHealthScript_id}/microsoft.graph.getGlobalScriptHighestAvailableVersion";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetGlobalScriptHighestAvailableVersionRequestBuilder instantiates a new GetGlobalScriptHighestAvailableVersionRequestBuilder and sets the default values.
func NewGetGlobalScriptHighestAvailableVersionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetGlobalScriptHighestAvailableVersionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetGlobalScriptHighestAvailableVersionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation update the Proprietary Device Health Script
func (m *GetGlobalScriptHighestAvailableVersionRequestBuilder) CreatePostRequestInformation(options *GetGlobalScriptHighestAvailableVersionRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Post update the Proprietary Device Health Script
func (m *GetGlobalScriptHighestAvailableVersionRequestBuilder) Post(options *GetGlobalScriptHighestAvailableVersionRequestBuilderPostOptions)(GetGlobalScriptHighestAvailableVersionResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetGlobalScriptHighestAvailableVersionResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetGlobalScriptHighestAvailableVersionResponseable), nil
}
