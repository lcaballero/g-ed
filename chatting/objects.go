package chatting

type Session struct {
	Id           string   `json:"id"`
	UserId       int   `json:"user"`
	RoomId       int   `json:"room"`
	LastPostedAt int64 `json:"lastPostAt"`
	EnteredAt    int64 `json:"enteredAt"`
}

type User struct {
	// Id (internal) so users can change their usernames
	Id       int    `json:"-"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Room struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Occupants []int `json:"occupants"`
}

type Post struct {
	Text   string
	UserId int
}
