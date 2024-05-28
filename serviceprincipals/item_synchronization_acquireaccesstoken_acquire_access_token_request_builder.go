package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSynchronizationAcquireaccesstokenAcquireAccessTokenRequestBuilder provides operations to call the acquireAccessToken method.
type ItemSynchronizationAcquireaccesstokenAcquireAccessTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSynchronizationAcquireaccesstokenAcquireAccessTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationAcquireaccesstokenAcquireAccessTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSynchronizationAcquireaccesstokenAcquireAccessTokenRequestBuilderInternal instantiates a new ItemSynchronizationAcquireaccesstokenAcquireAccessTokenRequestBuilder and sets the default values.
func NewItemSynchronizationAcquireaccesstokenAcquireAccessTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationAcquireaccesstokenAcquireAccessTokenRequestBuilder) {
    m := &ItemSynchronizationAcquireaccesstokenAcquireAccessTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/synchronization/acquireAccessToken", pathParameters),
    }
    return m
}
// NewItemSynchronizationAcquireaccesstokenAcquireAccessTokenRequestBuilder instantiates a new ItemSynchronizationAcquireaccesstokenAcquireAccessTokenRequestBuilder and sets the default values.
func NewItemSynchronizationAcquireaccesstokenAcquireAccessTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationAcquireaccesstokenAcquireAccessTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationAcquireaccesstokenAcquireAccessTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post acquire an OAuth Access token to authorize the Microsoft Entra provisioning service to provision users into an application.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/synchronization-synchronization-acquireaccesstoken?view=graph-rest-beta
func (m *ItemSynchronizationAcquireaccesstokenAcquireAccessTokenRequestBuilder) Post(ctx context.Context, body ItemSynchronizationAcquireaccesstokenAcquireAccessTokenPostRequestBodyable, requestConfiguration *ItemSynchronizationAcquireaccesstokenAcquireAccessTokenRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation acquire an OAuth Access token to authorize the Microsoft Entra provisioning service to provision users into an application.
// returns a *RequestInformation when successful
func (m *ItemSynchronizationAcquireaccesstokenAcquireAccessTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSynchronizationAcquireaccesstokenAcquireAccessTokenPostRequestBodyable, requestConfiguration *ItemSynchronizationAcquireaccesstokenAcquireAccessTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSynchronizationAcquireaccesstokenAcquireAccessTokenRequestBuilder when successful
func (m *ItemSynchronizationAcquireaccesstokenAcquireAccessTokenRequestBuilder) WithUrl(rawUrl string)(*ItemSynchronizationAcquireaccesstokenAcquireAccessTokenRequestBuilder) {
    return NewItemSynchronizationAcquireaccesstokenAcquireAccessTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
