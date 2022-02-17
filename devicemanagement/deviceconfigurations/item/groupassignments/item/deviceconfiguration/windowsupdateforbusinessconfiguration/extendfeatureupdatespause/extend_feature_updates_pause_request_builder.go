package extendfeatureupdatespause

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// ExtendFeatureUpdatesPauseRequestBuilder builds and executes requests for operations under \deviceManagement\deviceConfigurations\{deviceConfiguration-id}\groupAssignments\{deviceConfigurationGroupAssignment-id}\deviceConfiguration\microsoft.graph.windowsUpdateForBusinessConfiguration\microsoft.graph.extendFeatureUpdatesPause
type ExtendFeatureUpdatesPauseRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ExtendFeatureUpdatesPauseRequestBuilderPostOptions options for Post
type ExtendFeatureUpdatesPauseRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewExtendFeatureUpdatesPauseRequestBuilderInternal instantiates a new ExtendFeatureUpdatesPauseRequestBuilder and sets the default values.
func NewExtendFeatureUpdatesPauseRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExtendFeatureUpdatesPauseRequestBuilder) {
    m := &ExtendFeatureUpdatesPauseRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration_id}/groupAssignments/{deviceConfigurationGroupAssignment_id}/deviceConfiguration/microsoft.graph.windowsUpdateForBusinessConfiguration/microsoft.graph.extendFeatureUpdatesPause";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewExtendFeatureUpdatesPauseRequestBuilder instantiates a new ExtendFeatureUpdatesPauseRequestBuilder and sets the default values.
func NewExtendFeatureUpdatesPauseRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExtendFeatureUpdatesPauseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExtendFeatureUpdatesPauseRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation extend Feature Updates Pause for a Windows Update for Business ring.
func (m *ExtendFeatureUpdatesPauseRequestBuilder) CreatePostRequestInformation(options *ExtendFeatureUpdatesPauseRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post extend Feature Updates Pause for a Windows Update for Business ring.
func (m *ExtendFeatureUpdatesPauseRequestBuilder) Post(options *ExtendFeatureUpdatesPauseRequestBuilderPostOptions)(error) {
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
