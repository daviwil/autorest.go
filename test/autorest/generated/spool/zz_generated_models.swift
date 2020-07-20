// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.0.6302, generator: {generator})
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import Foundation

import (
	"encoding/json"
	"net/http"
	"time"
)

type AddThreadMembersRequest struct {
	// Members to add to a thread.
	Members *[]ThreadMember `json:"members,omitempty"`
}

// ChatAddThreadMembersOptions contains the optional parameters for the Chat.AddThreadMembers method.
type ChatAddThreadMembersOptions struct {
	// Thread members to be added to the thread.
	Body *AddThreadMembersRequest
	// Correlation vector, if a value is not provided a randomly generated correlation vector would be returned in the response
// headers.
	Mscv *string
}

// ChatCreateThreadOptions contains the optional parameters for the Chat.CreateThread method.
type ChatCreateThreadOptions struct {
	// Request payload for creating a chat thread.
	Body *CreateThreadRequest
	// Correlation vector, if a value is not provided a randomly generated correlation vector would be returned in the response
// headers.
	Mscv *string
}

// ChatDeleteMessageOptions contains the optional parameters for the Chat.DeleteMessage method.
type ChatDeleteMessageOptions struct {
	// Correlation vector, if a value is not provided a randomly generated correlation vector would be returned in the response
// headers.
	Mscv *string
}

// ChatDeleteThreadOptions contains the optional parameters for the Chat.DeleteThread method.
type ChatDeleteThreadOptions struct {
	// Correlation vector, if a value is not provided a randomly generated correlation vector would be returned in the response
// headers.
	Mscv *string
}

// ChatGetMessageOptions contains the optional parameters for the Chat.GetMessage method.
type ChatGetMessageOptions struct {
	// Correlation vector, if a value is not provided a randomly generated correlation vector would be returned in the response
// headers.
	Mscv *string
}

// ChatGetMessagesOptions contains the optional parameters for the Chat.GetMessages method.
type ChatGetMessagesOptions struct {
	// Correlation vector, if a value is not provided a randomly generated correlation vector would be returned in the response
// headers.
	Mscv *string
	// The number of messages being requested.
	PageSize *int32
	// The start time where the range query. This is represented by number of seconds since epoch time.
	StartTime *int64
	// The continuation token that previous request obtained. This is used for paging.
	SyncState *string
}

// ChatGetReadReceiptsOptions contains the optional parameters for the Chat.GetReadReceipts method.
type ChatGetReadReceiptsOptions struct {
	// Correlation vector, if a value is not provided a randomly generated correlation vector would be returned in the response
// headers.
	Mscv *string
}

// ChatGetThreadMembersOptions contains the optional parameters for the Chat.GetThreadMembers method.
type ChatGetThreadMembersOptions struct {
	// Correlation vector, if a value is not provided a randomly generated correlation vector would be returned in the response
// headers.
	Mscv *string
}

// ChatGetThreadOptions contains the optional parameters for the Chat.GetThread method.
type ChatGetThreadOptions struct {
	// Correlation vector, if a value is not provided a randomly generated correlation vector would be returned in the response
// headers.
	Mscv *string
}

// ChatGetThreadsOptions contains the optional parameters for the Chat.GetThreads method.
type ChatGetThreadsOptions struct {
	// Correlation vector, if a value is not provided a randomly generated correlation vector would be returned in the response
// headers.
	Mscv *string
	// The number of threads being requested.
	PageSize *int32
	// The start time where the range query. This is represented by number of seconds since epoch time.
	StartTime *int64
	// The continuation token that previous request obtained. This is used for paging.
	SyncState *string
}

// ChatNotifyUserTypingOptions contains the optional parameters for the Chat.NotifyUserTyping method.
type ChatNotifyUserTypingOptions struct {
	// Correlation vector, if a value is not provided a randomly generated correlation vector would be returned in the response
// headers.
	Mscv *string
}

// ChatRemoveThreadMemberOptions contains the optional parameters for the Chat.RemoveThreadMember method.
type ChatRemoveThreadMemberOptions struct {
	// Correlation vector, if a value is not provided a randomly generated correlation vector would be returned in the response
// headers.
	Mscv *string
}

// ChatSendMessageOptions contains the optional parameters for the Chat.SendMessage method.
type ChatSendMessageOptions struct {
	// Details of the message to create.
	Body *CreateMessageRequest
	// Correlation vector, if a value is not provided a randomly generated correlation vector would be returned in the response
// headers.
	Mscv *string
}

// ChatSendReadReceiptOptions contains the optional parameters for the Chat.SendReadReceipt method.
type ChatSendReadReceiptOptions struct {
	// Request payload for sending a read receipt.
	Body *PostReadReceiptRequest
	// Correlation vector, if a value is not provided a randomly generated correlation vector would be returned in the response
// headers.
	Mscv *string
}

// ChatUpdateMessageOptions contains the optional parameters for the Chat.UpdateMessage method.
type ChatUpdateMessageOptions struct {
	// Details of the request to update the message.
	Body *UpdateMessageRequest
	// Correlation vector, if a value is not provided a randomly generated correlation vector would be returned in the response
// headers.
	Mscv *string
}

// ChatUpdateThreadOptions contains the optional parameters for the Chat.UpdateThread method.
type ChatUpdateThreadOptions struct {
	// Request payload for updating a chat thread.
	Body *UpdateThreadRequest
	// Correlation vector, if a value is not provided a randomly generated correlation vector would be returned in the response
// headers.
	Mscv *string
}

type CreateMessageRequest struct {
	// This Id is a client-specific Id in a numeric unsigned Int64 format. It can be used for client deduping, among other client
// usages.
	ClientMessageID *string `json:"clientMessageId,omitempty"`

	// Chat message content.
	Content *string `json:"content,omitempty"`
	Priority *MessagePriority `json:"priority,omitempty"`

	// The display name of the message sender. This property is used to populate sender name for push notifications.
	SenderDisplayName *string `json:"senderDisplayName,omitempty"`
}

type CreateMessageResponse struct {
	// This Id is a client-specific Id in a numeric unsigned Int64 format. It can be used for client deduping, among other client
// usages.
	ClientMessageID *string `json:"clientMessageId,omitempty"`

	// The Id of the message. This Id is server generated.
	ID *string `json:"id,omitempty"`
}

// CreateMessageResponseResponse is the response envelope for operations that return a CreateMessageResponse type.
type CreateMessageResponseResponse struct {
	CreateMessageResponse *CreateMessageResponse

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type CreateThreadRequest struct {
	// Flag if a thread is sticky - sticky thread has an immutable member list, members cannot be added or removed.
// Sticky threads are only supported for 1-1 chat, i.e. with only two members.
	IsStickyThread *bool `json:"isStickyThread,omitempty"`

	// Members to be added to the thread.
	Members *[]ThreadMember `json:"members,omitempty"`

	// The thread topic.
	Topic *string `json:"topic,omitempty"`
}

type CreateThreadResponse struct {
	// Thread Id.
	ID *string `json:"id,omitempty"`
}

// CreateThreadResponseResponse is the response envelope for operations that return a CreateThreadResponse type.
type CreateThreadResponseResponse struct {
	CreateThreadResponse *CreateThreadResponse

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type GetMessagesResponse struct {
	// If there are more messages that can be retrieved, the backward link will be populated.
	BackwardLink *string `json:"backwardLink,omitempty"`

	// List of messages.
	Messages *[]Message `json:"messages,omitempty"`

	// Continuation link to get new and modified messages.
	SyncState *string `json:"syncState,omitempty"`
}

// GetMessagesResponseResponse is the response envelope for operations that return a GetMessagesResponse type.
type GetMessagesResponseResponse struct {
	GetMessagesResponse *GetMessagesResponse

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type GetThreadsResponse struct {
	// If there are more threads that can be retrieved, the backward link will be populated.
	BackwardLink *string `json:"backwardLink,omitempty"`

	// Continuation link to get new and modified threads.
	SyncState *string `json:"syncState,omitempty"`

	// List of threads.
	Threads *[]ThreadInfo `json:"threads,omitempty"`
}

// GetThreadsResponseResponse is the response envelope for operations that return a GetThreadsResponse type.
type GetThreadsResponseResponse struct {
	GetThreadsResponse *GetThreadsResponse

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type Message struct {
	// The client message Id specified when the message was sent.
// This Id is a client-specific Id in a numeric unsigned Int64 format. It can be used for client deduping, among other client
// usages.
	ClientMessageID *string `json:"clientMessageId,omitempty"`

	// Content of the message.
	Content *string `json:"content,omitempty"`

	// The timestamp when the message was deleted in Unix time (epoch time) in milliseconds.
	DeleteTime *int64 `json:"deleteTime,omitempty"`

	// The timestamp when the message was edited in Unix time (epoch time) in milliseconds.
	EditTime *int64 `json:"editTime,omitempty"`

	// The Id of the message sender.
	From *string `json:"from,omitempty"`

	// The id of the message. This id is server generated.
	ID *string `json:"id,omitempty"`

	// Type of the message.
//
// Possible values:
// - Text
// - ThreadActivity/TopicUpdate
// - ThreadActivity/AddMember
// - ThreadActivity/DeleteMember
	MessageType *string `json:"messageType,omitempty"`

	// The timestamp when the message arrived at the server. The timestamp is in ISO8601 format: `yyyy-MM-ddTHH:mm:ssZ`.
	OriginalArrivalTime *time.Time `json:"originalArrivalTime,omitempty"`
	Priority *MessagePriority `json:"priority,omitempty"`

	// The display name of the message sender. This property is used to populate sender name for push notifications.
	SenderDisplayName *string `json:"senderDisplayName,omitempty"`

	// Version of the message.
	Version *string `json:"version,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Message.
func (m Message) MarshalJSON() ([]byte, error) {
	type alias Message
	aux := &struct {
		*alias
		OriginalArrivalTime *timeRFC3339 `json:"originalArrivalTime"`
	}{
		alias: (*alias)(&m),
		OriginalArrivalTime: (*timeRFC3339)(m.OriginalArrivalTime),
	}
	return json.Marshal(aux)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Message.
func (m *Message) UnmarshalJSON(data []byte) error {
	type alias Message
	aux := &struct {
		*alias
		OriginalArrivalTime *timeRFC3339 `json:"originalArrivalTime"`
	}{
		alias: (*alias)(m),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	m.OriginalArrivalTime = (*time.Time)(aux.OriginalArrivalTime)
	return nil
}

// MessageResponse is the response envelope for operations that return a Message type.
type MessageResponse struct {
	Message *Message

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type PostReadReceiptRequest struct {
	// The client message Id specified when the message was sent.
// This Id is a client-specific Id in a numeric unsigned Int64 format. It can be used for client deduping, among other client
// usages.
	ClientMessageID *string `json:"clientMessageId,omitempty"`

	// Id of the latest message read by current user.
	MessageID *string `json:"messageId,omitempty"`
}

type ReadReceipt struct {
	// Client message id specified in the Microsoft.AzureCommunicationService.Gateway.Models.Client.CreateMessageRequest.
// This Id is a client-specific Id in a numeric unsigned Int64 format. It can be used for client deduping, among other client
// usages.
	ClientMessageID *string `json:"clientMessageId,omitempty"`

	// Id for the message that has been read. This id is server generated.
	MessageID *string `json:"messageId,omitempty"`

	// Read receipt timestamp.
	ReadTimestamp *int64 `json:"readTimestamp,omitempty"`

	// Read receipt sender id.
	UserID *string `json:"userId,omitempty"`
}

// ReadReceiptArrayResponse is the response envelope for operations that return a []ReadReceipt type.
type ReadReceiptArrayResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Array of ReadReceipt
	ReadReceiptArray *[]ReadReceipt
}

type Thread struct {
	// Thread creation time in Unix time (epoch time) in milliseconds.
	CreatedAt *string `json:"createdAt,omitempty"`

	// Id of the thread owner.
	CreatedBy *string `json:"createdBy,omitempty"`

	// Thread id.
	ID *string `json:"id,omitempty"`

	// Flag if a thread is sticky - sticky thread has an immutable member list, members cannot be added or removed.
	IsStickyThread *bool `json:"isStickyThread,omitempty"`

	// Thread members.
	Members *[]ThreadMember `json:"members,omitempty"`

	// Thread topic.
	Topic *string `json:"topic,omitempty"`
}

type ThreadInfo struct {
	// Thread id.
	ID *string `json:"id,omitempty"`

	// Flag if a thread is soft deleted.
	IsDeleted *bool `json:"isDeleted,omitempty"`

	// The timestamp when the last message arrived at the server. The timestamp is in ISO8601 format: `yyyy-MM-ddTHH:mm:ssZ`.
	LastMessageReceivedTime *time.Time `json:"lastMessageReceivedTime,omitempty"`

	// Thread topic.
	Topic *string `json:"topic,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ThreadInfo.
func (t ThreadInfo) MarshalJSON() ([]byte, error) {
	type alias ThreadInfo
	aux := &struct {
		*alias
		LastMessageReceivedTime *timeRFC3339 `json:"lastMessageReceivedTime"`
	}{
		alias: (*alias)(&t),
		LastMessageReceivedTime: (*timeRFC3339)(t.LastMessageReceivedTime),
	}
	return json.Marshal(aux)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ThreadInfo.
func (t *ThreadInfo) UnmarshalJSON(data []byte) error {
	type alias ThreadInfo
	aux := &struct {
		*alias
		LastMessageReceivedTime *timeRFC3339 `json:"lastMessageReceivedTime"`
	}{
		alias: (*alias)(t),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	t.LastMessageReceivedTime = (*time.Time)(aux.LastMessageReceivedTime)
	return nil
}

// A member of the thread.
type ThreadMember struct {
	// Display name for the thread member.
	DisplayName *string `json:"displayName,omitempty"`

	// The id of the thread member in the format `8:acs:ResourceId_AcsUserId`.
	ID *string `json:"id,omitempty"`

	// Role of the thread member. The valid value should be "User" or "Admin".
	MemberRole *MemberRole `json:"memberRole,omitempty"`

	// Time from which the group chat history is shared with the member in EPOCH time (milliseconds).
	ShareHistoryTime *string `json:"shareHistoryTime,omitempty"`
}

// ThreadMemberArrayResponse is the response envelope for operations that return a []ThreadMember type.
type ThreadMemberArrayResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Array of ThreadMember
	ThreadMemberArray *[]ThreadMember
}

// ThreadResponse is the response envelope for operations that return a Thread type.
type ThreadResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
	Thread *Thread
}

type UpdateMessageRequest struct {
	// Chat message content.
	Content *string `json:"content,omitempty"`
}

type UpdateThreadRequest struct {
	// Thread topic.
	Topic *string `json:"topic,omitempty"`
}

