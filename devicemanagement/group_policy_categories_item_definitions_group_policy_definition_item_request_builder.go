package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder provides operations to manage the definitions property of the microsoft.graph.groupPolicyCategory entity.
type GroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderGetQueryParameters the immediate GroupPolicyDefinition children of the category
type GroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderGetQueryParameters
}
// NewGroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderInternal instantiates a new GroupPolicyDefinitionItemRequestBuilder and sets the default values.
func NewGroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder) {
    m := &GroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyCategories/{groupPolicyCategory%2Did}/definitions/{groupPolicyDefinition%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewGroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder instantiates a new GroupPolicyDefinitionItemRequestBuilder and sets the default values.
func NewGroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the immediate GroupPolicyDefinition children of the category
func (m *GroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable), nil
}
// ToGetRequestInformation the immediate GroupPolicyDefinition children of the category
func (m *GroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *GroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder) WithUrl(rawUrl string)(*GroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder) {
    return NewGroupPolicyCategoriesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
