package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilder provides operations to call the createDownloadUrl method.
type ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilderInternal instantiates a new ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilder and sets the default values.
func NewItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilder) {
    m := &ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/logCollectionRequests/{deviceLogCollectionResponse%2Did}/createDownloadUrl", pathParameters),
    }
    return m
}
// NewItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilder instantiates a new ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilder and sets the default values.
func NewItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action createDownloadUrl
// Deprecated: This method is obsolete. Use PostAsCreateDownloadUrlPostResponse instead.
// returns a ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilderPostRequestConfiguration)(ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlResponseable), nil
}
// PostAsCreateDownloadUrlPostResponse invoke action createDownloadUrl
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilder) PostAsCreateDownloadUrlPostResponse(ctx context.Context, requestConfiguration *ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilderPostRequestConfiguration)(ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlPostResponseable), nil
}
// ToPostRequestInformation invoke action createDownloadUrl
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilder when successful
func (m *ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilder) WithUrl(rawUrl string)(*ItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilder) {
    return NewItemManagedDevicesItemLogCollectionRequestsItemCreateDownloadUrlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
