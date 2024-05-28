package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder provides operations to call the additionalAccess method.
type EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetQueryParameters invoke function additionalAccess
type EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetQueryParameters struct {
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
// EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, accessPackageId *string, incompatibleAccessPackageId *string)(*EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) {
    m := &EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/additionalAccess(accessPackageId='{accessPackageId}',incompatibleAccessPackageId='{incompatibleAccessPackageId}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if accessPackageId != nil {
        m.BaseRequestBuilder.PathParameters["accessPackageId"] = *accessPackageId
    }
    if incompatibleAccessPackageId != nil {
        m.BaseRequestBuilder.PathParameters["incompatibleAccessPackageId"] = *incompatibleAccessPackageId
    }
    return m
}
// NewEntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder instantiates a new EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function additionalAccess
// Deprecated: This method is obsolete. Use GetAsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponse instead.
// returns a EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetRequestConfiguration)(EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponseable), nil
}
// GetAsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponse invoke function additionalAccess
// returns a EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) GetAsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponse(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetRequestConfiguration)(EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponseable), nil
}
// ToGetRequestInformation invoke function additionalAccess
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentsAdditionalaccesswithaccesspackageidwithincompatibleaccesspackageidAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
