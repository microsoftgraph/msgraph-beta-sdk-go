package rotatebitlockerkeys

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// RotateBitLockerKeysRequestBuilder builds and executes requests for operations under \deviceManagement\deviceShellScripts\{deviceShellScript-id}\userRunStates\{deviceManagementScriptUserState-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.rotateBitLockerKeys
type RotateBitLockerKeysRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RotateBitLockerKeysRequestBuilderPostOptions options for Post
type RotateBitLockerKeysRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewRotateBitLockerKeysRequestBuilderInternal instantiates a new RotateBitLockerKeysRequestBuilder and sets the default values.
func NewRotateBitLockerKeysRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RotateBitLockerKeysRequestBuilder) {
    m := &RotateBitLockerKeysRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceShellScripts/{deviceShellScript_id}/userRunStates/{deviceManagementScriptUserState_id}/deviceRunStates/{deviceManagementScriptDeviceState_id}/managedDevice/microsoft.graph.rotateBitLockerKeys";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRotateBitLockerKeysRequestBuilder instantiates a new RotateBitLockerKeysRequestBuilder and sets the default values.
func NewRotateBitLockerKeysRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RotateBitLockerKeysRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRotateBitLockerKeysRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation rotate BitLockerKeys
func (m *RotateBitLockerKeysRequestBuilder) CreatePostRequestInformation(options *RotateBitLockerKeysRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Post rotate BitLockerKeys
func (m *RotateBitLockerKeysRequestBuilder) Post(options *RotateBitLockerKeysRequestBuilderPostOptions)(error) {
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
