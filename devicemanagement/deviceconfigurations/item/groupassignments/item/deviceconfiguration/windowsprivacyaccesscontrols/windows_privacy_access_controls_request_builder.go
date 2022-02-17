package windowsprivacyaccesscontrols

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// WindowsPrivacyAccessControlsRequestBuilder builds and executes requests for operations under \deviceManagement\deviceConfigurations\{deviceConfiguration-id}\groupAssignments\{deviceConfigurationGroupAssignment-id}\deviceConfiguration\microsoft.graph.windowsPrivacyAccessControls
type WindowsPrivacyAccessControlsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WindowsPrivacyAccessControlsRequestBuilderPostOptions options for Post
type WindowsPrivacyAccessControlsRequestBuilderPostOptions struct {
    // 
    Body *WindowsPrivacyAccessControlsRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewWindowsPrivacyAccessControlsRequestBuilderInternal instantiates a new WindowsPrivacyAccessControlsRequestBuilder and sets the default values.
func NewWindowsPrivacyAccessControlsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsPrivacyAccessControlsRequestBuilder) {
    m := &WindowsPrivacyAccessControlsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration_id}/groupAssignments/{deviceConfigurationGroupAssignment_id}/deviceConfiguration/microsoft.graph.windowsPrivacyAccessControls";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsPrivacyAccessControlsRequestBuilder instantiates a new WindowsPrivacyAccessControlsRequestBuilder and sets the default values.
func NewWindowsPrivacyAccessControlsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsPrivacyAccessControlsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsPrivacyAccessControlsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action windowsPrivacyAccessControls
func (m *WindowsPrivacyAccessControlsRequestBuilder) CreatePostRequestInformation(options *WindowsPrivacyAccessControlsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action windowsPrivacyAccessControls
func (m *WindowsPrivacyAccessControlsRequestBuilder) Post(options *WindowsPrivacyAccessControlsRequestBuilderPostOptions)(error) {
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
