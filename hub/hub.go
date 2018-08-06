package hub

type MessageRelay interface {
	Id() int
	Send([]byte) error
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	clients    map[int]MessageRelay // Registered clients.
	broadcast  chan []byte          // Inbound messages from the clients.
	register   chan MessageRelay    // Register requests from the clients.
	unregister chan MessageRelay    // Unregister requests from clients.
}

// NewHub creates an un-started Hub without any clients registered.
func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan MessageRelay),
		unregister: make(chan MessageRelay),
		clients:    make(map[int]MessageRelay),
	}
}

// Register saves the given client in order to broadcast messages to all
// clients registered.
func (h *Hub) Register(c *Client) {
	h.register <- c
}

func (h *Hub) Unregister(c *Client) {
	h.unregister <- c
}

// Run starts the hub and handles registering and de-registering clients.
func (h *Hub) Run() {
	go func() {
		for {
			select {
			case client := <-h.register:
				h.clients[client.Id()] = client

			case client := <-h.unregister:
				if client == nil {
					continue
				}

				if _, ok := h.clients[client.Id()]; ok {
					delete(h.clients, client.Id())
				}

			case message := <-h.broadcast:
				var removes []int

				for id, client := range h.clients {
					err := client.Send(message)
					if err != nil {
						removes = append(removes, id)
					}
				}

				if len(removes) > 0 {
					for _, id := range removes {
						delete(h.clients, id)
					}
				}

			}
		}
	}()
}
