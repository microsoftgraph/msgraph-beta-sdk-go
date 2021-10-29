package getassignmentfiltersstatusdetails

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// Builds and executes requests for operations under \deviceManagement\microsoft.graph.getAssignmentFiltersStatusDetails
type GetAssignmentFiltersStatusDetailsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Post
type GetAssignmentFiltersStatusDetailsRequestBuilderPostOptions struct {
    // 
    Body *GetAssignmentFiltersStatusDetailsRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Union type wrapper for classes assignmentFilterStatusDetails
type GetAssignmentFiltersStatusDetailsResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type assignmentFilterStatusDetails
    assignmentFilterStatusDetails *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AssignmentFilterStatusDetails;
}
// Instantiates a new getAssignmentFiltersStatusDetailsResponse and sets the default values.
func NewGetAssignmentFiltersStatusDetailsResponse()(*GetAssignmentFiltersStatusDetailsResponse) {
    m := &GetAssignmentFiltersStatusDetailsResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetAssignmentFiltersStatusDetailsResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the assignmentFilterStatusDetails property value. Union type representation for type assignmentFilterStatusDetails
func (m *GetAssignmentFiltersStatusDetailsResponse) GetAssignmentFilterStatusDetails()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AssignmentFilterStatusDetails) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterStatusDetails
    }
}
// The deserialization information for the current model
func (m *GetAssignmentFiltersStatusDetailsResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["assignmentFilterStatusDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAssignmentFilterStatusDetails() })
        if err != nil {
            return err
        }
        m.SetAssignmentFilterStatusDetails(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AssignmentFilterStatusDetails))
        return nil
    }
    return res
}
func (m *GetAssignmentFiltersStatusDetailsResponse) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetAssignmentFiltersStatusDetailsResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("assignmentFilterStatusDetails", m.GetAssignmentFilterStatusDetails())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *GetAssignmentFiltersStatusDetailsResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the assignmentFilterStatusDetails property value. Union type representation for type assignmentFilterStatusDetails
// Parameters:
//  - value : Value to set for the assignmentFilterStatusDetails property.
func (m *GetAssignmentFiltersStatusDetailsResponse) SetAssignmentFilterStatusDetails(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AssignmentFilterStatusDetails)() {
    m.assignmentFilterStatusDetails = value
}
// Instantiates a new GetAssignmentFiltersStatusDetailsRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewGetAssignmentFiltersStatusDetailsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetAssignmentFiltersStatusDetailsRequestBuilder) {
    m := &GetAssignmentFiltersStatusDetailsRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/microsoft.graph.getAssignmentFiltersStatusDetails";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new GetAssignmentFiltersStatusDetailsRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewGetAssignmentFiltersStatusDetailsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetAssignmentFiltersStatusDetailsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetAssignmentFiltersStatusDetailsRequestBuilderInternal(urlParams, requestAdapter)
}
// Invoke action getAssignmentFiltersStatusDetails
// Parameters:
//  - options : Options for the request
func (m *GetAssignmentFiltersStatusDetailsRequestBuilder) CreatePostRequestInformation(options *GetAssignmentFiltersStatusDetailsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Invoke action getAssignmentFiltersStatusDetails
// Parameters:
//  - options : Options for the request
func (m *GetAssignmentFiltersStatusDetailsRequestBuilder) Post(options *GetAssignmentFiltersStatusDetailsRequestBuilderPostOptions)(*GetAssignmentFiltersStatusDetailsResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGetAssignmentFiltersStatusDetailsResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*GetAssignmentFiltersStatusDetailsResponse), nil
}
