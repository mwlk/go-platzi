package main

import "fmt"
type INotificationFactory interface{
	SendNotification()
	GetSender() ISender
}

type ISender interface{
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct{

}

func (SMSNotification) SendNotification(){
	fmt.Println("sending via SMS")
}

type SMSNotificationSender struct{

}

func (SMSNotificationSender) GetSenderMethod() string{
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string{
	return "Twilio"
}

func (SMSNotification) GetSender() ISender{
	return SMSNotificationSender{}
}

type EmailNotification struct{

}

func (EmailNotification) SendNotification(){
	fmt.Println("Sending via Email")
}

type EmailNotificationSender struct{

}

func (EmailNotificationSender) GetSenderMethod() string{
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string{
	return "SMTP"
}

func (EmailNotification) GetSender() ISender{
	return EmailNotificationSender{}
}

func GetNotificationFactory(notificationType string)(INotificationFactory, error){
if notificationType == "SMS"{
	return &SMSNotification{}, nil
}

if notificationType == "Email"{
	return &EmailNotification{}, nil
}

return nil, fmt.Errorf("No Notification Type")
} 

func sendNotification(f INotificationFactory){
	f.SendNotification()
}

func getMethod(f INotificationFactory){
	 fmt.Println(f.GetSender().GetSenderMethod())
}

func main(){
smsFactory, _ := GetNotificationFactory("SMS")
emailFactory, _ := GetNotificationFactory("Email")

sendNotification(smsFactory)
sendNotification(emailFactory)

getMethod(smsFactory)
getMethod(emailFactory)
}