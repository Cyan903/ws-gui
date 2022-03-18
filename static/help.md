<!-- 
    Fyne rich text has bad spacing. The only way I've been able to get proper 
    spacing is by using empty #####.
-->

## Request
To make a request, click on the `request` tab and fill in the entry with a proper websocket url.  

##### Example:
- `ws://example.com`
- `wss://example.com` (for ssl only)

Fill in the request info entry and hit `send` to send a message.
#####
## Response
The responses are logged into a simple multi line entry. The format for messages is:
```
[MSG] for a received message
[YOU] for a message sent by you
[FATAL] when a connection error occurs
\> ... for connection info
```
The message history can be cleared in settings.
