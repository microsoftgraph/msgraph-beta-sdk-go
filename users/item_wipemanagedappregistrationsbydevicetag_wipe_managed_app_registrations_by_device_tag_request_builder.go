package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilder provides operations to call the wipeManagedAppRegistrationsByDeviceTag method.
type ItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal instantiates a new ItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilder and sets the default values.
func NewItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    m := &ItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/wipeManagedAppRegistrationsByDeviceTag", pathParameters),
    }
    return m
}
// NewItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilder instantiates a new ItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilder and sets the default values.
func NewItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(urlParams, requestAdapter)
}
// Post issues a wipe operation on an app registration with specified device tag.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilder) Post(ctx context.Context, body ItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagPostRequestBodyable, requestConfiguration *ItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilderPostRequestConfiguration)(error) {
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
func (m *ItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagPostRequestBodyable, requestConfiguration *ItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilder when successful
func (m *ItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilder) WithUrl(rawUrl string)(*ItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return NewItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
