package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i28844d8a5d325bb0845a1778d6df0388d975703f04633f34ffc2f4073b871389 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings/item/registration"
    i3a2bfcc71796b9365b1b580ed5a11589b09828c31bbc6758b6afbd6ee2ee53a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings/item/attendeereport"
    i973f10e60f8f0226efcd773b4162b5082a7c2927993ac09fbb3f58bd0fb4fa4d "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings/item/attendancereports"
    ia176566273a5afb80848389bc0868b65818d8d6d3ab8bc727ddcd5d4015393f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings/item/recording"
    ie765355a32ec6cb4ba90da216a62ec4fb6b7b232eb3780cefe410dc49551bbe5 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings/item/meetingattendancereport"
    ifa2556b906140d07529e0b82dd914940679909ab68a2bc38a29559eceb0566e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings/item/alternativerecording"
    ib99f82b2686be9ab3c0709927e690dd3df4c41d545906785c734827457ee1d56 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings/item/attendancereports/item"
)

// OnlineMeetingRequestBuilder builds and executes requests for operations under \communications\onlineMeetings\{onlineMeeting-id}
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
// OnlineMeetingRequestBuilderGetQueryParameters get onlineMeetings from communications
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
func (m *OnlineMeetingRequestBuilder) AlternativeRecording()(*ifa2556b906140d07529e0b82dd914940679909ab68a2bc38a29559eceb0566e9.AlternativeRecordingRequestBuilder) {
    return ifa2556b906140d07529e0b82dd914940679909ab68a2bc38a29559eceb0566e9.NewAlternativeRecordingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnlineMeetingRequestBuilder) AttendanceReports()(*i973f10e60f8f0226efcd773b4162b5082a7c2927993ac09fbb3f58bd0fb4fa4d.AttendanceReportsRequestBuilder) {
    return i973f10e60f8f0226efcd773b4162b5082a7c2927993ac09fbb3f58bd0fb4fa4d.NewAttendanceReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttendanceReportsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.communications.onlineMeetings.item.attendanceReports.item collection
func (m *OnlineMeetingRequestBuilder) AttendanceReportsById(id string)(*ib99f82b2686be9ab3c0709927e690dd3df4c41d545906785c734827457ee1d56.MeetingAttendanceReportRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["meetingAttendanceReport_id"] = id
    }
    return ib99f82b2686be9ab3c0709927e690dd3df4c41d545906785c734827457ee1d56.NewMeetingAttendanceReportRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnlineMeetingRequestBuilder) AttendeeReport()(*i3a2bfcc71796b9365b1b580ed5a11589b09828c31bbc6758b6afbd6ee2ee53a9.AttendeeReportRequestBuilder) {
    return i3a2bfcc71796b9365b1b580ed5a11589b09828c31bbc6758b6afbd6ee2ee53a9.NewAttendeeReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOnlineMeetingRequestBuilderInternal instantiates a new OnlineMeetingRequestBuilder and sets the default values.
func NewOnlineMeetingRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnlineMeetingRequestBuilder) {
    m := &OnlineMeetingRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/onlineMeetings/{onlineMeeting_id}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property onlineMeetings for communications
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
// CreateGetRequestInformation get onlineMeetings from communications
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
// CreatePatchRequestInformation update the navigation property onlineMeetings in communications
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
// Delete delete navigation property onlineMeetings for communications
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
// Get get onlineMeetings from communications
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
func (m *OnlineMeetingRequestBuilder) MeetingAttendanceReport()(*ie765355a32ec6cb4ba90da216a62ec4fb6b7b232eb3780cefe410dc49551bbe5.MeetingAttendanceReportRequestBuilder) {
    return ie765355a32ec6cb4ba90da216a62ec4fb6b7b232eb3780cefe410dc49551bbe5.NewMeetingAttendanceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property onlineMeetings in communications
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
func (m *OnlineMeetingRequestBuilder) Recording()(*ia176566273a5afb80848389bc0868b65818d8d6d3ab8bc727ddcd5d4015393f9.RecordingRequestBuilder) {
    return ia176566273a5afb80848389bc0868b65818d8d6d3ab8bc727ddcd5d4015393f9.NewRecordingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnlineMeetingRequestBuilder) Registration()(*i28844d8a5d325bb0845a1778d6df0388d975703f04633f34ffc2f4073b871389.RegistrationRequestBuilder) {
    return i28844d8a5d325bb0845a1778d6df0388d975703f04633f34ffc2f4073b871389.NewRegistrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
