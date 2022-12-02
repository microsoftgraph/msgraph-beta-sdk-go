package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeMailFoldersItemMessagesItemCreateForwardRequestBuilder provides operations to call the createForward method.
type MeMailFoldersItemMessagesItemCreateForwardRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeMailFoldersItemMessagesItemCreateForwardRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeMailFoldersItemMessagesItemCreateForwardRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeMailFoldersItemMessagesItemCreateForwardRequestBuilderInternal instantiates a new CreateForwardRequestBuilder and sets the default values.
func NewMeMailFoldersItemMessagesItemCreateForwardRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeMailFoldersItemMessagesItemCreateForwardRequestBuilder) {
    m := &MeMailFoldersItemMessagesItemCreateForwardRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/mailFolders/{mailFolder%2Did}/messages/{message%2Did}/microsoft.graph.createForward";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeMailFoldersItemMessagesItemCreateForwardRequestBuilder instantiates a new CreateForwardRequestBuilder and sets the default values.
func NewMeMailFoldersItemMessagesItemCreateForwardRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeMailFoldersItemMessagesItemCreateForwardRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeMailFoldersItemMessagesItemCreateForwardRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation create a draft to forward an existing message, in either JSON or MIME format. When using JSON format, you can: - Specify either a comment or the **body** property of the `message` parameter. Specifying both will return an HTTP 400 Bad Request error.- Specify either the `toRecipients` parameter or the **toRecipients** property of the `message` parameter. Specifying both or specifying neither will return an HTTP 400 Bad Request error.- Update the draft later to add content to the **body** or change other message properties. When using MIME format:- Provide the applicable Internet message headers and the MIME content, all encoded in **base64** format in the request body.- Add any attachments and S/MIME properties to the MIME content. Send the draft message in a subsequent operation. Alternatively, forward a message in a single operation.
func (m *MeMailFoldersItemMessagesItemCreateForwardRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeMailFoldersItemMessagesItemCreateForwardPostRequestBodyable, requestConfiguration *MeMailFoldersItemMessagesItemCreateForwardRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post create a draft to forward an existing message, in either JSON or MIME format. When using JSON format, you can: - Specify either a comment or the **body** property of the `message` parameter. Specifying both will return an HTTP 400 Bad Request error.- Specify either the `toRecipients` parameter or the **toRecipients** property of the `message` parameter. Specifying both or specifying neither will return an HTTP 400 Bad Request error.- Update the draft later to add content to the **body** or change other message properties. When using MIME format:- Provide the applicable Internet message headers and the MIME content, all encoded in **base64** format in the request body.- Add any attachments and S/MIME properties to the MIME content. Send the draft message in a subsequent operation. Alternatively, forward a message in a single operation.
func (m *MeMailFoldersItemMessagesItemCreateForwardRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeMailFoldersItemMessagesItemCreateForwardPostRequestBodyable, requestConfiguration *MeMailFoldersItemMessagesItemCreateForwardRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable), nil
}
