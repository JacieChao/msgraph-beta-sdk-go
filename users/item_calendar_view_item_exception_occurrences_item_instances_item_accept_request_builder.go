package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilder provides operations to call the accept method.
type ItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilderInternal instantiates a new AcceptRequestBuilder and sets the default values.
func NewItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilder) {
    m := &ItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}/accept", pathParameters),
    }
    return m
}
// NewItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilder instantiates a new AcceptRequestBuilder and sets the default values.
func NewItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilderInternal(urlParams, requestAdapter)
}
// Post accept the specified event in a user calendar. This API is supported in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/event-accept?view=graph-rest-1.0
func (m *ItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilder) Post(ctx context.Context, body ItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptPostRequestBodyable, requestConfiguration *ItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation accept the specified event in a user calendar. This API is supported in the following national cloud deployments.
func (m *ItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptPostRequestBodyable, requestConfiguration *ItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilder) {
    return NewItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
