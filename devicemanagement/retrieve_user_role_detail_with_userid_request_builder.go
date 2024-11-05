package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RetrieveUserRoleDetailWithUseridRequestBuilder provides operations to call the retrieveUserRoleDetail method.
type RetrieveUserRoleDetailWithUseridRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RetrieveUserRoleDetailWithUseridRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RetrieveUserRoleDetailWithUseridRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRetrieveUserRoleDetailWithUseridRequestBuilderInternal instantiates a new RetrieveUserRoleDetailWithUseridRequestBuilder and sets the default values.
func NewRetrieveUserRoleDetailWithUseridRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, userid *string)(*RetrieveUserRoleDetailWithUseridRequestBuilder) {
    m := &RetrieveUserRoleDetailWithUseridRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/retrieveUserRoleDetail(userid='{userid}')", pathParameters),
    }
    if userid != nil {
        m.BaseRequestBuilder.PathParameters["userid"] = *userid
    }
    return m
}
// NewRetrieveUserRoleDetailWithUseridRequestBuilder instantiates a new RetrieveUserRoleDetailWithUseridRequestBuilder and sets the default values.
func NewRetrieveUserRoleDetailWithUseridRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RetrieveUserRoleDetailWithUseridRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRetrieveUserRoleDetailWithUseridRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function retrieveUserRoleDetail
// returns a DeviceAndAppManagementAssignedRoleDetailable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RetrieveUserRoleDetailWithUseridRequestBuilder) Get(ctx context.Context, requestConfiguration *RetrieveUserRoleDetailWithUseridRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementAssignedRoleDetailable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceAndAppManagementAssignedRoleDetailFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementAssignedRoleDetailable), nil
}
// ToGetRequestInformation invoke function retrieveUserRoleDetail
// returns a *RequestInformation when successful
func (m *RetrieveUserRoleDetailWithUseridRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RetrieveUserRoleDetailWithUseridRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RetrieveUserRoleDetailWithUseridRequestBuilder when successful
func (m *RetrieveUserRoleDetailWithUseridRequestBuilder) WithUrl(rawUrl string)(*RetrieveUserRoleDetailWithUseridRequestBuilder) {
    return NewRetrieveUserRoleDetailWithUseridRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
