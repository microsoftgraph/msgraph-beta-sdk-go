package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i162499125dee18a8c32afea0c335b92a83c8bfbf514f6c85bec33a36bdaf7ca3 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/getmemberobjects"
    i29c75035d87e1e3a958b91c43025f79430cd4c9d3b0cc664ffe065bc407fda28 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/scopedrolemembers"
    i8dcb8bc7850e64bb1fd90771daf6714da6fc512dd4adae05039d541e9bb9419c "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/restore"
    i98d18deca5882a7f1ef7bba571654837ca98c441bb0eaa212a79efbf51074283 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/checkmembergroups"
    i9a39d6011e7f6e52d0463e43f4affbe7fe566deec74cb2cfa89d0164876f73fb "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/members"
    ib3392aac80c6073c1b4cb66c1c9a59ef64a733b53d9fbe3bb2fdeea43b1e8ed0 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/extensions"
    ib744a95755b350bc30b575b7a619bbb35c32385d740a0feca8f31cddc0e3af5d "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/getmembergroups"
    icfc4b43cabbf84440085018bda409b58924c8286c1e14bce978493163ca182b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/checkmemberobjects"
    i348fc6817b120c2ca4453cfe220f3a31e3dd76b71234d51b3cf59448c038898e "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/extensions/item"
    ibb066ad55db4b7866d16a995539af863ab771176674fded378a94af09016fc1b "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/scopedrolemembers/item"
    id21ccd20a38c8c744ae37c53448fd59c79e0a14764716cf4113d7e8fd6f62db0 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/members/item"
)

// AdministrativeUnitItemRequestBuilder provides operations to manage the administrativeUnits property of the microsoft.graph.directory entity.
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
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdministrativeUnitItemRequestBuilderGetQueryParameters conceptual container for user and group directory objects.
type AdministrativeUnitItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AdministrativeUnitItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdministrativeUnitItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AdministrativeUnitItemRequestBuilderGetQueryParameters
}
// AdministrativeUnitItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdministrativeUnitItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CheckMemberGroups the checkMemberGroups property
func (m *AdministrativeUnitItemRequestBuilder) CheckMemberGroups()(*i98d18deca5882a7f1ef7bba571654837ca98c441bb0eaa212a79efbf51074283.CheckMemberGroupsRequestBuilder) {
    return i98d18deca5882a7f1ef7bba571654837ca98c441bb0eaa212a79efbf51074283.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *AdministrativeUnitItemRequestBuilder) CheckMemberObjects()(*icfc4b43cabbf84440085018bda409b58924c8286c1e14bce978493163ca182b6.CheckMemberObjectsRequestBuilder) {
    return icfc4b43cabbf84440085018bda409b58924c8286c1e14bce978493163ca182b6.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAdministrativeUnitItemRequestBuilderInternal instantiates a new AdministrativeUnitItemRequestBuilder and sets the default values.
func NewAdministrativeUnitItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdministrativeUnitItemRequestBuilder) {
    m := &AdministrativeUnitItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/administrativeUnits/{administrativeUnit%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAdministrativeUnitItemRequestBuilder instantiates a new AdministrativeUnitItemRequestBuilder and sets the default values.
func NewAdministrativeUnitItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdministrativeUnitItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdministrativeUnitItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property administrativeUnits for directory
func (m *AdministrativeUnitItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property administrativeUnits for directory
func (m *AdministrativeUnitItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *AdministrativeUnitItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation conceptual container for user and group directory objects.
func (m *AdministrativeUnitItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration conceptual container for user and group directory objects.
func (m *AdministrativeUnitItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AdministrativeUnitItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property administrativeUnits in directory
func (m *AdministrativeUnitItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministrativeUnitable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property administrativeUnits in directory
func (m *AdministrativeUnitItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministrativeUnitable, requestConfiguration *AdministrativeUnitItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property administrativeUnits for directory
func (m *AdministrativeUnitItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AdministrativeUnitItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Extensions the extensions property
func (m *AdministrativeUnitItemRequestBuilder) Extensions()(*ib3392aac80c6073c1b4cb66c1c9a59ef64a733b53d9fbe3bb2fdeea43b1e8ed0.ExtensionsRequestBuilder) {
    return ib3392aac80c6073c1b4cb66c1c9a59ef64a733b53d9fbe3bb2fdeea43b1e8ed0.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.administrativeUnits.item.extensions.item collection
func (m *AdministrativeUnitItemRequestBuilder) ExtensionsById(id string)(*i348fc6817b120c2ca4453cfe220f3a31e3dd76b71234d51b3cf59448c038898e.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i348fc6817b120c2ca4453cfe220f3a31e3dd76b71234d51b3cf59448c038898e.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get conceptual container for user and group directory objects.
func (m *AdministrativeUnitItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AdministrativeUnitItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministrativeUnitable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAdministrativeUnitFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministrativeUnitable), nil
}
// GetMemberGroups the getMemberGroups property
func (m *AdministrativeUnitItemRequestBuilder) GetMemberGroups()(*ib744a95755b350bc30b575b7a619bbb35c32385d740a0feca8f31cddc0e3af5d.GetMemberGroupsRequestBuilder) {
    return ib744a95755b350bc30b575b7a619bbb35c32385d740a0feca8f31cddc0e3af5d.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *AdministrativeUnitItemRequestBuilder) GetMemberObjects()(*i162499125dee18a8c32afea0c335b92a83c8bfbf514f6c85bec33a36bdaf7ca3.GetMemberObjectsRequestBuilder) {
    return i162499125dee18a8c32afea0c335b92a83c8bfbf514f6c85bec33a36bdaf7ca3.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Members the members property
func (m *AdministrativeUnitItemRequestBuilder) Members()(*i9a39d6011e7f6e52d0463e43f4affbe7fe566deec74cb2cfa89d0164876f73fb.MembersRequestBuilder) {
    return i9a39d6011e7f6e52d0463e43f4affbe7fe566deec74cb2cfa89d0164876f73fb.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.administrativeUnits.item.members.item collection
func (m *AdministrativeUnitItemRequestBuilder) MembersById(id string)(*id21ccd20a38c8c744ae37c53448fd59c79e0a14764716cf4113d7e8fd6f62db0.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return id21ccd20a38c8c744ae37c53448fd59c79e0a14764716cf4113d7e8fd6f62db0.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property administrativeUnits in directory
func (m *AdministrativeUnitItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministrativeUnitable, requestConfiguration *AdministrativeUnitItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Restore the restore property
func (m *AdministrativeUnitItemRequestBuilder) Restore()(*i8dcb8bc7850e64bb1fd90771daf6714da6fc512dd4adae05039d541e9bb9419c.RestoreRequestBuilder) {
    return i8dcb8bc7850e64bb1fd90771daf6714da6fc512dd4adae05039d541e9bb9419c.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedRoleMembers the scopedRoleMembers property
func (m *AdministrativeUnitItemRequestBuilder) ScopedRoleMembers()(*i29c75035d87e1e3a958b91c43025f79430cd4c9d3b0cc664ffe065bc407fda28.ScopedRoleMembersRequestBuilder) {
    return i29c75035d87e1e3a958b91c43025f79430cd4c9d3b0cc664ffe065bc407fda28.NewScopedRoleMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedRoleMembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.administrativeUnits.item.scopedRoleMembers.item collection
func (m *AdministrativeUnitItemRequestBuilder) ScopedRoleMembersById(id string)(*ibb066ad55db4b7866d16a995539af863ab771176674fded378a94af09016fc1b.ScopedRoleMembershipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership%2Did"] = id
    }
    return ibb066ad55db4b7866d16a995539af863ab771176674fded378a94af09016fc1b.NewScopedRoleMembershipItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
