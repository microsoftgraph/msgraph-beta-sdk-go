package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i31a67f7bc2c45ba98338cd0edaf2867f688b9f08ec54f5d621086e29cbbbbbed "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onlinemeetings/item/alternativerecording"
    i325245ad62086769abb85404615f9161af71fd8a5187197632982e1dbb91a120 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onlinemeetings/item/recording"
    i5067990c676063eb8d3835deae020d7f6504ddad8405cb356682f6839dda6822 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onlinemeetings/item/meetingattendancereport"
    i65f8bf11ce791f596018398cdae14739b988a503719adcddd6149b76aa9ef7d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onlinemeetings/item/registration"
    i9fc238c5ea33c18d198b038ba9dc7600eb465be6a6ced5636f851fc096f6ea0d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onlinemeetings/item/attendancereports"
    ib5acddc80470812379ebb7f25e1a1c93a66d52604909cbf6927d11243ea3ef5c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onlinemeetings/item/attendeereport"
    i68329a1aa4c24d705e4cfdf254617301403e4cbff40bf8b7898a6b4a418033c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onlinemeetings/item/attendancereports/item"
)

// OnlineMeetingItemRequestBuilder provides operations to manage the onlineMeetings property of the microsoft.graph.user entity.
type OnlineMeetingItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OnlineMeetingItemRequestBuilderDeleteOptions options for Delete
type OnlineMeetingItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnlineMeetingItemRequestBuilderGetOptions options for Get
type OnlineMeetingItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OnlineMeetingItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnlineMeetingItemRequestBuilderGetQueryParameters get onlineMeetings from users
type OnlineMeetingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnlineMeetingItemRequestBuilderPatchOptions options for Patch
type OnlineMeetingItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnlineMeetingable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *OnlineMeetingItemRequestBuilder) AlternativeRecording()(*i31a67f7bc2c45ba98338cd0edaf2867f688b9f08ec54f5d621086e29cbbbbbed.AlternativeRecordingRequestBuilder) {
    return i31a67f7bc2c45ba98338cd0edaf2867f688b9f08ec54f5d621086e29cbbbbbed.NewAlternativeRecordingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnlineMeetingItemRequestBuilder) AttendanceReports()(*i9fc238c5ea33c18d198b038ba9dc7600eb465be6a6ced5636f851fc096f6ea0d.AttendanceReportsRequestBuilder) {
    return i9fc238c5ea33c18d198b038ba9dc7600eb465be6a6ced5636f851fc096f6ea0d.NewAttendanceReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttendanceReportsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.onlineMeetings.item.attendanceReports.item collection
func (m *OnlineMeetingItemRequestBuilder) AttendanceReportsById(id string)(*i68329a1aa4c24d705e4cfdf254617301403e4cbff40bf8b7898a6b4a418033c7.MeetingAttendanceReportItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["meetingAttendanceReport_id"] = id
    }
    return i68329a1aa4c24d705e4cfdf254617301403e4cbff40bf8b7898a6b4a418033c7.NewMeetingAttendanceReportItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnlineMeetingItemRequestBuilder) AttendeeReport()(*ib5acddc80470812379ebb7f25e1a1c93a66d52604909cbf6927d11243ea3ef5c.AttendeeReportRequestBuilder) {
    return ib5acddc80470812379ebb7f25e1a1c93a66d52604909cbf6927d11243ea3ef5c.NewAttendeeReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOnlineMeetingItemRequestBuilderInternal instantiates a new OnlineMeetingItemRequestBuilder and sets the default values.
func NewOnlineMeetingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnlineMeetingItemRequestBuilder) {
    m := &OnlineMeetingItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/onlineMeetings/{onlineMeeting_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnlineMeetingItemRequestBuilder instantiates a new OnlineMeetingItemRequestBuilder and sets the default values.
func NewOnlineMeetingItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnlineMeetingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlineMeetingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property onlineMeetings for users
func (m *OnlineMeetingItemRequestBuilder) CreateDeleteRequestInformation(options *OnlineMeetingItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get onlineMeetings from users
func (m *OnlineMeetingItemRequestBuilder) CreateGetRequestInformation(options *OnlineMeetingItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property onlineMeetings in users
func (m *OnlineMeetingItemRequestBuilder) CreatePatchRequestInformation(options *OnlineMeetingItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property onlineMeetings for users
func (m *OnlineMeetingItemRequestBuilder) Delete(options *OnlineMeetingItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get onlineMeetings from users
func (m *OnlineMeetingItemRequestBuilder) Get(options *OnlineMeetingItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnlineMeetingable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateOnlineMeetingFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnlineMeetingable), nil
}
func (m *OnlineMeetingItemRequestBuilder) MeetingAttendanceReport()(*i5067990c676063eb8d3835deae020d7f6504ddad8405cb356682f6839dda6822.MeetingAttendanceReportRequestBuilder) {
    return i5067990c676063eb8d3835deae020d7f6504ddad8405cb356682f6839dda6822.NewMeetingAttendanceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property onlineMeetings in users
func (m *OnlineMeetingItemRequestBuilder) Patch(options *OnlineMeetingItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *OnlineMeetingItemRequestBuilder) Recording()(*i325245ad62086769abb85404615f9161af71fd8a5187197632982e1dbb91a120.RecordingRequestBuilder) {
    return i325245ad62086769abb85404615f9161af71fd8a5187197632982e1dbb91a120.NewRecordingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnlineMeetingItemRequestBuilder) Registration()(*i65f8bf11ce791f596018398cdae14739b988a503719adcddd6149b76aa9ef7d7.RegistrationRequestBuilder) {
    return i65f8bf11ce791f596018398cdae14739b988a503719adcddd6149b76aa9ef7d7.NewRegistrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
