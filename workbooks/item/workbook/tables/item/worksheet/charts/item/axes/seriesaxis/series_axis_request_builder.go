package seriesaxis

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i302c594de48d394dfa434ee8cd8c665e67b018b8a9e6f50d3265dc97ea272db6 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/axes/seriesaxis/minorgridlines"
    i4094237c59704ba6c10249d038ec96c1112764b23cc969cbfacff90f0ebf28db "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/axes/seriesaxis/majorgridlines"
    i994b3bd7f633579dff24a58a17f0a5d9383a8a78ad45fb057ae316b64a59d9c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/axes/seriesaxis/format"
    ie0d318c5ad4c15fc1844cdce3273fa9190e3f7c2a43822cb049ac0cd1808c105 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/axes/seriesaxis/title"
)

// SeriesAxisRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\tables\{workbookTable-id}\worksheet\charts\{workbookChart-id}\axes\seriesAxis
type SeriesAxisRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SeriesAxisRequestBuilderDeleteOptions options for Delete
type SeriesAxisRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SeriesAxisRequestBuilderGetOptions options for Get
type SeriesAxisRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SeriesAxisRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SeriesAxisRequestBuilderGetQueryParameters represents the series axis of a 3-dimensional chart. Read-only.
type SeriesAxisRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SeriesAxisRequestBuilderPatchOptions options for Patch
type SeriesAxisRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChartAxis;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSeriesAxisRequestBuilderInternal instantiates a new SeriesAxisRequestBuilder and sets the default values.
func NewSeriesAxisRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SeriesAxisRequestBuilder) {
    m := &SeriesAxisRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/tables/{workbookTable_id}/worksheet/charts/{workbookChart_id}/axes/seriesAxis{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSeriesAxisRequestBuilder instantiates a new SeriesAxisRequestBuilder and sets the default values.
func NewSeriesAxisRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SeriesAxisRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSeriesAxisRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation represents the series axis of a 3-dimensional chart. Read-only.
func (m *SeriesAxisRequestBuilder) CreateDeleteRequestInformation(options *SeriesAxisRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation represents the series axis of a 3-dimensional chart. Read-only.
func (m *SeriesAxisRequestBuilder) CreateGetRequestInformation(options *SeriesAxisRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation represents the series axis of a 3-dimensional chart. Read-only.
func (m *SeriesAxisRequestBuilder) CreatePatchRequestInformation(options *SeriesAxisRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete represents the series axis of a 3-dimensional chart. Read-only.
func (m *SeriesAxisRequestBuilder) Delete(options *SeriesAxisRequestBuilderDeleteOptions)(error) {
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
func (m *SeriesAxisRequestBuilder) Format()(*i994b3bd7f633579dff24a58a17f0a5d9383a8a78ad45fb057ae316b64a59d9c3.FormatRequestBuilder) {
    return i994b3bd7f633579dff24a58a17f0a5d9383a8a78ad45fb057ae316b64a59d9c3.NewFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get represents the series axis of a 3-dimensional chart. Read-only.
func (m *SeriesAxisRequestBuilder) Get(options *SeriesAxisRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChartAxis, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWorkbookChartAxis() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChartAxis), nil
}
func (m *SeriesAxisRequestBuilder) MajorGridlines()(*i4094237c59704ba6c10249d038ec96c1112764b23cc969cbfacff90f0ebf28db.MajorGridlinesRequestBuilder) {
    return i4094237c59704ba6c10249d038ec96c1112764b23cc969cbfacff90f0ebf28db.NewMajorGridlinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SeriesAxisRequestBuilder) MinorGridlines()(*i302c594de48d394dfa434ee8cd8c665e67b018b8a9e6f50d3265dc97ea272db6.MinorGridlinesRequestBuilder) {
    return i302c594de48d394dfa434ee8cd8c665e67b018b8a9e6f50d3265dc97ea272db6.NewMinorGridlinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch represents the series axis of a 3-dimensional chart. Read-only.
func (m *SeriesAxisRequestBuilder) Patch(options *SeriesAxisRequestBuilderPatchOptions)(error) {
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
func (m *SeriesAxisRequestBuilder) Title()(*ie0d318c5ad4c15fc1844cdce3273fa9190e3f7c2a43822cb049ac0cd1808c105.TitleRequestBuilder) {
    return ie0d318c5ad4c15fc1844cdce3273fa9190e3f7c2a43822cb049ac0cd1808c105.NewTitleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
