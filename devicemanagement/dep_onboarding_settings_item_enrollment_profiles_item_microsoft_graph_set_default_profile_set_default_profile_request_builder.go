package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DepOnboardingSettingsItemEnrollmentProfilesItemMicrosoftGraphSetDefaultProfileSetDefaultProfileRequestBuilder provides operations to call the setDefaultProfile method.
type DepOnboardingSettingsItemEnrollmentProfilesItemMicrosoftGraphSetDefaultProfileSetDefaultProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DepOnboardingSettingsItemEnrollmentProfilesItemMicrosoftGraphSetDefaultProfileSetDefaultProfileRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DepOnboardingSettingsItemEnrollmentProfilesItemMicrosoftGraphSetDefaultProfileSetDefaultProfileRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDepOnboardingSettingsItemEnrollmentProfilesItemMicrosoftGraphSetDefaultProfileSetDefaultProfileRequestBuilderInternal instantiates a new SetDefaultProfileRequestBuilder and sets the default values.
func NewDepOnboardingSettingsItemEnrollmentProfilesItemMicrosoftGraphSetDefaultProfileSetDefaultProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingsItemEnrollmentProfilesItemMicrosoftGraphSetDefaultProfileSetDefaultProfileRequestBuilder) {
    m := &DepOnboardingSettingsItemEnrollmentProfilesItemMicrosoftGraphSetDefaultProfileSetDefaultProfileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/enrollmentProfiles/{enrollmentProfile%2Did}/microsoft.graph.setDefaultProfile";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDepOnboardingSettingsItemEnrollmentProfilesItemMicrosoftGraphSetDefaultProfileSetDefaultProfileRequestBuilder instantiates a new SetDefaultProfileRequestBuilder and sets the default values.
func NewDepOnboardingSettingsItemEnrollmentProfilesItemMicrosoftGraphSetDefaultProfileSetDefaultProfileRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingsItemEnrollmentProfilesItemMicrosoftGraphSetDefaultProfileSetDefaultProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDepOnboardingSettingsItemEnrollmentProfilesItemMicrosoftGraphSetDefaultProfileSetDefaultProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action setDefaultProfile
func (m *DepOnboardingSettingsItemEnrollmentProfilesItemMicrosoftGraphSetDefaultProfileSetDefaultProfileRequestBuilder) Post(ctx context.Context, requestConfiguration *DepOnboardingSettingsItemEnrollmentProfilesItemMicrosoftGraphSetDefaultProfileSetDefaultProfileRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action setDefaultProfile
func (m *DepOnboardingSettingsItemEnrollmentProfilesItemMicrosoftGraphSetDefaultProfileSetDefaultProfileRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DepOnboardingSettingsItemEnrollmentProfilesItemMicrosoftGraphSetDefaultProfileSetDefaultProfileRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
