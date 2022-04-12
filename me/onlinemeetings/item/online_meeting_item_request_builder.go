package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2a0e0d48f1a133f80557e834fceb9001d321597582147fa92c8d557b1e60c722 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings/item/recording"
    i47c2284118cc15348aa203315c0a4637fa39b84c65e76ecdc39fc175c32c90b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings/item/meetingattendancereport"
    i520e7c5dd1b2e99f3cf76ee8e21c2e881be13a393c94622086fa442c45f27e8f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings/item/attendancereports"
    i60f747f04da18ee740c2d7cec8aa333e174de51a71c041da85eebd852e1da259 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings/item/registration"
    i9751454805a6089ded23da943214fdbe9028e1de749d44a87219eee6d363460e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings/item/alternativerecording"
    i98d7ddf881a02ae894ac70424b059f80f8d245996309997a347c38739b78710b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings/item/attendeereport"
    id7c487e4aca56b8a93bfe6b7467589181f3783dae6bf5fc07a9339b86fdc342e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings/item/attendancereports/item"
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
// OnlineMeetingItemRequestBuilderGetQueryParameters get onlineMeetings from me
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
func (m *OnlineMeetingItemRequestBuilder) AlternativeRecording()(*i9751454805a6089ded23da943214fdbe9028e1de749d44a87219eee6d363460e.AlternativeRecordingRequestBuilder) {
    return i9751454805a6089ded23da943214fdbe9028e1de749d44a87219eee6d363460e.NewAlternativeRecordingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttendanceReports the attendanceReports property
func (m *OnlineMeetingItemRequestBuilder) AttendanceReports()(*i520e7c5dd1b2e99f3cf76ee8e21c2e881be13a393c94622086fa442c45f27e8f.AttendanceReportsRequestBuilder) {
    return i520e7c5dd1b2e99f3cf76ee8e21c2e881be13a393c94622086fa442c45f27e8f.NewAttendanceReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttendanceReportsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.onlineMeetings.item.attendanceReports.item collection
func (m *OnlineMeetingItemRequestBuilder) AttendanceReportsById(id string)(*id7c487e4aca56b8a93bfe6b7467589181f3783dae6bf5fc07a9339b86fdc342e.MeetingAttendanceReportItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["meetingAttendanceReport%2Did"] = id
    }
    return id7c487e4aca56b8a93bfe6b7467589181f3783dae6bf5fc07a9339b86fdc342e.NewMeetingAttendanceReportItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AttendeeReport the attendeeReport property
func (m *OnlineMeetingItemRequestBuilder) AttendeeReport()(*i98d7ddf881a02ae894ac70424b059f80f8d245996309997a347c38739b78710b.AttendeeReportRequestBuilder) {
    return i98d7ddf881a02ae894ac70424b059f80f8d245996309997a347c38739b78710b.NewAttendeeReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOnlineMeetingItemRequestBuilderInternal instantiates a new OnlineMeetingItemRequestBuilder and sets the default values.
func NewOnlineMeetingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingItemRequestBuilder) {
    m := &OnlineMeetingItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/onlineMeetings/{onlineMeeting%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property onlineMeetings for me
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
// CreateGetRequestInformation get onlineMeetings from me
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
// CreatePatchRequestInformation update the navigation property onlineMeetings in me
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
// Delete delete navigation property onlineMeetings for me
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
// Get get onlineMeetings from me
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
func (m *OnlineMeetingItemRequestBuilder) MeetingAttendanceReport()(*i47c2284118cc15348aa203315c0a4637fa39b84c65e76ecdc39fc175c32c90b9.MeetingAttendanceReportRequestBuilder) {
    return i47c2284118cc15348aa203315c0a4637fa39b84c65e76ecdc39fc175c32c90b9.NewMeetingAttendanceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property onlineMeetings in me
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
func (m *OnlineMeetingItemRequestBuilder) Recording()(*i2a0e0d48f1a133f80557e834fceb9001d321597582147fa92c8d557b1e60c722.RecordingRequestBuilder) {
    return i2a0e0d48f1a133f80557e834fceb9001d321597582147fa92c8d557b1e60c722.NewRecordingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Registration the registration property
func (m *OnlineMeetingItemRequestBuilder) Registration()(*i60f747f04da18ee740c2d7cec8aa333e174de51a71c041da85eebd852e1da259.RegistrationRequestBuilder) {
    return i60f747f04da18ee740c2d7cec8aa333e174de51a71c041da85eebd852e1da259.NewRegistrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
