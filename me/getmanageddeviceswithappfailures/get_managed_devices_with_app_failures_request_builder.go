package getmanageddeviceswithappfailures

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetManagedDevicesWithAppFailuresRequestBuilder builds and executes requests for operations under \me\microsoft.graph.getManagedDevicesWithAppFailures()
type GetManagedDevicesWithAppFailuresRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetManagedDevicesWithAppFailuresRequestBuilderGetOptions options for Get
type GetManagedDevicesWithAppFailuresRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetManagedDevicesWithAppFailuresRequestBuilderInternal instantiates a new GetManagedDevicesWithAppFailuresRequestBuilder and sets the default values.
func NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetManagedDevicesWithAppFailuresRequestBuilder) {
    m := &GetManagedDevicesWithAppFailuresRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/microsoft.graph.getManagedDevicesWithAppFailures()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetManagedDevicesWithAppFailuresRequestBuilder instantiates a new GetManagedDevicesWithAppFailuresRequestBuilder and sets the default values.
func NewGetManagedDevicesWithAppFailuresRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetManagedDevicesWithAppFailuresRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation retrieves the list of devices with failed apps
func (m *GetManagedDevicesWithAppFailuresRequestBuilder) CreateGetRequestInformation(options *GetManagedDevicesWithAppFailuresRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
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
// Get retrieves the list of devices with failed apps
func (m *GetManagedDevicesWithAppFailuresRequestBuilder) Get(options *GetManagedDevicesWithAppFailuresRequestBuilderGetOptions)([]string, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitiveCollectionAsync(*requestInfo, "string", nil, nil)
    if err != nil {
        return nil, err
    }
    val := make([]string, len(res))
    for i, v := range res {
        val[i] = *(v.(*string))
    }
    return val, nil
}
