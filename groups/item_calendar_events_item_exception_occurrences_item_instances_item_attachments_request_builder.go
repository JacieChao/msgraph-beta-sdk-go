package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilder provides operations to manage the attachments property of the microsoft.graph.event entity.
type ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilderGetQueryParameters retrieve a list of attachment objects attached to an event.
type ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilderGetQueryParameters
}
// ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilderInternal instantiates a new AttachmentsRequestBuilder and sets the default values.
func NewItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilder) {
    m := &ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendar/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}/attachments{?%24top,%24skip,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilder instantiates a new AttachmentsRequestBuilder and sets the default values.
func NewItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilder) Count()(*ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilder) {
    return NewItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateUploadSession provides operations to call the createUploadSession method.
func (m *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilder) CreateUploadSession()(*ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCreateUploadSessionRequestBuilder) {
    return NewItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get retrieve a list of attachment objects attached to an event.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/event-list-attachments?view=graph-rest-1.0
func (m *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttachmentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttachmentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttachmentCollectionResponseable), nil
}
// Post use this API to create a new Attachment. An attachment can be one of the following types: All these types of attachment resources are derived from the attachmentresource. 
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/eventmessage-post-attachments?view=graph-rest-1.0
func (m *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Attachmentable, requestConfiguration *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Attachmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttachmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Attachmentable), nil
}
// ToGetRequestInformation retrieve a list of attachment objects attached to an event.
func (m *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPostRequestInformation use this API to create a new Attachment. An attachment can be one of the following types: All these types of attachment resources are derived from the attachmentresource. 
func (m *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Attachmentable, requestConfiguration *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}