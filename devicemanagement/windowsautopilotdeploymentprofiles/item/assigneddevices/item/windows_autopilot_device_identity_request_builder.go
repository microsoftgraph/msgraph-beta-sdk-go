package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i158b2f1ed97ea7c14fab9f68b26bf476d3708afeb068765960eaaf699af1b5f8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeploymentprofiles/item/assigneddevices/item/updatedeviceproperties"
    i6305748a7c20ec648805e3bbbb3193245a7f1126da5d98e5b9b3c62a480a077f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeploymentprofiles/item/assigneddevices/item/assignusertodevice"
    i6d6560c766eb2ce774731049fc529d3b519832e9f2f53682fbfe7d35df7c1c3c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeploymentprofiles/item/assigneddevices/item/assignresourceaccounttodevice"
    i83c5f499364f220f25e381f8075396532f9c977c78fde505fd17f38a1411ece0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeploymentprofiles/item/assigneddevices/item/unassignuserfromdevice"
    ib73822b720982991e43913cc2d8c66067380da4aa5c84e88e3f2465bcd4f0a23 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeploymentprofiles/item/assigneddevices/item/deploymentprofile"
    ica77650b1b796d69d225beec92c6b9e45081ed70aef73755c178ffbf8b628d4f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeploymentprofiles/item/assigneddevices/item/unassignresourceaccountfromdevice"
    ifaa60ce9423a2f2f03b8b75f54d288b9a3382a7e5693b4c28468c616d6306ded "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeploymentprofiles/item/assigneddevices/item/intendeddeploymentprofile"
)

// WindowsAutopilotDeviceIdentityRequestBuilder builds and executes requests for operations under \deviceManagement\windowsAutopilotDeploymentProfiles\{windowsAutopilotDeploymentProfile-id}\assignedDevices\{windowsAutopilotDeviceIdentity-id}
type WindowsAutopilotDeviceIdentityRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WindowsAutopilotDeviceIdentityRequestBuilderDeleteOptions options for Delete
type WindowsAutopilotDeviceIdentityRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsAutopilotDeviceIdentityRequestBuilderGetOptions options for Get
type WindowsAutopilotDeviceIdentityRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WindowsAutopilotDeviceIdentityRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsAutopilotDeviceIdentityRequestBuilderGetQueryParameters the list of assigned devices for the profile.
type WindowsAutopilotDeviceIdentityRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// WindowsAutopilotDeviceIdentityRequestBuilderPatchOptions options for Patch
type WindowsAutopilotDeviceIdentityRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsAutopilotDeviceIdentity;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *WindowsAutopilotDeviceIdentityRequestBuilder) AssignResourceAccountToDevice()(*i6d6560c766eb2ce774731049fc529d3b519832e9f2f53682fbfe7d35df7c1c3c.AssignResourceAccountToDeviceRequestBuilder) {
    return i6d6560c766eb2ce774731049fc529d3b519832e9f2f53682fbfe7d35df7c1c3c.NewAssignResourceAccountToDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsAutopilotDeviceIdentityRequestBuilder) AssignUserToDevice()(*i6305748a7c20ec648805e3bbbb3193245a7f1126da5d98e5b9b3c62a480a077f.AssignUserToDeviceRequestBuilder) {
    return i6305748a7c20ec648805e3bbbb3193245a7f1126da5d98e5b9b3c62a480a077f.NewAssignUserToDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewWindowsAutopilotDeviceIdentityRequestBuilderInternal instantiates a new WindowsAutopilotDeviceIdentityRequestBuilder and sets the default values.
func NewWindowsAutopilotDeviceIdentityRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsAutopilotDeviceIdentityRequestBuilder) {
    m := &WindowsAutopilotDeviceIdentityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/windowsAutopilotDeploymentProfiles/{windowsAutopilotDeploymentProfile_id}/assignedDevices/{windowsAutopilotDeviceIdentity_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsAutopilotDeviceIdentityRequestBuilder instantiates a new WindowsAutopilotDeviceIdentityRequestBuilder and sets the default values.
func NewWindowsAutopilotDeviceIdentityRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsAutopilotDeviceIdentityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsAutopilotDeviceIdentityRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the list of assigned devices for the profile.
func (m *WindowsAutopilotDeviceIdentityRequestBuilder) CreateDeleteRequestInformation(options *WindowsAutopilotDeviceIdentityRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the list of assigned devices for the profile.
func (m *WindowsAutopilotDeviceIdentityRequestBuilder) CreateGetRequestInformation(options *WindowsAutopilotDeviceIdentityRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the list of assigned devices for the profile.
func (m *WindowsAutopilotDeviceIdentityRequestBuilder) CreatePatchRequestInformation(options *WindowsAutopilotDeviceIdentityRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the list of assigned devices for the profile.
func (m *WindowsAutopilotDeviceIdentityRequestBuilder) Delete(options *WindowsAutopilotDeviceIdentityRequestBuilderDeleteOptions)(error) {
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
func (m *WindowsAutopilotDeviceIdentityRequestBuilder) DeploymentProfile()(*ib73822b720982991e43913cc2d8c66067380da4aa5c84e88e3f2465bcd4f0a23.DeploymentProfileRequestBuilder) {
    return ib73822b720982991e43913cc2d8c66067380da4aa5c84e88e3f2465bcd4f0a23.NewDeploymentProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the list of assigned devices for the profile.
func (m *WindowsAutopilotDeviceIdentityRequestBuilder) Get(options *WindowsAutopilotDeviceIdentityRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsAutopilotDeviceIdentity, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWindowsAutopilotDeviceIdentity() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsAutopilotDeviceIdentity), nil
}
func (m *WindowsAutopilotDeviceIdentityRequestBuilder) IntendedDeploymentProfile()(*ifaa60ce9423a2f2f03b8b75f54d288b9a3382a7e5693b4c28468c616d6306ded.IntendedDeploymentProfileRequestBuilder) {
    return ifaa60ce9423a2f2f03b8b75f54d288b9a3382a7e5693b4c28468c616d6306ded.NewIntendedDeploymentProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the list of assigned devices for the profile.
func (m *WindowsAutopilotDeviceIdentityRequestBuilder) Patch(options *WindowsAutopilotDeviceIdentityRequestBuilderPatchOptions)(error) {
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
func (m *WindowsAutopilotDeviceIdentityRequestBuilder) UnassignResourceAccountFromDevice()(*ica77650b1b796d69d225beec92c6b9e45081ed70aef73755c178ffbf8b628d4f.UnassignResourceAccountFromDeviceRequestBuilder) {
    return ica77650b1b796d69d225beec92c6b9e45081ed70aef73755c178ffbf8b628d4f.NewUnassignResourceAccountFromDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsAutopilotDeviceIdentityRequestBuilder) UnassignUserFromDevice()(*i83c5f499364f220f25e381f8075396532f9c977c78fde505fd17f38a1411ece0.UnassignUserFromDeviceRequestBuilder) {
    return i83c5f499364f220f25e381f8075396532f9c977c78fde505fd17f38a1411ece0.NewUnassignUserFromDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsAutopilotDeviceIdentityRequestBuilder) UpdateDeviceProperties()(*i158b2f1ed97ea7c14fab9f68b26bf476d3708afeb068765960eaaf699af1b5f8.UpdateDevicePropertiesRequestBuilder) {
    return i158b2f1ed97ea7c14fab9f68b26bf476d3708afeb068765960eaaf699af1b5f8.NewUpdateDevicePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
