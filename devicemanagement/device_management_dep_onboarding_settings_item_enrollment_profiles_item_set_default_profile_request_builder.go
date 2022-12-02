package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemSetDefaultProfileRequestBuilder provides operations to call the setDefaultProfile method.
type DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemSetDefaultProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemSetDefaultProfileRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemSetDefaultProfileRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemSetDefaultProfileRequestBuilderInternal instantiates a new SetDefaultProfileRequestBuilder and sets the default values.
func NewDeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemSetDefaultProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemSetDefaultProfileRequestBuilder) {
    m := &DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemSetDefaultProfileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/enrollmentProfiles/{enrollmentProfile%2Did}/microsoft.graph.setDefaultProfile";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemSetDefaultProfileRequestBuilder instantiates a new SetDefaultProfileRequestBuilder and sets the default values.
func NewDeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemSetDefaultProfileRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemSetDefaultProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemSetDefaultProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action setDefaultProfile
func (m *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemSetDefaultProfileRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemSetDefaultProfileRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action setDefaultProfile
func (m *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemSetDefaultProfileRequestBuilder) Post(ctx context.Context, requestConfiguration *DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesItemSetDefaultProfileRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
