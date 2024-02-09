package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ZebraFotaDeploymentsItemCancelRequestBuilder provides operations to call the cancel method.
type ZebraFotaDeploymentsItemCancelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ZebraFotaDeploymentsItemCancelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebraFotaDeploymentsItemCancelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewZebraFotaDeploymentsItemCancelRequestBuilderInternal instantiates a new ZebraFotaDeploymentsItemCancelRequestBuilder and sets the default values.
func NewZebraFotaDeploymentsItemCancelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebraFotaDeploymentsItemCancelRequestBuilder) {
    m := &ZebraFotaDeploymentsItemCancelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/zebraFotaDeployments/{zebraFotaDeployment%2Did}/cancel", pathParameters),
    }
    return m
}
// NewZebraFotaDeploymentsItemCancelRequestBuilder instantiates a new ZebraFotaDeploymentsItemCancelRequestBuilder and sets the default values.
func NewZebraFotaDeploymentsItemCancelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebraFotaDeploymentsItemCancelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewZebraFotaDeploymentsItemCancelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action cancel
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a ZebraFotaDeploymentsItemCancelResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebraFotaDeploymentsItemCancelRequestBuilder) Post(ctx context.Context, requestConfiguration *ZebraFotaDeploymentsItemCancelRequestBuilderPostRequestConfiguration)(ZebraFotaDeploymentsItemCancelResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateZebraFotaDeploymentsItemCancelResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ZebraFotaDeploymentsItemCancelResponseable), nil
}
// PostAsCancelPostResponse invoke action cancel
// returns a ZebraFotaDeploymentsItemCancelPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebraFotaDeploymentsItemCancelRequestBuilder) PostAsCancelPostResponse(ctx context.Context, requestConfiguration *ZebraFotaDeploymentsItemCancelRequestBuilderPostRequestConfiguration)(ZebraFotaDeploymentsItemCancelPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateZebraFotaDeploymentsItemCancelPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ZebraFotaDeploymentsItemCancelPostResponseable), nil
}
// ToPostRequestInformation invoke action cancel
// returns a *RequestInformation when successful
func (m *ZebraFotaDeploymentsItemCancelRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ZebraFotaDeploymentsItemCancelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ZebraFotaDeploymentsItemCancelRequestBuilder when successful
func (m *ZebraFotaDeploymentsItemCancelRequestBuilder) WithUrl(rawUrl string)(*ZebraFotaDeploymentsItemCancelRequestBuilder) {
    return NewZebraFotaDeploymentsItemCancelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
