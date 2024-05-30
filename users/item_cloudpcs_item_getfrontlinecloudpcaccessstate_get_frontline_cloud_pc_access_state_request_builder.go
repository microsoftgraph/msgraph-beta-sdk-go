package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilder provides operations to call the getFrontlineCloudPcAccessState method.
type ItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilderInternal instantiates a new ItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilder and sets the default values.
func NewItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilder) {
    m := &ItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}/getFrontlineCloudPcAccessState()", pathParameters),
    }
    return m
}
// NewItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilder instantiates a new ItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilder and sets the default values.
func NewItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the frontlineCloudPcAccessState of a frontline Cloud PC.  This API only supports shared-use licenses. For more information, see cloudPcProvisioningPolicy. Shared-use licenses allow three users per license, with one user signed in at a time. Callers can get the latest frontline Cloud PC accessState and determine whether the frontline Cloud PC is accessible to a user.  If a web client needs to connect to a frontline Cloud PC, the sharedCloudPcAccessState validates the bookmark scenario. If sharedCloudPcAccessState isn't active/activating/standbyMode, the web client shows a bad bookmark.
// returns a *FrontlineCloudPcAccessState when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-getfrontlinecloudpcaccessstate?view=graph-rest-beta
func (m *ItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilderGetRequestConfiguration)(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FrontlineCloudPcAccessState, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendEnum(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseFrontlineCloudPcAccessState, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.FrontlineCloudPcAccessState), nil
}
// ToGetRequestInformation get the frontlineCloudPcAccessState of a frontline Cloud PC.  This API only supports shared-use licenses. For more information, see cloudPcProvisioningPolicy. Shared-use licenses allow three users per license, with one user signed in at a time. Callers can get the latest frontline Cloud PC accessState and determine whether the frontline Cloud PC is accessible to a user.  If a web client needs to connect to a frontline Cloud PC, the sharedCloudPcAccessState validates the bookmark scenario. If sharedCloudPcAccessState isn't active/activating/standbyMode, the web client shows a bad bookmark.
// returns a *RequestInformation when successful
func (m *ItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilder when successful
func (m *ItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilder) WithUrl(rawUrl string)(*ItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilder) {
    return NewItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
