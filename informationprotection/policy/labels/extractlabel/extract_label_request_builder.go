package extractlabel

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ExtractLabelRequestBuilder builds and executes requests for operations under \informationProtection\policy\labels\microsoft.graph.extractLabel
type ExtractLabelRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ExtractLabelRequestBuilderPostOptions options for Post
type ExtractLabelRequestBuilderPostOptions struct {
    // 
    Body *ExtractLabelRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ExtractLabelResponse union type wrapper for classes informationProtectionContentLabel
type ExtractLabelResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type informationProtectionContentLabel
    informationProtectionContentLabel *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InformationProtectionContentLabel;
}
// NewExtractLabelResponse instantiates a new extractLabelResponse and sets the default values.
func NewExtractLabelResponse()(*ExtractLabelResponse) {
    m := &ExtractLabelResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExtractLabelResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetInformationProtectionContentLabel gets the informationProtectionContentLabel property value. Union type representation for type informationProtectionContentLabel
func (m *ExtractLabelResponse) GetInformationProtectionContentLabel()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InformationProtectionContentLabel) {
    if m == nil {
        return nil
    } else {
        return m.informationProtectionContentLabel
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExtractLabelResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["informationProtectionContentLabel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewInformationProtectionContentLabel() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInformationProtectionContentLabel(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InformationProtectionContentLabel))
        }
        return nil
    }
    return res
}
func (m *ExtractLabelResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ExtractLabelResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("informationProtectionContentLabel", m.GetInformationProtectionContentLabel())
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
func (m *ExtractLabelResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetInformationProtectionContentLabel sets the informationProtectionContentLabel property value. Union type representation for type informationProtectionContentLabel
func (m *ExtractLabelResponse) SetInformationProtectionContentLabel(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InformationProtectionContentLabel)() {
    if m != nil {
        m.informationProtectionContentLabel = value
    }
}
// NewExtractLabelRequestBuilderInternal instantiates a new ExtractLabelRequestBuilder and sets the default values.
func NewExtractLabelRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExtractLabelRequestBuilder) {
    m := &ExtractLabelRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/informationProtection/policy/labels/microsoft.graph.extractLabel";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewExtractLabelRequestBuilder instantiates a new ExtractLabelRequestBuilder and sets the default values.
func NewExtractLabelRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExtractLabelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExtractLabelRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action extractLabel
func (m *ExtractLabelRequestBuilder) CreatePostRequestInformation(options *ExtractLabelRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action extractLabel
func (m *ExtractLabelRequestBuilder) Post(options *ExtractLabelRequestBuilderPostOptions)(*ExtractLabelResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExtractLabelResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ExtractLabelResponse), nil
}
