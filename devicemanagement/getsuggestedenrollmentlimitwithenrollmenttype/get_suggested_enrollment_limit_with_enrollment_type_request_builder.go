package getsuggestedenrollmentlimitwithenrollmenttype

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder builds and executes requests for operations under \deviceManagement\microsoft.graph.getSuggestedEnrollmentLimit(enrollmentType='{enrollmentType}')
type GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderGetOptions options for Get
type GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GetSuggestedEnrollmentLimitWithEnrollmentTypeResponse union type wrapper for classes suggestedEnrollmentLimit
type GetSuggestedEnrollmentLimitWithEnrollmentTypeResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type suggestedEnrollmentLimit
    suggestedEnrollmentLimit *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SuggestedEnrollmentLimit;
}
// NewGetSuggestedEnrollmentLimitWithEnrollmentTypeResponse instantiates a new getSuggestedEnrollmentLimitWithEnrollmentTypeResponse and sets the default values.
func NewGetSuggestedEnrollmentLimitWithEnrollmentTypeResponse()(*GetSuggestedEnrollmentLimitWithEnrollmentTypeResponse) {
    m := &GetSuggestedEnrollmentLimitWithEnrollmentTypeResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetSuggestedEnrollmentLimitWithEnrollmentTypeResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetSuggestedEnrollmentLimit gets the suggestedEnrollmentLimit property value. Union type representation for type suggestedEnrollmentLimit
func (m *GetSuggestedEnrollmentLimitWithEnrollmentTypeResponse) GetSuggestedEnrollmentLimit()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SuggestedEnrollmentLimit) {
    if m == nil {
        return nil
    } else {
        return m.suggestedEnrollmentLimit
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetSuggestedEnrollmentLimitWithEnrollmentTypeResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["suggestedEnrollmentLimit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSuggestedEnrollmentLimit() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuggestedEnrollmentLimit(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SuggestedEnrollmentLimit))
        }
        return nil
    }
    return res
}
func (m *GetSuggestedEnrollmentLimitWithEnrollmentTypeResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetSuggestedEnrollmentLimitWithEnrollmentTypeResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("suggestedEnrollmentLimit", m.GetSuggestedEnrollmentLimit())
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
func (m *GetSuggestedEnrollmentLimitWithEnrollmentTypeResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetSuggestedEnrollmentLimit sets the suggestedEnrollmentLimit property value. Union type representation for type suggestedEnrollmentLimit
func (m *GetSuggestedEnrollmentLimitWithEnrollmentTypeResponse) SetSuggestedEnrollmentLimit(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SuggestedEnrollmentLimit)() {
    if m != nil {
        m.suggestedEnrollmentLimit = value
    }
}
// NewGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderInternal instantiates a new GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder and sets the default values.
func NewGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, enrollmentType *string)(*GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder) {
    m := &GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoft.graph.getSuggestedEnrollmentLimit(enrollmentType='{enrollmentType}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if enrollmentType != nil {
        urlTplParams["enrollmentType"] = *enrollmentType
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder instantiates a new GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder and sets the default values.
func NewGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getSuggestedEnrollmentLimit
func (m *GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder) CreateGetRequestInformation(options *GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function getSuggestedEnrollmentLimit
func (m *GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder) Get(options *GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderGetOptions)(*GetSuggestedEnrollmentLimitWithEnrollmentTypeResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGetSuggestedEnrollmentLimitWithEnrollmentTypeResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*GetSuggestedEnrollmentLimitWithEnrollmentTypeResponse), nil
}
