package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsAutopilotDeviceIdentitiesItemAllowNextEnrollmentRequestBuilder provides operations to call the allowNextEnrollment method.
type WindowsAutopilotDeviceIdentitiesItemAllowNextEnrollmentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsAutopilotDeviceIdentitiesItemAllowNextEnrollmentRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsAutopilotDeviceIdentitiesItemAllowNextEnrollmentRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsAutopilotDeviceIdentitiesItemAllowNextEnrollmentRequestBuilderInternal instantiates a new AllowNextEnrollmentRequestBuilder and sets the default values.
func NewWindowsAutopilotDeviceIdentitiesItemAllowNextEnrollmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsAutopilotDeviceIdentitiesItemAllowNextEnrollmentRequestBuilder) {
    m := &WindowsAutopilotDeviceIdentitiesItemAllowNextEnrollmentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsAutopilotDeviceIdentities/{windowsAutopilotDeviceIdentity%2Did}/allowNextEnrollment", pathParameters),
    }
    return m
}
// NewWindowsAutopilotDeviceIdentitiesItemAllowNextEnrollmentRequestBuilder instantiates a new AllowNextEnrollmentRequestBuilder and sets the default values.
func NewWindowsAutopilotDeviceIdentitiesItemAllowNextEnrollmentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsAutopilotDeviceIdentitiesItemAllowNextEnrollmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsAutopilotDeviceIdentitiesItemAllowNextEnrollmentRequestBuilderInternal(urlParams, requestAdapter)
}
// Post unblocks next autopilot enrollment.
func (m *WindowsAutopilotDeviceIdentitiesItemAllowNextEnrollmentRequestBuilder) Post(ctx context.Context, requestConfiguration *WindowsAutopilotDeviceIdentitiesItemAllowNextEnrollmentRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation unblocks next autopilot enrollment.
func (m *WindowsAutopilotDeviceIdentitiesItemAllowNextEnrollmentRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *WindowsAutopilotDeviceIdentitiesItemAllowNextEnrollmentRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *WindowsAutopilotDeviceIdentitiesItemAllowNextEnrollmentRequestBuilder) WithUrl(rawUrl string)(*WindowsAutopilotDeviceIdentitiesItemAllowNextEnrollmentRequestBuilder) {
    return NewWindowsAutopilotDeviceIdentitiesItemAllowNextEnrollmentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
