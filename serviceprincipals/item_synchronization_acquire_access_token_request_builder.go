package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSynchronizationAcquireAccessTokenRequestBuilder provides operations to call the acquireAccessToken method.
type ItemSynchronizationAcquireAccessTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSynchronizationAcquireAccessTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationAcquireAccessTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSynchronizationAcquireAccessTokenRequestBuilderInternal instantiates a new AcquireAccessTokenRequestBuilder and sets the default values.
func NewItemSynchronizationAcquireAccessTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationAcquireAccessTokenRequestBuilder) {
    m := &ItemSynchronizationAcquireAccessTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/synchronization/acquireAccessToken", pathParameters),
    }
    return m
}
// NewItemSynchronizationAcquireAccessTokenRequestBuilder instantiates a new AcquireAccessTokenRequestBuilder and sets the default values.
func NewItemSynchronizationAcquireAccessTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationAcquireAccessTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationAcquireAccessTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post acquire an OAuth Access token to authorize the Microsoft Entra provisioning service to provision users into an application.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/synchronization-synchronization-acquireaccesstoken?view=graph-rest-1.0
func (m *ItemSynchronizationAcquireAccessTokenRequestBuilder) Post(ctx context.Context, body ItemSynchronizationAcquireAccessTokenPostRequestBodyable, requestConfiguration *ItemSynchronizationAcquireAccessTokenRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation acquire an OAuth Access token to authorize the Microsoft Entra provisioning service to provision users into an application.
func (m *ItemSynchronizationAcquireAccessTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSynchronizationAcquireAccessTokenPostRequestBodyable, requestConfiguration *ItemSynchronizationAcquireAccessTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemSynchronizationAcquireAccessTokenRequestBuilder) WithUrl(rawUrl string)(*ItemSynchronizationAcquireAccessTokenRequestBuilder) {
    return NewItemSynchronizationAcquireAccessTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
