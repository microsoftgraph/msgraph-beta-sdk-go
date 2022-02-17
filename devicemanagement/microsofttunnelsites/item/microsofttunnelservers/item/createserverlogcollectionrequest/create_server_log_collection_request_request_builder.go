package createserverlogcollectionrequest

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// CreateServerLogCollectionRequestRequestBuilder builds and executes requests for operations under \deviceManagement\microsoftTunnelSites\{microsoftTunnelSite-id}\microsoftTunnelServers\{microsoftTunnelServer-id}\microsoft.graph.createServerLogCollectionRequest
type CreateServerLogCollectionRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CreateServerLogCollectionRequestRequestBuilderPostOptions options for Post
type CreateServerLogCollectionRequestRequestBuilderPostOptions struct {
    // 
    Body *CreateServerLogCollectionRequestRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CreateServerLogCollectionRequestResponse union type wrapper for classes microsoftTunnelServerLogCollectionResponse
type CreateServerLogCollectionRequestResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type microsoftTunnelServerLogCollectionResponse
    microsoftTunnelServerLogCollectionResponse *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MicrosoftTunnelServerLogCollectionResponse;
}
// NewCreateServerLogCollectionRequestResponse instantiates a new createServerLogCollectionRequestResponse and sets the default values.
func NewCreateServerLogCollectionRequestResponse()(*CreateServerLogCollectionRequestResponse) {
    m := &CreateServerLogCollectionRequestResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreateServerLogCollectionRequestResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetMicrosoftTunnelServerLogCollectionResponse gets the microsoftTunnelServerLogCollectionResponse property value. Union type representation for type microsoftTunnelServerLogCollectionResponse
func (m *CreateServerLogCollectionRequestResponse) GetMicrosoftTunnelServerLogCollectionResponse()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MicrosoftTunnelServerLogCollectionResponse) {
    if m == nil {
        return nil
    } else {
        return m.microsoftTunnelServerLogCollectionResponse
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CreateServerLogCollectionRequestResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["microsoftTunnelServerLogCollectionResponse"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewMicrosoftTunnelServerLogCollectionResponse() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftTunnelServerLogCollectionResponse(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MicrosoftTunnelServerLogCollectionResponse))
        }
        return nil
    }
    return res
}
func (m *CreateServerLogCollectionRequestResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CreateServerLogCollectionRequestResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("microsoftTunnelServerLogCollectionResponse", m.GetMicrosoftTunnelServerLogCollectionResponse())
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
func (m *CreateServerLogCollectionRequestResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMicrosoftTunnelServerLogCollectionResponse sets the microsoftTunnelServerLogCollectionResponse property value. Union type representation for type microsoftTunnelServerLogCollectionResponse
func (m *CreateServerLogCollectionRequestResponse) SetMicrosoftTunnelServerLogCollectionResponse(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MicrosoftTunnelServerLogCollectionResponse)() {
    if m != nil {
        m.microsoftTunnelServerLogCollectionResponse = value
    }
}
// NewCreateServerLogCollectionRequestRequestBuilderInternal instantiates a new CreateServerLogCollectionRequestRequestBuilder and sets the default values.
func NewCreateServerLogCollectionRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CreateServerLogCollectionRequestRequestBuilder) {
    m := &CreateServerLogCollectionRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoftTunnelSites/{microsoftTunnelSite_id}/microsoftTunnelServers/{microsoftTunnelServer_id}/microsoft.graph.createServerLogCollectionRequest";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCreateServerLogCollectionRequestRequestBuilder instantiates a new CreateServerLogCollectionRequestRequestBuilder and sets the default values.
func NewCreateServerLogCollectionRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CreateServerLogCollectionRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCreateServerLogCollectionRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action createServerLogCollectionRequest
func (m *CreateServerLogCollectionRequestRequestBuilder) CreatePostRequestInformation(options *CreateServerLogCollectionRequestRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action createServerLogCollectionRequest
func (m *CreateServerLogCollectionRequestRequestBuilder) Post(options *CreateServerLogCollectionRequestRequestBuilderPostOptions)(*CreateServerLogCollectionRequestResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCreateServerLogCollectionRequestResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*CreateServerLogCollectionRequestResponse), nil
}
