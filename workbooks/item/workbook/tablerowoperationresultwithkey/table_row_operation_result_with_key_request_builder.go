package tablerowoperationresultwithkey

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// TableRowOperationResultWithKeyRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\microsoft.graph.tableRowOperationResult(key='{key}')
type TableRowOperationResultWithKeyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TableRowOperationResultWithKeyRequestBuilderGetOptions options for Get
type TableRowOperationResultWithKeyRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TableRowOperationResultWithKeyResponse union type wrapper for classes workbookTableRow
type TableRowOperationResultWithKeyResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type workbookTableRow
    workbookTableRow *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTableRow;
}
// NewTableRowOperationResultWithKeyResponse instantiates a new tableRowOperationResultWithKeyResponse and sets the default values.
func NewTableRowOperationResultWithKeyResponse()(*TableRowOperationResultWithKeyResponse) {
    m := &TableRowOperationResultWithKeyResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TableRowOperationResultWithKeyResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetWorkbookTableRow gets the workbookTableRow property value. Union type representation for type workbookTableRow
func (m *TableRowOperationResultWithKeyResponse) GetWorkbookTableRow()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTableRow) {
    if m == nil {
        return nil
    } else {
        return m.workbookTableRow
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TableRowOperationResultWithKeyResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["workbookTableRow"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWorkbookTableRow() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkbookTableRow(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTableRow))
        }
        return nil
    }
    return res
}
func (m *TableRowOperationResultWithKeyResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TableRowOperationResultWithKeyResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("workbookTableRow", m.GetWorkbookTableRow())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TableRowOperationResultWithKeyResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetWorkbookTableRow sets the workbookTableRow property value. Union type representation for type workbookTableRow
func (m *TableRowOperationResultWithKeyResponse) SetWorkbookTableRow(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookTableRow)() {
    if m != nil {
        m.workbookTableRow = value
    }
}
// NewTableRowOperationResultWithKeyRequestBuilderInternal instantiates a new TableRowOperationResultWithKeyRequestBuilder and sets the default values.
func NewTableRowOperationResultWithKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, key *string)(*TableRowOperationResultWithKeyRequestBuilder) {
    m := &TableRowOperationResultWithKeyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/microsoft.graph.tableRowOperationResult(key='{key}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if key != nil {
        urlTplParams["key"] = *key
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTableRowOperationResultWithKeyRequestBuilder instantiates a new TableRowOperationResultWithKeyRequestBuilder and sets the default values.
func NewTableRowOperationResultWithKeyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TableRowOperationResultWithKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTableRowOperationResultWithKeyRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function tableRowOperationResult
func (m *TableRowOperationResultWithKeyRequestBuilder) CreateGetRequestInformation(options *TableRowOperationResultWithKeyRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
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
// Get invoke function tableRowOperationResult
func (m *TableRowOperationResultWithKeyRequestBuilder) Get(options *TableRowOperationResultWithKeyRequestBuilderGetOptions)(*TableRowOperationResultWithKeyResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTableRowOperationResultWithKeyResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*TableRowOperationResultWithKeyResponse), nil
}
