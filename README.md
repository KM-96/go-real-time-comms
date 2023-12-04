# Go lang: Real time communications

The project aims at exploring different options for enabling real time communication between a client and server. 
The items mentioned in the topics section includes the different options that will be explored.

Application
There are 2 services as part of the application: a Client and a Server
- Client: A service that expects real time updates from a Server. The client will employ different strategies discussed in the topics section to get data from the Server.
- Server: A service that takes an input from the console of the user and is responsible to communicate the data to the client.

## Topics

### Polling

#### Short Polling
Sending repeated HTTP requests to the server at fixed time intervals.

##### Advantages:

- Simple implementation, as it leverages standard HTTP requests.
- Works across various browsers and platforms.

##### Disadvantages:

- Creates a significant amount of unnecessary network traffic, even when there are no updates.
- Real-time updates may experience delays depending on the polling interval.
- Puts an additional load on the server due to frequent requests.

Appropriate Use Case: Suitable for scenarios where data is expected to update every time a request is made, such as train or taxi location tracking.

#### Long Polling
Sending long-lived HTTP requests where the server does not respond unless there is an update or the resource becomes available.

##### Advantages

- Enables real-time updates with reduced latency compared to short polling.
- Reduces unnecessary network traffic by keeping the connection open until an update occurs.

##### Disadvantages

- Increased server load compared to traditional polling, as connections remain open for longer durations.
- Requires more complex server-side implementation to handle long-lived connections.
- Appropriate Use Case: Ideal for situations where data updates are infrequent, such as chat apps or push notifications.

### Web Sockets


### Webhooks

### gRPC 

### SSE (Server Sent Events)

## References
- 