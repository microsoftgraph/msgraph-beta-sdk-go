package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCreatePasswordSingleSignOnCredentialsRequestBuilder provides operations to call the createPasswordSingleSignOnCredentials method.
type ItemCreatePasswordSingleSignOnCredentialsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCreatePasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCreatePasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCreatePasswordSingleSignOnCredentialsRequestBuilderInternal instantiates a new CreatePasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewItemCreatePasswordSingleSignOnCredentialsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCreatePasswordSingleSignOnCredentialsRequestBuilder) {
    m := &ItemCreatePasswordSingleSignOnCredentialsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/createPasswordSingleSignOnCredentials", pathParameters),
    }
    return m
}
// NewItemCreatePasswordSingleSignOnCredentialsRequestBuilder instantiates a new CreatePasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewItemCreatePasswordSingleSignOnCredentialsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCreatePasswordSingleSignOnCredentialsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCreatePasswordSingleSignOnCredentialsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create single sign-on credentials using a password for a user or group. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipal-createpasswordsinglesignoncredentials?view=graph-rest-1.0
func (m *ItemCreatePasswordSingleSignOnCredentialsRequestBuilder) Post(ctx context.Context, body ItemCreatePasswordSingleSignOnCredentialsPostRequestBodyable, requestConfiguration *ItemCreatePasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PasswordSingleSignOnCredentialSetable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePasswordSingleSignOnCredentialSetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PasswordSingleSignOnCredentialSetable), nil
}
// ToPostRequestInformation create single sign-on credentials using a password for a user or group. This API is available in the following national cloud deployments.
func (m *ItemCreatePasswordSingleSignOnCredentialsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCreatePasswordSingleSignOnCredentialsPostRequestBodyable, requestConfiguration *ItemCreatePasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemCreatePasswordSingleSignOnCredentialsRequestBuilder) WithUrl(rawUrl string)(*ItemCreatePasswordSingleSignOnCredentialsRequestBuilder) {
    return NewItemCreatePasswordSingleSignOnCredentialsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
