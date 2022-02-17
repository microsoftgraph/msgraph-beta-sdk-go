package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i36d7f6a1b1906c3cdb16b4ede2175f668fa5388566026ae225a46e8d0be75ba4 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/attendancereports"
    i5792a57d1b509cf5853eddfc75cdd3d12d7fffac6c1d1ec273e69dfc55edec5d "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/attendeereport"
    i621490855f8aac21cf8d7032977ac3e4dd93569a78b9ec995379c44e1ed33eaf "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/alternativerecording"
    ia44b3dfc614f798ddda4e107c7b00a962e62b26be06acb8ba2c30fda4607647a "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/meetingattendancereport"
    ic4cea0d5545c1ec81a2c193eb0ee5d7a032aa549a87e4b5ea7d5e3d4250137fc "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/recording"
    ifc43ad7bd5674ca951e2271a8b4d99b63fa47b18f4fca9e422daedf0278acaf6 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/registration"
    i29fda42a3d65f7aed1e8a1e9c5006e031671b1aed329c71b8d3daa638210ed5a "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item/attendancereports/item"
)

// OnlineMeetingRequestBuilder builds and executes requests for operations under \app\onlineMeetings\{onlineMeeting-id}
type OnlineMeetingRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OnlineMeetingRequestBuilderDeleteOptions options for Delete
type OnlineMeetingRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnlineMeetingRequestBuilderGetOptions options for Get
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
// OnlineMeetingRequestBuilderGetQueryParameters get onlineMeetings from app
type OnlineMeetingRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnlineMeetingRequestBuilderPatchOptions options for Patch
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
func (m *OnlineMeetingRequestBuilder) AttendanceReports()(*i36d7f6a1b1906c3cdb16b4ede2175f668fa5388566026ae225a46e8d0be75ba4.AttendanceReportsRequestBuilder) {
    return i36d7f6a1b1906c3cdb16b4ede2175f668fa5388566026ae225a46e8d0be75ba4.NewAttendanceReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttendanceReportsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.app.onlineMeetings.item.attendanceReports.item collection
func (m *OnlineMeetingRequestBuilder) AttendanceReportsById(id string)(*i29fda42a3d65f7aed1e8a1e9c5006e031671b1aed329c71b8d3daa638210ed5a.MeetingAttendanceReportRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["meetingAttendanceReport_id"] = id
    }
    return i29fda42a3d65f7aed1e8a1e9c5006e031671b1aed329c71b8d3daa638210ed5a.NewMeetingAttendanceReportRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnlineMeetingRequestBuilder) AttendeeReport()(*i5792a57d1b509cf5853eddfc75cdd3d12d7fffac6c1d1ec273e69dfc55edec5d.AttendeeReportRequestBuilder) {
    return i5792a57d1b509cf5853eddfc75cdd3d12d7fffac6c1d1ec273e69dfc55edec5d.NewAttendeeReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOnlineMeetingRequestBuilderInternal instantiates a new OnlineMeetingRequestBuilder and sets the default values.
func NewOnlineMeetingRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnlineMeetingRequestBuilder) {
    m := &OnlineMeetingRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/app/onlineMeetings/{onlineMeeting_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnlineMeetingRequestBuilder instantiates a new OnlineMeetingRequestBuilder and sets the default values.
func NewOnlineMeetingRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnlineMeetingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlineMeetingRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property onlineMeetings for app
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
// CreateGetRequestInformation get onlineMeetings from app
func (m *OnlineMeetingRequestBuilder) CreateGetRequestInformation(options *OnlineMeetingRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property onlineMeetings in app
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
// Delete delete navigation property onlineMeetings for app
func (m *OnlineMeetingRequestBuilder) Delete(options *OnlineMeetingRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get get onlineMeetings from app
func (m *OnlineMeetingRequestBuilder) Get(options *OnlineMeetingRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnlineMeeting, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOnlineMeeting() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnlineMeeting), nil
}
func (m *OnlineMeetingRequestBuilder) MeetingAttendanceReport()(*ia44b3dfc614f798ddda4e107c7b00a962e62b26be06acb8ba2c30fda4607647a.MeetingAttendanceReportRequestBuilder) {
    return ia44b3dfc614f798ddda4e107c7b00a962e62b26be06acb8ba2c30fda4607647a.NewMeetingAttendanceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property onlineMeetings in app
func (m *OnlineMeetingRequestBuilder) Patch(options *OnlineMeetingRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
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
