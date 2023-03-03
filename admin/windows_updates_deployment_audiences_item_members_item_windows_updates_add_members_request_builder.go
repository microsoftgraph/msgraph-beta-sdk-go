package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesDeploymentAudiencesItemMembersItemWindowsUpdatesAddMembersRequestBuilder provides operations to call the addMembers method.
type WindowsUpdatesDeploymentAudiencesItemMembersItemWindowsUpdatesAddMembersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// WindowsUpdatesDeploymentAudiencesItemMembersItemWindowsUpdatesAddMembersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentAudiencesItemMembersItemWindowsUpdatesAddMembersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesDeploymentAudiencesItemMembersItemWindowsUpdatesAddMembersRequestBuilderInternal instantiates a new WindowsUpdatesAddMembersRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentAudiencesItemMembersItemWindowsUpdatesAddMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentAudiencesItemMembersItemWindowsUpdatesAddMembersRequestBuilder) {
    m := &WindowsUpdatesDeploymentAudiencesItemMembersItemWindowsUpdatesAddMembersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/deploymentAudiences/{deploymentAudience%2Did}/members/{updatableAsset%2Did}/windowsUpdates.addMembers";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewWindowsUpdatesDeploymentAudiencesItemMembersItemWindowsUpdatesAddMembersRequestBuilder instantiates a new WindowsUpdatesAddMembersRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentAudiencesItemMembersItemWindowsUpdatesAddMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentAudiencesItemMembersItemWindowsUpdatesAddMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesDeploymentAudiencesItemMembersItemWindowsUpdatesAddMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add members to an updatableAssetGroup. You can add azureADDevice resources as members, but may not add **updatableAssetGroup** resources as members. Adding an Azure AD device as a member of an updatable asset group automatically creates an **azureADDevice** object, if it does not already exist. You can also use the method addMembersById to add members.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/windowsupdates-updatableassetgroup-addmembers?view=graph-rest-1.0
func (m *WindowsUpdatesDeploymentAudiencesItemMembersItemWindowsUpdatesAddMembersRequestBuilder) Post(ctx context.Context, body WindowsUpdatesDeploymentAudiencesItemMembersItemWindowsUpdatesAddMembersAddMembersPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentAudiencesItemMembersItemWindowsUpdatesAddMembersRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation add members to an updatableAssetGroup. You can add azureADDevice resources as members, but may not add **updatableAssetGroup** resources as members. Adding an Azure AD device as a member of an updatable asset group automatically creates an **azureADDevice** object, if it does not already exist. You can also use the method addMembersById to add members.
func (m *WindowsUpdatesDeploymentAudiencesItemMembersItemWindowsUpdatesAddMembersRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesDeploymentAudiencesItemMembersItemWindowsUpdatesAddMembersAddMembersPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentAudiencesItemMembersItemWindowsUpdatesAddMembersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
