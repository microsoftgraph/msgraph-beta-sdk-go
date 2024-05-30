package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilder provides operations to call the updateStatus method.
type DeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilderInternal instantiates a new DeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilder and sets the default values.
func NewDeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilder) {
    m := &DeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/deviceAppManagementTasks/{deviceAppManagementTask%2Did}/updateStatus", pathParameters),
    }
    return m
}
// NewDeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilder instantiates a new DeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilder and sets the default values.
func NewDeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set the task's status and attach a note.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilder) Post(ctx context.Context, body DeviceappmanagementtasksItemUpdatestatusUpdateStatusPostRequestBodyable, requestConfiguration *DeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation set the task's status and attach a note.
// returns a *RequestInformation when successful
func (m *DeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeviceappmanagementtasksItemUpdatestatusUpdateStatusPostRequestBodyable, requestConfiguration *DeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilder when successful
func (m *DeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilder) WithUrl(rawUrl string)(*DeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilder) {
    return NewDeviceappmanagementtasksItemUpdatestatusUpdateStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
