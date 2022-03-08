package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DepOnboardingSettingItemRequestBuilderDeleteOptions options for Delete
type DepOnboardingSettingItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DepOnboardingSettingItemRequestBuilderGetOptions options for Get
type DepOnboardingSettingItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DepOnboardingSettingItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DepOnboardingSettingItemRequestBuilderGetQueryParameters this collections of multiple DEP tokens per-tenant.
type DepOnboardingSettingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DepOnboardingSettingItemRequestBuilderPatchOptions options for Patch
type DepOnboardingSettingItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DepOnboardingSettingable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDepOnboardingSettingItemRequestBuilderInternal instantiates a new DepOnboardingSettingItemRequestBuilder and sets the default values.
func NewDepOnboardingSettingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DepOnboardingSettingItemRequestBuilder) {
    m := &DepOnboardingSettingItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDepOnboardingSettingItemRequestBuilder instantiates a new DepOnboardingSettingItemRequestBuilder and sets the default values.
func NewDepOnboardingSettingItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DepOnboardingSettingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDepOnboardingSettingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property depOnboardingSettings for deviceManagement
func (m *DepOnboardingSettingItemRequestBuilder) CreateDeleteRequestInformation(options *DepOnboardingSettingItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation this collections of multiple DEP tokens per-tenant.
func (m *DepOnboardingSettingItemRequestBuilder) CreateGetRequestInformation(options *DepOnboardingSettingItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property depOnboardingSettings in deviceManagement
func (m *DepOnboardingSettingItemRequestBuilder) CreatePatchRequestInformation(options *DepOnboardingSettingItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DepOnboardingSettingItemRequestBuilder) DefaultIosEnrollmentProfile()(*ie8887b0ef7cf44f65a0cbcd3ff9572d708c15969aa877c22a6b342fba9ecacdd.DefaultIosEnrollmentProfileRequestBuilder) {
    return ie8887b0ef7cf44f65a0cbcd3ff9572d708c15969aa877c22a6b342fba9ecacdd.NewDefaultIosEnrollmentProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DepOnboardingSettingItemRequestBuilder) DefaultMacOsEnrollmentProfile()(*i8cbc14b32444f6bc2ca617287bc7edcc041b98aeb7e5731521363f58947dd8e5.DefaultMacOsEnrollmentProfileRequestBuilder) {
    return i8cbc14b32444f6bc2ca617287bc7edcc041b98aeb7e5731521363f58947dd8e5.NewDefaultMacOsEnrollmentProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property depOnboardingSettings for deviceManagement
func (m *DepOnboardingSettingItemRequestBuilder) Delete(options *DepOnboardingSettingItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *DepOnboardingSettingItemRequestBuilder) EnrollmentProfiles()(*i176618c3c34968adb403c4d65ed7688ded2af2f955513bc66fab436b03d406aa.EnrollmentProfilesRequestBuilder) {
    return i176618c3c34968adb403c4d65ed7688ded2af2f955513bc66fab436b03d406aa.NewEnrollmentProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnrollmentProfilesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.depOnboardingSettings.item.enrollmentProfiles.item collection
func (m *DepOnboardingSettingItemRequestBuilder) EnrollmentProfilesById(id string)(*ia63cbab7405477c100c61aae973c313555d3f6ecfa0f50ec772d43ee157d7a9d.EnrollmentProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["enrollmentProfile_id"] = id
    }
    return ia63cbab7405477c100c61aae973c313555d3f6ecfa0f50ec772d43ee157d7a9d.NewEnrollmentProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DepOnboardingSettingItemRequestBuilder) GenerateEncryptionPublicKey()(*idc472b110ec13510663e8f300d42a723cb0feabc33f9afe042fd0354e0c25f91.GenerateEncryptionPublicKeyRequestBuilder) {
    return idc472b110ec13510663e8f300d42a723cb0feabc33f9afe042fd0354e0c25f91.NewGenerateEncryptionPublicKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get this collections of multiple DEP tokens per-tenant.
func (m *DepOnboardingSettingItemRequestBuilder) Get(options *DepOnboardingSettingItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DepOnboardingSettingable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDepOnboardingSettingFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DepOnboardingSettingable), nil
}
// GetEncryptionPublicKey provides operations to call the getEncryptionPublicKey method.
func (m *DepOnboardingSettingItemRequestBuilder) GetEncryptionPublicKey()(*i0a6e758c70ffdd42835c37d486fd53bac47c3b3806c07e1d3709a2073bdf86d8.GetEncryptionPublicKeyRequestBuilder) {
    return i0a6e758c70ffdd42835c37d486fd53bac47c3b3806c07e1d3709a2073bdf86d8.NewGetEncryptionPublicKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DepOnboardingSettingItemRequestBuilder) ImportedAppleDeviceIdentities()(*i21c452abb08aebef53cfef108f20776c8c80e1b61fade3bc7b23262c36f6790e.ImportedAppleDeviceIdentitiesRequestBuilder) {
    return i21c452abb08aebef53cfef108f20776c8c80e1b61fade3bc7b23262c36f6790e.NewImportedAppleDeviceIdentitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ImportedAppleDeviceIdentitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.depOnboardingSettings.item.importedAppleDeviceIdentities.item collection
func (m *DepOnboardingSettingItemRequestBuilder) ImportedAppleDeviceIdentitiesById(id string)(*ia6717b64f1af972b7242219bd4db44a401fd79458567d783585ee41f0fd2c0b6.ImportedAppleDeviceIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["importedAppleDeviceIdentity_id"] = id
    }
    return ia6717b64f1af972b7242219bd4db44a401fd79458567d783585ee41f0fd2c0b6.NewImportedAppleDeviceIdentityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property depOnboardingSettings in deviceManagement
func (m *DepOnboardingSettingItemRequestBuilder) Patch(options *DepOnboardingSettingItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *DepOnboardingSettingItemRequestBuilder) ShareForSchoolDataSyncService()(*i989047b1c27e9cacd2dfc635896ba97eabdf1a1f2c9359c6782243264caa99ae.ShareForSchoolDataSyncServiceRequestBuilder) {
    return i989047b1c27e9cacd2dfc635896ba97eabdf1a1f2c9359c6782243264caa99ae.NewShareForSchoolDataSyncServiceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DepOnboardingSettingItemRequestBuilder) SyncWithAppleDeviceEnrollmentProgram()(*if9fb73889c241fb97e6deb0813e9e46d242f7a381291cf01dad122ab56ea043c.SyncWithAppleDeviceEnrollmentProgramRequestBuilder) {
    return if9fb73889c241fb97e6deb0813e9e46d242f7a381291cf01dad122ab56ea043c.NewSyncWithAppleDeviceEnrollmentProgramRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DepOnboardingSettingItemRequestBuilder) UnshareForSchoolDataSyncService()(*ie23d6848bd98c8bd4afdc914cf1374fd1fd8f96b0c4e1dbfadefc18a867a668a.UnshareForSchoolDataSyncServiceRequestBuilder) {
    return ie23d6848bd98c8bd4afdc914cf1374fd1fd8f96b0c4e1dbfadefc18a867a668a.NewUnshareForSchoolDataSyncServiceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DepOnboardingSettingItemRequestBuilder) UploadDepToken()(*id9d4a53b959ec383d97e9b95072f8672fabff3e6ae908bfa98f47839ef0cbb99.UploadDepTokenRequestBuilder) {
    return id9d4a53b959ec383d97e9b95072f8672fabff3e6ae908bfa98f47839ef0cbb99.NewUploadDepTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
