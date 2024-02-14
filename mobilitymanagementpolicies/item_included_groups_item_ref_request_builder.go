package mobilitymanagementpolicies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemIncludedGroupsItemRefRequestBuilder provides operations to manage the collection of mobilityManagementPolicy entities.
type ItemIncludedGroupsItemRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemIncludedGroupsItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemIncludedGroupsItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemIncludedGroupsItemRefRequestBuilderInternal instantiates a new ItemIncludedGroupsItemRefRequestBuilder and sets the default values.
func NewItemIncludedGroupsItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemIncludedGroupsItemRefRequestBuilder) {
    m := &ItemIncludedGroupsItemRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/mobilityManagementPolicies/{mobilityManagementPolicy%2Did}/includedGroups/{group%2Did}/$ref", pathParameters),
    }
    return m
}
// NewItemIncludedGroupsItemRefRequestBuilder instantiates a new ItemIncludedGroupsItemRefRequestBuilder and sets the default values.
func NewItemIncludedGroupsItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemIncludedGroupsItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemIncludedGroupsItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a group from the list of groups included in a mobile app management policy.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/mobileappmanagementpolicies-delete-includedgroups?view=graph-rest-1.0
func (m *ItemIncludedGroupsItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemIncludedGroupsItemRefRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation delete a group from the list of groups included in a mobile app management policy.
// returns a *RequestInformation when successful
func (m *ItemIncludedGroupsItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemIncludedGroupsItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemIncludedGroupsItemRefRequestBuilder when successful
func (m *ItemIncludedGroupsItemRefRequestBuilder) WithUrl(rawUrl string)(*ItemIncludedGroupsItemRefRequestBuilder) {
    return NewItemIncludedGroupsItemRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
