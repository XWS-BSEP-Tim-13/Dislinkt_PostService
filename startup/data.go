package startup

import (
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var posts = []*domain.Post{
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f82"),
		Date:     time.Now(),
		Username: "anagavrilovic",
		Content:  "Mrzim Go!",
		Image:    "3f74e912-6d37-4aef-92e8-3981d5bc9a23",
		Likes: []string{
			"srdjansukovic",
			"stefanljubovic",
			"anagavrilovic",
		},
		Dislikes: []string{},
		Comments: []domain.Comment{
			{
				Id:       getObjectId("623b0cc3a34d25d8567f9f83"),
				Content:  "I ja isto!",
				Date:     time.Now(),
				Username: "srdjansukovic",
			},
		},
	},
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f83"),
		Date:     time.Now(),
		Username: "srdjansukovic",
		Content:  "Mrzim Go!",
		Image:    "",
		Likes: []string{
			"srdjansukovic",
			"stefanljubovic",
		},
		Dislikes: []string{},
		Comments: []domain.Comment{
			{
				Id:       getObjectId("623b0cc3a34d25d8567f9f83"),
				Content:  "I ja isto!",
				Date:     time.Now(),
				Username: "srdjansukovic",
			},
		},
	},
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f84"),
		Date:     time.Now(),
		Username: "marijakljestan",
		Content:  "Mrzim Go!",
		Image:    "",
		Likes: []string{
			"srdjansukovic",
		},
		Dislikes: []string{},
		Comments: []domain.Comment{
			{
				Id:       getObjectId("623b0cc3a34d25d8567f9f83"),
				Content:  "I ja isto!",
				Date:     time.Now(),
				Username: "srdjansukovic",
			},
		},
	},
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f85"),
		Date:     time.Now(),
		Username: "marijakljestan",
		Content:  "Mrzim Go!",
		Image:    "",
		Likes: []string{
			"srdjansukovic",
			"stefanljubovic",
		},
		Dislikes: []string{},
		Comments: []domain.Comment{
			{
				Id:       getObjectId("623b0cc3a34d25d8567f9f83"),
				Content:  "I ja isto!",
				Date:     time.Now(),
				Username: "srdjansukovic",
			},
		},
	},
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f86"),
		Date:     time.Now(),
		Username: "anagavrilovic",
		Content:  "Mrzim Go!",
		Image:    "",
		Likes: []string{
			"srdjansukovic",
			"stefanljubovic",
		},
		Dislikes: []string{},
		Comments: []domain.Comment{
			{
				Id:       getObjectId("623b0cc3a34d25d8567f9f83"),
				Content:  "I ja isto!",
				Date:     time.Now(),
				Username: "srdjansukovic",
			},
		},
	},
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f87"),
		Date:     time.Now(),
		Username: "anagavrilovic",
		Content:  "Mrzim Go!",
		Image:    "",
		Likes: []string{
			"srdjansukovic",
			"stefanljubovic",
		},
		Dislikes: []string{},
		Comments: []domain.Comment{
			{
				Id:       getObjectId("623b0cc3a34d25d8567f9f83"),
				Content:  "I ja isto!",
				Date:     time.Now(),
				Username: "srdjansukovic",
			},
		},
	},
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f88"),
		Date:     time.Now(),
		Username: "anagavrilovic",
		Content:  "Mrzim Go!",
		Image:    "",
		Likes: []string{
			"srdjansukovic",
			"stefanljubovic",
		},
		Dislikes: []string{},
		Comments: []domain.Comment{
			{
				Id:       getObjectId("623b0cc3a34d25d8567f9f83"),
				Content:  "I ja isto!",
				Date:     time.Now(),
				Username: "srdjansukovic",
			},
		},
	},
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f89"),
		Date:     time.Now(),
		Username: "lenka",
		Content:  "Mrzim Go!",
		Image:    "",
		Likes: []string{
			"srdjansukovic",
			"stefanljubovic",
		},
		Dislikes: []string{},
		Comments: []domain.Comment{
			{
				Id:       getObjectId("623b0cc3a34d25d8567f9f83"),
				Content:  "I ja isto!",
				Date:     time.Now(),
				Username: "srdjansukovic",
			},
		},
	},
}

var messages = []*domain.MessageUsers{
	{
		Id:         getObjectId("623b0cc3a34d25d8567f9f83"),
		FirstUser:  "srdjansukovic",
		SecondUser: "stefanljubovic",
		Messages: []domain.Message{
			{
				Id:          getObjectId("623b0cc3a34d25d8567f9f93"),
				Date:        time.Now(),
				MessageTo:   "srdjansukovic",
				MessageFrom: "stefanljubovic",
				Content:     "Lorem ipsum lores..",
			},
			{
				Id:          getObjectId("623b0cc3a34d25d8567f9f94"),
				Date:        time.Now().Add(time.Hour),
				MessageTo:   "srdjansukovic",
				MessageFrom: "stefanljubovic",
				Content:     "Lorem ipsum lores lorem..",
			},
			{
				Id:          getObjectId("623b0cc3a34d25d8567f9f95"),
				Date:        time.Now().Add(time.Hour * 2),
				MessageTo:   "srdjansukovic",
				MessageFrom: "stefanljubovic",
				Content:     "Lorem ipsum lores lorem lor..",
			},
			{
				Id:          getObjectId("623b0cc3a34d25d8567f9f96"),
				Date:        time.Now().Add(time.Hour),
				MessageTo:   "stefanljubovic",
				MessageFrom: "srdjansukovic",
				Content:     "Lorem ipsum lores lorem lorem ipsum lorem..",
			},
			{
				Id:          getObjectId("623b0cc3a34d25d8567f9f97"),
				Date:        time.Now().Add(time.Hour),
				MessageTo:   "stefanljubovic",
				MessageFrom: "srdjansukovic",
				Content:     "Lorem ipsum lores lorem lorem ipsum saffsafsa..",
			},
		},
	},
	{
		Id:         getObjectId("623b0cc3a34d25d8567f9f84"),
		FirstUser:  "lenka",
		SecondUser: "stefanljubovic",
		Messages: []domain.Message{
			{
				Id:          getObjectId("623b0cc3a34d25d8567f9f98"),
				Date:        time.Now(),
				MessageTo:   "lenka",
				MessageFrom: "stefanljubovic",
				Content:     "Lorem ipsum lores..",
			},
		},
	},
}

var events = []*domain.Event{
	{
		Id:        getObjectId("623b0cc3a34d25d8567f9f92"),
		Action:    `Created post "Mrzim go!"`,
		User:      "stefanljubovic",
		Published: time.Now().Add(24 * time.Hour),
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
