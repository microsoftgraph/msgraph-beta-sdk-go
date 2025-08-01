// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilder provides operations to call the retrieveZebraFotaDeviceModels method.
type ZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilderGetQueryParameters invoke function retrieveZebraFotaDeviceModels
type ZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilderGetQueryParameters
}
// NewZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilderInternal instantiates a new ZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilder and sets the default values.
func NewZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilder) {
    m := &ZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/zebraFotaConnector/retrieveZebraFotaDeviceModels(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilder instantiates a new ZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilder and sets the default values.
func NewZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function retrieveZebraFotaDeviceModels
// Deprecated: This method is obsolete. Use GetAsRetrieveZebraFotaDeviceModelsGetResponse instead.
// returns a ZebraFotaConnectorRetrieveZebraFotaDeviceModelsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilder) Get(ctx context.Context, requestConfiguration *ZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilderGetRequestConfiguration)(ZebraFotaConnectorRetrieveZebraFotaDeviceModelsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateZebraFotaConnectorRetrieveZebraFotaDeviceModelsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ZebraFotaConnectorRetrieveZebraFotaDeviceModelsResponseable), nil
}
// GetAsRetrieveZebraFotaDeviceModelsGetResponse invoke function retrieveZebraFotaDeviceModels
// returns a ZebraFotaConnectorRetrieveZebraFotaDeviceModelsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilder) GetAsRetrieveZebraFotaDeviceModelsGetResponse(ctx context.Context, requestConfiguration *ZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilderGetRequestConfiguration)(ZebraFotaConnectorRetrieveZebraFotaDeviceModelsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateZebraFotaConnectorRetrieveZebraFotaDeviceModelsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ZebraFotaConnectorRetrieveZebraFotaDeviceModelsGetResponseable), nil
}
// ToGetRequestInformation invoke function retrieveZebraFotaDeviceModels
// returns a *RequestInformation when successful
func (m *ZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilder when successful
func (m *ZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilder) WithUrl(rawUrl string)(*ZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilder) {
    return NewZebraFotaConnectorRetrieveZebraFotaDeviceModelsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
