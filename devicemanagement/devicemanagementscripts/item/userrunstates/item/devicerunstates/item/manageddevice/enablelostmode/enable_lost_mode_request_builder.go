package enablelostmode

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// EnableLostModeRequestBuilder builds and executes requests for operations under \deviceManagement\deviceManagementScripts\{deviceManagementScript-id}\userRunStates\{deviceManagementScriptUserState-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.enableLostMode
type EnableLostModeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EnableLostModeRequestBuilderPostOptions options for Post
type EnableLostModeRequestBuilderPostOptions struct {
    // 
    Body *EnableLostModeRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewEnableLostModeRequestBuilderInternal instantiates a new EnableLostModeRequestBuilder and sets the default values.
func NewEnableLostModeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EnableLostModeRequestBuilder) {
    m := &EnableLostModeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceManagementScripts/{deviceManagementScript_id}/userRunStates/{deviceManagementScriptUserState_id}/deviceRunStates/{deviceManagementScriptDeviceState_id}/managedDevice/microsoft.graph.enableLostMode";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEnableLostModeRequestBuilder instantiates a new EnableLostModeRequestBuilder and sets the default values.
func NewEnableLostModeRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EnableLostModeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnableLostModeRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation enable lost mode
func (m *EnableLostModeRequestBuilder) CreatePostRequestInformation(options *EnableLostModeRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post enable lost mode
func (m *EnableLostModeRequestBuilder) Post(options *EnableLostModeRequestBuilderPostOptions)(error) {
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
