package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsInformationProtectionDeviceRegistrationsItemWipeRequestBuilder provides operations to call the wipe method.
type WindowsInformationProtectionDeviceRegistrationsItemWipeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsInformationProtectionDeviceRegistrationsItemWipeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsInformationProtectionDeviceRegistrationsItemWipeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsInformationProtectionDeviceRegistrationsItemWipeRequestBuilderInternal instantiates a new WipeRequestBuilder and sets the default values.
func NewWindowsInformationProtectionDeviceRegistrationsItemWipeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsInformationProtectionDeviceRegistrationsItemWipeRequestBuilder) {
    m := &WindowsInformationProtectionDeviceRegistrationsItemWipeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/windowsInformationProtectionDeviceRegistrations/{windowsInformationProtectionDeviceRegistration%2Did}/wipe", pathParameters),
    }
    return m
}
// NewWindowsInformationProtectionDeviceRegistrationsItemWipeRequestBuilder instantiates a new WipeRequestBuilder and sets the default values.
func NewWindowsInformationProtectionDeviceRegistrationsItemWipeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsInformationProtectionDeviceRegistrationsItemWipeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsInformationProtectionDeviceRegistrationsItemWipeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action wipe
func (m *WindowsInformationProtectionDeviceRegistrationsItemWipeRequestBuilder) Post(ctx context.Context, requestConfiguration *WindowsInformationProtectionDeviceRegistrationsItemWipeRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation invoke action wipe
func (m *WindowsInformationProtectionDeviceRegistrationsItemWipeRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *WindowsInformationProtectionDeviceRegistrationsItemWipeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *WindowsInformationProtectionDeviceRegistrationsItemWipeRequestBuilder) WithUrl(rawUrl string)(*WindowsInformationProtectionDeviceRegistrationsItemWipeRequestBuilder) {
    return NewWindowsInformationProtectionDeviceRegistrationsItemWipeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
