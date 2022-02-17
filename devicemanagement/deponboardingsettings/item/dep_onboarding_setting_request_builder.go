package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
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

// DepOnboardingSettingRequestBuilder builds and executes requests for operations under \deviceManagement\depOnboardingSettings\{depOnboardingSetting-id}
type DepOnboardingSettingRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DepOnboardingSettingRequestBuilderDeleteOptions options for Delete
type DepOnboardingSettingRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DepOnboardingSettingRequestBuilderGetOptions options for Get
type DepOnboardingSettingRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DepOnboardingSettingRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DepOnboardingSettingRequestBuilderGetQueryParameters this collections of multiple DEP tokens per-tenant.
type DepOnboardingSettingRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DepOnboardingSettingRequestBuilderPatchOptions options for Patch
type DepOnboardingSettingRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DepOnboardingSetting;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDepOnboardingSettingRequestBuilderInternal instantiates a new DepOnboardingSettingRequestBuilder and sets the default values.
func NewDepOnboardingSettingRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DepOnboardingSettingRequestBuilder) {
    m := &DepOnboardingSettingRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDepOnboardingSettingRequestBuilder instantiates a new DepOnboardingSettingRequestBuilder and sets the default values.
func NewDepOnboardingSettingRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DepOnboardingSettingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDepOnboardingSettingRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation this collections of multiple DEP tokens per-tenant.
func (m *DepOnboardingSettingRequestBuilder) CreateDeleteRequestInformation(options *DepOnboardingSettingRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DepOnboardingSettingRequestBuilder) CreateGetRequestInformation(options *DepOnboardingSettingRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation this collections of multiple DEP tokens per-tenant.
func (m *DepOnboardingSettingRequestBuilder) CreatePatchRequestInformation(options *DepOnboardingSettingRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DepOnboardingSettingRequestBuilder) DefaultIosEnrollmentProfile()(*ie8887b0ef7cf44f65a0cbcd3ff9572d708c15969aa877c22a6b342fba9ecacdd.DefaultIosEnrollmentProfileRequestBuilder) {
    return ie8887b0ef7cf44f65a0cbcd3ff9572d708c15969aa877c22a6b342fba9ecacdd.NewDefaultIosEnrollmentProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DepOnboardingSettingRequestBuilder) DefaultMacOsEnrollmentProfile()(*i8cbc14b32444f6bc2ca617287bc7edcc041b98aeb7e5731521363f58947dd8e5.DefaultMacOsEnrollmentProfileRequestBuilder) {
    return i8cbc14b32444f6bc2ca617287bc7edcc041b98aeb7e5731521363f58947dd8e5.NewDefaultMacOsEnrollmentProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete this collections of multiple DEP tokens per-tenant.
func (m *DepOnboardingSettingRequestBuilder) Delete(options *DepOnboardingSettingRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *DepOnboardingSettingRequestBuilder) EnrollmentProfiles()(*i176618c3c34968adb403c4d65ed7688ded2af2f955513bc66fab436b03d406aa.EnrollmentProfilesRequestBuilder) {
    return i176618c3c34968adb403c4d65ed7688ded2af2f955513bc66fab436b03d406aa.NewEnrollmentProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnrollmentProfilesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.depOnboardingSettings.item.enrollmentProfiles.item collection
func (m *DepOnboardingSettingRequestBuilder) EnrollmentProfilesById(id string)(*ia63cbab7405477c100c61aae973c313555d3f6ecfa0f50ec772d43ee157d7a9d.EnrollmentProfileRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["enrollmentProfile_id"] = id
    }
    return ia63cbab7405477c100c61aae973c313555d3f6ecfa0f50ec772d43ee157d7a9d.NewEnrollmentProfileRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DepOnboardingSettingRequestBuilder) GenerateEncryptionPublicKey()(*idc472b110ec13510663e8f300d42a723cb0feabc33f9afe042fd0354e0c25f91.GenerateEncryptionPublicKeyRequestBuilder) {
    return idc472b110ec13510663e8f300d42a723cb0feabc33f9afe042fd0354e0c25f91.NewGenerateEncryptionPublicKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get this collections of multiple DEP tokens per-tenant.
func (m *DepOnboardingSettingRequestBuilder) Get(options *DepOnboardingSettingRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DepOnboardingSetting, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDepOnboardingSetting() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DepOnboardingSetting), nil
}
// GetEncryptionPublicKey builds and executes requests for operations under \deviceManagement\depOnboardingSettings\{depOnboardingSetting-id}\microsoft.graph.getEncryptionPublicKey()
func (m *DepOnboardingSettingRequestBuilder) GetEncryptionPublicKey()(*i0a6e758c70ffdd42835c37d486fd53bac47c3b3806c07e1d3709a2073bdf86d8.GetEncryptionPublicKeyRequestBuilder) {
    return i0a6e758c70ffdd42835c37d486fd53bac47c3b3806c07e1d3709a2073bdf86d8.NewGetEncryptionPublicKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DepOnboardingSettingRequestBuilder) ImportedAppleDeviceIdentities()(*i21c452abb08aebef53cfef108f20776c8c80e1b61fade3bc7b23262c36f6790e.ImportedAppleDeviceIdentitiesRequestBuilder) {
    return i21c452abb08aebef53cfef108f20776c8c80e1b61fade3bc7b23262c36f6790e.NewImportedAppleDeviceIdentitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ImportedAppleDeviceIdentitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.depOnboardingSettings.item.importedAppleDeviceIdentities.item collection
func (m *DepOnboardingSettingRequestBuilder) ImportedAppleDeviceIdentitiesById(id string)(*ia6717b64f1af972b7242219bd4db44a401fd79458567d783585ee41f0fd2c0b6.ImportedAppleDeviceIdentityRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["importedAppleDeviceIdentity_id"] = id
    }
    return ia6717b64f1af972b7242219bd4db44a401fd79458567d783585ee41f0fd2c0b6.NewImportedAppleDeviceIdentityRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch this collections of multiple DEP tokens per-tenant.
func (m *DepOnboardingSettingRequestBuilder) Patch(options *DepOnboardingSettingRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *DepOnboardingSettingRequestBuilder) ShareForSchoolDataSyncService()(*i989047b1c27e9cacd2dfc635896ba97eabdf1a1f2c9359c6782243264caa99ae.ShareForSchoolDataSyncServiceRequestBuilder) {
    return i989047b1c27e9cacd2dfc635896ba97eabdf1a1f2c9359c6782243264caa99ae.NewShareForSchoolDataSyncServiceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DepOnboardingSettingRequestBuilder) SyncWithAppleDeviceEnrollmentProgram()(*if9fb73889c241fb97e6deb0813e9e46d242f7a381291cf01dad122ab56ea043c.SyncWithAppleDeviceEnrollmentProgramRequestBuilder) {
    return if9fb73889c241fb97e6deb0813e9e46d242f7a381291cf01dad122ab56ea043c.NewSyncWithAppleDeviceEnrollmentProgramRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DepOnboardingSettingRequestBuilder) UnshareForSchoolDataSyncService()(*ie23d6848bd98c8bd4afdc914cf1374fd1fd8f96b0c4e1dbfadefc18a867a668a.UnshareForSchoolDataSyncServiceRequestBuilder) {
    return ie23d6848bd98c8bd4afdc914cf1374fd1fd8f96b0c4e1dbfadefc18a867a668a.NewUnshareForSchoolDataSyncServiceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DepOnboardingSettingRequestBuilder) UploadDepToken()(*id9d4a53b959ec383d97e9b95072f8672fabff3e6ae908bfa98f47839ef0cbb99.UploadDepTokenRequestBuilder) {
    return id9d4a53b959ec383d97e9b95072f8672fabff3e6ae908bfa98f47839ef0cbb99.NewUploadDepTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
