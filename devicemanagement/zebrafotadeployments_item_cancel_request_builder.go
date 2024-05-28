package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ZebrafotadeploymentsItemCancelRequestBuilder provides operations to call the cancel method.
type ZebrafotadeploymentsItemCancelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ZebrafotadeploymentsItemCancelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebrafotadeploymentsItemCancelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewZebrafotadeploymentsItemCancelRequestBuilderInternal instantiates a new ZebrafotadeploymentsItemCancelRequestBuilder and sets the default values.
func NewZebrafotadeploymentsItemCancelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebrafotadeploymentsItemCancelRequestBuilder) {
    m := &ZebrafotadeploymentsItemCancelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/zebraFotaDeployments/{zebraFotaDeployment%2Did}/cancel", pathParameters),
    }
    return m
}
// NewZebrafotadeploymentsItemCancelRequestBuilder instantiates a new ZebrafotadeploymentsItemCancelRequestBuilder and sets the default values.
func NewZebrafotadeploymentsItemCancelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebrafotadeploymentsItemCancelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewZebrafotadeploymentsItemCancelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action cancel
// Deprecated: This method is obsolete. Use PostAsCancelPostResponse instead.
// returns a ZebrafotadeploymentsItemCancelResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebrafotadeploymentsItemCancelRequestBuilder) Post(ctx context.Context, requestConfiguration *ZebrafotadeploymentsItemCancelRequestBuilderPostRequestConfiguration)(ZebrafotadeploymentsItemCancelResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateZebrafotadeploymentsItemCancelResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ZebrafotadeploymentsItemCancelResponseable), nil
}
// PostAsCancelPostResponse invoke action cancel
// returns a ZebrafotadeploymentsItemCancelPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebrafotadeploymentsItemCancelRequestBuilder) PostAsCancelPostResponse(ctx context.Context, requestConfiguration *ZebrafotadeploymentsItemCancelRequestBuilderPostRequestConfiguration)(ZebrafotadeploymentsItemCancelPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateZebrafotadeploymentsItemCancelPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ZebrafotadeploymentsItemCancelPostResponseable), nil
}
// ToPostRequestInformation invoke action cancel
// returns a *RequestInformation when successful
func (m *ZebrafotadeploymentsItemCancelRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ZebrafotadeploymentsItemCancelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ZebrafotadeploymentsItemCancelRequestBuilder when successful
func (m *ZebrafotadeploymentsItemCancelRequestBuilder) WithUrl(rawUrl string)(*ZebrafotadeploymentsItemCancelRequestBuilder) {
    return NewZebrafotadeploymentsItemCancelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
