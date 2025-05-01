package main

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"time"

// 	"./mockup"
// 	"./model"

// 	log "github.com/sirupsen/logrus"

// 	firebase "firebase.google.com/go"

// 	"firebase.google.com/go/messaging"
// 	"google.golang.org/api/option"
// )

// var (
// 	fcmAppRepo *FcmAppRepo
// )

// type PushLink string

// const (
// 	LINK_ACTIVITY PushLink = "activity" // 활동알림
// 	LINK_NEWPOST  PushLink = "newpost"  // 새글알림
// 	LINK_CHAT     PushLink = "chat"     // 채팅
// )

// type FcmAppRepo struct {
// 	App *firebase.App
// }

// func repoInterface(app *firebase.App) *FcmAppRepo {
// 	return &FcmAppRepo{App: app}
// }

// // https://firebase.google.com/docs/cloud-messaging/server?hl=ko
// func initFcmApp() (*firebase.App, error) {
// 	fcmEnvInfo := os.Getenv("FcmEnvironmentFilename")
// 	opt := option.WithCredentialsFile("./json/" + fcmEnvInfo)
// 	pushApp, err := firebase.NewApp(context.Background(), nil, opt)
// 	if err != nil {
// 		log.Error(err)
// 		return nil, fmt.Errorf("error initializing app: %v", err)
// 	}
// 	return pushApp, nil
// }

// func getFcmRepo() *FcmAppRepo {
// 	if fcmAppRepo == nil {
// 		fcmApp, err := initFcmApp()
// 		if err != nil {
// 			log.Error(err.Error())
// 		}
// 		fcmAppRepo = repoInterface(fcmApp)
// 	}
// 	return fcmAppRepo
// }

// // 앱 푸쉬 보내기 (여러명 보내기)
// func (fcmApp *FcmAppRepo) MultiPushNotification(linkType PushLink, notiDB *model.Noti, userList []string) error {
// 	// fmt.Printf("fcmM [시작] 유저숫자: %d\n", len(userList))

// 	startTime := time.Now()
// 	// FCMList := GetMultiFCMTokenList(userList)
// 	FCMList := mockup.MockUpFCMList
// 	fmt.Printf("fcmM [시작] 유저숫자: %d\n", len(FCMList))

// 	var pushMessageList = make([]*messaging.Message, 0, len(FCMList))

// 	for _, fcm := range FCMList {
// 		messageData := map[string]string{
// 			"link":         string(linkType),
// 			"type":         notiDB.ArticleType,
// 			"articleNo":    fmt.Sprint(notiDB.ArticleNo),
// 			"id":           notiDB.ID.Hex(),
// 			"use_move_url": strconv.FormatBool(notiDB.UseMoveURL),
// 			"move_url":     notiDB.MoveURL,
// 		}
// 		message := CreateFCMMessaging(fcm.FCMToken, notiDB.NotiMessage, messageData)

// 		pushMessageList = append(pushMessageList, message)
// 	}
// 	if len(pushMessageList) == 0 {
// 		return nil
// 	}

// 	successCount, failureCount, err := fcmApp.SendAllPushMessage(pushMessageList)
// 	if err != nil {
// 		log.Error(err)
// 	}

// 	endTime := time.Now()
// 	fmt.Printf("fcm [파이널성공] 성공: %d / 실패: %d / 총알림: %d / 총걸린시간: %v\n", successCount, failureCount, len(pushMessageList), endTime.Sub(startTime))
// 	return nil
// }

// // fcm 받을 데이터 정리
// func CreateFCMMessaging(userToken, bodyMessage string, messageData map[string]string) *messaging.Message {
// 	return &messaging.Message{
// 		Token: userToken,
// 		Notification: &messaging.Notification{
// 			Title: "레사모(Re4mo)",
// 			Body:  bodyMessage,
// 		},
// 		Data: messageData,
// 		Android: &messaging.AndroidConfig{
// 			Notification: &messaging.AndroidNotification{
// 				Tag: "brace.formation",
// 			},
// 		},
// 		APNS: &messaging.APNSConfig{
// 			Payload: &messaging.APNSPayload{
// 				Aps: &messaging.Aps{
// 					MutableContent: true,
// 					Alert: &messaging.ApsAlert{
// 						Title: "레사모(Re4mo)",
// 						Body:  bodyMessage,
// 					},
// 					Sound: "default",
// 				},
// 			},
// 		},
// 	}
// }

// // 여러 기기에 알림 보내기
// func (fcmApp *FcmAppRepo) SendAllPushMessage(pushMessageList []*messaging.Message) (int, int, error) {

// 	if len(pushMessageList) == 0 {
// 		return 0, 0, nil
// 	}

// 	ctx := context.Background()
// 	client, err := fcmApp.App.Messaging(ctx)
// 	if err != nil {
// 		log.Error("푸시 전송 실패 : ", err)
// 		fmt.Printf("fcm [sendAll 실패1] 에러 : %s\n", err.Error())
// 	}

// 	res, err := client.SendAll(ctx, pushMessageList)
// 	if err != nil {
// 		log.Error("푸시 전송 실패 : ", err) // 일반적으로 토큰이 만료된 경우, 여기에 들어오고 푸시 실패하고 끝남
// 		fmt.Printf("fcm [sendAll 실패2] 에러 : %s\n", err.Error())
// 	}

// 	return res.SuccessCount, res.FailureCount, nil
// }
