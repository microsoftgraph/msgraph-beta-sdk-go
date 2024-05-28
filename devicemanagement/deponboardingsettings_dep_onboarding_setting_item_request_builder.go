package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeponboardingsettingsDepOnboardingSettingItemRequestBuilder provides operations to manage the depOnboardingSettings property of the microsoft.graph.deviceManagement entity.
type DeponboardingsettingsDepOnboardingSettingItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeponboardingsettingsDepOnboardingSettingItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeponboardingsettingsDepOnboardingSettingItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeponboardingsettingsDepOnboardingSettingItemRequestBuilderGetQueryParameters this collections of multiple DEP tokens per-tenant.
type DeponboardingsettingsDepOnboardingSettingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeponboardingsettingsDepOnboardingSettingItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeponboardingsettingsDepOnboardingSettingItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeponboardingsettingsDepOnboardingSettingItemRequestBuilderGetQueryParameters
}
// DeponboardingsettingsDepOnboardingSettingItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeponboardingsettingsDepOnboardingSettingItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeponboardingsettingsDepOnboardingSettingItemRequestBuilderInternal instantiates a new DeponboardingsettingsDepOnboardingSettingItemRequestBuilder and sets the default values.
func NewDeponboardingsettingsDepOnboardingSettingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeponboardingsettingsDepOnboardingSettingItemRequestBuilder) {
    m := &DeponboardingsettingsDepOnboardingSettingItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeponboardingsettingsDepOnboardingSettingItemRequestBuilder instantiates a new DeponboardingsettingsDepOnboardingSettingItemRequestBuilder and sets the default values.
func NewDeponboardingsettingsDepOnboardingSettingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeponboardingsettingsDepOnboardingSettingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeponboardingsettingsDepOnboardingSettingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// DefaultIosEnrollmentProfile provides operations to manage the defaultIosEnrollmentProfile property of the microsoft.graph.depOnboardingSetting entity.
// returns a *DeponboardingsettingsItemDefaultiosenrollmentprofileDefaultIosEnrollmentProfileRequestBuilder when successful
func (m *DeponboardingsettingsDepOnboardingSettingItemRequestBuilder) DefaultIosEnrollmentProfile()(*DeponboardingsettingsItemDefaultiosenrollmentprofileDefaultIosEnrollmentProfileRequestBuilder) {
    return NewDeponboardingsettingsItemDefaultiosenrollmentprofileDefaultIosEnrollmentProfileRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DefaultMacOsEnrollmentProfile provides operations to manage the defaultMacOsEnrollmentProfile property of the microsoft.graph.depOnboardingSetting entity.
// returns a *DeponboardingsettingsItemDefaultmacosenrollmentprofileDefaultMacOsEnrollmentProfileRequestBuilder when successful
func (m *DeponboardingsettingsDepOnboardingSettingItemRequestBuilder) DefaultMacOsEnrollmentProfile()(*DeponboardingsettingsItemDefaultmacosenrollmentprofileDefaultMacOsEnrollmentProfileRequestBuilder) {
    return NewDeponboardingsettingsItemDefaultmacosenrollmentprofileDefaultMacOsEnrollmentProfileRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property depOnboardingSettings for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeponboardingsettingsDepOnboardingSettingItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeponboardingsettingsDepOnboardingSettingItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// EnrollmentProfiles provides operations to manage the enrollmentProfiles property of the microsoft.graph.depOnboardingSetting entity.
// returns a *DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfilesRequestBuilder when successful
func (m *DeponboardingsettingsDepOnboardingSettingItemRequestBuilder) EnrollmentProfiles()(*DeponboardingsettingsItemEnrollmentprofilesEnrollmentProfilesRequestBuilder) {
    return NewDeponboardingsettingsItemEnrollmentprofilesEnrollmentProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GenerateEncryptionPublicKey provides operations to call the generateEncryptionPublicKey method.
// returns a *DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilder when successful
func (m *DeponboardingsettingsDepOnboardingSettingItemRequestBuilder) GenerateEncryptionPublicKey()(*DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilder) {
    return NewDeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get this collections of multiple DEP tokens per-tenant.
// returns a DepOnboardingSettingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeponboardingsettingsDepOnboardingSettingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeponboardingsettingsDepOnboardingSettingItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a *DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilder when successful
func (m *DeponboardingsettingsDepOnboardingSettingItemRequestBuilder) GetEncryptionPublicKey()(*DeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilder) {
    return NewDeponboardingsettingsItemGetencryptionpublickeyGetEncryptionPublicKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImportedAppleDeviceIdentities provides operations to manage the importedAppleDeviceIdentities property of the microsoft.graph.depOnboardingSetting entity.
// returns a *DeponboardingsettingsItemImportedappledeviceidentitiesImportedAppleDeviceIdentitiesRequestBuilder when successful
func (m *DeponboardingsettingsDepOnboardingSettingItemRequestBuilder) ImportedAppleDeviceIdentities()(*DeponboardingsettingsItemImportedappledeviceidentitiesImportedAppleDeviceIdentitiesRequestBuilder) {
    return NewDeponboardingsettingsItemImportedappledeviceidentitiesImportedAppleDeviceIdentitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property depOnboardingSettings in deviceManagement
// returns a DepOnboardingSettingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeponboardingsettingsDepOnboardingSettingItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable, requestConfiguration *DeponboardingsettingsDepOnboardingSettingItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a *DeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilder when successful
func (m *DeponboardingsettingsDepOnboardingSettingItemRequestBuilder) ShareForSchoolDataSyncService()(*DeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilder) {
    return NewDeponboardingsettingsItemShareforschooldatasyncserviceShareForSchoolDataSyncServiceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SyncWithAppleDeviceEnrollmentProgram provides operations to call the syncWithAppleDeviceEnrollmentProgram method.
// returns a *DeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilder when successful
func (m *DeponboardingsettingsDepOnboardingSettingItemRequestBuilder) SyncWithAppleDeviceEnrollmentProgram()(*DeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilder) {
    return NewDeponboardingsettingsItemSyncwithappledeviceenrollmentprogramSyncWithAppleDeviceEnrollmentProgramRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property depOnboardingSettings for deviceManagement
// returns a *RequestInformation when successful
func (m *DeponboardingsettingsDepOnboardingSettingItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeponboardingsettingsDepOnboardingSettingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation this collections of multiple DEP tokens per-tenant.
// returns a *RequestInformation when successful
func (m *DeponboardingsettingsDepOnboardingSettingItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeponboardingsettingsDepOnboardingSettingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RequestInformation when successful
func (m *DeponboardingsettingsDepOnboardingSettingItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable, requestConfiguration *DeponboardingsettingsDepOnboardingSettingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilder when successful
func (m *DeponboardingsettingsDepOnboardingSettingItemRequestBuilder) UnshareForSchoolDataSyncService()(*DeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilder) {
    return NewDeponboardingsettingsItemUnshareforschooldatasyncserviceUnshareForSchoolDataSyncServiceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UploadDepToken provides operations to call the uploadDepToken method.
// returns a *DeponboardingsettingsItemUploaddeptokenUploadDepTokenRequestBuilder when successful
func (m *DeponboardingsettingsDepOnboardingSettingItemRequestBuilder) UploadDepToken()(*DeponboardingsettingsItemUploaddeptokenUploadDepTokenRequestBuilder) {
    return NewDeponboardingsettingsItemUploaddeptokenUploadDepTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeponboardingsettingsDepOnboardingSettingItemRequestBuilder when successful
func (m *DeponboardingsettingsDepOnboardingSettingItemRequestBuilder) WithUrl(rawUrl string)(*DeponboardingsettingsDepOnboardingSettingItemRequestBuilder) {
    return NewDeponboardingsettingsDepOnboardingSettingItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
