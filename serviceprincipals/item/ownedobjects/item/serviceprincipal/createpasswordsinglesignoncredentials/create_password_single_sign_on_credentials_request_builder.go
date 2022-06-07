package createpasswordsinglesignoncredentials

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CreatePasswordSingleSignOnCredentialsRequestBuilder provides operations to call the createPasswordSingleSignOnCredentials method.
type CreatePasswordSingleSignOnCredentialsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CreatePasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CreatePasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCreatePasswordSingleSignOnCredentialsRequestBuilderInternal instantiates a new CreatePasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewCreatePasswordSingleSignOnCredentialsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CreatePasswordSingleSignOnCredentialsRequestBuilder) {
    m := &CreatePasswordSingleSignOnCredentialsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/ownedObjects/{directoryObject%2Did}/microsoft.graph.servicePrincipal/microsoft.graph.createPasswordSingleSignOnCredentials";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCreatePasswordSingleSignOnCredentialsRequestBuilder instantiates a new CreatePasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewCreatePasswordSingleSignOnCredentialsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CreatePasswordSingleSignOnCredentialsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCreatePasswordSingleSignOnCredentialsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action createPasswordSingleSignOnCredentials
func (m *CreatePasswordSingleSignOnCredentialsRequestBuilder) CreatePostRequestInformation(body CreatePasswordSingleSignOnCredentialsPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action createPasswordSingleSignOnCredentials
func (m *CreatePasswordSingleSignOnCredentialsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body CreatePasswordSingleSignOnCredentialsPostRequestBodyable, requestConfiguration *CreatePasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action createPasswordSingleSignOnCredentials
func (m *CreatePasswordSingleSignOnCredentialsRequestBuilder) Post(body CreatePasswordSingleSignOnCredentialsPostRequestBodyable)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PasswordSingleSignOnCredentialSetable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action createPasswordSingleSignOnCredentials
func (m *CreatePasswordSingleSignOnCredentialsRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body CreatePasswordSingleSignOnCredentialsPostRequestBodyable, requestConfiguration *CreatePasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PasswordSingleSignOnCredentialSetable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePasswordSingleSignOnCredentialSetFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PasswordSingleSignOnCredentialSetable), nil
}
