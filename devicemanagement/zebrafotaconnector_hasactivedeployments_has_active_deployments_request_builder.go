package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilder provides operations to call the hasActiveDeployments method.
type ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilderInternal instantiates a new ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilder and sets the default values.
func NewZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilder) {
    m := &ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/zebraFotaConnector/hasActiveDeployments", pathParameters),
    }
    return m
}
// NewZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilder instantiates a new ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilder and sets the default values.
func NewZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action hasActiveDeployments
// Deprecated: This method is obsolete. Use PostAsHasActiveDeploymentsPostResponse instead.
// returns a ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilder) Post(ctx context.Context, requestConfiguration *ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilderPostRequestConfiguration)(ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsResponseable), nil
}
// PostAsHasActiveDeploymentsPostResponse invoke action hasActiveDeployments
// returns a ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilder) PostAsHasActiveDeploymentsPostResponse(ctx context.Context, requestConfiguration *ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilderPostRequestConfiguration)(ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsPostResponseable), nil
}
// ToPostRequestInformation invoke action hasActiveDeployments
// returns a *RequestInformation when successful
func (m *ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilder when successful
func (m *ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilder) WithUrl(rawUrl string)(*ZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilder) {
    return NewZebrafotaconnectorHasactivedeploymentsHasActiveDeploymentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
