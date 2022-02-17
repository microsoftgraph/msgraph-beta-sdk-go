package windowsprotectionstate

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i7d799ec9df3d5c404eec113574d92a36c40d1176d5ba37bb9bbf6115c32cfd4a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/windowsprotectionstate/detectedmalwarestate"
    i7c28f247d1352ac8e9e8461bd9fb2ffbf0472cfcabfbfd3171c743b4b6ce4e64 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/item/windowsprotectionstate/detectedmalwarestate/item"
)

// WindowsProtectionStateRequestBuilder builds and executes requests for operations under \users\{user-id}\managedDevices\{managedDevice-id}\windowsProtectionState
type WindowsProtectionStateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WindowsProtectionStateRequestBuilderDeleteOptions options for Delete
type WindowsProtectionStateRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsProtectionStateRequestBuilderGetOptions options for Get
type WindowsProtectionStateRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WindowsProtectionStateRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsProtectionStateRequestBuilderGetQueryParameters the device protection status. This property is read-only.
type WindowsProtectionStateRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// WindowsProtectionStateRequestBuilderPatchOptions options for Patch
type WindowsProtectionStateRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsProtectionState;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewWindowsProtectionStateRequestBuilderInternal instantiates a new WindowsProtectionStateRequestBuilder and sets the default values.
func NewWindowsProtectionStateRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsProtectionStateRequestBuilder) {
    m := &WindowsProtectionStateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/managedDevices/{managedDevice_id}/windowsProtectionState{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsProtectionStateRequestBuilder instantiates a new WindowsProtectionStateRequestBuilder and sets the default values.
func NewWindowsProtectionStateRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsProtectionStateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsProtectionStateRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the device protection status. This property is read-only.
func (m *WindowsProtectionStateRequestBuilder) CreateDeleteRequestInformation(options *WindowsProtectionStateRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the device protection status. This property is read-only.
func (m *WindowsProtectionStateRequestBuilder) CreateGetRequestInformation(options *WindowsProtectionStateRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the device protection status. This property is read-only.
func (m *WindowsProtectionStateRequestBuilder) CreatePatchRequestInformation(options *WindowsProtectionStateRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the device protection status. This property is read-only.
func (m *WindowsProtectionStateRequestBuilder) Delete(options *WindowsProtectionStateRequestBuilderDeleteOptions)(error) {
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
func (m *WindowsProtectionStateRequestBuilder) DetectedMalwareState()(*i7d799ec9df3d5c404eec113574d92a36c40d1176d5ba37bb9bbf6115c32cfd4a.DetectedMalwareStateRequestBuilder) {
    return i7d799ec9df3d5c404eec113574d92a36c40d1176d5ba37bb9bbf6115c32cfd4a.NewDetectedMalwareStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DetectedMalwareStateById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.managedDevices.item.windowsProtectionState.detectedMalwareState.item collection
func (m *WindowsProtectionStateRequestBuilder) DetectedMalwareStateById(id string)(*i7c28f247d1352ac8e9e8461bd9fb2ffbf0472cfcabfbfd3171c743b4b6ce4e64.WindowsDeviceMalwareStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsDeviceMalwareState_id"] = id
    }
    return i7c28f247d1352ac8e9e8461bd9fb2ffbf0472cfcabfbfd3171c743b4b6ce4e64.NewWindowsDeviceMalwareStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the device protection status. This property is read-only.
func (m *WindowsProtectionStateRequestBuilder) Get(options *WindowsProtectionStateRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsProtectionState, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWindowsProtectionState() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsProtectionState), nil
}
// Patch the device protection status. This property is read-only.
func (m *WindowsProtectionStateRequestBuilder) Patch(options *WindowsProtectionStateRequestBuilderPatchOptions)(error) {
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
