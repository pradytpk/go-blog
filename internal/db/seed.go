package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	"github.com/pradytpk/go-blog/internal/store"
)

var usernames = []string{
	"john_doe", "jane_smith", "alice_wonder", "bob_builder", "charlie_brown",
	"danny_ocean", "eva_green", "frank_castle", "george_clooney", "harry_potter",
	"ian_fleming", "jack_sparrow", "kate_winslet", "luke_skywalker", "mary_jane",
	"nancy_drew", "oliver_queen", "peter_parker", "quentin_tarantino", "rachel_green",
	"sarah_connor", "tom_hanks", "ursula_sea", "victor_frankenstein", "wade_wilson",
	"xander_cage", "yara_greyjoy", "zach_galifianakis", "amber_heard", "brad_pitt",
	"chris_evans", "diana_prince", "edward_snowden", "fiona_apple", "george_rr_martin",
	"helen_keller", "indy_jones", "jerry_seinfeld", "kim_possible", "leonardo_dicaprio",
	"mike_tyson", "nina_simone", "oprah_winfrey", "paul_walker", "quentin_blake",
	"robin_williams", "selena_gomez", "tina_fey", "uma_thurman", "vin_diesel",
}

var titles = []string{
	"Master of Shadows", "The Silent Guardian", "Echoes of the Past", "Whispers in the Wind",
	"The Lost Kingdom", "Rise of the Phoenix", "The Forgotten Realms", "Tales of the Ancients",
	"The Final Frontier", "Beyond the Horizon", "Legends of the Night", "Secrets of the Deep",
	"Journey to the Unknown", "The Dark Prophecy", "Warriors of the Light", "The Eternal Quest",
	"Songs of the Stars", "The Hidden Treasure", "Battle for the Throne", "Mysteries of the Universe",
}

var content = []string{
	"Adventure begins when courage and bravery lead you through unknown lands.",
	"In the galaxy, stars shine brighter for those who explore fearlessly.",
	"Every journey starts with a single step towards a great discovery.",
	"Legends are born from stories of mystery, night, and hidden truths.",
	"The ocean holds secrets of power, waiting to be unveiled by heroes.",
	"Whispers of the past echo through time, revealing treasures long lost.",
	"Victory belongs to those who rise above challenges and face destiny head-on.",
	"With wisdom and valor, you can conquer kingdoms and fulfill your dreams.",
	"Hope fuels the quest for freedom, glory, and an eternal legacy.",
	"In the battle between light and darkness, harmony restores the balance.",
	"Destiny calls to those who seek the truth hidden in shadows.",
	"A hero's quest is never complete without bravery, wisdom, and hope combined.",
	"The universe expands for those who imagine endless possibilities and wonder.",
	"Power comes from within, guiding the brave on their journey forward.",
	"Strength lies in unity, where courage and honor forge unbreakable bonds.",
	"The night sky tells stories of galaxies far beyond our understanding.",
	"In every legend, there is a journey full of adventure and risk.",
	"Discoveries lead to enlightenment, where knowledge and courage pave the way.",
	"The rise of heroes brings light to dark corners of the world.",
	"Beyond the horizon, endless opportunities await those who dare to dream.",
}

var tags = []string{
	"adventure", "brave", "create", "discover", "explore",
	"future", "galaxy", "hero", "imagine", "journey",
	"kingdom", "legend", "mystery", "night", "ocean",
	"power", "quest", "rise", "shadow", "treasure",
	"universe", "victory", "whisper", "xenon", "yonder",
	"zeal", "wisdom", "courage", "honor", "destiny",
	"valor", "truth", "strength", "hope", "freedom",
	"glory", "eternal", "infinity", "light", "darkness",
	"harmony", "serenity", "chaos", "balance", "quest",
}

var commentsText = []string{
	"Great job on the project! Really impressive work.",
	"I found this very helpful. Thanks for sharing!",
	"This needs a bit more explanation. Could you elaborate?",
	"Excellent insights. Looking forward to more content like this.",
	"I had a similar issue, and your solution worked perfectly.",
	"The design could be improved for better user experience.",
	"Nice work! I especially liked the part about data handling.",
	"Could you provide more examples to illustrate this concept?",
	"I disagree with the approach. Have you considered an alternative?",
	"The performance improvements are noticeable. Well done!",
	"The code is clean but could benefit from more comments.",
	"Great explanation of the algorithm. It’s much clearer now.",
	"I appreciate the effort, but the documentation needs updates.",
	"This is a solid foundation, but more testing is needed.",
	"The implementation is good, but consider edge cases.",
	"The UI looks good, but there are some accessibility issues.",
	"I would suggest adding more unit tests for robustness.",
	"This feature is very useful. Thanks for implementing it!",
	"I’m experiencing some bugs. Could you look into it?",
	"Overall, a good contribution. Minor tweaks could make it better.",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()
	users := generateUsers(100)

	tx, _ := db.BeginTx(ctx, nil)
	for _, user := range users {
		if err := store.UsersIF.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()
			log.Println("Error Creating user:", err)
			return
		}
	}
	tx.Commit()
	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.PostsIF.Create(ctx, post); err != nil {
			_ = tx.Rollback()
			log.Println("Error Creating user:", err)
			return
		}
	}
	tx.Commit()

	comments := generateComment(500, users, posts)
	for _, comment := range comments {
		if err := store.CommentsIF.Create(ctx, comment); err != nil {
			_ = tx.Rollback()
			log.Println("Error Creating user:", err)
			return
		}
	}
	tx.Commit()

	log.Println("Seedding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)
	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			RoleID:   1,
		}
	}
	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: content[rand.Intn(len(content))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}
	return posts
}

func generateComment(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	comments := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		comments[i] = &store.Comment{
			UserID:  users[rand.Intn(len(users))].ID,
			PostID:  posts[rand.Intn(len(posts))].ID,
			Content: commentsText[rand.Intn(len(commentsText))],
		}
	}
	return comments
}
