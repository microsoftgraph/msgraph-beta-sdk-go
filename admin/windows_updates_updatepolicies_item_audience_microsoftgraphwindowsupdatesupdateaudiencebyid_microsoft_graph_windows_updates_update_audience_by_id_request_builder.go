package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder provides operations to call the updateAudienceById method.
type WindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilderInternal instantiates a new WindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder) {
    m := &WindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/updatePolicies/{updatePolicy%2Did}/audience/microsoft.graph.windowsUpdates.updateAudienceById", pathParameters),
    }
    return m
}
// NewWindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder instantiates a new WindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update the members and exclusions collections of a deploymentAudience with updatableAsset resources of the same type. Adding an azureADDevice to the members or exclusions collections of a deployment audience automatically creates a Microsoft Entra device object if it does not already exist. If the same updatableAsset gets included in the exclusions and members collections of a deploymentAudience, deployment will not apply to that asset. You can also use the method updateAudience to update the deploymentAudience.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/windowsupdates-deploymentaudience-updateaudiencebyid?view=graph-rest-beta
func (m *WindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder) Post(ctx context.Context, body WindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidUpdateAudienceByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation update the members and exclusions collections of a deploymentAudience with updatableAsset resources of the same type. Adding an azureADDevice to the members or exclusions collections of a deployment audience automatically creates a Microsoft Entra device object if it does not already exist. If the same updatableAsset gets included in the exclusions and members collections of a deploymentAudience, deployment will not apply to that asset. You can also use the method updateAudience to update the deploymentAudience.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidUpdateAudienceByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *WindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder when successful
func (m *WindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder) {
    return NewWindowsUpdatesUpdatepoliciesItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
