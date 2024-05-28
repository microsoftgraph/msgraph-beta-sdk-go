package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder provides operations to manage the transitiveRoleAssignments property of the microsoft.graph.rbacApplication entity.
type EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderGetQueryParameters get transitiveRoleAssignments from roleManagement
type EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderGetQueryParameters struct {
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
// EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderGetQueryParameters
}
// EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUnifiedRoleAssignmentId provides operations to manage the transitiveRoleAssignments property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementTransitiveroleassignmentsUnifiedRoleAssignmentItemRequestBuilder when successful
func (m *EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder) ByUnifiedRoleAssignmentId(unifiedRoleAssignmentId string)(*EntitlementmanagementTransitiveroleassignmentsUnifiedRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unifiedRoleAssignmentId != "" {
        urlTplParams["unifiedRoleAssignment%2Did"] = unifiedRoleAssignmentId
    }
    return NewEntitlementmanagementTransitiveroleassignmentsUnifiedRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderInternal instantiates a new EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder and sets the default values.
func NewEntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder) {
    m := &EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/entitlementManagement/transitiveRoleAssignments{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder instantiates a new EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder and sets the default values.
func NewEntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EntitlementmanagementTransitiveroleassignmentsCountRequestBuilder when successful
func (m *EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder) Count()(*EntitlementmanagementTransitiveroleassignmentsCountRequestBuilder) {
    return NewEntitlementmanagementTransitiveroleassignmentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get transitiveRoleAssignments from roleManagement
// returns a UnifiedRoleAssignmentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleAssignmentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentCollectionResponseable), nil
}
// Post create new navigation property to transitiveRoleAssignments for roleManagement
// returns a UnifiedRoleAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentable, requestConfiguration *EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentable), nil
}
// ToGetRequestInformation get transitiveRoleAssignments from roleManagement
// returns a *RequestInformation when successful
func (m *EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to transitiveRoleAssignments for roleManagement
// returns a *RequestInformation when successful
func (m *EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentable, requestConfiguration *EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder when successful
func (m *EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder) {
    return NewEntitlementmanagementTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
