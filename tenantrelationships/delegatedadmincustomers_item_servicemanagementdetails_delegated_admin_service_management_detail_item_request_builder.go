package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilder provides operations to manage the serviceManagementDetails property of the microsoft.graph.delegatedAdminCustomer entity.
type DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilderGetQueryParameters contains the management details of a service in the customer tenant that's managed by delegated administration.
type DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilderGetQueryParameters
}
// DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilderInternal instantiates a new DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilder and sets the default values.
func NewDelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilder) {
    m := &DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/delegatedAdminCustomers/{delegatedAdminCustomer%2Did}/serviceManagementDetails/{delegatedAdminServiceManagementDetail%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilder instantiates a new DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilder and sets the default values.
func NewDelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property serviceManagementDetails for tenantRelationships
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get contains the management details of a service in the customer tenant that's managed by delegated administration.
// returns a DelegatedAdminServiceManagementDetailable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminServiceManagementDetailable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDelegatedAdminServiceManagementDetailFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminServiceManagementDetailable), nil
}
// Patch update the navigation property serviceManagementDetails in tenantRelationships
// returns a DelegatedAdminServiceManagementDetailable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminServiceManagementDetailable, requestConfiguration *DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminServiceManagementDetailable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDelegatedAdminServiceManagementDetailFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminServiceManagementDetailable), nil
}
// ToDeleteRequestInformation delete navigation property serviceManagementDetails for tenantRelationships
// returns a *RequestInformation when successful
func (m *DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation contains the management details of a service in the customer tenant that's managed by delegated administration.
// returns a *RequestInformation when successful
func (m *DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property serviceManagementDetails in tenantRelationships
// returns a *RequestInformation when successful
func (m *DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminServiceManagementDetailable, requestConfiguration *DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilder when successful
func (m *DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilder) WithUrl(rawUrl string)(*DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilder) {
    return NewDelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
