package updatetiindicators

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UpdateTiIndicatorsRequestBuilder builds and executes requests for operations under \security\tiIndicators\microsoft.graph.updateTiIndicators
type UpdateTiIndicatorsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UpdateTiIndicatorsRequestBuilderPostOptions options for Post
type UpdateTiIndicatorsRequestBuilderPostOptions struct {
    // 
    Body *UpdateTiIndicatorsRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUpdateTiIndicatorsRequestBuilderInternal instantiates a new UpdateTiIndicatorsRequestBuilder and sets the default values.
func NewUpdateTiIndicatorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdateTiIndicatorsRequestBuilder) {
    m := &UpdateTiIndicatorsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/tiIndicators/microsoft.graph.updateTiIndicators";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUpdateTiIndicatorsRequestBuilder instantiates a new UpdateTiIndicatorsRequestBuilder and sets the default values.
func NewUpdateTiIndicatorsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdateTiIndicatorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUpdateTiIndicatorsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action updateTiIndicators
func (m *UpdateTiIndicatorsRequestBuilder) CreatePostRequestInformation(options *UpdateTiIndicatorsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action updateTiIndicators
func (m *UpdateTiIndicatorsRequestBuilder) Post(options *UpdateTiIndicatorsRequestBuilderPostOptions)([]UpdateTiIndicators, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendCollectionAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUpdateTiIndicators() }, nil, nil)
    if err != nil {
        return nil, err
    }
    val := make([]UpdateTiIndicators, len(res))
    for i, v := range res {
        val[i] = *(v.(*UpdateTiIndicators))
    }
    return val, nil
}
