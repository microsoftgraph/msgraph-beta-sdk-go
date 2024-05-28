package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetassignedroledetailsGetAssignedRoleDetailsRequestBuilder provides operations to call the getAssignedRoleDetails method.
type GetassignedroledetailsGetAssignedRoleDetailsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetassignedroledetailsGetAssignedRoleDetailsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetassignedroledetailsGetAssignedRoleDetailsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetassignedroledetailsGetAssignedRoleDetailsRequestBuilderInternal instantiates a new GetassignedroledetailsGetAssignedRoleDetailsRequestBuilder and sets the default values.
func NewGetassignedroledetailsGetAssignedRoleDetailsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetassignedroledetailsGetAssignedRoleDetailsRequestBuilder) {
    m := &GetassignedroledetailsGetAssignedRoleDetailsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/getAssignedRoleDetails()", pathParameters),
    }
    return m
}
// NewGetassignedroledetailsGetAssignedRoleDetailsRequestBuilder instantiates a new GetassignedroledetailsGetAssignedRoleDetailsRequestBuilder and sets the default values.
func NewGetassignedroledetailsGetAssignedRoleDetailsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetassignedroledetailsGetAssignedRoleDetailsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetassignedroledetailsGetAssignedRoleDetailsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieves the assigned role definitions and role assignments of the currently authenticated user.
// returns a DeviceAndAppManagementAssignedRoleDetailsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetassignedroledetailsGetAssignedRoleDetailsRequestBuilder) Get(ctx context.Context, requestConfiguration *GetassignedroledetailsGetAssignedRoleDetailsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementAssignedRoleDetailsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceAndAppManagementAssignedRoleDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementAssignedRoleDetailsable), nil
}
// ToGetRequestInformation retrieves the assigned role definitions and role assignments of the currently authenticated user.
// returns a *RequestInformation when successful
func (m *GetassignedroledetailsGetAssignedRoleDetailsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetassignedroledetailsGetAssignedRoleDetailsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetassignedroledetailsGetAssignedRoleDetailsRequestBuilder when successful
func (m *GetassignedroledetailsGetAssignedRoleDetailsRequestBuilder) WithUrl(rawUrl string)(*GetassignedroledetailsGetAssignedRoleDetailsRequestBuilder) {
    return NewGetassignedroledetailsGetAssignedRoleDetailsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
