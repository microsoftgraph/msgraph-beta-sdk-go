package manageddevice

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0204dc87398460d4287bddba221ba05169016444a01b33fa9e0115351af7c328 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/sendcustomnotificationtocompanyportal"
    i0caee8f1fcd1e35f3a745651c45eb9e041203cbff50e2330d6c6ba93246465ce "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/getoemwarranty"
    i0e79d1aff370896ba720fb518bf46a08ebe6591ef9f3a5e9383504e5db833878 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/createdevicelogcollectionrequest"
    i1514490d3ba4a568c72dd8130489159e5fcb3ba47bbfffb6c24e69f20541b627 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/ref"
    i24e276762d758f89618c5d08efe0d5a45f34e976c3dcefb7e9e5ebcb3541017f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/restorecloudpc"
    i27b8067f2cbef328e2ea92912dd0a95c10faacf26db22885f0f9194e45e67066 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/resetpasscode"
    i295788bf4a102f78f3c216f64170a22e460ba7e23a8041d6526640a26b0738f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/recoverpasscode"
    i2c446b375b2fd48b5eb5d4b370a970ffe6c57a0acdbfe0e28e8094475db38b82 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/getfilevaultkey"
    i2fdf5debc613c834402a1c5cb13c9b671471ef951677072fc939c53091966aac "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/windowsdefenderupdatesignatures"
    i3374e8447c02551ee2536075a785741a5e35ffe6b39d46bee586a40ebdde92fd "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/reenable"
    i34c9119fc415f067cfe3dc38398331b8c2eb9abb8b17597ee3f89316b1a05f2a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/playlostmodesound"
    i3649ef3901d2e347fcadf2f76814f9c8d60ec93688c1bbdc97444e13806de185 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/rotatebitlockerkeys"
    i4c1588ef782b785d9417fd290904a65cb77139848496db9afcc06e3814f42845 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/updatewindowsdeviceaccount"
    i4cad860cd6d907c8bf0feba31f0f39793c28667a7260d63fe06d0bf964dd9507 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/remotelock"
    i4d5b9ce4325ce655752fa49acfff7c227f74dc5772baddc9ab972d0099dd172c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/requestremoteassistance"
    i500b20bd7e2735a749bbc4006c139aa4bd48d59beb885ab46c11c96fb3c0fc2b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/deleteuserfromsharedappledevice"
    i517450f6e427b6ca5d73623bf73d81c28199cb35ccde33df9a5d88ae19ffa8cf "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/retire"
    i5797da546fe57ad4df31fe8a2b6f987573f3d6589723ad8662823d6c3770612e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/rebootnow"
    i5b4e1dd744c7393a2040ed414d84fe31f73999cf9023725687cbd7951cd22a93 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/deprovision"
    i5c328c2f039bfe4f71420dac168aac22c152ed72ee261f4f681036292ae71565 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/rotatefilevaultkey"
    i63a60c66200dc227be5ca2d5b63a465840df9eb4645ddcb19a2d1ea1a87389e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/activatedeviceesim"
    i710cdaa7adc294a42b9510684243ddf0d96fe328352f8be6599a5484966b9851 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/resizecloudpc"
    i8472707729f0a6774c2dd878ad69978368ef302ab88e38ce6c283ecf1da8f3a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/enablelostmode"
    i86a4110f1196e12494eb88d958b16fbbe5d9ea4da5db422cb4afa3fd12e095dd "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/disablelostmode"
    i894da0083e871f823adfb3e2b9c16dd9535f7c99e13e30184c686c7b6be9bffc "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/locatedevice"
    i90c924c51cccf9a7df3c8c92a8bb29b79e878b9b8aa5b97490ca2003a4eef1aa "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/cleanwindowsdevice"
    i98ffe717d2e5a6124f0902bee09dcd3cd721c8eb1b271550d6323b00e6881a93 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/overridecompliancestate"
    i9b0faeb5632998affc585543c944b6e0da86fbd5df8b7b3609cb93a92476cd9d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/revokeapplevpplicenses"
    i9e2c8c998231ca71595887e6a6762c5a478a80c31cd7c1d1e1cf0183be91e0da "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/disable"
    ia5abac056477bc9e56c3c443a224e7210e173907269c130743f2790e0c48bf56 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/getnoncompliantsettings"
    ia5d18fd3154dd7d32e5284edc2bb091979644691965d2f205095fb145d40b56e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/wipe"
    ia6d2866a6d4ed56c9cfe03f0687e9d0aed10c3867684d934f0c3c65550ca7914 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/reprovisioncloudpc"
    iaaa26e890bfc4f889515b7f7a5f6cfabdbeca35c6eeafee8800017cbe5fee901 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/syncdevice"
    iaeac50b20be1833b1d16f62dad05c25aac007e42bc47cacb2c470f67b39bdbbf "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/triggerconfigurationmanageraction"
    ib5eb03fe4e2262fc209b7f5bdeaae59cd44382255e7d0dda6ae31f759479a0ad "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/bypassactivationlock"
    ic9d1b67ed87b1196c8ff46b5a5223ed24d71f8b3ad69c2485a2d0bb87b6f6914 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/getcloudpcremoteactionresults"
    ica268eccd0ba8736b2d3e3fea1bf283e509165806066849ade9f77e4162d7dea "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/windowsdefenderscan"
    iccb4ba14219df49df245b73b4ec5b53a5671f483889848cf0a7ffa0cb923fe18 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/logoutsharedappledeviceactiveuser"
    idde2c45a4f52d7e2cb04fd8b1c3ffab68c73fad86a0386da005e7dfdf04bfd09 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/setdevicename"
    ie2655436ba44f850677e4b87f0f60605cc435a5c9710ebdc822de7d4f871ffcb "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item/manageddevice/shutdown"
)

// ManagedDeviceRequestBuilder builds and executes requests for operations under \deviceManagement\deviceManagementScripts\{deviceManagementScript-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice
type ManagedDeviceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagedDeviceRequestBuilderGetOptions options for Get
type ManagedDeviceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagedDeviceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedDeviceRequestBuilderGetQueryParameters the managed devices that executes the device management script.
type ManagedDeviceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
func (m *ManagedDeviceRequestBuilder) ActivateDeviceEsim()(*i63a60c66200dc227be5ca2d5b63a465840df9eb4645ddcb19a2d1ea1a87389e1.ActivateDeviceEsimRequestBuilder) {
    return i63a60c66200dc227be5ca2d5b63a465840df9eb4645ddcb19a2d1ea1a87389e1.NewActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) BypassActivationLock()(*ib5eb03fe4e2262fc209b7f5bdeaae59cd44382255e7d0dda6ae31f759479a0ad.BypassActivationLockRequestBuilder) {
    return ib5eb03fe4e2262fc209b7f5bdeaae59cd44382255e7d0dda6ae31f759479a0ad.NewBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) CleanWindowsDevice()(*i90c924c51cccf9a7df3c8c92a8bb29b79e878b9b8aa5b97490ca2003a4eef1aa.CleanWindowsDeviceRequestBuilder) {
    return i90c924c51cccf9a7df3c8c92a8bb29b79e878b9b8aa5b97490ca2003a4eef1aa.NewCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagedDeviceRequestBuilderInternal instantiates a new ManagedDeviceRequestBuilder and sets the default values.
func NewManagedDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceRequestBuilder) {
    m := &ManagedDeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceManagementScripts/{deviceManagementScript_id}/deviceRunStates/{deviceManagementScriptDeviceState_id}/managedDevice{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedDeviceRequestBuilder instantiates a new ManagedDeviceRequestBuilder and sets the default values.
func NewManagedDeviceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ManagedDeviceRequestBuilder) CreateDeviceLogCollectionRequest()(*i0e79d1aff370896ba720fb518bf46a08ebe6591ef9f3a5e9383504e5db833878.CreateDeviceLogCollectionRequestRequestBuilder) {
    return i0e79d1aff370896ba720fb518bf46a08ebe6591ef9f3a5e9383504e5db833878.NewCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the managed devices that executes the device management script.
func (m *ManagedDeviceRequestBuilder) CreateGetRequestInformation(options *ManagedDeviceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ManagedDeviceRequestBuilder) DeleteUserFromSharedAppleDevice()(*i500b20bd7e2735a749bbc4006c139aa4bd48d59beb885ab46c11c96fb3c0fc2b.DeleteUserFromSharedAppleDeviceRequestBuilder) {
    return i500b20bd7e2735a749bbc4006c139aa4bd48d59beb885ab46c11c96fb3c0fc2b.NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Deprovision()(*i5b4e1dd744c7393a2040ed414d84fe31f73999cf9023725687cbd7951cd22a93.DeprovisionRequestBuilder) {
    return i5b4e1dd744c7393a2040ed414d84fe31f73999cf9023725687cbd7951cd22a93.NewDeprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Disable()(*i9e2c8c998231ca71595887e6a6762c5a478a80c31cd7c1d1e1cf0183be91e0da.DisableRequestBuilder) {
    return i9e2c8c998231ca71595887e6a6762c5a478a80c31cd7c1d1e1cf0183be91e0da.NewDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DisableLostMode()(*i86a4110f1196e12494eb88d958b16fbbe5d9ea4da5db422cb4afa3fd12e095dd.DisableLostModeRequestBuilder) {
    return i86a4110f1196e12494eb88d958b16fbbe5d9ea4da5db422cb4afa3fd12e095dd.NewDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) EnableLostMode()(*i8472707729f0a6774c2dd878ad69978368ef302ab88e38ce6c283ecf1da8f3a1.EnableLostModeRequestBuilder) {
    return i8472707729f0a6774c2dd878ad69978368ef302ab88e38ce6c283ecf1da8f3a1.NewEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the managed devices that executes the device management script.
func (m *ManagedDeviceRequestBuilder) Get(options *ManagedDeviceRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDevice, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewManagedDevice() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDevice), nil
}
// GetCloudPcRemoteActionResults builds and executes requests for operations under \deviceManagement\deviceManagementScripts\{deviceManagementScript-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.getCloudPcRemoteActionResults()
func (m *ManagedDeviceRequestBuilder) GetCloudPcRemoteActionResults()(*ic9d1b67ed87b1196c8ff46b5a5223ed24d71f8b3ad69c2485a2d0bb87b6f6914.GetCloudPcRemoteActionResultsRequestBuilder) {
    return ic9d1b67ed87b1196c8ff46b5a5223ed24d71f8b3ad69c2485a2d0bb87b6f6914.NewGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFileVaultKey builds and executes requests for operations under \deviceManagement\deviceManagementScripts\{deviceManagementScript-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.getFileVaultKey()
func (m *ManagedDeviceRequestBuilder) GetFileVaultKey()(*i2c446b375b2fd48b5eb5d4b370a970ffe6c57a0acdbfe0e28e8094475db38b82.GetFileVaultKeyRequestBuilder) {
    return i2c446b375b2fd48b5eb5d4b370a970ffe6c57a0acdbfe0e28e8094475db38b82.NewGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetNonCompliantSettings builds and executes requests for operations under \deviceManagement\deviceManagementScripts\{deviceManagementScript-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.getNonCompliantSettings()
func (m *ManagedDeviceRequestBuilder) GetNonCompliantSettings()(*ia5abac056477bc9e56c3c443a224e7210e173907269c130743f2790e0c48bf56.GetNonCompliantSettingsRequestBuilder) {
    return ia5abac056477bc9e56c3c443a224e7210e173907269c130743f2790e0c48bf56.NewGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOemWarranty builds and executes requests for operations under \deviceManagement\deviceManagementScripts\{deviceManagementScript-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.getOemWarranty()
func (m *ManagedDeviceRequestBuilder) GetOemWarranty()(*i0caee8f1fcd1e35f3a745651c45eb9e041203cbff50e2330d6c6ba93246465ce.GetOemWarrantyRequestBuilder) {
    return i0caee8f1fcd1e35f3a745651c45eb9e041203cbff50e2330d6c6ba93246465ce.NewGetOemWarrantyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LocateDevice()(*i894da0083e871f823adfb3e2b9c16dd9535f7c99e13e30184c686c7b6be9bffc.LocateDeviceRequestBuilder) {
    return i894da0083e871f823adfb3e2b9c16dd9535f7c99e13e30184c686c7b6be9bffc.NewLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*iccb4ba14219df49df245b73b4ec5b53a5671f483889848cf0a7ffa0cb923fe18.LogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return iccb4ba14219df49df245b73b4ec5b53a5671f483889848cf0a7ffa0cb923fe18.NewLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) OverrideComplianceState()(*i98ffe717d2e5a6124f0902bee09dcd3cd721c8eb1b271550d6323b00e6881a93.OverrideComplianceStateRequestBuilder) {
    return i98ffe717d2e5a6124f0902bee09dcd3cd721c8eb1b271550d6323b00e6881a93.NewOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) PlayLostModeSound()(*i34c9119fc415f067cfe3dc38398331b8c2eb9abb8b17597ee3f89316b1a05f2a.PlayLostModeSoundRequestBuilder) {
    return i34c9119fc415f067cfe3dc38398331b8c2eb9abb8b17597ee3f89316b1a05f2a.NewPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RebootNow()(*i5797da546fe57ad4df31fe8a2b6f987573f3d6589723ad8662823d6c3770612e.RebootNowRequestBuilder) {
    return i5797da546fe57ad4df31fe8a2b6f987573f3d6589723ad8662823d6c3770612e.NewRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RecoverPasscode()(*i295788bf4a102f78f3c216f64170a22e460ba7e23a8041d6526640a26b0738f0.RecoverPasscodeRequestBuilder) {
    return i295788bf4a102f78f3c216f64170a22e460ba7e23a8041d6526640a26b0738f0.NewRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Reenable()(*i3374e8447c02551ee2536075a785741a5e35ffe6b39d46bee586a40ebdde92fd.ReenableRequestBuilder) {
    return i3374e8447c02551ee2536075a785741a5e35ffe6b39d46bee586a40ebdde92fd.NewReenableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Ref()(*i1514490d3ba4a568c72dd8130489159e5fcb3ba47bbfffb6c24e69f20541b627.RefRequestBuilder) {
    return i1514490d3ba4a568c72dd8130489159e5fcb3ba47bbfffb6c24e69f20541b627.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RemoteLock()(*i4cad860cd6d907c8bf0feba31f0f39793c28667a7260d63fe06d0bf964dd9507.RemoteLockRequestBuilder) {
    return i4cad860cd6d907c8bf0feba31f0f39793c28667a7260d63fe06d0bf964dd9507.NewRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ReprovisionCloudPc()(*ia6d2866a6d4ed56c9cfe03f0687e9d0aed10c3867684d934f0c3c65550ca7914.ReprovisionCloudPcRequestBuilder) {
    return ia6d2866a6d4ed56c9cfe03f0687e9d0aed10c3867684d934f0c3c65550ca7914.NewReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RequestRemoteAssistance()(*i4d5b9ce4325ce655752fa49acfff7c227f74dc5772baddc9ab972d0099dd172c.RequestRemoteAssistanceRequestBuilder) {
    return i4d5b9ce4325ce655752fa49acfff7c227f74dc5772baddc9ab972d0099dd172c.NewRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ResetPasscode()(*i27b8067f2cbef328e2ea92912dd0a95c10faacf26db22885f0f9194e45e67066.ResetPasscodeRequestBuilder) {
    return i27b8067f2cbef328e2ea92912dd0a95c10faacf26db22885f0f9194e45e67066.NewResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ResizeCloudPc()(*i710cdaa7adc294a42b9510684243ddf0d96fe328352f8be6599a5484966b9851.ResizeCloudPcRequestBuilder) {
    return i710cdaa7adc294a42b9510684243ddf0d96fe328352f8be6599a5484966b9851.NewResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RestoreCloudPc()(*i24e276762d758f89618c5d08efe0d5a45f34e976c3dcefb7e9e5ebcb3541017f.RestoreCloudPcRequestBuilder) {
    return i24e276762d758f89618c5d08efe0d5a45f34e976c3dcefb7e9e5ebcb3541017f.NewRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Retire()(*i517450f6e427b6ca5d73623bf73d81c28199cb35ccde33df9a5d88ae19ffa8cf.RetireRequestBuilder) {
    return i517450f6e427b6ca5d73623bf73d81c28199cb35ccde33df9a5d88ae19ffa8cf.NewRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RevokeAppleVppLicenses()(*i9b0faeb5632998affc585543c944b6e0da86fbd5df8b7b3609cb93a92476cd9d.RevokeAppleVppLicensesRequestBuilder) {
    return i9b0faeb5632998affc585543c944b6e0da86fbd5df8b7b3609cb93a92476cd9d.NewRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RotateBitLockerKeys()(*i3649ef3901d2e347fcadf2f76814f9c8d60ec93688c1bbdc97444e13806de185.RotateBitLockerKeysRequestBuilder) {
    return i3649ef3901d2e347fcadf2f76814f9c8d60ec93688c1bbdc97444e13806de185.NewRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RotateFileVaultKey()(*i5c328c2f039bfe4f71420dac168aac22c152ed72ee261f4f681036292ae71565.RotateFileVaultKeyRequestBuilder) {
    return i5c328c2f039bfe4f71420dac168aac22c152ed72ee261f4f681036292ae71565.NewRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SendCustomNotificationToCompanyPortal()(*i0204dc87398460d4287bddba221ba05169016444a01b33fa9e0115351af7c328.SendCustomNotificationToCompanyPortalRequestBuilder) {
    return i0204dc87398460d4287bddba221ba05169016444a01b33fa9e0115351af7c328.NewSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SetDeviceName()(*idde2c45a4f52d7e2cb04fd8b1c3ffab68c73fad86a0386da005e7dfdf04bfd09.SetDeviceNameRequestBuilder) {
    return idde2c45a4f52d7e2cb04fd8b1c3ffab68c73fad86a0386da005e7dfdf04bfd09.NewSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ShutDown()(*ie2655436ba44f850677e4b87f0f60605cc435a5c9710ebdc822de7d4f871ffcb.ShutDownRequestBuilder) {
    return ie2655436ba44f850677e4b87f0f60605cc435a5c9710ebdc822de7d4f871ffcb.NewShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SyncDevice()(*iaaa26e890bfc4f889515b7f7a5f6cfabdbeca35c6eeafee8800017cbe5fee901.SyncDeviceRequestBuilder) {
    return iaaa26e890bfc4f889515b7f7a5f6cfabdbeca35c6eeafee8800017cbe5fee901.NewSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) TriggerConfigurationManagerAction()(*iaeac50b20be1833b1d16f62dad05c25aac007e42bc47cacb2c470f67b39bdbbf.TriggerConfigurationManagerActionRequestBuilder) {
    return iaeac50b20be1833b1d16f62dad05c25aac007e42bc47cacb2c470f67b39bdbbf.NewTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) UpdateWindowsDeviceAccount()(*i4c1588ef782b785d9417fd290904a65cb77139848496db9afcc06e3814f42845.UpdateWindowsDeviceAccountRequestBuilder) {
    return i4c1588ef782b785d9417fd290904a65cb77139848496db9afcc06e3814f42845.NewUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderScan()(*ica268eccd0ba8736b2d3e3fea1bf283e509165806066849ade9f77e4162d7dea.WindowsDefenderScanRequestBuilder) {
    return ica268eccd0ba8736b2d3e3fea1bf283e509165806066849ade9f77e4162d7dea.NewWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderUpdateSignatures()(*i2fdf5debc613c834402a1c5cb13c9b671471ef951677072fc939c53091966aac.WindowsDefenderUpdateSignaturesRequestBuilder) {
    return i2fdf5debc613c834402a1c5cb13c9b671471ef951677072fc939c53091966aac.NewWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Wipe()(*ia5d18fd3154dd7d32e5284edc2bb091979644691965d2f205095fb145d40b56e.WipeRequestBuilder) {
    return ia5d18fd3154dd7d32e5284edc2bb091979644691965d2f205095fb145d40b56e.NewWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
