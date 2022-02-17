package bulkrestorecloudpc

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// BulkRestoreCloudPcRequestBuilder builds and executes requests for operations under \deviceManagement\comanagedDevices\microsoft.graph.bulkRestoreCloudPc
type BulkRestoreCloudPcRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// BulkRestoreCloudPcRequestBuilderPostOptions options for Post
type BulkRestoreCloudPcRequestBuilderPostOptions struct {
    // 
    Body *BulkRestoreCloudPcRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// BulkRestoreCloudPcResponse union type wrapper for classes cloudPcBulkRemoteActionResult
type BulkRestoreCloudPcResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type cloudPcBulkRemoteActionResult
    cloudPcBulkRemoteActionResult *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudPcBulkRemoteActionResult;
}
// NewBulkRestoreCloudPcResponse instantiates a new bulkRestoreCloudPcResponse and sets the default values.
func NewBulkRestoreCloudPcResponse()(*BulkRestoreCloudPcResponse) {
    m := &BulkRestoreCloudPcResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BulkRestoreCloudPcResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCloudPcBulkRemoteActionResult gets the cloudPcBulkRemoteActionResult property value. Union type representation for type cloudPcBulkRemoteActionResult
func (m *BulkRestoreCloudPcResponse) GetCloudPcBulkRemoteActionResult()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudPcBulkRemoteActionResult) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcBulkRemoteActionResult
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BulkRestoreCloudPcResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["cloudPcBulkRemoteActionResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewCloudPcBulkRemoteActionResult() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcBulkRemoteActionResult(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudPcBulkRemoteActionResult))
        }
        return nil
    }
    return res
}
func (m *BulkRestoreCloudPcResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *BulkRestoreCloudPcResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("cloudPcBulkRemoteActionResult", m.GetCloudPcBulkRemoteActionResult())
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
func (m *BulkRestoreCloudPcResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCloudPcBulkRemoteActionResult sets the cloudPcBulkRemoteActionResult property value. Union type representation for type cloudPcBulkRemoteActionResult
func (m *BulkRestoreCloudPcResponse) SetCloudPcBulkRemoteActionResult(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudPcBulkRemoteActionResult)() {
    if m != nil {
        m.cloudPcBulkRemoteActionResult = value
    }
}
// NewBulkRestoreCloudPcRequestBuilderInternal instantiates a new BulkRestoreCloudPcRequestBuilder and sets the default values.
func NewBulkRestoreCloudPcRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BulkRestoreCloudPcRequestBuilder) {
    m := &BulkRestoreCloudPcRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/comanagedDevices/microsoft.graph.bulkRestoreCloudPc";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBulkRestoreCloudPcRequestBuilder instantiates a new BulkRestoreCloudPcRequestBuilder and sets the default values.
func NewBulkRestoreCloudPcRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BulkRestoreCloudPcRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBulkRestoreCloudPcRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action bulkRestoreCloudPc
func (m *BulkRestoreCloudPcRequestBuilder) CreatePostRequestInformation(options *BulkRestoreCloudPcRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action bulkRestoreCloudPc
func (m *BulkRestoreCloudPcRequestBuilder) Post(options *BulkRestoreCloudPcRequestBuilderPostOptions)(*BulkRestoreCloudPcResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBulkRestoreCloudPcResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*BulkRestoreCloudPcResponse), nil
}
