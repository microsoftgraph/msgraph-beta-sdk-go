package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0265d37d4c192c128f26295070675911ea6710cb341155b618c000a508cfa91d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/shutdown"
    i06f050446fb7bd265e92527bb85f6f2a987efadd05b06c81151de522ed58e0ef "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/revokeapplevpplicenses"
    i145cd047a0f177065361e7e02b132b0f8ce1e436d675e88c887b86e5f37651a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/windowsdefenderscan"
    i155c4a30ec87223697d0afd5c81bef559d728417cc1bde95d7dabed695d3d9a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/assignmentfilterevaluationstatusdetails"
    i1947898e3bbb95e8575558ada088e5fe3120b250ffef60fca81cfa321a720d05 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/getnoncompliantsettings"
    i21902135131fa38cbfc951b49b1a03fc87545fa77bf484d3da70470837d99e8d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/enablelostmode"
    i23c0ee870e4a32a5f5380e1c4d9ac0000fa720e1310201d316c3d1ed1bc38429 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/detectedapps"
    i274fa57cac64761436e6b96aa0a35ef9ef6bac8d924e432b83868a7b27e98c4d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/cleanwindowsdevice"
    i2a36033ee9d2b5fd75a9ff828a2a4949c7937585f41afbe202af5de47127dd1b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/manageddevicemobileappconfigurationstates"
    i2a373154cbcecc73fa6269b089925214127c4b5c2d032275d9319e9099e011ba "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/getfilevaultkey"
    i2df51e7699d7e946c4c71f22c064f86f0b807928512258930f58dcfdb5db3cc6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/restorecloudpc"
    i310834334f4240786fcd6b5675b3ee8fbeb94f1c4b1b200e7e0b47c655de6a72 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/deviceconfigurationstates"
    i3190268115ebe3af8fbbb7f7b8f73864d5604f943cdd020b77129dffe05358c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/rebootnow"
    i31eb74f02efdd35067936e3939c2d3f35c2731c77e2ba3e2916be25ef567d1db "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/rotatebitlockerkeys"
    i397317cd5c150ef73c8dbfe046b3a5a67e66696c131b64044418b32848f78b02 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/sendcustomnotificationtocompanyportal"
    i3eeb66f94f27e28ae758717375265f4c4a2b99c3b3ebf1d81e6148870d288958 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/wipe"
    i474daedfd97083a2b70cfaaee5db2b7ac23cd62bfacff938adc2ec3734dca31a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/locatedevice"
    i4f5f8cca618fea6d5b5be4aaf673f53045f46c671985beb22ae0ab9f9e5532c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/deprovision"
    i523f2fe01a487001fe905fdd45d3f80a080ff6e83855beac7dba457c171312fd "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/disablelostmode"
    i5a2cb0524ba36f3452cc942efde3281af6e49bd918b07ec6d4553a95679f4244 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/getcloudpcremoteactionresults"
    i5c1b61352546e0e98f55231ab68a18b6252bb9ae9d227682a4961bb808486269 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/reprovisioncloudpc"
    i5f3cf0b3cdbd4ccbd408a7bbcfb1134713761ba963b39a1a0e97c2ada49c6e64 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/logcollectionrequests"
    i6273d621d2071ff7812d4ccc27138d3c37af406b8289aba4d972bb96c54c8c65 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/resizecloudpc"
    i66c8089368fd5b85bebc1b3902964a92f8983bc08608c1f9f7decfa3c2f65f26 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/bypassactivationlock"
    i69332727d899c5dfc77f80ae621a8729902a82541ca3df96c5050d6b45b50fe5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/setdevicename"
    i69683f121195c82c8db1f6f852405aa25762f58bd7b0efa9c356fb6260734ed3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/users"
    i6afbb8bdf1ff302cabb66d3e0a2c4212f688793e011fb7452c60d29077219ccc "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/logoutsharedappledeviceactiveuser"
    i7fe5a60c553129a9db80152bf36f67e1f4ee971a86dccf62668cdb5c4966077b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/getoemwarranty"
    i86b6189da9f672cea7ec2267f0c2bc9f2ab30cc1cce295e84df98894cde4ad21 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/windowsdefenderupdatesignatures"
    i8aba118121dfe13e2a7be43f8f0b1b73de8aaf2a1d21951caab199d8a2d84627 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/deleteuserfromsharedappledevice"
    i8c92d1f6b3158dce0d6c8344255235e2ea626a1c9b3ebd7ca2f5ab067d9e9e75 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/securitybaselinestates"
    i9da99e12e59b05474451fa26fb9ba2a61a1e7f3d2d11e16bb6107056d1bf1d18 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/resetpasscode"
    ia29b36d17801e439026505c34550ac7cee3a69c1ff4c5415c194f9ac625e44bf "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/syncdevice"
    ia9e35aa934485810b7f428c752d5b72ac6c7bb25b0203c31668c6ffeb9016989 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/requestremoteassistance"
    iaab41b7756caa3ffe90d804bde7b8e128221af504e2142a1fe034770c8fe02f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/rotatefilevaultkey"
    iab5d4b97db65f3530dcf7249ef64d333c528c95523c7d2f233e65624d9c3aed6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/playlostmodesound"
    ib5f935e060b5bc7e8d0d5d9de0d928d9a660e962724c3cea1fd5cab5cff38ecc "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/windowsprotectionstate"
    ib9cb36346c054b7f4cf11cfb47fbe3a66ed7d2cdfbab1e96a8571434aa769b46 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/activatedeviceesim"
    ibf345d4df2f0f8a237eae842ca384dd517ada2627beba3c7d7f53baa0aac9cd3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/devicecategory"
    ic814e3839bc6f855c24fd1a5a83ccb31ddd05002e9bd4af67e2782c60291bd1d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/retire"
    id23da8ff63dabfc5e2f8ae52fc2c56079ba902db00a78ac2b11ded16d7c94017 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/recoverpasscode"
    id475b3b19e63c8aea08d85fbab4ad351d8d0271fec24c7176711f0a0a8a45d62 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/triggerconfigurationmanageraction"
    ida779665108e67dcd8e766def39e06d84820ead54e3bd031f8b8be0d162618f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/reenable"
    idcab09b434ee10bc36b46bee67d1c1ee38cc5810295e2f26b4055875b7af37dc "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/disable"
    iddee818f039172537f45ef3559b2d06b97b848792f16ac94552dcfd877edcdf4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/updatewindowsdeviceaccount"
    ie4a14f1c75e88f719578d40ac5c4735db9978a85bb0f9336dd79da290c863f3a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/createdevicelogcollectionrequest"
    iedfe7deabe7e5f286b481b309b885e361a1fc70fcd9dd27c13cc31b607533a7d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/devicecompliancepolicystates"
    ifbbfe5c6464ce13a6e4499ee4f52892cd67d8abbfdced67cc6d8065e47e9249b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/remotelock"
    ife6aae505188bb4c2c63cb0aacab82413df394100290381167a1feb35be9f935 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/overridecompliancestate"
    i1bc1487940b05ce412b4ae41cbb75ff7066428095ef5956e8b49523af739ae12 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/logcollectionrequests/item"
    i1cf54773ea201b86cd3ff29dc00e0c4acb254a4bdf0fd5df1f70ef4013a36e72 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/devicecompliancepolicystates/item"
    i287d2bc1eaa2bb4787a75a13a96c399c86ed7cfa8a05bd662963aca0b4dc468f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/securitybaselinestates/item"
    i3854f67f765bed0207af3dba7d579b83beb1633e1d2a7d9e98b9579399e553ae "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/manageddevicemobileappconfigurationstates/item"
    i88c0ef2f8129d7022e7594e01e34ed8d304c73d763ba681efe40080e626ebd3e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/item/deviceconfigurationstates/item"
)

// ManagedDeviceRequestBuilder builds and executes requests for operations under \deviceManagement\comanagedDevices\{managedDevice-id}
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
// ManagedDeviceRequestBuilderGetQueryParameters the list of co-managed devices report
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
func (m *ManagedDeviceRequestBuilder) ActivateDeviceEsim()(*ib9cb36346c054b7f4cf11cfb47fbe3a66ed7d2cdfbab1e96a8571434aa769b46.ActivateDeviceEsimRequestBuilder) {
    return ib9cb36346c054b7f4cf11cfb47fbe3a66ed7d2cdfbab1e96a8571434aa769b46.NewActivateDeviceEsimRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) AssignmentFilterEvaluationStatusDetails()(*i155c4a30ec87223697d0afd5c81bef559d728417cc1bde95d7dabed695d3d9a3.AssignmentFilterEvaluationStatusDetailsRequestBuilder) {
    return i155c4a30ec87223697d0afd5c81bef559d728417cc1bde95d7dabed695d3d9a3.NewAssignmentFilterEvaluationStatusDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentFilterEvaluationStatusDetailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.comanagedDevices.item.assignmentFilterEvaluationStatusDetails.item collection
func (m *ManagedDeviceRequestBuilder) AssignmentFilterEvaluationStatusDetailsById(id string)(*i155c4a30ec87223697d0afd5c81bef559d728417cc1bde95d7dabed695d3d9a3.AssignmentFilterEvaluationStatusDetailsRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["assignmentFilterEvaluationStatusDetails_id"] = id
    }
    return i155c4a30ec87223697d0afd5c81bef559d728417cc1bde95d7dabed695d3d9a3.NewAssignmentFilterEvaluationStatusDetailsRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) BypassActivationLock()(*i66c8089368fd5b85bebc1b3902964a92f8983bc08608c1f9f7decfa3c2f65f26.BypassActivationLockRequestBuilder) {
    return i66c8089368fd5b85bebc1b3902964a92f8983bc08608c1f9f7decfa3c2f65f26.NewBypassActivationLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) CleanWindowsDevice()(*i274fa57cac64761436e6b96aa0a35ef9ef6bac8d924e432b83868a7b27e98c4d.CleanWindowsDeviceRequestBuilder) {
    return i274fa57cac64761436e6b96aa0a35ef9ef6bac8d924e432b83868a7b27e98c4d.NewCleanWindowsDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagedDeviceRequestBuilderInternal instantiates a new ManagedDeviceRequestBuilder and sets the default values.
func NewManagedDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceRequestBuilder) {
    m := &ManagedDeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice_id}{?select,expand}";
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
// CreateDeleteRequestInformation the list of co-managed devices report
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
func (m *ManagedDeviceRequestBuilder) CreateDeviceLogCollectionRequest()(*ie4a14f1c75e88f719578d40ac5c4735db9978a85bb0f9336dd79da290c863f3a.CreateDeviceLogCollectionRequestRequestBuilder) {
    return ie4a14f1c75e88f719578d40ac5c4735db9978a85bb0f9336dd79da290c863f3a.NewCreateDeviceLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the list of co-managed devices report
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
// CreatePatchRequestInformation the list of co-managed devices report
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
// Delete the list of co-managed devices report
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
func (m *ManagedDeviceRequestBuilder) DeleteUserFromSharedAppleDevice()(*i8aba118121dfe13e2a7be43f8f0b1b73de8aaf2a1d21951caab199d8a2d84627.DeleteUserFromSharedAppleDeviceRequestBuilder) {
    return i8aba118121dfe13e2a7be43f8f0b1b73de8aaf2a1d21951caab199d8a2d84627.NewDeleteUserFromSharedAppleDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Deprovision()(*i4f5f8cca618fea6d5b5be4aaf673f53045f46c671985beb22ae0ab9f9e5532c8.DeprovisionRequestBuilder) {
    return i4f5f8cca618fea6d5b5be4aaf673f53045f46c671985beb22ae0ab9f9e5532c8.NewDeprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DetectedApps()(*i23c0ee870e4a32a5f5380e1c4d9ac0000fa720e1310201d316c3d1ed1bc38429.DetectedAppsRequestBuilder) {
    return i23c0ee870e4a32a5f5380e1c4d9ac0000fa720e1310201d316c3d1ed1bc38429.NewDetectedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DeviceCategory()(*ibf345d4df2f0f8a237eae842ca384dd517ada2627beba3c7d7f53baa0aac9cd3.DeviceCategoryRequestBuilder) {
    return ibf345d4df2f0f8a237eae842ca384dd517ada2627beba3c7d7f53baa0aac9cd3.NewDeviceCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DeviceCompliancePolicyStates()(*iedfe7deabe7e5f286b481b309b885e361a1fc70fcd9dd27c13cc31b607533a7d.DeviceCompliancePolicyStatesRequestBuilder) {
    return iedfe7deabe7e5f286b481b309b885e361a1fc70fcd9dd27c13cc31b607533a7d.NewDeviceCompliancePolicyStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicyStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.comanagedDevices.item.deviceCompliancePolicyStates.item collection
func (m *ManagedDeviceRequestBuilder) DeviceCompliancePolicyStatesById(id string)(*i1cf54773ea201b86cd3ff29dc00e0c4acb254a4bdf0fd5df1f70ef4013a36e72.DeviceCompliancePolicyStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyState_id"] = id
    }
    return i1cf54773ea201b86cd3ff29dc00e0c4acb254a4bdf0fd5df1f70ef4013a36e72.NewDeviceCompliancePolicyStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DeviceConfigurationStates()(*i310834334f4240786fcd6b5675b3ee8fbeb94f1c4b1b200e7e0b47c655de6a72.DeviceConfigurationStatesRequestBuilder) {
    return i310834334f4240786fcd6b5675b3ee8fbeb94f1c4b1b200e7e0b47c655de6a72.NewDeviceConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.comanagedDevices.item.deviceConfigurationStates.item collection
func (m *ManagedDeviceRequestBuilder) DeviceConfigurationStatesById(id string)(*i88c0ef2f8129d7022e7594e01e34ed8d304c73d763ba681efe40080e626ebd3e.DeviceConfigurationStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationState_id"] = id
    }
    return i88c0ef2f8129d7022e7594e01e34ed8d304c73d763ba681efe40080e626ebd3e.NewDeviceConfigurationStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Disable()(*idcab09b434ee10bc36b46bee67d1c1ee38cc5810295e2f26b4055875b7af37dc.DisableRequestBuilder) {
    return idcab09b434ee10bc36b46bee67d1c1ee38cc5810295e2f26b4055875b7af37dc.NewDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) DisableLostMode()(*i523f2fe01a487001fe905fdd45d3f80a080ff6e83855beac7dba457c171312fd.DisableLostModeRequestBuilder) {
    return i523f2fe01a487001fe905fdd45d3f80a080ff6e83855beac7dba457c171312fd.NewDisableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) EnableLostMode()(*i21902135131fa38cbfc951b49b1a03fc87545fa77bf484d3da70470837d99e8d.EnableLostModeRequestBuilder) {
    return i21902135131fa38cbfc951b49b1a03fc87545fa77bf484d3da70470837d99e8d.NewEnableLostModeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the list of co-managed devices report
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
// GetCloudPcRemoteActionResults builds and executes requests for operations under \deviceManagement\comanagedDevices\{managedDevice-id}\microsoft.graph.getCloudPcRemoteActionResults()
func (m *ManagedDeviceRequestBuilder) GetCloudPcRemoteActionResults()(*i5a2cb0524ba36f3452cc942efde3281af6e49bd918b07ec6d4553a95679f4244.GetCloudPcRemoteActionResultsRequestBuilder) {
    return i5a2cb0524ba36f3452cc942efde3281af6e49bd918b07ec6d4553a95679f4244.NewGetCloudPcRemoteActionResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFileVaultKey builds and executes requests for operations under \deviceManagement\comanagedDevices\{managedDevice-id}\microsoft.graph.getFileVaultKey()
func (m *ManagedDeviceRequestBuilder) GetFileVaultKey()(*i2a373154cbcecc73fa6269b089925214127c4b5c2d032275d9319e9099e011ba.GetFileVaultKeyRequestBuilder) {
    return i2a373154cbcecc73fa6269b089925214127c4b5c2d032275d9319e9099e011ba.NewGetFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetNonCompliantSettings builds and executes requests for operations under \deviceManagement\comanagedDevices\{managedDevice-id}\microsoft.graph.getNonCompliantSettings()
func (m *ManagedDeviceRequestBuilder) GetNonCompliantSettings()(*i1947898e3bbb95e8575558ada088e5fe3120b250ffef60fca81cfa321a720d05.GetNonCompliantSettingsRequestBuilder) {
    return i1947898e3bbb95e8575558ada088e5fe3120b250ffef60fca81cfa321a720d05.NewGetNonCompliantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOemWarranty builds and executes requests for operations under \deviceManagement\comanagedDevices\{managedDevice-id}\microsoft.graph.getOemWarranty()
func (m *ManagedDeviceRequestBuilder) GetOemWarranty()(*i7fe5a60c553129a9db80152bf36f67e1f4ee971a86dccf62668cdb5c4966077b.GetOemWarrantyRequestBuilder) {
    return i7fe5a60c553129a9db80152bf36f67e1f4ee971a86dccf62668cdb5c4966077b.NewGetOemWarrantyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LocateDevice()(*i474daedfd97083a2b70cfaaee5db2b7ac23cd62bfacff938adc2ec3734dca31a.LocateDeviceRequestBuilder) {
    return i474daedfd97083a2b70cfaaee5db2b7ac23cd62bfacff938adc2ec3734dca31a.NewLocateDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LogCollectionRequests()(*i5f3cf0b3cdbd4ccbd408a7bbcfb1134713761ba963b39a1a0e97c2ada49c6e64.LogCollectionRequestsRequestBuilder) {
    return i5f3cf0b3cdbd4ccbd408a7bbcfb1134713761ba963b39a1a0e97c2ada49c6e64.NewLogCollectionRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LogCollectionRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.comanagedDevices.item.logCollectionRequests.item collection
func (m *ManagedDeviceRequestBuilder) LogCollectionRequestsById(id string)(*i1bc1487940b05ce412b4ae41cbb75ff7066428095ef5956e8b49523af739ae12.DeviceLogCollectionResponseRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceLogCollectionResponse_id"] = id
    }
    return i1bc1487940b05ce412b4ae41cbb75ff7066428095ef5956e8b49523af739ae12.NewDeviceLogCollectionResponseRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) LogoutSharedAppleDeviceActiveUser()(*i6afbb8bdf1ff302cabb66d3e0a2c4212f688793e011fb7452c60d29077219ccc.LogoutSharedAppleDeviceActiveUserRequestBuilder) {
    return i6afbb8bdf1ff302cabb66d3e0a2c4212f688793e011fb7452c60d29077219ccc.NewLogoutSharedAppleDeviceActiveUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ManagedDeviceMobileAppConfigurationStates()(*i2a36033ee9d2b5fd75a9ff828a2a4949c7937585f41afbe202af5de47127dd1b.ManagedDeviceMobileAppConfigurationStatesRequestBuilder) {
    return i2a36033ee9d2b5fd75a9ff828a2a4949c7937585f41afbe202af5de47127dd1b.NewManagedDeviceMobileAppConfigurationStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceMobileAppConfigurationStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.comanagedDevices.item.managedDeviceMobileAppConfigurationStates.item collection
func (m *ManagedDeviceRequestBuilder) ManagedDeviceMobileAppConfigurationStatesById(id string)(*i3854f67f765bed0207af3dba7d579b83beb1633e1d2a7d9e98b9579399e553ae.ManagedDeviceMobileAppConfigurationStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationState_id"] = id
    }
    return i3854f67f765bed0207af3dba7d579b83beb1633e1d2a7d9e98b9579399e553ae.NewManagedDeviceMobileAppConfigurationStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) OverrideComplianceState()(*ife6aae505188bb4c2c63cb0aacab82413df394100290381167a1feb35be9f935.OverrideComplianceStateRequestBuilder) {
    return ife6aae505188bb4c2c63cb0aacab82413df394100290381167a1feb35be9f935.NewOverrideComplianceStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the list of co-managed devices report
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
func (m *ManagedDeviceRequestBuilder) PlayLostModeSound()(*iab5d4b97db65f3530dcf7249ef64d333c528c95523c7d2f233e65624d9c3aed6.PlayLostModeSoundRequestBuilder) {
    return iab5d4b97db65f3530dcf7249ef64d333c528c95523c7d2f233e65624d9c3aed6.NewPlayLostModeSoundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RebootNow()(*i3190268115ebe3af8fbbb7f7b8f73864d5604f943cdd020b77129dffe05358c6.RebootNowRequestBuilder) {
    return i3190268115ebe3af8fbbb7f7b8f73864d5604f943cdd020b77129dffe05358c6.NewRebootNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RecoverPasscode()(*id23da8ff63dabfc5e2f8ae52fc2c56079ba902db00a78ac2b11ded16d7c94017.RecoverPasscodeRequestBuilder) {
    return id23da8ff63dabfc5e2f8ae52fc2c56079ba902db00a78ac2b11ded16d7c94017.NewRecoverPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Reenable()(*ida779665108e67dcd8e766def39e06d84820ead54e3bd031f8b8be0d162618f9.ReenableRequestBuilder) {
    return ida779665108e67dcd8e766def39e06d84820ead54e3bd031f8b8be0d162618f9.NewReenableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RemoteLock()(*ifbbfe5c6464ce13a6e4499ee4f52892cd67d8abbfdced67cc6d8065e47e9249b.RemoteLockRequestBuilder) {
    return ifbbfe5c6464ce13a6e4499ee4f52892cd67d8abbfdced67cc6d8065e47e9249b.NewRemoteLockRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ReprovisionCloudPc()(*i5c1b61352546e0e98f55231ab68a18b6252bb9ae9d227682a4961bb808486269.ReprovisionCloudPcRequestBuilder) {
    return i5c1b61352546e0e98f55231ab68a18b6252bb9ae9d227682a4961bb808486269.NewReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RequestRemoteAssistance()(*ia9e35aa934485810b7f428c752d5b72ac6c7bb25b0203c31668c6ffeb9016989.RequestRemoteAssistanceRequestBuilder) {
    return ia9e35aa934485810b7f428c752d5b72ac6c7bb25b0203c31668c6ffeb9016989.NewRequestRemoteAssistanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ResetPasscode()(*i9da99e12e59b05474451fa26fb9ba2a61a1e7f3d2d11e16bb6107056d1bf1d18.ResetPasscodeRequestBuilder) {
    return i9da99e12e59b05474451fa26fb9ba2a61a1e7f3d2d11e16bb6107056d1bf1d18.NewResetPasscodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ResizeCloudPc()(*i6273d621d2071ff7812d4ccc27138d3c37af406b8289aba4d972bb96c54c8c65.ResizeCloudPcRequestBuilder) {
    return i6273d621d2071ff7812d4ccc27138d3c37af406b8289aba4d972bb96c54c8c65.NewResizeCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RestoreCloudPc()(*i2df51e7699d7e946c4c71f22c064f86f0b807928512258930f58dcfdb5db3cc6.RestoreCloudPcRequestBuilder) {
    return i2df51e7699d7e946c4c71f22c064f86f0b807928512258930f58dcfdb5db3cc6.NewRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Retire()(*ic814e3839bc6f855c24fd1a5a83ccb31ddd05002e9bd4af67e2782c60291bd1d.RetireRequestBuilder) {
    return ic814e3839bc6f855c24fd1a5a83ccb31ddd05002e9bd4af67e2782c60291bd1d.NewRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RevokeAppleVppLicenses()(*i06f050446fb7bd265e92527bb85f6f2a987efadd05b06c81151de522ed58e0ef.RevokeAppleVppLicensesRequestBuilder) {
    return i06f050446fb7bd265e92527bb85f6f2a987efadd05b06c81151de522ed58e0ef.NewRevokeAppleVppLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RotateBitLockerKeys()(*i31eb74f02efdd35067936e3939c2d3f35c2731c77e2ba3e2916be25ef567d1db.RotateBitLockerKeysRequestBuilder) {
    return i31eb74f02efdd35067936e3939c2d3f35c2731c77e2ba3e2916be25ef567d1db.NewRotateBitLockerKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) RotateFileVaultKey()(*iaab41b7756caa3ffe90d804bde7b8e128221af504e2142a1fe034770c8fe02f4.RotateFileVaultKeyRequestBuilder) {
    return iaab41b7756caa3ffe90d804bde7b8e128221af504e2142a1fe034770c8fe02f4.NewRotateFileVaultKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SecurityBaselineStates()(*i8c92d1f6b3158dce0d6c8344255235e2ea626a1c9b3ebd7ca2f5ab067d9e9e75.SecurityBaselineStatesRequestBuilder) {
    return i8c92d1f6b3158dce0d6c8344255235e2ea626a1c9b3ebd7ca2f5ab067d9e9e75.NewSecurityBaselineStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecurityBaselineStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.comanagedDevices.item.securityBaselineStates.item collection
func (m *ManagedDeviceRequestBuilder) SecurityBaselineStatesById(id string)(*i287d2bc1eaa2bb4787a75a13a96c399c86ed7cfa8a05bd662963aca0b4dc468f.SecurityBaselineStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["securityBaselineState_id"] = id
    }
    return i287d2bc1eaa2bb4787a75a13a96c399c86ed7cfa8a05bd662963aca0b4dc468f.NewSecurityBaselineStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SendCustomNotificationToCompanyPortal()(*i397317cd5c150ef73c8dbfe046b3a5a67e66696c131b64044418b32848f78b02.SendCustomNotificationToCompanyPortalRequestBuilder) {
    return i397317cd5c150ef73c8dbfe046b3a5a67e66696c131b64044418b32848f78b02.NewSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SetDeviceName()(*i69332727d899c5dfc77f80ae621a8729902a82541ca3df96c5050d6b45b50fe5.SetDeviceNameRequestBuilder) {
    return i69332727d899c5dfc77f80ae621a8729902a82541ca3df96c5050d6b45b50fe5.NewSetDeviceNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) ShutDown()(*i0265d37d4c192c128f26295070675911ea6710cb341155b618c000a508cfa91d.ShutDownRequestBuilder) {
    return i0265d37d4c192c128f26295070675911ea6710cb341155b618c000a508cfa91d.NewShutDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) SyncDevice()(*ia29b36d17801e439026505c34550ac7cee3a69c1ff4c5415c194f9ac625e44bf.SyncDeviceRequestBuilder) {
    return ia29b36d17801e439026505c34550ac7cee3a69c1ff4c5415c194f9ac625e44bf.NewSyncDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) TriggerConfigurationManagerAction()(*id475b3b19e63c8aea08d85fbab4ad351d8d0271fec24c7176711f0a0a8a45d62.TriggerConfigurationManagerActionRequestBuilder) {
    return id475b3b19e63c8aea08d85fbab4ad351d8d0271fec24c7176711f0a0a8a45d62.NewTriggerConfigurationManagerActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) UpdateWindowsDeviceAccount()(*iddee818f039172537f45ef3559b2d06b97b848792f16ac94552dcfd877edcdf4.UpdateWindowsDeviceAccountRequestBuilder) {
    return iddee818f039172537f45ef3559b2d06b97b848792f16ac94552dcfd877edcdf4.NewUpdateWindowsDeviceAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Users()(*i69683f121195c82c8db1f6f852405aa25762f58bd7b0efa9c356fb6260734ed3.UsersRequestBuilder) {
    return i69683f121195c82c8db1f6f852405aa25762f58bd7b0efa9c356fb6260734ed3.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderScan()(*i145cd047a0f177065361e7e02b132b0f8ce1e436d675e88c887b86e5f37651a9.WindowsDefenderScanRequestBuilder) {
    return i145cd047a0f177065361e7e02b132b0f8ce1e436d675e88c887b86e5f37651a9.NewWindowsDefenderScanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsDefenderUpdateSignatures()(*i86b6189da9f672cea7ec2267f0c2bc9f2ab30cc1cce295e84df98894cde4ad21.WindowsDefenderUpdateSignaturesRequestBuilder) {
    return i86b6189da9f672cea7ec2267f0c2bc9f2ab30cc1cce295e84df98894cde4ad21.NewWindowsDefenderUpdateSignaturesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) WindowsProtectionState()(*ib5f935e060b5bc7e8d0d5d9de0d928d9a660e962724c3cea1fd5cab5cff38ecc.WindowsProtectionStateRequestBuilder) {
    return ib5f935e060b5bc7e8d0d5d9de0d928d9a660e962724c3cea1fd5cab5cff38ecc.NewWindowsProtectionStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceRequestBuilder) Wipe()(*i3eeb66f94f27e28ae758717375265f4c4a2b99c3b3ebf1d81e6148870d288958.WipeRequestBuilder) {
    return i3eeb66f94f27e28ae758717375265f4c4a2b99c3b3ebf1d81e6148870d288958.NewWipeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
