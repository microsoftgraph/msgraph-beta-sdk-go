package windowsinformationprotectiondeviceregistrations

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i982a2d08d81da2fa5b974b1400dc1c73b6ed0c5bffe779e528d8c9828a8e22ad "github.com/microsoftgraph/msgraph-beta-sdk-go/me/windowsinformationprotectiondeviceregistrations/ref"
)

// WindowsInformationProtectionDeviceRegistrationsRequestBuilder builds and executes requests for operations under \me\windowsInformationProtectionDeviceRegistrations
type WindowsInformationProtectionDeviceRegistrationsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WindowsInformationProtectionDeviceRegistrationsRequestBuilderGetOptions options for Get
type WindowsInformationProtectionDeviceRegistrationsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WindowsInformationProtectionDeviceRegistrationsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsInformationProtectionDeviceRegistrationsRequestBuilderGetQueryParameters zero or more WIP device registrations that belong to the user.
type WindowsInformationProtectionDeviceRegistrationsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// NewWindowsInformationProtectionDeviceRegistrationsRequestBuilderInternal instantiates a new WindowsInformationProtectionDeviceRegistrationsRequestBuilder and sets the default values.
func NewWindowsInformationProtectionDeviceRegistrationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsInformationProtectionDeviceRegistrationsRequestBuilder) {
    m := &WindowsInformationProtectionDeviceRegistrationsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/windowsInformationProtectionDeviceRegistrations{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsInformationProtectionDeviceRegistrationsRequestBuilder instantiates a new WindowsInformationProtectionDeviceRegistrationsRequestBuilder and sets the default values.
func NewWindowsInformationProtectionDeviceRegistrationsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsInformationProtectionDeviceRegistrationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsInformationProtectionDeviceRegistrationsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation zero or more WIP device registrations that belong to the user.
func (m *WindowsInformationProtectionDeviceRegistrationsRequestBuilder) CreateGetRequestInformation(options *WindowsInformationProtectionDeviceRegistrationsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get zero or more WIP device registrations that belong to the user.
func (m *WindowsInformationProtectionDeviceRegistrationsRequestBuilder) Get(options *WindowsInformationProtectionDeviceRegistrationsRequestBuilderGetOptions)(*WindowsInformationProtectionDeviceRegistrationsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionDeviceRegistrationsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*WindowsInformationProtectionDeviceRegistrationsResponse), nil
}
func (m *WindowsInformationProtectionDeviceRegistrationsRequestBuilder) Ref()(*i982a2d08d81da2fa5b974b1400dc1c73b6ed0c5bffe779e528d8c9828a8e22ad.RefRequestBuilder) {
    return i982a2d08d81da2fa5b974b1400dc1c73b6ed0c5bffe779e528d8c9828a8e22ad.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
