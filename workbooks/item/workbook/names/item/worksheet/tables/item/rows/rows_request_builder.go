package rows

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i09322a7bdf87b44a86a4bebe5033a6eee0194d0b3f2722c9d7f4c0e25c3aff26 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/rows/add"
    i5b826e8fb714eacc69c1dc3650ec2250e19dffa4cfbc5f679dbd923fb0f5f98b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/rows/itematwithindex"
    iabc2dbbeb528a3689fbe7bb263f070c18de0b4a7c5190ec9da8d12b0bb9de4a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/tables/item/rows/count"
)

// RowsRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\rows
type RowsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RowsRequestBuilderGetOptions options for Get
type RowsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RowsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RowsRequestBuilderGetQueryParameters represents a collection of all the rows in the table. Read-only.
type RowsRequestBuilderGetQueryParameters struct {
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
// RowsRequestBuilderPostOptions options for Post
type RowsRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTableRow;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *RowsRequestBuilder) Add()(*i09322a7bdf87b44a86a4bebe5033a6eee0194d0b3f2722c9d7f4c0e25c3aff26.AddRequestBuilder) {
    return i09322a7bdf87b44a86a4bebe5033a6eee0194d0b3f2722c9d7f4c0e25c3aff26.NewAddRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewRowsRequestBuilderInternal instantiates a new RowsRequestBuilder and sets the default values.
func NewRowsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RowsRequestBuilder) {
    m := &RowsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/names/{workbookNamedItem_id}/worksheet/tables/{workbookTable_id}/rows{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRowsRequestBuilder instantiates a new RowsRequestBuilder and sets the default values.
func NewRowsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RowsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRowsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\rows\microsoft.graph.count()
func (m *RowsRequestBuilder) Count()(*iabc2dbbeb528a3689fbe7bb263f070c18de0b4a7c5190ec9da8d12b0bb9de4a2.CountRequestBuilder) {
    return iabc2dbbeb528a3689fbe7bb263f070c18de0b4a7c5190ec9da8d12b0bb9de4a2.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation represents a collection of all the rows in the table. Read-only.
func (m *RowsRequestBuilder) CreateGetRequestInformation(options *RowsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation represents a collection of all the rows in the table. Read-only.
func (m *RowsRequestBuilder) CreatePostRequestInformation(options *RowsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get represents a collection of all the rows in the table. Read-only.
func (m *RowsRequestBuilder) Get(options *RowsRequestBuilderGetOptions)(*RowsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRowsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*RowsResponse), nil
}
// ItemAtWithIndex builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\tables\{workbookTable-id}\rows\microsoft.graph.itemAt(index={index})
func (m *RowsRequestBuilder) ItemAtWithIndex(index *int32)(*i5b826e8fb714eacc69c1dc3650ec2250e19dffa4cfbc5f679dbd923fb0f5f98b.ItemAtWithIndexRequestBuilder) {
    return i5b826e8fb714eacc69c1dc3650ec2250e19dffa4cfbc5f679dbd923fb0f5f98b.NewItemAtWithIndexRequestBuilderInternal(m.pathParameters, m.requestAdapter, index);
}
// Post represents a collection of all the rows in the table. Read-only.
func (m *RowsRequestBuilder) Post(options *RowsRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTableRow, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWorkbookTableRow() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTableRow), nil
}
