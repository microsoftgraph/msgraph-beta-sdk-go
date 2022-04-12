package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
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
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OnlineMeetingItemRequestBuilderDeleteOptions options for Delete
type OnlineMeetingItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// OnlineMeetingItemRequestBuilderGetOptions options for Get
type OnlineMeetingItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnlineMeetingItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// OnlineMeetingItemRequestBuilderGetQueryParameters get onlineMeetings from users
type OnlineMeetingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OnlineMeetingItemRequestBuilderPatchOptions options for Patch
type OnlineMeetingItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AlternativeRecording the alternativeRecording property
func (m *OnlineMeetingItemRequestBuilder) AlternativeRecording()(*i31a67f7bc2c45ba98338cd0edaf2867f688b9f08ec54f5d621086e29cbbbbbed.AlternativeRecordingRequestBuilder) {
    return i31a67f7bc2c45ba98338cd0edaf2867f688b9f08ec54f5d621086e29cbbbbbed.NewAlternativeRecordingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttendanceReports the attendanceReports property
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
        urlTplParams["meetingAttendanceReport%2Did"] = id
    }
    return i68329a1aa4c24d705e4cfdf254617301403e4cbff40bf8b7898a6b4a418033c7.NewMeetingAttendanceReportItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AttendeeReport the attendeeReport property
func (m *OnlineMeetingItemRequestBuilder) AttendeeReport()(*ib5acddc80470812379ebb7f25e1a1c93a66d52604909cbf6927d11243ea3ef5c.AttendeeReportRequestBuilder) {
    return ib5acddc80470812379ebb7f25e1a1c93a66d52604909cbf6927d11243ea3ef5c.NewAttendeeReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOnlineMeetingItemRequestBuilderInternal instantiates a new OnlineMeetingItemRequestBuilder and sets the default values.
func NewOnlineMeetingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingItemRequestBuilder) {
    m := &OnlineMeetingItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/onlineMeetings/{onlineMeeting%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnlineMeetingItemRequestBuilder instantiates a new OnlineMeetingItemRequestBuilder and sets the default values.
func NewOnlineMeetingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlineMeetingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property onlineMeetings for users
func (m *OnlineMeetingItemRequestBuilder) CreateDeleteRequestInformation(options *OnlineMeetingItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get onlineMeetings from users
func (m *OnlineMeetingItemRequestBuilder) CreateGetRequestInformation(options *OnlineMeetingItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property onlineMeetings in users
func (m *OnlineMeetingItemRequestBuilder) CreatePatchRequestInformation(options *OnlineMeetingItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
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
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get onlineMeetings from users
func (m *OnlineMeetingItemRequestBuilder) Get(options *OnlineMeetingItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnlineMeetingFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable), nil
}
// MeetingAttendanceReport the meetingAttendanceReport property
func (m *OnlineMeetingItemRequestBuilder) MeetingAttendanceReport()(*i5067990c676063eb8d3835deae020d7f6504ddad8405cb356682f6839dda6822.MeetingAttendanceReportRequestBuilder) {
    return i5067990c676063eb8d3835deae020d7f6504ddad8405cb356682f6839dda6822.NewMeetingAttendanceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property onlineMeetings in users
func (m *OnlineMeetingItemRequestBuilder) Patch(options *OnlineMeetingItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Recording the recording property
func (m *OnlineMeetingItemRequestBuilder) Recording()(*i325245ad62086769abb85404615f9161af71fd8a5187197632982e1dbb91a120.RecordingRequestBuilder) {
    return i325245ad62086769abb85404615f9161af71fd8a5187197632982e1dbb91a120.NewRecordingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Registration the registration property
func (m *OnlineMeetingItemRequestBuilder) Registration()(*i65f8bf11ce791f596018398cdae14739b988a503719adcddd6149b76aa9ef7d7.RegistrationRequestBuilder) {
    return i65f8bf11ce791f596018398cdae14739b988a503719adcddd6149b76aa9ef7d7.NewRegistrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
