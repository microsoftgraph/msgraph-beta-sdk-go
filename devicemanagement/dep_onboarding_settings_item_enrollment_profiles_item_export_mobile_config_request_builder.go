package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilder provides operations to call the exportMobileConfig method.
type DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
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
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/enrollmentProfiles/{enrollmentProfile%2Did}/exportMobileConfig()", pathParameters),
    }
    return m
}
// NewDepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilder instantiates a new ExportMobileConfigRequestBuilder and sets the default values.
func NewDepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilderInternal(urlParams, requestAdapter)
}
// Get exports the mobile configuration
// Deprecated: This method is obsolete. Use GetAsExportMobileConfigGetResponse instead.
func (m *DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilder) Get(ctx context.Context, requestConfiguration *DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilderGetRequestConfiguration)(DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigResponseable), nil
}
// GetAsExportMobileConfigGetResponse exports the mobile configuration
func (m *DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilder) GetAsExportMobileConfigGetResponse(ctx context.Context, requestConfiguration *DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilderGetRequestConfiguration)(DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigGetResponseable), nil
}
// ToGetRequestInformation exports the mobile configuration
func (m *DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilder) WithUrl(rawUrl string)(*DepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilder) {
    return NewDepOnboardingSettingsItemEnrollmentProfilesItemExportMobileConfigRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
