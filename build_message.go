package main

import "encoding/xml"

func buildMessage(cfg *Config) ([]byte, error) {
	mailboxes := make([]Mailbox, len(cfg.To))
	for i, addr := range cfg.To {
		mailboxes[i].EmailAddress = addr
	}

	createItem := &CreateItem{
		MessageDisposition: "SendAndSaveCopy",
		SavedItemFolderId: SavedItemFolderId{
			DistinguishedFolderId: DistinguishedFolderId{
				Id: "sentitems",
			},
		},
		Items: Messages{
			Message: []Message{
				{
					ItemClass: "IPM.Note",
					Subject:   cfg.Subject,
					Body: Body{
						BodyType: "Text",
						Body:     []byte(cfg.Body),
					},
					Sender: Sender{
						Mailbox: Mailbox{
							EmailAddress: cfg.From,
						},
					},
					ToRecipients: MailboxArray{
						mailboxes,
					},
				},
			},
		},
	}

	return xml.MarshalIndent(createItem, "", "  ")
}
