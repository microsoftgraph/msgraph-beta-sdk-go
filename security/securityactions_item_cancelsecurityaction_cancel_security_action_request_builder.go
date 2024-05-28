package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SecurityactionsItemCancelsecurityactionCancelSecurityActionRequestBuilder provides operations to call the cancelSecurityAction method.
type SecurityactionsItemCancelsecurityactionCancelSecurityActionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SecurityactionsItemCancelsecurityactionCancelSecurityActionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityactionsItemCancelsecurityactionCancelSecurityActionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSecurityactionsItemCancelsecurityactionCancelSecurityActionRequestBuilderInternal instantiates a new SecurityactionsItemCancelsecurityactionCancelSecurityActionRequestBuilder and sets the default values.
func NewSecurityactionsItemCancelsecurityactionCancelSecurityActionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityactionsItemCancelsecurityactionCancelSecurityActionRequestBuilder) {
    m := &SecurityactionsItemCancelsecurityactionCancelSecurityActionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/securityActions/{securityAction%2Did}/cancelSecurityAction", pathParameters),
    }
    return m
}
// NewSecurityactionsItemCancelsecurityactionCancelSecurityActionRequestBuilder instantiates a new SecurityactionsItemCancelsecurityactionCancelSecurityActionRequestBuilder and sets the default values.
func NewSecurityactionsItemCancelsecurityactionCancelSecurityActionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityactionsItemCancelsecurityactionCancelSecurityActionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityactionsItemCancelsecurityactionCancelSecurityActionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post cancel a security operation.
// Deprecated: The legacy Graph Security API is deprecated and will stop returning data on January 31, 2025. Please use the new Graph Security API. as of 2024-01/Deprecation
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/securityaction-cancelsecurityaction?view=graph-rest-beta
func (m *SecurityactionsItemCancelsecurityactionCancelSecurityActionRequestBuilder) Post(ctx context.Context, requestConfiguration *SecurityactionsItemCancelsecurityactionCancelSecurityActionRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation cancel a security operation.
// Deprecated: The legacy Graph Security API is deprecated and will stop returning data on January 31, 2025. Please use the new Graph Security API. as of 2024-01/Deprecation
// returns a *RequestInformation when successful
func (m *SecurityactionsItemCancelsecurityactionCancelSecurityActionRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *SecurityactionsItemCancelsecurityactionCancelSecurityActionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The legacy Graph Security API is deprecated and will stop returning data on January 31, 2025. Please use the new Graph Security API. as of 2024-01/Deprecation
// returns a *SecurityactionsItemCancelsecurityactionCancelSecurityActionRequestBuilder when successful
func (m *SecurityactionsItemCancelsecurityactionCancelSecurityActionRequestBuilder) WithUrl(rawUrl string)(*SecurityactionsItemCancelsecurityactionCancelSecurityActionRequestBuilder) {
    return NewSecurityactionsItemCancelsecurityactionCancelSecurityActionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
