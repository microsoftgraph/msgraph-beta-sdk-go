package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOnenoteSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder provides operations to call the copyToSectionGroup method.
type ItemOnenoteSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnenoteSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOnenoteSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilderInternal instantiates a new CopyToSectionGroupRequestBuilder and sets the default values.
func NewItemOnenoteSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder) {
    m := &ItemOnenoteSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/onenote/sectionGroups/{sectionGroup%2Did}/sections/{onenoteSection%2Did}/copyToSectionGroup", pathParameters),
    }
    return m
}
// NewItemOnenoteSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder instantiates a new CopyToSectionGroupRequestBuilder and sets the default values.
func NewItemOnenoteSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnenoteSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Post copies a section to a specific section group. For Copy operations, you follow an asynchronous calling pattern:  First call the Copy action, and then poll the operation endpoint for the result. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/section-copytosectiongroup?view=graph-rest-1.0
func (m *ItemOnenoteSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder) Post(ctx context.Context, body ItemOnenoteSectionGroupsItemSectionsItemCopyToSectionGroupPostRequestBodyable, requestConfiguration *ItemOnenoteSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteOperationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnenoteOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteOperationable), nil
}
// ToPostRequestInformation copies a section to a specific section group. For Copy operations, you follow an asynchronous calling pattern:  First call the Copy action, and then poll the operation endpoint for the result. This API is available in the following national cloud deployments.
func (m *ItemOnenoteSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemOnenoteSectionGroupsItemSectionsItemCopyToSectionGroupPostRequestBodyable, requestConfiguration *ItemOnenoteSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemOnenoteSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder) WithUrl(rawUrl string)(*ItemOnenoteSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder) {
    return NewItemOnenoteSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
