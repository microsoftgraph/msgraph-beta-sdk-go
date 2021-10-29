package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5792a57d1b509cf5853eddfc75cdd3d12d7fffac6c1d1ec273e69dfc55edec5d "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/attendeereport"
    i621490855f8aac21cf8d7032977ac3e4dd93569a78b9ec995379c44e1ed33eaf "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/alternativerecording"
    ia44b3dfc614f798ddda4e107c7b00a962e62b26be06acb8ba2c30fda4607647a "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/meetingattendancereport"
    ic4cea0d5545c1ec81a2c193eb0ee5d7a032aa549a87e4b5ea7d5e3d4250137fc "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/recording"
    ifc43ad7bd5674ca951e2271a8b4d99b63fa47b18f4fca9e422daedf0278acaf6 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/registration"
)

// Builds and executes requests for operations under \app\onlineMeetings\{onlineMeeting-id}
type OnlineMeetingRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type OnlineMeetingRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type OnlineMeetingRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OnlineMeetingRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get onlineMeetings from app
type OnlineMeetingRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type OnlineMeetingRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnlineMeeting;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *OnlineMeetingRequestBuilder) AlternativeRecording()(*i621490855f8aac21cf8d7032977ac3e4dd93569a78b9ec995379c44e1ed33eaf.AlternativeRecordingRequestBuilder) {
    return i621490855f8aac21cf8d7032977ac3e4dd93569a78b9ec995379c44e1ed33eaf.NewAlternativeRecordingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnlineMeetingRequestBuilder) AttendeeReport()(*i5792a57d1b509cf5853eddfc75cdd3d12d7fffac6c1d1ec273e69dfc55edec5d.AttendeeReportRequestBuilder) {
    return i5792a57d1b509cf5853eddfc75cdd3d12d7fffac6c1d1ec273e69dfc55edec5d.NewAttendeeReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new OnlineMeetingRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOnlineMeetingRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnlineMeetingRequestBuilder) {
    m := &OnlineMeetingRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/app/onlineMeetings/{onlineMeeting_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new OnlineMeetingRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOnlineMeetingRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnlineMeetingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlineMeetingRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete navigation property onlineMeetings for app
// Parameters:
//  - options : Options for the request
func (m *OnlineMeetingRequestBuilder) CreateDeleteRequestInformation(options *OnlineMeetingRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get onlineMeetings from app
// Parameters:
//  - options : Options for the request
func (m *OnlineMeetingRequestBuilder) CreateGetRequestInformation(options *OnlineMeetingRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Update the navigation property onlineMeetings in app
// Parameters:
//  - options : Options for the request
func (m *OnlineMeetingRequestBuilder) CreatePatchRequestInformation(options *OnlineMeetingRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete navigation property onlineMeetings for app
// Parameters:
//  - options : Options for the request
func (m *OnlineMeetingRequestBuilder) Delete(options *OnlineMeetingRequestBuilderDeleteOptions)(error) {
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
// Get onlineMeetings from app
// Parameters:
//  - options : Options for the request
func (m *OnlineMeetingRequestBuilder) Get(options *OnlineMeetingRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnlineMeeting, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOnlineMeeting() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnlineMeeting), nil
}
func (m *OnlineMeetingRequestBuilder) MeetingAttendanceReport()(*ia44b3dfc614f798ddda4e107c7b00a962e62b26be06acb8ba2c30fda4607647a.MeetingAttendanceReportRequestBuilder) {
    return ia44b3dfc614f798ddda4e107c7b00a962e62b26be06acb8ba2c30fda4607647a.NewMeetingAttendanceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Update the navigation property onlineMeetings in app
// Parameters:
//  - options : Options for the request
func (m *OnlineMeetingRequestBuilder) Patch(options *OnlineMeetingRequestBuilderPatchOptions)(error) {
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
func (m *OnlineMeetingRequestBuilder) Recording()(*ic4cea0d5545c1ec81a2c193eb0ee5d7a032aa549a87e4b5ea7d5e3d4250137fc.RecordingRequestBuilder) {
    return ic4cea0d5545c1ec81a2c193eb0ee5d7a032aa549a87e4b5ea7d5e3d4250137fc.NewRecordingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnlineMeetingRequestBuilder) Registration()(*ifc43ad7bd5674ca951e2271a8b4d99b63fa47b18f4fca9e422daedf0278acaf6.RegistrationRequestBuilder) {
    return ifc43ad7bd5674ca951e2271a8b4d99b63fa47b18f4fca9e422daedf0278acaf6.NewRegistrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
