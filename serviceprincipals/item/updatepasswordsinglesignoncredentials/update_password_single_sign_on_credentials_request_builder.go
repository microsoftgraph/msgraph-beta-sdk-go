package updatepasswordsinglesignoncredentials

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// UpdatePasswordSingleSignOnCredentialsRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\microsoft.graph.updatePasswordSingleSignOnCredentials
type UpdatePasswordSingleSignOnCredentialsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UpdatePasswordSingleSignOnCredentialsRequestBuilderPostOptions options for Post
type UpdatePasswordSingleSignOnCredentialsRequestBuilderPostOptions struct {
    // 
    Body *UpdatePasswordSingleSignOnCredentialsRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUpdatePasswordSingleSignOnCredentialsRequestBuilderInternal instantiates a new UpdatePasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewUpdatePasswordSingleSignOnCredentialsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdatePasswordSingleSignOnCredentialsRequestBuilder) {
    m := &UpdatePasswordSingleSignOnCredentialsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal_id}/microsoft.graph.updatePasswordSingleSignOnCredentials";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUpdatePasswordSingleSignOnCredentialsRequestBuilder instantiates a new UpdatePasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewUpdatePasswordSingleSignOnCredentialsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdatePasswordSingleSignOnCredentialsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUpdatePasswordSingleSignOnCredentialsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action updatePasswordSingleSignOnCredentials
func (m *UpdatePasswordSingleSignOnCredentialsRequestBuilder) CreatePostRequestInformation(options *UpdatePasswordSingleSignOnCredentialsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action updatePasswordSingleSignOnCredentials
func (m *UpdatePasswordSingleSignOnCredentialsRequestBuilder) Post(options *UpdatePasswordSingleSignOnCredentialsRequestBuilderPostOptions)(error) {
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
