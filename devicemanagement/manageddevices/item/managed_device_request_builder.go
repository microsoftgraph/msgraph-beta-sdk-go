package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i01b9c3c8f91d82e0acf6987d9d8a9d45e8970e9cc61e6faf4e24f9e53ef7d2c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/remotelock"
    i08036a752a811a186fe4cfbc410dc173a0482af2aea88d8ef8789f64be1a8a50 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/windowsprotectionstate"
    i0a9788a3556a4ca6e533dc1166cb3e642160b8b23aadcf53f0f54da24e25ee50 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/overridecompliancestate"
    i0e5c94a184309b3002cff816ad6e7906fadaccefe925a5a71b0eaf7da1cee573 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/devicecategory"
    i124459d67c21db7422395135441513bd6a695df9fe6596ae0672a66ea9e85780 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/reenable"
    i1a0e017de38d695b9d703bf70a4161eaea10b82614aa1d2ad5623c980ccc89b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/setdevicename"
    i1cfac124b8720ab5c719cf26bfc2d8d83f55a69d715b8813458367cee9373b60 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/getnoncompliantsettings"
    i1e9b4ce64039f5b2e7c5ab1ff6519bc2b54393233bd0de932f805ffa42b56ed5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/syncdevice"
    i1ed96ff0c3371d2b8be19e68b3491f7f051e944528a42666ad60e9a3df769b5d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/logoutsharedappledeviceactiveuser"
    i1f1456a44a9b0409729cadca40fbbdf79a42d5a2574cf3d332d1deb5f5f62202 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/rotatebitlockerkeys"
    i2011cd53e380c8be4e418adb3871ddb5cf1719dfb3ca6ee1b408b5211399c258 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/windowsdefenderupdatesignatures"
    i37fecb3bffa962bb92bae6d06b3cd6c4cc85201a2aa5cf7d9d6e15db1e249758 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/revokeapplevpplicenses"
    i3f7d5f2f782c35626e08180d2725ab649a5a2048dbc7a6dc1a4f3493ed5aed4e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/bypassactivationlock"
    i4014a9f77365c024ff8b81b517cc90c5db1deeaee89910692d948bd54130b4ab "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/users"
    i4566609add3e4bacdac10bdd501bfbc92cf44ca7aa6d3492d3c92cd84f86564a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/windowsdefenderscan"
    i4fa46045cdcf8f3536238ca61013da101771dfeaaa9fed96aaf28f151ec66162 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/devicecompliancepolicystates"
    i5103c90c1b0640a73cfc58b6d80783fcc03b5adb82eb2e098521278c8403ebbc "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/getoemwarranty"
    i511c6c2c2693c3294886a17581fa0e4993806d6e442616fc5613bb17d2f9a3e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/rebootnow"
    i54aaa976711b04a9d75b5a6b447501e988e9c492e4339da93ee652eda4f23f04 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/wipe"
    i5533b65c5e7ef6c9cf517e9d69b5af07f1cec79c9dc96399fcbe89d0ec21a145 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/restorecloudpc"
    i61cf99cf0bb474cf95f287f10d9cc99a9b9474bfc2c89d61e472f941642ebc12 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/enablelostmode"
    i63c852e0e601ed6586a514ddd8fcecf4281a23511e01d39fc62af91ac0292701 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/disablelostmode"
    i658174871e5ed8a97ec6404a2a51d6533ca0556946eb0312fbdea2092b5a0356 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/shutdown"
    i6e7ba191a5fdb7201d9a6092955773ebfa9ecb59bdbd1e57be20c7989ea54226 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/detectedapps"
    i6f21f93e6f07adafa56b75d1c7c78e9df57af9204e36460aad874e8c9617c701 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/playlostmodesound"
    i77ea5e92da231c3462e6daf6429f47a312c08417255e97ac007a61b747062e52 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/manageddevicemobileappconfigurationstates"
    i7d468003528ad5d962d3cce255cb40c4f894227cc59c71443390e09ab8ca40db "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/getcloudpcremoteactionresults"
    i7d4fd4db8c253e46353ca32adf5b0f49f79c9b16b0709c7adc865da540dc64f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/logcollectionrequests"
    i86e677b9939bcc1c1f4b1438dd0178ec6f297b7d3c5baa413730c16ab30ff0a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/cleanwindowsdevice"
    i9479492bd537ee0da9d7cbaf1a7bda614d8af9b2db9f340f4d162cfdab45d7d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/resetpasscode"
    i964b81837112ab4b6bfd1cf6220e1e8cb634c0c56d1238300c9c00340217450d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/resizecloudpc"
    i9b07e3d1ad25d8bac7ab04fef250c221091aa5e13eba25491edfadca2f67508b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/triggerconfigurationmanageraction"
    i9c36cb050999ff8fab3434eaa458ff65d63d50ae824d7965d7bd33cff08b058a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/rotatefilevaultkey"
    ia36af35bae5da4cc0bc0059a8dd1595ea3098bbaf66a723ad498408f9329538d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/requestremoteassistance"
    ia8bf08ad44640f5a1416c953b553221a5b755b2af9acc8770a00a1522e4650a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/locatedevice"
    iabf61e946f47f041682806eff385ca77fa52f76f6756d068e2e853bebc2c381b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/sendcustomnotificationtocompanyportal"
    ib4d6bc26f8783e0021130bc5307c6d883c818e4a41f6cb4f64641101753c3a9c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/deprovision"
    iba0e44812d3f93778c7d884706c3eb3fad191cee346e2a62205385f2c63dc366 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/getfilevaultkey"
    ic41b8c2fdd1ab9c9e0d3c3744785e36b1c26031b43bec2708dcfb79053ed891b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/activatedeviceesim"
    ic4a4e554ac185065d3036fce9e0be725936f6f53aefb12ab4ac38f451bb309a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/createdevicelogcollectionrequest"
    ic7b58bd1f7ec7e59ddafda5fb2b6789ce5abdfcba5b1e47d8b82e90e8daeb055 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/reprovisioncloudpc"
    icb6ec007efa06597d8fa269fd84ddc90f7dd27cc757bf2d9fc8f6249c2efc2da "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/deleteuserfromsharedappledevice"
    id1fd9f59ff6efc1a25cf2a7067af3ac8949a4357a777c40891113b5c6f9a3a97 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/updatewindowsdeviceaccount"
    ida5dd88d263aad80639be0db6c9f33ac2ef69ddb1701e98adef543f2d2e39c09 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/disable"
    ie1efb01c1dca39b9cc510aab2614914c0ae675e30562a196406520ea8c0163e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/securitybaselinestates"
    ie44ea0c4c2478526c3387d0067c7beeb549f1ececdbea0c42e565e76c317dd35 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/deviceconfigurationstates"
    if75bb98073f1131ee02501bf098984c52a24d708147db3112ea968e2fc7b2c03 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/retire"
    ifb1c8b230156b71d830111db728223a0fcfda4e39c490d27d18e1957c1c35ab9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/recoverpasscode"
    iff1aefed3994b5a74430f7805d686fd8fecf6a36de439aad9dfcb6522e0b8308 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/assignmentfilterevaluationstatusdetails"
    i0f7c3711eb6736d4154b8664f763bba6cd17670f2f39b4f1a6cf7fd15aa16312 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/devicecompliancepolicystates/item"
    i299912d62f14997dd65aac2ade28936d603ff421576c4948af1d264ab7e90bd0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/manageddevicemobileappconfigurationstates/item"
    i8bf3e9602de99585216828ab1a7ba1b28f64f44b09c1192ca701225a40bdb050 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/securitybaselinestates/item"
    ib34591ad39ce97862312af86b9741ca3b1a7d0746218c1f1269fecd90b2e0ee6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/logcollectionrequests/item"
    ibf3c9129b1fe1c1f3482c4943cfae633d290cb95cfe0dec78735ebb6afbf6ac2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/manageddevices/item/deviceconfigurationstates/item"
)

// ManagedDeviceRequestBuilder builds and executes requests for operations under \deviceManagement\managedDevices\{managedDevice-id}
type ManagedDeviceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagedDeviceRequestBuilderDeleteOptions options for Delete
type ManagedDeviceRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
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
// ManagedDeviceRequestBuilderGetQueryParameters the list of managed devices.
type ManagedDeviceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ManagedDeviceRequestBuilderPatchOptions options for Patch
type ManagedDeviceRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDevice;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ManagedDeviceRequestBuilder) ActivateDeviceEsim()(*ic41b8c2fdd1ab9c9e0d3c3744785e36b1c26031b43bec2708dcfb79053ed891b.ActivateDeviceEsimRequestBuilder) {
    return ic41b8c2fdd1ab9c9e0d3c3744785e36b1c26031b43bec2708dcfb79053ed891b.NewActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) AssignmentFilterEvaluationStatusDetails()(*iff1aefed3994b5a74430f7805d686fd8fecf6a36de439aad9dfcb6522e0b8308.AssignmentFilterEvaluationStatusDetailsRequestBuilder) {
    return iff1aefed3994b5a74430f7805d686fd8fecf6a36de439aad9dfcb6522e0b8308.NewAssignmentFilterEvaluationStatusDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentFilterEvaluationStatusDetailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.managedDevices.item.assignmentFilterEvaluationStatusDetails.item collection
func (m *ManagedDeviceRequestBuilder) AssignmentFilterEvaluationStatusDetailsById(id string)(*iff1aefed3994b5a74430f7805d686fd8fecf6a36de439aad9dfcb6522e0b8308.AssignmentFilterEvaluationStatusDetailsRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["assignmentFilterEvaluationStatusDetails_id"] = id
    }
    return iff1aefed3994b5a74430f7805d686fd8fecf6a36de439aad9dfcb6522e0b8308.NewAssignmentFilterEvaluationStatusDetailsRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) BypassActivationLock()(*i3f7d5f2f782c35626e08180d2725ab649a5a2048dbc7a6dc1a4f3493ed5aed4e.BypassActivationLockRequestBuilder) {
    return i3f7d5f2f782c35626e08180d2725ab649a5a2048dbc7a6dc1a4f3493ed5aed4e.NewBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) CleanWindowsDevice()(*i86e677b9939bcc1c1f4b1438dd0178ec6f297b7d3c5baa413730c16ab30ff0a9.CleanWindowsDeviceRequestBuilder) {
    return i86e677b9939bcc1c1f4b1438dd0178ec6f297b7d3c5baa413730c16ab30ff0a9.NewCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagedDeviceRequestBuilderInternal instantiates a new ManagedDeviceRequestBuilder and sets the default values.
func NewManagedDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceRequestBuilder) {
    m := &ManagedDeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/managedDevices/{managedDevice_id}{?select,expand}";
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
// CreateDeleteRequestInformation the list of managed devices.
func (m *ManagedDeviceRequestBuilder) CreateDeleteRequestInformation(options *ManagedDeviceRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ManagedDeviceRequestBuilder) CreateDeviceLogCollectionRequest()(*ic4a4e554ac185065d3036fce9e0be725936f6f53aefb12ab4ac38f451bb309a7.CreateDeviceLogCollectionRequestRequestBuilder) {
    return ic4a4e554ac185065d3036fce9e0be725936f6f53aefb12ab4ac38f451bb309a7.NewCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the list of managed devices.
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
// CreatePatchRequestInformation the list of managed devices.
func (m *ManagedDeviceRequestBuilder) CreatePatchRequestInformation(options *ManagedDeviceRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the list of managed devices.
func (m *ManagedDeviceRequestBuilder) Delete(options *ManagedDeviceRequestBuilderDeleteOptions)(error) {
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
func (m *ManagedDeviceRequestBuilder) DeleteUserFromSharedAppleDevice()(*icb6ec007efa06597d8fa269fd84ddc90f7dd27cc757bf2d9fc8f6249c2efc2da.DeleteUserFromSharedAppleDeviceRequestBuilder) {
    return icb6ec007efa06597d8fa269fd84ddc90f7dd27cc757bf2d9fc8f6249c2efc2da.NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Deprovision()(*ib4d6bc26f8783e0021130bc5307c6d883c818e4a41f6cb4f64641101753c3a9c.DeprovisionRequestBuilder) {
    return ib4d6bc26f8783e0021130bc5307c6d883c818e4a41f6cb4f64641101753c3a9c.NewDeprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DetectedApps()(*i6e7ba191a5fdb7201d9a6092955773ebfa9ecb59bdbd1e57be20c7989ea54226.DetectedAppsRequestBuilder) {
    return i6e7ba191a5fdb7201d9a6092955773ebfa9ecb59bdbd1e57be20c7989ea54226.NewDetectedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DeviceCategory()(*i0e5c94a184309b3002cff816ad6e7906fadaccefe925a5a71b0eaf7da1cee573.DeviceCategoryRequestBuilder) {
    return i0e5c94a184309b3002cff816ad6e7906fadaccefe925a5a71b0eaf7da1cee573.NewDeviceCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DeviceCompliancePolicyStates()(*i4fa46045cdcf8f3536238ca61013da101771dfeaaa9fed96aaf28f151ec66162.DeviceCompliancePolicyStatesRequestBuilder) {
    return i4fa46045cdcf8f3536238ca61013da101771dfeaaa9fed96aaf28f151ec66162.NewDeviceCompliancePolicyStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicyStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.managedDevices.item.deviceCompliancePolicyStates.item collection
func (m *ManagedDeviceRequestBuilder) DeviceCompliancePolicyStatesById(id string)(*i0f7c3711eb6736d4154b8664f763bba6cd17670f2f39b4f1a6cf7fd15aa16312.DeviceCompliancePolicyStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyState_id"] = id
    }
    return i0f7c3711eb6736d4154b8664f763bba6cd17670f2f39b4f1a6cf7fd15aa16312.NewDeviceCompliancePolicyStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DeviceConfigurationStates()(*ie44ea0c4c2478526c3387d0067c7beeb549f1ececdbea0c42e565e76c317dd35.DeviceConfigurationStatesRequestBuilder) {
    return ie44ea0c4c2478526c3387d0067c7beeb549f1ececdbea0c42e565e76c317dd35.NewDeviceConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.managedDevices.item.deviceConfigurationStates.item collection
func (m *ManagedDeviceRequestBuilder) DeviceConfigurationStatesById(id string)(*ibf3c9129b1fe1c1f3482c4943cfae633d290cb95cfe0dec78735ebb6afbf6ac2.DeviceConfigurationStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationState_id"] = id
    }
    return ibf3c9129b1fe1c1f3482c4943cfae633d290cb95cfe0dec78735ebb6afbf6ac2.NewDeviceConfigurationStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Disable()(*ida5dd88d263aad80639be0db6c9f33ac2ef69ddb1701e98adef543f2d2e39c09.DisableRequestBuilder) {
    return ida5dd88d263aad80639be0db6c9f33ac2ef69ddb1701e98adef543f2d2e39c09.NewDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DisableLostMode()(*i63c852e0e601ed6586a514ddd8fcecf4281a23511e01d39fc62af91ac0292701.DisableLostModeRequestBuilder) {
    return i63c852e0e601ed6586a514ddd8fcecf4281a23511e01d39fc62af91ac0292701.NewDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) EnableLostMode()(*i61cf99cf0bb474cf95f287f10d9cc99a9b9474bfc2c89d61e472f941642ebc12.EnableLostModeRequestBuilder) {
    return i61cf99cf0bb474cf95f287f10d9cc99a9b9474bfc2c89d61e472f941642ebc12.NewEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the list of managed devices.
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
// GetCloudPcRemoteActionResults builds and executes requests for operations under \deviceManagement\managedDevices\{managedDevice-id}\microsoft.graph.getCloudPcRemoteActionResults()
func (m *ManagedDeviceRequestBuilder) GetCloudPcRemoteActionResults()(*i7d468003528ad5d962d3cce255cb40c4f894227cc59c71443390e09ab8ca40db.GetCloudPcRemoteActionResultsRequestBuilder) {
    return i7d468003528ad5d962d3cce255cb40c4f894227cc59c71443390e09ab8ca40db.NewGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFileVaultKey builds and executes requests for operations under \deviceManagement\managedDevices\{managedDevice-id}\microsoft.graph.getFileVaultKey()
func (m *ManagedDeviceRequestBuilder) GetFileVaultKey()(*iba0e44812d3f93778c7d884706c3eb3fad191cee346e2a62205385f2c63dc366.GetFileVaultKeyRequestBuilder) {
    return iba0e44812d3f93778c7d884706c3eb3fad191cee346e2a62205385f2c63dc366.NewGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetNonCompliantSettings builds and executes requests for operations under \deviceManagement\managedDevices\{managedDevice-id}\microsoft.graph.getNonCompliantSettings()
func (m *ManagedDeviceRequestBuilder) GetNonCompliantSettings()(*i1cfac124b8720ab5c719cf26bfc2d8d83f55a69d715b8813458367cee9373b60.GetNonCompliantSettingsRequestBuilder) {
    return i1cfac124b8720ab5c719cf26bfc2d8d83f55a69d715b8813458367cee9373b60.NewGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOemWarranty builds and executes requests for operations under \deviceManagement\managedDevices\{managedDevice-id}\microsoft.graph.getOemWarranty()
func (m *ManagedDeviceRequestBuilder) GetOemWarranty()(*i5103c90c1b0640a73cfc58b6d80783fcc03b5adb82eb2e098521278c8403ebbc.GetOemWarrantyRequestBuilder) {
    return i5103c90c1b0640a73cfc58b6d80783fcc03b5adb82eb2e098521278c8403ebbc.NewGetOemWarrantyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LocateDevice()(*ia8bf08ad44640f5a1416c953b553221a5b755b2af9acc8770a00a1522e4650a6.LocateDeviceRequestBuilder) {
    return ia8bf08ad44640f5a1416c953b553221a5b755b2af9acc8770a00a1522e4650a6.NewLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LogCollectionRequests()(*i7d4fd4db8c253e46353ca32adf5b0f49f79c9b16b0709c7adc865da540dc64f4.LogCollectionRequestsRequestBuilder) {
    return i7d4fd4db8c253e46353ca32adf5b0f49f79c9b16b0709c7adc865da540dc64f4.NewLogCollectionRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LogCollectionRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.managedDevices.item.logCollectionRequests.item collection
func (m *ManagedDeviceRequestBuilder) LogCollectionRequestsById(id string)(*ib34591ad39ce97862312af86b9741ca3b1a7d0746218c1f1269fecd90b2e0ee6.DeviceLogCollectionResponseRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceLogCollectionResponse_id"] = id
    }
    return ib34591ad39ce97862312af86b9741ca3b1a7d0746218c1f1269fecd90b2e0ee6.NewDeviceLogCollectionResponseRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*i1ed96ff0c3371d2b8be19e68b3491f7f051e944528a42666ad60e9a3df769b5d.LogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return i1ed96ff0c3371d2b8be19e68b3491f7f051e944528a42666ad60e9a3df769b5d.NewLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ManagedDeviceMobileAppConfigurationStates()(*i77ea5e92da231c3462e6daf6429f47a312c08417255e97ac007a61b747062e52.ManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    return i77ea5e92da231c3462e6daf6429f47a312c08417255e97ac007a61b747062e52.NewManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceMobileAppConfigurationStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.managedDevices.item.managedDeviceMobileAppConfigurationStates.item collection
func (m *ManagedDeviceRequestBuilder) ManagedDeviceMobileAppConfigurationStatesById(id string)(*i299912d62f14997dd65aac2ade28936d603ff421576c4948af1d264ab7e90bd0.ManagedDeviceMobileAppConfigurationStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationState_id"] = id
    }
    return i299912d62f14997dd65aac2ade28936d603ff421576c4948af1d264ab7e90bd0.NewManagedDeviceMobileAppConfigurationStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) OverrideComplianceState()(*i0a9788a3556a4ca6e533dc1166cb3e642160b8b23aadcf53f0f54da24e25ee50.OverrideComplianceStateRequestBuilder) {
    return i0a9788a3556a4ca6e533dc1166cb3e642160b8b23aadcf53f0f54da24e25ee50.NewOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the list of managed devices.
func (m *ManagedDeviceRequestBuilder) Patch(options *ManagedDeviceRequestBuilderPatchOptions)(error) {
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
func (m *ManagedDeviceRequestBuilder) PlayLostModeSound()(*i6f21f93e6f07adafa56b75d1c7c78e9df57af9204e36460aad874e8c9617c701.PlayLostModeSoundRequestBuilder) {
    return i6f21f93e6f07adafa56b75d1c7c78e9df57af9204e36460aad874e8c9617c701.NewPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RebootNow()(*i511c6c2c2693c3294886a17581fa0e4993806d6e442616fc5613bb17d2f9a3e1.RebootNowRequestBuilder) {
    return i511c6c2c2693c3294886a17581fa0e4993806d6e442616fc5613bb17d2f9a3e1.NewRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RecoverPasscode()(*ifb1c8b230156b71d830111db728223a0fcfda4e39c490d27d18e1957c1c35ab9.RecoverPasscodeRequestBuilder) {
    return ifb1c8b230156b71d830111db728223a0fcfda4e39c490d27d18e1957c1c35ab9.NewRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Reenable()(*i124459d67c21db7422395135441513bd6a695df9fe6596ae0672a66ea9e85780.ReenableRequestBuilder) {
    return i124459d67c21db7422395135441513bd6a695df9fe6596ae0672a66ea9e85780.NewReenableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RemoteLock()(*i01b9c3c8f91d82e0acf6987d9d8a9d45e8970e9cc61e6faf4e24f9e53ef7d2c6.RemoteLockRequestBuilder) {
    return i01b9c3c8f91d82e0acf6987d9d8a9d45e8970e9cc61e6faf4e24f9e53ef7d2c6.NewRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ReprovisionCloudPc()(*ic7b58bd1f7ec7e59ddafda5fb2b6789ce5abdfcba5b1e47d8b82e90e8daeb055.ReprovisionCloudPcRequestBuilder) {
    return ic7b58bd1f7ec7e59ddafda5fb2b6789ce5abdfcba5b1e47d8b82e90e8daeb055.NewReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RequestRemoteAssistance()(*ia36af35bae5da4cc0bc0059a8dd1595ea3098bbaf66a723ad498408f9329538d.RequestRemoteAssistanceRequestBuilder) {
    return ia36af35bae5da4cc0bc0059a8dd1595ea3098bbaf66a723ad498408f9329538d.NewRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ResetPasscode()(*i9479492bd537ee0da9d7cbaf1a7bda614d8af9b2db9f340f4d162cfdab45d7d4.ResetPasscodeRequestBuilder) {
    return i9479492bd537ee0da9d7cbaf1a7bda614d8af9b2db9f340f4d162cfdab45d7d4.NewResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ResizeCloudPc()(*i964b81837112ab4b6bfd1cf6220e1e8cb634c0c56d1238300c9c00340217450d.ResizeCloudPcRequestBuilder) {
    return i964b81837112ab4b6bfd1cf6220e1e8cb634c0c56d1238300c9c00340217450d.NewResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RestoreCloudPc()(*i5533b65c5e7ef6c9cf517e9d69b5af07f1cec79c9dc96399fcbe89d0ec21a145.RestoreCloudPcRequestBuilder) {
    return i5533b65c5e7ef6c9cf517e9d69b5af07f1cec79c9dc96399fcbe89d0ec21a145.NewRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Retire()(*if75bb98073f1131ee02501bf098984c52a24d708147db3112ea968e2fc7b2c03.RetireRequestBuilder) {
    return if75bb98073f1131ee02501bf098984c52a24d708147db3112ea968e2fc7b2c03.NewRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RevokeAppleVppLicenses()(*i37fecb3bffa962bb92bae6d06b3cd6c4cc85201a2aa5cf7d9d6e15db1e249758.RevokeAppleVppLicensesRequestBuilder) {
    return i37fecb3bffa962bb92bae6d06b3cd6c4cc85201a2aa5cf7d9d6e15db1e249758.NewRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RotateBitLockerKeys()(*i1f1456a44a9b0409729cadca40fbbdf79a42d5a2574cf3d332d1deb5f5f62202.RotateBitLockerKeysRequestBuilder) {
    return i1f1456a44a9b0409729cadca40fbbdf79a42d5a2574cf3d332d1deb5f5f62202.NewRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RotateFileVaultKey()(*i9c36cb050999ff8fab3434eaa458ff65d63d50ae824d7965d7bd33cff08b058a.RotateFileVaultKeyRequestBuilder) {
    return i9c36cb050999ff8fab3434eaa458ff65d63d50ae824d7965d7bd33cff08b058a.NewRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SecurityBaselineStates()(*ie1efb01c1dca39b9cc510aab2614914c0ae675e30562a196406520ea8c0163e9.SecurityBaselineStatesRequestBuilder) {
    return ie1efb01c1dca39b9cc510aab2614914c0ae675e30562a196406520ea8c0163e9.NewSecurityBaselineStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecurityBaselineStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.managedDevices.item.securityBaselineStates.item collection
func (m *ManagedDeviceRequestBuilder) SecurityBaselineStatesById(id string)(*i8bf3e9602de99585216828ab1a7ba1b28f64f44b09c1192ca701225a40bdb050.SecurityBaselineStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["securityBaselineState_id"] = id
    }
    return i8bf3e9602de99585216828ab1a7ba1b28f64f44b09c1192ca701225a40bdb050.NewSecurityBaselineStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SendCustomNotificationToCompanyPortal()(*iabf61e946f47f041682806eff385ca77fa52f76f6756d068e2e853bebc2c381b.SendCustomNotificationToCompanyPortalRequestBuilder) {
    return iabf61e946f47f041682806eff385ca77fa52f76f6756d068e2e853bebc2c381b.NewSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SetDeviceName()(*i1a0e017de38d695b9d703bf70a4161eaea10b82614aa1d2ad5623c980ccc89b4.SetDeviceNameRequestBuilder) {
    return i1a0e017de38d695b9d703bf70a4161eaea10b82614aa1d2ad5623c980ccc89b4.NewSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ShutDown()(*i658174871e5ed8a97ec6404a2a51d6533ca0556946eb0312fbdea2092b5a0356.ShutDownRequestBuilder) {
    return i658174871e5ed8a97ec6404a2a51d6533ca0556946eb0312fbdea2092b5a0356.NewShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SyncDevice()(*i1e9b4ce64039f5b2e7c5ab1ff6519bc2b54393233bd0de932f805ffa42b56ed5.SyncDeviceRequestBuilder) {
    return i1e9b4ce64039f5b2e7c5ab1ff6519bc2b54393233bd0de932f805ffa42b56ed5.NewSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) TriggerConfigurationManagerAction()(*i9b07e3d1ad25d8bac7ab04fef250c221091aa5e13eba25491edfadca2f67508b.TriggerConfigurationManagerActionRequestBuilder) {
    return i9b07e3d1ad25d8bac7ab04fef250c221091aa5e13eba25491edfadca2f67508b.NewTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) UpdateWindowsDeviceAccount()(*id1fd9f59ff6efc1a25cf2a7067af3ac8949a4357a777c40891113b5c6f9a3a97.UpdateWindowsDeviceAccountRequestBuilder) {
    return id1fd9f59ff6efc1a25cf2a7067af3ac8949a4357a777c40891113b5c6f9a3a97.NewUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Users()(*i4014a9f77365c024ff8b81b517cc90c5db1deeaee89910692d948bd54130b4ab.UsersRequestBuilder) {
    return i4014a9f77365c024ff8b81b517cc90c5db1deeaee89910692d948bd54130b4ab.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderScan()(*i4566609add3e4bacdac10bdd501bfbc92cf44ca7aa6d3492d3c92cd84f86564a.WindowsDefenderScanRequestBuilder) {
    return i4566609add3e4bacdac10bdd501bfbc92cf44ca7aa6d3492d3c92cd84f86564a.NewWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderUpdateSignatures()(*i2011cd53e380c8be4e418adb3871ddb5cf1719dfb3ca6ee1b408b5211399c258.WindowsDefenderUpdateSignaturesRequestBuilder) {
    return i2011cd53e380c8be4e418adb3871ddb5cf1719dfb3ca6ee1b408b5211399c258.NewWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsProtectionState()(*i08036a752a811a186fe4cfbc410dc173a0482af2aea88d8ef8789f64be1a8a50.WindowsProtectionStateRequestBuilder) {
    return i08036a752a811a186fe4cfbc410dc173a0482af2aea88d8ef8789f64be1a8a50.NewWindowsProtectionStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Wipe()(*i54aaa976711b04a9d75b5a6b447501e988e9c492e4339da93ee652eda4f23f04.WipeRequestBuilder) {
    return i54aaa976711b04a9d75b5a6b447501e988e9c492e4339da93ee652eda4f23f04.NewWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
