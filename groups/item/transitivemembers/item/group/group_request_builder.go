package group

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i03cfa61a81909438f028781e4b7ac3565a57ac14b89b07971e0157cfd406dea4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/item/group/unsubscribebymail"
    i066edf936e6b5a813b1df769792f6bce474030ded4a2be056d0da3808cc7e80a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/item/group/checkgrantedpermissionsforapp"
    i23bafc51dd52fdd67340aaacfca8dabaa784f8384abd856528f3caba51ab4257 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/item/group/resetunseencount"
    i2412e4c3f296b0fef04f16a80ac1785de6794aca002536303152042746717b3b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/item/group/getmembergroups"
    i461a3b5ebb8db72f0662796dfd53cd2fa75496956951cc01df2e2b48d8ff0db0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/item/group/assignlicense"
    i52e3bc57ae007d4f94003cb2089e52c9df36ab21abf403c24d41bf46a45ebd2a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/item/group/checkmembergroups"
    i6ea03a88cb25cc4c9b63b1fcdcdeb3420699cf69a02a625536b2655ae4e953fc "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/item/group/getmemberobjects"
    i8ba605edfb2242ff37fbbd10aa395f6330789d44d5025b73063bccef33aa4af0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/item/group/evaluatedynamicmembership"
    i90422894c9f9144c44644c14a88f4d2913353ec943067cfdaa624af13130cead "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/item/group/checkmemberobjects"
    i91a5b1cdbb34bef9e6974a354522b436a8454c57458b844c92b1b0ca01fbb5a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/item/group/validateproperties"
    ia970d2e5f1c114071e3deb123135c55ac626cc495bcf5f26b90bec441bcd0e50 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/item/group/subscribebymail"
    ib0cf3790c46575bdf1ff1e79117c0c34d0a32b86c8c913e6aae735d33f5bc6cb "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/item/group/restore"
    ic2b20e6df5c13c1b61444a0dd6231aa249f27092d549a6d7016e76d3e86eb377 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/item/group/renew"
    ic4e72874ce5e1a87014323049c102655804cdd58281a1aa5b2f04d0e046aa1b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/item/group/addfavorite"
    iff81518c6e73f7ca85375b4c55387f53b0481f23a8bf81f8a201fb9e99c7971d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivemembers/item/group/removefavorite"
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
func (m *GroupRequestBuilder) AddFavorite()(*ic4e72874ce5e1a87014323049c102655804cdd58281a1aa5b2f04d0e046aa1b8.AddFavoriteRequestBuilder) {
    return ic4e72874ce5e1a87014323049c102655804cdd58281a1aa5b2f04d0e046aa1b8.NewAddFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *GroupRequestBuilder) AssignLicense()(*i461a3b5ebb8db72f0662796dfd53cd2fa75496956951cc01df2e2b48d8ff0db0.AssignLicenseRequestBuilder) {
    return i461a3b5ebb8db72f0662796dfd53cd2fa75496956951cc01df2e2b48d8ff0db0.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckGrantedPermissionsForApp the checkGrantedPermissionsForApp property
func (m *GroupRequestBuilder) CheckGrantedPermissionsForApp()(*i066edf936e6b5a813b1df769792f6bce474030ded4a2be056d0da3808cc7e80a.CheckGrantedPermissionsForAppRequestBuilder) {
    return i066edf936e6b5a813b1df769792f6bce474030ded4a2be056d0da3808cc7e80a.NewCheckGrantedPermissionsForAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *GroupRequestBuilder) CheckMemberGroups()(*i52e3bc57ae007d4f94003cb2089e52c9df36ab21abf403c24d41bf46a45ebd2a.CheckMemberGroupsRequestBuilder) {
    return i52e3bc57ae007d4f94003cb2089e52c9df36ab21abf403c24d41bf46a45ebd2a.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *GroupRequestBuilder) CheckMemberObjects()(*i90422894c9f9144c44644c14a88f4d2913353ec943067cfdaa624af13130cead.CheckMemberObjectsRequestBuilder) {
    return i90422894c9f9144c44644c14a88f4d2913353ec943067cfdaa624af13130cead.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupRequestBuilderInternal instantiates a new GroupRequestBuilder and sets the default values.
func NewGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupRequestBuilder) {
    m := &GroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/transitiveMembers/{directoryObject%2Did}/microsoft.graph.group{?%24select,%24expand}";
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
func (m *GroupRequestBuilder) EvaluateDynamicMembership()(*i8ba605edfb2242ff37fbbd10aa395f6330789d44d5025b73063bccef33aa4af0.EvaluateDynamicMembershipRequestBuilder) {
    return i8ba605edfb2242ff37fbbd10aa395f6330789d44d5025b73063bccef33aa4af0.NewEvaluateDynamicMembershipRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.group
func (m *GroupRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMemberGroups the getMemberGroups property
func (m *GroupRequestBuilder) GetMemberGroups()(*i2412e4c3f296b0fef04f16a80ac1785de6794aca002536303152042746717b3b.GetMemberGroupsRequestBuilder) {
    return i2412e4c3f296b0fef04f16a80ac1785de6794aca002536303152042746717b3b.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *GroupRequestBuilder) GetMemberObjects()(*i6ea03a88cb25cc4c9b63b1fcdcdeb3420699cf69a02a625536b2655ae4e953fc.GetMemberObjectsRequestBuilder) {
    return i6ea03a88cb25cc4c9b63b1fcdcdeb3420699cf69a02a625536b2655ae4e953fc.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *GroupRequestBuilder) RemoveFavorite()(*iff81518c6e73f7ca85375b4c55387f53b0481f23a8bf81f8a201fb9e99c7971d.RemoveFavoriteRequestBuilder) {
    return iff81518c6e73f7ca85375b4c55387f53b0481f23a8bf81f8a201fb9e99c7971d.NewRemoveFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Renew the renew property
func (m *GroupRequestBuilder) Renew()(*ic2b20e6df5c13c1b61444a0dd6231aa249f27092d549a6d7016e76d3e86eb377.RenewRequestBuilder) {
    return ic2b20e6df5c13c1b61444a0dd6231aa249f27092d549a6d7016e76d3e86eb377.NewRenewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetUnseenCount the resetUnseenCount property
func (m *GroupRequestBuilder) ResetUnseenCount()(*i23bafc51dd52fdd67340aaacfca8dabaa784f8384abd856528f3caba51ab4257.ResetUnseenCountRequestBuilder) {
    return i23bafc51dd52fdd67340aaacfca8dabaa784f8384abd856528f3caba51ab4257.NewResetUnseenCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *GroupRequestBuilder) Restore()(*ib0cf3790c46575bdf1ff1e79117c0c34d0a32b86c8c913e6aae735d33f5bc6cb.RestoreRequestBuilder) {
    return ib0cf3790c46575bdf1ff1e79117c0c34d0a32b86c8c913e6aae735d33f5bc6cb.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscribeByMail the subscribeByMail property
func (m *GroupRequestBuilder) SubscribeByMail()(*ia970d2e5f1c114071e3deb123135c55ac626cc495bcf5f26b90bec441bcd0e50.SubscribeByMailRequestBuilder) {
    return ia970d2e5f1c114071e3deb123135c55ac626cc495bcf5f26b90bec441bcd0e50.NewSubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnsubscribeByMail the unsubscribeByMail property
func (m *GroupRequestBuilder) UnsubscribeByMail()(*i03cfa61a81909438f028781e4b7ac3565a57ac14b89b07971e0157cfd406dea4.UnsubscribeByMailRequestBuilder) {
    return i03cfa61a81909438f028781e4b7ac3565a57ac14b89b07971e0157cfd406dea4.NewUnsubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties the validateProperties property
func (m *GroupRequestBuilder) ValidateProperties()(*i91a5b1cdbb34bef9e6974a354522b436a8454c57458b844c92b1b0ca01fbb5a0.ValidatePropertiesRequestBuilder) {
    return i91a5b1cdbb34bef9e6974a354522b436a8454c57458b844c92b1b0ca01fbb5a0.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
