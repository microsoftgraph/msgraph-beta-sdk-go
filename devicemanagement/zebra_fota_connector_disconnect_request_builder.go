package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ZebraFotaConnectorDisconnectRequestBuilder provides operations to call the disconnect method.
type ZebraFotaConnectorDisconnectRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ZebraFotaConnectorDisconnectRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebraFotaConnectorDisconnectRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewZebraFotaConnectorDisconnectRequestBuilderInternal instantiates a new ZebraFotaConnectorDisconnectRequestBuilder and sets the default values.
func NewZebraFotaConnectorDisconnectRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebraFotaConnectorDisconnectRequestBuilder) {
    m := &ZebraFotaConnectorDisconnectRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/zebraFotaConnector/disconnect", pathParameters),
    }
    return m
}
// NewZebraFotaConnectorDisconnectRequestBuilder instantiates a new ZebraFotaConnectorDisconnectRequestBuilder and sets the default values.
func NewZebraFotaConnectorDisconnectRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebraFotaConnectorDisconnectRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewZebraFotaConnectorDisconnectRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action disconnect
// Deprecated: This method is obsolete. Use PostAsDisconnectPostResponse instead.
// returns a ZebraFotaConnectorDisconnectResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebraFotaConnectorDisconnectRequestBuilder) Post(ctx context.Context, requestConfiguration *ZebraFotaConnectorDisconnectRequestBuilderPostRequestConfiguration)(ZebraFotaConnectorDisconnectResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateZebraFotaConnectorDisconnectResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ZebraFotaConnectorDisconnectResponseable), nil
}
// PostAsDisconnectPostResponse invoke action disconnect
// returns a ZebraFotaConnectorDisconnectPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebraFotaConnectorDisconnectRequestBuilder) PostAsDisconnectPostResponse(ctx context.Context, requestConfiguration *ZebraFotaConnectorDisconnectRequestBuilderPostRequestConfiguration)(ZebraFotaConnectorDisconnectPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateZebraFotaConnectorDisconnectPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ZebraFotaConnectorDisconnectPostResponseable), nil
}
// ToPostRequestInformation invoke action disconnect
// returns a *RequestInformation when successful
func (m *ZebraFotaConnectorDisconnectRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ZebraFotaConnectorDisconnectRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ZebraFotaConnectorDisconnectRequestBuilder when successful
func (m *ZebraFotaConnectorDisconnectRequestBuilder) WithUrl(rawUrl string)(*ZebraFotaConnectorDisconnectRequestBuilder) {
    return NewZebraFotaConnectorDisconnectRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
