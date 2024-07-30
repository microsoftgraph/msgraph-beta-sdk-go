package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilder provides operations to manage the questions property of the microsoft.graph.virtualEventRegistrationConfiguration entity.
type VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilderGetQueryParameters get a list of all registration questions for a webinar. The list can include either predefined registration questions or custom registration questions.
type VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilderGetQueryParameters struct {
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
// VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilderGetQueryParameters
}
// VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByVirtualEventRegistrationQuestionBaseId provides operations to manage the questions property of the microsoft.graph.virtualEventRegistrationConfiguration entity.
// returns a *VirtualEventsWebinarsItemRegistrationConfigurationQuestionsVirtualEventRegistrationQuestionBaseItemRequestBuilder when successful
func (m *VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilder) ByVirtualEventRegistrationQuestionBaseId(virtualEventRegistrationQuestionBaseId string)(*VirtualEventsWebinarsItemRegistrationConfigurationQuestionsVirtualEventRegistrationQuestionBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if virtualEventRegistrationQuestionBaseId != "" {
        urlTplParams["virtualEventRegistrationQuestionBase%2Did"] = virtualEventRegistrationQuestionBaseId
    }
    return NewVirtualEventsWebinarsItemRegistrationConfigurationQuestionsVirtualEventRegistrationQuestionBaseItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilderInternal instantiates a new VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilder) {
    m := &VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/webinars/{virtualEventWebinar%2Did}/registrationConfiguration/questions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilder instantiates a new VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilder and sets the default values.
func NewVirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *VirtualEventsWebinarsItemRegistrationConfigurationQuestionsCountRequestBuilder when successful
func (m *VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilder) Count()(*VirtualEventsWebinarsItemRegistrationConfigurationQuestionsCountRequestBuilder) {
    return NewVirtualEventsWebinarsItemRegistrationConfigurationQuestionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of all registration questions for a webinar. The list can include either predefined registration questions or custom registration questions.
// returns a VirtualEventRegistrationQuestionBaseCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualeventregistrationconfiguration-list-questions?view=graph-rest-beta
func (m *VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventRegistrationQuestionBaseCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVirtualEventRegistrationQuestionBaseCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventRegistrationQuestionBaseCollectionResponseable), nil
}
// Post create a registration question for a webinar. You can create either a predefined registration question or a custom registration question.
// returns a VirtualEventRegistrationQuestionBaseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualeventregistrationconfiguration-post-questions?view=graph-rest-beta
func (m *VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventRegistrationQuestionBaseable, requestConfiguration *VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventRegistrationQuestionBaseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVirtualEventRegistrationQuestionBaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventRegistrationQuestionBaseable), nil
}
// ToGetRequestInformation get a list of all registration questions for a webinar. The list can include either predefined registration questions or custom registration questions.
// returns a *RequestInformation when successful
func (m *VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a registration question for a webinar. You can create either a predefined registration question or a custom registration question.
// returns a *RequestInformation when successful
func (m *VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventRegistrationQuestionBaseable, requestConfiguration *VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilder when successful
func (m *VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilder) WithUrl(rawUrl string)(*VirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilder) {
    return NewVirtualEventsWebinarsItemRegistrationConfigurationQuestionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
