package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilder provides operations to manage the incompatibleGroups property of the microsoft.graph.accessPackage entity.
type EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilderGetQueryParameters the groups whose members are ineligible to be assigned this access package.
type EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilderGetQueryParameters struct {
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
// EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilderGetQueryParameters
}
// ByGroupId gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageCatalogs.item.accessPackages.item.incompatibleGroups.item collection
func (m *EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilder) ByGroupId(groupId string)(*EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if groupId != "" {
        urlTplParams["group%2Did"] = groupId
    }
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilderInternal instantiates a new IncompatibleGroupsRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilder) {
    m := &EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageCatalogs/{accessPackageCatalog%2Did}/accessPackages/{accessPackage%2Did}/incompatibleGroups{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilder instantiates a new IncompatibleGroupsRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilder) Count()(*EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsCountRequestBuilder) {
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the groups whose members are ineligible to be assigned this access package.
func (m *EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// Ref provides operations to manage the collection of identityGovernance entities.
func (m *EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilder) Ref()(*EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRefRequestBuilder) {
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the groups whose members are ineligible to be assigned this access package.
func (m *EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemIncompatibleGroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
