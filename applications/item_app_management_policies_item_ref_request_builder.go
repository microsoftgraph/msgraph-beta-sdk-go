package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAppManagementPoliciesItemRefRequestBuilder provides operations to manage the collection of application entities.
type ItemAppManagementPoliciesItemRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAppManagementPoliciesItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAppManagementPoliciesItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemAppManagementPoliciesItemRefRequestBuilderInternal instantiates a new ItemAppManagementPoliciesItemRefRequestBuilder and sets the default values.
func NewItemAppManagementPoliciesItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAppManagementPoliciesItemRefRequestBuilder) {
    m := &ItemAppManagementPoliciesItemRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/appManagementPolicies/{appManagementPolicy%2Did}/$ref", pathParameters),
    }
    return m
}
// NewItemAppManagementPoliciesItemRefRequestBuilder instantiates a new ItemAppManagementPoliciesItemRefRequestBuilder and sets the default values.
func NewItemAppManagementPoliciesItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAppManagementPoliciesItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAppManagementPoliciesItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove an appManagementPolicy policy object from an application or service principal object. When you remove the appManagementPolicy, the application or service principal adopts the tenant-wide tenantAppManagementPolicy setting. 
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/appmanagementpolicy-delete-appliesto?view=graph-rest-beta
func (m *ItemAppManagementPoliciesItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemAppManagementPoliciesItemRefRequestBuilderDeleteRequestConfiguration)(error) {
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
// ToDeleteRequestInformation remove an appManagementPolicy policy object from an application or service principal object. When you remove the appManagementPolicy, the application or service principal adopts the tenant-wide tenantAppManagementPolicy setting. 
// returns a *RequestInformation when successful
func (m *ItemAppManagementPoliciesItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemAppManagementPoliciesItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAppManagementPoliciesItemRefRequestBuilder when successful
func (m *ItemAppManagementPoliciesItemRefRequestBuilder) WithUrl(rawUrl string)(*ItemAppManagementPoliciesItemRefRequestBuilder) {
    return NewItemAppManagementPoliciesItemRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
