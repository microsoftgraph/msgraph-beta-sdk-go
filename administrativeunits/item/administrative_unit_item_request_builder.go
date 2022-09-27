package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0f36168ca30e64e3bea01bb77d0f73154f5020719384d2f3ab434667a1277d0c "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/scopedrolemembers"
    i1fa710cc359019f3bf0614158e392cfc3471c6427ce07ada8e4627100cec5a82 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/members"
    i349829ba4fce29dd7987c6df11dbcaf800247af0fa3219b01dcdc6e1343f1956 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/checkmembergroups"
    i4e5a3bd6a5396f5b0939bfce85a44230296cab8d3446ee6e3a5a29204173428d "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/restore"
    ia216e040c36eeacd190ee60a98d2779e9f151c69020cf34b69a5a127c5a6db73 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/getmembergroups"
    ib099fb63054d0dd1784a47669d75c804d03b727cd11a2480311fa2d37e24ed6c "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/checkmemberobjects"
    ic10bcb37ee210a709c1b4b767679d61e4e281098278ef4698ae36b4863f0598c "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/extensions"
    if3ef57b69f282bfad9e1e97bd91a5b30f4da893d66532d98d784f6d054b6fb74 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/getmemberobjects"
    ia0eae455c54ecade94b5dc033e0f792d508d4ae4d0e22ed0d19c83e392692780 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/scopedrolemembers/item"
    ide3d97019bbb4651e57fb602931f7d0fab23c99c26f2f79b556726ee7c595db8 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/members/item"
    ieb753c78202e5499a95653d7cfff3918b3c7124418d5e65033612f2c7fbc09c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/extensions/item"
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
    Headers map[string]string
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
func (m *AdministrativeUnitItemRequestBuilder) CheckMemberGroups()(*i349829ba4fce29dd7987c6df11dbcaf800247af0fa3219b01dcdc6e1343f1956.CheckMemberGroupsRequestBuilder) {
    return i349829ba4fce29dd7987c6df11dbcaf800247af0fa3219b01dcdc6e1343f1956.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *AdministrativeUnitItemRequestBuilder) CheckMemberObjects()(*ib099fb63054d0dd1784a47669d75c804d03b727cd11a2480311fa2d37e24ed6c.CheckMemberObjectsRequestBuilder) {
    return ib099fb63054d0dd1784a47669d75c804d03b727cd11a2480311fa2d37e24ed6c.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// CreateDeleteRequestInformation delete an administrativeUnit.
func (m *AdministrativeUnitItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete an administrativeUnit.
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
// CreateGetRequestInformation retrieve the properties and relationships of an administrativeUnit object. Since the **administrativeUnit** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in an **administrativeUnit** instance.
func (m *AdministrativeUnitItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration retrieve the properties and relationships of an administrativeUnit object. Since the **administrativeUnit** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in an **administrativeUnit** instance.
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
// CreatePatchRequestInformation update the properties of an administrativeUnit object.
func (m *AdministrativeUnitItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministrativeUnitable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the properties of an administrativeUnit object.
func (m *AdministrativeUnitItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministrativeUnitable, requestConfiguration *AdministrativeUnitItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete an administrativeUnit.
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
func (m *AdministrativeUnitItemRequestBuilder) Extensions()(*ic10bcb37ee210a709c1b4b767679d61e4e281098278ef4698ae36b4863f0598c.ExtensionsRequestBuilder) {
    return ic10bcb37ee210a709c1b4b767679d61e4e281098278ef4698ae36b4863f0598c.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.administrativeUnits.item.extensions.item collection
func (m *AdministrativeUnitItemRequestBuilder) ExtensionsById(id string)(*ieb753c78202e5499a95653d7cfff3918b3c7124418d5e65033612f2c7fbc09c2.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ieb753c78202e5499a95653d7cfff3918b3c7124418d5e65033612f2c7fbc09c2.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get retrieve the properties and relationships of an administrativeUnit object. Since the **administrativeUnit** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in an **administrativeUnit** instance.
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
func (m *AdministrativeUnitItemRequestBuilder) GetMemberGroups()(*ia216e040c36eeacd190ee60a98d2779e9f151c69020cf34b69a5a127c5a6db73.GetMemberGroupsRequestBuilder) {
    return ia216e040c36eeacd190ee60a98d2779e9f151c69020cf34b69a5a127c5a6db73.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *AdministrativeUnitItemRequestBuilder) GetMemberObjects()(*if3ef57b69f282bfad9e1e97bd91a5b30f4da893d66532d98d784f6d054b6fb74.GetMemberObjectsRequestBuilder) {
    return if3ef57b69f282bfad9e1e97bd91a5b30f4da893d66532d98d784f6d054b6fb74.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Members the members property
func (m *AdministrativeUnitItemRequestBuilder) Members()(*i1fa710cc359019f3bf0614158e392cfc3471c6427ce07ada8e4627100cec5a82.MembersRequestBuilder) {
    return i1fa710cc359019f3bf0614158e392cfc3471c6427ce07ada8e4627100cec5a82.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.administrativeUnits.item.members.item collection
func (m *AdministrativeUnitItemRequestBuilder) MembersById(id string)(*ide3d97019bbb4651e57fb602931f7d0fab23c99c26f2f79b556726ee7c595db8.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return ide3d97019bbb4651e57fb602931f7d0fab23c99c26f2f79b556726ee7c595db8.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the properties of an administrativeUnit object.
func (m *AdministrativeUnitItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministrativeUnitable, requestConfiguration *AdministrativeUnitItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministrativeUnitable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// Restore the restore property
func (m *AdministrativeUnitItemRequestBuilder) Restore()(*i4e5a3bd6a5396f5b0939bfce85a44230296cab8d3446ee6e3a5a29204173428d.RestoreRequestBuilder) {
    return i4e5a3bd6a5396f5b0939bfce85a44230296cab8d3446ee6e3a5a29204173428d.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedRoleMembers the scopedRoleMembers property
func (m *AdministrativeUnitItemRequestBuilder) ScopedRoleMembers()(*i0f36168ca30e64e3bea01bb77d0f73154f5020719384d2f3ab434667a1277d0c.ScopedRoleMembersRequestBuilder) {
    return i0f36168ca30e64e3bea01bb77d0f73154f5020719384d2f3ab434667a1277d0c.NewScopedRoleMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedRoleMembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.administrativeUnits.item.scopedRoleMembers.item collection
func (m *AdministrativeUnitItemRequestBuilder) ScopedRoleMembersById(id string)(*ia0eae455c54ecade94b5dc033e0f792d508d4ae4d0e22ed0d19c83e392692780.ScopedRoleMembershipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership%2Did"] = id
    }
    return ia0eae455c54ecade94b5dc033e0f792d508d4ae4d0e22ed0d19c83e392692780.NewScopedRoleMembershipItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
