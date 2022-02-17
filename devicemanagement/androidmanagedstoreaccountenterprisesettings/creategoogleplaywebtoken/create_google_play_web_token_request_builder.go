package creategoogleplaywebtoken

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// CreateGooglePlayWebTokenRequestBuilder builds and executes requests for operations under \deviceManagement\androidManagedStoreAccountEnterpriseSettings\microsoft.graph.createGooglePlayWebToken
type CreateGooglePlayWebTokenRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CreateGooglePlayWebTokenRequestBuilderPostOptions options for Post
type CreateGooglePlayWebTokenRequestBuilderPostOptions struct {
    // 
    Body *CreateGooglePlayWebTokenRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewCreateGooglePlayWebTokenRequestBuilderInternal instantiates a new CreateGooglePlayWebTokenRequestBuilder and sets the default values.
func NewCreateGooglePlayWebTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CreateGooglePlayWebTokenRequestBuilder) {
    m := &CreateGooglePlayWebTokenRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/androidManagedStoreAccountEnterpriseSettings/microsoft.graph.createGooglePlayWebToken";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCreateGooglePlayWebTokenRequestBuilder instantiates a new CreateGooglePlayWebTokenRequestBuilder and sets the default values.
func NewCreateGooglePlayWebTokenRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CreateGooglePlayWebTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCreateGooglePlayWebTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation generates a web token that is used in an embeddable component.
func (m *CreateGooglePlayWebTokenRequestBuilder) CreatePostRequestInformation(options *CreateGooglePlayWebTokenRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post generates a web token that is used in an embeddable component.
func (m *CreateGooglePlayWebTokenRequestBuilder) Post(options *CreateGooglePlayWebTokenRequestBuilderPostOptions)(*string, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(*requestInfo, "string", nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*string), nil
}
