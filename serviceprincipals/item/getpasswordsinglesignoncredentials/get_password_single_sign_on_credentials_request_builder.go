package getpasswordsinglesignoncredentials

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetPasswordSingleSignOnCredentialsRequestBuilder provides operations to call the getPasswordSingleSignOnCredentials method.
type GetPasswordSingleSignOnCredentialsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetPasswordSingleSignOnCredentialsRequestBuilderPostOptions options for Post
type GetPasswordSingleSignOnCredentialsRequestBuilderPostOptions struct {
    // 
    Body GetPasswordSingleSignOnCredentialsRequestBodyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetPasswordSingleSignOnCredentialsRequestBuilderInternal instantiates a new GetPasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewGetPasswordSingleSignOnCredentialsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetPasswordSingleSignOnCredentialsRequestBuilder) {
    m := &GetPasswordSingleSignOnCredentialsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal_id}/microsoft.graph.getPasswordSingleSignOnCredentials";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetPasswordSingleSignOnCredentialsRequestBuilder instantiates a new GetPasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewGetPasswordSingleSignOnCredentialsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetPasswordSingleSignOnCredentialsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetPasswordSingleSignOnCredentialsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getPasswordSingleSignOnCredentials
func (m *GetPasswordSingleSignOnCredentialsRequestBuilder) CreatePostRequestInformation(options *GetPasswordSingleSignOnCredentialsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action getPasswordSingleSignOnCredentials
func (m *GetPasswordSingleSignOnCredentialsRequestBuilder) Post(options *GetPasswordSingleSignOnCredentialsRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PasswordSingleSignOnCredentialSetable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreatePasswordSingleSignOnCredentialSetFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PasswordSingleSignOnCredentialSetable), nil
}
