package group

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i03ecb7ca34f92b5605d7245fc90dc2c0b72dc43f58a22a6c7e21f723dad5b2d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/item/group/restore"
    i17092f7261264a2f0255a5ac04f42580fce1e959fea0a681c4a91bf0638fb515 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/item/group/getmembergroups"
    i1c539874b822563762684a2fbf019b0fd92c430ebadb0ff0d8b8345f60438d48 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/item/group/unsubscribebymail"
    i26f9002a7d2e4776787df240ec14cad75b3b2a1686be3f1127a6bf92ffeab099 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/item/group/addfavorite"
    i2a2bd0ca2efee011e8eaaea4a0bb2124144ad7bbb11247260d59f91238a429c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/item/group/subscribebymail"
    i3efec48d3704d9e37ec0caf7bf7392c1cc9544203611a18ac0474ebe63248ac8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/item/group/getmemberobjects"
    i5ba239b7b916932a59648738c9c857f641c1a79a37ca8b919003cdf2b2b425a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/item/group/checkmemberobjects"
    i6f733851fd18148e44adbae5ec3e14e4daf59da5accd703f72885d24108bbaf7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/item/group/validateproperties"
    i7ba94979f26931d406d3be0e7542846413bff6fa16c87f2d4341f65531091785 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/item/group/removefavorite"
    i876930b27688b7989eecbe8cc0bb8d08ce2346b975da5a1bb6f724bfe0d55a0d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/item/group/resetunseencount"
    i99671a59b4f1302345143fa82c5025ff0fa9b31e6a5d9ea0a738887f79d0f5cb "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/item/group/renew"
    ib77b6ef94b8f11e8e3447f3eb6adca4cada604440e45269b0dbe11880a0dab37 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/item/group/checkmembergroups"
    icf7be9f0ea3a07f80dca0715b03d3296a783331210a41a8f4c89d2d2d0433f13 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/item/group/assignlicense"
    id71215f18227366f93e6a7751d0838fa8e720356a7b68ce2350d10c20c658a8b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/item/group/evaluatedynamicmembership"
    ief4021e7b13991e6921047b490cf96ebd1b3f9b8cb51f53f88e6d22b13625aca "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/item/group/checkgrantedpermissionsforapp"
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
func (m *GroupRequestBuilder) AddFavorite()(*i26f9002a7d2e4776787df240ec14cad75b3b2a1686be3f1127a6bf92ffeab099.AddFavoriteRequestBuilder) {
    return i26f9002a7d2e4776787df240ec14cad75b3b2a1686be3f1127a6bf92ffeab099.NewAddFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *GroupRequestBuilder) AssignLicense()(*icf7be9f0ea3a07f80dca0715b03d3296a783331210a41a8f4c89d2d2d0433f13.AssignLicenseRequestBuilder) {
    return icf7be9f0ea3a07f80dca0715b03d3296a783331210a41a8f4c89d2d2d0433f13.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckGrantedPermissionsForApp the checkGrantedPermissionsForApp property
func (m *GroupRequestBuilder) CheckGrantedPermissionsForApp()(*ief4021e7b13991e6921047b490cf96ebd1b3f9b8cb51f53f88e6d22b13625aca.CheckGrantedPermissionsForAppRequestBuilder) {
    return ief4021e7b13991e6921047b490cf96ebd1b3f9b8cb51f53f88e6d22b13625aca.NewCheckGrantedPermissionsForAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *GroupRequestBuilder) CheckMemberGroups()(*ib77b6ef94b8f11e8e3447f3eb6adca4cada604440e45269b0dbe11880a0dab37.CheckMemberGroupsRequestBuilder) {
    return ib77b6ef94b8f11e8e3447f3eb6adca4cada604440e45269b0dbe11880a0dab37.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *GroupRequestBuilder) CheckMemberObjects()(*i5ba239b7b916932a59648738c9c857f641c1a79a37ca8b919003cdf2b2b425a5.CheckMemberObjectsRequestBuilder) {
    return i5ba239b7b916932a59648738c9c857f641c1a79a37ca8b919003cdf2b2b425a5.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupRequestBuilderInternal instantiates a new GroupRequestBuilder and sets the default values.
func NewGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupRequestBuilder) {
    m := &GroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.group{?%24select,%24expand}";
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
func (m *GroupRequestBuilder) EvaluateDynamicMembership()(*id71215f18227366f93e6a7751d0838fa8e720356a7b68ce2350d10c20c658a8b.EvaluateDynamicMembershipRequestBuilder) {
    return id71215f18227366f93e6a7751d0838fa8e720356a7b68ce2350d10c20c658a8b.NewEvaluateDynamicMembershipRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.group
func (m *GroupRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMemberGroups the getMemberGroups property
func (m *GroupRequestBuilder) GetMemberGroups()(*i17092f7261264a2f0255a5ac04f42580fce1e959fea0a681c4a91bf0638fb515.GetMemberGroupsRequestBuilder) {
    return i17092f7261264a2f0255a5ac04f42580fce1e959fea0a681c4a91bf0638fb515.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *GroupRequestBuilder) GetMemberObjects()(*i3efec48d3704d9e37ec0caf7bf7392c1cc9544203611a18ac0474ebe63248ac8.GetMemberObjectsRequestBuilder) {
    return i3efec48d3704d9e37ec0caf7bf7392c1cc9544203611a18ac0474ebe63248ac8.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *GroupRequestBuilder) RemoveFavorite()(*i7ba94979f26931d406d3be0e7542846413bff6fa16c87f2d4341f65531091785.RemoveFavoriteRequestBuilder) {
    return i7ba94979f26931d406d3be0e7542846413bff6fa16c87f2d4341f65531091785.NewRemoveFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Renew the renew property
func (m *GroupRequestBuilder) Renew()(*i99671a59b4f1302345143fa82c5025ff0fa9b31e6a5d9ea0a738887f79d0f5cb.RenewRequestBuilder) {
    return i99671a59b4f1302345143fa82c5025ff0fa9b31e6a5d9ea0a738887f79d0f5cb.NewRenewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetUnseenCount the resetUnseenCount property
func (m *GroupRequestBuilder) ResetUnseenCount()(*i876930b27688b7989eecbe8cc0bb8d08ce2346b975da5a1bb6f724bfe0d55a0d.ResetUnseenCountRequestBuilder) {
    return i876930b27688b7989eecbe8cc0bb8d08ce2346b975da5a1bb6f724bfe0d55a0d.NewResetUnseenCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *GroupRequestBuilder) Restore()(*i03ecb7ca34f92b5605d7245fc90dc2c0b72dc43f58a22a6c7e21f723dad5b2d8.RestoreRequestBuilder) {
    return i03ecb7ca34f92b5605d7245fc90dc2c0b72dc43f58a22a6c7e21f723dad5b2d8.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscribeByMail the subscribeByMail property
func (m *GroupRequestBuilder) SubscribeByMail()(*i2a2bd0ca2efee011e8eaaea4a0bb2124144ad7bbb11247260d59f91238a429c1.SubscribeByMailRequestBuilder) {
    return i2a2bd0ca2efee011e8eaaea4a0bb2124144ad7bbb11247260d59f91238a429c1.NewSubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnsubscribeByMail the unsubscribeByMail property
func (m *GroupRequestBuilder) UnsubscribeByMail()(*i1c539874b822563762684a2fbf019b0fd92c430ebadb0ff0d8b8345f60438d48.UnsubscribeByMailRequestBuilder) {
    return i1c539874b822563762684a2fbf019b0fd92c430ebadb0ff0d8b8345f60438d48.NewUnsubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties the validateProperties property
func (m *GroupRequestBuilder) ValidateProperties()(*i6f733851fd18148e44adbae5ec3e14e4daf59da5accd703f72885d24108bbaf7.ValidatePropertiesRequestBuilder) {
    return i6f733851fd18148e44adbae5ec3e14e4daf59da5accd703f72885d24108bbaf7.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
