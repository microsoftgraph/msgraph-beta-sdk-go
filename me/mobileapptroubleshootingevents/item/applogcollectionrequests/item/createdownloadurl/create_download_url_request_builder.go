package createdownloadurl

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// CreateDownloadUrlRequestBuilder builds and executes requests for operations under \me\mobileAppTroubleshootingEvents\{mobileAppTroubleshootingEvent-id}\appLogCollectionRequests\{appLogCollectionRequest-id}\microsoft.graph.createDownloadUrl
type CreateDownloadUrlRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CreateDownloadUrlRequestBuilderPostOptions options for Post
type CreateDownloadUrlRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CreateDownloadUrlResponse union type wrapper for classes appLogCollectionDownloadDetails
type CreateDownloadUrlResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type appLogCollectionDownloadDetails
    appLogCollectionDownloadDetails *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AppLogCollectionDownloadDetails;
}
// NewCreateDownloadUrlResponse instantiates a new createDownloadUrlResponse and sets the default values.
func NewCreateDownloadUrlResponse()(*CreateDownloadUrlResponse) {
    m := &CreateDownloadUrlResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreateDownloadUrlResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAppLogCollectionDownloadDetails gets the appLogCollectionDownloadDetails property value. Union type representation for type appLogCollectionDownloadDetails
func (m *CreateDownloadUrlResponse) GetAppLogCollectionDownloadDetails()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AppLogCollectionDownloadDetails) {
    if m == nil {
        return nil
    } else {
        return m.appLogCollectionDownloadDetails
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CreateDownloadUrlResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["appLogCollectionDownloadDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAppLogCollectionDownloadDetails() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppLogCollectionDownloadDetails(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AppLogCollectionDownloadDetails))
        }
        return nil
    }
    return res
}
func (m *CreateDownloadUrlResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CreateDownloadUrlResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("appLogCollectionDownloadDetails", m.GetAppLogCollectionDownloadDetails())
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
func (m *CreateDownloadUrlResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAppLogCollectionDownloadDetails sets the appLogCollectionDownloadDetails property value. Union type representation for type appLogCollectionDownloadDetails
func (m *CreateDownloadUrlResponse) SetAppLogCollectionDownloadDetails(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AppLogCollectionDownloadDetails)() {
    if m != nil {
        m.appLogCollectionDownloadDetails = value
    }
}
// NewCreateDownloadUrlRequestBuilderInternal instantiates a new CreateDownloadUrlRequestBuilder and sets the default values.
func NewCreateDownloadUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CreateDownloadUrlRequestBuilder) {
    m := &CreateDownloadUrlRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/mobileAppTroubleshootingEvents/{mobileAppTroubleshootingEvent_id}/appLogCollectionRequests/{appLogCollectionRequest_id}/microsoft.graph.createDownloadUrl";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCreateDownloadUrlRequestBuilder instantiates a new CreateDownloadUrlRequestBuilder and sets the default values.
func NewCreateDownloadUrlRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CreateDownloadUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCreateDownloadUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action createDownloadUrl
func (m *CreateDownloadUrlRequestBuilder) CreatePostRequestInformation(options *CreateDownloadUrlRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Post invoke action createDownloadUrl
func (m *CreateDownloadUrlRequestBuilder) Post(options *CreateDownloadUrlRequestBuilderPostOptions)(*CreateDownloadUrlResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCreateDownloadUrlResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*CreateDownloadUrlResponse), nil
}
