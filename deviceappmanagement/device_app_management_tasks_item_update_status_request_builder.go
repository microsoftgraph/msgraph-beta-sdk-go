package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceAppManagementTasksItemUpdateStatusRequestBuilder provides operations to call the updateStatus method.
type DeviceAppManagementTasksItemUpdateStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceAppManagementTasksItemUpdateStatusRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementTasksItemUpdateStatusRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceAppManagementTasksItemUpdateStatusRequestBuilderInternal instantiates a new UpdateStatusRequestBuilder and sets the default values.
func NewDeviceAppManagementTasksItemUpdateStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementTasksItemUpdateStatusRequestBuilder) {
    m := &DeviceAppManagementTasksItemUpdateStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/deviceAppManagementTasks/{deviceAppManagementTask%2Did}/updateStatus", pathParameters),
    }
    return m
}
// NewDeviceAppManagementTasksItemUpdateStatusRequestBuilder instantiates a new UpdateStatusRequestBuilder and sets the default values.
func NewDeviceAppManagementTasksItemUpdateStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementTasksItemUpdateStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceAppManagementTasksItemUpdateStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set the task's status and attach a note.
func (m *DeviceAppManagementTasksItemUpdateStatusRequestBuilder) Post(ctx context.Context, body DeviceAppManagementTasksItemUpdateStatusPostRequestBodyable, requestConfiguration *DeviceAppManagementTasksItemUpdateStatusRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation set the task's status and attach a note.
func (m *DeviceAppManagementTasksItemUpdateStatusRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeviceAppManagementTasksItemUpdateStatusPostRequestBodyable, requestConfiguration *DeviceAppManagementTasksItemUpdateStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *DeviceAppManagementTasksItemUpdateStatusRequestBuilder) WithUrl(rawUrl string)(*DeviceAppManagementTasksItemUpdateStatusRequestBuilder) {
    return NewDeviceAppManagementTasksItemUpdateStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
