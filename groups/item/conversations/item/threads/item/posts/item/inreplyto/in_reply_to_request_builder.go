package inreplyto

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i01a68f9780cd1f46a79f2dc863055bd2c5fd0b52d8fbc97136600da9a39be794 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto/multivalueextendedproperties"
    i1d2403878bf8da0493447df951161dbafd3178e061d2670a89a318f9bfc06ad6 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto/singlevalueextendedproperties"
    i235cd8bebbce382226651ea1277f498ee89c8f0f383184bfadf8d37fbd548d80 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto/forward"
    i340b1e7815d335dd75428026204f77ed63767269b5f30448456d513e1be7aee1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto/attachments"
    i96a63a39bd69c9fa664b91c15cc7805f22764600de5c802e46d1cc329315cce5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto/mentions"
    i96f246ed552519bb033453ca223a1f6da6e5f2067b722bc22ab9f8838083c5b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto/reply"
    ifa5655936308bf17dffd49a6bd3ba8b75023a0ad8079ecd7ed031e904d94a8e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto/extensions"
    i4d4f4736b20b533874b21d43afdeb8387415582e42ec98e9b5f3a0834591f636 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto/attachments/item"
    i5ae608db00267db06b6efd5d156fb58202137670bc8f0b2db988c7647811c893 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto/singlevalueextendedproperties/item"
    i627ac24b1c5154b00c54703eaed3391660a6f1b701cbc4e8f4cafc57218df556 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto/mentions/item"
    ibdb87644e5d9ad4b691e589620d652f65dd8e136822c1a12d53aa5971dd79ace "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto/extensions/item"
    id25bdfa22a85826c2d415ddac3ea84ff50c4050e7452b38f37b55fbaa6e42145 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/conversations/item/threads/item/posts/item/inreplyto/multivalueextendedproperties/item"
)

// InReplyToRequestBuilder provides operations to manage the inReplyTo property of the microsoft.graph.post entity.
type InReplyToRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// InReplyToRequestBuilderDeleteOptions options for Delete
type InReplyToRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// InReplyToRequestBuilderGetOptions options for Get
type InReplyToRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *InReplyToRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// InReplyToRequestBuilderGetQueryParameters read-only. Supports $expand.
type InReplyToRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// InReplyToRequestBuilderPatchOptions options for Patch
type InReplyToRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Attachments the attachments property
func (m *InReplyToRequestBuilder) Attachments()(*i340b1e7815d335dd75428026204f77ed63767269b5f30448456d513e1be7aee1.AttachmentsRequestBuilder) {
    return i340b1e7815d335dd75428026204f77ed63767269b5f30448456d513e1be7aee1.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.conversations.item.threads.item.posts.item.inReplyTo.attachments.item collection
func (m *InReplyToRequestBuilder) AttachmentsById(id string)(*i4d4f4736b20b533874b21d43afdeb8387415582e42ec98e9b5f3a0834591f636.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i4d4f4736b20b533874b21d43afdeb8387415582e42ec98e9b5f3a0834591f636.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewInReplyToRequestBuilderInternal instantiates a new InReplyToRequestBuilder and sets the default values.
func NewInReplyToRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InReplyToRequestBuilder) {
    m := &InReplyToRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/conversations/{conversation_id}/threads/{conversationThread_id}/posts/{post_id}/inReplyTo{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewInReplyToRequestBuilder instantiates a new InReplyToRequestBuilder and sets the default values.
func NewInReplyToRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InReplyToRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInReplyToRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property inReplyTo for groups
func (m *InReplyToRequestBuilder) CreateDeleteRequestInformation(options *InReplyToRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation read-only. Supports $expand.
func (m *InReplyToRequestBuilder) CreateGetRequestInformation(options *InReplyToRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property inReplyTo in groups
func (m *InReplyToRequestBuilder) CreatePatchRequestInformation(options *InReplyToRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property inReplyTo for groups
func (m *InReplyToRequestBuilder) Delete(options *InReplyToRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Extensions the extensions property
func (m *InReplyToRequestBuilder) Extensions()(*ifa5655936308bf17dffd49a6bd3ba8b75023a0ad8079ecd7ed031e904d94a8e2.ExtensionsRequestBuilder) {
    return ifa5655936308bf17dffd49a6bd3ba8b75023a0ad8079ecd7ed031e904d94a8e2.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.conversations.item.threads.item.posts.item.inReplyTo.extensions.item collection
func (m *InReplyToRequestBuilder) ExtensionsById(id string)(*ibdb87644e5d9ad4b691e589620d652f65dd8e136822c1a12d53aa5971dd79ace.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ibdb87644e5d9ad4b691e589620d652f65dd8e136822c1a12d53aa5971dd79ace.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *InReplyToRequestBuilder) Forward()(*i235cd8bebbce382226651ea1277f498ee89c8f0f383184bfadf8d37fbd548d80.ForwardRequestBuilder) {
    return i235cd8bebbce382226651ea1277f498ee89c8f0f383184bfadf8d37fbd548d80.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get read-only. Supports $expand.
func (m *InReplyToRequestBuilder) Get(options *InReplyToRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePostFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable), nil
}
// Mentions the mentions property
func (m *InReplyToRequestBuilder) Mentions()(*i96a63a39bd69c9fa664b91c15cc7805f22764600de5c802e46d1cc329315cce5.MentionsRequestBuilder) {
    return i96a63a39bd69c9fa664b91c15cc7805f22764600de5c802e46d1cc329315cce5.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.conversations.item.threads.item.posts.item.inReplyTo.mentions.item collection
func (m *InReplyToRequestBuilder) MentionsById(id string)(*i627ac24b1c5154b00c54703eaed3391660a6f1b701cbc4e8f4cafc57218df556.MentionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention_id"] = id
    }
    return i627ac24b1c5154b00c54703eaed3391660a6f1b701cbc4e8f4cafc57218df556.NewMentionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *InReplyToRequestBuilder) MultiValueExtendedProperties()(*i01a68f9780cd1f46a79f2dc863055bd2c5fd0b52d8fbc97136600da9a39be794.MultiValueExtendedPropertiesRequestBuilder) {
    return i01a68f9780cd1f46a79f2dc863055bd2c5fd0b52d8fbc97136600da9a39be794.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.conversations.item.threads.item.posts.item.inReplyTo.multiValueExtendedProperties.item collection
func (m *InReplyToRequestBuilder) MultiValueExtendedPropertiesById(id string)(*id25bdfa22a85826c2d415ddac3ea84ff50c4050e7452b38f37b55fbaa6e42145.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return id25bdfa22a85826c2d415ddac3ea84ff50c4050e7452b38f37b55fbaa6e42145.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property inReplyTo in groups
func (m *InReplyToRequestBuilder) Patch(options *InReplyToRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Reply the reply property
func (m *InReplyToRequestBuilder) Reply()(*i96f246ed552519bb033453ca223a1f6da6e5f2067b722bc22ab9f8838083c5b8.ReplyRequestBuilder) {
    return i96f246ed552519bb033453ca223a1f6da6e5f2067b722bc22ab9f8838083c5b8.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *InReplyToRequestBuilder) SingleValueExtendedProperties()(*i1d2403878bf8da0493447df951161dbafd3178e061d2670a89a318f9bfc06ad6.SingleValueExtendedPropertiesRequestBuilder) {
    return i1d2403878bf8da0493447df951161dbafd3178e061d2670a89a318f9bfc06ad6.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.conversations.item.threads.item.posts.item.inReplyTo.singleValueExtendedProperties.item collection
func (m *InReplyToRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i5ae608db00267db06b6efd5d156fb58202137670bc8f0b2db988c7647811c893.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i5ae608db00267db06b6efd5d156fb58202137670bc8f0b2db988c7647811c893.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
