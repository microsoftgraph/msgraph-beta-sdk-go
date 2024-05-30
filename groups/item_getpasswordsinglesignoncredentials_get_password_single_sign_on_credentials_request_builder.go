package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilder provides operations to call the getPasswordSingleSignOnCredentials method.
type ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilderInternal instantiates a new ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilder) {
    m := &ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/getPasswordSingleSignOnCredentials", pathParameters),
    }
    return m
}
// NewItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilder instantiates a new ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getPasswordSingleSignOnCredentials
// Deprecated: This method is obsolete. Use PostAsGetPasswordSingleSignOnCredentialsPostResponse instead.
// returns a ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration)(ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsResponseable), nil
}
// PostAsGetPasswordSingleSignOnCredentialsPostResponse invoke action getPasswordSingleSignOnCredentials
// returns a ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilder) PostAsGetPasswordSingleSignOnCredentialsPostResponse(ctx context.Context, requestConfiguration *ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration)(ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsPostResponseable), nil
}
// ToPostRequestInformation invoke action getPasswordSingleSignOnCredentials
// returns a *RequestInformation when successful
func (m *ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilder when successful
func (m *ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilder) WithUrl(rawUrl string)(*ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilder) {
    return NewItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
