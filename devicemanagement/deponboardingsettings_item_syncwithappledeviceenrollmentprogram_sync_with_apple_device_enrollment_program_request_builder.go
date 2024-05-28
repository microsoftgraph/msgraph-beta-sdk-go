package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilder provides operations to call the syncWithAppleDeviceEnrollmentProgram method.
type DeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilderInternal instantiates a new DeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilder and sets the default values.
func NewDeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilder) {
    m := &DeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/syncWithAppleDeviceEnrollmentProgram", pathParameters),
    }
    return m
}
// NewDeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilder instantiates a new DeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilder and sets the default values.
func NewDeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilderInternal(urlParams, requestAdapter)
}
// Post synchronizes between Apple Device Enrollment Program and Intune
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilder) Post(ctx context.Context, requestConfiguration *DeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilderPostRequestConfiguration)(error) {
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
func (m *DeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilder when successful
func (m *DeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilder) WithUrl(rawUrl string)(*DeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilder) {
    return NewDeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
