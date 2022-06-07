package group

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0cd9cfbe684ee0d081656b6b044aa404e92a6fee67e60faa808f1b24c4926751 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item/group/renew"
    i0d3c3634d2d28af7e8f97f199e7b5219edbbf26fc8438334e8991a5f40817a6a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item/group/addfavorite"
    i26753c341fc34b82e26241757b8843d8fab85ca8e4fca9363142f4d32f3a6251 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item/group/unsubscribebymail"
    i3377e51d162dea85f5bbd70a21bc1a9145ca6595d25dc3cb0a66d9943dbff189 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item/group/removefavorite"
    i3b315585b2a99ed6d9b8586a4f29b32c5d75a958e5cfcc10866a6bbe3e60b853 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item/group/getmembergroups"
    i44a8a7ac67a8aeffa6167f1984f4f9260ca55a2d506f4ae5e4178612f681e76b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item/group/resetunseencount"
    i602ddb8310bf992ae6454519eba55ee23eb7b0d3a908ea05421cd8f168d5df0c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item/group/restore"
    i610f4f48e386835d59d5599084cce62669d3c320100fe0ae8d8d8455d92eea5b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item/group/validateproperties"
    i6573227eb120c9b3cf38211cd18c4de2fd567b450d0053f90b21a809d92f2a71 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item/group/subscribebymail"
    i91287d254279e5725feddff5a826cbb74e829f3c7947aaa9e89d691132c22c8a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item/group/checkmemberobjects"
    iba2e337a141d5a7bc14d5483e07dde2657f0a5491e2b002256f65f64c2d8719f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item/group/assignlicense"
    ica7841d2df637caa7f439d686b91bff67860f2c4c6d95f7cf8ac1947b6502d4f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item/group/checkmembergroups"
    idfaea751205908ca32de7997c39bad502db5c054fbd9df7b8bdbb5593ecff6b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item/group/checkgrantedpermissionsforapp"
    ie55159ee857a5315e85e695eb95c39368c1063c9d50ef28d974a296d9265387e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item/group/evaluatedynamicmembership"
    ied2ed1ce180917e9674884aa66f019f53c74faec8652ee7441a0e55f970f708e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item/group/getmemberobjects"
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
func (m *GroupRequestBuilder) AddFavorite()(*i0d3c3634d2d28af7e8f97f199e7b5219edbbf26fc8438334e8991a5f40817a6a.AddFavoriteRequestBuilder) {
    return i0d3c3634d2d28af7e8f97f199e7b5219edbbf26fc8438334e8991a5f40817a6a.NewAddFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *GroupRequestBuilder) AssignLicense()(*iba2e337a141d5a7bc14d5483e07dde2657f0a5491e2b002256f65f64c2d8719f.AssignLicenseRequestBuilder) {
    return iba2e337a141d5a7bc14d5483e07dde2657f0a5491e2b002256f65f64c2d8719f.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckGrantedPermissionsForApp the checkGrantedPermissionsForApp property
func (m *GroupRequestBuilder) CheckGrantedPermissionsForApp()(*idfaea751205908ca32de7997c39bad502db5c054fbd9df7b8bdbb5593ecff6b2.CheckGrantedPermissionsForAppRequestBuilder) {
    return idfaea751205908ca32de7997c39bad502db5c054fbd9df7b8bdbb5593ecff6b2.NewCheckGrantedPermissionsForAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *GroupRequestBuilder) CheckMemberGroups()(*ica7841d2df637caa7f439d686b91bff67860f2c4c6d95f7cf8ac1947b6502d4f.CheckMemberGroupsRequestBuilder) {
    return ica7841d2df637caa7f439d686b91bff67860f2c4c6d95f7cf8ac1947b6502d4f.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *GroupRequestBuilder) CheckMemberObjects()(*i91287d254279e5725feddff5a826cbb74e829f3c7947aaa9e89d691132c22c8a.CheckMemberObjectsRequestBuilder) {
    return i91287d254279e5725feddff5a826cbb74e829f3c7947aaa9e89d691132c22c8a.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupRequestBuilderInternal instantiates a new GroupRequestBuilder and sets the default values.
func NewGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupRequestBuilder) {
    m := &GroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/members/{directoryObject%2Did}/microsoft.graph.group{?%24select,%24expand}";
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
func (m *GroupRequestBuilder) EvaluateDynamicMembership()(*ie55159ee857a5315e85e695eb95c39368c1063c9d50ef28d974a296d9265387e.EvaluateDynamicMembershipRequestBuilder) {
    return ie55159ee857a5315e85e695eb95c39368c1063c9d50ef28d974a296d9265387e.NewEvaluateDynamicMembershipRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.group
func (m *GroupRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMemberGroups the getMemberGroups property
func (m *GroupRequestBuilder) GetMemberGroups()(*i3b315585b2a99ed6d9b8586a4f29b32c5d75a958e5cfcc10866a6bbe3e60b853.GetMemberGroupsRequestBuilder) {
    return i3b315585b2a99ed6d9b8586a4f29b32c5d75a958e5cfcc10866a6bbe3e60b853.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *GroupRequestBuilder) GetMemberObjects()(*ied2ed1ce180917e9674884aa66f019f53c74faec8652ee7441a0e55f970f708e.GetMemberObjectsRequestBuilder) {
    return ied2ed1ce180917e9674884aa66f019f53c74faec8652ee7441a0e55f970f708e.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *GroupRequestBuilder) RemoveFavorite()(*i3377e51d162dea85f5bbd70a21bc1a9145ca6595d25dc3cb0a66d9943dbff189.RemoveFavoriteRequestBuilder) {
    return i3377e51d162dea85f5bbd70a21bc1a9145ca6595d25dc3cb0a66d9943dbff189.NewRemoveFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Renew the renew property
func (m *GroupRequestBuilder) Renew()(*i0cd9cfbe684ee0d081656b6b044aa404e92a6fee67e60faa808f1b24c4926751.RenewRequestBuilder) {
    return i0cd9cfbe684ee0d081656b6b044aa404e92a6fee67e60faa808f1b24c4926751.NewRenewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetUnseenCount the resetUnseenCount property
func (m *GroupRequestBuilder) ResetUnseenCount()(*i44a8a7ac67a8aeffa6167f1984f4f9260ca55a2d506f4ae5e4178612f681e76b.ResetUnseenCountRequestBuilder) {
    return i44a8a7ac67a8aeffa6167f1984f4f9260ca55a2d506f4ae5e4178612f681e76b.NewResetUnseenCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *GroupRequestBuilder) Restore()(*i602ddb8310bf992ae6454519eba55ee23eb7b0d3a908ea05421cd8f168d5df0c.RestoreRequestBuilder) {
    return i602ddb8310bf992ae6454519eba55ee23eb7b0d3a908ea05421cd8f168d5df0c.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscribeByMail the subscribeByMail property
func (m *GroupRequestBuilder) SubscribeByMail()(*i6573227eb120c9b3cf38211cd18c4de2fd567b450d0053f90b21a809d92f2a71.SubscribeByMailRequestBuilder) {
    return i6573227eb120c9b3cf38211cd18c4de2fd567b450d0053f90b21a809d92f2a71.NewSubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnsubscribeByMail the unsubscribeByMail property
func (m *GroupRequestBuilder) UnsubscribeByMail()(*i26753c341fc34b82e26241757b8843d8fab85ca8e4fca9363142f4d32f3a6251.UnsubscribeByMailRequestBuilder) {
    return i26753c341fc34b82e26241757b8843d8fab85ca8e4fca9363142f4d32f3a6251.NewUnsubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties the validateProperties property
func (m *GroupRequestBuilder) ValidateProperties()(*i610f4f48e386835d59d5599084cce62669d3c320100fe0ae8d8d8455d92eea5b.ValidatePropertiesRequestBuilder) {
    return i610f4f48e386835d59d5599084cce62669d3c320100fe0ae8d8d8455d92eea5b.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
