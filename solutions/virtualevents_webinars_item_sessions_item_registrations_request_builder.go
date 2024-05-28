package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilder provides operations to manage the registrations property of the microsoft.graph.virtualEventSession entity.
type VirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilderGetQueryParameters get registrations from solutions
type VirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilderGetQueryParameters struct {
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
// VirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilderGetQueryParameters
}
// ByVirtualEventRegistrationId provides operations to manage the registrations property of the microsoft.graph.virtualEventSession entity.
// returns a *VirtualeventsWebinarsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilder when successful
func (m *VirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilder) ByVirtualEventRegistrationId(virtualEventRegistrationId string)(*VirtualeventsWebinarsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if virtualEventRegistrationId != "" {
        urlTplParams["virtualEventRegistration%2Did"] = virtualEventRegistrationId
    }
    return NewVirtualeventsWebinarsItemSessionsItemRegistrationsVirtualEventRegistrationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilderInternal instantiates a new VirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilder and sets the default values.
func NewVirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilder) {
    m := &VirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/webinars/{virtualEventWebinar%2Did}/sessions/{virtualEventSession%2Did}/registrations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilder instantiates a new VirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilder and sets the default values.
func NewVirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *VirtualeventsWebinarsItemSessionsItemRegistrationsCountRequestBuilder when successful
func (m *VirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilder) Count()(*VirtualeventsWebinarsItemSessionsItemRegistrationsCountRequestBuilder) {
    return NewVirtualeventsWebinarsItemSessionsItemRegistrationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get registrations from solutions
// returns a VirtualEventRegistrationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventRegistrationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVirtualEventRegistrationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEventRegistrationCollectionResponseable), nil
}
// ToGetRequestInformation get registrations from solutions
// returns a *RequestInformation when successful
func (m *VirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilder when successful
func (m *VirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilder) WithUrl(rawUrl string)(*VirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilder) {
    return NewVirtualeventsWebinarsItemSessionsItemRegistrationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
