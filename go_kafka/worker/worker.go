package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2/log"
)

const (
	Topic          = "push-notifications"
	Partition      = int32(0)
	DefaultOffset  = sarama.OffsetOldest
	FCMCredentials = "./json/mation-test-firebase-adminsdk.json"
)

var (
	BrokersURLs = []string{"localhost:29092"}
)

func main() {

	// connect kafka
	worker, err := connectConsumer(BrokersURLs)
	if err != nil {
		log.Fatalf("Error connecting to Kafka: %v", err)
	}
	defer worker.Close()

	// offset := int64(10) // i want to get consumer.HighWaterMarkOffset()?
	// offset := sarama.OffsetOldest // Always start consuming from the beginning
	offset, err := getHighWaterMarkOffset(Topic, Partition)
	if err != nil {
		log.Fatalf("Error getting high water mark offset: %v", err)
	}

	consumer, err := worker.ConsumePartition(Topic, Partition, offset)
	if err != nil {
		log.Fatalf("Error consuming partition: %v", err)
	}

	fmt.Println("consumer started")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	msgCount := 0
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCount++
				// fmt.Printf("Received message count %d | topic (%s) | message(%s)\n", msgCount, string(msg.Topic), string(msg.Value))
				fmt.Printf("Received message count %d | topic (%s)\n", msgCount, string(msg.Topic))

				// 받은 메시지를 파싱하여 푸시 알림을 전송하는 로직 추가
				var notification AppPush
				err := json.Unmarshal(msg.Value, &notification)
				if err != nil {
					fmt.Println("Error parsing notification:", err)
					continue
				}

				// 여기서 푸시 알림을 전송하는 로직 구현
				if notification.LinkType != "" && notification.NotiDB != nil {
					getFcmRepo().MultiPushNotification(notification.LinkType, notification.NotiDB, nil)
				}

			case <-sigchan:
				fmt.Println("Interruption detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("processed", msgCount, "messges")

}

func connectConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func getHighWaterMarkOffset(topic string, partition int32) (int64, error) {
	// Create a Kafka client
	client, err := sarama.NewClient(BrokersURLs, nil)
	if err != nil {
		return 0, err
	}
	defer client.Close()

	// Get the high water mark offset
	offset, err := client.GetOffset(topic, partition, sarama.OffsetNewest)
	if err != nil {
		return 0, err
	}

	return offset, nil
}

// =========================================

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

var (
	fcmAppRepo *FcmAppRepo
)

type PushLink string

const (
	LINK_ACTIVITY PushLink = "activity" // 활동알림
	LINK_NEWPOST  PushLink = "newpost"  // 새글알림
	LINK_CHAT     PushLink = "chat"     // 채팅
)

type FcmAppRepo struct {
	App *firebase.App
}

func repoInterface(app *firebase.App) *FcmAppRepo {
	return &FcmAppRepo{App: app}
}

// https://firebase.google.com/docs/cloud-messaging/server?hl=ko
func initFcmApp() (*firebase.App, error) {
	opt := option.WithCredentialsFile("./json/mation-test-firebase-adminsdk.json")
	pushApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return pushApp, nil
}

func getFcmRepo() *FcmAppRepo {
	if fcmAppRepo == nil {
		fcmApp, err := initFcmApp()
		if err != nil {
			log.Error(err.Error())
			return nil // 오류 발생 시 nil 반환
		}
		fcmAppRepo = repoInterface(fcmApp)
	}
	return fcmAppRepo
}

// 앱 푸쉬 보내기 (여러명 보내기)
func (fcmApp *FcmAppRepo) MultiPushNotification(linkType PushLink, notiDB *Noti, userList []string) error {
	// fmt.Printf("fcmM [시작] 유저숫자: %d\n", len(userList))

	startTime := time.Now()

	FCMList := MockUpFCMList
	fmt.Printf("fcmM [시작] 유저숫자: %d\n", len(FCMList))

	var pushMessageList = make([]*messaging.Message, 0, len(FCMList))

	for _, fcm := range FCMList {
		messageData := map[string]string{
			"link":         string(linkType),
			"type":         notiDB.ArticleType,
			"articleNo":    fmt.Sprint(notiDB.ArticleNo),
			"use_move_url": strconv.FormatBool(notiDB.UseMoveURL),
			// "id":           notiDB.ID.Hex(),
			// "move_url":     notiDB.MoveURL,
		}
		message := CreateFCMMessaging(fcm.FCMToken, notiDB.NotiMessage, messageData)

		pushMessageList = append(pushMessageList, message)
	}
	if len(pushMessageList) == 0 {
		return nil
	}

	successCount, failureCount, err := fcmApp.SendAllPushMessage(pushMessageList)
	if err != nil {
		log.Error(err)
	}

	endTime := time.Now()
	fmt.Printf("fcm [파이널성공] 성공: %d / 실패: %d / 총알림: %d / 총걸린시간: %v\n", successCount, failureCount, len(pushMessageList), endTime.Sub(startTime))
	return nil
}

// fcm 받을 데이터 정리
func CreateFCMMessaging(userToken, bodyMessage string, messageData map[string]string) *messaging.Message {
	return &messaging.Message{
		Token: userToken,
		Notification: &messaging.Notification{
			Title: "레사모(Re4mo)",
			Body:  bodyMessage,
		},
		Data: messageData,
		Android: &messaging.AndroidConfig{
			Notification: &messaging.AndroidNotification{
				Tag: "brace.formation",
			},
		},
		APNS: &messaging.APNSConfig{
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					MutableContent: true,
					Alert: &messaging.ApsAlert{
						Title: "레사모(Re4mo)",
						Body:  bodyMessage,
					},
					Sound: "default",
				},
			},
		},
	}
}

// 여러 기기에 알림 보내기
func (fcmApp *FcmAppRepo) SendAllPushMessage(pushMessageList []*messaging.Message) (int, int, error) {

	if len(pushMessageList) == 0 {
		return 0, 0, nil
	}

	ctx := context.Background()
	client, err := fcmApp.App.Messaging(ctx)
	if err != nil {
		log.Error("푸시 전송 실패 : ", err)
		fmt.Printf("fcm [sendAll 실패1] 에러 : %s\n", err.Error())
	}

	res, err := client.SendAll(ctx, pushMessageList)
	if err != nil {
		log.Error("푸시 전송 실패 : ", err) // 일반적으로 토큰이 만료된 경우, 여기에 들어오고 푸시 실패하고 끝남
		fmt.Printf("fcm [sendAll 실패2] 에러 : %s\n", err.Error())
	}

	return res.SuccessCount, res.FailureCount, nil
}

var MockUpFCMList = []FcmToken{
	{
		FCMToken: "d1tiCYD4QKaS1AbJN4SQU4:APA91bE6GlA9iajEUkbYC9QYl2i23nA5v6srT35xJYbR9r3BVeqSoJOQ-bX3g_qMYTfUyCkCf_Pdo2qpLxCnRONo5r7qXJ5DSF37og23359G4vHrj0QU4E54j8fR8sQWIy9s5BZ_WDFP",
		Username: "naver",
	},
	{
		FCMToken: "ePFvAD85T5CSI6HH6vpEXS:APA91bGRuWUgu0mLPF7ty2RFCVhHKI4ECeNLhy4shHSqSCKJqOmocbA7XgoDLKDdI-sCOyRXm2JAFblrCN1_upV4ZwxcTzCC9ZYSMKkay2LzJ-YYvAS55JSX4IHBWHVjOzbvncvD8qff",
		Username: "jexists",
	},
	{
		Username: "jexists",
		FCMToken: "ePFvAD85T5CSI6HH6vpEXS:APA91bGRuWUgu0mLPF7ty2RFCVhHKI4ECeNLhy4shHSqSCKJqOmocbA7XgoDLKDdI-sCOyRXm2JAFblrCN1_upV4ZwxcTzCC9ZYSMKkay2LzJ-YYvAS55JSX4IHBWHVjOzbvncvD8qff",
	},
	{
		Username: "naver",
		FCMToken: "ePFvAD85T5CSI6HH6vpEXS:APA91bGRuWUgu0mLPF7ty2RFCVhHKI4ECeNLhy4shHSqSCKJqOmocbA7XgoDLKDdI-sCOyRXm2JAFblrCN1_upV4ZwxcTzCC9ZYSMKkay2LzJ-YYvAS55JSX4IHBWHVjOzbvncvD8qff",
	},
	{
		FCMToken: "eR5R-4fSQkCNeij01K0SEk:APA91bFQ6SaD4f9rh-VMsw6VS6vWXemtdGGJQKd1novG9iqrytWefE1yI4iBMsBwXyCsXXU3scNJUuVLrrCVEuA10fC0RmEiMdJH4NUxjtNr7MBQPCWFiZgG2OSVRpyRULVfXA1XXp9h",
		Username: "jexists",
	},
}
