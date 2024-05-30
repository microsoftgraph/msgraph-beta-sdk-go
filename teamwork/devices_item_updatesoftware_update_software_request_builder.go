package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicesItemUpdatesoftwareUpdateSoftwareRequestBuilder provides operations to call the updateSoftware method.
type DevicesItemUpdatesoftwareUpdateSoftwareRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicesItemUpdatesoftwareUpdateSoftwareRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicesItemUpdatesoftwareUpdateSoftwareRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDevicesItemUpdatesoftwareUpdateSoftwareRequestBuilderInternal instantiates a new DevicesItemUpdatesoftwareUpdateSoftwareRequestBuilder and sets the default values.
func NewDevicesItemUpdatesoftwareUpdateSoftwareRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicesItemUpdatesoftwareUpdateSoftwareRequestBuilder) {
    m := &DevicesItemUpdatesoftwareUpdateSoftwareRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/devices/{teamworkDevice%2Did}/updateSoftware", pathParameters),
    }
    return m
}
// NewDevicesItemUpdatesoftwareUpdateSoftwareRequestBuilder instantiates a new DevicesItemUpdatesoftwareUpdateSoftwareRequestBuilder and sets the default values.
func NewDevicesItemUpdatesoftwareUpdateSoftwareRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicesItemUpdatesoftwareUpdateSoftwareRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicesItemUpdatesoftwareUpdateSoftwareRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update the software for a Microsoft Teams-enabled device. This API triggers a long-running operation.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/teamworkdevice-updatesoftware?view=graph-rest-beta
func (m *DevicesItemUpdatesoftwareUpdateSoftwareRequestBuilder) Post(ctx context.Context, body DevicesItemUpdatesoftwareUpdateSoftwarePostRequestBodyable, requestConfiguration *DevicesItemUpdatesoftwareUpdateSoftwareRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation update the software for a Microsoft Teams-enabled device. This API triggers a long-running operation.
// returns a *RequestInformation when successful
func (m *DevicesItemUpdatesoftwareUpdateSoftwareRequestBuilder) ToPostRequestInformation(ctx context.Context, body DevicesItemUpdatesoftwareUpdateSoftwarePostRequestBodyable, requestConfiguration *DevicesItemUpdatesoftwareUpdateSoftwareRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DevicesItemUpdatesoftwareUpdateSoftwareRequestBuilder when successful
func (m *DevicesItemUpdatesoftwareUpdateSoftwareRequestBuilder) WithUrl(rawUrl string)(*DevicesItemUpdatesoftwareUpdateSoftwareRequestBuilder) {
    return NewDevicesItemUpdatesoftwareUpdateSoftwareRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
