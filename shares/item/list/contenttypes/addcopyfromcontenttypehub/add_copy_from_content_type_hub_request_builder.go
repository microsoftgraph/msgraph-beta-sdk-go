package addcopyfromcontenttypehub

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// AddCopyFromContentTypeHubRequestBuilder builds and executes requests for operations under \shares\{sharedDriveItem-id}\list\contentTypes\microsoft.graph.addCopyFromContentTypeHub
type AddCopyFromContentTypeHubRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AddCopyFromContentTypeHubRequestBuilderPostOptions options for Post
type AddCopyFromContentTypeHubRequestBuilderPostOptions struct {
    // 
    Body *AddCopyFromContentTypeHubRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AddCopyFromContentTypeHubResponse union type wrapper for classes contentType
type AddCopyFromContentTypeHubResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type contentType
    contentType *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType;
}
// NewAddCopyFromContentTypeHubResponse instantiates a new addCopyFromContentTypeHubResponse and sets the default values.
func NewAddCopyFromContentTypeHubResponse()(*AddCopyFromContentTypeHubResponse) {
    m := &AddCopyFromContentTypeHubResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AddCopyFromContentTypeHubResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContentType gets the contentType property value. Union type representation for type contentType
func (m *AddCopyFromContentTypeHubResponse) GetContentType()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AddCopyFromContentTypeHubResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["contentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewContentType() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType))
        }
        return nil
    }
    return res
}
func (m *AddCopyFromContentTypeHubResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AddCopyFromContentTypeHubResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("contentType", m.GetContentType())
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
func (m *AddCopyFromContentTypeHubResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContentType sets the contentType property value. Union type representation for type contentType
func (m *AddCopyFromContentTypeHubResponse) SetContentType(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType)() {
    if m != nil {
        m.contentType = value
    }
}
// NewAddCopyFromContentTypeHubRequestBuilderInternal instantiates a new AddCopyFromContentTypeHubRequestBuilder and sets the default values.
func NewAddCopyFromContentTypeHubRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AddCopyFromContentTypeHubRequestBuilder) {
    m := &AddCopyFromContentTypeHubRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem_id}/list/contentTypes/microsoft.graph.addCopyFromContentTypeHub";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAddCopyFromContentTypeHubRequestBuilder instantiates a new AddCopyFromContentTypeHubRequestBuilder and sets the default values.
func NewAddCopyFromContentTypeHubRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AddCopyFromContentTypeHubRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAddCopyFromContentTypeHubRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action addCopyFromContentTypeHub
func (m *AddCopyFromContentTypeHubRequestBuilder) CreatePostRequestInformation(options *AddCopyFromContentTypeHubRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action addCopyFromContentTypeHub
func (m *AddCopyFromContentTypeHubRequestBuilder) Post(options *AddCopyFromContentTypeHubRequestBuilderPostOptions)(*AddCopyFromContentTypeHubResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAddCopyFromContentTypeHubResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*AddCopyFromContentTypeHubResponse), nil
}
