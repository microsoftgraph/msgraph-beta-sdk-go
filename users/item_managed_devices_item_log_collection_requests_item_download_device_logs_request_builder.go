package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder provides operations to call the downloadDeviceLogs method.
type ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilderInternal instantiates a new ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder and sets the default values.
func NewItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder) {
    m := &ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/logCollectionRequests/{deviceLogCollectionResponse%2Did}/downloadDeviceLogs", pathParameters),
    }
    return m
}
// NewItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder instantiates a new ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder and sets the default values.
func NewItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action downloadDeviceLogs
// Deprecated: This method is obsolete. Use PostAsDownloadDeviceLogsPostResponse instead.
// returns a ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilderPostRequestConfiguration)(ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsResponseable), nil
}
// PostAsDownloadDeviceLogsPostResponse invoke action downloadDeviceLogs
// returns a ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder) PostAsDownloadDeviceLogsPostResponse(ctx context.Context, requestConfiguration *ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilderPostRequestConfiguration)(ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsPostResponseable), nil
}
// ToPostRequestInformation invoke action downloadDeviceLogs
// returns a *RequestInformation when successful
func (m *ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder when successful
func (m *ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder) WithUrl(rawUrl string)(*ItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder) {
    return NewItemManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
