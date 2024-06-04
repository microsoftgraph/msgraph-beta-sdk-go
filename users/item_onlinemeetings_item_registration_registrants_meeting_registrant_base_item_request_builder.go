package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilder provides operations to manage the registrants property of the microsoft.graph.meetingRegistrationBase entity.
type ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilderGetQueryParameters registrants of the online meeting.
type ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilderGetQueryParameters
}
// ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilderInternal instantiates a new ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilder and sets the default values.
func NewItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilder) {
    m := &ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/onlineMeetings/{onlineMeeting%2Did}/registration/registrants/{meetingRegistrantBase%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilder instantiates a new ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilder and sets the default values.
func NewItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete cancel an onlineMeeting registration for a meetingRegistrant on behalf of the registrant. Only use this method when the allowedRegistrant property of the meetingRegistration object has a value of organization and the registrant's delegated permission was used to enroll. When the allowedRegistrant value is everyone, registrants can only use the link in the email they receive to cancel their registration.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/meetingregistrant-delete?view=graph-rest-beta
func (m *ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get registrants of the online meeting.
// returns a MeetingRegistrantBaseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingRegistrantBaseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMeetingRegistrantBaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingRegistrantBaseable), nil
}
// Patch update the navigation property registrants in users
// returns a MeetingRegistrantBaseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingRegistrantBaseable, requestConfiguration *ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingRegistrantBaseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMeetingRegistrantBaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingRegistrantBaseable), nil
}
// ToDeleteRequestInformation cancel an onlineMeeting registration for a meetingRegistrant on behalf of the registrant. Only use this method when the allowedRegistrant property of the meetingRegistration object has a value of organization and the registrant's delegated permission was used to enroll. When the allowedRegistrant value is everyone, registrants can only use the link in the email they receive to cancel their registration.
// returns a *RequestInformation when successful
func (m *ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation registrants of the online meeting.
// returns a *RequestInformation when successful
func (m *ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property registrants in users
// returns a *RequestInformation when successful
func (m *ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingRegistrantBaseable, requestConfiguration *ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilder when successful
func (m *ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilder) WithUrl(rawUrl string)(*ItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilder) {
    return NewItemOnlinemeetingsItemRegistrationRegistrantsMeetingRegistrantBaseItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
