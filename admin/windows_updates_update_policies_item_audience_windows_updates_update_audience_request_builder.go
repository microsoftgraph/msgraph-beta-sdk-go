package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceRequestBuilder provides operations to call the updateAudience method.
type WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceRequestBuilderInternal instantiates a new WindowsUpdatesUpdateAudienceRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceRequestBuilder) {
    m := &WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/updatePolicies/{updatePolicy%2Did}/audience/windowsUpdates.updateAudience", pathParameters),
    }
    return m
}
// NewWindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceRequestBuilder instantiates a new WindowsUpdatesUpdateAudienceRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update the members and exclusions collections of a deploymentAudience. Adding an azureADDevice to the members or exclusions collections of a deployment audience automatically creates an Azure AD device object, if it does not already exist. If the same updatableAsset gets included in the **exclusions** and **members** collections of a **deploymentAudience**, deployment will not apply to that asset. If all **updatableAsset** objects are the same type, you can also use the method updateAudienceById to update the **deploymentAudience**.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/windowsupdates-deploymentaudience-updateaudience?view=graph-rest-1.0
func (m *WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceRequestBuilder) Post(ctx context.Context, body WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceUpdateAudiencePostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation update the members and exclusions collections of a deploymentAudience. Adding an azureADDevice to the members or exclusions collections of a deployment audience automatically creates an Azure AD device object, if it does not already exist. If the same updatableAsset gets included in the **exclusions** and **members** collections of a **deploymentAudience**, deployment will not apply to that asset. If all **updatableAsset** objects are the same type, you can also use the method updateAudienceById to update the **deploymentAudience**.
func (m *WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceUpdateAudiencePostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
