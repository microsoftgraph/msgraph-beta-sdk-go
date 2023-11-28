package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder provides operations to call the getCloudPcRemoteActionResults method.
type ItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetQueryParameters invoke function getCloudPcRemoteActionResults
type ItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetQueryParameters struct {
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
// ItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetQueryParameters
}
// NewItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderInternal instantiates a new GetCloudPcRemoteActionResultsRequestBuilder and sets the default values.
func NewItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) {
    m := &ItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/getCloudPcRemoteActionResults(){?%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
    }
    return m
}
// NewItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder instantiates a new GetCloudPcRemoteActionResultsRequestBuilder and sets the default values.
func NewItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getCloudPcRemoteActionResults
// Deprecated: This method is obsolete. Use GetAsGetCloudPcRemoteActionResultsGetResponse instead.
func (m *ItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration)(ItemManagedDevicesItemGetCloudPcRemoteActionResultsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemManagedDevicesItemGetCloudPcRemoteActionResultsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemManagedDevicesItemGetCloudPcRemoteActionResultsResponseable), nil
}
// GetAsGetCloudPcRemoteActionResultsGetResponse invoke function getCloudPcRemoteActionResults
func (m *ItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) GetAsGetCloudPcRemoteActionResultsGetResponse(ctx context.Context, requestConfiguration *ItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration)(ItemManagedDevicesItemGetCloudPcRemoteActionResultsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemManagedDevicesItemGetCloudPcRemoteActionResultsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemManagedDevicesItemGetCloudPcRemoteActionResultsGetResponseable), nil
}
// ToGetRequestInformation invoke function getCloudPcRemoteActionResults
func (m *ItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) WithUrl(rawUrl string)(*ItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) {
    return NewItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
