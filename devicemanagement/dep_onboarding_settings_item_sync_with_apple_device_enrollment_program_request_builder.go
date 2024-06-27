package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilder provides operations to call the syncWithAppleDeviceEnrollmentProgram method.
type DepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilderInternal instantiates a new DepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilder and sets the default values.
func NewDepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilder) {
    m := &DepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/syncWithAppleDeviceEnrollmentProgram", pathParameters),
    }
    return m
}
// NewDepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilder instantiates a new DepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilder and sets the default values.
func NewDepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilderInternal(urlParams, requestAdapter)
}
// Post synchronizes between Apple Device Enrollment Program and Intune
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilder) Post(ctx context.Context, requestConfiguration *DepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation synchronizes between Apple Device Enrollment Program and Intune
// returns a *RequestInformation when successful
func (m *DepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilder when successful
func (m *DepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilder) WithUrl(rawUrl string)(*DepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilder) {
    return NewDepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
