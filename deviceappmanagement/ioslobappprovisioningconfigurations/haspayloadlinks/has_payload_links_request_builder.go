package haspayloadlinks

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// HasPayloadLinksRequestBuilder provides operations to call the hasPayloadLinks method.
type HasPayloadLinksRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// HasPayloadLinksRequestBuilderPostOptions options for Post
type HasPayloadLinksRequestBuilderPostOptions struct {
    // 
    Body HasPayloadLinksRequestBodyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewHasPayloadLinksRequestBuilderInternal instantiates a new HasPayloadLinksRequestBuilder and sets the default values.
func NewHasPayloadLinksRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*HasPayloadLinksRequestBuilder) {
    m := &HasPayloadLinksRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/iosLobAppProvisioningConfigurations/microsoft.graph.hasPayloadLinks";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewHasPayloadLinksRequestBuilder instantiates a new HasPayloadLinksRequestBuilder and sets the default values.
func NewHasPayloadLinksRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*HasPayloadLinksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHasPayloadLinksRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action hasPayloadLinks
func (m *HasPayloadLinksRequestBuilder) CreatePostRequestInformation(options *HasPayloadLinksRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Post invoke action hasPayloadLinks
func (m *HasPayloadLinksRequestBuilder) Post(options *HasPayloadLinksRequestBuilderPostOptions)(HasPayloadLinksResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateHasPayloadLinksResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(HasPayloadLinksResponseable), nil
}
