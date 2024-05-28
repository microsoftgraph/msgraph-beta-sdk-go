package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManageddevicesItemRebootnowRebootNowRequestBuilder provides operations to call the rebootNow method.
type ManageddevicesItemRebootnowRebootNowRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesItemRebootnowRebootNowRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemRebootnowRebootNowRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManageddevicesItemRebootnowRebootNowRequestBuilderInternal instantiates a new ManageddevicesItemRebootnowRebootNowRequestBuilder and sets the default values.
func NewManageddevicesItemRebootnowRebootNowRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemRebootnowRebootNowRequestBuilder) {
    m := &ManageddevicesItemRebootnowRebootNowRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/rebootNow", pathParameters),
    }
    return m
}
// NewManageddevicesItemRebootnowRebootNowRequestBuilder instantiates a new ManageddevicesItemRebootnowRebootNowRequestBuilder and sets the default values.
func NewManageddevicesItemRebootnowRebootNowRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemRebootnowRebootNowRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesItemRebootnowRebootNowRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reboot device
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemRebootnowRebootNowRequestBuilder) Post(ctx context.Context, requestConfiguration *ManageddevicesItemRebootnowRebootNowRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation reboot device
// returns a *RequestInformation when successful
func (m *ManageddevicesItemRebootnowRebootNowRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesItemRebootnowRebootNowRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ManageddevicesItemRebootnowRebootNowRequestBuilder when successful
func (m *ManageddevicesItemRebootnowRebootNowRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesItemRebootnowRebootNowRequestBuilder) {
    return NewManageddevicesItemRebootnowRebootNowRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
