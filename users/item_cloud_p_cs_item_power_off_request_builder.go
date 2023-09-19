package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudPCsItemPowerOffRequestBuilder provides operations to call the powerOff method.
type ItemCloudPCsItemPowerOffRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCloudPCsItemPowerOffRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudPCsItemPowerOffRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCloudPCsItemPowerOffRequestBuilderInternal instantiates a new PowerOffRequestBuilder and sets the default values.
func NewItemCloudPCsItemPowerOffRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsItemPowerOffRequestBuilder) {
    m := &ItemCloudPCsItemPowerOffRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}/powerOff", pathParameters),
    }
    return m
}
// NewItemCloudPCsItemPowerOffRequestBuilder instantiates a new PowerOffRequestBuilder and sets the default values.
func NewItemCloudPCsItemPowerOffRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsItemPowerOffRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudPCsItemPowerOffRequestBuilderInternal(urlParams, requestAdapter)
}
// Post power off a Windows 365 Frontline Cloud PC. This action supports Microsoft Endpoint Manager (MEM) admin scenarios.  After a Windows 365 Frontline Cloud PC is powered off, it's deallocated, and licenses are revoked immediately. Only IT admin users can perform this action. 
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-poweroff?view=graph-rest-1.0
func (m *ItemCloudPCsItemPowerOffRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCloudPCsItemPowerOffRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation power off a Windows 365 Frontline Cloud PC. This action supports Microsoft Endpoint Manager (MEM) admin scenarios.  After a Windows 365 Frontline Cloud PC is powered off, it's deallocated, and licenses are revoked immediately. Only IT admin users can perform this action. 
func (m *ItemCloudPCsItemPowerOffRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCloudPCsItemPowerOffRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemCloudPCsItemPowerOffRequestBuilder) WithUrl(rawUrl string)(*ItemCloudPCsItemPowerOffRequestBuilder) {
    return NewItemCloudPCsItemPowerOffRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
