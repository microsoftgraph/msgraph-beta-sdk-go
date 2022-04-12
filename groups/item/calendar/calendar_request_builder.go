package calendar

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i06cb71e1758c94dc508ff81b35d70e6b87e40dc3c43fc277a55fc5b07c51c025 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events"
    i4e394637a163af92a315f34da04f62ec4dc1d045e17f3716acaced9ea35ad045 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/getschedule"
    ib7ee09a2619abb2c78701384def8fcdab068e606d667dea6587234b16d250e88 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview"
    ic79a06981e01676a5b41be5ce0766fd4aeeb0f4a22bbea5fca16431948ae127a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarpermissions"
    ic906c3aa3cff8041006e1db1afa8c5ecce79b293f1b9649a48c05726e6f788f8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/multivalueextendedproperties"
    icd7a045db01e44db897afd8ba5343e7d38c1599413b6a17400da01b527973e57 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/allowedcalendarsharingroleswithuser"
    ie357553e3f2994fd5e3a6ab96e22c5c927c07c3e3c999ee4350aca9434117b9c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/singlevalueextendedproperties"
    i2877c9d97209d99b4d909de67ba8cd762e6a506898bb79d3f55b0659900ca31c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item"
    i2b0a7eca3f232e3b28e09ce2398345a2762596c5fc5f1b3b203b447ca9d62dcd "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/multivalueextendedproperties/item"
    i6389000acb4a221ddef87343d30894dc16987ef37595f24421ad461fe767083f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarpermissions/item"
    i7135f234b9b77ac7d1f38566664eb079da2d1647d2b78e7e1558a59a4f6a8ae5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/singlevalueextendedproperties/item"
    iafdc3a5c49cb6bb0133e40dda5d4c139576d35278bf232fa62de70f01714f934 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item"
)

// CalendarRequestBuilder provides operations to manage the calendar property of the microsoft.graph.group entity.
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
// CalendarRequestBuilderGetQueryParameters the group's calendar. Read-only.
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
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*icd7a045db01e44db897afd8ba5343e7d38c1599413b6a17400da01b527973e57.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return icd7a045db01e44db897afd8ba5343e7d38c1599413b6a17400da01b527973e57.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
// CalendarPermissions the calendarPermissions property
func (m *CalendarRequestBuilder) CalendarPermissions()(*ic79a06981e01676a5b41be5ce0766fd4aeeb0f4a22bbea5fca16431948ae127a.CalendarPermissionsRequestBuilder) {
    return ic79a06981e01676a5b41be5ce0766fd4aeeb0f4a22bbea5fca16431948ae127a.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarPermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarPermissions.item collection
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*i6389000acb4a221ddef87343d30894dc16987ef37595f24421ad461fe767083f.CalendarPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission%2Did"] = id
    }
    return i6389000acb4a221ddef87343d30894dc16987ef37595f24421ad461fe767083f.NewCalendarPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CalendarView the calendarView property
func (m *CalendarRequestBuilder) CalendarView()(*ib7ee09a2619abb2c78701384def8fcdab068e606d667dea6587234b16d250e88.CalendarViewRequestBuilder) {
    return ib7ee09a2619abb2c78701384def8fcdab068e606d667dea6587234b16d250e88.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item collection
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*i2877c9d97209d99b4d909de67ba8cd762e6a506898bb79d3f55b0659900ca31c.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return i2877c9d97209d99b4d909de67ba8cd762e6a506898bb79d3f55b0659900ca31c.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCalendarRequestBuilderInternal instantiates a new CalendarRequestBuilder and sets the default values.
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendar{?%24select}";
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
// CreateDeleteRequestInformation delete navigation property calendar for groups
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
// CreateGetRequestInformation the group's calendar. Read-only.
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
// CreatePatchRequestInformation update the navigation property calendar in groups
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
// Delete delete navigation property calendar for groups
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
func (m *CalendarRequestBuilder) Events()(*i06cb71e1758c94dc508ff81b35d70e6b87e40dc3c43fc277a55fc5b07c51c025.EventsRequestBuilder) {
    return i06cb71e1758c94dc508ff81b35d70e6b87e40dc3c43fc277a55fc5b07c51c025.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item collection
func (m *CalendarRequestBuilder) EventsById(id string)(*iafdc3a5c49cb6bb0133e40dda5d4c139576d35278bf232fa62de70f01714f934.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return iafdc3a5c49cb6bb0133e40dda5d4c139576d35278bf232fa62de70f01714f934.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the group's calendar. Read-only.
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
func (m *CalendarRequestBuilder) GetSchedule()(*i4e394637a163af92a315f34da04f62ec4dc1d045e17f3716acaced9ea35ad045.GetScheduleRequestBuilder) {
    return i4e394637a163af92a315f34da04f62ec4dc1d045e17f3716acaced9ea35ad045.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*ic906c3aa3cff8041006e1db1afa8c5ecce79b293f1b9649a48c05726e6f788f8.MultiValueExtendedPropertiesRequestBuilder) {
    return ic906c3aa3cff8041006e1db1afa8c5ecce79b293f1b9649a48c05726e6f788f8.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.multiValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i2b0a7eca3f232e3b28e09ce2398345a2762596c5fc5f1b3b203b447ca9d62dcd.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i2b0a7eca3f232e3b28e09ce2398345a2762596c5fc5f1b3b203b447ca9d62dcd.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calendar in groups
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*ie357553e3f2994fd5e3a6ab96e22c5c927c07c3e3c999ee4350aca9434117b9c.SingleValueExtendedPropertiesRequestBuilder) {
    return ie357553e3f2994fd5e3a6ab96e22c5c927c07c3e3c999ee4350aca9434117b9c.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.singleValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i7135f234b9b77ac7d1f38566664eb079da2d1647d2b78e7e1558a59a4f6a8ae5.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i7135f234b9b77ac7d1f38566664eb079da2d1647d2b78e7e1558a59a4f6a8ae5.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
