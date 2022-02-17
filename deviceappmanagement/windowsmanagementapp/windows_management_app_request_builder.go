package windowsmanagementapp

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i6c384b2192ee3a8f21fc65abdf0cc687b856853682a74de978836a231eeaa218 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/windowsmanagementapp/setasmanagedinstaller"
    ie68815be795e5c46a6b775db4a53c48c75956027c4be75e1ee5c82e410c26bfd "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/windowsmanagementapp/ref"
)

// WindowsManagementAppRequestBuilder builds and executes requests for operations under \deviceAppManagement\windowsManagementApp
type WindowsManagementAppRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WindowsManagementAppRequestBuilderGetOptions options for Get
type WindowsManagementAppRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WindowsManagementAppRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsManagementAppRequestBuilderGetQueryParameters windows management app.
type WindowsManagementAppRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewWindowsManagementAppRequestBuilderInternal instantiates a new WindowsManagementAppRequestBuilder and sets the default values.
func NewWindowsManagementAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsManagementAppRequestBuilder) {
    m := &WindowsManagementAppRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/windowsManagementApp{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsManagementAppRequestBuilder instantiates a new WindowsManagementAppRequestBuilder and sets the default values.
func NewWindowsManagementAppRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsManagementAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsManagementAppRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation windows management app.
func (m *WindowsManagementAppRequestBuilder) CreateGetRequestInformation(options *WindowsManagementAppRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get windows management app.
func (m *WindowsManagementAppRequestBuilder) Get(options *WindowsManagementAppRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsManagementApp, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWindowsManagementApp() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsManagementApp), nil
}
func (m *WindowsManagementAppRequestBuilder) Ref()(*ie68815be795e5c46a6b775db4a53c48c75956027c4be75e1ee5c82e410c26bfd.RefRequestBuilder) {
    return ie68815be795e5c46a6b775db4a53c48c75956027c4be75e1ee5c82e410c26bfd.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsManagementAppRequestBuilder) SetAsManagedInstaller()(*i6c384b2192ee3a8f21fc65abdf0cc687b856853682a74de978836a231eeaa218.SetAsManagedInstallerRequestBuilder) {
    return i6c384b2192ee3a8f21fc65abdf0cc687b856853682a74de978836a231eeaa218.NewSetAsManagedInstallerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
