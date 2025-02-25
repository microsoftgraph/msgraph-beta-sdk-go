package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilder provides operations to manage the customQuestions property of the microsoft.graph.meetingRegistration entity.
type ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilderGetQueryParameters custom registration questions.
type ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilderGetQueryParameters
}
// ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilderInternal instantiates a new ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilder and sets the default values.
func NewItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilder) {
    m := &ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/onlineMeetings/{onlineMeeting%2Did}/registration/customQuestions/{meetingRegistrationQuestion%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilder instantiates a new ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilder and sets the default values.
func NewItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property customQuestions for users
// Deprecated: The meetingRegistrationBase Entity is deprecated and will stop returning data on Dec 12th, 2024. Please use the new webinar APIs. as of 2024-04/meetingRegistrationDeprecation on 2024-04-01 and will be removed 2024-12-12
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get custom registration questions.
// Deprecated: The meetingRegistrationBase Entity is deprecated and will stop returning data on Dec 12th, 2024. Please use the new webinar APIs. as of 2024-04/meetingRegistrationDeprecation on 2024-04-01 and will be removed 2024-12-12
// returns a MeetingRegistrationQuestionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingRegistrationQuestionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// Patch update the navigation property customQuestions in users
// Deprecated: The meetingRegistrationBase Entity is deprecated and will stop returning data on Dec 12th, 2024. Please use the new webinar APIs. as of 2024-04/meetingRegistrationDeprecation on 2024-04-01 and will be removed 2024-12-12
// returns a MeetingRegistrationQuestionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingRegistrationQuestionable, requestConfiguration *ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingRegistrationQuestionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToDeleteRequestInformation delete navigation property customQuestions for users
// Deprecated: The meetingRegistrationBase Entity is deprecated and will stop returning data on Dec 12th, 2024. Please use the new webinar APIs. as of 2024-04/meetingRegistrationDeprecation on 2024-04-01 and will be removed 2024-12-12
// returns a *RequestInformation when successful
func (m *ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation custom registration questions.
// Deprecated: The meetingRegistrationBase Entity is deprecated and will stop returning data on Dec 12th, 2024. Please use the new webinar APIs. as of 2024-04/meetingRegistrationDeprecation on 2024-04-01 and will be removed 2024-12-12
// returns a *RequestInformation when successful
func (m *ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property customQuestions in users
// Deprecated: The meetingRegistrationBase Entity is deprecated and will stop returning data on Dec 12th, 2024. Please use the new webinar APIs. as of 2024-04/meetingRegistrationDeprecation on 2024-04-01 and will be removed 2024-12-12
// returns a *RequestInformation when successful
func (m *ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingRegistrationQuestionable, requestConfiguration *ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The meetingRegistrationBase Entity is deprecated and will stop returning data on Dec 12th, 2024. Please use the new webinar APIs. as of 2024-04/meetingRegistrationDeprecation on 2024-04-01 and will be removed 2024-12-12
// returns a *ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilder when successful
func (m *ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilder) WithUrl(rawUrl string)(*ItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilder) {
    return NewItemOnlineMeetingsItemRegistrationCustomQuestionsMeetingRegistrationQuestionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
