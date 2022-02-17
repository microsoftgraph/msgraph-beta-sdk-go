package meetingattendancereport

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i4ba55ed63fb75b5bb5dd1bfef40702604868b33bbe659b17afbc2c5f7b8861fe "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings/item/meetingattendancereport/attendancerecords"
    ifde72711ce2bcb491f303f9337972e86208ba1236f423b2f91521635edc8ab24 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/onlinemeetings/item/meetingattendancereport/attendancerecords/item"
)

// MeetingAttendanceReportRequestBuilder builds and executes requests for operations under \communications\onlineMeetings\{onlineMeeting-id}\meetingAttendanceReport
type MeetingAttendanceReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MeetingAttendanceReportRequestBuilderDeleteOptions options for Delete
type MeetingAttendanceReportRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MeetingAttendanceReportRequestBuilderGetOptions options for Get
type MeetingAttendanceReportRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MeetingAttendanceReportRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MeetingAttendanceReportRequestBuilderGetQueryParameters get meetingAttendanceReport from communications
type MeetingAttendanceReportRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// MeetingAttendanceReportRequestBuilderPatchOptions options for Patch
type MeetingAttendanceReportRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MeetingAttendanceReport;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *MeetingAttendanceReportRequestBuilder) AttendanceRecords()(*i4ba55ed63fb75b5bb5dd1bfef40702604868b33bbe659b17afbc2c5f7b8861fe.AttendanceRecordsRequestBuilder) {
    return i4ba55ed63fb75b5bb5dd1bfef40702604868b33bbe659b17afbc2c5f7b8861fe.NewAttendanceRecordsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttendanceRecordsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.communications.onlineMeetings.item.meetingAttendanceReport.attendanceRecords.item collection
func (m *MeetingAttendanceReportRequestBuilder) AttendanceRecordsById(id string)(*ifde72711ce2bcb491f303f9337972e86208ba1236f423b2f91521635edc8ab24.AttendanceRecordRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attendanceRecord_id"] = id
    }
    return ifde72711ce2bcb491f303f9337972e86208ba1236f423b2f91521635edc8ab24.NewAttendanceRecordRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMeetingAttendanceReportRequestBuilderInternal instantiates a new MeetingAttendanceReportRequestBuilder and sets the default values.
func NewMeetingAttendanceReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MeetingAttendanceReportRequestBuilder) {
    m := &MeetingAttendanceReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/onlineMeetings/{onlineMeeting_id}/meetingAttendanceReport{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeetingAttendanceReportRequestBuilder instantiates a new MeetingAttendanceReportRequestBuilder and sets the default values.
func NewMeetingAttendanceReportRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MeetingAttendanceReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeetingAttendanceReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property meetingAttendanceReport for communications
func (m *MeetingAttendanceReportRequestBuilder) CreateDeleteRequestInformation(options *MeetingAttendanceReportRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get meetingAttendanceReport from communications
func (m *MeetingAttendanceReportRequestBuilder) CreateGetRequestInformation(options *MeetingAttendanceReportRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property meetingAttendanceReport in communications
func (m *MeetingAttendanceReportRequestBuilder) CreatePatchRequestInformation(options *MeetingAttendanceReportRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property meetingAttendanceReport for communications
func (m *MeetingAttendanceReportRequestBuilder) Delete(options *MeetingAttendanceReportRequestBuilderDeleteOptions)(error) {
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
// Get get meetingAttendanceReport from communications
func (m *MeetingAttendanceReportRequestBuilder) Get(options *MeetingAttendanceReportRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MeetingAttendanceReport, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewMeetingAttendanceReport() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MeetingAttendanceReport), nil
}
// Patch update the navigation property meetingAttendanceReport in communications
func (m *MeetingAttendanceReportRequestBuilder) Patch(options *MeetingAttendanceReportRequestBuilderPatchOptions)(error) {
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
