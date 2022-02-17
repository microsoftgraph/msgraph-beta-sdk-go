package unassignresourceaccountfromdevice

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// UnassignResourceAccountFromDeviceRequestBuilder builds and executes requests for operations under \deviceManagement\windowsAutopilotDeploymentProfiles\{windowsAutopilotDeploymentProfile-id}\assignedDevices\{windowsAutopilotDeviceIdentity-id}\microsoft.graph.unassignResourceAccountFromDevice
type UnassignResourceAccountFromDeviceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UnassignResourceAccountFromDeviceRequestBuilderPostOptions options for Post
type UnassignResourceAccountFromDeviceRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUnassignResourceAccountFromDeviceRequestBuilderInternal instantiates a new UnassignResourceAccountFromDeviceRequestBuilder and sets the default values.
func NewUnassignResourceAccountFromDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnassignResourceAccountFromDeviceRequestBuilder) {
    m := &UnassignResourceAccountFromDeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/windowsAutopilotDeploymentProfiles/{windowsAutopilotDeploymentProfile_id}/assignedDevices/{windowsAutopilotDeviceIdentity_id}/microsoft.graph.unassignResourceAccountFromDevice";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnassignResourceAccountFromDeviceRequestBuilder instantiates a new UnassignResourceAccountFromDeviceRequestBuilder and sets the default values.
func NewUnassignResourceAccountFromDeviceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnassignResourceAccountFromDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnassignResourceAccountFromDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation unassigns the resource account from an Autopilot device.
func (m *UnassignResourceAccountFromDeviceRequestBuilder) CreatePostRequestInformation(options *UnassignResourceAccountFromDeviceRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post unassigns the resource account from an Autopilot device.
func (m *UnassignResourceAccountFromDeviceRequestBuilder) Post(options *UnassignResourceAccountFromDeviceRequestBuilderPostOptions)(error) {
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
