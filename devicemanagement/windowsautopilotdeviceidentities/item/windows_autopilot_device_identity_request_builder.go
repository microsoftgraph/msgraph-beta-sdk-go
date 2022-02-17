package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0603021b313096a6726ff00fcc4dfd8a8b9c2fc89775400414461f4d36fd7d04 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/assignresourceaccounttodevice"
    i08637003c24136cfb79d3be758b142dd1115777f5410f7f104df5f6983cc4367 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/deploymentprofile"
    i3f74d7cb02f41203c9cd7d155fac63204acbe2e5d81f2402d7b639e0c9a7b5f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/assignusertodevice"
    i7336d162e1b3abcfc9696b9aee7f016a3c7fa2dbb98e24cf3154a91e3365d661 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/unassignresourceaccountfromdevice"
    i83461ac3f25e4e8a3194b7af8d988758e07ed71d4aa5b3f61ded8939dc620c98 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/updatedeviceproperties"
    id6142e63ae5bd7453e3630d6c28f631a110e460306d06d64699f247d9bb218f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/intendeddeploymentprofile"
    iea2a6744b4e629234021d6f6c8e9fa076a3ab8a7876d7c57abc69043cf535a5b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/unassignuserfromdevice"
)

// WindowsAutopilotDeviceIdentityRequestBuilder builds and executes requests for operations under \deviceManagement\windowsAutopilotDeviceIdentities\{windowsAutopilotDeviceIdentity-id}
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
// WindowsAutopilotDeviceIdentityRequestBuilderGetQueryParameters the Windows autopilot device identities contained collection.
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
func (m *WindowsAutopilotDeviceIdentityRequestBuilder) AssignResourceAccountToDevice()(*i0603021b313096a6726ff00fcc4dfd8a8b9c2fc89775400414461f4d36fd7d04.AssignResourceAccountToDeviceRequestBuilder) {
    return i0603021b313096a6726ff00fcc4dfd8a8b9c2fc89775400414461f4d36fd7d04.NewAssignResourceAccountToDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsAutopilotDeviceIdentityRequestBuilder) AssignUserToDevice()(*i3f74d7cb02f41203c9cd7d155fac63204acbe2e5d81f2402d7b639e0c9a7b5f6.AssignUserToDeviceRequestBuilder) {
    return i3f74d7cb02f41203c9cd7d155fac63204acbe2e5d81f2402d7b639e0c9a7b5f6.NewAssignUserToDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewWindowsAutopilotDeviceIdentityRequestBuilderInternal instantiates a new WindowsAutopilotDeviceIdentityRequestBuilder and sets the default values.
func NewWindowsAutopilotDeviceIdentityRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsAutopilotDeviceIdentityRequestBuilder) {
    m := &WindowsAutopilotDeviceIdentityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/windowsAutopilotDeviceIdentities/{windowsAutopilotDeviceIdentity_id}{?select,expand}";
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
// CreateDeleteRequestInformation the Windows autopilot device identities contained collection.
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
// CreateGetRequestInformation the Windows autopilot device identities contained collection.
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
// CreatePatchRequestInformation the Windows autopilot device identities contained collection.
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
// Delete the Windows autopilot device identities contained collection.
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
func (m *WindowsAutopilotDeviceIdentityRequestBuilder) DeploymentProfile()(*i08637003c24136cfb79d3be758b142dd1115777f5410f7f104df5f6983cc4367.DeploymentProfileRequestBuilder) {
    return i08637003c24136cfb79d3be758b142dd1115777f5410f7f104df5f6983cc4367.NewDeploymentProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the Windows autopilot device identities contained collection.
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
func (m *WindowsAutopilotDeviceIdentityRequestBuilder) IntendedDeploymentProfile()(*id6142e63ae5bd7453e3630d6c28f631a110e460306d06d64699f247d9bb218f3.IntendedDeploymentProfileRequestBuilder) {
    return id6142e63ae5bd7453e3630d6c28f631a110e460306d06d64699f247d9bb218f3.NewIntendedDeploymentProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the Windows autopilot device identities contained collection.
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
func (m *WindowsAutopilotDeviceIdentityRequestBuilder) UnassignResourceAccountFromDevice()(*i7336d162e1b3abcfc9696b9aee7f016a3c7fa2dbb98e24cf3154a91e3365d661.UnassignResourceAccountFromDeviceRequestBuilder) {
    return i7336d162e1b3abcfc9696b9aee7f016a3c7fa2dbb98e24cf3154a91e3365d661.NewUnassignResourceAccountFromDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsAutopilotDeviceIdentityRequestBuilder) UnassignUserFromDevice()(*iea2a6744b4e629234021d6f6c8e9fa076a3ab8a7876d7c57abc69043cf535a5b.UnassignUserFromDeviceRequestBuilder) {
    return iea2a6744b4e629234021d6f6c8e9fa076a3ab8a7876d7c57abc69043cf535a5b.NewUnassignUserFromDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsAutopilotDeviceIdentityRequestBuilder) UpdateDeviceProperties()(*i83461ac3f25e4e8a3194b7af8d988758e07ed71d4aa5b3f61ded8939dc620c98.UpdateDevicePropertiesRequestBuilder) {
    return i83461ac3f25e4e8a3194b7af8d988758e07ed71d4aa5b3f61ded8939dc620c98.NewUpdateDevicePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
