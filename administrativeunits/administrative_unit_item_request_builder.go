package administrativeunits

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AdministrativeUnitItemRequestBuilder provides operations to manage the collection of administrativeUnit entities.
type AdministrativeUnitItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AdministrativeUnitItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdministrativeUnitItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdministrativeUnitItemRequestBuilderGetQueryParameters retrieve the properties and relationships of an administrativeUnit object. Since the **administrativeUnit** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in an **administrativeUnit** instance.
type AdministrativeUnitItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AdministrativeUnitItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdministrativeUnitItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AdministrativeUnitItemRequestBuilderGetQueryParameters
}
// AdministrativeUnitItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdministrativeUnitItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAdministrativeUnitItemRequestBuilderInternal instantiates a new AdministrativeUnitItemRequestBuilder and sets the default values.
func NewAdministrativeUnitItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdministrativeUnitItemRequestBuilder) {
    m := &AdministrativeUnitItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/administrativeUnits/{administrativeUnit%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAdministrativeUnitItemRequestBuilder instantiates a new AdministrativeUnitItemRequestBuilder and sets the default values.
func NewAdministrativeUnitItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdministrativeUnitItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdministrativeUnitItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete an administrativeUnit.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/administrativeunit-delete?view=graph-rest-1.0
func (m *AdministrativeUnitItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AdministrativeUnitItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.administrativeUnit entity.
func (m *AdministrativeUnitItemRequestBuilder) Extensions()(*ItemExtensionsRequestBuilder) {
    return NewItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.administrativeUnit entity.
func (m *AdministrativeUnitItemRequestBuilder) ExtensionsById(id string)(*ItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get retrieve the properties and relationships of an administrativeUnit object. Since the **administrativeUnit** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in an **administrativeUnit** instance.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/administrativeunit-get?view=graph-rest-1.0
func (m *AdministrativeUnitItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AdministrativeUnitItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministrativeUnitable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAdministrativeUnitFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministrativeUnitable), nil
}
// Members provides operations to manage the members property of the microsoft.graph.administrativeUnit entity.
func (m *AdministrativeUnitItemRequestBuilder) Members()(*ItemMembersRequestBuilder) {
    return NewItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.administrativeUnits.item.members.item collection
func (m *AdministrativeUnitItemRequestBuilder) MembersById(id string)(*ItemMembersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemMembersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MicrosoftGraphCheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *AdministrativeUnitItemRequestBuilder) MicrosoftGraphCheckMemberGroups()(*ItemMicrosoftGraphCheckMemberGroupsRequestBuilder) {
    return NewItemMicrosoftGraphCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *AdministrativeUnitItemRequestBuilder) MicrosoftGraphCheckMemberObjects()(*ItemMicrosoftGraphCheckMemberObjectsRequestBuilder) {
    return NewItemMicrosoftGraphCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetMemberGroups provides operations to call the getMemberGroups method.
func (m *AdministrativeUnitItemRequestBuilder) MicrosoftGraphGetMemberGroups()(*ItemMicrosoftGraphGetMemberGroupsRequestBuilder) {
    return NewItemMicrosoftGraphGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetMemberObjects provides operations to call the getMemberObjects method.
func (m *AdministrativeUnitItemRequestBuilder) MicrosoftGraphGetMemberObjects()(*ItemMicrosoftGraphGetMemberObjectsRequestBuilder) {
    return NewItemMicrosoftGraphGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRestore provides operations to call the restore method.
func (m *AdministrativeUnitItemRequestBuilder) MicrosoftGraphRestore()(*ItemMicrosoftGraphRestoreRequestBuilder) {
    return NewItemMicrosoftGraphRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the properties of an administrativeUnit object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/administrativeunit-update?view=graph-rest-1.0
func (m *AdministrativeUnitItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministrativeUnitable, requestConfiguration *AdministrativeUnitItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministrativeUnitable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAdministrativeUnitFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministrativeUnitable), nil
}
// ScopedRoleMembers provides operations to manage the scopedRoleMembers property of the microsoft.graph.administrativeUnit entity.
func (m *AdministrativeUnitItemRequestBuilder) ScopedRoleMembers()(*ItemScopedRoleMembersRequestBuilder) {
    return NewItemScopedRoleMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ScopedRoleMembersById provides operations to manage the scopedRoleMembers property of the microsoft.graph.administrativeUnit entity.
func (m *AdministrativeUnitItemRequestBuilder) ScopedRoleMembersById(id string)(*ItemScopedRoleMembersScopedRoleMembershipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership%2Did"] = id
    }
    return NewItemScopedRoleMembersScopedRoleMembershipItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete an administrativeUnit.
func (m *AdministrativeUnitItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AdministrativeUnitItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties and relationships of an administrativeUnit object. Since the **administrativeUnit** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in an **administrativeUnit** instance.
func (m *AdministrativeUnitItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AdministrativeUnitItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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
// ToPatchRequestInformation update the properties of an administrativeUnit object.
func (m *AdministrativeUnitItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministrativeUnitable, requestConfiguration *AdministrativeUnitItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
