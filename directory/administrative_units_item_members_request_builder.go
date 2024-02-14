package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AdministrativeUnitsItemMembersRequestBuilder provides operations to manage the members property of the microsoft.graph.administrativeUnit entity.
type AdministrativeUnitsItemMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AdministrativeUnitsItemMembersRequestBuilderGetQueryParameters users and groups that are members of this administrative unit. Supports $expand.
type AdministrativeUnitsItemMembersRequestBuilderGetQueryParameters struct {
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
// AdministrativeUnitsItemMembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdministrativeUnitsItemMembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AdministrativeUnitsItemMembersRequestBuilderGetQueryParameters
}
// AdministrativeUnitsItemMembersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdministrativeUnitsItemMembersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDirectoryObjectId gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.administrativeUnits.item.members.item collection
// returns a *AdministrativeUnitsItemMembersDirectoryObjectItemRequestBuilder when successful
func (m *AdministrativeUnitsItemMembersRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*AdministrativeUnitsItemMembersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewAdministrativeUnitsItemMembersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAdministrativeUnitsItemMembersRequestBuilderInternal instantiates a new AdministrativeUnitsItemMembersRequestBuilder and sets the default values.
func NewAdministrativeUnitsItemMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdministrativeUnitsItemMembersRequestBuilder) {
    m := &AdministrativeUnitsItemMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/administrativeUnits/{administrativeUnit%2Did}/members{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAdministrativeUnitsItemMembersRequestBuilder instantiates a new AdministrativeUnitsItemMembersRequestBuilder and sets the default values.
func NewAdministrativeUnitsItemMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdministrativeUnitsItemMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdministrativeUnitsItemMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AdministrativeUnitsItemMembersCountRequestBuilder when successful
func (m *AdministrativeUnitsItemMembersRequestBuilder) Count()(*AdministrativeUnitsItemMembersCountRequestBuilder) {
    return NewAdministrativeUnitsItemMembersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get users and groups that are members of this administrative unit. Supports $expand.
// returns a DirectoryObjectCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AdministrativeUnitsItemMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *AdministrativeUnitsItemMembersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable), nil
}
// GraphApplication casts the previous resource to application.
// returns a *AdministrativeUnitsItemMembersGraphApplicationRequestBuilder when successful
func (m *AdministrativeUnitsItemMembersRequestBuilder) GraphApplication()(*AdministrativeUnitsItemMembersGraphApplicationRequestBuilder) {
    return NewAdministrativeUnitsItemMembersGraphApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphDevice casts the previous resource to device.
// returns a *AdministrativeUnitsItemMembersGraphDeviceRequestBuilder when successful
func (m *AdministrativeUnitsItemMembersRequestBuilder) GraphDevice()(*AdministrativeUnitsItemMembersGraphDeviceRequestBuilder) {
    return NewAdministrativeUnitsItemMembersGraphDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphGroup casts the previous resource to group.
// returns a *AdministrativeUnitsItemMembersGraphGroupRequestBuilder when successful
func (m *AdministrativeUnitsItemMembersRequestBuilder) GraphGroup()(*AdministrativeUnitsItemMembersGraphGroupRequestBuilder) {
    return NewAdministrativeUnitsItemMembersGraphGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphOrgContact casts the previous resource to orgContact.
// returns a *AdministrativeUnitsItemMembersGraphOrgContactRequestBuilder when successful
func (m *AdministrativeUnitsItemMembersRequestBuilder) GraphOrgContact()(*AdministrativeUnitsItemMembersGraphOrgContactRequestBuilder) {
    return NewAdministrativeUnitsItemMembersGraphOrgContactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
// returns a *AdministrativeUnitsItemMembersGraphServicePrincipalRequestBuilder when successful
func (m *AdministrativeUnitsItemMembersRequestBuilder) GraphServicePrincipal()(*AdministrativeUnitsItemMembersGraphServicePrincipalRequestBuilder) {
    return NewAdministrativeUnitsItemMembersGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
// returns a *AdministrativeUnitsItemMembersGraphUserRequestBuilder when successful
func (m *AdministrativeUnitsItemMembersRequestBuilder) GraphUser()(*AdministrativeUnitsItemMembersGraphUserRequestBuilder) {
    return NewAdministrativeUnitsItemMembersGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to members for directory
// returns a DirectoryObjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/administrativeunit-post-members?view=graph-rest-1.0
func (m *AdministrativeUnitsItemMembersRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, requestConfiguration *AdministrativeUnitsItemMembersRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable), nil
}
// Ref provides operations to manage the collection of directory entities.
// returns a *AdministrativeUnitsItemMembersRefRequestBuilder when successful
func (m *AdministrativeUnitsItemMembersRequestBuilder) Ref()(*AdministrativeUnitsItemMembersRefRequestBuilder) {
    return NewAdministrativeUnitsItemMembersRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation users and groups that are members of this administrative unit. Supports $expand.
// returns a *RequestInformation when successful
func (m *AdministrativeUnitsItemMembersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AdministrativeUnitsItemMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to members for directory
// returns a *RequestInformation when successful
func (m *AdministrativeUnitsItemMembersRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, requestConfiguration *AdministrativeUnitsItemMembersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/directory/administrativeUnits/{administrativeUnit%2Did}/members", m.BaseRequestBuilder.PathParameters)
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
// returns a *AdministrativeUnitsItemMembersRequestBuilder when successful
func (m *AdministrativeUnitsItemMembersRequestBuilder) WithUrl(rawUrl string)(*AdministrativeUnitsItemMembersRequestBuilder) {
    return NewAdministrativeUnitsItemMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
