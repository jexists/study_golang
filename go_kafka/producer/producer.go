package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
)

type Comment struct {
	Text string `form:"text" json:"text"`
}

func main() {
	app := fiber.New()
	api := app.Group("/api/v1")
	api.Post("/comments", createComment)
	app.Listen(":3000")
}

func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func PushCommentToQueue(topic string, message []byte) error {
	brokersUrl := []string{"localhost:29092"}
	producer, err := ConnectProducer(brokersUrl)
	if err != nil {
		return err
	}
	defer producer.Close()
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}
	fmt.Printf("message is sotre in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}

func createComment(c *fiber.Ctx) error {
	cmt := new(AppPush)
	if err := c.BodyParser(cmt); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}
	cmtInBytes, err := json.Marshal(cmt)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating notification.",
		})
		return err
	}

	PushCommentToQueue("push-notifications", cmtInBytes)

	err = c.JSON(&fiber.Map{
		"success": true,
		"message": "notification pushed sucessfully.",
		"comment": cmt,
	})

	return err
}

// ============
type PushLink string

const (
	LINK_ACTIVITY PushLink = "activity" // 활동알림
	LINK_NEWPOST  PushLink = "newpost"  // 새글알림
	LINK_CHAT     PushLink = "chat"     // 채팅
)

type AppPush struct {
	LinkType PushLink `form:"link_type" json:"link_type"`
	NotiDB   *Noti    `form:"noti_db" json:"noti_db"`
}

type FcmToken struct {
	Username string `bson:"username" json:"username"`
	FCMToken string `bson:"fcm_token" json:"fcm_token"`
}

// Noti is 알림DB
type Noti struct {
	ArticleType string `json:"article_type,omitempty" bson:"article_type,omitempty"` // moveURL 만들때 사용 (talk/market/등)
	ArticleNo   int    `json:"article_no,omitempty" bson:"article_no,omitempty"`
	UseMoveURL  bool   `json:"use_move_url" bson:"use_move_url"` // true: 앱에서 moveURL을 사용 (webview)
	NotiMessage string `json:"noti_message" bson:"noti_message"` // 알림 내용
	// ID                primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	// NotiUsername      string             `json:"noti_username" bson:"noti_username"`       // 알림 받는사람
	// RelatedUsername   string             `json:"related_username" bson:"related_username"` // 알림 전송하는 사람
	// NotiTime          time.Time          `json:"noti_time" bson:"noti_time"`
	// NotiRead          bool               `json:"noti_read" bson:"noti_read"`
	// MoveURL           string             `json:"move_url,omitempty" bson:"move_url,omitempty"`
	// NotiType          string             `json:"noti_type" bson:"noti_type"`                           //comment, offer, counteroffer, club, contact, report, keyword
	// ArticleID         primitive.ObjectID `json:"article_id,omitempty" bson:"article_id,omitempty"`
	// ArticleUsername   string             `json:"article_username,omitempty" bson:"article_username,omitempty"`       // 첫댓글인 경우 글쓴이에게 알림
	// ArticleTitle      string             `json:"article_title,omitempty" bson:"article_title,omitempty"`             // 새글알림
	// ArticleMainImage  string             `json:"article_mainimage,omitempty" bson:"article_mainimage,omitempty"`     // 새글알림
	// CmtParentUsername string             `json:"cmt_parent_username,omitempty" bson:"cmt_parent_username,omitempty"` // 대댓글인 경우 부모댓글쓴이에게 알림
	// CmtID             primitive.ObjectID `json:"cmt_id,omitempty" bson:"cmt_id,omitempty"`                           // 댓글인 경우 댓글 ID
	// CmtContent        string             `json:"cmt_content,omitempty" bson:"cmt_content,omitempty"`
	// OfferID           primitive.ObjectID `json:"offer_id,omitempty" bson:"offer_id,omitempty"`
	// NotiImage         string             `json:"noti_image" bson:"noti_image"`
	// NotiTitle         string             `json:"noti_title" bson:"noti_title"` // 닉네임 / 클럽이름 / 기타
	// UsePopUp          bool               `json:"use_popup" bson:"use_popup"`
	// BrdNo             int                `json:"brd_no,omitempty" bson:"brd_no,omitempty"` // 클럽 사용
}
