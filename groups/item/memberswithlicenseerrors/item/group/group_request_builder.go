package group

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i13100d95d92a1ece3101c01fdf32f7dc63523e383aa97dc42ce5f56ff1151443 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/item/group/getmemberobjects"
    i18921f4df2edd6b129f7f35acb180851a4ff22cef45509ab75ea67520f7b4e49 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/item/group/removefavorite"
    i1b50bc5fac9fc4a42f5f6698ab93b1874df80067a469bf5a1248c6214a1b2ea7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/item/group/checkgrantedpermissionsforapp"
    i2820cfde39f2f605cedf77c1a1eb4d05ffc500dd15483b9a32d1a5534c74f197 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/item/group/addfavorite"
    i2bb9e81c9bd2bd802a5f2dac96b1aabc1d8c1d28fbc1af1c529142c735ce5940 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/item/group/assignlicense"
    i2c0fc72004f231d479b96e19c684cd25386cafdda521ed56fb978b8884137221 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/item/group/restore"
    i370173ffe84269acb9c8063b6dc1392c076bda0c5b3f7bfe8ae8a4541ab269c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/item/group/getmembergroups"
    i4c476796bf3731ae4ff71c6dfc9be00c16d2e289cba3f83597f144f193a71d73 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/item/group/validateproperties"
    i69c0b00afc77c603ba47ac4bea84e68b57dc996ba61e36fde64c6f0d7c1a6d77 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/item/group/renew"
    i6f1ac0a399aa68e54e89834e24e81061487312d879d2b0daace7f4d08bfc0f83 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/item/group/unsubscribebymail"
    i73a24de5c925e7bf6d91209385e95099d69ea92872709503af4be32edbddc1ec "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/item/group/subscribebymail"
    i84622b6a316432bea304345a5322c4d9e521b5105a52021d7910ae4d86ae8330 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/item/group/checkmembergroups"
    i9a70ed774ca6d298bd0d35c4f9fdfb15fe6cd37007f7646ed122c33e34f274c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/item/group/checkmemberobjects"
    ib69f3638aedbab79a7169587713b24daf56a631f883992337feadc799ae01983 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/item/group/evaluatedynamicmembership"
    ifed102d6ee7cd513dfbe6fa0a2134a6e8b4a14c0359c53425047f11f51027e30 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/memberswithlicenseerrors/item/group/resetunseencount"
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
func (m *GroupRequestBuilder) AddFavorite()(*i2820cfde39f2f605cedf77c1a1eb4d05ffc500dd15483b9a32d1a5534c74f197.AddFavoriteRequestBuilder) {
    return i2820cfde39f2f605cedf77c1a1eb4d05ffc500dd15483b9a32d1a5534c74f197.NewAddFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *GroupRequestBuilder) AssignLicense()(*i2bb9e81c9bd2bd802a5f2dac96b1aabc1d8c1d28fbc1af1c529142c735ce5940.AssignLicenseRequestBuilder) {
    return i2bb9e81c9bd2bd802a5f2dac96b1aabc1d8c1d28fbc1af1c529142c735ce5940.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckGrantedPermissionsForApp the checkGrantedPermissionsForApp property
func (m *GroupRequestBuilder) CheckGrantedPermissionsForApp()(*i1b50bc5fac9fc4a42f5f6698ab93b1874df80067a469bf5a1248c6214a1b2ea7.CheckGrantedPermissionsForAppRequestBuilder) {
    return i1b50bc5fac9fc4a42f5f6698ab93b1874df80067a469bf5a1248c6214a1b2ea7.NewCheckGrantedPermissionsForAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *GroupRequestBuilder) CheckMemberGroups()(*i84622b6a316432bea304345a5322c4d9e521b5105a52021d7910ae4d86ae8330.CheckMemberGroupsRequestBuilder) {
    return i84622b6a316432bea304345a5322c4d9e521b5105a52021d7910ae4d86ae8330.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *GroupRequestBuilder) CheckMemberObjects()(*i9a70ed774ca6d298bd0d35c4f9fdfb15fe6cd37007f7646ed122c33e34f274c0.CheckMemberObjectsRequestBuilder) {
    return i9a70ed774ca6d298bd0d35c4f9fdfb15fe6cd37007f7646ed122c33e34f274c0.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupRequestBuilderInternal instantiates a new GroupRequestBuilder and sets the default values.
func NewGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupRequestBuilder) {
    m := &GroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/membersWithLicenseErrors/{directoryObject%2Did}/microsoft.graph.group{?%24select,%24expand}";
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
func (m *GroupRequestBuilder) EvaluateDynamicMembership()(*ib69f3638aedbab79a7169587713b24daf56a631f883992337feadc799ae01983.EvaluateDynamicMembershipRequestBuilder) {
    return ib69f3638aedbab79a7169587713b24daf56a631f883992337feadc799ae01983.NewEvaluateDynamicMembershipRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.group
func (m *GroupRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMemberGroups the getMemberGroups property
func (m *GroupRequestBuilder) GetMemberGroups()(*i370173ffe84269acb9c8063b6dc1392c076bda0c5b3f7bfe8ae8a4541ab269c6.GetMemberGroupsRequestBuilder) {
    return i370173ffe84269acb9c8063b6dc1392c076bda0c5b3f7bfe8ae8a4541ab269c6.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *GroupRequestBuilder) GetMemberObjects()(*i13100d95d92a1ece3101c01fdf32f7dc63523e383aa97dc42ce5f56ff1151443.GetMemberObjectsRequestBuilder) {
    return i13100d95d92a1ece3101c01fdf32f7dc63523e383aa97dc42ce5f56ff1151443.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *GroupRequestBuilder) RemoveFavorite()(*i18921f4df2edd6b129f7f35acb180851a4ff22cef45509ab75ea67520f7b4e49.RemoveFavoriteRequestBuilder) {
    return i18921f4df2edd6b129f7f35acb180851a4ff22cef45509ab75ea67520f7b4e49.NewRemoveFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Renew the renew property
func (m *GroupRequestBuilder) Renew()(*i69c0b00afc77c603ba47ac4bea84e68b57dc996ba61e36fde64c6f0d7c1a6d77.RenewRequestBuilder) {
    return i69c0b00afc77c603ba47ac4bea84e68b57dc996ba61e36fde64c6f0d7c1a6d77.NewRenewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetUnseenCount the resetUnseenCount property
func (m *GroupRequestBuilder) ResetUnseenCount()(*ifed102d6ee7cd513dfbe6fa0a2134a6e8b4a14c0359c53425047f11f51027e30.ResetUnseenCountRequestBuilder) {
    return ifed102d6ee7cd513dfbe6fa0a2134a6e8b4a14c0359c53425047f11f51027e30.NewResetUnseenCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *GroupRequestBuilder) Restore()(*i2c0fc72004f231d479b96e19c684cd25386cafdda521ed56fb978b8884137221.RestoreRequestBuilder) {
    return i2c0fc72004f231d479b96e19c684cd25386cafdda521ed56fb978b8884137221.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscribeByMail the subscribeByMail property
func (m *GroupRequestBuilder) SubscribeByMail()(*i73a24de5c925e7bf6d91209385e95099d69ea92872709503af4be32edbddc1ec.SubscribeByMailRequestBuilder) {
    return i73a24de5c925e7bf6d91209385e95099d69ea92872709503af4be32edbddc1ec.NewSubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnsubscribeByMail the unsubscribeByMail property
func (m *GroupRequestBuilder) UnsubscribeByMail()(*i6f1ac0a399aa68e54e89834e24e81061487312d879d2b0daace7f4d08bfc0f83.UnsubscribeByMailRequestBuilder) {
    return i6f1ac0a399aa68e54e89834e24e81061487312d879d2b0daace7f4d08bfc0f83.NewUnsubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties the validateProperties property
func (m *GroupRequestBuilder) ValidateProperties()(*i4c476796bf3731ae4ff71c6dfc9be00c16d2e289cba3f83597f144f193a71d73.ValidatePropertiesRequestBuilder) {
    return i4c476796bf3731ae4ff71c6dfc9be00c16d2e289cba3f83597f144f193a71d73.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
