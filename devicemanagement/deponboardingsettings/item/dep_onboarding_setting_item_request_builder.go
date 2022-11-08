package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0a6e758c70ffdd42835c37d486fd53bac47c3b3806c07e1d3709a2073bdf86d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deponboardingsettings/item/getencryptionpublickey"
    i176618c3c34968adb403c4d65ed7688ded2af2f955513bc66fab436b03d406aa "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deponboardingsettings/item/enrollmentprofiles"
    i21c452abb08aebef53cfef108f20776c8c80e1b61fade3bc7b23262c36f6790e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deponboardingsettings/item/importedappledeviceidentities"
    i8cbc14b32444f6bc2ca617287bc7edcc041b98aeb7e5731521363f58947dd8e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deponboardingsettings/item/defaultmacosenrollmentprofile"
    i989047b1c27e9cacd2dfc635896ba97eabdf1a1f2c9359c6782243264caa99ae "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deponboardingsettings/item/shareforschooldatasyncservice"
    id9d4a53b959ec383d97e9b95072f8672fabff3e6ae908bfa98f47839ef0cbb99 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deponboardingsettings/item/uploaddeptoken"
    idc472b110ec13510663e8f300d42a723cb0feabc33f9afe042fd0354e0c25f91 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deponboardingsettings/item/generateencryptionpublickey"
    ie23d6848bd98c8bd4afdc914cf1374fd1fd8f96b0c4e1dbfadefc18a867a668a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deponboardingsettings/item/unshareforschooldatasyncservice"
    ie8887b0ef7cf44f65a0cbcd3ff9572d708c15969aa877c22a6b342fba9ecacdd "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deponboardingsettings/item/defaultiosenrollmentprofile"
    if9fb73889c241fb97e6deb0813e9e46d242f7a381291cf01dad122ab56ea043c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deponboardingsettings/item/syncwithappledeviceenrollmentprogram"
    ia63cbab7405477c100c61aae973c313555d3f6ecfa0f50ec772d43ee157d7a9d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deponboardingsettings/item/enrollmentprofiles/item"
    ia6717b64f1af972b7242219bd4db44a401fd79458567d783585ee41f0fd2c0b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deponboardingsettings/item/importedappledeviceidentities/item"
)

// DepOnboardingSettingItemRequestBuilder provides operations to manage the depOnboardingSettings property of the microsoft.graph.deviceManagement entity.
type DepOnboardingSettingItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DepOnboardingSettingItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DepOnboardingSettingItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DepOnboardingSettingItemRequestBuilderGetQueryParameters this collections of multiple DEP tokens per-tenant.
type DepOnboardingSettingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DepOnboardingSettingItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DepOnboardingSettingItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DepOnboardingSettingItemRequestBuilderGetQueryParameters
}
// DepOnboardingSettingItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DepOnboardingSettingItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDepOnboardingSettingItemRequestBuilderInternal instantiates a new DepOnboardingSettingItemRequestBuilder and sets the default values.
func NewDepOnboardingSettingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingItemRequestBuilder) {
    m := &DepOnboardingSettingItemRequestBuilder{
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
// NewDepOnboardingSettingItemRequestBuilder instantiates a new DepOnboardingSettingItemRequestBuilder and sets the default values.
func NewDepOnboardingSettingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDepOnboardingSettingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property depOnboardingSettings for deviceManagement
func (m *DepOnboardingSettingItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DepOnboardingSettingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DepOnboardingSettingItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DepOnboardingSettingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DepOnboardingSettingItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable, requestConfiguration *DepOnboardingSettingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DepOnboardingSettingItemRequestBuilder) DefaultIosEnrollmentProfile()(*ie8887b0ef7cf44f65a0cbcd3ff9572d708c15969aa877c22a6b342fba9ecacdd.DefaultIosEnrollmentProfileRequestBuilder) {
    return ie8887b0ef7cf44f65a0cbcd3ff9572d708c15969aa877c22a6b342fba9ecacdd.NewDefaultIosEnrollmentProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DefaultMacOsEnrollmentProfile provides operations to manage the defaultMacOsEnrollmentProfile property of the microsoft.graph.depOnboardingSetting entity.
func (m *DepOnboardingSettingItemRequestBuilder) DefaultMacOsEnrollmentProfile()(*i8cbc14b32444f6bc2ca617287bc7edcc041b98aeb7e5731521363f58947dd8e5.DefaultMacOsEnrollmentProfileRequestBuilder) {
    return i8cbc14b32444f6bc2ca617287bc7edcc041b98aeb7e5731521363f58947dd8e5.NewDefaultMacOsEnrollmentProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property depOnboardingSettings for deviceManagement
func (m *DepOnboardingSettingItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DepOnboardingSettingItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *DepOnboardingSettingItemRequestBuilder) EnrollmentProfiles()(*i176618c3c34968adb403c4d65ed7688ded2af2f955513bc66fab436b03d406aa.EnrollmentProfilesRequestBuilder) {
    return i176618c3c34968adb403c4d65ed7688ded2af2f955513bc66fab436b03d406aa.NewEnrollmentProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnrollmentProfilesById provides operations to manage the enrollmentProfiles property of the microsoft.graph.depOnboardingSetting entity.
func (m *DepOnboardingSettingItemRequestBuilder) EnrollmentProfilesById(id string)(*ia63cbab7405477c100c61aae973c313555d3f6ecfa0f50ec772d43ee157d7a9d.EnrollmentProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["enrollmentProfile%2Did"] = id
    }
    return ia63cbab7405477c100c61aae973c313555d3f6ecfa0f50ec772d43ee157d7a9d.NewEnrollmentProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GenerateEncryptionPublicKey provides operations to call the generateEncryptionPublicKey method.
func (m *DepOnboardingSettingItemRequestBuilder) GenerateEncryptionPublicKey()(*idc472b110ec13510663e8f300d42a723cb0feabc33f9afe042fd0354e0c25f91.GenerateEncryptionPublicKeyRequestBuilder) {
    return idc472b110ec13510663e8f300d42a723cb0feabc33f9afe042fd0354e0c25f91.NewGenerateEncryptionPublicKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get this collections of multiple DEP tokens per-tenant.
func (m *DepOnboardingSettingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DepOnboardingSettingItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable, error) {
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
func (m *DepOnboardingSettingItemRequestBuilder) GetEncryptionPublicKey()(*i0a6e758c70ffdd42835c37d486fd53bac47c3b3806c07e1d3709a2073bdf86d8.GetEncryptionPublicKeyRequestBuilder) {
    return i0a6e758c70ffdd42835c37d486fd53bac47c3b3806c07e1d3709a2073bdf86d8.NewGetEncryptionPublicKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ImportedAppleDeviceIdentities provides operations to manage the importedAppleDeviceIdentities property of the microsoft.graph.depOnboardingSetting entity.
func (m *DepOnboardingSettingItemRequestBuilder) ImportedAppleDeviceIdentities()(*i21c452abb08aebef53cfef108f20776c8c80e1b61fade3bc7b23262c36f6790e.ImportedAppleDeviceIdentitiesRequestBuilder) {
    return i21c452abb08aebef53cfef108f20776c8c80e1b61fade3bc7b23262c36f6790e.NewImportedAppleDeviceIdentitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ImportedAppleDeviceIdentitiesById provides operations to manage the importedAppleDeviceIdentities property of the microsoft.graph.depOnboardingSetting entity.
func (m *DepOnboardingSettingItemRequestBuilder) ImportedAppleDeviceIdentitiesById(id string)(*ia6717b64f1af972b7242219bd4db44a401fd79458567d783585ee41f0fd2c0b6.ImportedAppleDeviceIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["importedAppleDeviceIdentity%2Did"] = id
    }
    return ia6717b64f1af972b7242219bd4db44a401fd79458567d783585ee41f0fd2c0b6.NewImportedAppleDeviceIdentityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property depOnboardingSettings in deviceManagement
func (m *DepOnboardingSettingItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable, requestConfiguration *DepOnboardingSettingItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DepOnboardingSettingable, error) {
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
func (m *DepOnboardingSettingItemRequestBuilder) ShareForSchoolDataSyncService()(*i989047b1c27e9cacd2dfc635896ba97eabdf1a1f2c9359c6782243264caa99ae.ShareForSchoolDataSyncServiceRequestBuilder) {
    return i989047b1c27e9cacd2dfc635896ba97eabdf1a1f2c9359c6782243264caa99ae.NewShareForSchoolDataSyncServiceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SyncWithAppleDeviceEnrollmentProgram provides operations to call the syncWithAppleDeviceEnrollmentProgram method.
func (m *DepOnboardingSettingItemRequestBuilder) SyncWithAppleDeviceEnrollmentProgram()(*if9fb73889c241fb97e6deb0813e9e46d242f7a381291cf01dad122ab56ea043c.SyncWithAppleDeviceEnrollmentProgramRequestBuilder) {
    return if9fb73889c241fb97e6deb0813e9e46d242f7a381291cf01dad122ab56ea043c.NewSyncWithAppleDeviceEnrollmentProgramRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnshareForSchoolDataSyncService provides operations to call the unshareForSchoolDataSyncService method.
func (m *DepOnboardingSettingItemRequestBuilder) UnshareForSchoolDataSyncService()(*ie23d6848bd98c8bd4afdc914cf1374fd1fd8f96b0c4e1dbfadefc18a867a668a.UnshareForSchoolDataSyncServiceRequestBuilder) {
    return ie23d6848bd98c8bd4afdc914cf1374fd1fd8f96b0c4e1dbfadefc18a867a668a.NewUnshareForSchoolDataSyncServiceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UploadDepToken provides operations to call the uploadDepToken method.
func (m *DepOnboardingSettingItemRequestBuilder) UploadDepToken()(*id9d4a53b959ec383d97e9b95072f8672fabff3e6ae908bfa98f47839ef0cbb99.UploadDepTokenRequestBuilder) {
    return id9d4a53b959ec383d97e9b95072f8672fabff3e6ae908bfa98f47839ef0cbb99.NewUploadDepTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
