package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilder provides operations to call the approveFotaApps method.
type ZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilderInternal instantiates a new ZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilder and sets the default values.
func NewZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilder) {
    m := &ZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/zebraFotaConnector/approveFotaApps", pathParameters),
    }
    return m
}
// NewZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilder instantiates a new ZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilder and sets the default values.
func NewZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action approveFotaApps
// Deprecated: This method is obsolete. Use PostAsApproveFotaAppsPostResponse instead.
// returns a ZebrafotaconnectorApprovefotaappsApproveFotaAppsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilder) Post(ctx context.Context, requestConfiguration *ZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilderPostRequestConfiguration)(ZebrafotaconnectorApprovefotaappsApproveFotaAppsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateZebrafotaconnectorApprovefotaappsApproveFotaAppsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ZebrafotaconnectorApprovefotaappsApproveFotaAppsResponseable), nil
}
// PostAsApproveFotaAppsPostResponse invoke action approveFotaApps
// returns a ZebrafotaconnectorApprovefotaappsApproveFotaAppsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilder) PostAsApproveFotaAppsPostResponse(ctx context.Context, requestConfiguration *ZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilderPostRequestConfiguration)(ZebrafotaconnectorApprovefotaappsApproveFotaAppsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateZebrafotaconnectorApprovefotaappsApproveFotaAppsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ZebrafotaconnectorApprovefotaappsApproveFotaAppsPostResponseable), nil
}
// ToPostRequestInformation invoke action approveFotaApps
// returns a *RequestInformation when successful
func (m *ZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilder when successful
func (m *ZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilder) WithUrl(rawUrl string)(*ZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilder) {
    return NewZebrafotaconnectorApprovefotaappsApproveFotaAppsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
