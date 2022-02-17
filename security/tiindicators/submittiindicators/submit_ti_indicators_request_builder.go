package submittiindicators

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SubmitTiIndicatorsRequestBuilder builds and executes requests for operations under \security\tiIndicators\microsoft.graph.submitTiIndicators
type SubmitTiIndicatorsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SubmitTiIndicatorsRequestBuilderPostOptions options for Post
type SubmitTiIndicatorsRequestBuilderPostOptions struct {
    // 
    Body *SubmitTiIndicatorsRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSubmitTiIndicatorsRequestBuilderInternal instantiates a new SubmitTiIndicatorsRequestBuilder and sets the default values.
func NewSubmitTiIndicatorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SubmitTiIndicatorsRequestBuilder) {
    m := &SubmitTiIndicatorsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/tiIndicators/microsoft.graph.submitTiIndicators";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSubmitTiIndicatorsRequestBuilder instantiates a new SubmitTiIndicatorsRequestBuilder and sets the default values.
func NewSubmitTiIndicatorsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SubmitTiIndicatorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSubmitTiIndicatorsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action submitTiIndicators
func (m *SubmitTiIndicatorsRequestBuilder) CreatePostRequestInformation(options *SubmitTiIndicatorsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action submitTiIndicators
func (m *SubmitTiIndicatorsRequestBuilder) Post(options *SubmitTiIndicatorsRequestBuilderPostOptions)([]SubmitTiIndicators, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendCollectionAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSubmitTiIndicators() }, nil, nil)
    if err != nil {
        return nil, err
    }
    val := make([]SubmitTiIndicators, len(res))
    for i, v := range res {
        val[i] = *(v.(*SubmitTiIndicators))
    }
    return val, nil
}
