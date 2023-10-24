package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilder provides operations to manage the authenticationContext property of the microsoft.graph.unifiedRbacResourceAction entity.
type DeviceManagementResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceManagementResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilderGetQueryParameters get authenticationContext from roleManagement
type DeviceManagementResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilderGetQueryParameters
}
// NewDeviceManagementResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilderInternal instantiates a new AuthenticationContextRequestBuilder and sets the default values.
func NewDeviceManagementResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilder) {
    m := &DeviceManagementResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/deviceManagement/resourceNamespaces/{unifiedRbacResourceNamespace%2Did}/resourceActions/{unifiedRbacResourceAction%2Did}/authenticationContext{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewDeviceManagementResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilder instantiates a new AuthenticationContextRequestBuilder and sets the default values.
func NewDeviceManagementResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get authenticationContext from roleManagement
func (m *DeviceManagementResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationContextClassReferenceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationContextClassReferenceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationContextClassReferenceable), nil
}
// ToGetRequestInformation get authenticationContext from roleManagement
func (m *DeviceManagementResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DeviceManagementResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilder) WithUrl(rawUrl string)(*DeviceManagementResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilder) {
    return NewDeviceManagementResourceNamespacesItemResourceActionsItemAuthenticationContextRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
