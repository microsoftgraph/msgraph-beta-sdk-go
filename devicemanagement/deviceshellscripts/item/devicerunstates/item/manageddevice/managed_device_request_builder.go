package manageddevice

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i075ab85f8d81a3ce4511ddfd8ade2d911af355a5199370e48801496e30273f39 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/createdevicelogcollectionrequest"
    i15ca6af0e6b6c5c390635c45f90c596f328a454e8bfc6d9a3f06edd49271c7f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/resetpasscode"
    i1720b615d9f3da31d18959222152e03ad929cb8799a3416c8b212a7a39b7f2c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/enablelostmode"
    i1c754f5b5654c2adad97b39eedb7de507b94222fdaf3ba9dfdbaff13a0c54ba9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/getfilevaultkey"
    i202c0fcc4ca2b084a611211f63c799dce1f5cfb935e9b6740fb6b2f63892e72f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/locatedevice"
    i20d05c0e6f7d0a3ff8197b873d1e4212292dcb0b7d4b95a0328f418c5cc09ca0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/ref"
    i292f88a2da536d1fe0bf25217a048605139040781a22aa17d47c18473e6f87fb "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/revokeapplevpplicenses"
    i2d07bed4f2b25936d265c489664a4630f77855756776534642d1a4e330ac1adf "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/cleanwindowsdevice"
    i364c273814fea86036ca235fc8feb6f93cdd10fb971ec54309b8711dadc3417f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/overridecompliancestate"
    i3f51f1e5e246358bc85e59ff4df3954ef8212fbf7cef665eae49b1c9927b4542 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/getoemwarranty"
    i49fe0bdac5cce938ce158b64d6391353d50a07b821419879396d76d90b9b6892 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/activatedeviceesim"
    i506dac435c764b7ff2cb4ca6b7d554ec3e1714924ea1806d55199fdab554b00f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/disable"
    i513d26c57552e229459d8e925026ef945c822c96a958d6cd1b1ec9a3adfef063 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/disablelostmode"
    i5ef2c3f841190c5f3aa5d5f553a1b067dc3757c9b7873f4a3e50b80f1717c6af "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/recoverpasscode"
    i5f21d0c942ff159c8708035b5c9b8cfeb0c4d8f2f1048a191348a3e4a6604304 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/updatewindowsdeviceaccount"
    i642b520fdcb1e1ad4fb535f39ecc3f52604cd017159f9637c4c0cb16388f159c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/syncdevice"
    i64d3215e4ca57aa09d5abc04d941765fe0ce21707667c5e3eb73bb9635f3c376 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/deleteuserfromsharedappledevice"
    i66a1abe948af272beda292d04a2d9a65b0b36911f42bb9193674147731f5d711 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/sendcustomnotificationtocompanyportal"
    i67d0be7987ca4df1c83601b7968147013918ea3309efca584a79c1b9cf4913d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/getcloudpcremoteactionresults"
    i6b5591e92bd46fcbdb21547493098fd29ea77d5f89772c9577b6cec7212417bf "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/windowsdefenderscan"
    i7898fa55af47d40246e0597eddf2485582cb961c243a950631855be1f22dd418 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/shutdown"
    i7a107dd97eefbe87c814914314db9099bf4f7561810742920b4cffa039be0bc1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/rotatebitlockerkeys"
    i8b5e9250ef9fd933d2164369a797cacfe846dbaa402165b19be5a07f65fd7e33 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/wipe"
    i903d69992327eb5d98ca375b478086a1a11734cafc662b9574e2180925bccfc1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/resizecloudpc"
    i9ca50c0daaec6fc92d7473b8c066352b6ae5d2d65bcf42dc98d4ded0be56e760 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/deprovision"
    ia366bfbd83d43a567c2102ce31034f7ac98c855597157bb6cae6e0ed19dd67a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/bypassactivationlock"
    ia65e1d9cc5ad18067df3ef979a27e91b211083d731ec2a54f3372bf301de55c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/setdevicename"
    ib1652025becd68bac80c41c9f08b651fb8a9293aec66ac4a6cf9eefbbfec2845 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/playlostmodesound"
    ib58770b246034fbac0a317ea89bf7353d2cb08b3941f59be42c3cb39093867f8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/triggerconfigurationmanageraction"
    ib87b543df792fbcc1840ef6f8495c54d6831ae77c96608d0c11c9d443eba9d60 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/requestremoteassistance"
    iba05d9be3b043899bcf2b75443f07dcf3a5470f347632a73480aa6b1470bc2fd "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/rotatefilevaultkey"
    ic6b107131f4ed8591d912604a935480f9d82988fc0188a1f065c9b70b10883f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/windowsdefenderupdatesignatures"
    ic89b86f7b12d86c1c6c03a3be0a00f8ac782faf06eea8a6c6892e5d9d8d421bb "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/retire"
    id2f9bc65912bc09f2835b52428e51ebfc0cbe1324e9aeafca8e85e614775e860 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/rebootnow"
    id7c0e19f2a7413d66b45a5b2f34c9c45f4fc369a654753f0764640e10e6ef898 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/restorecloudpc"
    id9929975793477daf5c46b290aa0225cb77d075c58ea2b78917ade5383abe58c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/reprovisioncloudpc"
    idb1115a323076a79a0947d66a1c0ab4b1ece6912a9032006b1b7a55f3c34705b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/remotelock"
    if746c9fc6a875057e2a705867ad44c6391298556bbe57f5b0d87171f3ef25912 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/reenable"
    ifb113177ed8acd5c787f5abba1840a5426cd6ee763d4562251d931d4977b16ab "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/logoutsharedappledeviceactiveuser"
    ifeccb9ae22be9f4e0f8038485b5750f1cc68b259ad6ffbd176438bec4dafa43b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item/manageddevice/getnoncompliantsettings"
)

// ManagedDeviceRequestBuilder builds and executes requests for operations under \deviceManagement\deviceShellScripts\{deviceShellScript-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice
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
func (m *ManagedDeviceRequestBuilder) ActivateDeviceEsim()(*i49fe0bdac5cce938ce158b64d6391353d50a07b821419879396d76d90b9b6892.ActivateDeviceEsimRequestBuilder) {
    return i49fe0bdac5cce938ce158b64d6391353d50a07b821419879396d76d90b9b6892.NewActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) BypassActivationLock()(*ia366bfbd83d43a567c2102ce31034f7ac98c855597157bb6cae6e0ed19dd67a2.BypassActivationLockRequestBuilder) {
    return ia366bfbd83d43a567c2102ce31034f7ac98c855597157bb6cae6e0ed19dd67a2.NewBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) CleanWindowsDevice()(*i2d07bed4f2b25936d265c489664a4630f77855756776534642d1a4e330ac1adf.CleanWindowsDeviceRequestBuilder) {
    return i2d07bed4f2b25936d265c489664a4630f77855756776534642d1a4e330ac1adf.NewCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagedDeviceRequestBuilderInternal instantiates a new ManagedDeviceRequestBuilder and sets the default values.
func NewManagedDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceRequestBuilder) {
    m := &ManagedDeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceShellScripts/{deviceShellScript_id}/deviceRunStates/{deviceManagementScriptDeviceState_id}/managedDevice{?select,expand}";
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
func (m *ManagedDeviceRequestBuilder) CreateDeviceLogCollectionRequest()(*i075ab85f8d81a3ce4511ddfd8ade2d911af355a5199370e48801496e30273f39.CreateDeviceLogCollectionRequestRequestBuilder) {
    return i075ab85f8d81a3ce4511ddfd8ade2d911af355a5199370e48801496e30273f39.NewCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ManagedDeviceRequestBuilder) DeleteUserFromSharedAppleDevice()(*i64d3215e4ca57aa09d5abc04d941765fe0ce21707667c5e3eb73bb9635f3c376.DeleteUserFromSharedAppleDeviceRequestBuilder) {
    return i64d3215e4ca57aa09d5abc04d941765fe0ce21707667c5e3eb73bb9635f3c376.NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Deprovision()(*i9ca50c0daaec6fc92d7473b8c066352b6ae5d2d65bcf42dc98d4ded0be56e760.DeprovisionRequestBuilder) {
    return i9ca50c0daaec6fc92d7473b8c066352b6ae5d2d65bcf42dc98d4ded0be56e760.NewDeprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Disable()(*i506dac435c764b7ff2cb4ca6b7d554ec3e1714924ea1806d55199fdab554b00f.DisableRequestBuilder) {
    return i506dac435c764b7ff2cb4ca6b7d554ec3e1714924ea1806d55199fdab554b00f.NewDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DisableLostMode()(*i513d26c57552e229459d8e925026ef945c822c96a958d6cd1b1ec9a3adfef063.DisableLostModeRequestBuilder) {
    return i513d26c57552e229459d8e925026ef945c822c96a958d6cd1b1ec9a3adfef063.NewDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) EnableLostMode()(*i1720b615d9f3da31d18959222152e03ad929cb8799a3416c8b212a7a39b7f2c3.EnableLostModeRequestBuilder) {
    return i1720b615d9f3da31d18959222152e03ad929cb8799a3416c8b212a7a39b7f2c3.NewEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// GetCloudPcRemoteActionResults builds and executes requests for operations under \deviceManagement\deviceShellScripts\{deviceShellScript-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.getCloudPcRemoteActionResults()
func (m *ManagedDeviceRequestBuilder) GetCloudPcRemoteActionResults()(*i67d0be7987ca4df1c83601b7968147013918ea3309efca584a79c1b9cf4913d1.GetCloudPcRemoteActionResultsRequestBuilder) {
    return i67d0be7987ca4df1c83601b7968147013918ea3309efca584a79c1b9cf4913d1.NewGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFileVaultKey builds and executes requests for operations under \deviceManagement\deviceShellScripts\{deviceShellScript-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.getFileVaultKey()
func (m *ManagedDeviceRequestBuilder) GetFileVaultKey()(*i1c754f5b5654c2adad97b39eedb7de507b94222fdaf3ba9dfdbaff13a0c54ba9.GetFileVaultKeyRequestBuilder) {
    return i1c754f5b5654c2adad97b39eedb7de507b94222fdaf3ba9dfdbaff13a0c54ba9.NewGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetNonCompliantSettings builds and executes requests for operations under \deviceManagement\deviceShellScripts\{deviceShellScript-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.getNonCompliantSettings()
func (m *ManagedDeviceRequestBuilder) GetNonCompliantSettings()(*ifeccb9ae22be9f4e0f8038485b5750f1cc68b259ad6ffbd176438bec4dafa43b.GetNonCompliantSettingsRequestBuilder) {
    return ifeccb9ae22be9f4e0f8038485b5750f1cc68b259ad6ffbd176438bec4dafa43b.NewGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOemWarranty builds and executes requests for operations under \deviceManagement\deviceShellScripts\{deviceShellScript-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.getOemWarranty()
func (m *ManagedDeviceRequestBuilder) GetOemWarranty()(*i3f51f1e5e246358bc85e59ff4df3954ef8212fbf7cef665eae49b1c9927b4542.GetOemWarrantyRequestBuilder) {
    return i3f51f1e5e246358bc85e59ff4df3954ef8212fbf7cef665eae49b1c9927b4542.NewGetOemWarrantyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LocateDevice()(*i202c0fcc4ca2b084a611211f63c799dce1f5cfb935e9b6740fb6b2f63892e72f.LocateDeviceRequestBuilder) {
    return i202c0fcc4ca2b084a611211f63c799dce1f5cfb935e9b6740fb6b2f63892e72f.NewLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*ifb113177ed8acd5c787f5abba1840a5426cd6ee763d4562251d931d4977b16ab.LogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return ifb113177ed8acd5c787f5abba1840a5426cd6ee763d4562251d931d4977b16ab.NewLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) OverrideComplianceState()(*i364c273814fea86036ca235fc8feb6f93cdd10fb971ec54309b8711dadc3417f.OverrideComplianceStateRequestBuilder) {
    return i364c273814fea86036ca235fc8feb6f93cdd10fb971ec54309b8711dadc3417f.NewOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) PlayLostModeSound()(*ib1652025becd68bac80c41c9f08b651fb8a9293aec66ac4a6cf9eefbbfec2845.PlayLostModeSoundRequestBuilder) {
    return ib1652025becd68bac80c41c9f08b651fb8a9293aec66ac4a6cf9eefbbfec2845.NewPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RebootNow()(*id2f9bc65912bc09f2835b52428e51ebfc0cbe1324e9aeafca8e85e614775e860.RebootNowRequestBuilder) {
    return id2f9bc65912bc09f2835b52428e51ebfc0cbe1324e9aeafca8e85e614775e860.NewRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RecoverPasscode()(*i5ef2c3f841190c5f3aa5d5f553a1b067dc3757c9b7873f4a3e50b80f1717c6af.RecoverPasscodeRequestBuilder) {
    return i5ef2c3f841190c5f3aa5d5f553a1b067dc3757c9b7873f4a3e50b80f1717c6af.NewRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Reenable()(*if746c9fc6a875057e2a705867ad44c6391298556bbe57f5b0d87171f3ef25912.ReenableRequestBuilder) {
    return if746c9fc6a875057e2a705867ad44c6391298556bbe57f5b0d87171f3ef25912.NewReenableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Ref()(*i20d05c0e6f7d0a3ff8197b873d1e4212292dcb0b7d4b95a0328f418c5cc09ca0.RefRequestBuilder) {
    return i20d05c0e6f7d0a3ff8197b873d1e4212292dcb0b7d4b95a0328f418c5cc09ca0.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RemoteLock()(*idb1115a323076a79a0947d66a1c0ab4b1ece6912a9032006b1b7a55f3c34705b.RemoteLockRequestBuilder) {
    return idb1115a323076a79a0947d66a1c0ab4b1ece6912a9032006b1b7a55f3c34705b.NewRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ReprovisionCloudPc()(*id9929975793477daf5c46b290aa0225cb77d075c58ea2b78917ade5383abe58c.ReprovisionCloudPcRequestBuilder) {
    return id9929975793477daf5c46b290aa0225cb77d075c58ea2b78917ade5383abe58c.NewReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RequestRemoteAssistance()(*ib87b543df792fbcc1840ef6f8495c54d6831ae77c96608d0c11c9d443eba9d60.RequestRemoteAssistanceRequestBuilder) {
    return ib87b543df792fbcc1840ef6f8495c54d6831ae77c96608d0c11c9d443eba9d60.NewRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ResetPasscode()(*i15ca6af0e6b6c5c390635c45f90c596f328a454e8bfc6d9a3f06edd49271c7f9.ResetPasscodeRequestBuilder) {
    return i15ca6af0e6b6c5c390635c45f90c596f328a454e8bfc6d9a3f06edd49271c7f9.NewResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ResizeCloudPc()(*i903d69992327eb5d98ca375b478086a1a11734cafc662b9574e2180925bccfc1.ResizeCloudPcRequestBuilder) {
    return i903d69992327eb5d98ca375b478086a1a11734cafc662b9574e2180925bccfc1.NewResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RestoreCloudPc()(*id7c0e19f2a7413d66b45a5b2f34c9c45f4fc369a654753f0764640e10e6ef898.RestoreCloudPcRequestBuilder) {
    return id7c0e19f2a7413d66b45a5b2f34c9c45f4fc369a654753f0764640e10e6ef898.NewRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Retire()(*ic89b86f7b12d86c1c6c03a3be0a00f8ac782faf06eea8a6c6892e5d9d8d421bb.RetireRequestBuilder) {
    return ic89b86f7b12d86c1c6c03a3be0a00f8ac782faf06eea8a6c6892e5d9d8d421bb.NewRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RevokeAppleVppLicenses()(*i292f88a2da536d1fe0bf25217a048605139040781a22aa17d47c18473e6f87fb.RevokeAppleVppLicensesRequestBuilder) {
    return i292f88a2da536d1fe0bf25217a048605139040781a22aa17d47c18473e6f87fb.NewRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RotateBitLockerKeys()(*i7a107dd97eefbe87c814914314db9099bf4f7561810742920b4cffa039be0bc1.RotateBitLockerKeysRequestBuilder) {
    return i7a107dd97eefbe87c814914314db9099bf4f7561810742920b4cffa039be0bc1.NewRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RotateFileVaultKey()(*iba05d9be3b043899bcf2b75443f07dcf3a5470f347632a73480aa6b1470bc2fd.RotateFileVaultKeyRequestBuilder) {
    return iba05d9be3b043899bcf2b75443f07dcf3a5470f347632a73480aa6b1470bc2fd.NewRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SendCustomNotificationToCompanyPortal()(*i66a1abe948af272beda292d04a2d9a65b0b36911f42bb9193674147731f5d711.SendCustomNotificationToCompanyPortalRequestBuilder) {
    return i66a1abe948af272beda292d04a2d9a65b0b36911f42bb9193674147731f5d711.NewSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SetDeviceName()(*ia65e1d9cc5ad18067df3ef979a27e91b211083d731ec2a54f3372bf301de55c3.SetDeviceNameRequestBuilder) {
    return ia65e1d9cc5ad18067df3ef979a27e91b211083d731ec2a54f3372bf301de55c3.NewSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ShutDown()(*i7898fa55af47d40246e0597eddf2485582cb961c243a950631855be1f22dd418.ShutDownRequestBuilder) {
    return i7898fa55af47d40246e0597eddf2485582cb961c243a950631855be1f22dd418.NewShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SyncDevice()(*i642b520fdcb1e1ad4fb535f39ecc3f52604cd017159f9637c4c0cb16388f159c.SyncDeviceRequestBuilder) {
    return i642b520fdcb1e1ad4fb535f39ecc3f52604cd017159f9637c4c0cb16388f159c.NewSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) TriggerConfigurationManagerAction()(*ib58770b246034fbac0a317ea89bf7353d2cb08b3941f59be42c3cb39093867f8.TriggerConfigurationManagerActionRequestBuilder) {
    return ib58770b246034fbac0a317ea89bf7353d2cb08b3941f59be42c3cb39093867f8.NewTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) UpdateWindowsDeviceAccount()(*i5f21d0c942ff159c8708035b5c9b8cfeb0c4d8f2f1048a191348a3e4a6604304.UpdateWindowsDeviceAccountRequestBuilder) {
    return i5f21d0c942ff159c8708035b5c9b8cfeb0c4d8f2f1048a191348a3e4a6604304.NewUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderScan()(*i6b5591e92bd46fcbdb21547493098fd29ea77d5f89772c9577b6cec7212417bf.WindowsDefenderScanRequestBuilder) {
    return i6b5591e92bd46fcbdb21547493098fd29ea77d5f89772c9577b6cec7212417bf.NewWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderUpdateSignatures()(*ic6b107131f4ed8591d912604a935480f9d82988fc0188a1f065c9b70b10883f1.WindowsDefenderUpdateSignaturesRequestBuilder) {
    return ic6b107131f4ed8591d912604a935480f9d82988fc0188a1f065c9b70b10883f1.NewWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Wipe()(*i8b5e9250ef9fd933d2164369a797cacfe846dbaa402165b19be5a07f65fd7e33.WipeRequestBuilder) {
    return i8b5e9250ef9fd933d2164369a797cacfe846dbaa402165b19be5a07f65fd7e33.NewWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
