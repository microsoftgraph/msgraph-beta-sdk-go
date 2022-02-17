package getuserownedobjects

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetUserOwnedObjectsRequestBuilder builds and executes requests for operations under \contracts\microsoft.graph.getUserOwnedObjects
type GetUserOwnedObjectsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetUserOwnedObjectsRequestBuilderPostOptions options for Post
type GetUserOwnedObjectsRequestBuilderPostOptions struct {
    // 
    Body *GetUserOwnedObjectsRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GetUserOwnedObjectsResponse union type wrapper for classes directoryObject
type GetUserOwnedObjectsResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type directoryObject
    directoryObject *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryObject;
}
// NewGetUserOwnedObjectsResponse instantiates a new getUserOwnedObjectsResponse and sets the default values.
func NewGetUserOwnedObjectsResponse()(*GetUserOwnedObjectsResponse) {
    m := &GetUserOwnedObjectsResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetUserOwnedObjectsResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDirectoryObject gets the directoryObject property value. Union type representation for type directoryObject
func (m *GetUserOwnedObjectsResponse) GetDirectoryObject()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.directoryObject
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetUserOwnedObjectsResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["directoryObject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryObject(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryObject))
        }
        return nil
    }
    return res
}
func (m *GetUserOwnedObjectsResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetUserOwnedObjectsResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("directoryObject", m.GetDirectoryObject())
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
func (m *GetUserOwnedObjectsResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDirectoryObject sets the directoryObject property value. Union type representation for type directoryObject
func (m *GetUserOwnedObjectsResponse) SetDirectoryObject(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryObject)() {
    if m != nil {
        m.directoryObject = value
    }
}
// NewGetUserOwnedObjectsRequestBuilderInternal instantiates a new GetUserOwnedObjectsRequestBuilder and sets the default values.
func NewGetUserOwnedObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetUserOwnedObjectsRequestBuilder) {
    m := &GetUserOwnedObjectsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/contracts/microsoft.graph.getUserOwnedObjects";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetUserOwnedObjectsRequestBuilder instantiates a new GetUserOwnedObjectsRequestBuilder and sets the default values.
func NewGetUserOwnedObjectsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetUserOwnedObjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetUserOwnedObjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getUserOwnedObjects
func (m *GetUserOwnedObjectsRequestBuilder) CreatePostRequestInformation(options *GetUserOwnedObjectsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action getUserOwnedObjects
func (m *GetUserOwnedObjectsRequestBuilder) Post(options *GetUserOwnedObjectsRequestBuilderPostOptions)(*GetUserOwnedObjectsResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGetUserOwnedObjectsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*GetUserOwnedObjectsResponse), nil
}
