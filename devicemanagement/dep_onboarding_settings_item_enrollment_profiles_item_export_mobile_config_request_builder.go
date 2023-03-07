package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilder provides operations to call the exportMobileConfig method.
type DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilderInternal instantiates a new ExportMobileConfigRequestBuilder and sets the default values.
func NewDepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilder) {
    m := &DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/enrollmentProfiles/{enrollmentProfile%2Did}/exportMobileConfig()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilder instantiates a new ExportMobileConfigRequestBuilder and sets the default values.
func NewDepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilderInternal(urlParams, requestAdapter)
}
// Get exports the mobile configuration
func (m *DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilder) Get(ctx context.Context, requestConfiguration *DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilderGetRequestConfiguration)(DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateDepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigResponseable), nil
}
// ToGetRequestInformation exports the mobile configuration
func (m *DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
