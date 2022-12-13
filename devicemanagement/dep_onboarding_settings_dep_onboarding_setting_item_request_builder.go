package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder provides operations to manage the depOnboardingSettings property of the microsoft.graph.deviceManagement entity.
type DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
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
// NewDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder instantiates a new DepOnboardingSettingItemRequestBuilder and sets the default values.
func NewDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property depOnboardingSettings for deviceManagement
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation this collections of multiple DEP tokens per-tenant.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property depOnboardingSettings in deviceManagement
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable, requestConfiguration *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DefaultIosEnrollmentProfile provides operations to manage the defaultIosEnrollmentProfile property of the microsoft.graph.depOnboardingSetting entity.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) DefaultIosEnrollmentProfile()(*DepOnboardingSettingsItemDefaultIosEnrollmentProfileRequestBuilder) {
    return NewDepOnboardingSettingsItemDefaultIosEnrollmentProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DefaultMacOsEnrollmentProfile provides operations to manage the defaultMacOsEnrollmentProfile property of the microsoft.graph.depOnboardingSetting entity.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) DefaultMacOsEnrollmentProfile()(*DepOnboardingSettingsItemDefaultMacOsEnrollmentProfileRequestBuilder) {
    return NewDepOnboardingSettingsItemDefaultMacOsEnrollmentProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property depOnboardingSettings for deviceManagement
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) EnrollmentProfiles()(*DepOnboardingSettingsItemEnrollmentProfilesRequestBuilder) {
    return NewDepOnboardingSettingsItemEnrollmentProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnrollmentProfilesById provides operations to manage the enrollmentProfiles property of the microsoft.graph.depOnboardingSetting entity.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) EnrollmentProfilesById(id string)(*DepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["enrollmentProfile%2Did"] = id
    }
    return NewDepOnboardingSettingsItemEnrollmentProfilesEnrollmentProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GenerateEncryptionPublicKey provides operations to call the generateEncryptionPublicKey method.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) GenerateEncryptionPublicKey()(*DepOnboardingSettingsItemGenerateEncryptionPublicKeyRequestBuilder) {
    return NewDepOnboardingSettingsItemGenerateEncryptionPublicKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get this collections of multiple DEP tokens per-tenant.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable, error) {
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
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) GetEncryptionPublicKey()(*DepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilder) {
    return NewDepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ImportedAppleDeviceIdentities provides operations to manage the importedAppleDeviceIdentities property of the microsoft.graph.depOnboardingSetting entity.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) ImportedAppleDeviceIdentities()(*DepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilder) {
    return NewDepOnboardingSettingsItemImportedAppleDeviceIdentitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ImportedAppleDeviceIdentitiesById provides operations to manage the importedAppleDeviceIdentities property of the microsoft.graph.depOnboardingSetting entity.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) ImportedAppleDeviceIdentitiesById(id string)(*DepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportedAppleDeviceIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["importedAppleDeviceIdentity%2Did"] = id
    }
    return NewDepOnboardingSettingsItemImportedAppleDeviceIdentitiesImportedAppleDeviceIdentityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property depOnboardingSettings in deviceManagement
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable, requestConfiguration *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable, error) {
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
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) ShareForSchoolDataSyncService()(*DepOnboardingSettingsItemShareForSchoolDataSyncServiceRequestBuilder) {
    return NewDepOnboardingSettingsItemShareForSchoolDataSyncServiceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SyncWithAppleDeviceEnrollmentProgram provides operations to call the syncWithAppleDeviceEnrollmentProgram method.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) SyncWithAppleDeviceEnrollmentProgram()(*DepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilder) {
    return NewDepOnboardingSettingsItemSyncWithAppleDeviceEnrollmentProgramRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnshareForSchoolDataSyncService provides operations to call the unshareForSchoolDataSyncService method.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) UnshareForSchoolDataSyncService()(*DepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilder) {
    return NewDepOnboardingSettingsItemUnshareForSchoolDataSyncServiceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UploadDepToken provides operations to call the uploadDepToken method.
func (m *DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) UploadDepToken()(*DepOnboardingSettingsItemUploadDepTokenRequestBuilder) {
    return NewDepOnboardingSettingsItemUploadDepTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
