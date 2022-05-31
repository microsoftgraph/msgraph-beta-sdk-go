package group

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i00a9a0475d4d49e9bc305954ef8a50ab684cf641ab06430d76c8a03891e57254 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item/group/getmemberobjects"
    i3aaabb15e1bfa745f17aecd017f86807a9242138af5ed8715aeb50942d57d34f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item/group/checkmemberobjects"
    i913f1dabbdc21fed35d2eea36fd05e9d228112d660c02eed2e364e11aed32dee "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item/group/checkgrantedpermissionsforapp"
    i985b35f376dd52d0b316c74f52583d282b0035393711a1054b1b12e73358c178 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item/group/resetunseencount"
    i9d6e3845b22df449d53cd964a37c8a5279e0d35a32c2b6457767056ad07cc356 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item/group/subscribebymail"
    iad505319754ea109eb4acbe9bab7b2b1e5da681dc782803aa37c7b4ad78aab01 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item/group/evaluatedynamicmembership"
    ib75b06c21ea01cdf6e2d9b3e34572d5c96d1b8f9e4c8d4fb13c2a740f4d6a5ad "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item/group/addfavorite"
    ib93a4e6d1002003529f06cf97f862447fb7471f3475262b5d9604641eba83b25 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item/group/unsubscribebymail"
    ic0994e0a82dec65b9e8249eea2ec47a023294fe519c202f62f9ea34ed8394627 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item/group/renew"
    idc1dfec26eebff80a26edd5a391cc3e2f1a91bdb3313638361019b03147b7413 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item/group/removefavorite"
    ie89b8d794cebf21ee9b5d41e4c3a69dd1a1cd35934d87039c1423bf6687b1448 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item/group/validateproperties"
    ieeb071baee60c381c46ae1a524c8a219e504dfbb7f67a161b717708bb1dba8f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item/group/getmembergroups"
    ief1abfcddd67bf85dc640ffd970d03f195d9a5a919bc05e93be67fd9893e8971 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item/group/restore"
    if81f97fbde5d2efa861495598691f7a2da80d0e565649467b8b45beb220341cc "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item/group/checkmembergroups"
    if984b5b927434384759163b7acae04740f86067e2d79f42c4a7e651f744355f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item/group/assignlicense"
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
func (m *GroupRequestBuilder) AddFavorite()(*ib75b06c21ea01cdf6e2d9b3e34572d5c96d1b8f9e4c8d4fb13c2a740f4d6a5ad.AddFavoriteRequestBuilder) {
    return ib75b06c21ea01cdf6e2d9b3e34572d5c96d1b8f9e4c8d4fb13c2a740f4d6a5ad.NewAddFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *GroupRequestBuilder) AssignLicense()(*if984b5b927434384759163b7acae04740f86067e2d79f42c4a7e651f744355f2.AssignLicenseRequestBuilder) {
    return if984b5b927434384759163b7acae04740f86067e2d79f42c4a7e651f744355f2.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckGrantedPermissionsForApp the checkGrantedPermissionsForApp property
func (m *GroupRequestBuilder) CheckGrantedPermissionsForApp()(*i913f1dabbdc21fed35d2eea36fd05e9d228112d660c02eed2e364e11aed32dee.CheckGrantedPermissionsForAppRequestBuilder) {
    return i913f1dabbdc21fed35d2eea36fd05e9d228112d660c02eed2e364e11aed32dee.NewCheckGrantedPermissionsForAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *GroupRequestBuilder) CheckMemberGroups()(*if81f97fbde5d2efa861495598691f7a2da80d0e565649467b8b45beb220341cc.CheckMemberGroupsRequestBuilder) {
    return if81f97fbde5d2efa861495598691f7a2da80d0e565649467b8b45beb220341cc.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *GroupRequestBuilder) CheckMemberObjects()(*i3aaabb15e1bfa745f17aecd017f86807a9242138af5ed8715aeb50942d57d34f.CheckMemberObjectsRequestBuilder) {
    return i3aaabb15e1bfa745f17aecd017f86807a9242138af5ed8715aeb50942d57d34f.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupRequestBuilderInternal instantiates a new GroupRequestBuilder and sets the default values.
func NewGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupRequestBuilder) {
    m := &GroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/owners/{directoryObject%2Did}/microsoft.graph.group{?%24select,%24expand}";
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
func (m *GroupRequestBuilder) EvaluateDynamicMembership()(*iad505319754ea109eb4acbe9bab7b2b1e5da681dc782803aa37c7b4ad78aab01.EvaluateDynamicMembershipRequestBuilder) {
    return iad505319754ea109eb4acbe9bab7b2b1e5da681dc782803aa37c7b4ad78aab01.NewEvaluateDynamicMembershipRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.group
func (m *GroupRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMemberGroups the getMemberGroups property
func (m *GroupRequestBuilder) GetMemberGroups()(*ieeb071baee60c381c46ae1a524c8a219e504dfbb7f67a161b717708bb1dba8f1.GetMemberGroupsRequestBuilder) {
    return ieeb071baee60c381c46ae1a524c8a219e504dfbb7f67a161b717708bb1dba8f1.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *GroupRequestBuilder) GetMemberObjects()(*i00a9a0475d4d49e9bc305954ef8a50ab684cf641ab06430d76c8a03891e57254.GetMemberObjectsRequestBuilder) {
    return i00a9a0475d4d49e9bc305954ef8a50ab684cf641ab06430d76c8a03891e57254.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *GroupRequestBuilder) RemoveFavorite()(*idc1dfec26eebff80a26edd5a391cc3e2f1a91bdb3313638361019b03147b7413.RemoveFavoriteRequestBuilder) {
    return idc1dfec26eebff80a26edd5a391cc3e2f1a91bdb3313638361019b03147b7413.NewRemoveFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Renew the renew property
func (m *GroupRequestBuilder) Renew()(*ic0994e0a82dec65b9e8249eea2ec47a023294fe519c202f62f9ea34ed8394627.RenewRequestBuilder) {
    return ic0994e0a82dec65b9e8249eea2ec47a023294fe519c202f62f9ea34ed8394627.NewRenewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetUnseenCount the resetUnseenCount property
func (m *GroupRequestBuilder) ResetUnseenCount()(*i985b35f376dd52d0b316c74f52583d282b0035393711a1054b1b12e73358c178.ResetUnseenCountRequestBuilder) {
    return i985b35f376dd52d0b316c74f52583d282b0035393711a1054b1b12e73358c178.NewResetUnseenCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *GroupRequestBuilder) Restore()(*ief1abfcddd67bf85dc640ffd970d03f195d9a5a919bc05e93be67fd9893e8971.RestoreRequestBuilder) {
    return ief1abfcddd67bf85dc640ffd970d03f195d9a5a919bc05e93be67fd9893e8971.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscribeByMail the subscribeByMail property
func (m *GroupRequestBuilder) SubscribeByMail()(*i9d6e3845b22df449d53cd964a37c8a5279e0d35a32c2b6457767056ad07cc356.SubscribeByMailRequestBuilder) {
    return i9d6e3845b22df449d53cd964a37c8a5279e0d35a32c2b6457767056ad07cc356.NewSubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnsubscribeByMail the unsubscribeByMail property
func (m *GroupRequestBuilder) UnsubscribeByMail()(*ib93a4e6d1002003529f06cf97f862447fb7471f3475262b5d9604641eba83b25.UnsubscribeByMailRequestBuilder) {
    return ib93a4e6d1002003529f06cf97f862447fb7471f3475262b5d9604641eba83b25.NewUnsubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties the validateProperties property
func (m *GroupRequestBuilder) ValidateProperties()(*ie89b8d794cebf21ee9b5d41e4c3a69dd1a1cd35934d87039c1423bf6687b1448.ValidatePropertiesRequestBuilder) {
    return ie89b8d794cebf21ee9b5d41e4c3a69dd1a1cd35934d87039c1423bf6687b1448.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
