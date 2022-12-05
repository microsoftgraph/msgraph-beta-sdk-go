package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AdminWindowsUpdatesDeploymentsItemAudienceUpdateAudienceRequestBuilder provides operations to call the updateAudience method.
type AdminWindowsUpdatesDeploymentsItemAudienceUpdateAudienceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AdminWindowsUpdatesDeploymentsItemAudienceUpdateAudienceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminWindowsUpdatesDeploymentsItemAudienceUpdateAudienceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAdminWindowsUpdatesDeploymentsItemAudienceUpdateAudienceRequestBuilderInternal instantiates a new UpdateAudienceRequestBuilder and sets the default values.
func NewAdminWindowsUpdatesDeploymentsItemAudienceUpdateAudienceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminWindowsUpdatesDeploymentsItemAudienceUpdateAudienceRequestBuilder) {
    m := &AdminWindowsUpdatesDeploymentsItemAudienceUpdateAudienceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/deployments/{deployment%2Did}/audience/microsoft.graph.windowsUpdates.updateAudience";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAdminWindowsUpdatesDeploymentsItemAudienceUpdateAudienceRequestBuilder instantiates a new UpdateAudienceRequestBuilder and sets the default values.
func NewAdminWindowsUpdatesDeploymentsItemAudienceUpdateAudienceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminWindowsUpdatesDeploymentsItemAudienceUpdateAudienceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdminWindowsUpdatesDeploymentsItemAudienceUpdateAudienceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation update the members and exclusions collections of a deploymentAudience. Adding an azureADDevice to the members or exclusions collections of a deployment audience automatically creates an Azure AD device object, if it does not already exist. If the same updatableAsset gets included in the **exclusions** and **members** collections of a **deploymentAudience**, deployment will not apply to that asset. If all **updatableAsset** objects are the same type, you can also use the method updateAudienceById to update the **deploymentAudience**.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceUpdateAudienceRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdminWindowsUpdatesDeploymentsItemAudienceUpdateAudiencePostRequestBodyable, requestConfiguration *AdminWindowsUpdatesDeploymentsItemAudienceUpdateAudienceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post update the members and exclusions collections of a deploymentAudience. Adding an azureADDevice to the members or exclusions collections of a deployment audience automatically creates an Azure AD device object, if it does not already exist. If the same updatableAsset gets included in the **exclusions** and **members** collections of a **deploymentAudience**, deployment will not apply to that asset. If all **updatableAsset** objects are the same type, you can also use the method updateAudienceById to update the **deploymentAudience**.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceUpdateAudienceRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdminWindowsUpdatesDeploymentsItemAudienceUpdateAudiencePostRequestBodyable, requestConfiguration *AdminWindowsUpdatesDeploymentsItemAudienceUpdateAudienceRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
