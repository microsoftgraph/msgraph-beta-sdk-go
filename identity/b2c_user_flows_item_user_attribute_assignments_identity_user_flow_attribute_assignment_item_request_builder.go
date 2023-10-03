package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder provides operations to manage the userAttributeAssignments property of the microsoft.graph.b2cIdentityUserFlow entity.
type B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderGetQueryParameters read the properties and relationships of an identityUserFlowAttributeAssignment object. This API is supported in the following national cloud deployments.
type B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderGetQueryParameters
}
// B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewB2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderInternal instantiates a new IdentityUserFlowAttributeAssignmentItemRequestBuilder and sets the default values.
func NewB2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) {
    m := &B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow%2Did}/userAttributeAssignments/{identityUserFlowAttributeAssignment%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewB2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder instantiates a new IdentityUserFlowAttributeAssignmentItemRequestBuilder and sets the default values.
func NewB2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete an identityUserFlowAttributeAssignment object. This API is supported in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityuserflowattributeassignment-delete?view=graph-rest-1.0
func (m *B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an identityUserFlowAttributeAssignment object. This API is supported in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityuserflowattributeassignment-get?view=graph-rest-1.0
func (m *B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityUserFlowAttributeAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentityUserFlowAttributeAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityUserFlowAttributeAssignmentable), nil
}
// Patch update the properties of a identityUserFlowAttributeAssignment object. This API is supported in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityuserflowattributeassignment-update?view=graph-rest-1.0
func (m *B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityUserFlowAttributeAssignmentable, requestConfiguration *B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityUserFlowAttributeAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentityUserFlowAttributeAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityUserFlowAttributeAssignmentable), nil
}
// ToDeleteRequestInformation delete an identityUserFlowAttributeAssignment object. This API is supported in the following national cloud deployments.
func (m *B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation read the properties and relationships of an identityUserFlowAttributeAssignment object. This API is supported in the following national cloud deployments.
func (m *B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a identityUserFlowAttributeAssignment object. This API is supported in the following national cloud deployments.
func (m *B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityUserFlowAttributeAssignmentable, requestConfiguration *B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UserAttribute provides operations to manage the userAttribute property of the microsoft.graph.identityUserFlowAttributeAssignment entity.
func (m *B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) UserAttribute()(*B2cUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilder) {
    return NewB2cUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*B2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) {
    return NewB2cUserFlowsItemUserAttributeAssignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
