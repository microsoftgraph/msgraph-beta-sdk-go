package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i7cc870ef048a207a9014ea9ac0661af8ca9c9e902f9300f0bddcd8017bd04dc1 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/series/item/format"
    ie639119ad6424d58827d1cb29471e73d99947256632b8cf6454ea57a4b25ca75 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/series/item/points"
    if6f600e401d23fe3acf281ff9864255f1af2b032ed20578a65ffc5593fee695d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/series/item/points/item"
)

// WorkbookChartSeriesRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\charts\{workbookChart-id}\series\{workbookChartSeries-id}
type WorkbookChartSeriesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WorkbookChartSeriesRequestBuilderDeleteOptions options for Delete
type WorkbookChartSeriesRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorkbookChartSeriesRequestBuilderGetOptions options for Get
type WorkbookChartSeriesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WorkbookChartSeriesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorkbookChartSeriesRequestBuilderGetQueryParameters represents either a single series or collection of series in the chart. Read-only.
type WorkbookChartSeriesRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// WorkbookChartSeriesRequestBuilderPatchOptions options for Patch
type WorkbookChartSeriesRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChartSeries;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewWorkbookChartSeriesRequestBuilderInternal instantiates a new WorkbookChartSeriesRequestBuilder and sets the default values.
func NewWorkbookChartSeriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookChartSeriesRequestBuilder) {
    m := &WorkbookChartSeriesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/names/{workbookNamedItem_id}/worksheet/charts/{workbookChart_id}/series/{workbookChartSeries_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWorkbookChartSeriesRequestBuilder instantiates a new WorkbookChartSeriesRequestBuilder and sets the default values.
func NewWorkbookChartSeriesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookChartSeriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookChartSeriesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation represents either a single series or collection of series in the chart. Read-only.
func (m *WorkbookChartSeriesRequestBuilder) CreateDeleteRequestInformation(options *WorkbookChartSeriesRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// CreateGetRequestInformation represents either a single series or collection of series in the chart. Read-only.
func (m *WorkbookChartSeriesRequestBuilder) CreateGetRequestInformation(options *WorkbookChartSeriesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation represents either a single series or collection of series in the chart. Read-only.
func (m *WorkbookChartSeriesRequestBuilder) CreatePatchRequestInformation(options *WorkbookChartSeriesRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
// Delete represents either a single series or collection of series in the chart. Read-only.
func (m *WorkbookChartSeriesRequestBuilder) Delete(options *WorkbookChartSeriesRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *WorkbookChartSeriesRequestBuilder) Format()(*i7cc870ef048a207a9014ea9ac0661af8ca9c9e902f9300f0bddcd8017bd04dc1.FormatRequestBuilder) {
    return i7cc870ef048a207a9014ea9ac0661af8ca9c9e902f9300f0bddcd8017bd04dc1.NewFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get represents either a single series or collection of series in the chart. Read-only.
func (m *WorkbookChartSeriesRequestBuilder) Get(options *WorkbookChartSeriesRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChartSeries, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWorkbookChartSeries() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChartSeries), nil
}
// Patch represents either a single series or collection of series in the chart. Read-only.
func (m *WorkbookChartSeriesRequestBuilder) Patch(options *WorkbookChartSeriesRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *WorkbookChartSeriesRequestBuilder) Points()(*ie639119ad6424d58827d1cb29471e73d99947256632b8cf6454ea57a4b25ca75.PointsRequestBuilder) {
    return ie639119ad6424d58827d1cb29471e73d99947256632b8cf6454ea57a4b25ca75.NewPointsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PointsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.workbook.names.item.worksheet.charts.item.series.item.points.item collection
func (m *WorkbookChartSeriesRequestBuilder) PointsById(id string)(*if6f600e401d23fe3acf281ff9864255f1af2b032ed20578a65ffc5593fee695d.WorkbookChartPointRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookChartPoint_id"] = id
    }
    return if6f600e401d23fe3acf281ff9864255f1af2b032ed20578a65ffc5593fee695d.NewWorkbookChartPointRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
