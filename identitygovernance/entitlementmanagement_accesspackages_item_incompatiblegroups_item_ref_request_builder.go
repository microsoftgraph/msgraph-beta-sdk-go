package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemRefRequestBuilder provides operations to manage the collection of identityGovernance entities.
type EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsItemRefRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemRefRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemRefRequestBuilder) {
    m := &EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/incompatibleGroups/{group%2Did}/$ref", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsItemRefRequestBuilder instantiates a new EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemRefRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove a group from the list of groups marked as incompatible on an accessPackage.  
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackage-delete-incompatiblegroup?view=graph-rest-beta
func (m *EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemRefRequestBuilderDeleteRequestConfiguration)(error) {
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
// ToDeleteRequestInformation remove a group from the list of groups marked as incompatible on an accessPackage.  
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemRefRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemRefRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagesItemIncompatiblegroupsItemRefRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsItemRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
