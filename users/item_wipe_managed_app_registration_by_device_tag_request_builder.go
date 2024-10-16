package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemWipeManagedAppRegistrationByDeviceTagRequestBuilder provides operations to call the wipeManagedAppRegistrationByDeviceTag method.
type ItemWipeManagedAppRegistrationByDeviceTagRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemWipeManagedAppRegistrationByDeviceTagRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemWipeManagedAppRegistrationByDeviceTagRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal instantiates a new ItemWipeManagedAppRegistrationByDeviceTagRequestBuilder and sets the default values.
func NewItemWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    m := &ItemWipeManagedAppRegistrationByDeviceTagRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/wipeManagedAppRegistrationByDeviceTag", pathParameters),
    }
    return m
}
// NewItemWipeManagedAppRegistrationByDeviceTagRequestBuilder instantiates a new ItemWipeManagedAppRegistrationByDeviceTagRequestBuilder and sets the default values.
func NewItemWipeManagedAppRegistrationByDeviceTagRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(urlParams, requestAdapter)
}
// Post issues a wipe operation on an app registration with specified device tag.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemWipeManagedAppRegistrationByDeviceTagRequestBuilder) Post(ctx context.Context, body ItemWipeManagedAppRegistrationByDeviceTagPostRequestBodyable, requestConfiguration *ItemWipeManagedAppRegistrationByDeviceTagRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation issues a wipe operation on an app registration with specified device tag.
// returns a *RequestInformation when successful
func (m *ItemWipeManagedAppRegistrationByDeviceTagRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemWipeManagedAppRegistrationByDeviceTagPostRequestBodyable, requestConfiguration *ItemWipeManagedAppRegistrationByDeviceTagRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemWipeManagedAppRegistrationByDeviceTagRequestBuilder when successful
func (m *ItemWipeManagedAppRegistrationByDeviceTagRequestBuilder) WithUrl(rawUrl string)(*ItemWipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return NewItemWipeManagedAppRegistrationByDeviceTagRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
