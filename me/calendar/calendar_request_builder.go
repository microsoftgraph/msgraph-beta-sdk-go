package calendar

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2f1046c95e29e9dc0a0e26221e25babc15231d926fd21540bc4f8be55206acba "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarpermissions"
    i310a0918126f2fd1a54cb6d55a6aa248b29b49ce4b8f6944a5a18dfe96111178 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events"
    i84f41c2d8a59ad72a2557098f3b161d03f5d8ac8e3987507b0fa7ecdf05601bf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/multivalueextendedproperties"
    i85085bfff31dfdfaf3940c26b7428d4dd4f5f214e1afb26f8ff92de42666214c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/allowedcalendarsharingroleswithuser"
    ia8c270bfd48992dd6d8586b4eb88c58708cf3fbfc388c4660cde8b22341c2d61 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview"
    iafafdd8f3e81db193ebe1f176611ba8d070c598779572146f18a7a2320d2b8fc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/singlevalueextendedproperties"
    id1ef0f108571f18aefac2c61a51615b16187fb53daeb8778a7d3fc99ee5b1a30 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/getschedule"
    i1339b57f4311b596f1e163d07c2af7ba9105f71396c487bac3b1051937b9a343 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/singlevalueextendedproperties/item"
    i32007971daba2b6bfd608251e3fa3037f0c6076b83bf11d1cebc0be6d191b58f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item"
    i7923064a9f4c4b75d44193b15f31eea4e3de5a03a690127e925545ca3fcb8ce5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item"
    icc49a296c7df21a7cc4befcd6f1a34f299f25134015701387f39f2be77ec1b62 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarpermissions/item"
    iec1f5e677b964978263004630240f212d437a10091543ccef2397757ae2b5c2e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/multivalueextendedproperties/item"
)

// CalendarRequestBuilder provides operations to manage the calendar property of the microsoft.graph.user entity.
type CalendarRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CalendarRequestBuilderDeleteOptions options for Delete
type CalendarRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// CalendarRequestBuilderGetOptions options for Get
type CalendarRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CalendarRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// CalendarRequestBuilderGetQueryParameters the user's primary calendar. Read-only.
type CalendarRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CalendarRequestBuilderPatchOptions options for Patch
type CalendarRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Calendarable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AllowedCalendarSharingRolesWithUser provides operations to call the allowedCalendarSharingRoles method.
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*i85085bfff31dfdfaf3940c26b7428d4dd4f5f214e1afb26f8ff92de42666214c.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return i85085bfff31dfdfaf3940c26b7428d4dd4f5f214e1afb26f8ff92de42666214c.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
// CalendarPermissions the calendarPermissions property
func (m *CalendarRequestBuilder) CalendarPermissions()(*i2f1046c95e29e9dc0a0e26221e25babc15231d926fd21540bc4f8be55206acba.CalendarPermissionsRequestBuilder) {
    return i2f1046c95e29e9dc0a0e26221e25babc15231d926fd21540bc4f8be55206acba.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarPermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarPermissions.item collection
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*icc49a296c7df21a7cc4befcd6f1a34f299f25134015701387f39f2be77ec1b62.CalendarPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission%2Did"] = id
    }
    return icc49a296c7df21a7cc4befcd6f1a34f299f25134015701387f39f2be77ec1b62.NewCalendarPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CalendarView the calendarView property
func (m *CalendarRequestBuilder) CalendarView()(*ia8c270bfd48992dd6d8586b4eb88c58708cf3fbfc388c4660cde8b22341c2d61.CalendarViewRequestBuilder) {
    return ia8c270bfd48992dd6d8586b4eb88c58708cf3fbfc388c4660cde8b22341c2d61.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item collection
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*i32007971daba2b6bfd608251e3fa3037f0c6076b83bf11d1cebc0be6d191b58f.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return i32007971daba2b6bfd608251e3fa3037f0c6076b83bf11d1cebc0be6d191b58f.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCalendarRequestBuilderInternal instantiates a new CalendarRequestBuilder and sets the default values.
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendar{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCalendarRequestBuilder instantiates a new CalendarRequestBuilder and sets the default values.
func NewCalendarRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property calendar for me
func (m *CalendarRequestBuilder) CreateDeleteRequestInformation(options *CalendarRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the user's primary calendar. Read-only.
func (m *CalendarRequestBuilder) CreateGetRequestInformation(options *CalendarRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property calendar in me
func (m *CalendarRequestBuilder) CreatePatchRequestInformation(options *CalendarRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property calendar for me
func (m *CalendarRequestBuilder) Delete(options *CalendarRequestBuilderDeleteOptions)(error) {
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
// Events the events property
func (m *CalendarRequestBuilder) Events()(*i310a0918126f2fd1a54cb6d55a6aa248b29b49ce4b8f6944a5a18dfe96111178.EventsRequestBuilder) {
    return i310a0918126f2fd1a54cb6d55a6aa248b29b49ce4b8f6944a5a18dfe96111178.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.events.item collection
func (m *CalendarRequestBuilder) EventsById(id string)(*i7923064a9f4c4b75d44193b15f31eea4e3de5a03a690127e925545ca3fcb8ce5.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return i7923064a9f4c4b75d44193b15f31eea4e3de5a03a690127e925545ca3fcb8ce5.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the user's primary calendar. Read-only.
func (m *CalendarRequestBuilder) Get(options *CalendarRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Calendarable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCalendarFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Calendarable), nil
}
// GetSchedule the getSchedule property
func (m *CalendarRequestBuilder) GetSchedule()(*id1ef0f108571f18aefac2c61a51615b16187fb53daeb8778a7d3fc99ee5b1a30.GetScheduleRequestBuilder) {
    return id1ef0f108571f18aefac2c61a51615b16187fb53daeb8778a7d3fc99ee5b1a30.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*i84f41c2d8a59ad72a2557098f3b161d03f5d8ac8e3987507b0fa7ecdf05601bf.MultiValueExtendedPropertiesRequestBuilder) {
    return i84f41c2d8a59ad72a2557098f3b161d03f5d8ac8e3987507b0fa7ecdf05601bf.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.multiValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iec1f5e677b964978263004630240f212d437a10091543ccef2397757ae2b5c2e.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return iec1f5e677b964978263004630240f212d437a10091543ccef2397757ae2b5c2e.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calendar in me
func (m *CalendarRequestBuilder) Patch(options *CalendarRequestBuilderPatchOptions)(error) {
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
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*iafafdd8f3e81db193ebe1f176611ba8d070c598779572146f18a7a2320d2b8fc.SingleValueExtendedPropertiesRequestBuilder) {
    return iafafdd8f3e81db193ebe1f176611ba8d070c598779572146f18a7a2320d2b8fc.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.singleValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i1339b57f4311b596f1e163d07c2af7ba9105f71396c487bac3b1051937b9a343.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i1339b57f4311b596f1e163d07c2af7ba9105f71396c487bac3b1051937b9a343.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
