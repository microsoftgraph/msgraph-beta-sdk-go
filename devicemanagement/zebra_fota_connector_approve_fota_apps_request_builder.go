package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ZebraFotaConnectorApproveFotaAppsRequestBuilder provides operations to call the approveFotaApps method.
type ZebraFotaConnectorApproveFotaAppsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ZebraFotaConnectorApproveFotaAppsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebraFotaConnectorApproveFotaAppsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewZebraFotaConnectorApproveFotaAppsRequestBuilderInternal instantiates a new ZebraFotaConnectorApproveFotaAppsRequestBuilder and sets the default values.
func NewZebraFotaConnectorApproveFotaAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebraFotaConnectorApproveFotaAppsRequestBuilder) {
    m := &ZebraFotaConnectorApproveFotaAppsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/zebraFotaConnector/approveFotaApps", pathParameters),
    }
    return m
}
// NewZebraFotaConnectorApproveFotaAppsRequestBuilder instantiates a new ZebraFotaConnectorApproveFotaAppsRequestBuilder and sets the default values.
func NewZebraFotaConnectorApproveFotaAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebraFotaConnectorApproveFotaAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewZebraFotaConnectorApproveFotaAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action approveFotaApps
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a ZebraFotaConnectorApproveFotaAppsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebraFotaConnectorApproveFotaAppsRequestBuilder) Post(ctx context.Context, requestConfiguration *ZebraFotaConnectorApproveFotaAppsRequestBuilderPostRequestConfiguration)(ZebraFotaConnectorApproveFotaAppsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateZebraFotaConnectorApproveFotaAppsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ZebraFotaConnectorApproveFotaAppsResponseable), nil
}
// PostAsApproveFotaAppsPostResponse invoke action approveFotaApps
// returns a ZebraFotaConnectorApproveFotaAppsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebraFotaConnectorApproveFotaAppsRequestBuilder) PostAsApproveFotaAppsPostResponse(ctx context.Context, requestConfiguration *ZebraFotaConnectorApproveFotaAppsRequestBuilderPostRequestConfiguration)(ZebraFotaConnectorApproveFotaAppsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateZebraFotaConnectorApproveFotaAppsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ZebraFotaConnectorApproveFotaAppsPostResponseable), nil
}
// ToPostRequestInformation invoke action approveFotaApps
// returns a *RequestInformation when successful
func (m *ZebraFotaConnectorApproveFotaAppsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ZebraFotaConnectorApproveFotaAppsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ZebraFotaConnectorApproveFotaAppsRequestBuilder when successful
func (m *ZebraFotaConnectorApproveFotaAppsRequestBuilder) WithUrl(rawUrl string)(*ZebraFotaConnectorApproveFotaAppsRequestBuilder) {
    return NewZebraFotaConnectorApproveFotaAppsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
