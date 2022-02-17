package muteall

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// MuteAllRequestBuilder builds and executes requests for operations under \app\calls\{call-id}\participants\microsoft.graph.muteAll
type MuteAllRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MuteAllRequestBuilderPostOptions options for Post
type MuteAllRequestBuilderPostOptions struct {
    // 
    Body *MuteAllRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MuteAllResponse union type wrapper for classes muteParticipantsOperation
type MuteAllResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type muteParticipantsOperation
    muteParticipantsOperation *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MuteParticipantsOperation;
}
// NewMuteAllResponse instantiates a new muteAllResponse and sets the default values.
func NewMuteAllResponse()(*MuteAllResponse) {
    m := &MuteAllResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MuteAllResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetMuteParticipantsOperation gets the muteParticipantsOperation property value. Union type representation for type muteParticipantsOperation
func (m *MuteAllResponse) GetMuteParticipantsOperation()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MuteParticipantsOperation) {
    if m == nil {
        return nil
    } else {
        return m.muteParticipantsOperation
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MuteAllResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["muteParticipantsOperation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewMuteParticipantsOperation() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMuteParticipantsOperation(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MuteParticipantsOperation))
        }
        return nil
    }
    return res
}
func (m *MuteAllResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MuteAllResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("muteParticipantsOperation", m.GetMuteParticipantsOperation())
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
func (m *MuteAllResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMuteParticipantsOperation sets the muteParticipantsOperation property value. Union type representation for type muteParticipantsOperation
func (m *MuteAllResponse) SetMuteParticipantsOperation(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MuteParticipantsOperation)() {
    if m != nil {
        m.muteParticipantsOperation = value
    }
}
// NewMuteAllRequestBuilderInternal instantiates a new MuteAllRequestBuilder and sets the default values.
func NewMuteAllRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MuteAllRequestBuilder) {
    m := &MuteAllRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/app/calls/{call_id}/participants/microsoft.graph.muteAll";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMuteAllRequestBuilder instantiates a new MuteAllRequestBuilder and sets the default values.
func NewMuteAllRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MuteAllRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMuteAllRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action muteAll
func (m *MuteAllRequestBuilder) CreatePostRequestInformation(options *MuteAllRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action muteAll
func (m *MuteAllRequestBuilder) Post(options *MuteAllRequestBuilderPostOptions)(*MuteAllResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMuteAllResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*MuteAllResponse), nil
}
