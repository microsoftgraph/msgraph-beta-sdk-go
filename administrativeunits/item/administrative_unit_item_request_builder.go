package item

import (
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AdministrativeUnitItemRequestBuilderDeleteOptions options for Delete
type AdministrativeUnitItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AdministrativeUnitItemRequestBuilderGetOptions options for Get
type AdministrativeUnitItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *AdministrativeUnitItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AdministrativeUnitItemRequestBuilderGetQueryParameters get entity from administrativeUnits by key
type AdministrativeUnitItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AdministrativeUnitItemRequestBuilderPatchOptions options for Patch
type AdministrativeUnitItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministrativeUnitable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
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
    m.urlTemplate = "{+baseurl}/administrativeUnits/{administrativeUnit_id}{?select,expand}";
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
// CreateDeleteRequestInformation delete entity from administrativeUnits
func (m *AdministrativeUnitItemRequestBuilder) CreateDeleteRequestInformation(options *AdministrativeUnitItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get entity from administrativeUnits by key
func (m *AdministrativeUnitItemRequestBuilder) CreateGetRequestInformation(options *AdministrativeUnitItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update entity in administrativeUnits
func (m *AdministrativeUnitItemRequestBuilder) CreatePatchRequestInformation(options *AdministrativeUnitItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete entity from administrativeUnits
func (m *AdministrativeUnitItemRequestBuilder) Delete(options *AdministrativeUnitItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
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
        urlTplParams["extension_id"] = id
    }
    return ieb753c78202e5499a95653d7cfff3918b3c7124418d5e65033612f2c7fbc09c2.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get entity from administrativeUnits by key
func (m *AdministrativeUnitItemRequestBuilder) Get(options *AdministrativeUnitItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministrativeUnitable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAdministrativeUnitFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
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
        urlTplParams["directoryObject_id"] = id
    }
    return ide3d97019bbb4651e57fb602931f7d0fab23c99c26f2f79b556726ee7c595db8.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update entity in administrativeUnits
func (m *AdministrativeUnitItemRequestBuilder) Patch(options *AdministrativeUnitItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
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
        urlTplParams["scopedRoleMembership_id"] = id
    }
    return ia0eae455c54ecade94b5dc033e0f792d508d4ae4d0e22ed0d19c83e392692780.NewScopedRoleMembershipItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
