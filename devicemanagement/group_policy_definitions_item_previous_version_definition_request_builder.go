package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder provides operations to manage the previousVersionDefinition property of the microsoft.graph.groupPolicyDefinition entity.
type GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilderGetQueryParameters definition of the previous version of this definition
type GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilderGetQueryParameters
}
// GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Category provides operations to manage the category property of the microsoft.graph.groupPolicyDefinition entity.
// returns a *GroupPolicyDefinitionsItemPreviousVersionDefinitionCategoryRequestBuilder when successful
func (m *GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder) Category()(*GroupPolicyDefinitionsItemPreviousVersionDefinitionCategoryRequestBuilder) {
    return NewGroupPolicyDefinitionsItemPreviousVersionDefinitionCategoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewGroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilderInternal instantiates a new GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder and sets the default values.
func NewGroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder) {
    m := &GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyDefinitions/{groupPolicyDefinition%2Did}/previousVersionDefinition{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewGroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder instantiates a new GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder and sets the default values.
func NewGroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// DefinitionFile provides operations to manage the definitionFile property of the microsoft.graph.groupPolicyDefinition entity.
// returns a *GroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilder when successful
func (m *GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder) DefinitionFile()(*GroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilder) {
    return NewGroupPolicyDefinitionsItemPreviousVersionDefinitionDefinitionFileRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property previousVersionDefinition for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get definition of the previous version of this definition
// returns a GroupPolicyDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// NextVersionDefinition provides operations to manage the nextVersionDefinition property of the microsoft.graph.groupPolicyDefinition entity.
// returns a *GroupPolicyDefinitionsItemPreviousVersionDefinitionNextVersionDefinitionRequestBuilder when successful
func (m *GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder) NextVersionDefinition()(*GroupPolicyDefinitionsItemPreviousVersionDefinitionNextVersionDefinitionRequestBuilder) {
    return NewGroupPolicyDefinitionsItemPreviousVersionDefinitionNextVersionDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property previousVersionDefinition in deviceManagement
// returns a GroupPolicyDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, requestConfiguration *GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// Presentations provides operations to manage the presentations property of the microsoft.graph.groupPolicyDefinition entity.
// returns a *GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsRequestBuilder when successful
func (m *GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder) Presentations()(*GroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsRequestBuilder) {
    return NewGroupPolicyDefinitionsItemPreviousVersionDefinitionPresentationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property previousVersionDefinition for deviceManagement
// returns a *RequestInformation when successful
func (m *GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation definition of the previous version of this definition
// returns a *RequestInformation when successful
func (m *GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property previousVersionDefinition in deviceManagement
// returns a *RequestInformation when successful
func (m *GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, requestConfiguration *GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder when successful
func (m *GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder) WithUrl(rawUrl string)(*GroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder) {
    return NewGroupPolicyDefinitionsItemPreviousVersionDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
