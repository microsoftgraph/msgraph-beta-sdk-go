package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesDeploymentAudiencesItemMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder provides operations to call the updateAudienceById method.
type WindowsUpdatesDeploymentAudiencesItemMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesDeploymentAudiencesItemMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentAudiencesItemMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesDeploymentAudiencesItemMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilderInternal instantiates a new MicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentAudiencesItemMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentAudiencesItemMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder) {
    m := &WindowsUpdatesDeploymentAudiencesItemMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/deploymentAudiences/{deploymentAudience%2Did}/microsoft.graph.windowsUpdates.updateAudienceById", pathParameters),
    }
    return m
}
// NewWindowsUpdatesDeploymentAudiencesItemMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder instantiates a new MicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentAudiencesItemMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentAudiencesItemMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesDeploymentAudiencesItemMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update the members and exclusions collections of a deploymentAudience with updatableAsset resources of the same type. Adding an azureADDevice to the members or exclusions collections of a deployment audience automatically creates an Azure AD device object if it does not already exist. If the same updatableAsset gets included in the exclusions and members collections of a deploymentAudience, deployment will not apply to that asset. You can also use the method updateAudience to update the deploymentAudience.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/windowsupdates-deploymentaudience-updateaudiencebyid?view=graph-rest-1.0
func (m *WindowsUpdatesDeploymentAudiencesItemMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder) Post(ctx context.Context, body WindowsUpdatesDeploymentAudiencesItemMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentAudiencesItemMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation update the members and exclusions collections of a deploymentAudience with updatableAsset resources of the same type. Adding an azureADDevice to the members or exclusions collections of a deployment audience automatically creates an Azure AD device object if it does not already exist. If the same updatableAsset gets included in the exclusions and members collections of a deploymentAudience, deployment will not apply to that asset. You can also use the method updateAudience to update the deploymentAudience.
func (m *WindowsUpdatesDeploymentAudiencesItemMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesDeploymentAudiencesItemMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentAudiencesItemMicrosoftGraphWindowsUpdatesUpdateAudienceByIdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
