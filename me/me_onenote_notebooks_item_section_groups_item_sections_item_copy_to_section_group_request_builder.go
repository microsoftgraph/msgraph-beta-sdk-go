package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder provides operations to call the copyToSectionGroup method.
type MeOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilderInternal instantiates a new CopyToSectionGroupRequestBuilder and sets the default values.
func NewMeOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder) {
    m := &MeOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/onenote/notebooks/{notebook%2Did}/sectionGroups/{sectionGroup%2Did}/sections/{onenoteSection%2Did}/microsoft.graph.copyToSectionGroup";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder instantiates a new CopyToSectionGroupRequestBuilder and sets the default values.
func NewMeOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation copies a section to a specific section group. For Copy operations, you follow an asynchronous calling pattern:  First call the Copy action, and then poll the operation endpoint for the result.
func (m *MeOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToSectionGroupPostRequestBodyable, requestConfiguration *MeOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post copies a section to a specific section group. For Copy operations, you follow an asynchronous calling pattern:  First call the Copy action, and then poll the operation endpoint for the result.
func (m *MeOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToSectionGroupPostRequestBodyable, requestConfiguration *MeOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteOperationable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnenoteOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteOperationable), nil
}
