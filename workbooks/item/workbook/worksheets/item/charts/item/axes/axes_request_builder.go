package axes

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ia319f6626dcc9e9626aa310a8ef3a5133444552f1ada1f1d6a59b585b602c983 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/axes/seriesaxis"
    ibfcc6fac91eaf45ab1fa084261b841bb4e8b0bfbdda6e0094d272b58bfb8d83b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/axes/valueaxis"
    ieb51a913c30bd6e3a9419fd5803b77bcb45da99fd6360761eff0c6d4c7e79144 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/worksheets/item/charts/item/axes/categoryaxis"
)

// AxesRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\worksheets\{workbookWorksheet-id}\charts\{workbookChart-id}\axes
type AxesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AxesRequestBuilderDeleteOptions options for Delete
type AxesRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AxesRequestBuilderGetOptions options for Get
type AxesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AxesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AxesRequestBuilderGetQueryParameters represents chart axes. Read-only.
type AxesRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AxesRequestBuilderPatchOptions options for Patch
type AxesRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChartAxes;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AxesRequestBuilder) CategoryAxis()(*ieb51a913c30bd6e3a9419fd5803b77bcb45da99fd6360761eff0c6d4c7e79144.CategoryAxisRequestBuilder) {
    return ieb51a913c30bd6e3a9419fd5803b77bcb45da99fd6360761eff0c6d4c7e79144.NewCategoryAxisRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAxesRequestBuilderInternal instantiates a new AxesRequestBuilder and sets the default values.
func NewAxesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AxesRequestBuilder) {
    m := &AxesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/worksheets/{workbookWorksheet_id}/charts/{workbookChart_id}/axes{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAxesRequestBuilder instantiates a new AxesRequestBuilder and sets the default values.
func NewAxesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AxesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAxesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation represents chart axes. Read-only.
func (m *AxesRequestBuilder) CreateDeleteRequestInformation(options *AxesRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation represents chart axes. Read-only.
func (m *AxesRequestBuilder) CreateGetRequestInformation(options *AxesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation represents chart axes. Read-only.
func (m *AxesRequestBuilder) CreatePatchRequestInformation(options *AxesRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete represents chart axes. Read-only.
func (m *AxesRequestBuilder) Delete(options *AxesRequestBuilderDeleteOptions)(error) {
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
// Get represents chart axes. Read-only.
func (m *AxesRequestBuilder) Get(options *AxesRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChartAxes, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWorkbookChartAxes() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChartAxes), nil
}
// Patch represents chart axes. Read-only.
func (m *AxesRequestBuilder) Patch(options *AxesRequestBuilderPatchOptions)(error) {
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
func (m *AxesRequestBuilder) SeriesAxis()(*ia319f6626dcc9e9626aa310a8ef3a5133444552f1ada1f1d6a59b585b602c983.SeriesAxisRequestBuilder) {
    return ia319f6626dcc9e9626aa310a8ef3a5133444552f1ada1f1d6a59b585b602c983.NewSeriesAxisRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AxesRequestBuilder) ValueAxis()(*ibfcc6fac91eaf45ab1fa084261b841bb4e8b0bfbdda6e0094d272b58bfb8d83b.ValueAxisRequestBuilder) {
    return ibfcc6fac91eaf45ab1fa084261b841bb4e8b0bfbdda6e0094d272b58bfb8d83b.NewValueAxisRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
