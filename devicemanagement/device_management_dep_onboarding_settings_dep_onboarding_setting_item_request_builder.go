package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder provides operations to manage the depOnboardingSettings property of the microsoft.graph.deviceManagement entity.
type DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderGetQueryParameters this collections of multiple DEP tokens per-tenant.
type DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderGetQueryParameters
}
// DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderInternal instantiates a new DepOnboardingSettingItemRequestBuilder and sets the default values.
func NewDeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) {
    m := &DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder instantiates a new DepOnboardingSettingItemRequestBuilder and sets the default values.
func NewDeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property depOnboardingSettings for deviceManagement
func (m *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation this collections of multiple DEP tokens per-tenant.
func (m *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property depOnboardingSettings in deviceManagement
func (m *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable, requestConfiguration *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DefaultIosEnrollmentProfile provides operations to manage the defaultIosEnrollmentProfile property of the microsoft.graph.depOnboardingSetting entity.
func (m *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) DefaultIosEnrollmentProfile()(*DeviceManagementDepOnboardingSettingsItemDefaultIosEnrollmentProfileRequestBuilder) {
    return NewDeviceManagementDepOnboardingSettingsItemDefaultIosEnrollmentProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DefaultMacOsEnrollmentProfile provides operations to manage the defaultMacOsEnrollmentProfile property of the microsoft.graph.depOnboardingSetting entity.
func (m *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) DefaultMacOsEnrollmentProfile()(*DeviceManagementDepOnboardingSettingsItemDefaultMacOsEnrollmentProfileRequestBuilder) {
    return NewDeviceManagementDepOnboardingSettingsItemDefaultMacOsEnrollmentProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property depOnboardingSettings for deviceManagement
func (m *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// EnrollmentProfiles provides operations to manage the enrollmentProfiles property of the microsoft.graph.depOnboardingSetting entity.
func (m *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) EnrollmentProfiles()(*DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesRequestBuilder) {
    return NewDeviceManagementDepOnboardingSettingsItemEnrollmentProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnrollmentProfilesById provides operations to manage the enrollmentProfiles property of the microsoft.graph.depOnboardingSetting entity.
func (m *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) EnrollmentProfilesById(id string)(*DeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["enrollmentProfile%2Did"] = id
    }
    return NewDeviceManagementDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GenerateEncryptionPublicKey provides operations to call the generateEncryptionPublicKey method.
func (m *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) GenerateEncryptionPublicKey()(*DeviceManagementDepOnboardingSettingsItemGenerateEncryptionPublicKeyRequestBuilder) {
    return NewDeviceManagementDepOnboardingSettingsItemGenerateEncryptionPublicKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get this collections of multiple DEP tokens per-tenant.
func (m *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDepOnboardingSettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable), nil
}
// GetEncryptionPublicKey provides operations to call the getEncryptionPublicKey method.
func (m *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) GetEncryptionPublicKey()(*DeviceManagementDepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilder) {
    return NewDeviceManagementDepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ImportedAppleDeviceIdentities provides operations to manage the importedAppleDeviceIdentities property of the microsoft.graph.depOnboardingSetting entity.
func (m *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) ImportedAppleDeviceIdentities()(*DeviceManagementDepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder) {
    return NewDeviceManagementDepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ImportedAppleDeviceIdentitiesById provides operations to manage the importedAppleDeviceIdentities property of the microsoft.graph.depOnboardingSetting entity.
func (m *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) ImportedAppleDeviceIdentitiesById(id string)(*DeviceManagementDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportedAppleDeviceIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["importedAppleDeviceIdentity%2Did"] = id
    }
    return NewDeviceManagementDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportedAppleDeviceIdentityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property depOnboardingSettings in deviceManagement
func (m *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable, requestConfiguration *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDepOnboardingSettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable), nil
}
// ShareForSchoolDataSyncService provides operations to call the shareForSchoolDataSyncService method.
func (m *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) ShareForSchoolDataSyncService()(*DeviceManagementDepOnboardingSettingsItemShareForSchoolDataSyncServiceRequestBuilder) {
    return NewDeviceManagementDepOnboardingSettingsItemShareForSchoolDataSyncServiceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SyncWithAppleDeviceEnrollmentProgram provides operations to call the syncWithAppleDeviceEnrollmentProgram method.
func (m *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) SyncWithAppleDeviceEnrollmentProgram()(*DeviceManagementDepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilder) {
    return NewDeviceManagementDepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnshareForSchoolDataSyncService provides operations to call the unshareForSchoolDataSyncService method.
func (m *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) UnshareForSchoolDataSyncService()(*DeviceManagementDepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilder) {
    return NewDeviceManagementDepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UploadDepToken provides operations to call the uploadDepToken method.
func (m *DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) UploadDepToken()(*DeviceManagementDepOnboardingSettingsItemUploadDepTokenRequestBuilder) {
    return NewDeviceManagementDepOnboardingSettingsItemUploadDepTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
