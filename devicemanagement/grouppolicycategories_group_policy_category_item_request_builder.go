package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder provides operations to manage the groupPolicyCategories property of the microsoft.graph.deviceManagement entity.
type GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilderGetQueryParameters the available group policy categories for this account.
type GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilderGetQueryParameters
}
// GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Children provides operations to manage the children property of the microsoft.graph.groupPolicyCategory entity.
// returns a *GrouppolicycategoriesItemChildrenRequestBuilder when successful
func (m *GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder) Children()(*GrouppolicycategoriesItemChildrenRequestBuilder) {
    return NewGrouppolicycategoriesItemChildrenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewGrouppolicycategoriesGroupPolicyCategoryItemRequestBuilderInternal instantiates a new GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder and sets the default values.
func NewGrouppolicycategoriesGroupPolicyCategoryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder) {
    m := &GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyCategories/{groupPolicyCategory%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewGrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder instantiates a new GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder and sets the default values.
func NewGrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGrouppolicycategoriesGroupPolicyCategoryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// DefinitionFile provides operations to manage the definitionFile property of the microsoft.graph.groupPolicyCategory entity.
// returns a *GrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilder when successful
func (m *GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder) DefinitionFile()(*GrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilder) {
    return NewGrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Definitions provides operations to manage the definitions property of the microsoft.graph.groupPolicyCategory entity.
// returns a *GrouppolicycategoriesItemDefinitionsRequestBuilder when successful
func (m *GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder) Definitions()(*GrouppolicycategoriesItemDefinitionsRequestBuilder) {
    return NewGrouppolicycategoriesItemDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property groupPolicyCategories for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the available group policy categories for this account.
// returns a GroupPolicyCategoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyCategoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable), nil
}
// Parent provides operations to manage the parent property of the microsoft.graph.groupPolicyCategory entity.
// returns a *GrouppolicycategoriesItemParentRequestBuilder when successful
func (m *GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder) Parent()(*GrouppolicycategoriesItemParentRequestBuilder) {
    return NewGrouppolicycategoriesItemParentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property groupPolicyCategories in deviceManagement
// returns a GroupPolicyCategoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable, requestConfiguration *GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyCategoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable), nil
}
// ToDeleteRequestInformation delete navigation property groupPolicyCategories for deviceManagement
// returns a *RequestInformation when successful
func (m *GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the available group policy categories for this account.
// returns a *RequestInformation when successful
func (m *GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property groupPolicyCategories in deviceManagement
// returns a *RequestInformation when successful
func (m *GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyCategoryable, requestConfiguration *GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder when successful
func (m *GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder) WithUrl(rawUrl string)(*GrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder) {
    return NewGrouppolicycategoriesGroupPolicyCategoryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
