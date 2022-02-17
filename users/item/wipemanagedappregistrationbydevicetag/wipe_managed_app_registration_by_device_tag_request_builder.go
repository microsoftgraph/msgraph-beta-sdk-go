package wipemanagedappregistrationbydevicetag

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// WipeManagedAppRegistrationByDeviceTagRequestBuilder builds and executes requests for operations under \users\{user-id}\microsoft.graph.wipeManagedAppRegistrationByDeviceTag
type WipeManagedAppRegistrationByDeviceTagRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WipeManagedAppRegistrationByDeviceTagRequestBuilderPostOptions options for Post
type WipeManagedAppRegistrationByDeviceTagRequestBuilderPostOptions struct {
    // 
    Body *WipeManagedAppRegistrationByDeviceTagRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal instantiates a new WipeManagedAppRegistrationByDeviceTagRequestBuilder and sets the default values.
func NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    m := &WipeManagedAppRegistrationByDeviceTagRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/microsoft.graph.wipeManagedAppRegistrationByDeviceTag";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWipeManagedAppRegistrationByDeviceTagRequestBuilder instantiates a new WipeManagedAppRegistrationByDeviceTagRequestBuilder and sets the default values.
func NewWipeManagedAppRegistrationByDeviceTagRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation issues a wipe operation on an app registration with specified device tag.
func (m *WipeManagedAppRegistrationByDeviceTagRequestBuilder) CreatePostRequestInformation(options *WipeManagedAppRegistrationByDeviceTagRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Post issues a wipe operation on an app registration with specified device tag.
func (m *WipeManagedAppRegistrationByDeviceTagRequestBuilder) Post(options *WipeManagedAppRegistrationByDeviceTagRequestBuilderPostOptions)(error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
