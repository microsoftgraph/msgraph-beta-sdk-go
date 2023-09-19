package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemDeviceEnrollmentConfigurationsItemAssignRequestBuilder provides operations to call the assign method.
type ItemDeviceEnrollmentConfigurationsItemAssignRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDeviceEnrollmentConfigurationsItemAssignRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDeviceEnrollmentConfigurationsItemAssignRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemDeviceEnrollmentConfigurationsItemAssignRequestBuilderInternal instantiates a new AssignRequestBuilder and sets the default values.
func NewItemDeviceEnrollmentConfigurationsItemAssignRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDeviceEnrollmentConfigurationsItemAssignRequestBuilder) {
    m := &ItemDeviceEnrollmentConfigurationsItemAssignRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/deviceEnrollmentConfigurations/{deviceEnrollmentConfiguration%2Did}/assign", pathParameters),
    }
    return m
}
// NewItemDeviceEnrollmentConfigurationsItemAssignRequestBuilder instantiates a new AssignRequestBuilder and sets the default values.
func NewItemDeviceEnrollmentConfigurationsItemAssignRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDeviceEnrollmentConfigurationsItemAssignRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDeviceEnrollmentConfigurationsItemAssignRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action assign
func (m *ItemDeviceEnrollmentConfigurationsItemAssignRequestBuilder) Post(ctx context.Context, body ItemDeviceEnrollmentConfigurationsItemAssignPostRequestBodyable, requestConfiguration *ItemDeviceEnrollmentConfigurationsItemAssignRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action assign
func (m *ItemDeviceEnrollmentConfigurationsItemAssignRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemDeviceEnrollmentConfigurationsItemAssignPostRequestBodyable, requestConfiguration *ItemDeviceEnrollmentConfigurationsItemAssignRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemDeviceEnrollmentConfigurationsItemAssignRequestBuilder) WithUrl(rawUrl string)(*ItemDeviceEnrollmentConfigurationsItemAssignRequestBuilder) {
    return NewItemDeviceEnrollmentConfigurationsItemAssignRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
