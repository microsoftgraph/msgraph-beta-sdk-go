package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3e31f706fc76ba4e16c152454d85c280f8ac53d7c171bf7d0e1e53bb8e2eca1c "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/wdacsupplementalpolicies/item/devicestatuses/item/policy"
)

// WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder builds and executes requests for operations under \deviceAppManagement\wdacSupplementalPolicies\{windowsDefenderApplicationControlSupplementalPolicy-id}\deviceStatuses\{windowsDefenderApplicationControlSupplementalPolicyDeploymentStatus-id}
type WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilderDeleteOptions options for Delete
type WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilderGetOptions options for Get
type WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilderGetQueryParameters the list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
type WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilderPatchOptions options for Patch
type WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilderInternal instantiates a new WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder and sets the default values.
func NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder) {
    m := &WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/wdacSupplementalPolicies/{windowsDefenderApplicationControlSupplementalPolicy_id}/deviceStatuses/{windowsDefenderApplicationControlSupplementalPolicyDeploymentStatus_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder instantiates a new WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder and sets the default values.
func NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder) CreateDeleteRequestInformation(options *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder) CreateGetRequestInformation(options *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder) CreatePatchRequestInformation(options *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder) Delete(options *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilderDeleteOptions)(error) {
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
// Get the list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder) Get(options *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus), nil
}
// Patch the list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder) Patch(options *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilderPatchOptions)(error) {
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
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder) Policy()(*i3e31f706fc76ba4e16c152454d85c280f8ac53d7c171bf7d0e1e53bb8e2eca1c.PolicyRequestBuilder) {
    return i3e31f706fc76ba4e16c152454d85c280f8ac53d7c171bf7d0e1e53bb8e2eca1c.NewPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
