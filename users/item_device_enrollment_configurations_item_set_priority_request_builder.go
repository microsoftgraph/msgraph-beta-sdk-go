package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemDeviceEnrollmentConfigurationsItemSetPriorityRequestBuilder provides operations to call the setPriority method.
type ItemDeviceEnrollmentConfigurationsItemSetPriorityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDeviceEnrollmentConfigurationsItemSetPriorityRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDeviceEnrollmentConfigurationsItemSetPriorityRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemDeviceEnrollmentConfigurationsItemSetPriorityRequestBuilderInternal instantiates a new SetPriorityRequestBuilder and sets the default values.
func NewItemDeviceEnrollmentConfigurationsItemSetPriorityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDeviceEnrollmentConfigurationsItemSetPriorityRequestBuilder) {
    m := &ItemDeviceEnrollmentConfigurationsItemSetPriorityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/deviceEnrollmentConfigurations/{deviceEnrollmentConfiguration%2Did}/setPriority", pathParameters),
    }
    return m
}
// NewItemDeviceEnrollmentConfigurationsItemSetPriorityRequestBuilder instantiates a new SetPriorityRequestBuilder and sets the default values.
func NewItemDeviceEnrollmentConfigurationsItemSetPriorityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDeviceEnrollmentConfigurationsItemSetPriorityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDeviceEnrollmentConfigurationsItemSetPriorityRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action setPriority
func (m *ItemDeviceEnrollmentConfigurationsItemSetPriorityRequestBuilder) Post(ctx context.Context, body ItemDeviceEnrollmentConfigurationsItemSetPriorityPostRequestBodyable, requestConfiguration *ItemDeviceEnrollmentConfigurationsItemSetPriorityRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action setPriority
func (m *ItemDeviceEnrollmentConfigurationsItemSetPriorityRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemDeviceEnrollmentConfigurationsItemSetPriorityPostRequestBodyable, requestConfiguration *ItemDeviceEnrollmentConfigurationsItemSetPriorityRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemDeviceEnrollmentConfigurationsItemSetPriorityRequestBuilder) WithUrl(rawUrl string)(*ItemDeviceEnrollmentConfigurationsItemSetPriorityRequestBuilder) {
    return NewItemDeviceEnrollmentConfigurationsItemSetPriorityRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
