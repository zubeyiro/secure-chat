package events

const (
	PING                 = "ping" // Server
	PONG                 = "pong" // user
	NEW_USER_CONNECTED   = "user_connected"
	USER_DISCONNECTED    = "user_disconnected"
	USER_INFO_REQUESTED  = "user_info_requested"
	USER_INFO_RECEIVED   = "user_info_received"
	USER_LIST_REQUESTED  = "user_list_requested"
	USER_LIST_RECEIVED   = "user_list_received"
	NEW_MESSAGE_SENT     = "new_message_sent"
	NEW_MESSAGE_RECEIVED = "new_message_received"
)

/*
- Server
	- Events
		- server_started
			- create public and private keys
		- user_connected
			- Add to users list
			- Update all other users
		- new_message
			- Relay message
		- user_list_request
			- Send list to requester
		- user_disconnected
- Client
	- Events
		- Incoming
			- connected
				- save servers public key
			- user_list_received
			- message_received
		- Outgoing
			- user_list_request
			- send_message
- Encryption/Decryption
	- client <> server
	- client <> client
*/
