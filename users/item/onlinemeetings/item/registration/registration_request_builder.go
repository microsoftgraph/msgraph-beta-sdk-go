package registration

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i8fffe15359b63d385c4d09428627fd086207d8559033c154283ebae61be2989b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onlinemeetings/item/registration/registrants"
    ic80f4dba15f8a2e6471c424c8052facd225497424b2ebc8ec26c277612211cf3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onlinemeetings/item/registration/customquestions"
    i2ad1071fdb0ce5c7f13d818f7600904c2cfb682f002cf4a3d1323a2d717ccb6e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onlinemeetings/item/registration/customquestions/item"
    i34835e4f604126539f2c916992cff760482a1c7bf03f1d5d4928a2099ac5183d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onlinemeetings/item/registration/registrants/item"
)

// RegistrationRequestBuilder builds and executes requests for operations under \users\{user-id}\onlineMeetings\{onlineMeeting-id}\registration
type RegistrationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RegistrationRequestBuilderDeleteOptions options for Delete
type RegistrationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RegistrationRequestBuilderGetOptions options for Get
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
// RegistrationRequestBuilderGetQueryParameters the registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
type RegistrationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// RegistrationRequestBuilderPatchOptions options for Patch
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
// NewRegistrationRequestBuilderInternal instantiates a new RegistrationRequestBuilder and sets the default values.
func NewRegistrationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RegistrationRequestBuilder) {
    m := &RegistrationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/onlineMeetings/{onlineMeeting_id}/registration{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRegistrationRequestBuilder instantiates a new RegistrationRequestBuilder and sets the default values.
func NewRegistrationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RegistrationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRegistrationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
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
// CreateGetRequestInformation the registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
func (m *RegistrationRequestBuilder) CreateGetRequestInformation(options *RegistrationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
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
// CreatePatchRequestInformation the registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
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
func (m *RegistrationRequestBuilder) CustomQuestions()(*ic80f4dba15f8a2e6471c424c8052facd225497424b2ebc8ec26c277612211cf3.CustomQuestionsRequestBuilder) {
    return ic80f4dba15f8a2e6471c424c8052facd225497424b2ebc8ec26c277612211cf3.NewCustomQuestionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomQuestionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onlineMeetings.item.registration.customQuestions.item collection
func (m *RegistrationRequestBuilder) CustomQuestionsById(id string)(*i2ad1071fdb0ce5c7f13d818f7600904c2cfb682f002cf4a3d1323a2d717ccb6e.MeetingRegistrationQuestionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["meetingRegistrationQuestion_id"] = id
    }
    return i2ad1071fdb0ce5c7f13d818f7600904c2cfb682f002cf4a3d1323a2d717ccb6e.NewMeetingRegistrationQuestionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete the registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
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
// Get the registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
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
// Patch the registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
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
func (m *RegistrationRequestBuilder) Registrants()(*i8fffe15359b63d385c4d09428627fd086207d8559033c154283ebae61be2989b.RegistrantsRequestBuilder) {
    return i8fffe15359b63d385c4d09428627fd086207d8559033c154283ebae61be2989b.NewRegistrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegistrantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onlineMeetings.item.registration.registrants.item collection
func (m *RegistrationRequestBuilder) RegistrantsById(id string)(*i34835e4f604126539f2c916992cff760482a1c7bf03f1d5d4928a2099ac5183d.MeetingRegistrantRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["meetingRegistrant_id"] = id
    }
    return i34835e4f604126539f2c916992cff760482a1c7bf03f1d5d4928a2099ac5183d.NewMeetingRegistrantRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
