package group

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i113a0834c4c20c1cdcea6f1313099440a249b429741f60f4d54ea8477690dc9b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberof/item/group/subscribebymail"
    i18c585357c5e00ecfe7a802ebe4a0284dd317834ec613f19ee1a2b0ef620f5b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberof/item/group/checkmemberobjects"
    i2d1c59fd9ba5f6aa98b411ee6108fb9c75a208eb6537b469d143c2268dd61ef1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberof/item/group/evaluatedynamicmembership"
    i2da5a409b8a0957ef4d680525e4929d7b0aaa2f10d1f994c43fc0dc019ca7ce9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberof/item/group/checkgrantedpermissionsforapp"
    i31865e45fb1eb1c06b1619b2e52705f849f403a5da5578e9b598ce35d8b42475 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberof/item/group/addfavorite"
    i5477b950114d48b52a734b659ec486d04e29fadd1a08b305d0468184293c499b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberof/item/group/removefavorite"
    i5972db511d57896eed439e3bbb569e77685fa12a70d1ec93478b485a76722223 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberof/item/group/getmembergroups"
    i620c35d7ae31a5b2ec36b6dc849b576a8263f975da9af8e32b4a27f8c534b36f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberof/item/group/renew"
    i79ff5a6ccb53c64c9acbffd3801d0bf040075f505b2695e36344ba4f2fa6e2f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberof/item/group/getmemberobjects"
    i94e9ff0ac1f40f388a059c93abf5ea6b9c75670872ce9c191c1eeee62dbbca1a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberof/item/group/checkmembergroups"
    ia8431c38111c17494e9540569170e5cae8cebe11c66609f22920dca4985f0893 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberof/item/group/unsubscribebymail"
    ib206a21134ce1dd20c75c45e85261a7b295426dcaa4accab852780434cc8adfc "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberof/item/group/assignlicense"
    ib8c1831dfb795382ed55a779855bc0bc7c3ce68d2de81578fa900993e731dc99 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberof/item/group/restore"
    ie234ab14458a38c7333348821a23c2b99297a5b7ae104e918ce22621b69e2a19 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberof/item/group/resetunseencount"
    iee25434fdc149327b735c70e9dfab72d7aeced980e5648cb350b5b968f1f7910 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberof/item/group/validateproperties"
)

// GroupRequestBuilder casts the previous resource to group.
type GroupRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupRequestBuilderGetQueryParameters get the item of type microsoft.graph.directoryObject as microsoft.graph.group
type GroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupRequestBuilderGetQueryParameters
}
// AddFavorite the addFavorite property
func (m *GroupRequestBuilder) AddFavorite()(*i31865e45fb1eb1c06b1619b2e52705f849f403a5da5578e9b598ce35d8b42475.AddFavoriteRequestBuilder) {
    return i31865e45fb1eb1c06b1619b2e52705f849f403a5da5578e9b598ce35d8b42475.NewAddFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *GroupRequestBuilder) AssignLicense()(*ib206a21134ce1dd20c75c45e85261a7b295426dcaa4accab852780434cc8adfc.AssignLicenseRequestBuilder) {
    return ib206a21134ce1dd20c75c45e85261a7b295426dcaa4accab852780434cc8adfc.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckGrantedPermissionsForApp the checkGrantedPermissionsForApp property
func (m *GroupRequestBuilder) CheckGrantedPermissionsForApp()(*i2da5a409b8a0957ef4d680525e4929d7b0aaa2f10d1f994c43fc0dc019ca7ce9.CheckGrantedPermissionsForAppRequestBuilder) {
    return i2da5a409b8a0957ef4d680525e4929d7b0aaa2f10d1f994c43fc0dc019ca7ce9.NewCheckGrantedPermissionsForAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *GroupRequestBuilder) CheckMemberGroups()(*i94e9ff0ac1f40f388a059c93abf5ea6b9c75670872ce9c191c1eeee62dbbca1a.CheckMemberGroupsRequestBuilder) {
    return i94e9ff0ac1f40f388a059c93abf5ea6b9c75670872ce9c191c1eeee62dbbca1a.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *GroupRequestBuilder) CheckMemberObjects()(*i18c585357c5e00ecfe7a802ebe4a0284dd317834ec613f19ee1a2b0ef620f5b4.CheckMemberObjectsRequestBuilder) {
    return i18c585357c5e00ecfe7a802ebe4a0284dd317834ec613f19ee1a2b0ef620f5b4.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupRequestBuilderInternal instantiates a new GroupRequestBuilder and sets the default values.
func NewGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupRequestBuilder) {
    m := &GroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/memberOf/{directoryObject%2Did}/microsoft.graph.group{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupRequestBuilder instantiates a new GroupRequestBuilder and sets the default values.
func NewGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get the item of type microsoft.graph.directoryObject as microsoft.graph.group
func (m *GroupRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get the item of type microsoft.graph.directoryObject as microsoft.graph.group
func (m *GroupRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GroupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// EvaluateDynamicMembership the evaluateDynamicMembership property
func (m *GroupRequestBuilder) EvaluateDynamicMembership()(*i2d1c59fd9ba5f6aa98b411ee6108fb9c75a208eb6537b469d143c2268dd61ef1.EvaluateDynamicMembershipRequestBuilder) {
    return i2d1c59fd9ba5f6aa98b411ee6108fb9c75a208eb6537b469d143c2268dd61ef1.NewEvaluateDynamicMembershipRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.group
func (m *GroupRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMemberGroups the getMemberGroups property
func (m *GroupRequestBuilder) GetMemberGroups()(*i5972db511d57896eed439e3bbb569e77685fa12a70d1ec93478b485a76722223.GetMemberGroupsRequestBuilder) {
    return i5972db511d57896eed439e3bbb569e77685fa12a70d1ec93478b485a76722223.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *GroupRequestBuilder) GetMemberObjects()(*i79ff5a6ccb53c64c9acbffd3801d0bf040075f505b2695e36344ba4f2fa6e2f9.GetMemberObjectsRequestBuilder) {
    return i79ff5a6ccb53c64c9acbffd3801d0bf040075f505b2695e36344ba4f2fa6e2f9.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler get the item of type microsoft.graph.directoryObject as microsoft.graph.group
func (m *GroupRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *GroupRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable), nil
}
// RemoveFavorite the removeFavorite property
func (m *GroupRequestBuilder) RemoveFavorite()(*i5477b950114d48b52a734b659ec486d04e29fadd1a08b305d0468184293c499b.RemoveFavoriteRequestBuilder) {
    return i5477b950114d48b52a734b659ec486d04e29fadd1a08b305d0468184293c499b.NewRemoveFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Renew the renew property
func (m *GroupRequestBuilder) Renew()(*i620c35d7ae31a5b2ec36b6dc849b576a8263f975da9af8e32b4a27f8c534b36f.RenewRequestBuilder) {
    return i620c35d7ae31a5b2ec36b6dc849b576a8263f975da9af8e32b4a27f8c534b36f.NewRenewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetUnseenCount the resetUnseenCount property
func (m *GroupRequestBuilder) ResetUnseenCount()(*ie234ab14458a38c7333348821a23c2b99297a5b7ae104e918ce22621b69e2a19.ResetUnseenCountRequestBuilder) {
    return ie234ab14458a38c7333348821a23c2b99297a5b7ae104e918ce22621b69e2a19.NewResetUnseenCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *GroupRequestBuilder) Restore()(*ib8c1831dfb795382ed55a779855bc0bc7c3ce68d2de81578fa900993e731dc99.RestoreRequestBuilder) {
    return ib8c1831dfb795382ed55a779855bc0bc7c3ce68d2de81578fa900993e731dc99.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscribeByMail the subscribeByMail property
func (m *GroupRequestBuilder) SubscribeByMail()(*i113a0834c4c20c1cdcea6f1313099440a249b429741f60f4d54ea8477690dc9b.SubscribeByMailRequestBuilder) {
    return i113a0834c4c20c1cdcea6f1313099440a249b429741f60f4d54ea8477690dc9b.NewSubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnsubscribeByMail the unsubscribeByMail property
func (m *GroupRequestBuilder) UnsubscribeByMail()(*ia8431c38111c17494e9540569170e5cae8cebe11c66609f22920dca4985f0893.UnsubscribeByMailRequestBuilder) {
    return ia8431c38111c17494e9540569170e5cae8cebe11c66609f22920dca4985f0893.NewUnsubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties the validateProperties property
func (m *GroupRequestBuilder) ValidateProperties()(*iee25434fdc149327b735c70e9dfab72d7aeced980e5648cb350b5b968f1f7910.ValidatePropertiesRequestBuilder) {
    return iee25434fdc149327b735c70e9dfab72d7aeced980e5648cb350b5b968f1f7910.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
