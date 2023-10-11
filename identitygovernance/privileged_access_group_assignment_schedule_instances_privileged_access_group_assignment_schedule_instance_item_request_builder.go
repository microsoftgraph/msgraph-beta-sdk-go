package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder provides operations to manage the assignmentScheduleInstances property of the microsoft.graph.privilegedAccessGroup entity.
type PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderGetQueryParameters read the properties and relationships of a privilegedAccessGroupAssignmentScheduleInstance object. This API is available in the following national cloud deployments.
type PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderGetQueryParameters
}
// PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivatedUsing provides operations to manage the activatedUsing property of the microsoft.graph.privilegedAccessGroupAssignmentScheduleInstance entity.
func (m *PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) ActivatedUsing()(*PrivilegedAccessGroupAssignmentScheduleInstancesItemActivatedUsingRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentScheduleInstancesItemActivatedUsingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderInternal instantiates a new PrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) {
    m := &PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/assignmentScheduleInstances/{privilegedAccessGroupAssignmentScheduleInstance%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewPrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder instantiates a new PrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignmentScheduleInstances for identityGovernance
func (m *PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a privilegedAccessGroupAssignmentScheduleInstance object. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/privilegedaccessgroupassignmentscheduleinstance-get?view=graph-rest-1.0
func (m *PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupAssignmentScheduleInstanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegedAccessGroupAssignmentScheduleInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupAssignmentScheduleInstanceable), nil
}
// Group provides operations to manage the group property of the microsoft.graph.privilegedAccessGroupAssignmentScheduleInstance entity.
func (m *PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) Group()(*PrivilegedAccessGroupAssignmentScheduleInstancesItemGroupRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentScheduleInstancesItemGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property assignmentScheduleInstances in identityGovernance
func (m *PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupAssignmentScheduleInstanceable, requestConfiguration *PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupAssignmentScheduleInstanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegedAccessGroupAssignmentScheduleInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupAssignmentScheduleInstanceable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.privilegedAccessGroupAssignmentScheduleInstance entity.
func (m *PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) Principal()(*PrivilegedAccessGroupAssignmentScheduleInstancesItemPrincipalRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentScheduleInstancesItemPrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property assignmentScheduleInstances for identityGovernance
func (m *PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a privilegedAccessGroupAssignmentScheduleInstance object. This API is available in the following national cloud deployments.
func (m *PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignmentScheduleInstances in identityGovernance
func (m *PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupAssignmentScheduleInstanceable, requestConfiguration *PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) WithUrl(rawUrl string)(*PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
