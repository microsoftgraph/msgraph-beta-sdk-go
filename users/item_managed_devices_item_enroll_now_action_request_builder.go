package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManagedDevicesItemEnrollNowActionRequestBuilder provides operations to call the enrollNowAction method.
type ItemManagedDevicesItemEnrollNowActionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManagedDevicesItemEnrollNowActionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesItemEnrollNowActionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManagedDevicesItemEnrollNowActionRequestBuilderInternal instantiates a new ItemManagedDevicesItemEnrollNowActionRequestBuilder and sets the default values.
func NewItemManagedDevicesItemEnrollNowActionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemEnrollNowActionRequestBuilder) {
    m := &ItemManagedDevicesItemEnrollNowActionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/enrollNowAction", pathParameters),
    }
    return m
}
// NewItemManagedDevicesItemEnrollNowActionRequestBuilder instantiates a new ItemManagedDevicesItemEnrollNowActionRequestBuilder and sets the default values.
func NewItemManagedDevicesItemEnrollNowActionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemEnrollNowActionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesItemEnrollNowActionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post trigger comanagement enrollment action on ConfigurationManager client
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManagedDevicesItemEnrollNowActionRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemManagedDevicesItemEnrollNowActionRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation trigger comanagement enrollment action on ConfigurationManager client
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemManagedDevicesItemEnrollNowActionRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemManagedDevicesItemEnrollNowActionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemManagedDevicesItemEnrollNowActionRequestBuilder when successful
func (m *ItemManagedDevicesItemEnrollNowActionRequestBuilder) WithUrl(rawUrl string)(*ItemManagedDevicesItemEnrollNowActionRequestBuilder) {
    return NewItemManagedDevicesItemEnrollNowActionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
