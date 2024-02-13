package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilder provides operations to call the triggerConfigurationManagerAction method.
type ComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilderInternal instantiates a new ComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilder and sets the default values.
func NewComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilder) {
    m := &ComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/triggerConfigurationManagerAction", pathParameters),
    }
    return m
}
// NewComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilder instantiates a new ComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilder and sets the default values.
func NewComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post trigger action on ConfigurationManager client
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilder) Post(ctx context.Context, body ComanagedDevicesItemTriggerConfigurationManagerActionPostRequestBodyable, requestConfiguration *ComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation trigger action on ConfigurationManager client
// returns a *RequestInformation when successful
func (m *ComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ComanagedDevicesItemTriggerConfigurationManagerActionPostRequestBodyable, requestConfiguration *ComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilder when successful
func (m *ComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilder) WithUrl(rawUrl string)(*ComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilder) {
    return NewComanagedDevicesItemTriggerConfigurationManagerActionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
