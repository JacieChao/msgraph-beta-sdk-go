package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilder provides operations to call the snoozeReminder method.
type ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilderInternal instantiates a new SnoozeReminderRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilder) {
    m := &ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}/snoozeReminder", pathParameters),
    }
    return m
}
// NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilder instantiates a new SnoozeReminderRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action snoozeReminder
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilder) Post(ctx context.Context, body ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSnoozeReminderPostRequestBodyable, requestConfiguration *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action snoozeReminder
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSnoozeReminderPostRequestBodyable, requestConfiguration *ItemCalendarGroupsItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
