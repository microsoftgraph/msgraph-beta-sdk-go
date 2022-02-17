package charts

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i72f0fceb753e8d599c7ff603d3f0f9775ad152c17ed9f9395e9ff10c3031e1c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/count"
    iaae83489ea294bb7dd8d9e59fd2a2aca00b5108777851284733d90a873d09944 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/add"
    if12b5f364787664c00eedfbeb5f5ed16d386d218fa4367fd86ac91b6c46a8461 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/itemwithname"
    if9805569b100180f5dbffb55100a765cf174796896ba0b919bd1e38e84c60bd8 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/itematwithindex"
)

// ChartsRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\charts
type ChartsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ChartsRequestBuilderGetOptions options for Get
type ChartsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ChartsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ChartsRequestBuilderGetQueryParameters returns collection of charts that are part of the worksheet. Read-only.
type ChartsRequestBuilderGetQueryParameters struct {
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
// ChartsRequestBuilderPostOptions options for Post
type ChartsRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChart;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ChartsRequestBuilder) Add()(*iaae83489ea294bb7dd8d9e59fd2a2aca00b5108777851284733d90a873d09944.AddRequestBuilder) {
    return iaae83489ea294bb7dd8d9e59fd2a2aca00b5108777851284733d90a873d09944.NewAddRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewChartsRequestBuilderInternal instantiates a new ChartsRequestBuilder and sets the default values.
func NewChartsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChartsRequestBuilder) {
    m := &ChartsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/tables/{workbookTable_id}/worksheet/charts{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChartsRequestBuilder instantiates a new ChartsRequestBuilder and sets the default values.
func NewChartsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChartsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChartsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\charts\microsoft.graph.count()
func (m *ChartsRequestBuilder) Count()(*i72f0fceb753e8d599c7ff603d3f0f9775ad152c17ed9f9395e9ff10c3031e1c0.CountRequestBuilder) {
    return i72f0fceb753e8d599c7ff603d3f0f9775ad152c17ed9f9395e9ff10c3031e1c0.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation returns collection of charts that are part of the worksheet. Read-only.
func (m *ChartsRequestBuilder) CreateGetRequestInformation(options *ChartsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation returns collection of charts that are part of the worksheet. Read-only.
func (m *ChartsRequestBuilder) CreatePostRequestInformation(options *ChartsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get returns collection of charts that are part of the worksheet. Read-only.
func (m *ChartsRequestBuilder) Get(options *ChartsRequestBuilderGetOptions)(*ChartsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChartsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ChartsResponse), nil
}
// ItemAtWithIndex builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\charts\microsoft.graph.itemAt(index={index})
func (m *ChartsRequestBuilder) ItemAtWithIndex(index *int32)(*if9805569b100180f5dbffb55100a765cf174796896ba0b919bd1e38e84c60bd8.ItemAtWithIndexRequestBuilder) {
    return if9805569b100180f5dbffb55100a765cf174796896ba0b919bd1e38e84c60bd8.NewItemAtWithIndexRequestBuilderInternal(m.pathParameters, m.requestAdapter, index);
}
// ItemWithName builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\charts\microsoft.graph.item(name='{name}')
func (m *ChartsRequestBuilder) ItemWithName(name *string)(*if12b5f364787664c00eedfbeb5f5ed16d386d218fa4367fd86ac91b6c46a8461.ItemWithNameRequestBuilder) {
    return if12b5f364787664c00eedfbeb5f5ed16d386d218fa4367fd86ac91b6c46a8461.NewItemWithNameRequestBuilderInternal(m.pathParameters, m.requestAdapter, name);
}
// Post returns collection of charts that are part of the worksheet. Read-only.
func (m *ChartsRequestBuilder) Post(options *ChartsRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChart, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWorkbookChart() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChart), nil
}
