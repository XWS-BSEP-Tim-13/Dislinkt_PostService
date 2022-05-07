package startup

import (
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var posts = []*domain.Post{
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f82"),
		Date:     time.Time{},
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
				Date:     time.Time{},
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
