package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder provides operations to manage the eligibilityScheduleRequests property of the microsoft.graph.privilegedAccessGroup entity.
type PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderGetQueryParameters read the properties and relationships of a privilegedAccessGroupEligibilityScheduleRequest object.
type PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderGetQueryParameters
}
// PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Cancel provides operations to call the cancel method.
func (m *PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) Cancel()(*PrivilegedAccessGroupEligibilityScheduleRequestsItemCancelRequestBuilder) {
    return NewPrivilegedAccessGroupEligibilityScheduleRequestsItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderInternal instantiates a new PrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) {
    m := &PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests/{privilegedAccessGroupEligibilityScheduleRequest%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewPrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder instantiates a new PrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property eligibilityScheduleRequests for identityGovernance
func (m *PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read the properties and relationships of a privilegedAccessGroupEligibilityScheduleRequest object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/privilegedaccessgroupeligibilityschedulerequest-get?view=graph-rest-1.0
func (m *PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupEligibilityScheduleRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Group provides operations to manage the group property of the microsoft.graph.privilegedAccessGroupEligibilityScheduleRequest entity.
func (m *PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) Group()(*PrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilder) {
    return NewPrivilegedAccessGroupEligibilityScheduleRequestsItemGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property eligibilityScheduleRequests in identityGovernance
func (m *PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupEligibilityScheduleRequestable, requestConfiguration *PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupEligibilityScheduleRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// Principal provides operations to manage the principal property of the microsoft.graph.privilegedAccessGroupEligibilityScheduleRequest entity.
func (m *PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) Principal()(*PrivilegedAccessGroupEligibilityScheduleRequestsItemPrincipalRequestBuilder) {
    return NewPrivilegedAccessGroupEligibilityScheduleRequestsItemPrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TargetSchedule provides operations to manage the targetSchedule property of the microsoft.graph.privilegedAccessGroupEligibilityScheduleRequest entity.
func (m *PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) TargetSchedule()(*PrivilegedAccessGroupEligibilityScheduleRequestsItemTargetScheduleRequestBuilder) {
    return NewPrivilegedAccessGroupEligibilityScheduleRequestsItemTargetScheduleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property eligibilityScheduleRequests for identityGovernance
func (m *PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a privilegedAccessGroupEligibilityScheduleRequest object.
func (m *PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property eligibilityScheduleRequests in identityGovernance
func (m *PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupEligibilityScheduleRequestable, requestConfiguration *PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) WithUrl(rawUrl string)(*PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) {
    return NewPrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
