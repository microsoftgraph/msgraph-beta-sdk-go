package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder provides operations to manage the userAttributeAssignments property of the microsoft.graph.b2cIdentityUserFlow entity.
type B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderGetQueryParameters read the properties and relationships of an identityUserFlowAttributeAssignment object.
type B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderGetQueryParameters
}
// B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewB2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderInternal instantiates a new B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder and sets the default values.
func NewB2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) {
    m := &B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow%2Did}/userAttributeAssignments/{identityUserFlowAttributeAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewB2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder instantiates a new B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder and sets the default values.
func NewB2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete an identityUserFlowAttributeAssignment object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityuserflowattributeassignment-delete?view=graph-rest-beta
func (m *B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an identityUserFlowAttributeAssignment object.
// returns a IdentityUserFlowAttributeAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityuserflowattributeassignment-get?view=graph-rest-beta
func (m *B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityUserFlowAttributeAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// Patch update the properties of a identityUserFlowAttributeAssignment object.
// returns a IdentityUserFlowAttributeAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityuserflowattributeassignment-update?view=graph-rest-beta
func (m *B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityUserFlowAttributeAssignmentable, requestConfiguration *B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityUserFlowAttributeAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToDeleteRequestInformation delete an identityUserFlowAttributeAssignment object.
// returns a *RequestInformation when successful
func (m *B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an identityUserFlowAttributeAssignment object.
// returns a *RequestInformation when successful
func (m *B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a identityUserFlowAttributeAssignment object.
// returns a *RequestInformation when successful
func (m *B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityUserFlowAttributeAssignmentable, requestConfiguration *B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UserAttribute provides operations to manage the userAttribute property of the microsoft.graph.identityUserFlowAttributeAssignment entity.
// returns a *B2cuserflowsItemUserattributeassignmentsItemUserattributeUserAttributeRequestBuilder when successful
func (m *B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) UserAttribute()(*B2cuserflowsItemUserattributeassignmentsItemUserattributeUserAttributeRequestBuilder) {
    return NewB2cuserflowsItemUserattributeassignmentsItemUserattributeUserAttributeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder when successful
func (m *B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*B2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) {
    return NewB2cuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
