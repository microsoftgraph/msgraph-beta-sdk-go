package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder provides operations to call the getProvisionedCloudPCs method.
type ItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetQueryParameters invoke function getProvisionedCloudPCs
type ItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
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
// ItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetQueryParameters
}
// NewItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderInternal instantiates a new GetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder and sets the default values.
func NewItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, groupId *string, servicePlanId *string)(*ItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder) {
    m := &ItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/cloudPCs/getProvisionedCloudPCs(groupId='{groupId}',servicePlanId='{servicePlanId}'){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}", pathParameters),
    }
    if groupId != nil {
        m.BaseRequestBuilder.PathParameters["groupId"] = *groupId
    }
    if servicePlanId != nil {
        m.BaseRequestBuilder.PathParameters["servicePlanId"] = *servicePlanId
    }
    return m
}
// NewItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder instantiates a new GetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder and sets the default values.
func NewItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function getProvisionedCloudPCs
// Deprecated: This method is obsolete. Use GetAsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdGetResponse instead.
func (m *ItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetRequestConfiguration)(ItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdResponseable), nil
}
// GetAsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdGetResponse invoke function getProvisionedCloudPCs
func (m *ItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder) GetAsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdGetResponse(ctx context.Context, requestConfiguration *ItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetRequestConfiguration)(ItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdGetResponseable), nil
}
// ToGetRequestInformation invoke function getProvisionedCloudPCs
func (m *ItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder) WithUrl(rawUrl string)(*ItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder) {
    return NewItemCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
