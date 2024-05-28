package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder provides operations to call the updateAudienceById method.
type WindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilderInternal instantiates a new WindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder) {
    m := &WindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/deployments/{deployment%2Did}/audience/microsoft.graph.windowsUpdates.updateAudienceById", pathParameters),
    }
    return m
}
// NewWindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder instantiates a new WindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update the members and exclusions collections of a deploymentAudience with updatableAsset resources of the same type. Adding an azureADDevice to the members or exclusions collections of a deployment audience automatically creates a Microsoft Entra device object if it does not already exist. If the same updatableAsset gets included in the exclusions and members collections of a deploymentAudience, deployment will not apply to that asset. You can also use the method updateAudience to update the deploymentAudience.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/windowsupdates-deploymentaudience-updateaudiencebyid?view=graph-rest-beta
func (m *WindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder) Post(ctx context.Context, body WindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidUpdateAudienceByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilderPostRequestConfiguration)(error) {
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
func (m *WindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidUpdateAudienceByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder when successful
func (m *WindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceMicrosoftgraphwindowsupdatesupdateaudiencebyidMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
