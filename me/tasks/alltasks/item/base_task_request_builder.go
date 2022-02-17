package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i4d914957eb3084c5bc657dbc0bcea87fd6de701ecb15755c642a6ec8d4991865 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/alltasks/item/move"
    i5ff907d6b73cbfca73bba51a106f2a6632bf7114dcb344eff85a8986c70e23f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/alltasks/item/checklistitems"
    i63ee81a3dfe76eef9ca33b8a554278668a0c2207db00c252f06339df970ed8a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/alltasks/item/extensions"
    ide349d375a635575a72282b91fadf447d228db56856a5a438bddd7d5c0f84140 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/alltasks/item/linkedresources"
    ifd0bb0f9f690dbfe8c39f0fc370f4acb50e5071e4d99b8d5c83b43cc562eb0ed "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/alltasks/item/parentlist"
    i6820be8e1566091a81ecbd31899376a78e7c325b99449ad78a8ce0df252e0f9e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/alltasks/item/linkedresources/item"
    idf980e1c1f1f1a265b11cc926e016e5a7cf411227fcd5921660f5f9ccf16d5a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/alltasks/item/extensions/item"
    iff25132f9ec78dc6599b2dbf74fafad86c86d8fce4ee83662fd3257c45299fb9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks/alltasks/item/checklistitems/item"
)

// BaseTaskRequestBuilder builds and executes requests for operations under \me\tasks\alltasks\{baseTask-id}
type BaseTaskRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// BaseTaskRequestBuilderDeleteOptions options for Delete
type BaseTaskRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// BaseTaskRequestBuilderGetOptions options for Get
type BaseTaskRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *BaseTaskRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// BaseTaskRequestBuilderGetQueryParameters all tasks in the users mailbox.
type BaseTaskRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// BaseTaskRequestBuilderPatchOptions options for Patch
type BaseTaskRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BaseTask;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *BaseTaskRequestBuilder) ChecklistItems()(*i5ff907d6b73cbfca73bba51a106f2a6632bf7114dcb344eff85a8986c70e23f2.ChecklistItemsRequestBuilder) {
    return i5ff907d6b73cbfca73bba51a106f2a6632bf7114dcb344eff85a8986c70e23f2.NewChecklistItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChecklistItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.tasks.alltasks.item.checklistItems.item collection
func (m *BaseTaskRequestBuilder) ChecklistItemsById(id string)(*iff25132f9ec78dc6599b2dbf74fafad86c86d8fce4ee83662fd3257c45299fb9.ChecklistItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["checklistItem_id"] = id
    }
    return iff25132f9ec78dc6599b2dbf74fafad86c86d8fce4ee83662fd3257c45299fb9.NewChecklistItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewBaseTaskRequestBuilderInternal instantiates a new BaseTaskRequestBuilder and sets the default values.
func NewBaseTaskRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseTaskRequestBuilder) {
    m := &BaseTaskRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/tasks/alltasks/{baseTask_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBaseTaskRequestBuilder instantiates a new BaseTaskRequestBuilder and sets the default values.
func NewBaseTaskRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BaseTaskRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBaseTaskRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation all tasks in the users mailbox.
func (m *BaseTaskRequestBuilder) CreateDeleteRequestInformation(options *BaseTaskRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation all tasks in the users mailbox.
func (m *BaseTaskRequestBuilder) CreateGetRequestInformation(options *BaseTaskRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation all tasks in the users mailbox.
func (m *BaseTaskRequestBuilder) CreatePatchRequestInformation(options *BaseTaskRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete all tasks in the users mailbox.
func (m *BaseTaskRequestBuilder) Delete(options *BaseTaskRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *BaseTaskRequestBuilder) Extensions()(*i63ee81a3dfe76eef9ca33b8a554278668a0c2207db00c252f06339df970ed8a8.ExtensionsRequestBuilder) {
    return i63ee81a3dfe76eef9ca33b8a554278668a0c2207db00c252f06339df970ed8a8.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.tasks.alltasks.item.extensions.item collection
func (m *BaseTaskRequestBuilder) ExtensionsById(id string)(*idf980e1c1f1f1a265b11cc926e016e5a7cf411227fcd5921660f5f9ccf16d5a5.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return idf980e1c1f1f1a265b11cc926e016e5a7cf411227fcd5921660f5f9ccf16d5a5.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get all tasks in the users mailbox.
func (m *BaseTaskRequestBuilder) Get(options *BaseTaskRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BaseTask, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewBaseTask() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BaseTask), nil
}
func (m *BaseTaskRequestBuilder) LinkedResources()(*ide349d375a635575a72282b91fadf447d228db56856a5a438bddd7d5c0f84140.LinkedResourcesRequestBuilder) {
    return ide349d375a635575a72282b91fadf447d228db56856a5a438bddd7d5c0f84140.NewLinkedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LinkedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.tasks.alltasks.item.linkedResources.item collection
func (m *BaseTaskRequestBuilder) LinkedResourcesById(id string)(*i6820be8e1566091a81ecbd31899376a78e7c325b99449ad78a8ce0df252e0f9e.LinkedResource_v2RequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["linkedResource_v2_id"] = id
    }
    return i6820be8e1566091a81ecbd31899376a78e7c325b99449ad78a8ce0df252e0f9e.NewLinkedResource_v2RequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *BaseTaskRequestBuilder) Move()(*i4d914957eb3084c5bc657dbc0bcea87fd6de701ecb15755c642a6ec8d4991865.MoveRequestBuilder) {
    return i4d914957eb3084c5bc657dbc0bcea87fd6de701ecb15755c642a6ec8d4991865.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *BaseTaskRequestBuilder) ParentList()(*ifd0bb0f9f690dbfe8c39f0fc370f4acb50e5071e4d99b8d5c83b43cc562eb0ed.ParentListRequestBuilder) {
    return ifd0bb0f9f690dbfe8c39f0fc370f4acb50e5071e4d99b8d5c83b43cc562eb0ed.NewParentListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch all tasks in the users mailbox.
func (m *BaseTaskRequestBuilder) Patch(options *BaseTaskRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
