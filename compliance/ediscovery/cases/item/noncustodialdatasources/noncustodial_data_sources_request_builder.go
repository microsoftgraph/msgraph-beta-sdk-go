package noncustodialdatasources

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i6be3c26afae412e709f8c29691ed322d7a709004190807c07df909b750acf5fb "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/noncustodialdatasources/applyhold"
    ifd3d9f64cbcf567f464389efaaaf04cd06578883272854f51583af1c0ebb48f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/noncustodialdatasources/removehold"
)

// NoncustodialDataSourcesRequestBuilder builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}\noncustodialDataSources
type NoncustodialDataSourcesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// NoncustodialDataSourcesRequestBuilderGetOptions options for Get
type NoncustodialDataSourcesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *NoncustodialDataSourcesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NoncustodialDataSourcesRequestBuilderGetQueryParameters returns a list of case noncustodialDataSource objects for this case.  Nullable.
type NoncustodialDataSourcesRequestBuilderGetQueryParameters struct {
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
// NoncustodialDataSourcesRequestBuilderPostOptions options for Post
type NoncustodialDataSourcesRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NoncustodialDataSource;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *NoncustodialDataSourcesRequestBuilder) ApplyHold()(*i6be3c26afae412e709f8c29691ed322d7a709004190807c07df909b750acf5fb.ApplyHoldRequestBuilder) {
    return i6be3c26afae412e709f8c29691ed322d7a709004190807c07df909b750acf5fb.NewApplyHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewNoncustodialDataSourcesRequestBuilderInternal instantiates a new NoncustodialDataSourcesRequestBuilder and sets the default values.
func NewNoncustodialDataSourcesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*NoncustodialDataSourcesRequestBuilder) {
    m := &NoncustodialDataSourcesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case_id}/noncustodialDataSources{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewNoncustodialDataSourcesRequestBuilder instantiates a new NoncustodialDataSourcesRequestBuilder and sets the default values.
func NewNoncustodialDataSourcesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*NoncustodialDataSourcesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNoncustodialDataSourcesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation returns a list of case noncustodialDataSource objects for this case.  Nullable.
func (m *NoncustodialDataSourcesRequestBuilder) CreateGetRequestInformation(options *NoncustodialDataSourcesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation returns a list of case noncustodialDataSource objects for this case.  Nullable.
func (m *NoncustodialDataSourcesRequestBuilder) CreatePostRequestInformation(options *NoncustodialDataSourcesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get returns a list of case noncustodialDataSource objects for this case.  Nullable.
func (m *NoncustodialDataSourcesRequestBuilder) Get(options *NoncustodialDataSourcesRequestBuilderGetOptions)(*NoncustodialDataSourcesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNoncustodialDataSourcesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*NoncustodialDataSourcesResponse), nil
}
// Post returns a list of case noncustodialDataSource objects for this case.  Nullable.
func (m *NoncustodialDataSourcesRequestBuilder) Post(options *NoncustodialDataSourcesRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NoncustodialDataSource, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewNoncustodialDataSource() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NoncustodialDataSource), nil
}
func (m *NoncustodialDataSourcesRequestBuilder) RemoveHold()(*ifd3d9f64cbcf567f464389efaaaf04cd06578883272854f51583af1c0ebb48f9.RemoveHoldRequestBuilder) {
    return ifd3d9f64cbcf567f464389efaaaf04cd06578883272854f51583af1c0ebb48f9.NewRemoveHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
