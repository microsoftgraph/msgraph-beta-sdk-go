package activatedeviceesim

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// ActivateDeviceEsimRequestBuilder builds and executes requests for operations under \deviceManagement\deviceCustomAttributeShellScripts\{deviceCustomAttributeShellScript-id}\userRunStates\{deviceManagementScriptUserState-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.activateDeviceEsim
type ActivateDeviceEsimRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ActivateDeviceEsimRequestBuilderPostOptions options for Post
type ActivateDeviceEsimRequestBuilderPostOptions struct {
    // 
    Body *ActivateDeviceEsimRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewActivateDeviceEsimRequestBuilderInternal instantiates a new ActivateDeviceEsimRequestBuilder and sets the default values.
func NewActivateDeviceEsimRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ActivateDeviceEsimRequestBuilder) {
    m := &ActivateDeviceEsimRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCustomAttributeShellScripts/{deviceCustomAttributeShellScript_id}/userRunStates/{deviceManagementScriptUserState_id}/deviceRunStates/{deviceManagementScriptDeviceState_id}/managedDevice/microsoft.graph.activateDeviceEsim";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewActivateDeviceEsimRequestBuilder instantiates a new ActivateDeviceEsimRequestBuilder and sets the default values.
func NewActivateDeviceEsimRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ActivateDeviceEsimRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewActivateDeviceEsimRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation activate eSIM on the device.
func (m *ActivateDeviceEsimRequestBuilder) CreatePostRequestInformation(options *ActivateDeviceEsimRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post activate eSIM on the device.
func (m *ActivateDeviceEsimRequestBuilder) Post(options *ActivateDeviceEsimRequestBuilderPostOptions)(error) {
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
