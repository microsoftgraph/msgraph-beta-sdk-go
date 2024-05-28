package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilder provides operations to call the additionalAccess method.
type EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilderGetQueryParameters in Microsoft Entra Entitlement Management, retrieve a collection of accessPackageAssignment objects that indicate a target user has an assignment to a specified access package and also an assignment to another, potentially incompatible, access package.  It can be used to prepare to configure the incompatible access packages for a specific access package.
type EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilderGetQueryParameters struct {
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
// EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilder) {
    m := &EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/additionalAccess(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilder instantiates a new EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilderInternal(urlParams, requestAdapter)
}
// Get in Microsoft Entra Entitlement Management, retrieve a collection of accessPackageAssignment objects that indicate a target user has an assignment to a specified access package and also an assignment to another, potentially incompatible, access package.  It can be used to prepare to configure the incompatible access packages for a specific access package.
// Deprecated: This method is obsolete. Use GetAsAdditionalAccessGetResponse instead.
// returns a EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackageassignment-additionalaccess?view=graph-rest-beta
func (m *EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilderGetRequestConfiguration)(EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessResponseable), nil
}
// GetAsAdditionalAccessGetResponse in Microsoft Entra Entitlement Management, retrieve a collection of accessPackageAssignment objects that indicate a target user has an assignment to a specified access package and also an assignment to another, potentially incompatible, access package.  It can be used to prepare to configure the incompatible access packages for a specific access package.
// returns a EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackageassignment-additionalaccess?view=graph-rest-beta
func (m *EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilder) GetAsAdditionalAccessGetResponse(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilderGetRequestConfiguration)(EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessGetResponseable), nil
}
// ToGetRequestInformation in Microsoft Entra Entitlement Management, retrieve a collection of accessPackageAssignment objects that indicate a target user has an assignment to a specified access package and also an assignment to another, potentially incompatible, access package.  It can be used to prepare to configure the incompatible access packages for a specific access package.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilder when successful
func (m *EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentsAdditionalaccessAdditionalAccessRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
