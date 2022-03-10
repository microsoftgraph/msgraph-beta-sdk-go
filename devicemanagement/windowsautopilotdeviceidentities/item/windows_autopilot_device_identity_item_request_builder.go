package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0603021b313096a6726ff00fcc4dfd8a8b9c2fc89775400414461f4d36fd7d04 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/assignresourceaccounttodevice"
    i08637003c24136cfb79d3be758b142dd1115777f5410f7f104df5f6983cc4367 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/deploymentprofile"
    i3f74d7cb02f41203c9cd7d155fac63204acbe2e5d81f2402d7b639e0c9a7b5f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/assignusertodevice"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i7336d162e1b3abcfc9696b9aee7f016a3c7fa2dbb98e24cf3154a91e3365d661 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/unassignresourceaccountfromdevice"
    i83461ac3f25e4e8a3194b7af8d988758e07ed71d4aa5b3f61ded8939dc620c98 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/updatedeviceproperties"
    id6142e63ae5bd7453e3630d6c28f631a110e460306d06d64699f247d9bb218f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/intendeddeploymentprofile"
    iea2a6744b4e629234021d6f6c8e9fa076a3ab8a7876d7c57abc69043cf535a5b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/unassignuserfromdevice"
)

// WindowsAutopilotDeviceIdentityItemRequestBuilder provides operations to manage the windowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
type WindowsAutopilotDeviceIdentityItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WindowsAutopilotDeviceIdentityItemRequestBuilderDeleteOptions options for Delete
type WindowsAutopilotDeviceIdentityItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsAutopilotDeviceIdentityItemRequestBuilderGetOptions options for Get
type WindowsAutopilotDeviceIdentityItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters the Windows autopilot device identities contained collection.
type WindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// WindowsAutopilotDeviceIdentityItemRequestBuilderPatchOptions options for Patch
type WindowsAutopilotDeviceIdentityItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsAutopilotDeviceIdentityable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) AssignResourceAccountToDevice()(*i0603021b313096a6726ff00fcc4dfd8a8b9c2fc89775400414461f4d36fd7d04.AssignResourceAccountToDeviceRequestBuilder) {
    return i0603021b313096a6726ff00fcc4dfd8a8b9c2fc89775400414461f4d36fd7d04.NewAssignResourceAccountToDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) AssignUserToDevice()(*i3f74d7cb02f41203c9cd7d155fac63204acbe2e5d81f2402d7b639e0c9a7b5f6.AssignUserToDeviceRequestBuilder) {
    return i3f74d7cb02f41203c9cd7d155fac63204acbe2e5d81f2402d7b639e0c9a7b5f6.NewAssignUserToDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewWindowsAutopilotDeviceIdentityItemRequestBuilderInternal instantiates a new WindowsAutopilotDeviceIdentityItemRequestBuilder and sets the default values.
func NewWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsAutopilotDeviceIdentityItemRequestBuilder) {
    m := &WindowsAutopilotDeviceIdentityItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/windowsAutopilotDeviceIdentities/{windowsAutopilotDeviceIdentity_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsAutopilotDeviceIdentityItemRequestBuilder instantiates a new WindowsAutopilotDeviceIdentityItemRequestBuilder and sets the default values.
func NewWindowsAutopilotDeviceIdentityItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsAutopilotDeviceIdentityItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property windowsAutopilotDeviceIdentities for deviceManagement
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) CreateDeleteRequestInformation(options *WindowsAutopilotDeviceIdentityItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) CreateGetRequestInformation(options *WindowsAutopilotDeviceIdentityItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property windowsAutopilotDeviceIdentities in deviceManagement
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) CreatePatchRequestInformation(options *WindowsAutopilotDeviceIdentityItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property windowsAutopilotDeviceIdentities for deviceManagement
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) Delete(options *WindowsAutopilotDeviceIdentityItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) DeploymentProfile()(*i08637003c24136cfb79d3be758b142dd1115777f5410f7f104df5f6983cc4367.DeploymentProfileRequestBuilder) {
    return i08637003c24136cfb79d3be758b142dd1115777f5410f7f104df5f6983cc4367.NewDeploymentProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the Windows autopilot device identities contained collection.
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) Get(options *WindowsAutopilotDeviceIdentityItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsAutopilotDeviceIdentityable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateWindowsAutopilotDeviceIdentityFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsAutopilotDeviceIdentityable), nil
}
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) IntendedDeploymentProfile()(*id6142e63ae5bd7453e3630d6c28f631a110e460306d06d64699f247d9bb218f3.IntendedDeploymentProfileRequestBuilder) {
    return id6142e63ae5bd7453e3630d6c28f631a110e460306d06d64699f247d9bb218f3.NewIntendedDeploymentProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property windowsAutopilotDeviceIdentities in deviceManagement
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) Patch(options *WindowsAutopilotDeviceIdentityItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) UnassignResourceAccountFromDevice()(*i7336d162e1b3abcfc9696b9aee7f016a3c7fa2dbb98e24cf3154a91e3365d661.UnassignResourceAccountFromDeviceRequestBuilder) {
    return i7336d162e1b3abcfc9696b9aee7f016a3c7fa2dbb98e24cf3154a91e3365d661.NewUnassignResourceAccountFromDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) UnassignUserFromDevice()(*iea2a6744b4e629234021d6f6c8e9fa076a3ab8a7876d7c57abc69043cf535a5b.UnassignUserFromDeviceRequestBuilder) {
    return iea2a6744b4e629234021d6f6c8e9fa076a3ab8a7876d7c57abc69043cf535a5b.NewUnassignUserFromDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsAutopilotDeviceIdentityItemRequestBuilder) UpdateDeviceProperties()(*i83461ac3f25e4e8a3194b7af8d988758e07ed71d4aa5b3f61ded8939dc620c98.UpdateDevicePropertiesRequestBuilder) {
    return i83461ac3f25e4e8a3194b7af8d988758e07ed71d4aa5b3f61ded8939dc620c98.NewUpdateDevicePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
