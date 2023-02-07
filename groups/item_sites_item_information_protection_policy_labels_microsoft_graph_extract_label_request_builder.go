package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemInformationProtectionPolicyLabelsMicrosoftGraphExtractLabelRequestBuilder provides operations to call the extractLabel method.
type ItemSitesItemInformationProtectionPolicyLabelsMicrosoftGraphExtractLabelRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemInformationProtectionPolicyLabelsMicrosoftGraphExtractLabelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionPolicyLabelsMicrosoftGraphExtractLabelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemInformationProtectionPolicyLabelsMicrosoftGraphExtractLabelRequestBuilderInternal instantiates a new MicrosoftGraphExtractLabelRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionPolicyLabelsMicrosoftGraphExtractLabelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionPolicyLabelsMicrosoftGraphExtractLabelRequestBuilder) {
    m := &ItemSitesItemInformationProtectionPolicyLabelsMicrosoftGraphExtractLabelRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection/policy/labels/microsoft.graph.extractLabel";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesItemInformationProtectionPolicyLabelsMicrosoftGraphExtractLabelRequestBuilder instantiates a new MicrosoftGraphExtractLabelRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionPolicyLabelsMicrosoftGraphExtractLabelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionPolicyLabelsMicrosoftGraphExtractLabelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemInformationProtectionPolicyLabelsMicrosoftGraphExtractLabelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post using the metadata that exists on an already-labeled piece of information, resolve the metadata to a specific sensitivity label. The contentInfo input is resolved to informationProtectionContentLabel.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/informationprotectionlabel-extractlabel?view=graph-rest-1.0
func (m *ItemSitesItemInformationProtectionPolicyLabelsMicrosoftGraphExtractLabelRequestBuilder) Post(ctx context.Context, body ItemSitesItemInformationProtectionPolicyLabelsMicrosoftGraphExtractLabelExtractLabelPostRequestBodyable, requestConfiguration *ItemSitesItemInformationProtectionPolicyLabelsMicrosoftGraphExtractLabelRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionContentLabelable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInformationProtectionContentLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionContentLabelable), nil
}
// ToPostRequestInformation using the metadata that exists on an already-labeled piece of information, resolve the metadata to a specific sensitivity label. The contentInfo input is resolved to informationProtectionContentLabel.
func (m *ItemSitesItemInformationProtectionPolicyLabelsMicrosoftGraphExtractLabelRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSitesItemInformationProtectionPolicyLabelsMicrosoftGraphExtractLabelExtractLabelPostRequestBodyable, requestConfiguration *ItemSitesItemInformationProtectionPolicyLabelsMicrosoftGraphExtractLabelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
