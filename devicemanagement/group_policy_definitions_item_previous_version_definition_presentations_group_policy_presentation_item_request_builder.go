package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder provides operations to manage the presentations property of the microsoft.graph.groupPolicyDefinition entity.
type GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetQueryParameters the group policy presentations associated with the definition.
type GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetQueryParameters
}
// GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderInternal instantiates a new GroupPolicyPresentationItemRequestBuilder and sets the default values.
func NewGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) {
    m := &GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyDefinitions/{groupPolicyDefinition%2Did}/previousVersionDefinition/presentations/{groupPolicyPresentation%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder instantiates a new GroupPolicyPresentationItemRequestBuilder and sets the default values.
func NewGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Definition provides operations to manage the definition property of the microsoft.graph.groupPolicyPresentation entity.
func (m *GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) Definition()(*GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilder) {
    return NewGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsItemDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property presentations for deviceManagement
func (m *GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the group policy presentations associated with the definition.
func (m *GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyPresentationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable), nil
}
// Patch update the navigation property presentations in deviceManagement
func (m *GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable, requestConfiguration *GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyPresentationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable), nil
}
// ToDeleteRequestInformation delete navigation property presentations for deviceManagement
func (m *GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    requestInfo.Headers.TryAdd("Accept", "application/json, application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the group policy presentations associated with the definition.
func (m *GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property presentations in deviceManagement
func (m *GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationable, requestConfiguration *GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) WithUrl(rawUrl string)(*GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder) {
    return NewGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsGroupPolicyPresentationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
