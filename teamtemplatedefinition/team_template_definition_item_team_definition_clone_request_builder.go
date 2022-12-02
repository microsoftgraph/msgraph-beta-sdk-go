package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplateDefinitionItemTeamDefinitionCloneRequestBuilder provides operations to call the clone method.
type TeamTemplateDefinitionItemTeamDefinitionCloneRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamTemplateDefinitionItemTeamDefinitionCloneRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplateDefinitionItemTeamDefinitionCloneRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTeamTemplateDefinitionItemTeamDefinitionCloneRequestBuilderInternal instantiates a new CloneRequestBuilder and sets the default values.
func NewTeamTemplateDefinitionItemTeamDefinitionCloneRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplateDefinitionItemTeamDefinitionCloneRequestBuilder) {
    m := &TeamTemplateDefinitionItemTeamDefinitionCloneRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/microsoft.graph.clone";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamTemplateDefinitionItemTeamDefinitionCloneRequestBuilder instantiates a new CloneRequestBuilder and sets the default values.
func NewTeamTemplateDefinitionItemTeamDefinitionCloneRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplateDefinitionItemTeamDefinitionCloneRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplateDefinitionItemTeamDefinitionCloneRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation create a copy of a team. This operation also creates a copy of the corresponding group.You can specify which parts of the team to clone: When tabs are cloned, they are put into an unconfigured state -- they are displayed on the tab bar in Microsoft Teams, and the first time you open them, you'll go through the configuration screen. (If the person opening the tab does not have permission to configure apps, they will see a message explaining that the tab hasn't been configured.) Cloning is a long-running operation.After the POST clone returns, you need to GET the operation returned by the Location: header to see if it's 'running' or 'succeeded' or 'failed'. You should continue to GET until the status is not 'running'. The recommended delay between GETs is 5 seconds.
func (m *TeamTemplateDefinitionItemTeamDefinitionCloneRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamTemplateDefinitionItemTeamDefinitionClonePostRequestBodyable, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionCloneRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post create a copy of a team. This operation also creates a copy of the corresponding group.You can specify which parts of the team to clone: When tabs are cloned, they are put into an unconfigured state -- they are displayed on the tab bar in Microsoft Teams, and the first time you open them, you'll go through the configuration screen. (If the person opening the tab does not have permission to configure apps, they will see a message explaining that the tab hasn't been configured.) Cloning is a long-running operation.After the POST clone returns, you need to GET the operation returned by the Location: header to see if it's 'running' or 'succeeded' or 'failed'. You should continue to GET until the status is not 'running'. The recommended delay between GETs is 5 seconds.
func (m *TeamTemplateDefinitionItemTeamDefinitionCloneRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamTemplateDefinitionItemTeamDefinitionClonePostRequestBodyable, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionCloneRequestBuilderPostRequestConfiguration)(error) {
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
