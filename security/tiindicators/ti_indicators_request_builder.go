package tiindicators

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5fd228a8811b3ee0e602ce99917149a6ef997b242fa74579dc28eb98aeff4610 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/tiindicators/updatetiindicators"
    i84faad3a80813a60a94e3416bfa220426811428a0f17f7a125a8a8389a71931a "github.com/microsoftgraph/msgraph-beta-sdk-go/security/tiindicators/deletetiindicatorsbyexternalid"
    id9554bf1b52f05e494c0e5962b41bb42f06710f09ff2286b921260ce9ef0c431 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/tiindicators/submittiindicators"
    ie38c9a1118b781a8cc486e08009768123690676d39af853ebdf06098c5110f09 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/tiindicators/deletetiindicators"
)

// TiIndicatorsRequestBuilder builds and executes requests for operations under \security\tiIndicators
type TiIndicatorsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TiIndicatorsRequestBuilderGetOptions options for Get
type TiIndicatorsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TiIndicatorsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TiIndicatorsRequestBuilderGetQueryParameters get tiIndicators from security
type TiIndicatorsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// TiIndicatorsRequestBuilderPostOptions options for Post
type TiIndicatorsRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TiIndicator;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewTiIndicatorsRequestBuilderInternal instantiates a new TiIndicatorsRequestBuilder and sets the default values.
func NewTiIndicatorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TiIndicatorsRequestBuilder) {
    m := &TiIndicatorsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/tiIndicators{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTiIndicatorsRequestBuilder instantiates a new TiIndicatorsRequestBuilder and sets the default values.
func NewTiIndicatorsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TiIndicatorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTiIndicatorsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get tiIndicators from security
func (m *TiIndicatorsRequestBuilder) CreateGetRequestInformation(options *TiIndicatorsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
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
// CreatePostRequestInformation create new navigation property to tiIndicators for security
func (m *TiIndicatorsRequestBuilder) CreatePostRequestInformation(options *TiIndicatorsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *TiIndicatorsRequestBuilder) DeleteTiIndicators()(*ie38c9a1118b781a8cc486e08009768123690676d39af853ebdf06098c5110f09.DeleteTiIndicatorsRequestBuilder) {
    return ie38c9a1118b781a8cc486e08009768123690676d39af853ebdf06098c5110f09.NewDeleteTiIndicatorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TiIndicatorsRequestBuilder) DeleteTiIndicatorsByExternalId()(*i84faad3a80813a60a94e3416bfa220426811428a0f17f7a125a8a8389a71931a.DeleteTiIndicatorsByExternalIdRequestBuilder) {
    return i84faad3a80813a60a94e3416bfa220426811428a0f17f7a125a8a8389a71931a.NewDeleteTiIndicatorsByExternalIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get tiIndicators from security
func (m *TiIndicatorsRequestBuilder) Get(options *TiIndicatorsRequestBuilderGetOptions)(*TiIndicatorsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTiIndicatorsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*TiIndicatorsResponse), nil
}
// Post create new navigation property to tiIndicators for security
func (m *TiIndicatorsRequestBuilder) Post(options *TiIndicatorsRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TiIndicator, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTiIndicator() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TiIndicator), nil
}
func (m *TiIndicatorsRequestBuilder) SubmitTiIndicators()(*id9554bf1b52f05e494c0e5962b41bb42f06710f09ff2286b921260ce9ef0c431.SubmitTiIndicatorsRequestBuilder) {
    return id9554bf1b52f05e494c0e5962b41bb42f06710f09ff2286b921260ce9ef0c431.NewSubmitTiIndicatorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TiIndicatorsRequestBuilder) UpdateTiIndicators()(*i5fd228a8811b3ee0e602ce99917149a6ef997b242fa74579dc28eb98aeff4610.UpdateTiIndicatorsRequestBuilder) {
    return i5fd228a8811b3ee0e602ce99917149a6ef997b242fa74579dc28eb98aeff4610.NewUpdateTiIndicatorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
