package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilder provides operations to manage the customQuestions property of the microsoft.graph.meetingRegistration entity.
type VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilderGetQueryParameters get a custom registration question associated with a meetingRegistration object on behalf of the organizer.
type VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilderGetQueryParameters
}
// VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMeetingRegistrationQuestionId provides operations to manage the customQuestions property of the microsoft.graph.meetingRegistration entity.
func (m *VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilder) ByMeetingRegistrationQuestionId(meetingRegistrationQuestionId string)(*VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if meetingRegistrationQuestionId != "" {
        urlTplParams["meetingRegistrationQuestion%2Did"] = meetingRegistrationQuestionId
    }
    return NewVirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilderInternal instantiates a new CustomQuestionsRequestBuilder and sets the default values.
func NewVirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilder) {
    m := &VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/events/{virtualEvent%2Did}/sessions/{virtualEventSession%2Did}/registration/customQuestions{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewVirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilder instantiates a new CustomQuestionsRequestBuilder and sets the default values.
func NewVirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilder) Count()(*VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsCountRequestBuilder) {
    return NewVirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a custom registration question associated with a meetingRegistration object on behalf of the organizer.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/meetingregistration-list-customquestions?view=graph-rest-1.0
func (m *VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingRegistrationQuestionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMeetingRegistrationQuestionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingRegistrationQuestionCollectionResponseable), nil
}
// Post create a custom registration question associated with a meetingRegistration object on behalf of the organizer.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/meetingregistration-post-customquestions?view=graph-rest-1.0
func (m *VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingRegistrationQuestionable, requestConfiguration *VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingRegistrationQuestionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMeetingRegistrationQuestionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingRegistrationQuestionable), nil
}
// ToGetRequestInformation get a custom registration question associated with a meetingRegistration object on behalf of the organizer.
func (m *VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPostRequestInformation create a custom registration question associated with a meetingRegistration object on behalf of the organizer.
func (m *VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingRegistrationQuestionable, requestConfiguration *VirtualEventsEventsItemSessionsItemRegistrationCustomQuestionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
