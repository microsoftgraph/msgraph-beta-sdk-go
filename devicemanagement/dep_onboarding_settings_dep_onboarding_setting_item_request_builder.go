package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder provides operations to manage the depOnboardingSettings property of the microsoft.graph.deviceManagement entity.
type DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DepOnboardingSettingsDepOnboardingSettingItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DepOnboardingSettingsDepOnboardingSettingItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DepOnboardingSettingsDepOnboardingSettingItemRequestBuilderGetQueryParameters this collections of multiple DEP tokens per-tenant.
type DepOnboardingSettingsDepOnboardingSettingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DepOnboardingSettingsDepOnboardingSettingItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DepOnboardingSettingsDepOnboardingSettingItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilderGetQueryParameters
}
// DepOnboardingSettingsDepOnboardingSettingItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DepOnboardingSettingsDepOnboardingSettingItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderInternal instantiates a new DepOnboardingSettingItemRequestBuilder and sets the default values.
func NewDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) {
    m := &DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder instantiates a new DepOnboardingSettingItemRequestBuilder and sets the default values.
func NewDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// DefaultIosEnrollmentProfile provides operations to manage the defaultIosEnrollmentProfile property of the microsoft.graph.depOnboardingSetting entity.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) DefaultIosEnrollmentProfile()(*DepOnboardingSettingsItemDefaultIosEnrollmentProfileRequestBuilder) {
    return NewDepOnboardingSettingsItemDefaultIosEnrollmentProfileRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DefaultMacOsEnrollmentProfile provides operations to manage the defaultMacOsEnrollmentProfile property of the microsoft.graph.depOnboardingSetting entity.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) DefaultMacOsEnrollmentProfile()(*DepOnboardingSettingsItemDefaultMacOsEnrollmentProfileRequestBuilder) {
    return NewDepOnboardingSettingsItemDefaultMacOsEnrollmentProfileRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property depOnboardingSettings for deviceManagement
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// EnrollmentProfiles provides operations to manage the enrollmentProfiles property of the microsoft.graph.depOnboardingSetting entity.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) EnrollmentProfiles()(*DepOnboardingSettingsItemEnrollmentProfilesRequestBuilder) {
    return NewDepOnboardingSettingsItemEnrollmentProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GenerateEncryptionPublicKey provides operations to call the generateEncryptionPublicKey method.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) GenerateEncryptionPublicKey()(*DepOnboardingSettingsItemGenerateEncryptionPublicKeyRequestBuilder) {
    return NewDepOnboardingSettingsItemGenerateEncryptionPublicKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get this collections of multiple DEP tokens per-tenant.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDepOnboardingSettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable), nil
}
// GetEncryptionPublicKey provides operations to call the getEncryptionPublicKey method.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) GetEncryptionPublicKey()(*DepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilder) {
    return NewDepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImportedAppleDeviceIdentities provides operations to manage the importedAppleDeviceIdentities property of the microsoft.graph.depOnboardingSetting entity.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) ImportedAppleDeviceIdentities()(*DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder) {
    return NewDepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property depOnboardingSettings in deviceManagement
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable, requestConfiguration *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDepOnboardingSettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable), nil
}
// ShareForSchoolDataSyncService provides operations to call the shareForSchoolDataSyncService method.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) ShareForSchoolDataSyncService()(*DepOnboardingSettingsItemShareForSchoolDataSyncServiceRequestBuilder) {
    return NewDepOnboardingSettingsItemShareForSchoolDataSyncServiceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SyncWithAppleDeviceEnrollmentProgram provides operations to call the syncWithAppleDeviceEnrollmentProgram method.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) SyncWithAppleDeviceEnrollmentProgram()(*DepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilder) {
    return NewDepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property depOnboardingSettings for deviceManagement
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation this collections of multiple DEP tokens per-tenant.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property depOnboardingSettings in deviceManagement
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable, requestConfiguration *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// UnshareForSchoolDataSyncService provides operations to call the unshareForSchoolDataSyncService method.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) UnshareForSchoolDataSyncService()(*DepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilder) {
    return NewDepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UploadDepToken provides operations to call the uploadDepToken method.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) UploadDepToken()(*DepOnboardingSettingsItemUploadDepTokenRequestBuilder) {
    return NewDepOnboardingSettingsItemUploadDepTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) WithUrl(rawUrl string)(*DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) {
    return NewDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
