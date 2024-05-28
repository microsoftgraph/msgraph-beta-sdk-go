package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilder provides operations to manage the scopedRoleMembers property of the microsoft.graph.administrativeUnit entity.
type AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilderGetQueryParameters scoped-role members of this administrative unit.
type AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilderGetQueryParameters struct {
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
// AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilderGetQueryParameters
}
// AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByScopedRoleMembershipId provides operations to manage the scopedRoleMembers property of the microsoft.graph.administrativeUnit entity.
// returns a *AdministrativeunitsItemScopedrolemembersScopedRoleMembershipItemRequestBuilder when successful
func (m *AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilder) ByScopedRoleMembershipId(scopedRoleMembershipId string)(*AdministrativeunitsItemScopedrolemembersScopedRoleMembershipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if scopedRoleMembershipId != "" {
        urlTplParams["scopedRoleMembership%2Did"] = scopedRoleMembershipId
    }
    return NewAdministrativeunitsItemScopedrolemembersScopedRoleMembershipItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilderInternal instantiates a new AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilder and sets the default values.
func NewAdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilder) {
    m := &AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/administrativeUnits/{administrativeUnit%2Did}/scopedRoleMembers{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilder instantiates a new AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilder and sets the default values.
func NewAdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AdministrativeunitsItemScopedrolemembersCountRequestBuilder when successful
func (m *AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilder) Count()(*AdministrativeunitsItemScopedrolemembersCountRequestBuilder) {
    return NewAdministrativeunitsItemScopedrolemembersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get scoped-role members of this administrative unit.
// returns a ScopedRoleMembershipCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScopedRoleMembershipCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateScopedRoleMembershipCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScopedRoleMembershipCollectionResponseable), nil
}
// Post create new navigation property to scopedRoleMembers for directory
// returns a ScopedRoleMembershipable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScopedRoleMembershipable, requestConfiguration *AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScopedRoleMembershipable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateScopedRoleMembershipFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScopedRoleMembershipable), nil
}
// ToGetRequestInformation scoped-role members of this administrative unit.
// returns a *RequestInformation when successful
func (m *AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to scopedRoleMembers for directory
// returns a *RequestInformation when successful
func (m *AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScopedRoleMembershipable, requestConfiguration *AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilder when successful
func (m *AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilder) WithUrl(rawUrl string)(*AdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilder) {
    return NewAdministrativeunitsItemScopedrolemembersScopedRoleMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
