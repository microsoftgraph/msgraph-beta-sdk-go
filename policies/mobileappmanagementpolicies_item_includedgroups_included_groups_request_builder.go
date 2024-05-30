package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilder provides operations to manage the includedGroups property of the microsoft.graph.mobilityManagementPolicy entity.
type MobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilderGetQueryParameters get the list of groups that are included in a mobile app management policy.
type MobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// MobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilderGetQueryParameters
}
// ByGroupId gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.mobileAppManagementPolicies.item.includedGroups.item collection
// returns a *MobileappmanagementpoliciesItemIncludedgroupsGroupItemRequestBuilder when successful
func (m *MobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilder) ByGroupId(groupId string)(*MobileappmanagementpoliciesItemIncludedgroupsGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if groupId != "" {
        urlTplParams["group%2Did"] = groupId
    }
    return NewMobileappmanagementpoliciesItemIncludedgroupsGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilderInternal instantiates a new MobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilder and sets the default values.
func NewMobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilder) {
    m := &MobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/mobileAppManagementPolicies/{mobilityManagementPolicy%2Did}/includedGroups{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewMobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilder instantiates a new MobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilder and sets the default values.
func NewMobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *MobileappmanagementpoliciesItemIncludedgroupsCountRequestBuilder when successful
func (m *MobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilder) Count()(*MobileappmanagementpoliciesItemIncludedgroupsCountRequestBuilder) {
    return NewMobileappmanagementpoliciesItemIncludedgroupsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the list of groups that are included in a mobile app management policy.
// returns a GroupCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/mobileappmanagementpolicies-list-includedgroups?view=graph-rest-beta
func (m *MobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupCollectionResponseable), nil
}
// Ref provides operations to manage the collection of policyRoot entities.
// returns a *MobileappmanagementpoliciesItemIncludedgroupsRefRequestBuilder when successful
func (m *MobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilder) Ref()(*MobileappmanagementpoliciesItemIncludedgroupsRefRequestBuilder) {
    return NewMobileappmanagementpoliciesItemIncludedgroupsRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get the list of groups that are included in a mobile app management policy.
// returns a *RequestInformation when successful
func (m *MobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilder when successful
func (m *MobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilder) WithUrl(rawUrl string)(*MobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilder) {
    return NewMobileappmanagementpoliciesItemIncludedgroupsIncludedGroupsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
