package registration

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i319d65f6c34b3fc2ed3543b9da4ffd50724671e1d8c1448e35dab886365ee566 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/registration/customquestions"
    id6cc9fa5871bc7bc0f0cdb14eefb4c23aab98b3f48d589f2b88c26747ddac518 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/registration/registrants"
    i1e3e5e104931ac0fdec551f3570e723ef72d8298672e117c6541fc33bb340f2e "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/registration/customquestions/item"
    i30c315dc061d49533ccaa40d62807a1351e7a46bd7fc3e004240034ba112b51d "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/registration/registrants/item"
)

// Builds and executes requests for operations under \app\onlineMeetings\{onlineMeeting-id}\registration
type RegistrationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type RegistrationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type RegistrationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RegistrationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
type RegistrationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type RegistrationRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MeetingRegistration;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new RegistrationRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewRegistrationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RegistrationRequestBuilder) {
    m := &RegistrationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/app/onlineMeetings/{onlineMeeting_id}/registration{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new RegistrationRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewRegistrationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RegistrationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRegistrationRequestBuilderInternal(urlParams, requestAdapter)
}
// The registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
// Parameters:
//  - options : Options for the request
func (m *RegistrationRequestBuilder) CreateDeleteRequestInformation(options *RegistrationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// The registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
// Parameters:
//  - options : Options for the request
func (m *RegistrationRequestBuilder) CreateGetRequestInformation(options *RegistrationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
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
// The registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
// Parameters:
//  - options : Options for the request
func (m *RegistrationRequestBuilder) CreatePatchRequestInformation(options *RegistrationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
func (m *RegistrationRequestBuilder) CustomQuestions()(*i319d65f6c34b3fc2ed3543b9da4ffd50724671e1d8c1448e35dab886365ee566.CustomQuestionsRequestBuilder) {
    return i319d65f6c34b3fc2ed3543b9da4ffd50724671e1d8c1448e35dab886365ee566.NewCustomQuestionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.app.onlineMeetings.item.registration.customQuestions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *RegistrationRequestBuilder) CustomQuestionsById(id string)(*i1e3e5e104931ac0fdec551f3570e723ef72d8298672e117c6541fc33bb340f2e.MeetingRegistrationQuestionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["meetingRegistrationQuestion_id"] = id
    }
    return i1e3e5e104931ac0fdec551f3570e723ef72d8298672e117c6541fc33bb340f2e.NewMeetingRegistrationQuestionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
// Parameters:
//  - options : Options for the request
func (m *RegistrationRequestBuilder) Delete(options *RegistrationRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// The registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
// Parameters:
//  - options : Options for the request
func (m *RegistrationRequestBuilder) Get(options *RegistrationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MeetingRegistration, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewMeetingRegistration() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MeetingRegistration), nil
}
// The registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
// Parameters:
//  - options : Options for the request
func (m *RegistrationRequestBuilder) Patch(options *RegistrationRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *RegistrationRequestBuilder) Registrants()(*id6cc9fa5871bc7bc0f0cdb14eefb4c23aab98b3f48d589f2b88c26747ddac518.RegistrantsRequestBuilder) {
    return id6cc9fa5871bc7bc0f0cdb14eefb4c23aab98b3f48d589f2b88c26747ddac518.NewRegistrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.app.onlineMeetings.item.registration.registrants.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *RegistrationRequestBuilder) RegistrantsById(id string)(*i30c315dc061d49533ccaa40d62807a1351e7a46bd7fc3e004240034ba112b51d.MeetingRegistrantRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["meetingRegistrant_id"] = id
    }
    return i30c315dc061d49533ccaa40d62807a1351e7a46bd7fc3e004240034ba112b51d.NewMeetingRegistrantRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
