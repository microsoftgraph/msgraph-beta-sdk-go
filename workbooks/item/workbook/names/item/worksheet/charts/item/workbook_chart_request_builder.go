package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0770e5d3edaa6343454250456213e74ae7c4e1cfb0f63309a95db6bfb30a0e0b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/setdata"
    i0d63b5b38c4dbd31f208465d6b1167bfe0f43830a02ea3c7d73b67fc5bdcee7f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/series"
    i135cb12d4e690a4ca476139b511d15056c1663cfd6f888d7a33995a43fe4995f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/image"
    i29a75b9541809284953ed8d6146949483ae74828fa9b2de853fe532996dd389e "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/datalabels"
    i3a0b55bb01168ff80d60a792c5b6aed8135a82e5b8168e2c956effdd54094075 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/worksheet"
    i8eebd3d0441e66b7ce5a05ea0d2d1c768cafe61634ad6d5db06f0d2c89c78b36 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/format"
    i8ef2f438e7be159bc344065cbbfde74cae052f9185c2372612c8f536054a0fa0 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/imagewithwidthwithheightwithfittingmode"
    ib1c2c73e175132eeb6d72e8bc938280042ab81e9a96fc30d0712980b2973146f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/imagewithwidthwithheight"
    ic268d93b7b72d9e6e54c68478f6133b8566ab4853d35a2d927c3cca02d9efa15 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/setposition"
    ic374e3aaec8919c26ff1c82541429f68388bc12c0ab2d935bd6fc21027da25a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/title"
    id432606ecec713da621d6ff2b62be55fbd25209f1adae0f3a8dbc266642eb089 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/legend"
    idc4909f345126c7783d965170e3f44946d32f0e405ead79f744ad10fa66b6745 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/axes"
    if6d6b845fc4da1172ee81c2a3ddcf3d4660aab5b65d6e3e5ae70375260d9d2fa "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/imagewithwidth"
    i2cab366b5da7f96e96cd7080085a3fe753c1da31c7a96686d0f0ec7089bdbd4e "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/series/item"
)

// WorkbookChartRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\charts\{workbookChart-id}
type WorkbookChartRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WorkbookChartRequestBuilderDeleteOptions options for Delete
type WorkbookChartRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorkbookChartRequestBuilderGetOptions options for Get
type WorkbookChartRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WorkbookChartRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorkbookChartRequestBuilderGetQueryParameters returns collection of charts that are part of the worksheet. Read-only.
type WorkbookChartRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// WorkbookChartRequestBuilderPatchOptions options for Patch
type WorkbookChartRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChart;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *WorkbookChartRequestBuilder) Axes()(*idc4909f345126c7783d965170e3f44946d32f0e405ead79f744ad10fa66b6745.AxesRequestBuilder) {
    return idc4909f345126c7783d965170e3f44946d32f0e405ead79f744ad10fa66b6745.NewAxesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewWorkbookChartRequestBuilderInternal instantiates a new WorkbookChartRequestBuilder and sets the default values.
func NewWorkbookChartRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookChartRequestBuilder) {
    m := &WorkbookChartRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/names/{workbookNamedItem_id}/worksheet/charts/{workbookChart_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWorkbookChartRequestBuilder instantiates a new WorkbookChartRequestBuilder and sets the default values.
func NewWorkbookChartRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookChartRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookChartRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation returns collection of charts that are part of the worksheet. Read-only.
func (m *WorkbookChartRequestBuilder) CreateDeleteRequestInformation(options *WorkbookChartRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation returns collection of charts that are part of the worksheet. Read-only.
func (m *WorkbookChartRequestBuilder) CreateGetRequestInformation(options *WorkbookChartRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation returns collection of charts that are part of the worksheet. Read-only.
func (m *WorkbookChartRequestBuilder) CreatePatchRequestInformation(options *WorkbookChartRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WorkbookChartRequestBuilder) DataLabels()(*i29a75b9541809284953ed8d6146949483ae74828fa9b2de853fe532996dd389e.DataLabelsRequestBuilder) {
    return i29a75b9541809284953ed8d6146949483ae74828fa9b2de853fe532996dd389e.NewDataLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete returns collection of charts that are part of the worksheet. Read-only.
func (m *WorkbookChartRequestBuilder) Delete(options *WorkbookChartRequestBuilderDeleteOptions)(error) {
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
func (m *WorkbookChartRequestBuilder) Format()(*i8eebd3d0441e66b7ce5a05ea0d2d1c768cafe61634ad6d5db06f0d2c89c78b36.FormatRequestBuilder) {
    return i8eebd3d0441e66b7ce5a05ea0d2d1c768cafe61634ad6d5db06f0d2c89c78b36.NewFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get returns collection of charts that are part of the worksheet. Read-only.
func (m *WorkbookChartRequestBuilder) Get(options *WorkbookChartRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChart, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWorkbookChart() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChart), nil
}
// Image builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\charts\{workbookChart-id}\microsoft.graph.image()
func (m *WorkbookChartRequestBuilder) Image()(*i135cb12d4e690a4ca476139b511d15056c1663cfd6f888d7a33995a43fe4995f.ImageRequestBuilder) {
    return i135cb12d4e690a4ca476139b511d15056c1663cfd6f888d7a33995a43fe4995f.NewImageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ImageWithWidth builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\charts\{workbookChart-id}\microsoft.graph.image(width={width})
func (m *WorkbookChartRequestBuilder) ImageWithWidth(width *int32)(*if6d6b845fc4da1172ee81c2a3ddcf3d4660aab5b65d6e3e5ae70375260d9d2fa.ImageWithWidthRequestBuilder) {
    return if6d6b845fc4da1172ee81c2a3ddcf3d4660aab5b65d6e3e5ae70375260d9d2fa.NewImageWithWidthRequestBuilderInternal(m.pathParameters, m.requestAdapter, width);
}
// ImageWithWidthWithHeight builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\charts\{workbookChart-id}\microsoft.graph.image(width={width},height={height})
func (m *WorkbookChartRequestBuilder) ImageWithWidthWithHeight(width *int32, height *int32)(*ib1c2c73e175132eeb6d72e8bc938280042ab81e9a96fc30d0712980b2973146f.ImageWithWidthWithHeightRequestBuilder) {
    return ib1c2c73e175132eeb6d72e8bc938280042ab81e9a96fc30d0712980b2973146f.NewImageWithWidthWithHeightRequestBuilderInternal(m.pathParameters, m.requestAdapter, width, height);
}
// ImageWithWidthWithHeightWithFittingMode builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\charts\{workbookChart-id}\microsoft.graph.image(width={width},height={height},fittingMode='{fittingMode}')
func (m *WorkbookChartRequestBuilder) ImageWithWidthWithHeightWithFittingMode(width *int32, height *int32, fittingMode *string)(*i8ef2f438e7be159bc344065cbbfde74cae052f9185c2372612c8f536054a0fa0.ImageWithWidthWithHeightWithFittingModeRequestBuilder) {
    return i8ef2f438e7be159bc344065cbbfde74cae052f9185c2372612c8f536054a0fa0.NewImageWithWidthWithHeightWithFittingModeRequestBuilderInternal(m.pathParameters, m.requestAdapter, width, height, fittingMode);
}
func (m *WorkbookChartRequestBuilder) Legend()(*id432606ecec713da621d6ff2b62be55fbd25209f1adae0f3a8dbc266642eb089.LegendRequestBuilder) {
    return id432606ecec713da621d6ff2b62be55fbd25209f1adae0f3a8dbc266642eb089.NewLegendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch returns collection of charts that are part of the worksheet. Read-only.
func (m *WorkbookChartRequestBuilder) Patch(options *WorkbookChartRequestBuilderPatchOptions)(error) {
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
func (m *WorkbookChartRequestBuilder) Series()(*i0d63b5b38c4dbd31f208465d6b1167bfe0f43830a02ea3c7d73b67fc5bdcee7f.SeriesRequestBuilder) {
    return i0d63b5b38c4dbd31f208465d6b1167bfe0f43830a02ea3c7d73b67fc5bdcee7f.NewSeriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SeriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.workbook.names.item.worksheet.charts.item.series.item collection
func (m *WorkbookChartRequestBuilder) SeriesById(id string)(*i2cab366b5da7f96e96cd7080085a3fe753c1da31c7a96686d0f0ec7089bdbd4e.WorkbookChartSeriesRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["workbookChartSeries_id"] = id
    }
    return i2cab366b5da7f96e96cd7080085a3fe753c1da31c7a96686d0f0ec7089bdbd4e.NewWorkbookChartSeriesRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) SetData()(*i0770e5d3edaa6343454250456213e74ae7c4e1cfb0f63309a95db6bfb30a0e0b.SetDataRequestBuilder) {
    return i0770e5d3edaa6343454250456213e74ae7c4e1cfb0f63309a95db6bfb30a0e0b.NewSetDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) SetPosition()(*ic268d93b7b72d9e6e54c68478f6133b8566ab4853d35a2d927c3cca02d9efa15.SetPositionRequestBuilder) {
    return ic268d93b7b72d9e6e54c68478f6133b8566ab4853d35a2d927c3cca02d9efa15.NewSetPositionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) Title()(*ic374e3aaec8919c26ff1c82541429f68388bc12c0ab2d935bd6fc21027da25a3.TitleRequestBuilder) {
    return ic374e3aaec8919c26ff1c82541429f68388bc12c0ab2d935bd6fc21027da25a3.NewTitleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) Worksheet()(*i3a0b55bb01168ff80d60a792c5b6aed8135a82e5b8168e2c956effdd54094075.WorksheetRequestBuilder) {
    return i3a0b55bb01168ff80d60a792c5b6aed8135a82e5b8168e2c956effdd54094075.NewWorksheetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
