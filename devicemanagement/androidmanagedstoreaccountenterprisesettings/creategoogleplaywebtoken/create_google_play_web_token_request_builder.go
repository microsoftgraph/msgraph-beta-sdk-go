package creategoogleplaywebtoken

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CreateGooglePlayWebTokenRequestBuilder provides operations to call the createGooglePlayWebToken method.
type CreateGooglePlayWebTokenRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CreateGooglePlayWebTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CreateGooglePlayWebTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCreateGooglePlayWebTokenRequestBuilderInternal instantiates a new CreateGooglePlayWebTokenRequestBuilder and sets the default values.
func NewCreateGooglePlayWebTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CreateGooglePlayWebTokenRequestBuilder) {
    m := &CreateGooglePlayWebTokenRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/androidManagedStoreAccountEnterpriseSettings/microsoft.graph.createGooglePlayWebToken";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCreateGooglePlayWebTokenRequestBuilder instantiates a new CreateGooglePlayWebTokenRequestBuilder and sets the default values.
func NewCreateGooglePlayWebTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CreateGooglePlayWebTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCreateGooglePlayWebTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation generates a web token that is used in an embeddable component.
func (m *CreateGooglePlayWebTokenRequestBuilder) CreatePostRequestInformation(body CreateGooglePlayWebTokenPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration generates a web token that is used in an embeddable component.
func (m *CreateGooglePlayWebTokenRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body CreateGooglePlayWebTokenPostRequestBodyable, requestConfiguration *CreateGooglePlayWebTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post generates a web token that is used in an embeddable component.
func (m *CreateGooglePlayWebTokenRequestBuilder) Post(body CreateGooglePlayWebTokenPostRequestBodyable)(CreateGooglePlayWebTokenResponseable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler generates a web token that is used in an embeddable component.
func (m *CreateGooglePlayWebTokenRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body CreateGooglePlayWebTokenPostRequestBodyable, requestConfiguration *CreateGooglePlayWebTokenRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(CreateGooglePlayWebTokenResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateCreateGooglePlayWebTokenResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(CreateGooglePlayWebTokenResponseable), nil
}
