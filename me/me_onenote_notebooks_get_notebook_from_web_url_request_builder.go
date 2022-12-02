package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeOnenoteNotebooksGetNotebookFromWebUrlRequestBuilder provides operations to call the getNotebookFromWebUrl method.
type MeOnenoteNotebooksGetNotebookFromWebUrlRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeOnenoteNotebooksGetNotebookFromWebUrlRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeOnenoteNotebooksGetNotebookFromWebUrlRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeOnenoteNotebooksGetNotebookFromWebUrlRequestBuilderInternal instantiates a new GetNotebookFromWebUrlRequestBuilder and sets the default values.
func NewMeOnenoteNotebooksGetNotebookFromWebUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeOnenoteNotebooksGetNotebookFromWebUrlRequestBuilder) {
    m := &MeOnenoteNotebooksGetNotebookFromWebUrlRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/onenote/notebooks/microsoft.graph.getNotebookFromWebUrl";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeOnenoteNotebooksGetNotebookFromWebUrlRequestBuilder instantiates a new GetNotebookFromWebUrlRequestBuilder and sets the default values.
func NewMeOnenoteNotebooksGetNotebookFromWebUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeOnenoteNotebooksGetNotebookFromWebUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeOnenoteNotebooksGetNotebookFromWebUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation retrieve the properties and relationships of a notebook object by using its URL path. The location can be user notebooks on Microsoft 365, group notebooks, or SharePoint site-hosted team notebooks on Microsoft 365.
func (m *MeOnenoteNotebooksGetNotebookFromWebUrlRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeOnenoteNotebooksGetNotebookFromWebUrlPostRequestBodyable, requestConfiguration *MeOnenoteNotebooksGetNotebookFromWebUrlRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post retrieve the properties and relationships of a notebook object by using its URL path. The location can be user notebooks on Microsoft 365, group notebooks, or SharePoint site-hosted team notebooks on Microsoft 365.
func (m *MeOnenoteNotebooksGetNotebookFromWebUrlRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeOnenoteNotebooksGetNotebookFromWebUrlPostRequestBodyable, requestConfiguration *MeOnenoteNotebooksGetNotebookFromWebUrlRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CopyNotebookModelable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCopyNotebookModelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CopyNotebookModelable), nil
}
