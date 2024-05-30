package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RolescopetagsRoleScopeTagsRequestBuilder provides operations to manage the roleScopeTags property of the microsoft.graph.deviceManagement entity.
type RolescopetagsRoleScopeTagsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RolescopetagsRoleScopeTagsRequestBuilderGetQueryParameters the Role Scope Tags.
type RolescopetagsRoleScopeTagsRequestBuilderGetQueryParameters struct {
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
// RolescopetagsRoleScopeTagsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolescopetagsRoleScopeTagsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RolescopetagsRoleScopeTagsRequestBuilderGetQueryParameters
}
// RolescopetagsRoleScopeTagsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RolescopetagsRoleScopeTagsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByRoleScopeTagId provides operations to manage the roleScopeTags property of the microsoft.graph.deviceManagement entity.
// returns a *RolescopetagsRoleScopeTagItemRequestBuilder when successful
func (m *RolescopetagsRoleScopeTagsRequestBuilder) ByRoleScopeTagId(roleScopeTagId string)(*RolescopetagsRoleScopeTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if roleScopeTagId != "" {
        urlTplParams["roleScopeTag%2Did"] = roleScopeTagId
    }
    return NewRolescopetagsRoleScopeTagItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewRolescopetagsRoleScopeTagsRequestBuilderInternal instantiates a new RolescopetagsRoleScopeTagsRequestBuilder and sets the default values.
func NewRolescopetagsRoleScopeTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolescopetagsRoleScopeTagsRequestBuilder) {
    m := &RolescopetagsRoleScopeTagsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/roleScopeTags{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewRolescopetagsRoleScopeTagsRequestBuilder instantiates a new RolescopetagsRoleScopeTagsRequestBuilder and sets the default values.
func NewRolescopetagsRoleScopeTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RolescopetagsRoleScopeTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRolescopetagsRoleScopeTagsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *RolescopetagsCountRequestBuilder when successful
func (m *RolescopetagsRoleScopeTagsRequestBuilder) Count()(*RolescopetagsCountRequestBuilder) {
    return NewRolescopetagsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the Role Scope Tags.
// returns a RoleScopeTagCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolescopetagsRoleScopeTagsRequestBuilder) Get(ctx context.Context, requestConfiguration *RolescopetagsRoleScopeTagsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleScopeTagCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRoleScopeTagCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleScopeTagCollectionResponseable), nil
}
// GetRoleScopeTagsById provides operations to call the getRoleScopeTagsById method.
// returns a *RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilder when successful
func (m *RolescopetagsRoleScopeTagsRequestBuilder) GetRoleScopeTagsById()(*RolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilder) {
    return NewRolescopetagsGetrolescopetagsbyidGetRoleScopeTagsByIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HasCustomRoleScopeTag provides operations to call the hasCustomRoleScopeTag method.
// returns a *RolescopetagsHascustomrolescopetagHasCustomRoleScopeTagRequestBuilder when successful
func (m *RolescopetagsRoleScopeTagsRequestBuilder) HasCustomRoleScopeTag()(*RolescopetagsHascustomrolescopetagHasCustomRoleScopeTagRequestBuilder) {
    return NewRolescopetagsHascustomrolescopetagHasCustomRoleScopeTagRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to roleScopeTags for deviceManagement
// returns a RoleScopeTagable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RolescopetagsRoleScopeTagsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleScopeTagable, requestConfiguration *RolescopetagsRoleScopeTagsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleScopeTagable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRoleScopeTagFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleScopeTagable), nil
}
// ToGetRequestInformation the Role Scope Tags.
// returns a *RequestInformation when successful
func (m *RolescopetagsRoleScopeTagsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RolescopetagsRoleScopeTagsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to roleScopeTags for deviceManagement
// returns a *RequestInformation when successful
func (m *RolescopetagsRoleScopeTagsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleScopeTagable, requestConfiguration *RolescopetagsRoleScopeTagsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RolescopetagsRoleScopeTagsRequestBuilder when successful
func (m *RolescopetagsRoleScopeTagsRequestBuilder) WithUrl(rawUrl string)(*RolescopetagsRoleScopeTagsRequestBuilder) {
    return NewRolescopetagsRoleScopeTagsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
