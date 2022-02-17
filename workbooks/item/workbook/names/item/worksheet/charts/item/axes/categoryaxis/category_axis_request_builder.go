package categoryaxis

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i14f0bf5992087601f32dbae817d91c6ea4ea138a1a26914ca064a03fa0e94ebb "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/axes/categoryaxis/minorgridlines"
    ia17bcf44d9f7e964052502ce1f1f560a0ad6d110c5531b7bd9a3ff274671b29b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/axes/categoryaxis/majorgridlines"
    ie08c3f6e89ef6b12263a455dbfb9e91f18aa1bef6481e64c1fcb187d46a3aa12 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/axes/categoryaxis/format"
    ief543dee382a94b5904d8b0767931e2f5105a756111ff08842e074504a88b4bf "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/axes/categoryaxis/title"
)

// CategoryAxisRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\charts\{workbookChart-id}\axes\categoryAxis
type CategoryAxisRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CategoryAxisRequestBuilderDeleteOptions options for Delete
type CategoryAxisRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CategoryAxisRequestBuilderGetOptions options for Get
type CategoryAxisRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CategoryAxisRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CategoryAxisRequestBuilderGetQueryParameters represents the category axis in a chart. Read-only.
type CategoryAxisRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// CategoryAxisRequestBuilderPatchOptions options for Patch
type CategoryAxisRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChartAxis;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewCategoryAxisRequestBuilderInternal instantiates a new CategoryAxisRequestBuilder and sets the default values.
func NewCategoryAxisRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CategoryAxisRequestBuilder) {
    m := &CategoryAxisRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/names/{workbookNamedItem_id}/worksheet/charts/{workbookChart_id}/axes/categoryAxis{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCategoryAxisRequestBuilder instantiates a new CategoryAxisRequestBuilder and sets the default values.
func NewCategoryAxisRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CategoryAxisRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCategoryAxisRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation represents the category axis in a chart. Read-only.
func (m *CategoryAxisRequestBuilder) CreateDeleteRequestInformation(options *CategoryAxisRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation represents the category axis in a chart. Read-only.
func (m *CategoryAxisRequestBuilder) CreateGetRequestInformation(options *CategoryAxisRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation represents the category axis in a chart. Read-only.
func (m *CategoryAxisRequestBuilder) CreatePatchRequestInformation(options *CategoryAxisRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete represents the category axis in a chart. Read-only.
func (m *CategoryAxisRequestBuilder) Delete(options *CategoryAxisRequestBuilderDeleteOptions)(error) {
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
func (m *CategoryAxisRequestBuilder) Format()(*ie08c3f6e89ef6b12263a455dbfb9e91f18aa1bef6481e64c1fcb187d46a3aa12.FormatRequestBuilder) {
    return ie08c3f6e89ef6b12263a455dbfb9e91f18aa1bef6481e64c1fcb187d46a3aa12.NewFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get represents the category axis in a chart. Read-only.
func (m *CategoryAxisRequestBuilder) Get(options *CategoryAxisRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChartAxis, error) {
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
func (m *CategoryAxisRequestBuilder) MajorGridlines()(*ia17bcf44d9f7e964052502ce1f1f560a0ad6d110c5531b7bd9a3ff274671b29b.MajorGridlinesRequestBuilder) {
    return ia17bcf44d9f7e964052502ce1f1f560a0ad6d110c5531b7bd9a3ff274671b29b.NewMajorGridlinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CategoryAxisRequestBuilder) MinorGridlines()(*i14f0bf5992087601f32dbae817d91c6ea4ea138a1a26914ca064a03fa0e94ebb.MinorGridlinesRequestBuilder) {
    return i14f0bf5992087601f32dbae817d91c6ea4ea138a1a26914ca064a03fa0e94ebb.NewMinorGridlinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch represents the category axis in a chart. Read-only.
func (m *CategoryAxisRequestBuilder) Patch(options *CategoryAxisRequestBuilderPatchOptions)(error) {
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
func (m *CategoryAxisRequestBuilder) Title()(*ief543dee382a94b5904d8b0767931e2f5105a756111ff08842e074504a88b4bf.TitleRequestBuilder) {
    return ief543dee382a94b5904d8b0767931e2f5105a756111ff08842e074504a88b4bf.NewTitleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
