package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilder provides operations to manage the eligibilityScheduleRequests property of the microsoft.graph.privilegedAccessGroup entity.
type PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilderGetQueryParameters get a list of the privilegedAccessGroupEligibilityScheduleRequest objects and their properties.
type PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilderGetQueryParameters struct {
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
// PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilderGetQueryParameters
}
// PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPrivilegedAccessGroupEligibilityScheduleRequestId provides operations to manage the eligibilityScheduleRequests property of the microsoft.graph.privilegedAccessGroup entity.
func (m *PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilder) ByPrivilegedAccessGroupEligibilityScheduleRequestId(privilegedAccessGroupEligibilityScheduleRequestId string)(*PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if privilegedAccessGroupEligibilityScheduleRequestId != "" {
        urlTplParams["privilegedAccessGroupEligibilityScheduleRequest%2Did"] = privilegedAccessGroupEligibilityScheduleRequestId
    }
    return NewPrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilderInternal instantiates a new EligibilityScheduleRequestsRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilder) {
    m := &PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewPrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilder instantiates a new EligibilityScheduleRequestsRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilder) Count()(*PrivilegedAccessGroupEligibilityScheduleRequestsCountRequestBuilder) {
    return NewPrivilegedAccessGroupEligibilityScheduleRequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
func (m *PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilder) FilterByCurrentUserWithOn(on *string)(*PrivilegedAccessGroupEligibilityScheduleRequestsFilterByCurrentUserWithOnRequestBuilder) {
    return NewPrivilegedAccessGroupEligibilityScheduleRequestsFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
// Get get a list of the privilegedAccessGroupEligibilityScheduleRequest objects and their properties.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/privilegedaccessgroup-list-eligibilityschedulerequests?view=graph-rest-1.0
func (m *PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupEligibilityScheduleRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegedAccessGroupEligibilityScheduleRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupEligibilityScheduleRequestCollectionResponseable), nil
}
// Post create a new privilegedAccessGroupEligibilityScheduleRequest object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/privilegedaccessgroup-post-eligibilityschedulerequests?view=graph-rest-1.0
func (m *PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupEligibilityScheduleRequestable, requestConfiguration *PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupEligibilityScheduleRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegedAccessGroupEligibilityScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupEligibilityScheduleRequestable), nil
}
// ToGetRequestInformation get a list of the privilegedAccessGroupEligibilityScheduleRequest objects and their properties.
func (m *PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new privilegedAccessGroupEligibilityScheduleRequest object.
func (m *PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupEligibilityScheduleRequestable, requestConfiguration *PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
