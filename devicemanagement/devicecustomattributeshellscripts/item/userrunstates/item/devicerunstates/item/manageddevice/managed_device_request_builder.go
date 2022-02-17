package manageddevice

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0113d2cc3aa964427a0c9379f43c318cc0e551f02f3a72b4025f274b01fed38a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/recoverpasscode"
    i01b93502d4de25ec824f88e0b70a439c36adb94f17c7578aad64be89b5f79918 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/setdevicename"
    i09be9fe34d96d07cb6e65e0d2a0d2f8cf9308afacf6a2769bfd8c613ea22d353 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/resetpasscode"
    i0a3dcc91a1d415c87312ed76158a0e4d7d8d01deb115304dc3ce0d3ec3853620 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/disable"
    i1531b77f7bc6a368df7d044cbfdfb3378c63960e09011ec04c59279090e99c2b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/rebootnow"
    i16679e82632a0af1970adc5f1ebaf464e7f89484a6bebefc05f75a7e9d5a6755 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/revokeapplevpplicenses"
    i205c1cfc860e83160600f12ef33c84e6ed923ab733411d7b226224570eac89a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/enablelostmode"
    i209e2b2083ba318400dcea6a198367a0ae2caa69fb8460a19cd2b9ed653d15fc "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/wipe"
    i27d9afdb18763182e192baf05c28ce45899c4d65431970d0a75deee72b4d2ffa "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/bypassactivationlock"
    i2b107af841dcd31d3787426bcfede2c74e8e60c00ae771694d98d0149871948f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/windowsdefenderupdatesignatures"
    i33cca811806b97b76f83d558cdfede5c6c359d0477fc2c72135fda72cb309abc "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/cleanwindowsdevice"
    i48ba070dfb8f7572a8e826d347a53938137f29189195b1a8484c7f15abdf8c99 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/getnoncompliantsettings"
    i4933e01d55a93120599dff251f04036f1824b4bf0ca7f144878f51ed7f3e8dc5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/getcloudpcremoteactionresults"
    i4f07a523637d00bd8e338db10e0f5d86e8b6c55ec27a721cd84cd58d43bdffd5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/remotelock"
    i55fae017c747062e34edf59c747820d3c475fc00d0605b4cee70ff8c8c5b86be "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/syncdevice"
    i68308209a28acab8084c101176f21b048cc94ec1045d5572e8543afad487fa2a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/ref"
    i6d4fbf7b6134c31ecf346531b93b45886bc72b7a2af5160bc799ac76ad57f37d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/deleteuserfromsharedappledevice"
    i6f5de1549895878de97c34bfe3ac5c81ddd8a01e59d0b8a144636970af8315c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/resizecloudpc"
    i7a2c8a05481b2e4fae31d5574b42bcbebd8b0313ac4b146bf7b10f4145ecea87 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/overridecompliancestate"
    i9cb44657461195d581585e0e93cfa9f0c3e14feda164bac4b1cb0674607a3f0b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/createdevicelogcollectionrequest"
    ia6c73d13d750ef033236a81ca18d1def68002d39a8eb3ddd73542c7df36a3d5c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/sendcustomnotificationtocompanyportal"
    ia7495b826275c2eca57401876526f1db50af5e09b8c826c532de1b3e7221d797 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/rotatebitlockerkeys"
    iade03315799c4d8b07f7c30f9d1ea91a993acada1b486dc95845dd80c336ea85 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/reenable"
    iaf44a1cab77523322bd155de12aa075474471071524a8fb9c17f63b6148056e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/updatewindowsdeviceaccount"
    ib7da0b5bf7390ffce3e75771c76951a12b144b905d78c2a9eb711857d643a51e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/playlostmodesound"
    ib9b0dd6c653a1c3bb67182ac1b426664c84a225760b81a3f63a9a0d5ef2b7afa "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/locatedevice"
    ibac66010f8fbcb24085780c279c9d76fe576318ea0716249c52014ff8a72c2b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/getoemwarranty"
    ibd3d6ea6bfb0e2c933bac96c03b78652a9823ef84f135036962ef11eda4a74c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/windowsdefenderscan"
    ic30313d5d3476a57710091d621f36d146d2295984e0b30bb107154f92d084885 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/logoutsharedappledeviceactiveuser"
    ic5698227b1db7f9abcdfadb8c5540c8985757efeb27219f9dbf60166677ba974 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/reprovisioncloudpc"
    ic8ba7cb6b98f0ca735a5f8fcf2e365121f1c6ef28c8c634c415454efcb1c7d03 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/shutdown"
    icd6f0d318f67de28a6579d1f7ca3be4f8d4d4af0409e39bd4b240e13b1c37249 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/rotatefilevaultkey"
    idb628250968416c6e069d568084d71fc3f20848662bcd54fb56074002904ad09 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/disablelostmode"
    idca09cc99e497e8dbed398de1cc01d7f761e347b09b4dc1ec14a5e8e081ebe2e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/retire"
    ie6ba3f71553a104c36f4a18cddbd6c1872f85ddaccddd9dba1cb64181fce1795 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/getfilevaultkey"
    ie8fd8384b64acb2a6cf58fa67b11e6e06bf6fd0ecbc4f5cdcb8ece85665ff532 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/deprovision"
    iea8a25234cff13df4e31feaa9c2e96c0c4567ec51ec17f270a86d1bde5cd4d01 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/restorecloudpc"
    if4f72a4702988e258e58788447d0d45da69d316b63ff76a480b565e3191b185c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/activatedeviceesim"
    ifa10b07246d3d7ff899d4579c5a6a6beeb408d6d3601a833dd2c7dd4147a4118 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/requestremoteassistance"
    ifaa214a686668e2595da116c49578208861fcd6f75a1ac2a99a03caeba4d0133 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item/devicerunstates/item/manageddevice/triggerconfigurationmanageraction"
)

// ManagedDeviceRequestBuilder builds and executes requests for operations under \deviceManagement\deviceCustomAttributeShellScripts\{deviceCustomAttributeShellScript-id}\userRunStates\{deviceManagementScriptUserState-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice
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
func (m *ManagedDeviceRequestBuilder) ActivateDeviceEsim()(*if4f72a4702988e258e58788447d0d45da69d316b63ff76a480b565e3191b185c.ActivateDeviceEsimRequestBuilder) {
    return if4f72a4702988e258e58788447d0d45da69d316b63ff76a480b565e3191b185c.NewActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) BypassActivationLock()(*i27d9afdb18763182e192baf05c28ce45899c4d65431970d0a75deee72b4d2ffa.BypassActivationLockRequestBuilder) {
    return i27d9afdb18763182e192baf05c28ce45899c4d65431970d0a75deee72b4d2ffa.NewBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) CleanWindowsDevice()(*i33cca811806b97b76f83d558cdfede5c6c359d0477fc2c72135fda72cb309abc.CleanWindowsDeviceRequestBuilder) {
    return i33cca811806b97b76f83d558cdfede5c6c359d0477fc2c72135fda72cb309abc.NewCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagedDeviceRequestBuilderInternal instantiates a new ManagedDeviceRequestBuilder and sets the default values.
func NewManagedDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceRequestBuilder) {
    m := &ManagedDeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCustomAttributeShellScripts/{deviceCustomAttributeShellScript_id}/userRunStates/{deviceManagementScriptUserState_id}/deviceRunStates/{deviceManagementScriptDeviceState_id}/managedDevice{?select,expand}";
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
func (m *ManagedDeviceRequestBuilder) CreateDeviceLogCollectionRequest()(*i9cb44657461195d581585e0e93cfa9f0c3e14feda164bac4b1cb0674607a3f0b.CreateDeviceLogCollectionRequestRequestBuilder) {
    return i9cb44657461195d581585e0e93cfa9f0c3e14feda164bac4b1cb0674607a3f0b.NewCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ManagedDeviceRequestBuilder) DeleteUserFromSharedAppleDevice()(*i6d4fbf7b6134c31ecf346531b93b45886bc72b7a2af5160bc799ac76ad57f37d.DeleteUserFromSharedAppleDeviceRequestBuilder) {
    return i6d4fbf7b6134c31ecf346531b93b45886bc72b7a2af5160bc799ac76ad57f37d.NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Deprovision()(*ie8fd8384b64acb2a6cf58fa67b11e6e06bf6fd0ecbc4f5cdcb8ece85665ff532.DeprovisionRequestBuilder) {
    return ie8fd8384b64acb2a6cf58fa67b11e6e06bf6fd0ecbc4f5cdcb8ece85665ff532.NewDeprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Disable()(*i0a3dcc91a1d415c87312ed76158a0e4d7d8d01deb115304dc3ce0d3ec3853620.DisableRequestBuilder) {
    return i0a3dcc91a1d415c87312ed76158a0e4d7d8d01deb115304dc3ce0d3ec3853620.NewDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DisableLostMode()(*idb628250968416c6e069d568084d71fc3f20848662bcd54fb56074002904ad09.DisableLostModeRequestBuilder) {
    return idb628250968416c6e069d568084d71fc3f20848662bcd54fb56074002904ad09.NewDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) EnableLostMode()(*i205c1cfc860e83160600f12ef33c84e6ed923ab733411d7b226224570eac89a7.EnableLostModeRequestBuilder) {
    return i205c1cfc860e83160600f12ef33c84e6ed923ab733411d7b226224570eac89a7.NewEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// GetCloudPcRemoteActionResults builds and executes requests for operations under \deviceManagement\deviceCustomAttributeShellScripts\{deviceCustomAttributeShellScript-id}\userRunStates\{deviceManagementScriptUserState-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.getCloudPcRemoteActionResults()
func (m *ManagedDeviceRequestBuilder) GetCloudPcRemoteActionResults()(*i4933e01d55a93120599dff251f04036f1824b4bf0ca7f144878f51ed7f3e8dc5.GetCloudPcRemoteActionResultsRequestBuilder) {
    return i4933e01d55a93120599dff251f04036f1824b4bf0ca7f144878f51ed7f3e8dc5.NewGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFileVaultKey builds and executes requests for operations under \deviceManagement\deviceCustomAttributeShellScripts\{deviceCustomAttributeShellScript-id}\userRunStates\{deviceManagementScriptUserState-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.getFileVaultKey()
func (m *ManagedDeviceRequestBuilder) GetFileVaultKey()(*ie6ba3f71553a104c36f4a18cddbd6c1872f85ddaccddd9dba1cb64181fce1795.GetFileVaultKeyRequestBuilder) {
    return ie6ba3f71553a104c36f4a18cddbd6c1872f85ddaccddd9dba1cb64181fce1795.NewGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetNonCompliantSettings builds and executes requests for operations under \deviceManagement\deviceCustomAttributeShellScripts\{deviceCustomAttributeShellScript-id}\userRunStates\{deviceManagementScriptUserState-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.getNonCompliantSettings()
func (m *ManagedDeviceRequestBuilder) GetNonCompliantSettings()(*i48ba070dfb8f7572a8e826d347a53938137f29189195b1a8484c7f15abdf8c99.GetNonCompliantSettingsRequestBuilder) {
    return i48ba070dfb8f7572a8e826d347a53938137f29189195b1a8484c7f15abdf8c99.NewGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOemWarranty builds and executes requests for operations under \deviceManagement\deviceCustomAttributeShellScripts\{deviceCustomAttributeShellScript-id}\userRunStates\{deviceManagementScriptUserState-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.getOemWarranty()
func (m *ManagedDeviceRequestBuilder) GetOemWarranty()(*ibac66010f8fbcb24085780c279c9d76fe576318ea0716249c52014ff8a72c2b1.GetOemWarrantyRequestBuilder) {
    return ibac66010f8fbcb24085780c279c9d76fe576318ea0716249c52014ff8a72c2b1.NewGetOemWarrantyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LocateDevice()(*ib9b0dd6c653a1c3bb67182ac1b426664c84a225760b81a3f63a9a0d5ef2b7afa.LocateDeviceRequestBuilder) {
    return ib9b0dd6c653a1c3bb67182ac1b426664c84a225760b81a3f63a9a0d5ef2b7afa.NewLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*ic30313d5d3476a57710091d621f36d146d2295984e0b30bb107154f92d084885.LogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return ic30313d5d3476a57710091d621f36d146d2295984e0b30bb107154f92d084885.NewLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) OverrideComplianceState()(*i7a2c8a05481b2e4fae31d5574b42bcbebd8b0313ac4b146bf7b10f4145ecea87.OverrideComplianceStateRequestBuilder) {
    return i7a2c8a05481b2e4fae31d5574b42bcbebd8b0313ac4b146bf7b10f4145ecea87.NewOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) PlayLostModeSound()(*ib7da0b5bf7390ffce3e75771c76951a12b144b905d78c2a9eb711857d643a51e.PlayLostModeSoundRequestBuilder) {
    return ib7da0b5bf7390ffce3e75771c76951a12b144b905d78c2a9eb711857d643a51e.NewPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RebootNow()(*i1531b77f7bc6a368df7d044cbfdfb3378c63960e09011ec04c59279090e99c2b.RebootNowRequestBuilder) {
    return i1531b77f7bc6a368df7d044cbfdfb3378c63960e09011ec04c59279090e99c2b.NewRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RecoverPasscode()(*i0113d2cc3aa964427a0c9379f43c318cc0e551f02f3a72b4025f274b01fed38a.RecoverPasscodeRequestBuilder) {
    return i0113d2cc3aa964427a0c9379f43c318cc0e551f02f3a72b4025f274b01fed38a.NewRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Reenable()(*iade03315799c4d8b07f7c30f9d1ea91a993acada1b486dc95845dd80c336ea85.ReenableRequestBuilder) {
    return iade03315799c4d8b07f7c30f9d1ea91a993acada1b486dc95845dd80c336ea85.NewReenableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Ref()(*i68308209a28acab8084c101176f21b048cc94ec1045d5572e8543afad487fa2a.RefRequestBuilder) {
    return i68308209a28acab8084c101176f21b048cc94ec1045d5572e8543afad487fa2a.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RemoteLock()(*i4f07a523637d00bd8e338db10e0f5d86e8b6c55ec27a721cd84cd58d43bdffd5.RemoteLockRequestBuilder) {
    return i4f07a523637d00bd8e338db10e0f5d86e8b6c55ec27a721cd84cd58d43bdffd5.NewRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ReprovisionCloudPc()(*ic5698227b1db7f9abcdfadb8c5540c8985757efeb27219f9dbf60166677ba974.ReprovisionCloudPcRequestBuilder) {
    return ic5698227b1db7f9abcdfadb8c5540c8985757efeb27219f9dbf60166677ba974.NewReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RequestRemoteAssistance()(*ifa10b07246d3d7ff899d4579c5a6a6beeb408d6d3601a833dd2c7dd4147a4118.RequestRemoteAssistanceRequestBuilder) {
    return ifa10b07246d3d7ff899d4579c5a6a6beeb408d6d3601a833dd2c7dd4147a4118.NewRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ResetPasscode()(*i09be9fe34d96d07cb6e65e0d2a0d2f8cf9308afacf6a2769bfd8c613ea22d353.ResetPasscodeRequestBuilder) {
    return i09be9fe34d96d07cb6e65e0d2a0d2f8cf9308afacf6a2769bfd8c613ea22d353.NewResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ResizeCloudPc()(*i6f5de1549895878de97c34bfe3ac5c81ddd8a01e59d0b8a144636970af8315c8.ResizeCloudPcRequestBuilder) {
    return i6f5de1549895878de97c34bfe3ac5c81ddd8a01e59d0b8a144636970af8315c8.NewResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RestoreCloudPc()(*iea8a25234cff13df4e31feaa9c2e96c0c4567ec51ec17f270a86d1bde5cd4d01.RestoreCloudPcRequestBuilder) {
    return iea8a25234cff13df4e31feaa9c2e96c0c4567ec51ec17f270a86d1bde5cd4d01.NewRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Retire()(*idca09cc99e497e8dbed398de1cc01d7f761e347b09b4dc1ec14a5e8e081ebe2e.RetireRequestBuilder) {
    return idca09cc99e497e8dbed398de1cc01d7f761e347b09b4dc1ec14a5e8e081ebe2e.NewRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RevokeAppleVppLicenses()(*i16679e82632a0af1970adc5f1ebaf464e7f89484a6bebefc05f75a7e9d5a6755.RevokeAppleVppLicensesRequestBuilder) {
    return i16679e82632a0af1970adc5f1ebaf464e7f89484a6bebefc05f75a7e9d5a6755.NewRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RotateBitLockerKeys()(*ia7495b826275c2eca57401876526f1db50af5e09b8c826c532de1b3e7221d797.RotateBitLockerKeysRequestBuilder) {
    return ia7495b826275c2eca57401876526f1db50af5e09b8c826c532de1b3e7221d797.NewRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RotateFileVaultKey()(*icd6f0d318f67de28a6579d1f7ca3be4f8d4d4af0409e39bd4b240e13b1c37249.RotateFileVaultKeyRequestBuilder) {
    return icd6f0d318f67de28a6579d1f7ca3be4f8d4d4af0409e39bd4b240e13b1c37249.NewRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SendCustomNotificationToCompanyPortal()(*ia6c73d13d750ef033236a81ca18d1def68002d39a8eb3ddd73542c7df36a3d5c.SendCustomNotificationToCompanyPortalRequestBuilder) {
    return ia6c73d13d750ef033236a81ca18d1def68002d39a8eb3ddd73542c7df36a3d5c.NewSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SetDeviceName()(*i01b93502d4de25ec824f88e0b70a439c36adb94f17c7578aad64be89b5f79918.SetDeviceNameRequestBuilder) {
    return i01b93502d4de25ec824f88e0b70a439c36adb94f17c7578aad64be89b5f79918.NewSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ShutDown()(*ic8ba7cb6b98f0ca735a5f8fcf2e365121f1c6ef28c8c634c415454efcb1c7d03.ShutDownRequestBuilder) {
    return ic8ba7cb6b98f0ca735a5f8fcf2e365121f1c6ef28c8c634c415454efcb1c7d03.NewShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SyncDevice()(*i55fae017c747062e34edf59c747820d3c475fc00d0605b4cee70ff8c8c5b86be.SyncDeviceRequestBuilder) {
    return i55fae017c747062e34edf59c747820d3c475fc00d0605b4cee70ff8c8c5b86be.NewSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) TriggerConfigurationManagerAction()(*ifaa214a686668e2595da116c49578208861fcd6f75a1ac2a99a03caeba4d0133.TriggerConfigurationManagerActionRequestBuilder) {
    return ifaa214a686668e2595da116c49578208861fcd6f75a1ac2a99a03caeba4d0133.NewTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) UpdateWindowsDeviceAccount()(*iaf44a1cab77523322bd155de12aa075474471071524a8fb9c17f63b6148056e7.UpdateWindowsDeviceAccountRequestBuilder) {
    return iaf44a1cab77523322bd155de12aa075474471071524a8fb9c17f63b6148056e7.NewUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderScan()(*ibd3d6ea6bfb0e2c933bac96c03b78652a9823ef84f135036962ef11eda4a74c8.WindowsDefenderScanRequestBuilder) {
    return ibd3d6ea6bfb0e2c933bac96c03b78652a9823ef84f135036962ef11eda4a74c8.NewWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderUpdateSignatures()(*i2b107af841dcd31d3787426bcfede2c74e8e60c00ae771694d98d0149871948f.WindowsDefenderUpdateSignaturesRequestBuilder) {
    return i2b107af841dcd31d3787426bcfede2c74e8e60c00ae771694d98d0149871948f.NewWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Wipe()(*i209e2b2083ba318400dcea6a198367a0ae2caa69fb8460a19cd2b9ed653d15fc.WipeRequestBuilder) {
    return i209e2b2083ba318400dcea6a198367a0ae2caa69fb8460a19cd2b9ed653d15fc.NewWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
