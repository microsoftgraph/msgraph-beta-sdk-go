package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementPartnersItemTerminateRequestBuilder provides operations to call the terminate method.
type DeviceManagementPartnersItemTerminateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceManagementPartnersItemTerminateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementPartnersItemTerminateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementPartnersItemTerminateRequestBuilderInternal instantiates a new DeviceManagementPartnersItemTerminateRequestBuilder and sets the default values.
func NewDeviceManagementPartnersItemTerminateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementPartnersItemTerminateRequestBuilder) {
    m := &DeviceManagementPartnersItemTerminateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceManagementPartners/{deviceManagementPartner%2Did}/terminate", pathParameters),
    }
    return m
}
// NewDeviceManagementPartnersItemTerminateRequestBuilder instantiates a new DeviceManagementPartnersItemTerminateRequestBuilder and sets the default values.
func NewDeviceManagementPartnersItemTerminateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementPartnersItemTerminateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementPartnersItemTerminateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action terminate
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceManagementPartnersItemTerminateRequestBuilder) Post(ctx context.Context, requestConfiguration *DeviceManagementPartnersItemTerminateRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action terminate
// returns a *RequestInformation when successful
func (m *DeviceManagementPartnersItemTerminateRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementPartnersItemTerminateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeviceManagementPartnersItemTerminateRequestBuilder when successful
func (m *DeviceManagementPartnersItemTerminateRequestBuilder) WithUrl(rawUrl string)(*DeviceManagementPartnersItemTerminateRequestBuilder) {
    return NewDeviceManagementPartnersItemTerminateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
