package chatting

const (
	keySessions = "sessions"
	keyUsers = "users"
	keyPosts = "posts"
	keyRooms = "rooms"
)

type Domain struct {
	repo Repo
}

// Users returns a copy of UserID to User map.
func (d Domain) Users() map[int]User {
	return make(map[int]User)
}

func (d Domain) ReplaceSession(old string, session Session) error {
	current := d.Sessions()
	delete(current, old)
	current[session.Id] = session
	d.repo.Set(keySessions, current)
	return nil
}

// AddSession contains the session key and the other information
func (d Domain) AddSession(session Session) error {
	current := d.Sessions()
	key := session.Id
	current[key] = session
	d.repo.Set(keySessions, current)
	return nil
}

// AddPost updates the posts in the room
func (d Domain) AddPost(roomId int, post Post) error {
	roomPosts := d.Posts()
	posts, ok := roomPosts[roomId]
	if !ok {
		posts = make([]Post, 0)
	}
	posts = append(posts, post)
	roomPosts[roomId] = posts
	return d.repo.Set(keyPosts, roomPosts)
}

// Rooms returns a copy of RoomID to Room map.
func (d Domain) Rooms() map[int]Room {
	vals, ok := d.repo.Get(keyRooms).(map[int]Room)
	rooms := make(map[int]Room)
	if !ok {
		return rooms
	}
	for k,room := range vals {
		rooms[k] = room
	}
	return rooms
}

func (d Domain) AddRoom(room Room) error {
	current := d.Rooms()
	current[room.Id] = room
	return d.repo.Set(keyRooms, current)
}

// Posts returns a copy of RoomID to []Post map.
func (d Domain) Posts() map[int][]Post {
	vals, ok := d.repo.Get(keyPosts).(map[int][]Post)
	roomPosts := make(map[int][]Post)
	if !ok {
		return roomPosts
	}

	for k,posts := range vals {
		clone := make([]Post, len(posts))
		copy(clone, posts)
		roomPosts[k] = clone
	}

	return roomPosts
}

// Sessions returns a copy of SessionKey to Session map.
func (d Domain) Sessions() map[string]Session {
	vals, ok := d.repo.Get(keySessions).(map[string]Session)
	sessions := make(map[string]Session)
	if !ok {
		return sessions
	}
	for k,session := range vals {
		sessions[k] = session
	}
	return sessions
}
