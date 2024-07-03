package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder provides operations to manage the previousVersionDefinition property of the microsoft.graph.groupPolicyDefinition entity.
type GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilderGetQueryParameters definition of the previous version of this definition
type GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilderGetQueryParameters
}
// GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Category provides operations to manage the category property of the microsoft.graph.groupPolicyDefinition entity.
// returns a *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionCategoryRequestBuilder when successful
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder) Category()(*GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionCategoryRequestBuilder) {
    return NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionCategoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilderInternal instantiates a new GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder and sets the default values.
func NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder) {
    m := &GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyDefinitions/{groupPolicyDefinition%2Did}/nextVersionDefinition/previousVersionDefinition{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder instantiates a new GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder and sets the default values.
func NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// DefinitionFile provides operations to manage the definitionFile property of the microsoft.graph.groupPolicyDefinition entity.
// returns a *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionDefinitionFileRequestBuilder when successful
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder) DefinitionFile()(*GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionDefinitionFileRequestBuilder) {
    return NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionDefinitionFileRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property previousVersionDefinition for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, error) {
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
// Patch update the navigation property previousVersionDefinition in deviceManagement
// returns a GroupPolicyDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, requestConfiguration *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, error) {
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
// returns a *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsRequestBuilder when successful
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder) Presentations()(*GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsRequestBuilder) {
    return NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionPresentationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property previousVersionDefinition for deviceManagement
// returns a *RequestInformation when successful
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, requestConfiguration *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder when successful
func (m *GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder) WithUrl(rawUrl string)(*GroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder) {
    return NewGroupPolicyDefinitionsItemNextVersionDefinitionPreviousVersionDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
