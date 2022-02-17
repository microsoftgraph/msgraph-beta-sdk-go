package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3183896baf94d618a94be9010348fed36cad7fe793350064ebf7c3d92813417c "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedapppolicies/item/targetedmanagedappprotection"
    i5a94562070daac291339d688f48e02dc86f2c6f7ff6090f260af1284334d30d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedapppolicies/item/targetapps"
    i87ae477f7833a79caa4e8069280d2e383547bd57a2317fb29a659611abf5837f "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedapppolicies/item/managedappprotection"
    idee3b8f515a95822717488ac94a52d8e67cb6ffe59cd110e93b47bf85402c06b "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedapppolicies/item/windowsinformationprotection"
)

// ManagedAppPolicyRequestBuilder builds and executes requests for operations under \deviceAppManagement\managedAppPolicies\{managedAppPolicy-id}
type ManagedAppPolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagedAppPolicyRequestBuilderDeleteOptions options for Delete
type ManagedAppPolicyRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedAppPolicyRequestBuilderGetOptions options for Get
type ManagedAppPolicyRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagedAppPolicyRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedAppPolicyRequestBuilderGetQueryParameters managed app policies.
type ManagedAppPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ManagedAppPolicyRequestBuilderPatchOptions options for Patch
type ManagedAppPolicyRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedAppPolicy;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewManagedAppPolicyRequestBuilderInternal instantiates a new ManagedAppPolicyRequestBuilder and sets the default values.
func NewManagedAppPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppPolicyRequestBuilder) {
    m := &ManagedAppPolicyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedAppPolicies/{managedAppPolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedAppPolicyRequestBuilder instantiates a new ManagedAppPolicyRequestBuilder and sets the default values.
func NewManagedAppPolicyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedAppPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation managed app policies.
func (m *ManagedAppPolicyRequestBuilder) CreateDeleteRequestInformation(options *ManagedAppPolicyRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation managed app policies.
func (m *ManagedAppPolicyRequestBuilder) CreateGetRequestInformation(options *ManagedAppPolicyRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation managed app policies.
func (m *ManagedAppPolicyRequestBuilder) CreatePatchRequestInformation(options *ManagedAppPolicyRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete managed app policies.
func (m *ManagedAppPolicyRequestBuilder) Delete(options *ManagedAppPolicyRequestBuilderDeleteOptions)(error) {
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
// Get managed app policies.
func (m *ManagedAppPolicyRequestBuilder) Get(options *ManagedAppPolicyRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedAppPolicy, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewManagedAppPolicy() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedAppPolicy), nil
}
func (m *ManagedAppPolicyRequestBuilder) ManagedAppProtection()(*i87ae477f7833a79caa4e8069280d2e383547bd57a2317fb29a659611abf5837f.ManagedAppProtectionRequestBuilder) {
    return i87ae477f7833a79caa4e8069280d2e383547bd57a2317fb29a659611abf5837f.NewManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch managed app policies.
func (m *ManagedAppPolicyRequestBuilder) Patch(options *ManagedAppPolicyRequestBuilderPatchOptions)(error) {
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
func (m *ManagedAppPolicyRequestBuilder) TargetApps()(*i5a94562070daac291339d688f48e02dc86f2c6f7ff6090f260af1284334d30d8.TargetAppsRequestBuilder) {
    return i5a94562070daac291339d688f48e02dc86f2c6f7ff6090f260af1284334d30d8.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedAppPolicyRequestBuilder) TargetedManagedAppProtection()(*i3183896baf94d618a94be9010348fed36cad7fe793350064ebf7c3d92813417c.TargetedManagedAppProtectionRequestBuilder) {
    return i3183896baf94d618a94be9010348fed36cad7fe793350064ebf7c3d92813417c.NewTargetedManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedAppPolicyRequestBuilder) WindowsInformationProtection()(*idee3b8f515a95822717488ac94a52d8e67cb6ffe59cd110e93b47bf85402c06b.WindowsInformationProtectionRequestBuilder) {
    return idee3b8f515a95822717488ac94a52d8e67cb6ffe59cd110e93b47bf85402c06b.NewWindowsInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
