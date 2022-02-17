package deletepasswordsinglesignoncredentials

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// DeletePasswordSingleSignOnCredentialsRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\microsoft.graph.deletePasswordSingleSignOnCredentials
type DeletePasswordSingleSignOnCredentialsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeletePasswordSingleSignOnCredentialsRequestBuilderPostOptions options for Post
type DeletePasswordSingleSignOnCredentialsRequestBuilderPostOptions struct {
    // 
    Body *DeletePasswordSingleSignOnCredentialsRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDeletePasswordSingleSignOnCredentialsRequestBuilderInternal instantiates a new DeletePasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewDeletePasswordSingleSignOnCredentialsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeletePasswordSingleSignOnCredentialsRequestBuilder) {
    m := &DeletePasswordSingleSignOnCredentialsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal_id}/microsoft.graph.deletePasswordSingleSignOnCredentials";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeletePasswordSingleSignOnCredentialsRequestBuilder instantiates a new DeletePasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewDeletePasswordSingleSignOnCredentialsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeletePasswordSingleSignOnCredentialsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletePasswordSingleSignOnCredentialsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action deletePasswordSingleSignOnCredentials
func (m *DeletePasswordSingleSignOnCredentialsRequestBuilder) CreatePostRequestInformation(options *DeletePasswordSingleSignOnCredentialsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action deletePasswordSingleSignOnCredentials
func (m *DeletePasswordSingleSignOnCredentialsRequestBuilder) Post(options *DeletePasswordSingleSignOnCredentialsRequestBuilderPostOptions)(error) {
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
