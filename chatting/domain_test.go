package chatting

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Add_Session(t *testing.T) {
	d := Domain{ repo: make(MockRepo) }
	nixTime := time.Now().Unix()
	expected := Session{
		Id: "first-session",
		UserId: 2,
		RoomId: 3,
		LastPostedAt: nixTime,
		EnteredAt: nixTime,
	}

	d.AddSession(expected)
	sessions := d.Sessions()

	assert.Equal(t, 1, len(sessions))

	actual := sessions[expected.Id]
	assert.NotNil(t, actual)
	assert.Equal(t, expected.Id, actual.Id)
}

func Test_Add_Post(t *testing.T) {
	d := Domain{ repo: make(MockRepo) }
	expected := Post{
		Text: "I'm batman",
		UserId: 1,
	}
	d.AddPost(1, expected)
	roomToPosts := d.Posts()
	posts, ok := roomToPosts[1]

	assert.True(t, ok)
	assert.Equal(t, 1, len(posts))
}

func Test_Get_Sessions(t *testing.T) {
	d := Domain{ repo: make(MockRepo) }
	sessions := make(map[string]Session)
	nixTime := time.Now().Unix()
	expected := Session{
		Id: "room-1",
		UserId: 2,
		RoomId: 3,
		LastPostedAt: nixTime,
		EnteredAt: nixTime,
	}
	sessions[expected.Id] = expected

	d.repo.Set("sessions", sessions)

	set := d.Sessions()
	assert.NotNil(t, set)

	actual, ok := set[expected.Id]
	assert.True(t, ok)
	assert.Equal(t, expected.Id, actual.Id)
	assert.Equal(t, expected.UserId, actual.UserId)
	assert.Equal(t, expected.RoomId, actual.RoomId)
	assert.Equal(t, expected.LastPostedAt, actual.LastPostedAt)
	assert.Equal(t, expected.EnteredAt, actual.EnteredAt)
}

func Test_Rooms(t *testing.T) {
	d := Domain{ repo: make(MockRepo) }
	rooms := make(map[int]Room)
	expected := Room{
		Occupants: []int{ 1 },
		Id: 1,
	}
	rooms[expected.Id] = expected

	d.repo.Set("rooms", rooms)

	set := d.Rooms()
	assert.NotNil(t, set)

	actual, ok := set[1]
	assert.True(t, ok)
	assert.Equal(t, expected.Id, actual.Id)
}

func Test_Domain_Users_From_Empty_Repo(t *testing.T) {
	d := Domain{ repo: make(MockRepo) }
	assert.NotNil(t, d.Users())
}

func Test_Domain_Rooms_From_Empty_Repo(t *testing.T) {
	d := Domain{ repo: make(MockRepo) }
	assert.NotNil(t, d.Rooms())
}

func Test_Domain_Posts_From_Empty_Repo(t *testing.T) {
	d := Domain{ repo: make(MockRepo) }
	assert.NotNil(t, d.Posts())
}

func Test_Domain_Sessions_From_Empty_Repo(t *testing.T) {
	d := Domain{ repo: make(MockRepo) }
	assert.NotNil(t, d.Posts())
}
