package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder provides operations to manage the webinars property of the microsoft.graph.virtualEventsRoot entity.
type VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderGetQueryParameters read the properties and relationships of a virtualEventWebinar object.
type VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderGetQueryParameters
}
// VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderInternal instantiates a new VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder and sets the default values.
func NewVirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) {
    m := &VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/webinars/{virtualEventWebinar%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder instantiates a new VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder and sets the default values.
func NewVirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property webinars for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a virtualEventWebinar object.
// returns a VirtualEventWebinarable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualeventwebinar-get?view=graph-rest-beta
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventWebinarable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVirtualEventWebinarFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventWebinarable), nil
}
// Patch update the properties of a virtualEventWebinar object.
// returns a VirtualEventWebinarable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualeventwebinar-update?view=graph-rest-beta
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventWebinarable, requestConfiguration *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventWebinarable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVirtualEventWebinarFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventWebinarable), nil
}
// Presenters provides operations to manage the presenters property of the microsoft.graph.virtualEvent entity.
// returns a *VirtualeventsWebinarsItemPresentersRequestBuilder when successful
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) Presenters()(*VirtualeventsWebinarsItemPresentersRequestBuilder) {
    return NewVirtualeventsWebinarsItemPresentersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RegistrationConfiguration provides operations to manage the registrationConfiguration property of the microsoft.graph.virtualEventWebinar entity.
// returns a *VirtualeventsWebinarsItemRegistrationconfigurationRegistrationConfigurationRequestBuilder when successful
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) RegistrationConfiguration()(*VirtualeventsWebinarsItemRegistrationconfigurationRegistrationConfigurationRequestBuilder) {
    return NewVirtualeventsWebinarsItemRegistrationconfigurationRegistrationConfigurationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Registrations provides operations to manage the registrations property of the microsoft.graph.virtualEventWebinar entity.
// returns a *VirtualeventsWebinarsItemRegistrationsRequestBuilder when successful
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) Registrations()(*VirtualeventsWebinarsItemRegistrationsRequestBuilder) {
    return NewVirtualeventsWebinarsItemRegistrationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RegistrationsWithEmail provides operations to manage the registrations property of the microsoft.graph.virtualEventWebinar entity.
// returns a *VirtualeventsWebinarsItemRegistrationswithemailRegistrationsWithEmailRequestBuilder when successful
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) RegistrationsWithEmail(email *string)(*VirtualeventsWebinarsItemRegistrationswithemailRegistrationsWithEmailRequestBuilder) {
    return NewVirtualeventsWebinarsItemRegistrationswithemailRegistrationsWithEmailRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, email)
}
// RegistrationsWithUserId provides operations to manage the registrations property of the microsoft.graph.virtualEventWebinar entity.
// returns a *VirtualeventsWebinarsItemRegistrationswithuseridRegistrationsWithUserIdRequestBuilder when successful
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) RegistrationsWithUserId(userId *string)(*VirtualeventsWebinarsItemRegistrationswithuseridRegistrationsWithUserIdRequestBuilder) {
    return NewVirtualeventsWebinarsItemRegistrationswithuseridRegistrationsWithUserIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, userId)
}
// Sessions provides operations to manage the sessions property of the microsoft.graph.virtualEvent entity.
// returns a *VirtualeventsWebinarsItemSessionsRequestBuilder when successful
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) Sessions()(*VirtualeventsWebinarsItemSessionsRequestBuilder) {
    return NewVirtualeventsWebinarsItemSessionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SessionsWithJoinWebUrl provides operations to manage the sessions property of the microsoft.graph.virtualEvent entity.
// returns a *VirtualeventsWebinarsItemSessionswithjoinweburlSessionsWithJoinWebUrlRequestBuilder when successful
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) SessionsWithJoinWebUrl(joinWebUrl *string)(*VirtualeventsWebinarsItemSessionswithjoinweburlSessionsWithJoinWebUrlRequestBuilder) {
    return NewVirtualeventsWebinarsItemSessionswithjoinweburlSessionsWithJoinWebUrlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, joinWebUrl)
}
// ToDeleteRequestInformation delete navigation property webinars for solutions
// returns a *RequestInformation when successful
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a virtualEventWebinar object.
// returns a *RequestInformation when successful
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a virtualEventWebinar object.
// returns a *RequestInformation when successful
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventWebinarable, requestConfiguration *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder when successful
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) WithUrl(rawUrl string)(*VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) {
    return NewVirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
