package updateaddomainpassword

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// UpdateAdDomainPasswordRequestBuilder builds and executes requests for operations under \deviceManagement\virtualEndpoint\onPremisesConnections\{cloudPcOnPremisesConnection-id}\microsoft.graph.updateAdDomainPassword
type UpdateAdDomainPasswordRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UpdateAdDomainPasswordRequestBuilderPostOptions options for Post
type UpdateAdDomainPasswordRequestBuilderPostOptions struct {
    // 
    Body *UpdateAdDomainPasswordRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUpdateAdDomainPasswordRequestBuilderInternal instantiates a new UpdateAdDomainPasswordRequestBuilder and sets the default values.
func NewUpdateAdDomainPasswordRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdateAdDomainPasswordRequestBuilder) {
    m := &UpdateAdDomainPasswordRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/onPremisesConnections/{cloudPcOnPremisesConnection_id}/microsoft.graph.updateAdDomainPassword";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUpdateAdDomainPasswordRequestBuilder instantiates a new UpdateAdDomainPasswordRequestBuilder and sets the default values.
func NewUpdateAdDomainPasswordRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdateAdDomainPasswordRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUpdateAdDomainPasswordRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action updateAdDomainPassword
func (m *UpdateAdDomainPasswordRequestBuilder) CreatePostRequestInformation(options *UpdateAdDomainPasswordRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action updateAdDomainPassword
func (m *UpdateAdDomainPasswordRequestBuilder) Post(options *UpdateAdDomainPasswordRequestBuilderPostOptions)(error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
