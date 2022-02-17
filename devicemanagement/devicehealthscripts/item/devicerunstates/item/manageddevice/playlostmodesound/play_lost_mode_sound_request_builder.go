package playlostmodesound

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// PlayLostModeSoundRequestBuilder builds and executes requests for operations under \deviceManagement\deviceHealthScripts\{deviceHealthScript-id}\deviceRunStates\{deviceHealthScriptDeviceState-id}\managedDevice\microsoft.graph.playLostModeSound
type PlayLostModeSoundRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PlayLostModeSoundRequestBuilderPostOptions options for Post
type PlayLostModeSoundRequestBuilderPostOptions struct {
    // 
    Body *PlayLostModeSoundRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewPlayLostModeSoundRequestBuilderInternal instantiates a new PlayLostModeSoundRequestBuilder and sets the default values.
func NewPlayLostModeSoundRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlayLostModeSoundRequestBuilder) {
    m := &PlayLostModeSoundRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceHealthScripts/{deviceHealthScript_id}/deviceRunStates/{deviceHealthScriptDeviceState_id}/managedDevice/microsoft.graph.playLostModeSound";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPlayLostModeSoundRequestBuilder instantiates a new PlayLostModeSoundRequestBuilder and sets the default values.
func NewPlayLostModeSoundRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlayLostModeSoundRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlayLostModeSoundRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation play lost mode sound
func (m *PlayLostModeSoundRequestBuilder) CreatePostRequestInformation(options *PlayLostModeSoundRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post play lost mode sound
func (m *PlayLostModeSoundRequestBuilder) Post(options *PlayLostModeSoundRequestBuilderPostOptions)(error) {
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
