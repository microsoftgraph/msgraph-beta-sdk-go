package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i13a4289526bad4d09f46c002819318815cb5009ad9d6118977993c39c71eb713 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/notebooks/item/sectiongroups"
    i8cae1002445f13514cd17214ebde1602548f0fdeccb72b40e8c2641de30b8f95 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/notebooks/item/copynotebook"
    iaa797de12b2ff9b5e75d4e1fdd4aba3c50d394de5195ed44b0962261d1d45ffd "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/notebooks/item/sections"
    i4e0c85405d8104211426a14e5e081e5618c5ddabddb25835ec8c34eeeca7eb1d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/notebooks/item/sections/item"
    i8c0c1cccb46ede336960e9bc9e30e28f1d0b03c22b304bcc9b4c015ff449625e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/notebooks/item/sectiongroups/item"
)

// NotebookRequestBuilder builds and executes requests for operations under \groups\{group-id}\onenote\notebooks\{notebook-id}
type NotebookRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// NotebookRequestBuilderDeleteOptions options for Delete
type NotebookRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NotebookRequestBuilderGetOptions options for Get
type NotebookRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *NotebookRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NotebookRequestBuilderGetQueryParameters the collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable.
type NotebookRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NotebookRequestBuilderPatchOptions options for Patch
type NotebookRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Notebook;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewNotebookRequestBuilderInternal instantiates a new NotebookRequestBuilder and sets the default values.
func NewNotebookRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*NotebookRequestBuilder) {
    m := &NotebookRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/onenote/notebooks/{notebook_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewNotebookRequestBuilder instantiates a new NotebookRequestBuilder and sets the default values.
func NewNotebookRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*NotebookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNotebookRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *NotebookRequestBuilder) CopyNotebook()(*i8cae1002445f13514cd17214ebde1602548f0fdeccb72b40e8c2641de30b8f95.CopyNotebookRequestBuilder) {
    return i8cae1002445f13514cd17214ebde1602548f0fdeccb72b40e8c2641de30b8f95.NewCopyNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable.
func (m *NotebookRequestBuilder) CreateDeleteRequestInformation(options *NotebookRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable.
func (m *NotebookRequestBuilder) CreateGetRequestInformation(options *NotebookRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable.
func (m *NotebookRequestBuilder) CreatePatchRequestInformation(options *NotebookRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable.
func (m *NotebookRequestBuilder) Delete(options *NotebookRequestBuilderDeleteOptions)(error) {
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
// Get the collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable.
func (m *NotebookRequestBuilder) Get(options *NotebookRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Notebook, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewNotebook() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Notebook), nil
}
// Patch the collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable.
func (m *NotebookRequestBuilder) Patch(options *NotebookRequestBuilderPatchOptions)(error) {
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
func (m *NotebookRequestBuilder) SectionGroups()(*i13a4289526bad4d09f46c002819318815cb5009ad9d6118977993c39c71eb713.SectionGroupsRequestBuilder) {
    return i13a4289526bad4d09f46c002819318815cb5009ad9d6118977993c39c71eb713.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.onenote.notebooks.item.sectionGroups.item collection
func (m *NotebookRequestBuilder) SectionGroupsById(id string)(*i8c0c1cccb46ede336960e9bc9e30e28f1d0b03c22b304bcc9b4c015ff449625e.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id"] = id
    }
    return i8c0c1cccb46ede336960e9bc9e30e28f1d0b03c22b304bcc9b4c015ff449625e.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *NotebookRequestBuilder) Sections()(*iaa797de12b2ff9b5e75d4e1fdd4aba3c50d394de5195ed44b0962261d1d45ffd.SectionsRequestBuilder) {
    return iaa797de12b2ff9b5e75d4e1fdd4aba3c50d394de5195ed44b0962261d1d45ffd.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.onenote.notebooks.item.sections.item collection
func (m *NotebookRequestBuilder) SectionsById(id string)(*i4e0c85405d8104211426a14e5e081e5618c5ddabddb25835ec8c34eeeca7eb1d.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return i4e0c85405d8104211426a14e5e081e5618c5ddabddb25835ec8c34eeeca7eb1d.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
