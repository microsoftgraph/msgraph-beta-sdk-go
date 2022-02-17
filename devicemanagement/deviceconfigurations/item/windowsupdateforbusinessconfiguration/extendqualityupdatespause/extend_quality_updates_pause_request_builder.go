package extendqualityupdatespause

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// ExtendQualityUpdatesPauseRequestBuilder builds and executes requests for operations under \deviceManagement\deviceConfigurations\{deviceConfiguration-id}\microsoft.graph.windowsUpdateForBusinessConfiguration\microsoft.graph.extendQualityUpdatesPause
type ExtendQualityUpdatesPauseRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ExtendQualityUpdatesPauseRequestBuilderPostOptions options for Post
type ExtendQualityUpdatesPauseRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewExtendQualityUpdatesPauseRequestBuilderInternal instantiates a new ExtendQualityUpdatesPauseRequestBuilder and sets the default values.
func NewExtendQualityUpdatesPauseRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExtendQualityUpdatesPauseRequestBuilder) {
    m := &ExtendQualityUpdatesPauseRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration_id}/microsoft.graph.windowsUpdateForBusinessConfiguration/microsoft.graph.extendQualityUpdatesPause";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewExtendQualityUpdatesPauseRequestBuilder instantiates a new ExtendQualityUpdatesPauseRequestBuilder and sets the default values.
func NewExtendQualityUpdatesPauseRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExtendQualityUpdatesPauseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExtendQualityUpdatesPauseRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation extend Quality Updates Pause for a Windows Update for Business ring.
func (m *ExtendQualityUpdatesPauseRequestBuilder) CreatePostRequestInformation(options *ExtendQualityUpdatesPauseRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post extend Quality Updates Pause for a Windows Update for Business ring.
func (m *ExtendQualityUpdatesPauseRequestBuilder) Post(options *ExtendQualityUpdatesPauseRequestBuilderPostOptions)(error) {
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
