package main

import "time"

type Config struct {
	URL      string
	Domain   string
	Username string
	Password string
	From     string
	To       []string
	Subject  string
	Body     string
	Timeout  time.Duration
}

// Web services reference
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createitem-operation-email-message

type CreateItem struct {
	XMLName            struct{}          `xml:"m:CreateItem"`
	MessageDisposition string            `xml:"MessageDisposition,attr"`
	SavedItemFolderId  SavedItemFolderId `xml:"m:SavedItemFolderId"`
	Items              Messages          `xml:"m:Items"`
}

type Messages struct {
	Message []Message `xml:"t:Message"`
}

type SavedItemFolderId struct {
	DistinguishedFolderId DistinguishedFolderId `xml:"t:DistinguishedFolderId"`
}

type DistinguishedFolderId struct {
	Id string `xml:"Id,attr"`
}

type Message struct {
	ItemClass    string       `xml:"t:ItemClass"`
	Subject      string       `xml:"t:Subject"`
	Body         Body         `xml:"t:Body"`
	Sender       Sender       `xml:"t:Sender"`
	ToRecipients MailboxArray `xml:"t:ToRecipients"`
}

type Body struct {
	BodyType string `xml:"BodyType,attr"`
	Body     []byte `xml:",chardata"`
}

type Sender struct {
	Mailbox Mailbox `xml:"t:Mailbox"`
}

type MailboxArray struct {
	Mailbox []Mailbox `xml:"t:Mailbox"`
}

type Mailbox struct {
	EmailAddress string `xml:"t:EmailAddress"`
}
