package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilder provides operations to call the accept method.
type ItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilderInternal instantiates a new AcceptRequestBuilder and sets the default values.
func NewItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilder) {
    m := &ItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}/microsoft.graph.accept";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilder instantiates a new AcceptRequestBuilder and sets the default values.
func NewItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation accept the specified event in a user calendar.
func (m *ItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptPostRequestBodyable, requestConfiguration *ItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post accept the specified event in a user calendar.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/event-accept?view=graph-rest-1.0
func (m *ItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilder) Post(ctx context.Context, body ItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptPostRequestBodyable, requestConfiguration *ItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
