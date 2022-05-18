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

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
