package setdevicename

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// SetDeviceNameRequestBuilder builds and executes requests for operations under \deviceManagement\deviceCustomAttributeShellScripts\{deviceCustomAttributeShellScript-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.setDeviceName
type SetDeviceNameRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SetDeviceNameRequestBuilderPostOptions options for Post
type SetDeviceNameRequestBuilderPostOptions struct {
    // 
    Body *SetDeviceNameRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSetDeviceNameRequestBuilderInternal instantiates a new SetDeviceNameRequestBuilder and sets the default values.
func NewSetDeviceNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetDeviceNameRequestBuilder) {
    m := &SetDeviceNameRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCustomAttributeShellScripts/{deviceCustomAttributeShellScript_id}/deviceRunStates/{deviceManagementScriptDeviceState_id}/managedDevice/microsoft.graph.setDeviceName";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSetDeviceNameRequestBuilder instantiates a new SetDeviceNameRequestBuilder and sets the default values.
func NewSetDeviceNameRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetDeviceNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetDeviceNameRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation set device name of the device.
func (m *SetDeviceNameRequestBuilder) CreatePostRequestInformation(options *SetDeviceNameRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post set device name of the device.
func (m *SetDeviceNameRequestBuilder) Post(options *SetDeviceNameRequestBuilderPostOptions)(error) {
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
