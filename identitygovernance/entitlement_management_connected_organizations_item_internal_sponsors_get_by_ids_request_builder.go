package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsRequestBuilder provides operations to call the getByIds method.
type EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsRequestBuilderInternal instantiates a new GetByIdsRequestBuilder and sets the default values.
func NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsRequestBuilder) {
    m := &EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/connectedOrganizations/{connectedOrganization%2Did}/internalSponsors/getByIds", pathParameters),
    }
    return m
}
// NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsRequestBuilder instantiates a new GetByIdsRequestBuilder and sets the default values.
func NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post return the directory objects specified in a list of IDs. Some common uses for this function are to:
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directoryobject-getbyids?view=graph-rest-1.0
func (m *EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsRequestBuilder) Post(ctx context.Context, body EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostRequestBodyable, requestConfiguration *EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsRequestBuilderPostRequestConfiguration)(EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsResponseable), nil
}
// ToPostRequestInformation return the directory objects specified in a list of IDs. Some common uses for this function are to:
func (m *EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsRequestBuilder) ToPostRequestInformation(ctx context.Context, body EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostRequestBodyable, requestConfiguration *EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsRequestBuilder) {
    return NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
