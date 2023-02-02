package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemInformationProtectionMicrosoftGraphEncryptBufferEncryptBufferRequestBuilder provides operations to call the encryptBuffer method.
type ItemSitesItemInformationProtectionMicrosoftGraphEncryptBufferEncryptBufferRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemInformationProtectionMicrosoftGraphEncryptBufferEncryptBufferRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionMicrosoftGraphEncryptBufferEncryptBufferRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemInformationProtectionMicrosoftGraphEncryptBufferEncryptBufferRequestBuilderInternal instantiates a new EncryptBufferRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionMicrosoftGraphEncryptBufferEncryptBufferRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionMicrosoftGraphEncryptBufferEncryptBufferRequestBuilder) {
    m := &ItemSitesItemInformationProtectionMicrosoftGraphEncryptBufferEncryptBufferRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection/microsoft.graph.encryptBuffer";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesItemInformationProtectionMicrosoftGraphEncryptBufferEncryptBufferRequestBuilder instantiates a new EncryptBufferRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionMicrosoftGraphEncryptBufferEncryptBufferRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionMicrosoftGraphEncryptBufferEncryptBufferRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemInformationProtectionMicrosoftGraphEncryptBufferEncryptBufferRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action encryptBuffer
func (m *ItemSitesItemInformationProtectionMicrosoftGraphEncryptBufferEncryptBufferRequestBuilder) Post(ctx context.Context, body ItemSitesItemInformationProtectionMicrosoftGraphEncryptBufferEncryptBufferPostRequestBodyable, requestConfiguration *ItemSitesItemInformationProtectionMicrosoftGraphEncryptBufferEncryptBufferRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BufferEncryptionResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBufferEncryptionResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BufferEncryptionResultable), nil
}
// ToPostRequestInformation invoke action encryptBuffer
func (m *ItemSitesItemInformationProtectionMicrosoftGraphEncryptBufferEncryptBufferRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSitesItemInformationProtectionMicrosoftGraphEncryptBufferEncryptBufferPostRequestBodyable, requestConfiguration *ItemSitesItemInformationProtectionMicrosoftGraphEncryptBufferEncryptBufferRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
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
