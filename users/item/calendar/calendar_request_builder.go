package calendar

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0d974efbbcde965f86cd784750b0d398768a58c4f0c31e8aa6eebab38714cfb3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview"
    i5f2b87b4b8ae2f7ac0f87df9716f56e4739dd16a5886917392d33a88c18c770c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events"
    i7f99ab356098627c58430897ce6db68b26eb164759df0d5c099175b468c773cd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/singlevalueextendedproperties"
    ia046aa77ed3651894558d377e9f2faf10b90f9fa6ea4b8d58fa81d2c0bf352b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarpermissions"
    ib11c18a31cdd7cea653dae6c94f9cb3c8880c5d2ff9d71db4311df3cb8ef6acb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/allowedcalendarsharingroleswithuser"
    ic0da2257b14ab755d200b83c581b125370f630995453afb2a431a7b9a660b4fb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/getschedule"
    ie4f67284862d68d4348bd087a24d1359caa53a620bec669a57587285e12ce7b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/multivalueextendedproperties"
    i35fa4f4f3eb3bdb367848efc1a7e0f1f004ad8371e5de0506e46fa00e172f96a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item"
    i646e54f4392ea806f149eff3b9db0f4c56b5fd5bf2023a287a9335a93382fc50 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/singlevalueextendedproperties/item"
    i6acf51572e8a0bbb972c27da38b03c2fa4c3df3e13192ade1cbb983349a39367 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item"
    i94baab4924959ccc5f343b01737d73626b9abc213582a99069414c9ca8dd6fc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarpermissions/item"
    ibe6df2514f80166133d14155ec9a4ce76559c25b036bd39eee4eb72b15967d12 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/multivalueextendedproperties/item"
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
// CalendarRequestBuilderGetQueryParameters get the properties and relationships of a calendar object. The calendar can be one for a user, or the default calendar of a Microsoft 365 group. There are two scenarios where an app can get another user's calendar:
type CalendarRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CalendarRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CalendarRequestBuilderGetQueryParameters
}
// CalendarRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllowedCalendarSharingRolesWithUser provides operations to call the allowedCalendarSharingRoles method.
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*ib11c18a31cdd7cea653dae6c94f9cb3c8880c5d2ff9d71db4311df3cb8ef6acb.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return ib11c18a31cdd7cea653dae6c94f9cb3c8880c5d2ff9d71db4311df3cb8ef6acb.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
// CalendarPermissions provides operations to manage the calendarPermissions property of the microsoft.graph.calendar entity.
func (m *CalendarRequestBuilder) CalendarPermissions()(*ia046aa77ed3651894558d377e9f2faf10b90f9fa6ea4b8d58fa81d2c0bf352b7.CalendarPermissionsRequestBuilder) {
    return ia046aa77ed3651894558d377e9f2faf10b90f9fa6ea4b8d58fa81d2c0bf352b7.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarPermissionsById provides operations to manage the calendarPermissions property of the microsoft.graph.calendar entity.
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*i94baab4924959ccc5f343b01737d73626b9abc213582a99069414c9ca8dd6fc2.CalendarPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission%2Did"] = id
    }
    return i94baab4924959ccc5f343b01737d73626b9abc213582a99069414c9ca8dd6fc2.NewCalendarPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CalendarView provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
func (m *CalendarRequestBuilder) CalendarView()(*i0d974efbbcde965f86cd784750b0d398768a58c4f0c31e8aa6eebab38714cfb3.CalendarViewRequestBuilder) {
    return i0d974efbbcde965f86cd784750b0d398768a58c4f0c31e8aa6eebab38714cfb3.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*i35fa4f4f3eb3bdb367848efc1a7e0f1f004ad8371e5de0506e46fa00e172f96a.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return i35fa4f4f3eb3bdb367848efc1a7e0f1f004ad8371e5de0506e46fa00e172f96a.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCalendarRequestBuilderInternal instantiates a new CalendarRequestBuilder and sets the default values.
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendar{?%24select}";
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
// CreateGetRequestInformation get the properties and relationships of a calendar object. The calendar can be one for a user, or the default calendar of a Microsoft 365 group. There are two scenarios where an app can get another user's calendar:
func (m *CalendarRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *CalendarRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the properties of a calendar object. The calendar can be one for a user, or the default calendar of a Microsoft 365 group.
func (m *CalendarRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Calendarable, requestConfiguration *CalendarRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Events provides operations to manage the events property of the microsoft.graph.calendar entity.
func (m *CalendarRequestBuilder) Events()(*i5f2b87b4b8ae2f7ac0f87df9716f56e4739dd16a5886917392d33a88c18c770c.EventsRequestBuilder) {
    return i5f2b87b4b8ae2f7ac0f87df9716f56e4739dd16a5886917392d33a88c18c770c.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById provides operations to manage the events property of the microsoft.graph.calendar entity.
func (m *CalendarRequestBuilder) EventsById(id string)(*i6acf51572e8a0bbb972c27da38b03c2fa4c3df3e13192ade1cbb983349a39367.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return i6acf51572e8a0bbb972c27da38b03c2fa4c3df3e13192ade1cbb983349a39367.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get the properties and relationships of a calendar object. The calendar can be one for a user, or the default calendar of a Microsoft 365 group. There are two scenarios where an app can get another user's calendar:
func (m *CalendarRequestBuilder) Get(ctx context.Context, requestConfiguration *CalendarRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Calendarable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCalendarFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Calendarable), nil
}
// GetSchedule provides operations to call the getSchedule method.
func (m *CalendarRequestBuilder) GetSchedule()(*ic0da2257b14ab755d200b83c581b125370f630995453afb2a431a7b9a660b4fb.GetScheduleRequestBuilder) {
    return ic0da2257b14ab755d200b83c581b125370f630995453afb2a431a7b9a660b4fb.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.calendar entity.
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*ie4f67284862d68d4348bd087a24d1359caa53a620bec669a57587285e12ce7b3.MultiValueExtendedPropertiesRequestBuilder) {
    return ie4f67284862d68d4348bd087a24d1359caa53a620bec669a57587285e12ce7b3.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.calendar entity.
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ibe6df2514f80166133d14155ec9a4ce76559c25b036bd39eee4eb72b15967d12.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ibe6df2514f80166133d14155ec9a4ce76559c25b036bd39eee4eb72b15967d12.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the properties of a calendar object. The calendar can be one for a user, or the default calendar of a Microsoft 365 group.
func (m *CalendarRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Calendarable, requestConfiguration *CalendarRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Calendarable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCalendarFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Calendarable), nil
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.calendar entity.
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*i7f99ab356098627c58430897ce6db68b26eb164759df0d5c099175b468c773cd.SingleValueExtendedPropertiesRequestBuilder) {
    return i7f99ab356098627c58430897ce6db68b26eb164759df0d5c099175b468c773cd.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.calendar entity.
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i646e54f4392ea806f149eff3b9db0f4c56b5fd5bf2023a287a9335a93382fc50.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i646e54f4392ea806f149eff3b9db0f4c56b5fd5bf2023a287a9335a93382fc50.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
