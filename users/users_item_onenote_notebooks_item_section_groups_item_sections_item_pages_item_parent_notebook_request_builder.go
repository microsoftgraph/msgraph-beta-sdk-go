package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilder provides operations to manage the parentNotebook property of the microsoft.graph.onenotePage entity.
type UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilderGetQueryParameters the notebook that contains the page.  Read-only.
type UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilderGetQueryParameters
}
// NewUsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilderInternal instantiates a new ParentNotebookRequestBuilder and sets the default values.
func NewUsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilder) {
    m := &UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/onenote/notebooks/{notebook%2Did}/sectionGroups/{sectionGroup%2Did}/sections/{onenoteSection%2Did}/pages/{onenotePage%2Did}/parentNotebook{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilder instantiates a new ParentNotebookRequestBuilder and sets the default values.
func NewUsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the notebook that contains the page.  Read-only.
func (m *UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get the notebook that contains the page.  Read-only.
func (m *UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Notebookable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateNotebookFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Notebookable), nil
}
