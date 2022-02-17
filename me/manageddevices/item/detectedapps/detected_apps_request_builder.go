package detectedapps

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ied1996a8deb71ecff4f79904335c95fc01694a40de1d003102afbae5071685dd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item/detectedapps/ref"
)

// DetectedAppsRequestBuilder builds and executes requests for operations under \me\managedDevices\{managedDevice-id}\detectedApps
type DetectedAppsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DetectedAppsRequestBuilderGetOptions options for Get
type DetectedAppsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DetectedAppsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DetectedAppsRequestBuilderGetQueryParameters all applications currently installed on the device
type DetectedAppsRequestBuilderGetQueryParameters struct {
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
// NewDetectedAppsRequestBuilderInternal instantiates a new DetectedAppsRequestBuilder and sets the default values.
func NewDetectedAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DetectedAppsRequestBuilder) {
    m := &DetectedAppsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/managedDevices/{managedDevice_id}/detectedApps{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDetectedAppsRequestBuilder instantiates a new DetectedAppsRequestBuilder and sets the default values.
func NewDetectedAppsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DetectedAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDetectedAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation all applications currently installed on the device
func (m *DetectedAppsRequestBuilder) CreateGetRequestInformation(options *DetectedAppsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get all applications currently installed on the device
func (m *DetectedAppsRequestBuilder) Get(options *DetectedAppsRequestBuilderGetOptions)(*DetectedAppsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDetectedAppsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*DetectedAppsResponse), nil
}
func (m *DetectedAppsRequestBuilder) Ref()(*ied1996a8deb71ecff4f79904335c95fc01694a40de1d003102afbae5071685dd.RefRequestBuilder) {
    return ied1996a8deb71ecff4f79904335c95fc01694a40de1d003102afbae5071685dd.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
