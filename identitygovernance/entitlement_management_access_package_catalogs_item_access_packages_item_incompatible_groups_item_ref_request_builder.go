package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsItemRefRequestBuilder provides operations to manage the collection of identityGovernance entities.
type EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsItemRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsItemRefRequestBuilderDeleteQueryParameters delete ref of navigation property incompatibleGroups for identityGovernance
type EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsItemRefRequestBuilderDeleteQueryParameters struct {
    // Delete Uri
    Id *string `uriparametername:"%40id"`
}
// EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsItemRefRequestBuilderDeleteQueryParameters
}
// NewEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsItemRefRequestBuilderInternal instantiates a new RefRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsItemRefRequestBuilder) {
    m := &EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsItemRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageCatalogs/{accessPackageCatalog%2Did}/accessPackages/{accessPackage%2Did}/incompatibleGroups/{group%2Did}/$ref{?%40id*}", pathParameters),
    }
    return m
}
// NewEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsItemRefRequestBuilder instantiates a new RefRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete ref of navigation property incompatibleGroups for identityGovernance
func (m *EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsItemRefRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation delete ref of navigation property incompatibleGroups for identityGovernance
func (m *EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
