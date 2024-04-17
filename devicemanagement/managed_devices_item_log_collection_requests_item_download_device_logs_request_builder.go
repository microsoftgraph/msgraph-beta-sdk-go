package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder provides operations to call the downloadDeviceLogs method.
type ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilderInternal instantiates a new ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder and sets the default values.
func NewManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder) {
    m := &ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/logCollectionRequests/{deviceLogCollectionResponse%2Did}/downloadDeviceLogs", pathParameters),
    }
    return m
}
// NewManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder instantiates a new ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder and sets the default values.
func NewManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action downloadDeviceLogs
// Deprecated: This method is obsolete. Use PostAsDownloadDeviceLogsPostResponse instead.
// returns a ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder) Post(ctx context.Context, requestConfiguration *ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilderPostRequestConfiguration)(ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsResponseable), nil
}
// PostAsDownloadDeviceLogsPostResponse invoke action downloadDeviceLogs
// returns a ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder) PostAsDownloadDeviceLogsPostResponse(ctx context.Context, requestConfiguration *ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilderPostRequestConfiguration)(ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsPostResponseable), nil
}
// ToPostRequestInformation invoke action downloadDeviceLogs
// returns a *RequestInformation when successful
func (m *ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder when successful
func (m *ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder) WithUrl(rawUrl string)(*ManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder) {
    return NewManagedDevicesItemLogCollectionRequestsItemDownloadDeviceLogsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
