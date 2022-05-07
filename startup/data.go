package startup

import (
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var posts = []*domain.Post{
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f82"),
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
				Date:     "",
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
