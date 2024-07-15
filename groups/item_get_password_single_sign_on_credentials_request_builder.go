package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemGetPasswordSingleSignOnCredentialsRequestBuilder provides operations to call the getPasswordSingleSignOnCredentials method.
type ItemGetPasswordSingleSignOnCredentialsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGetPasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetPasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemGetPasswordSingleSignOnCredentialsRequestBuilderInternal instantiates a new ItemGetPasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewItemGetPasswordSingleSignOnCredentialsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetPasswordSingleSignOnCredentialsRequestBuilder) {
    m := &ItemGetPasswordSingleSignOnCredentialsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/getPasswordSingleSignOnCredentials", pathParameters),
    }
    return m
}
// NewItemGetPasswordSingleSignOnCredentialsRequestBuilder instantiates a new ItemGetPasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewItemGetPasswordSingleSignOnCredentialsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetPasswordSingleSignOnCredentialsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGetPasswordSingleSignOnCredentialsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get the list of password-based single sign-on credentials for a group. This API returns the encrypted passwords as null.
// Deprecated: This method is obsolete. Use PostAsGetPasswordSingleSignOnCredentialsPostResponse instead.
// returns a ItemGetPasswordSingleSignOnCredentialsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/group-getpasswordsinglesignoncredentials?view=graph-rest-beta
func (m *ItemGetPasswordSingleSignOnCredentialsRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemGetPasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration)(ItemGetPasswordSingleSignOnCredentialsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetPasswordSingleSignOnCredentialsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetPasswordSingleSignOnCredentialsResponseable), nil
}
// PostAsGetPasswordSingleSignOnCredentialsPostResponse get the list of password-based single sign-on credentials for a group. This API returns the encrypted passwords as null.
// returns a ItemGetPasswordSingleSignOnCredentialsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/group-getpasswordsinglesignoncredentials?view=graph-rest-beta
func (m *ItemGetPasswordSingleSignOnCredentialsRequestBuilder) PostAsGetPasswordSingleSignOnCredentialsPostResponse(ctx context.Context, requestConfiguration *ItemGetPasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration)(ItemGetPasswordSingleSignOnCredentialsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetPasswordSingleSignOnCredentialsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetPasswordSingleSignOnCredentialsPostResponseable), nil
}
// ToPostRequestInformation get the list of password-based single sign-on credentials for a group. This API returns the encrypted passwords as null.
// returns a *RequestInformation when successful
func (m *ItemGetPasswordSingleSignOnCredentialsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemGetPasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemGetPasswordSingleSignOnCredentialsRequestBuilder when successful
func (m *ItemGetPasswordSingleSignOnCredentialsRequestBuilder) WithUrl(rawUrl string)(*ItemGetPasswordSingleSignOnCredentialsRequestBuilder) {
    return NewItemGetPasswordSingleSignOnCredentialsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
