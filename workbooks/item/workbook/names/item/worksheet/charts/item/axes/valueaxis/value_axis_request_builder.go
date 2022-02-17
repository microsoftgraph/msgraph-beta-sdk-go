package valueaxis

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i9b8f1fd7fc0fb1c3edab4c74c3bf73762248f8affab55ae747a9dd11390d6054 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/axes/valueaxis/majorgridlines"
    iaeb9820d934daf1dc4a1be69e9c08b782edb6306dcbea92ea9b5a693b933252e "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/axes/valueaxis/minorgridlines"
    ic16e608129abc5bc8a26b7e02dda75c6a96b358af06c4e23ce3627040823fd4a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/axes/valueaxis/format"
    ic46e9f6fee18cf088eaca5a8e8a2410f17e24c87e11277a837f24104c18b89e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/names/item/worksheet/charts/item/axes/valueaxis/title"
)

// ValueAxisRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\names\{workbookNamedItem-id}\worksheet\charts\{workbookChart-id}\axes\valueAxis
type ValueAxisRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ValueAxisRequestBuilderDeleteOptions options for Delete
type ValueAxisRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ValueAxisRequestBuilderGetOptions options for Get
type ValueAxisRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ValueAxisRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ValueAxisRequestBuilderGetQueryParameters represents the value axis in an axis. Read-only.
type ValueAxisRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ValueAxisRequestBuilderPatchOptions options for Patch
type ValueAxisRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChartAxis;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewValueAxisRequestBuilderInternal instantiates a new ValueAxisRequestBuilder and sets the default values.
func NewValueAxisRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ValueAxisRequestBuilder) {
    m := &ValueAxisRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/names/{workbookNamedItem_id}/worksheet/charts/{workbookChart_id}/axes/valueAxis{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewValueAxisRequestBuilder instantiates a new ValueAxisRequestBuilder and sets the default values.
func NewValueAxisRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ValueAxisRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewValueAxisRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation represents the value axis in an axis. Read-only.
func (m *ValueAxisRequestBuilder) CreateDeleteRequestInformation(options *ValueAxisRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation represents the value axis in an axis. Read-only.
func (m *ValueAxisRequestBuilder) CreateGetRequestInformation(options *ValueAxisRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation represents the value axis in an axis. Read-only.
func (m *ValueAxisRequestBuilder) CreatePatchRequestInformation(options *ValueAxisRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete represents the value axis in an axis. Read-only.
func (m *ValueAxisRequestBuilder) Delete(options *ValueAxisRequestBuilderDeleteOptions)(error) {
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
func (m *ValueAxisRequestBuilder) Format()(*ic16e608129abc5bc8a26b7e02dda75c6a96b358af06c4e23ce3627040823fd4a.FormatRequestBuilder) {
    return ic16e608129abc5bc8a26b7e02dda75c6a96b358af06c4e23ce3627040823fd4a.NewFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get represents the value axis in an axis. Read-only.
func (m *ValueAxisRequestBuilder) Get(options *ValueAxisRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChartAxis, error) {
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
func (m *ValueAxisRequestBuilder) MajorGridlines()(*i9b8f1fd7fc0fb1c3edab4c74c3bf73762248f8affab55ae747a9dd11390d6054.MajorGridlinesRequestBuilder) {
    return i9b8f1fd7fc0fb1c3edab4c74c3bf73762248f8affab55ae747a9dd11390d6054.NewMajorGridlinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ValueAxisRequestBuilder) MinorGridlines()(*iaeb9820d934daf1dc4a1be69e9c08b782edb6306dcbea92ea9b5a693b933252e.MinorGridlinesRequestBuilder) {
    return iaeb9820d934daf1dc4a1be69e9c08b782edb6306dcbea92ea9b5a693b933252e.NewMinorGridlinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch represents the value axis in an axis. Read-only.
func (m *ValueAxisRequestBuilder) Patch(options *ValueAxisRequestBuilderPatchOptions)(error) {
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
func (m *ValueAxisRequestBuilder) Title()(*ic46e9f6fee18cf088eaca5a8e8a2410f17e24c87e11277a837f24104c18b89e7.TitleRequestBuilder) {
    return ic46e9f6fee18cf088eaca5a8e8a2410f17e24c87e11277a837f24104c18b89e7.NewTitleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
